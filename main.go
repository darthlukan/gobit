package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/stretchr/objx"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	mashKey    = "Ax1B5LkBRdmshDXrwUOszDMYNKsUp1UP9ErjsnLOZfENrxPx2a"
	baseUrl    = "https://community-bitcointy.p.mashape.com/"
	fromBTCUrl = "price/"
	toBTCUrl   = "convert/"
)

func request(url string) (objx.Map, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Machape-Key", mashKey)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	data, err := objx.FromJSON(string(body))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func FromBTC(currency string) string {
	url := fmt.Sprintf("%v%v%v", baseUrl, fromBTCUrl, currency)
	data, err := request(url)
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}
	return fmt.Sprintf("1 BTC == %v %v\n", data, currency)
}

func ToBTC(value, currency string) string {
	url := fmt.Sprintf("%v%v%v%v", baseUrl, toBTCUrl, value, currency)
	data, err := request(url)
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	currVal := data.Get("value")
	curr := data.Get("currency")

	return fmt.Sprint("%v BTC == %v %v\n", value, currVal, curr)

}

func main() {
	app := cli.NewApp()
	app.Name = "gobit"
	app.Usage = "Perform currency conversions between major currencies and Bitcoin."
	app.Version = "0.1"
	app.Action = func(c *cli.Context) {
		fmt.Printf("Args: %v", c.Args())
		return
	}
	app.Run(os.Args)
}
