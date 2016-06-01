package vkapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	API_METHOD_URL      = "https://api.vk.com/method/"
	AUTH_HOST           = "https://oauth.vk.com/authorize"
	AUTH_HOST_GET_TOKEN = "https://oauth.vk.com/access_token"
)

type Api struct {
	AccessToken string
	UserId      int
	ExpiresIn   int
	debug       bool
}

func (vk Api) Request(method_name string, params map[string]string) string {
	u, err := url.Parse(API_METHOD_URL + method_name)
	if err != nil {
		fmt.Print(err)
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	q.Set("access_token", vk.AccessToken)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}

	return string(content)
}

func GetResponse(m string, parametr string) interface{} {
	data := []byte(m)
	var parsed interface{}
	err := json.Unmarshal(data, &parsed)
	if err != nil {
		fmt.Print(err)
	}
	par, _ := parsed.(map[string]interface{})
	var va interface{}
	for _, v := range par {
		va = v
	}
	return (append([]interface{}{}, va))[0].([]interface{})[1].(map[string]interface{})[parametr]

}
