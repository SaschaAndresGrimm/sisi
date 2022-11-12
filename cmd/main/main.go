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
	sisi <ip> get <module> <param> <key> [-a <api> -t <timeout> please]
	sisi <ip> set <module> <param> <key> <value> [-a <api> -t <timeout> please]
	sisi <ip> do <module> <task> [-a <api> -t <timeout> please]
	sisi -h | --help

	Options:
	-a <api>      API version [default: 1.8.0]
	-t <timeout>  request timeout in seconds [default: 2]
	-h --help     Show this help screen.
	please        just being friendly, optionally `

	args, _ := docopt.ParseDoc(usage)
	timeout, _ := args.Int("-t")

	c := utils.NewClient(args["<ip>"].(string),
		args["-a"].(string),
		timeout)

	if args["get"] == true {
		c.Get(args["<module>"].(string),
			args["<param>"].(string),
			args["<key>"].(string))

	} else if args["set"] == true {
		c.Set(args["<module>"].(string),
			args["<param>"].(string),
			args["<key>"].(string),
			args["<value>"].(string),
		)

	} else if args["do"] == true {
		c.Do(args["<module>"].(string),
			args["<task>"].(string),
		)

	} else {
		fmt.Println(usage)
	}

	if args["please"] == true {
		fmt.Println("You're weclome!")
	}

}
