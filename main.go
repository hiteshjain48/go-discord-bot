package main

import(
	"fmt"

)

func main(){
	err := config.ReadConfig()

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	bot.Start()
	
	 <- make(chan struct{})
	 return
}