//爬虫
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error :status code ", resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)

	//将GBK编码解码为UTF8
	uft8Reader := transform.NewReader(resp.Body,
		e.NewDecoder())

	all, err := ioutil.ReadAll(uft8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

//通过1024字节判定数据编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
