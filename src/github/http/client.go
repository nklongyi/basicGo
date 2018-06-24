package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

const url string = "http://www.baidu.com"
func main()  {
	request,err := http.NewRequest(http.MethodGet,url,nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")
	//检查重定向
	client :=http.Client{
		CheckRedirect:func(req *http.Request, via []*http.Request) error{
			fmt.Println("Redirect:",req)
			return nil
		},
	}

	resp,err :=client.Do(request)
	//resp,err :=http.Get()
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	s,err := httputil.DumpResponse(resp,true)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",s)
}
