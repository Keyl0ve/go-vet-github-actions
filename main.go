package main

import "fmt"

func main() {
	fmt.Print("test")
}

type User struct {
	userName string
	age      int
	country  Country
}

type Country struct {
	address    string
	userNumber string
	userFriend Friend
}

type Friend struct {
	userNum int
}

// 構造体の term check
func (u User) a() {
	u.userName = "John"
	u.age = 10
	u.country.address = "Japan"
	u.country.userNumber = "aaa"
}

// func (uu User) b() {
// 	uu.userName = "aaa"
// 	uu.age = 10
// 	uu.country.userFriend.userNum = 10
// }

// func (uuu User) c() {
// 	uuu.userName = "aaa"
// 	uuu.age = 10
// }

// func (uuuu User) d() {
// 	uuuu.userName = "aaa"
// 	uuuu.age = 10
// }

// func (user User) e() {
// 	user.userName = "John"
// 	user.country.userNumber = "aaa"
// 	user.country.address = "Japan"
// }
