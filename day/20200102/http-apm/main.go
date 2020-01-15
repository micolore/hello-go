package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	go func() {
		for {
			log.Println(Add("hello......"))
		}
	}()
	http.ListenAndServe("localhost:8090", nil)
}

var datas []string

func Add(str string) string {

	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)
	return sData
}
