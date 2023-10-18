/*
Copyright © 2023 Nitin Panwar <contact@nitinp.dev>
*/
package cmd

import (
	"regexp"

	"github.com/nitintf/iptracker/internal/ip"
	"github.com/nitintf/iptracker/internal/print"
	"github.com/spf13/cobra"
)

func isValidIPAddress(ip string) bool {
	// Regular expression for a valid IPv4 or IPv6 address
	ipPattern := `^([0-9]{1,3}\.){3}[0-9]{1,3}$|` +
		`^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|` +
		`([0-9a-fA-F]{1,4}:){1,7}:|` +
		`([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|` +
		`([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|` +
		`([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|` +
		`([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|` +
		`([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|` +
		`[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|` +
		`:((:[0-9a-fA-F]{1,4}){1,7}|:)|` +
		`fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|` +
		`::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|` +
		`([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`

	match, err := regexp.MatchString(ipPattern, ip)
	if err != nil {
		return false
	}
	return match
}

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track IP address information",
	Long:  `sub-command for the IP address CLI tool, allowing you to track and gather information about a specific IP address`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			ipAdd := args[0]
			isValid := isValidIPAddress(ipAdd)

			if isValid {
				ip.ShowIpData(ipAdd)
			} else {
				print.Error("Invalid IP address.")
			}
		} else {
			print.Error("Please provide IP to trace.")
		}
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
}
