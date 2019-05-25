package core

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/robinstechprojects/Butler-Robin/util"
	"strconv"
	
)


//MakeConnection connects the Bot to Discord
func MakeConnection(token string){

    

	d, err := discordgo.New(token)

    if err != nil {
        fmt.Println("failed to create discord session", err)
    }

   // bot, err := d.User("@me")

    if err != nil {
        fmt.Println("failed to access account", err)
    }

    d.AddHandler(handleCmd)
    err = d.Open()

    if err != nil {
        fmt.Println("unable to establish connection", err)
    }

    defer d.Close()

    <-make(chan struct{})
}

func handleCmd(d *discordgo.Session, msg *discordgo.MessageCreate) {
    //user := msg.Author
    //if user.ID == bid ||user.Bot {
    //    return
    //}
	sender := msg.Author.ID
	bot, err := d.User("@me")
	if sender == bot.ID{
		fmt.Println(err)
	}
    content := msg.Content

    if (content == "!test") {
        switch{
        case Lang == "DE":
            d.ChannelMessageSend(msg.ChannelID, "Test bestanden !")
    
        case Lang == "EN":
            d.ChannelMessageSend(msg.ChannelID, "Test successfull !")
        
        default:
            fmt.Println("Your Language Setting is invalig")
            d.ChannelMessageSend(msg.ChannelID, "Internal System Error, Look into Console for detailed Information")
        
        }
	}
	
	if (content == "!dice"){
		d.ChannelMessageSend(msg.ChannelID, strconv.Itoa(util.DiceInt())) //Number is generated in util
	}

	if (content == "!coinflip"){ 
		if util.CoinFlipInt() == 1 {
            switch{
            case Lang == "DE":
                d.ChannelMessageSend(msg.ChannelID, "Kopf!")
            case Lang == "EN":
                d.ChannelMessageSend(msg.ChannelID, "Heads!")
            }
		}else{
            switch{
            case Lang == "DE":
                d.ChannelMessageSend(msg.ChannelID, "Zahl!")
            case Lang == "EN":
                d.ChannelMessageSend(msg.ChannelID, "Tails!")
            }
        }
    

	}
    //fmt.Printf("Message: %+v\n", msg.Message)
}