package main

import (
	"errors"
	"fmt"
	"net/http"
)

type reqResult struct {
	url string
	status string  //int든 자유임
}


var errRequestFailed = errors.New("request failed")


//c chan result의 대안 중 하나. 
//c chan<- result  : 이 채널은 send only 기능만 시키겠다.
//참고: c <- result{~}: 채널로 보내기
//        <- c: 채널로부터 받기

func hitURL(url string, c chan<- reqResult) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400{
		status = "FAILED"  //재활용이니 =써라. := 말고.
	}
	c <- reqResult{url: url, status: status}
	

	return nil
}


func main() {


	results := map[string]string{}

	//위에서 정의한 result 타입를 송수신
	c:= make(chan reqResult)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
		}







	for _, url := range urls {
		go hitURL(url, c)
	}

	for i:=0; i < len(urls); i++ {
		// fmt.Println(<-c) // 한번에 다 나옴. 가장 오래 걸린 사이트만큼만 시간 소요됨.
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

}