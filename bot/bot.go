package bot

import(
	"fmt"
	"github.com/hiteshjain48/go-discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start(){
	goBot, err := discordgo.New("Bot "+config.Token)
	if err != nil {
		fmt.Println("session err")
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")
	if err != nil {
		fmt.Println("user error")
		fmt.Println(err.Error())
		return
	}
	BotID = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println("open error!")
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Bot is running!")
}

func messageHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {

	if msg.Author.ID == BotID {
		return
	}

	if msg.Content == "ting" {
		_, _ = session.ChannelMessageSend(msg.ChannelID, "tong")
	}
}