package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// http://httpbin.org/get

func getSample() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println(resp)

	fmt.Println("Status:", resp.Status)
	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println("Proto:", resp.Proto)
	fmt.Println("ProtoMajor:", resp.ProtoMajor)
	fmt.Println("ProtoMinor:", resp.ProtoMinor)
	fmt.Println("Header:", resp.Header)
	fmt.Println("Body:", resp.Body)
	fmt.Println("ContentLength:", resp.ContentLength)
	fmt.Println("TransferEncoding:", resp.TransferEncoding)
	fmt.Println("Close:", resp.Close)
	fmt.Println("Uncompressed:", resp.Uncompressed)
	fmt.Println("Trailer:", resp.Trailer)
	fmt.Println("Request:", resp.Request)
	fmt.Println("TLS:", resp.TLS)
}

func getTextSample() {
	res, err := http.Get("https://www.baidu.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

type OriginObject struct {
	Origin string `json:"origin"`
}

func getJsonSample() {
	res, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var obj OriginObject

	// decode 方案一

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// err = json.Unmarshal(body, &obj)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// decode 方案二
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&obj)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", obj.Origin)
}

func main() {
	getSample()
	fmt.Println("============")
	getTextSample()
	fmt.Println("============")
	getJsonSample()
}
