package main

import(
	"fmt"
	"github.com/hiteshjain48/go-discord-bot/bot" 
	"github.com/hiteshjain48/go-discord-bot/config"

)

func main(){
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Start()
	
	 <- make(chan struct{})
	 return
}