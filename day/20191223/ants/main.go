package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/panjf2000/ants"
)

type Request struct {
	Param  []byte
	Result chan []byte
}

func main() {
	fmt.Println("ants...")
	pool, _ := ants.NewPoolWithFunc(10000, func(patload interface{}) {

		request, ok := patload.(*Request)
		if !ok {
			return
		}

		//定义一个匿名函数处理请求的body
		reversParam := func(s []byte) []byte {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			return s
		}(request.Param)

		//把处理过的数据设置到result字段
		request.Result <- reversParam
	})

	defer pool.Release()

	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello ants")
		param, err := ioutil.ReadAll(r.Body)
		//fmt.Println("params", param)
		if err != nil {
			http.Error(w, "request error", http.StatusInternalServerError)
		}
		defer r.Body.Close()
		request := &Request{Param: param, Result: make(chan []byte)}
		if err := pool.Invoke(request); err != nil {
			http.Error(w, "throttle limit error", http.StatusInternalServerError)
		}
		w.Write(<-request.Result)
	})
	http.ListenAndServe("127.0.0.1:8080", nil)

}
