package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ketamine" {

		formats := []string{
			"You need to be TRANQULISED!",
			"Where are the party girls?",
			"EXTEND! EXTEND! EXTEND!",
			"DAMAGE AHHHHHH",
			"Small dose, big dose, or VIETNAMESE DOSE?",
		}

		line := formats[rand.Intn(len(formats))]

		_, err := s.ChannelMessageSend(m.ChannelID, line)
		if err != nil {
			fmt.Println(err)
		}
	}

	if strings.Contains(m.Content, "alvin") || strings.Contains(m.Content, "Alvin") {

		_, err := s.ChannelMessageSend(m.ChannelID, "https://c6899f20efcf82aa3b68-31e49c3dbc881ef6f419c0bc595d3dfb.ssl.cf4.rackcdn.com/alvin.JPG")
		if err != nil {
			fmt.Println(err)
		}

	}

	if strings.Contains(m.Content, "chun") || strings.Contains(m.Content, "Chun") {

		_, err := s.ChannelMessageSend(m.ChannelID, "https://c6899f20efcf82aa3b68-31e49c3dbc881ef6f419c0bc595d3dfb.ssl.cf4.rackcdn.com/chun.JPG")
		if err != nil {
			fmt.Println(err)
		}

	}

	if strings.Contains(m.Content, "rio") || strings.Contains(m.Content, "Rio") {

		_, err := s.ChannelMessageSend(m.ChannelID, "https://c6899f20efcf82aa3b68-31e49c3dbc881ef6f419c0bc595d3dfb.ssl.cf4.rackcdn.com/rio.JPG")
		if err != nil {
			fmt.Println(err)
		}

	}

	if strings.Contains(m.Content, "cedric") || strings.Contains(m.Content, "Cedric") {

		_, err := s.ChannelMessageSend(m.ChannelID, "https://c6899f20efcf82aa3b68-31e49c3dbc881ef6f419c0bc595d3dfb.ssl.cf4.rackcdn.com/cedric.JPG")
		if err != nil {
			fmt.Println(err)
		}

	}

	if strings.Contains(m.Content, "joe") || strings.Contains(m.Content, "Joe") {

		_, err := s.ChannelMessageSend(m.ChannelID, "https://c6899f20efcf82aa3b68-31e49c3dbc881ef6f419c0bc595d3dfb.ssl.cf4.rackcdn.com/joe.JPG")
		if err != nil {
			fmt.Println(err)
		}

	}

	if strings.Contains(m.Content, "house") || strings.Contains(m.Content, "House") {

		_, err := s.ChannelMessageSend(m.ChannelID, "https://c6899f20efcf82aa3b68-31e49c3dbc881ef6f419c0bc595d3dfb.ssl.cf4.rackcdn.com/house.JPG")
		if err != nil {
			fmt.Println(err)
		}

	}

	if strings.Contains(m.Content, "bryan") || strings.Contains(m.Content, "Bryan") {

		_, err := s.ChannelMessageSend(m.ChannelID, "https://c6899f20efcf82aa3b68-31e49c3dbc881ef6f419c0bc595d3dfb.ssl.cf4.rackcdn.com/bryan.JPG")
		if err != nil {
			fmt.Println(err)
		}

	}

	if m.Content == "!peacock" {

		_, err := s.ChannelMessageSend(m.ChannelID, "https://c.tenor.com/qdXw4s3RjlIAAAAC/the-other-guys-mark-wahlberg.gif")
		if err != nil {
			fmt.Println(err)
		}

	}

	if m.Content == "<@!957090513696211025>" {

		help := "Yes m'lord? How can I help?"

		// Send a text message with the list of Gophers
		_, err := s.ChannelMessageSend(m.ChannelID, help)
		if err != nil {
			fmt.Println(err)
		}
	}

}
