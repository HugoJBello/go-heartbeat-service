
package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HugoJBello/go-heartbeat-service/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetActivity(t *testing.T) {

	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/activity/last?limit=1&skip=0", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	bodyGetResp, err := ioutil.ReadAll(resp.Body)
	var activity models.ServiceActivityResponse

	assert.Nil(t, err)

	json.Unmarshal(bodyGetResp, &activity)

	assert.NotNil(t, activity)
	assert.Equal(t, len(activity.Data), 1)

	fmt.Println(activity.Data)

}


func TestCreateActivity(t *testing.T) {

	body := models.CreateServiceActivity{ServiceId: "test",ActivityContent: "doing things", ServiceName:"test", ActivityContentType:"test",Date: time.Now()}
	jsonBody, _ := json.Marshal(body)

	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.Post(fmt.Sprintf("%s/activity/new", ts.URL), "application/json",  bytes.NewBuffer(jsonBody))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	bodyPostResp, err := ioutil.ReadAll(resp.Body)
	var activity models.ServiceActivityResponse

	assert.Nil(t, err)

	json.Unmarshal(bodyPostResp, &activity)

	assert.NotNil(t, activity)
	fmt.Println(activity.Data)


}