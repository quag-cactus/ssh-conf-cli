# ssh-conf-cli
Cross-Platform Command Line Interface for editing ssh config file.

## Installation

1. You can download the binary file from [release page](https://github.com/quag-cactus/ssh-conf-cli/releases)

2. Unzip the file (`ssh-conf-cli_linux_amd64_<versions>.zip`) and put it in your `$PATH` directory as needed.
    * For linux
        ```bash
        $ unzip ssh-conf-cli_linux_amd64.zip -d ~/.local/bin
        ```
    * For Windows
        Unzip the file (`ssh-conf-cli_win_amd64_<versions>.zip`) using your favorite unzip tool and put the file `ssh-conf-cli.exe` in your `$PATH` directory.

## Usage

To get the version:
```bash
$ ssh-conf-cli version
```

To Edit value of any targeted-host:
```bash
$ ssh-conf-cli edit -T [target-pattern] <subcommand> 
```
For example, to edit the value of `HostName` configlated for host `remote-machine`:
```bash
$ ssh-conf-cli edit -T remote-machine hostName -n 0.0.0.0
```

## Source build

To make binary file for your platform:
```bash
go build -ldflags "-s -w -X github.com/quag-cactus/ssh-conf-cli/cmd.version=<VERSION_STRING>" -trimpath ./...
```