package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/leonkaihao/nycab/pkg/cli"
	log "github.com/sirupsen/logrus"
)

func main() {
	parser := cli.NewParser()
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
