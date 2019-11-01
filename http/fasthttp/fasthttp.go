// fasthttp
package main

import (
	//"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("Hello fasthttp!")

	//fasthttp_get()

	url := "http://192.168.1.229:23333/around-api"

	fasthttp_post_json(url)
}

func fasthttp_get() {
	url := "http://localhost:23333/around-api"

	status, resp, err := fasthttp.Get(nil, url)

	if err != nil {
		fmt.Println("request error", err.Error())
		return
	}
	if status != fasthttp.StatusOK {
		fmt.Println("request no success", status)
	}
	fmt.Println(string(resp))
}

func fasthttp_post() {

	url := "http://192.168.1.229:23333/around-api"

	args := &fasthttp.Args{}
	args.Add("name", "l")
	args.Add("age", "19")

	status, resp, err := fasthttp.Post(nil, url, args)

	if err != nil {
		fmt.Println("request error", err.Error())
		return
	}
	if status != fasthttp.StatusOK {
		fmt.Println("request no success", status)
	}
	fmt.Println(string(resp))
}

//转发好像有点问题，还有，fh是用来构建服务器的而不只是单独的发送http请求
func fasthttp_post_json(url string) {

	req := &fasthttp.Request{}

	fmt.Println(url)
	req.SetRequestURI(url)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod(fasthttp.MethodGet)

	resp := &fasthttp.Response{}

	jsonStr := "1ssd" //"{\"name\":\"li\"}"
	//json, err := json.Marshal(jsonStr)
	req.SetBody([]byte(jsonStr))
	//fmt.Println("fasthttp_post_json error", err.Error())

	client := &fasthttp.Client{}

	err := client.Do(req, resp)

	b := resp.Body()
	if err != nil {
		err.Error()
	}
	fmt.Println("result:\r\n", string(b))

}
