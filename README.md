<h1 align="center">
  TallusX
</h1>

<h4 align="center">Tool to scan a network for alive host (!nbtscan - Host Lookup).</h4>
<div style="text-align:center">
    <div style="align:center">
    <img src="https://img.shields.io/badge/Author-sc4rfurry-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Author">
    <img src="https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Version">
    <img src="https://img.shields.io/badge/Go_Version-1.18.1-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17" alt="Go Version">
    <img src="https://img.shields.io/badge/OS-Linux-informational?style=flat-square&logo=ubuntu&logoColor=green&color=5194f0&bgcolor=110d17" alt="OS">
    <img src="https://img.shields.io/badge/Go_Library-wafme0w-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17" alt="Go Library">
    </div>
</div>

#

## Table of Contents

- [Installation](#installation)
- [Features](#features)
- [Running TallusX](#running-tallusx)
- [Options](#options)
- [Building TallusX](#building-tallusx)
- [Pre-Compiled Binaries](#pre-compiled-binaries)
- [References](#references)
- [Contributing](#contributing)
- [License](#license)


#

### 🔧 Technologies & Tools

![](https://img.shields.io/badge/Editor-VS_Code-informational?style=flat-square&logo=visual-studio&logoColor=blue&color=5194f0)
![](https://img.shields.io/badge/Language-Go-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17)
![](https://img.shields.io/badge/Go_Version-1.18.1-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17)

#

### 📚 Requirements
> - Go 18.1 linux/amd64

#
### Installation

- sudo apt-get update && sudo apt-get golang
- git clone https://github.com/sc4rfurry/T4llusX.git
- cd T4llusX
- go get .
- go build main.go
    - or use the `builder.sh` script to build the tool.


### Features

- Can scan a network range for Alive hosts
- Static Binary (No Dependencies almost :)
#

## Running TallusX
```sh
go run main.go --help
go run main.go <subnet> [options]
```


### Example

To run the tool on a target, just use the following command.

```console
go run main.go 192.168.0 -d
```
#

## Options
```sh
	Usage: ./main <subnet> [options]
__________________________________________________________________________________

	Options: 
		subnet 		Network to scan. (Ex: 192.168.0)
		-d/--debug 	Enable debug mode
		-V/--version 	Show version
		-h/--help 	Show this help menu


[info]  Subnet: First three octets. Ex: ./main 192.168.0
```

## Building TallusX
> To build the tool, you can use the following command.
```sh
env GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w -extldflags "-static"' -o TallusX main.go
```

> You can also use the bultin Bash script to build the tool.

- Before running the script, make sure to give it execution permissions.
- The bash script can build both Linux and Windows binaries.
- Binaries will be Stripped and Compressed. (lolcat, strip and upx are required)
```sh
chmod +x builder.sh
./builder.sh main.go
```
#
## Pre-Compiled Binaries
<div>
<div style="text-align:center">
    <a href="https://github.com/sc4rfurry/T4llusX/releases/tag/v1.0.0">
    <img src="https://img.shields.io/badge/Download-v1.0.0-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Download">
    </a>
<div style="text-align:center">
    <img src="https://img.shields.io/badge/Status-Active-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Status">
</div>
</div>
</div>

#

## References
* Special thanks to the following projects (they are the base of this tool):
    - [goHackTools](https://github.com/dreddsa5dies/goHackTools)

#

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

[def]: https://img.shields.io/badge/OS-Linux-informational?style=flat-square&logo=ubuntu&logoColor=green&color=5194f0&bgcolor=110d17