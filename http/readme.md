# HTTP 模块

> `import "net/http"`

## 请求

最好不要 用 url 作为变量名，因为 golang 又 "net/url" 包

* 请求虽然有 `Get`、`Post`、`PUT`、`HEAD` 之分，但是返回都是 Response 的指针
* 请求返回后，养成 defer resp.Body.Close() 的好习惯

```Go
type Response struct {
        Status     string // e.g. "200 OK"
        StatusCode int    // e.g. 200
        Proto      string // e.g. "HTTP/1.0"
        ProtoMajor int    // e.g. 1
        ProtoMinor int    // e.g. 0
        Header Header

        Body io.ReadCloser

        ContentLength int64

        TransferEncoding []string

        Close bool

        Uncompressed bool

        Trailer Header

        Request *Request

        TLS *tls.ConnectionState
}
```

### Get

`func Get(url string) (resp *Response, err error)`

```Golang
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
```

* Get Plain Text
* Get Json

json -> Struct 两种方案

*  body, _ := ioutil.ReadAll(res.Body) -> json.Unmarshal(body, &obj)
*  decoder := json.NewDecoder(res.Body) -> err := decoder.Decode(&obj)

### Post

* `func Post(url string, contentType string, body io.Reader) (resp *Response, err error)`
* `func PostForm(url string, data url.Values) (resp *Response, err error)`

四种形式，这里主要介绍前两种

* application/x-www-form-urlencoded (默认，不提供contentType)
* application/json 
* multipart/form-data
* text/xml （XML RPC）


### Client

这个最全面，也能完成最多的需求

```Golang
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```

* 可以设置Header
* context


