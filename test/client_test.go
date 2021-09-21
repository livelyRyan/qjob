package test

import (
	"encoding/json"
	"fmt"
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

	resp, err := http.Post("http://localhost:9999/jobs/save", "application/json", strings.NewReader(string(bs)))
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
