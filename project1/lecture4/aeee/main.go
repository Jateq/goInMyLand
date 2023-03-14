package main

import (
	"fmt"
	"project1/lecture4/aeee/accing"
	"project1/lecture4/aeee/account"
)

func main() {

	a := accing.CreateAccount(2, "Elaman", "password")
	fmt.Println("Created account:", a)
	accs, err := account.LoadFromJson("account.json")
	if err != nil {
		panic(err)
	}

	accing.UpdatePassword(a, "newpassword")
	fmt.Println("Updated password:", a)

	balance := accing.CheckBalance(a)
	fmt.Println("Balance:", balance)
	fmt.Println(accs)
	err1 := a.SaveToJSON("account.json")
	if err1 != nil {
		panic(err1)
	}
}
