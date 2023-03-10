#!/bin/bash

set -e

ascii_art='\r
  _____           _   _                __  __
 |_   _|   __ _  | | | |  _   _   ___  \ \/ /
   | |    / _` | | | | | | | | | / __|  \  /
   | |   | (_| | | | | | | |_| | \__ \  /  \
   |_|    \__,_| |_| |_|  \__,_| |___/ /_/\_\
'

desc="--> T4llusX - Tool to scan a network for alive host (!nbtscan - Host Lookup)."
banner="$ascii_art \r\n $desc \r \n"

if ! command -v lolcat &>/dev/null; then
    echo "lolcat could not be found"
    echo "Installing lolcat"
    sudo apt-get install lolcat
fi

if [[ $# -lt 1 ]]; then
    echo "Usage: $0 <path-to-go-code>" | lolcat -a -d 5
    exit 1
fi

if [[ ! -f "$1" ]]; then
    echo "Error: File not found: $1" | lolcat -a -d 5
    exit 1
fi

echo -e "$banner" | lolcat -a -d 5

echo -e "\r\n [+] Building TallusX \r\n" | lolcat -a -d 5

echo -e "\r\n [+] Checking for Go installation \r\n" | lolcat -a -d 5
# check for go installation
if ! command -v go &>/dev/null; then
    echo "go could not be found"
    echo "Installing go"
    sudo apt-get install golang
fi
echo -e "\r\n [+] Go installation found \r\n" | lolcat -a -d 5

echo -e "\r\n [+] Installing dependencies \r\n" | lolcat -a -d 5
# install go dependencies
go get . && go mod tidy &>/dev/null || handle_error "Failed to install dependencies. Please check your Go installation."

echo -e "\r\n [+] Dependencies installed \r\n" | lolcat -a -d 5

echo -e "----------------------------------------------------------------------------------------------------\r\n" | lolcat -a -d 5

platforms=("windows" "linux")
architectures=("386" "amd64")

output_dir="bin"

if [[ ! -d "$output_dir" ]]; then
    mkdir "$output_dir"
fi

function handle_error {
    echo "Error: $1"
    exit 1
}

trap 'handle_error "An error occurred on line $LINENO: $BASH_COMMAND"' ERR

for platform in "${platforms[@]}"; do
    for arch in "${architectures[@]}"; do
        echo -e "[+] Building for $platform-$arch\r\n" | lolcat -a -d 5
        if [[ "$platform" == "windows" ]]; then
            output="$output_dir/TallusX-$platform-$arch.exe"
        else
            output="$output_dir/TallusX-$platform-$arch"
        fi
        env GOOS=$platform GOARCH=$arch go build -a -ldflags '-s -w -extldflags "-static"' -o "$output" "$1" || handle_error "Failed to build $platform-$arch binary. Please check your Go installation."
        echo -e "\t[+] Stripping $platform-$arch binary" | lolcat -a -d 5
        strip "$output" || handle_error "Failed to strip $platform-$arch binary."
        if ! command -v upx &>/dev/null; then
            echo "upx could not be found"
            echo "Installing upx"
            sudo apt-get install upx
        fi
        echo -e "\t[+] Compressing $platform-$arch binary" | lolcat -a -d 5
        upx -9 "$output" &>/dev/null || handle_error "Failed to compress $platform-$arch binary."
        echo -e "\r\n\t[!] Binary created: $output\r\n" | lolcat -a -d 5
    done
done