Ruler - User & Technical Guide

This guide covers everything you need to know to install, configure, build, and run Ruler, a Go-based moving target network defense tool.

Command Syntax
sudo go run Ruler.go [options] <network_interface>

Options
-interval <minutes> : Set the duration between IP rotations (default: 7 minutes).
Example
To run Ruler on your network interface (e.g., eth0) with a 5-minute rotation cycle:
sudo go run main.go -interval 5 eth0

Building a Binary
If you want to compile Ruler into a standalone executable file:
go build -o ruler Ruler.go
sudo ./ruler -interval 5 eth0

Exit & Cleanup
You can stop Ruler at any time by pressing Ctrl + C or sending a termination signal (SIGTERM). The application will automatically catch the signal, purge all active virtual decoy IP addresses from your network interface, and restore your system state cleanly.

IMPORTANT: This software is provided strictly for educational, defensive research, and authorized network security testing purposes. The author assumes no liability for misuse or damage caused by this program. Always ensure you have explicit authorization before running network-modifying tools.
