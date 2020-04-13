Go-lang 스터디 진행현황
=============
##### 노마드코더-쉽고 빠른 Go 시작하기
> 2020.04.13 
>> *1.4 Functions part Two*
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
>>*1.5 for, range, ...args*
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
>>*1.6 If with a Twist*(특별한게 없어서 코드패스)
>>*1.7 Switch*

