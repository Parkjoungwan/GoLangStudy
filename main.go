package main

import "fmt"

func main() {
	names := []string{"nico", "lynn", "dal"}
	//slices는 length가 없는 배열 (?Vector와의 차이점은?)
	names = append(names, "what?") //append의 반환값은 새로운 slice다. 따라서 names에 반환
	fmt.Println(names)
	fmt.Println(&names[0])
}
