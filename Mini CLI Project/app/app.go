package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// This function will return de CLI application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "IP and domain scrapper - CLI"
	app.Usage = "Searches for ips and domains in the internet and returns info about it."

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "searchIps",
			Usage:  "Serches IP addresses in the internet.",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "searchServers",
			Usage:  "Serches server names in the internet.",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	// net package
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ips := range ips {
		fmt.Printf("Found the ip %s for host %s. \n", ips, host)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	// net package
	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Printf("Found the name server %s for host %s. \n", server.Host, host)
	}
}
