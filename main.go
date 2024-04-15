package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/cypher-org/cypher/plugins"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	if test_token := os.Getenv("TEST_TOKEN"); len(test_token) != 0 {
		token = test_token
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error occured during inititalization: %v", err)
		os.Exit(1)
	}

	session.AddHandler(plugins.Handler)

	err = session.Open()
	if err != nil {
		log.Fatalf("Error while opening a websocket connection: %v", err)
		os.Exit(1)
	}

	defer session.Close()

	plugins.ApplyCommands(session)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to Stop")
	<-stop

	log.Println("Stopping..")
}
