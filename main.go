package main

import (
	"fmt"
	"time"
)

// import "main/synchronous"

// step2
// 주고받을 자료 bool
func isGood(person string, c chan bool) {
	time.Sleep(time.Second * 2)
	c <- true //send true to channel a.k.a. c
}

// step1
func main() {
	// synchronous.Run()

	//gorountine은 프로그램이 작동할동안만 live.(즉 메인함수와 함께 꺼짐)
	//기다려주지 않음.
	//Channel은 고루틴간, 고루틴-메인 간 커뮤니케이션 가능

	//NOTE: 어떤 종류의 자료를 주고받을 것인가 bool
	c := make(chan bool)

	people := [2]string{"jesper", "flynn"}

	//2개의 고루틴 생성
	for _, person := range people {
		//고루틴이니 result := go isGood(person) 이런 거 불가.. 받기 전에 끝날 수도
		//고로 저 채널을 isGood으로 보내서 메인 함수와 연계.
		go isGood(person, c)
	}

	//NOTE: await. blocking operation임 채널로부터 한 메시지를 받을 동안 await. go runtime이 멈춤.
	result := <-c //isGood이 보내준 걸 consume. receive from channel a.k.a. c


	fmt.Println(result)
	fmt.Println(<-c) //남은 한 고루틴 메시지 받기. 둘 다 거의 동시에 끝남.
	// fmt.Println(<-c) //all gourintes are asleep - deadlock ERROR. 고루틴 2개 돌렸음..
	

}
