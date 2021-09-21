package test

import (
	"encoding/json"
	"fmt"
	restful "github.com/emicklei/go-restful"
	"io/ioutil"
	"net/http"
	"qjob/common"
	"strings"
	"testing"
)

func TestAddJob(t *testing.T)  {
	job := common.Job{
		Name:    "test",
		Command: "date",
		CornExp: "5 * * * *",
	}
	bs, err := json.Marshal(job)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:9999/jobs/save", restful.MIME_JSON, strings.NewReader(string(bs)))
	if err != nil {
		fmt.Println(err)
		return
	}
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response common.Response
	err = json.Unmarshal(bs, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}

// /jobs/kill?jobName=a
func TestKillJob(t *testing.T) {
	resp, err := http.Post("http://localhost:9999/jobs/kill?jobName=a", restful.MIME_JSON, nil)
	if err != nil {
		panic(err)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response common.Response
	err = json.Unmarshal(bs, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}