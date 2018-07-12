package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/urfave/cli.v2"
)

func CommandHandler(c *cli.Context) error {
	// fmt.Println(c.Command.FullName())
	uri := fmt.Sprintf("/%s", strings.Replace(c.Command.FullName(), " ", "/", -1))

	res, err := wnc.DoRequest(uri, http.MethodGet)
	if err == nil {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	return nil
}
