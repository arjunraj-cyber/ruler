package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func showBanner() {
	cyan := "\033[36m"
	reset := "\033[0m"

	fmt.Println(cyan + `
    ██████╗ ██╗   ██╗██╗     ███████╗██████╗ 
    ██╔══██╗██║   ██║██║     ██╔════╝██╔══██╗
    ██████╔╝██║   ██║██║     █████╗  ██████╔╝
    ██╔══██╗██║   ██║██║     ██╔══╝  ██╔══██╗
    ██║  ██║╚██████╔╝███████╗███████╗██║  ██║
    ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝
    [ Secure Defense Layer - Fake IP Decoy System ]
	` + reset)
}

func generateFakeIPs() []string {
	var ips []string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {
		ip := fmt.Sprintf("10.255.254.%d", r.Intn(200)+10)
		ips = append(ips, ip)
	}
	return ips
}

func validateInterface(ifaceName string) error {
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return fmt.Errorf("interface %s does not exist", ifaceName)
	}
	if iface.Flags&net.FlagUp == 0 {
		return fmt.Errorf("interface %s is down", ifaceName)
	}
	return nil
}

func applyDecoys(iface string, oldIPs []string, newIPs []string) {
	for _, ip := range oldIPs {
		cmd := exec.Command("ip", "addr", "del", ip+"/24", "dev", iface)
		_ = cmd.Run()
	}

	fmt.Printf("\n[+] Rotating Decoy Grid on interface [%s]...\n", iface)
	for _, ip := range newIPs {
		cmd := exec.Command("ip", "addr", "add", ip+"/24", "dev", iface)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("[-] Failed to bind fake IP %s (Are you root?)\n", ip)
		} else {
			fmt.Printf("[*] Deployed Decoy Active -> %s\n", ip)
		}
	}
}

func clearAll(iface string, ips []string) {
	fmt.Println("\n[!] Shutting down Ruler... Purging fake IP traces.")
	for _, ip := range ips {
		cmd := exec.Command("ip", "addr", "del", ip+"/24", "dev", iface)
		_ = cmd.Run()
	}
	fmt.Println("[+] Network restored to clean state.")
}

func main() {
	showBanner()

	intervalFlag := flag.Int("interval", 7, "Rotation interval in minutes")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: sudo go run main.go [options] <network_interface>")
		fmt.Println("Example: sudo go run main.go -interval 5 eth0")
		os.Exit(1)
	}

	iface := args[0]

	if err := validateInterface(iface); err != nil {
		fmt.Printf("[-] Validation Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[+] Initializing Ruler on verified target interface: %s\n", iface)
	fmt.Printf("[+] Rotation cycle set to every %d minutes.\n", *intervalFlag)

	var currentIPs []string
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(time.Duration(*intervalFlag) * time.Minute)
	defer ticker.Stop()

	currentIPs = generateFakeIPs()
	applyDecoys(iface, []string{}, currentIPs)

	go func() {
		for range ticker.C {
			newIPs := generateFakeIPs()
			applyDecoys(iface, currentIPs, newIPs)
			currentIPs = newIPs
		}
	}()

	<-sigChan
	clearAll(iface, currentIPs)
}
