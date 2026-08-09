Ruler 👑
A lightweight, high-performance network defense tool written in Go designed to confuse active reconnaissance and network mapping attempts. Ruler dynamically binds and rotates a block of randomized decoy IP addresses on your network interface at custom intervals, protecting your system's real IP presence.
Generates and rotates a block of 10 randomized decoy IP addresses within an isolated private subnet block (10.255.254.x) to prevent ARP collisions.
Customize how frequently your decoy grid shifts using the -interval flag to bypass rapid IDS/IPS signature triggers.
Automatically validates that your target network interface exists and is active before deploying decoys.
Listens for interrupt signals (Ctrl + C, SIGTERM) to automatically purge all fake IP traces and restore your network to a clean state upon exit.

Disclaimer
This software is provided for educational, defensive research, and authorized network security testing purposes only. The author and contributors assume no liability and are not responsible for any misuse or damage caused by this program. Users are entirely responsible for complying with local, state, and federal laws regarding network administration and security testing. Do not use this tool for unauthorized or illegal activities.

