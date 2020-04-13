Go-lang 스터디 진행현황
=============
##### 노마드코더-쉽고 빠른 Go 시작하기
> 2020.04.13 
>> *1.5 for, range, ...args*
>> ```go
>> package main
>>
>> import (
>> 	"fmt"
>> 	"strings"
>> ) 

>> func lenAndUpper(name string) (lenght int, uppercase string) {
>> 	defer fmt.Println("I'm done")
>> 	lenght = len(name)
>> 	uppercase = strings.ToUpper(name)
>> 	return
>> }

>> func main() {
>> 	totalLenght, upper := lenAndUpper("park")
>> 	fmt.Println(totalLenght, upper)
>> }
>> ```

