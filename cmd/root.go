/*
Copyright © 2023 Nitin Panwar <contact@nitinp.dev>
*/
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/nitintf/iptracker/internal/ip"
	"github.com/nitintf/iptracker/internal/print"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iptracker",
	Short: "Track IP addresses",
	Long:  `A Go based CLI tool to track IP addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		ipAdd, err := getPrivateIpAdress()

		if err != nil {
			print.Error("Unable to retrive Private IP Address")
		}

		publicIpAddr, err := getPublicIPAddress()

		if err != nil {
			print.Error("Unable to retrive Public IP Address")
			return
		}

		print.Info("\nPrivate IP Address: " + ipAdd)
		ip.ShowIpData(publicIpAddr)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

func getPrivateIpAdress() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	var ipAddress string

	for _, address := range addrs {
		// Check if the address is an IP address and not a loopback or multicast address
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddress = ipnet.IP.String()
				break // Use the first IPv4 address found
			}
		}
	}

	if ipAddress == "" {
		return "", err
	}

	return ipAddress, nil
}

type PublicIp struct {
	Origin string `json:"origin"`
}

func getPublicIPAddress() (string, error) {
	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	data := PublicIp{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return "", err
	}

	return data.Origin, nil
}
