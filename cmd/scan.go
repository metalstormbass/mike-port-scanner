/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"
	"net/netip"
	"os"
	"regexp"
	"strings"

	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan an IP address or range",
	Long:  `Scan an ip range by using: scan --network=192.168.0.0/24`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This will work eventually.....%v\n",
			emoji.ClownFace)

		networkInput := args[0]

		// Validate and Parse Input
		ipAddress, netmask, networkAddress := ParseNetwork(networkInput)

		fmt.Println(ipAddress, netmask, networkAddress)

		//Scan()

		//FormatOutput()

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Function to validate input and to seperate mask and IP into different variables
func ParseNetwork(networkInput string) (ipAddress string, netmask string, networkAddress []netip.Addr) {

	// Define default netmask
	netmask = "32"

	// Regex to do initial Validation
	valid := validateInput(networkInput)

	if !valid {
		fmt.Println("Bad Input")
		os.Exit(2)
	} else {

		// Switch to split up variables
		switch {
		case strings.Contains(networkInput, "/"):
			ipAddress = strings.Split(networkInput, "/")[0]
			netmask = strings.Split(networkInput, "/")[1]
		default:
			ipAddress = networkInput
		}
	}
	// Get Network Address
	_, ipv4Net, err := net.ParseCIDR(ipAddress + "/" + netmask)
	if err != nil {
		println("Failed to ParseCIDR")
		os.Exit(2)
	}

	// Dereference Pointer (output of get ParseCIDR)
	derefIP := *ipv4Net

	// Convert to String
	netAdd := derefIP.String()
	netAdd = strings.Split(netAdd, "/")[0]

	fmt.Println(ipv4Net)
	// Parse Prefix
	prefix, err := netip.ParsePrefix(netAdd + "/" + netmask)
	if err != nil {
		panic(err)
	}

	// Get IP Addresses to Test
	//var networkRange []string

	//for ip := ip.Mask(ipv4Net.Mask); ipv4Net.Contains(ip); ip = ip.Next() {
	//	networkRange = append(networkRange, ip.String())
	//}

	var networkRange []netip.Addr
	for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
		networkRange = append(networkRange, addr)
	}

	//fmt.Println(networkRange)
	return ipAddress, netmask, networkRange[1 : len(networkRange)-1]
}

//Regex Function

func validateInput(s string) bool {
	return regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}(\/([0-9]|[1-2][0-9]|3[0-2]))?$`).MatchString(s)
}

// Validate IP Address
//func checkIPAddress(ip string) {
//	if net.ParseIP(ip) == nil {
//		fmt.Printf("IP Address %s is invalid. Please enter a valid IP address\n", ip)
//	}
//}

func Scan() {
	fmt.Println("testing flow to Scan function")
}

func FormatOutput() {
	fmt.Println("testing flow to MainHandle function")

}
