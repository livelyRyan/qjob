package master

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"qjob/common"
	"strconv"
	"time"
)

type ApiServer struct {
	httpServer *http.Server
}

var GApiServer *ApiServer

// InitApiServer 初始化服务
func InitApiServer() error {
	if GApiServer != nil {
		return nil
	}
	// 创建路由
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/jobs/save", saveJob)
	GApiServer = &ApiServer{
		httpServer: &http.Server{
			ReadTimeout:  time.Duration(Config.ApiReadTimeout) * time.Millisecond,
			WriteTimeout: time.Duration(Config.ApiWriteTimeout) * time.Millisecond,
			Handler:      serveMux,
		},
	}

	// 监听端口
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(Config.ApiPort))
	if err != nil {
		return err
	}

	// 协程启动服务
	go func() {
		err := GApiServer.httpServer.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()

	return err
}

// 保存 job
func saveJob(resp http.ResponseWriter, req *http.Request) {
	var (
		err  error
		body []byte
		oldJob *common.Job
	)
	body, err = ioutil.ReadAll(req.Body)
	defer func(er error, resp http.ResponseWriter) {
		if er != nil {
			bs, er := common.BuildResponse(-1, er.Error(), nil)
			if er != nil {
				if _, er = resp.Write(bs); er != nil {
					fmt.Println(er)
				}
			}
		}
	}(err, resp)
	if err != nil {
		return
	}

	var job common.Job
	err = json.Unmarshal(body, &job)
	if err != nil {
		return
	}
	// 保存到 ETCD 中
	oldJob, err = GJobMgr.SaveJob(&job)
	if err != nil {
		return
	}
	// 响应
	bytes, _ := common.BuildResponse(0, "Success", oldJob)
	_, err = resp.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}
