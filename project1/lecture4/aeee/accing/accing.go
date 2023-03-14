package accing

import (
	"project1/lecture4/aeee/account"
)

func CreateAccount(id int, userName, password string) *account.Account {
	return &account.Account{
		ID:       id,
		UserName: userName,
		Password: password,
		Balance:  0.0,
	}
}

func UpdatePassword(a *account.Account, newPassword string) {
	a.UpdatePassword(newPassword)
}

func CheckBalance(a *account.Account) float64 {
	return a.CheckBalance()
}
func TransferFunds(from, to *account.Account, amount float64) {
	if from.Balance >= amount {
		from.Balance -= amount
		to.Balance += amount
	}
}
