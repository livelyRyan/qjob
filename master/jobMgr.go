package master

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"qjob/common"
	"time"
)

type JobMgr struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

var GJobMgr *JobMgr

func InitJobMgr() error {
	// 形成配置
	config := clientv3.Config{
		Endpoints:   Config.EtcdEndpoints,
		DialTimeout: time.Duration(Config.EtcdDialTimeout) * time.Millisecond,
	}

	// 建立连接
	client, err := clientv3.New(config)
	if err != nil {
		return err
	}

	// 探测 client 连通性
	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancelFunc()
	_, err = client.Status(ctx, config.Endpoints[0])
	if err != nil {
		panic(err)
	}

	GJobMgr = &JobMgr{
		client: client,
		kv:     client.KV,
		lease:  client.Lease,
	}
	return nil
}

func (jobMgr *JobMgr) SaveJob(job *common.Job) (*common.Job, error) {
	key := common.ETCD_JOB_DIR + job.Name
	bs, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}

	oldJobBs, err := jobMgr.kv.Put(context.TODO(), key, string(bs), clientv3.WithPrevKV())
	if err != nil {
		return nil, err
	}
	if oldJobBs.PrevKv != nil {
		var oldJob common.Job
		err = json.Unmarshal(oldJobBs.PrevKv.Value, &oldJob)
		if err != nil {
			// 不影响返回值
			fmt.Println(err)
			return nil, nil
		}
		return &oldJob, nil
	}
	return nil, nil
}

func (jobMgr *JobMgr) KillJob(jobName string) error {
	// 创建一个 1 秒过期时间的 lease
	leaseResp, err := GJobMgr.lease.Grant(context.TODO(), 1)
	if err != nil {
		return err
	}
	key := common.ETCD_JOB_KILL_DIR + jobName

	// 在 etcd 中存放 key，过期时间是为了节省 etcd 中的空间
	_, err = GJobMgr.kv.Put(context.TODO(), key, "", clientv3.WithLease(leaseResp.ID))
	return err
}
