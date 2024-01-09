package main

import (
	"fmt"
	"time"
)

// import "main/synchronous"

// step2
// 주고받을 자료 string
func isGood(person string, c chan string) {
	time.Sleep(time.Second * 2)
	// c <- true //send true to channel a.k.a. c
	c <- person + " is great!"
}

// step1
func main() {
	// synchronous.Run()

	//gorountine은 프로그램이 작동할동안만 live.(즉 메인함수와 함께 꺼짐)
	//기다려주지 않음.

	//Channel은 고루틴간, 고루틴-메인 간 커뮤니케이션 가능
	//NOTE: 어떤 종류의 자료를 주고받을 것인가 string
	c := make(chan string)

	people := [5]string{"jesper", "flynn", "pan", "juan", "kim"}

	//TODO: 2개의 고루틴 생성
	for _, person := range people {

		//NOTE: 고루틴은 go 키워드
		//고루틴이니 result := go isGood(person) 이런 거 불가.. 받기 전에 끝날 수도
		//고로 저 채널을 isGood으로 보내서 메인 함수와 연계.
		go isGood(person, c)
	}

	for i := 0; i < len(people); i++ {
		//NOTE: step 3. await. blocking operation임 채널로부터 한 메시지를 받을 동안 await. 메인(go runtime)이 멈춰 기다림.
		//다만 병렬적으로 도니까 2초 뒤에 다 주루룩 나옴 ㅎㅎ. Concurrency로 인해 누가 먼저 튀어나오는 지는 모름
		fmt.Println(<-c)
	}


	//걍 위 방식대로 하셈. 아래는 기존 코드 참고.
	// //NOTE: await. blocking operation임 채널로부터 한 메시지를 받을 동안 await. go runtime이 멈춤.
	// result := <-c //isGood이 보내준 걸 consume. receive from channel a.k.a. c


	// fmt.Println(result)
	// fmt.Println(<-c) //남은 한 고루틴 메시지 받기. 둘 다 거의 동시에 끝남.
	// // fmt.Println(<-c) //all gourintes are asleep - deadlock ERROR. 고루틴 2개 돌렸음.. 받을 게 없는데 대기.. 안 끝남. 고로 에러 내줌. 그러니 그냥 저렇게 고루틴 갯수만큼 도는 for loop에 넣자.
	

}
