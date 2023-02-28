package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/common-nighthawk/go-figure"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var ok = color.Bold + color.Green + "{ Hit } " + color.Reset
var info = color.Bold + color.Yellow + "[info] " + color.Reset
var error_msg = color.Bold + color.Red + "[error] " + color.Reset

var isVerbose bool = false
var timeout int = 300

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

func banner() {
	var author string = "sc4rfurry"
	var version string = "1.0.0"
	var go_version string = "1.18.1 or higher"
	var github string = "https://github.com/sc4rfurry"
	var description string = "Tool to scan a network for alive host " + color.Bold + color.Cyan + "(!nbtscan - Host Lookup)." + color.Reset
	banner := figure.NewColorFigure("TallusX", "", "cyan", true)
	banner.Print()
	println("\n")
	fmt.Println(color.Bold + color.Green + "  " + description + "  " + color.Reset + "\n")
	fmt.Println("\t  Author: " + author)
	fmt.Println("\t  Version: " + version)
	fmt.Println("\t  Go Version: " + go_version)
	fmt.Println("\t  Github: " + github + "\n")
	fmt.Println("-------------------------------------------------------------------------------->")
	println("\n")

}

func help() {
	println(color.Bold + color.Green + "\n\t\t~ Help Menu ~" + color.Reset)
	println("\n\t" + color.Bold + color.Cyan + "Usage: " + color.Reset + "./main <subnet>")
	println(color.Bold + color.Gray + "__________________________________________________________________________________" + color.Reset)
	println("\n\t" + color.Bold + color.Green + "Options: " + color.Reset)
	println("\t\t" + color.Bold + color.Yellow + "subnet " + color.Reset + "\t\tNetwork to scan. (Ex: 192.168.0)")
	println("\t\t" + color.Bold + color.Yellow + "-d/--debug " + color.Reset + "\tEnable debug mode")
	println("\t\t" + color.Bold + color.Yellow + "-V/--version " + color.Reset + "\tShow version")
	println("\t\t" + color.Bold + color.Yellow + "-h/--help " + color.Reset + "\tShow this help menu" + "\n")
	println("\n" + info + color.Bold + color.Gray + " Subnet: " + color.Reset + "First three octets." + color.Bold + color.Cyan + " Ex: ./main 192.168.0" + color.Reset + "\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		help()
	}
	for _, arg := range os.Args {
		if arg == "-h" || arg == "--help" {
			help()
		} else if arg == "-d" || arg == "--debug" {
			isVerbose = true
		} else if arg == "-V" || arg == "--version" {
			println(color.Bold + color.Green + "TallusX " + color.Reset + color.Bold + color.Cyan + "v1.0.0" + color.Reset)
			os.Exit(0)
		}
	}
	subnetToScan := os.Args[1]
	banner()

	activeThreads := 0
	doneChannel := make(chan bool)

	println(info + "Scanning " + color.Bold + color.Cyan + subnetToScan + ".0/24" + color.Reset + " for alive hosts.")
	if !isVerbose {
		println(info + "Use " + color.Bold + color.Cyan + "-d/--debug" + color.Reset + " to enable debug mode." + "\n")
	} else {
		println(info + "Debug mode enabled." + "\n")
	}
	for ip := 0; ip <= 255; ip++ {
		fullIP := subnetToScan + "." + strconv.Itoa(ip)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		if isVerbose {
			go resolve(fullIP, doneChannel, true)
			activeThreads++
		} else {
			go resolve(fullIP, doneChannel, false)
			activeThreads++
		}
	}

	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func resolve(ip string, doneChannel chan bool, verbose bool) {
	addresses, err := net.LookupAddr(ip)
	if verbose {
		if err == nil {
			println(ok + " " + ip + " - " + color.Bold + color.Cyan + strings.Join(addresses, ", ") + color.Reset)
		} else {
			println(error_msg + " " + ip + " - " + color.Bold + color.Red + err.Error() + color.Reset)
		}
		doneChannel <- true
	} else {
		if err == nil {
			println(ok + " " + ip + " - " + color.Bold + color.Cyan + strings.Join(addresses, ", ") + color.Reset)
		}
		doneChannel <- true
	}
}
