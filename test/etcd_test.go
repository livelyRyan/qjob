package test

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestEtcd(t *testing.T) {
	config := clientv3.Config{
		Endpoints:   []string{"192.168.44.14:2379"},
		DialTimeout: time.Second * 3,
	}

	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancelFunc()
	status, err := client.Status(ctx, config.Endpoints[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(status)

	_, err = client.KV.Put(ctx, "k", "v")
	if err != nil {
		fmt.Println(err)
	}
}
