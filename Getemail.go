package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	QQemail = `(\d+)@qq.com`
)

func GetEamil() {
	resq, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	ErrorHander(err, "get error")
	defer resq.Body.Close()
	pageByte, err := ioutil.ReadAll(resq.Body)
	ErrorHander(err, "read error")
	pageStr := string(pageByte)
	re := regexp.MustCompile(QQemail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("QQ:", result[1])
	}
}

func main() {
	GetEamil()
}

func ErrorHander(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}
