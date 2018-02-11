package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// COMMAND
// curl -X POST -H "Content-Type:application/json" --data '{"name":"bonfy","salary":90}' http://httpbin.org/post

// {
// 	"args": {},
// 	"data": "{\"name\":\"bonfy\",\"salary\":90}",
// 	"files": {},
// 	"form": {},
// 	"headers": {
// 	  "Accept": "*/*",
// 	  "Connection": "close",
// 	  "Content-Length": "28",
// 	  "Content-Type": "application/json",
// 	  "Host": "httpbin.org",
// 	  "User-Agent": "curl/7.54.0"
// 	},
// 	"json": {
// 	  "name": "bonfy",
// 	  "salary": 90
// 	},
// 	"origin": "xxx.xxx.xx.xx",
// 	"url": "http://httpbin.org/post"
//   }

// Use Post
func postJson() {
	target := "http://httpbin.org/post"
	// 也可以 先用 url.Values构建 然后 用 Encode 变成这种方式
	jsonStr := []byte(`{"name":"bonfy","salary":90}`)
	contentType := "application/json"
	res, err := http.Post(target, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	fmt.Println("resonse Headers:", res.Header)

	jsonBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json := string(jsonBytes)
	fmt.Println("response Json:", json)
}

// Use Post and post data params
// 这样是没有 json的
func postJsonWrongWay() {
	target := "http://httpbin.org/post"
	params := "name=bonfy&salary=90"
	contentType := "application/json"
	res, err := http.Post(target, contentType, strings.NewReader(params))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	fmt.Println("resonse Headers:", res.Header)

	jsonBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json := string(jsonBytes)
	fmt.Println("response Json:", json)
}

func main() {
	postJson()
}
