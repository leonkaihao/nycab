package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/leonkaihao/nycab/pkg/cli"
	"github.com/leonkaihao/nycab/pkg/config"
	log "github.com/sirupsen/logrus"
)

func getConfig() (conf *config.Cli, err error) {
	v, err := config.LoadConfig("cli.nycab")
	if err != nil {
		return
	}
	conf, err = config.GetCliConfig(v)
	if err != nil {
		return
	}
	return
}

func main() {
	conf, err := getConfig()
	if err != nil {
		log.Fatalln(err)
	}
	parser := cli.NewParser(conf)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Command: ")
	cmdline, err := reader.ReadString('\n')
	for err == nil {
		err = parser.Parse(cmdline)
		if err != nil {
			log.Errorln(err)
		}
		fmt.Print("Enter Command: ")
		cmdline, err = reader.ReadString('\n')
	}
}
