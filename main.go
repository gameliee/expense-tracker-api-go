package main

import "gamelieelearn/expense-tracker-api-go/cmd"

func main() {
	err := cmd.InitializeApplication()
	if err != nil {
		panic(err)
	}
	cmd.StartApplication()
}
