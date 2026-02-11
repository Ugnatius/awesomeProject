package main

import "fmt"

func main() {
	score := 5
	if score > 10 {
		if score > 15 {
			fmt.Println("You are too good.")
		} else {
			fmt.Println("You are just good.")
		}

	} else {
		fmt.Println("You must to learn how to use this tool.")
	}

	mark := 8
	if mark >= 15 {
		fmt.Println("You are too good.")
	} else {
		if mark > 10 {
			fmt.Println("You are just good.")

		} else {
			fmt.Println("You must to learn how to use this tool.")
		}
	}

	age := 10
	if age == 20 {
		fmt.Println("You are too young.")
	} else if age == 50 {
		fmt.Println("You are middle")
	} else if age == 10 {
		fmt.Println("You are baby.")
	} else {
		fmt.Println("You must to be another")
	}

	subscribed := true
	if subscribed {
		fmt.Println("You are subscribed.")
	} else {
		fmt.Println("You are not subscribed.")
	}

	ochko := 16
	if ochko > 5 && ochko < 15 {
		fmt.Println("Ты лох")
	} else {
		fmt.Println("Ты просто обосрался")
	}
}
