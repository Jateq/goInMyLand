package account

import (
	"encoding/json"
	"os"
)

type Account struct {
	ID       int
	UserName string
	Password string
	Balance  float64
}

func (a *Account) UpdatePassword(newPassword string) {
	a.Password = newPassword
}

func (a *Account) CheckBalance() float64 {
	return a.Balance
}

func (a *Account) SaveToJSON(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(a)
}
