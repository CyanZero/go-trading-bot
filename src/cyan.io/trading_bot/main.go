package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"cyan.io/apis"
	"cyan.io/cli"
	"cyan.io/utils"
)

func main() {

	auth := os.Getenv("CAL_BOT_AUTH")
	sec := os.Getenv("CAL_BOT_SECRET")
	logPath := os.Getenv("CAL_BOT_LOG_PATH")

	// Setup API credentials
	if len(auth) > 0:
		fmt.Printf("Auth from env: %s***\n", auth[:5])
	fmt.Printf("Log path: %s\n", logPath)
	if auth == "" || sec == "" {
		log.Fatal("No valid credentail found, quit..")

		os.Exit(1)
	}

	utils.Authorization = auth
	utils.Secret = []byte(sec)

	// Setup app options, supported options:
	// -v, verbose
	// -c, commandline mode
	verbose := flag.Bool("v", false, "a bool")
	commandLineMode := flag.Bool("c", false, "Initiate commandline mode")
	flag.Parse()

	// Setup app logs
	f, err := os.OpenFile(logPath+"bot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("This is a test log entry")

	if *verbose {
		log.SetOutput(os.Stderr)
	}

	// The main interface
	// Default in web service mode
	// Unless commandLine Mode is indicated
	if *commandLineMode {
		cli.CommandLineMode()
	} else {
		router := apis.NewRouter()

		log.Println("Rest API server started via port: " + utils.PORT)
		log.Fatal(http.ListenAndServe(utils.PORT, router))
	}

}
