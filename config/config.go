package config

import(
	"fmt"
	"encoding/json"
	"os"
)

var (
	Token string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file..")

	file, err := os.ReadFile("./con.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(config)
		fmt.Print(string(file))
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}