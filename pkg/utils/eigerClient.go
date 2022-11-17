package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type client struct {
	//client struct to access SIMPLON rest-like API
	ip, api string
	timeout int
	client  *resty.Client
}

func NewClient(ip, api string, timeout int) *client {
	//constructor returning a Client struct to interact with SIMPLON rest-like API
	c := new(client)
	c.ip = ip
	c.api = api
	c.timeout = timeout

	c.client = resty.New()
	c.client.SetTimeout(time.Duration(c.timeout) * time.Second)
	c.client.SetHeader("Accept", "application/json")

	return c
}

func (c client) url(module, param, key string) string {
	//return API url
	url := fmt.Sprintf("http://%s/%s/api/%s/%s/%s",
		c.ip,
		module,
		c.api,
		param,
		key)
	return url
}

func (c client) Get(module, param, key string) (resty.Response, error) {
	// get request for ressources

	resp, err := c.client.R().
		Get(c.url(module, param, key))

	return *resp, err
}

func (c client) Set(module, param, key, value string) (resty.Response, error) {
	//put request for ressources
	var body string

	if isNumeric(value) {
		body = fmt.Sprintf("{\"value\":%s}", value)
	} else if v, err := strconv.ParseBool(value); err == nil {
		body = fmt.Sprintf("{\"value\":%v}", v)
	} else {
		body = fmt.Sprintf("{\"value\":\"%s\"}", value)
	}

	resp, err := c.client.R().
		SetBody(body).
		Put(c.url(module, param, key))

	return *resp, err
}

func (c client) Do(module, task string) (resty.Response, error) {
	// put request for command ressource
	resp, err := c.client.R().
		Put(c.url(module, "command", task))

	return *resp, err
}

func isNumeric(s string) bool {
	//check if string is a numeric value
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func (c client) Print(resp resty.Response, err error) {
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
