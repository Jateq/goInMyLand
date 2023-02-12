package main

import (
	"fmt"
)

func main() {

	fmt.Println()
	fmt.Println("\033[1;33mWelcome to Jateq's market\033[0m")
	LoadDataFromJSON("users.json", &Users)
	LoadDataFromJSON("items.json", &Items)
	fmt.Println("Do you have an account? (y or n)")
	var input, username, password, corrector string
	fmt.Scanln(&input)
	if input == "y" {
		fmt.Println("\033[1;33mPlease log in\033[0m")
		fmt.Scanln(&username, &password)
		if Authorization(username, password) {
			fmt.Print("\033[1;32m")
			fmt.Println("Authorization succeeded for", username)
			fmt.Print("\033[0m")
		} else {
			for !Authorization(username, password) {
				fmt.Println("\033[1;31mAuthorization failed, try again\033[0m")
				fmt.Scanln(&username, &password)
			}
			fmt.Print("\033[1;32m")
			fmt.Println("Authorization succeeded for", username)
			fmt.Print("\033[0m")

		}
	} else if input == "n" {
		fmt.Println("Please register (username, password)")
		fmt.Scanln(&username, &password)
		fmt.Println("Please retype your password again")
		fmt.Scanln(&corrector)
		for password != corrector {
			fmt.Println("They are not match, please make sure everything is ok")
			fmt.Scanln(&corrector)
		}
		Registration(username, password)
		fmt.Println("You are welcome to Jateq's Store")

	} else {
		fmt.Println("\033[1;31mInvalid input. Please enter either 'y' or 'n'\033[0m")
	}
	if input == "y" || input == "n" {
		fmt.Println("\n\033[1;33mOffers for today:\033[0m")
		fmt.Println(Store(Items))
		//fmt.Println("-----------------------------")

		fmt.Println("\033[1;33mSearch in market:\033[0m")
		var name string
		fmt.Scanln(&name)
		items := SearchItemsByName(name)
		if len(items) == 0 {
			fmt.Println("Sorry", name, "are out of stock")
		} else {
			for _, item := range items {
				fmt.Println("Name:", item.Name)
				fmt.Println("Price:", item.Price)
				fmt.Println("Rating:", item.Rating)
				fmt.Println("-----------------------------")
			}
			var answer string
			var rate float64
			fmt.Println("\033[1;33mDo you want to buy it? (y or n)\033[0m")
			fmt.Scanln(&answer)
			if answer == "y" {
				fmt.Println("\033[1;33mDid you liked it? Give us your rating (1 - 5)\033[0m")
				fmt.Scanln(&rate)
				GiveRating(name, rate)
				fmt.Println("\033[1;32mThanks for your feedback\033[0m")
			} else {
				fmt.Println("Filter?")
			}
		}
	}
	SaveDataToJSON("users.json", &Users)
	SaveDataToJSON("items.json", &Items)
}
