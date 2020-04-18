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
>>
>>**1.10 Maps**
>>
>>```go
>>package main
>>
>>import "fmt"
>>
>>func main() {
>>	nico := map[string]string{"name": "nico", "age": "12"}
>>	for _, value := range nico {
>>		fmt.Println(value)
>>	}
>>}
>>```
>>
>
> __20.04.15__
>
>>**1.11 structs**
>>
>>```go
>>package main
>>
>>import "fmt"
>>
>>type person struct {
>>	name    string
>>	age     int
>>	favFood []string
>>}
>>
>>func main() {
>>	favFood := []string{"kimchi", "ramen"}
>>	nico := person{name: "nico", age: 18, favFood: favFood} // field:value or value >>통일 시켜야 한다.
>>	fmt.Println(nico.name)
>>}//c에서의 구조체와 유사
>>```
>>
>
> __20.04.16__
>
>>**2.0 Account + NewAccount**
>>```go
>>package main
>>// main.go
>>import (
>>	"fmt"
>>
>>	"github.com/park/learngo/accounts"
>>)
>>
>>func main() {
>>	account := accounts.NewAccount("park")
>>	fmt.Println(account)
>>}
>>```
>>```go
>>// acounts.go
>>package accounts
>>
>>// Account strunt
>>type Account struct {//이 struct는 private
>>	owner   string
>>	balance int
>>}
>>
>>// NewAccount creates Account
>>func NewAccount(owner string) *Account {//constructor 역할을 함수가 대신함
>>	account := Account{owner: owner, balance: 0}
>>	return &account //포인터를 이용해서 sturct를 새로만들지 않기 위해 만든 sturct의 주소값을 반환
>>}
>>//struct를 이용해서 은행 계좌 만들기
>>```
>>
>>**2.1 Methods part One, Two**
>>
>>```go
>>package accounts
>>//accounts.go
>>
>>// Account strunt
>>type Account struct {
>>	owner   string
>>	balance int
>>}
>>
>>var errnoMoney = errors.New("Can't withdraw") //에러메세지
>>
>>// NewAccount creates Account
>>func NewAccount(owner string) *Account {
>>	account := Account{owner: owner, balance: 0}
>>	return &account
>>}
>>
>>// Deposit x amount on your account 계좌에 금액추가
>>func (a *Account) Deposit(amount int) {// 우리가 처음 만들었던 Account에 접근하려면 주소로 접근해야한다.(그냥 Account를 사용하면 '복사본'에 접근한다.) 따라서 포인터로 receiver도 포인터로 생성
>>	a.balance += amount
>>}
>>
>>// Balance of your account 계좌 확인
>>func (a Account) Balance() int {//method를 생성할 땐 func 와 이름 사이에 receiver를 만들어줌. 이때 이름은 해당 struct의 첫 글자의 소문자여야 한다 a <- 'A'ccount
>>	return a.balance
>>}
>>// Withdraw x amount from your account
>>func (a *Account) Withdraw(amount int) error {
>>	if a.balance < amount {// 에러체크
>>		return errnoMoney
>>	}
>>	a.balance -= amount
>>	return nil
>>}
>>```
>>
>>```go
>>package main
>>//main.go
>>import (
>>	"fmt"
>>
>>	"github.com/park/learngo/accounts"
>>)
>>
>>func main() {
>>	account := accounts.NewAccount("park")
>>	account.Deposit(10)
>>	fmt.Println(account.Balance())
>>	err := account.Withdraw(20)
>>	if err != nil {// error 체크를 직접 만들어줘야함. 아니면 아무것도 출력하지 않는다.
>>		fmt.Println(err)
>>	}
>>	fmt.Println(account.Balance())
>>
>>}
>>```
>>
>
>__20.04.18__
>
>>**2.3 Finishing Up**
>>
>>```go
>>package main
>>
>>import (
>>	"fmt"
>>
>>	"github.com/park/learngo/accounts"
>>)
>>
>>func main() {
>>	account := accounts.NewAccount("park")
>>	account.Deposit(10)
>>
>>	fmt.Println(account)//기본적으로 go 에서 지원하는 method 호출
>>
>>}
>>```
>>
>>```go
>>package accounts
>>
>>import (
>>	"errors"
>>	"fmt"
>>)
>>
>>// Account strunt
>>type Account struct {
>>	owner   string
>>	balance int
>>}
>>
>>var errnoMoney = errors.New("Can't withdraw")
>>
>>// NewAccount creates Account
>>func NewAccount(owner string) *Account {
>>	account := Account{owner: owner, balance: 0}
>>	return &account
>>}
>>
>>// Deposit x amount on your account
>>func (a *Account) Deposit(amount int) {
>>	a.balance += amount
>>}
>>
>>// Balance of your account
>>func (a Account) Balance() int {
>>	return a.balance
>>}
>>
>>// Withdraw x amount from your account
>>func (a *Account) Withdraw(amount int) error {
>>	if a.balance < amount {
>>		return errnoMoney
>>	}
>>	a.balance -= amount
>>	return nil
>>}
>>
>>// ChangeOwner of the account
>>func (a *Account) ChangeOwner(newOwner string) {
>>	a.owner = newOwner
>>}
>>
>>// Owner of the account
>>func (a Account) Owner() string {
>>	return a.owner
>>}
>>
>>func (a Account) String() string { //기본적으로 go에서 지원하는 method
>>	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
>>}
>>```
>>

