package main

import (
	"fmt"

	"sisi/pkg/utils"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `sisi -- Simple SIMPLON API CLI
	Tell Sisi if she should get, set, or do anything for you. Be friendly and say 'please'.

	Usage:
	sisi <ip> get <module> <param> <key> [-a <api> -t <timeout> please -p <port>]
	sisi <ip> set <module> <param> <key> <value> [-a <api> -t <timeout> please -p <port>]
	sisi <ip> do <module> <task> [-a <api> -t <timeout> please -p <port>]
	sisi -h | --help

	Options:
	-p <port>	  API port [default: 80]
	-a <api>      API version [default: 1.8.0]
	-t <timeout>  request timeout in seconds [default: 5]
	-h --help     Show this help screen.
	please        just being friendly, optionally `

	args, _ := docopt.ParseDoc(usage)
	timeout, _ := args.Int("-t")
	port, _ := args.Int("-p")

	c := utils.NewClient(
		args["<ip>"].(string),
		port,
		args["-a"].(string),
		timeout)

	if args["get"] == true {
		resp, err := c.Get(args["<module>"].(string),
			args["<param>"].(string),
			args["<key>"].(string))

		c.Print(resp, err)

	} else if args["set"] == true {
		resp, err := c.Set(args["<module>"].(string),
			args["<param>"].(string),
			args["<key>"].(string),
			args["<value>"].(string),
		)

		c.Print(resp, err)

	} else if args["do"] == true {
		resp, err := c.Do(args["<module>"].(string),
			args["<task>"].(string),
		)

		c.Print(resp, err)

	} else {
		fmt.Println(usage)
	}

	if args["please"] == true {
		fmt.Println("You're weclome!")
	}

}
