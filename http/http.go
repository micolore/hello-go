// http
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unsafe"
)

func main() {
	fmt.Println("Hello  Http!")
	var base_url = "http://localhost:23333/around-api"
	http_post_json(base_url+"/content/index", "page=1")

	// song := make(map[string]string)
	// song["name"] = "李白"
	// song["timelength"] = "128"
	// song["author"] = "李荣浩"
	// samplePost(song)
}

//http get
func http_get() {
	response, err := http.Get("http://localhost:23333/around-api")
	if err != nil {
		fmt.Println("error ....")
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("请求成功，返回数据: %s\n", string(body))
}

//http post
func http_post_json(url string, params string) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		// handle error
	}
	//common
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	//bussiness
	req.Header.Set("android_id", "")
	req.Header.Set("android_version", "")
	req.Header.Set("build_num", "")
	req.Header.Set("channel", "")
	req.Header.Set("country", "")
	req.Header.Set("google_account", "")
	req.Header.Set("imei", "")
	req.Header.Set("imsi", "")
	req.Header.Set("imsi2", "")
	req.Header.Set("ip", "")
	req.Header.Set("longitude", "")
	req.Header.Set("latitude", "")
	req.Header.Set("version_name", "")
	req.Header.Set("version_code", "")
	req.Header.Set("serial", "")
	req.Header.Set("resolution", "")
	req.Header.Set("phone_model", "")
	req.Header.Set("net_work", "")
	req.Header.Set("mac", "")
	req.Header.Set("gpReferrer", "")
	req.Header.Set("deviceId", "")
	req.Header.Set("timestamp", "")
	req.Header.Set("sign", "")
	req.Header.Set("user_id", "")
	req.Header.Set("token", "")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

type JsonPostSample struct {
}

func samplePost(params map[string]string) {

	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "http://localhost:23333/around-api"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}
