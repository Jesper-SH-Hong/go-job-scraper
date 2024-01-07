package synchronous

import (
	"errors"
	"fmt"
	"net/http"
)


var errRequestFailed = errors.New("request failed")

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400{
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}


func Run() {

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


	results := map[string]string{}  //이렇게 {}로 초기화 가능. 
	//혹은 var results = map[string]string{},
	//혹은 var results = make(map[string]string).
	//저렇게들 해줘야 map이 nil이 아니라서 값 추가가 가능함.






	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}



}