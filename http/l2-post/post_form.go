package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

// COMMAND
// curl -X POST -d "name=bonfy&salary=90" http://httpbin.org/post

// {
// 	"args": {},
// 	"data": "",
// 	"files": {},
// 	"form": {
// 	  "name": "bonfy",
// 	  "salary": "90"
// 	},
// 	"headers": {
// 	  "Accept": "*/*",
// 	  "Connection": "close",
// 	  "Content-Length": "20",
// 	  "Content-Type": "application/x-www-form-urlencoded",
// 	  "Host": "httpbin.org",
// 	  "User-Agent": "curl/7.54.0"
// 	},
// 	"json": null,
// 	"origin": "xxx,xxx,xx,xx",
// 	"url": "http://httpbin.org/post"
//   }

// Use Post
func postFormSample1() {
	target := "http://httpbin.org/post"
	// 也可以 先用 url.Values构建 然后 用 Encode 变成这种方式
	params := "name=bonfy&salary=90"
	contentType := "application/x-www-form-urlencoded"
	res, err := http.Post(target, contentType, strings.NewReader(params))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println(res)

	jsonBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	json := string(jsonBytes)
	name := gjson.Get(json, "form.name")
	fmt.Println("Form:", name)
	headers := gjson.Get(json, "headers")
	fmt.Println("HEADER:", headers)
}

// Use PostForm
func postFormSample2() {
	target := "http://httpbin.org/post"
	// params := "name=bonfy&salary=90"
	// contentType := "application/x-www-form-urlencoded"

	v := url.Values{}
	v.Set("name", "bonfy")
	v.Add("salary", "90")
	res, err := http.PostForm(target, v)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// fmt.Println(res)

	fmt.Println("Status:", res.Status)
	fmt.Println("StatusCode:", res.StatusCode)
	fmt.Println("Proto:", res.Proto)
	fmt.Println("ProtoMajor:", res.ProtoMajor)
	fmt.Println("ProtoMinor:", res.ProtoMinor)
	fmt.Println("Header:", res.Header)
	fmt.Println("Body:", res.Body)
	fmt.Println("ContentLength:", res.ContentLength)
	fmt.Println("TransferEncoding:", res.TransferEncoding)
	fmt.Println("Close:", res.Close)
	fmt.Println("Uncompressed:", res.Uncompressed)
	fmt.Println("Trailer:", res.Trailer)
	fmt.Println("Request:", res.Request)
	fmt.Println("TLS:", res.TLS)

	jsonBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	json := string(jsonBytes)
	name := gjson.Get(json, "form.name")
	fmt.Println("Form:", name)
	headers := gjson.Get(json, "headers")
	fmt.Println("HEADER:", headers)
}

func main() {
	postFormSample1()

	postFormSample2()
}
