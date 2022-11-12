package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	//client struct to access SIMPLON rest-like API
	ip, api string
	timeout int
	client  *resty.Client
}

func NewClient(ip, api string, timeout int) *Client {
	//constructor returning a Client struct to interact with SIMPLON rest-like API
	c := new(Client)
	c.ip = ip
	c.api = api
	c.timeout = timeout

	c.client = resty.New()
	c.client.SetTimeout(time.Duration(c.timeout) * time.Second)
	c.client.SetHeader("Accept", "application/json")

	return c
}

func (c Client) url(module, param, key string) string {
	//return API url
	url := fmt.Sprintf("http://%s/%s/api/%s/%s/%s",
		c.ip,
		module,
		c.api,
		param,
		key)
	return url
}

func (c Client) get(module, param, key string) (resty.Response, error) {
	// get request for ressources

	resp, err := c.client.R().
		Get(c.url(module, param, key))

	print(*resp, err)
	return *resp, err
}

func (c Client) set(module, param, key, value string) (resty.Response, error) {
	//put request for ressources
	body := ""

	if isNumeric(value) {
		body = fmt.Sprintf("{\"value\":%s}", value)
	} else {
		body = fmt.Sprintf("{\"value\":\"%s\"}", value)
	}

	resp, err := c.client.R().
		SetBody(body).
		Put(c.url(module, param, key))

	print(*resp, err)

	return *resp, err
}

func (c Client) do(module, task string) (resty.Response, error) {
	// put request for command ressource
	body := ""

	resp, err := c.client.R().
		SetBody(body).
		Put(c.url(module, "command", task))

	print(*resp, err)

	return *resp, err
}

func isNumeric(s string) bool {
	//check if string is a numeric value
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func print(resp resty.Response, err error) {
	//pretty printing of SIMPLON API replies

	if resp.Request.Body != nil {
		fmt.Println(resp.Request.Method,
			resp.Request.URL,
			resp.Request.Body,
			resp.Status(),
			resp.Time())
	} else {
		fmt.Println(resp.Request.Method,
			resp.Request.URL,
			resp.Status(),
			resp.Time())
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	//print dict and list entries
	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		if len(resp.String()) > 2 {
			s := resp.String()
			s = s[1 : len(s)-1]
			for _, val := range strings.Split(s, ",") {
				fmt.Println(" ", val)
			}
		}
	} else if resp.StatusCode() != 404 {
		fmt.Println(resp.String())
	}
}

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

	c := NewClient(args["<ip>"].(string),
		args["-a"].(string),
		timeout)

	if args["get"] == true {
		c.get(args["<module>"].(string),
			args["<param>"].(string),
			args["<key>"].(string))

	} else if args["set"] == true {
		c.set(args["<module>"].(string),
			args["<param>"].(string),
			args["<key>"].(string),
			args["<value>"].(string),
		)

	} else if args["do"] == true {
		c.do(args["<module>"].(string),
			args["<task>"].(string),
		)

	} else {
		fmt.Println(usage)
	}

	if args["please"] == true {
		fmt.Println("You're weclome!")
	}

}
