Go-lang 스터디 진행현황
=============
##### 노마드코더-쉽고 빠른 Go 시작하기
> __2020.04.13__ 
>>
>> **1.4 Functions part Two**
>>
>> ```go
>> package main
>>
>> import (
>> 	"fmt" 
>> 	"strings"
>> ) 
>>
>> func lenAndUpper(name string) (lenght int, uppercase string) {
>> 	defer fmt.Println("I'm done") //함수가 종료된 후에 실행되는 defer
>> 	lenght = len(name)
>> 	uppercase = strings.ToUpper(name)
>> 	return
>> }
>>
>> func main() {
>> 	totalLenght, upper := lenAndUpper("park") //go에서는 여러 개의 값을 반환 가능(c에서는 안됬는데...)
>> 	fmt.Println(totalLenght, upper)
>> }
>> ```
>>
>>**1.5 for, range, ...args**
>>
>>```go
>>package main
>>import "fmt"
>>
>>func superAdd(numbers ...int) int { //arguments로 여러 개를 가져오고 싶다면 '...'를 사용
>>	total := 0
>>	for _, number := range numbers {//range: 배열을 반복할 때 사용함
>>		total += number
>>  }
>>	return total
>>}
>>
>>func main() {
>>	result := superAdd(1, 2, 3, 4, 5, 6) 
>>	fmt.Println(result)
>>}
>>```
>> **1.6 If with a Twist(특별한게 없어서 코드패스)**
>>
>> **1.7 Switch**
>>
>>```go
>>package main
>>
>>import "fmt"
>>
>>func canIDrink(age int) bool {
>>	switch koreanAge := age + 2; koreanAge {
>>	case 10:
>>		return false
>>	case 18:
>>		return true
>>		/*case age < 10: //go에서는 switch에 조건을 만들 수 있다. if, else if 처럼 사용 가능
>>			return false
>>		case age == 18:
>>			return true
>>		case age > 50:
>>			return false*/
>>	}
>>	return false
>>}
>>
>>func main() {
>>	fmt.Println(canIDrink(18))
>>}
>>```
>>> __다른 프로그래밍 언어와 다르지 않은 기능들__
>
> __20.04.14__
>
>>**1.8 Pointers!**
>>
>>```go
>>package main
>>
>>import "fmt"
>>
>>func main() {
>>	a := 2
>>	b := &a
>>	a = 10
>>	*b = 20
>>	fmt.Println(a, *b) //Low-level programing
>>} //c의 POINTER 개념하고 똑같네
>>```
>>
>>**1.9 Arrays and Slices**
>>
>>```go
>>package main
>>
>>import "fmt"
>>
>>func main() {
>>	names := []string{"nico", "lynn", "dal"}
>>	//slices는 length가 없는 배열 (?Vector와의 차이점은?)
>>	names = append(names, "what?") //append의 반환값은 새로운 slice다. 따라서 names에 반환
>>	fmt.Println(names)
>>	fmt.Println(&names[0])
>>}// vector와 비슷한 개념인 듯 하다.
>>```
