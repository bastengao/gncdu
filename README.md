# gncdu

gncdu implements [NCurses Disk Usage](https://dev.yorhel.nl/ncdu)(ncdu) with golang, and is at leaset twice faster as ncdu.

## Install

### Binaries

macOS (Apple Silicon)

    wget -O https://github.com/bastengao/gncdu/releases/download/v0.7.0/gncdu-darwin-arm64 && chmod +x /usr/local/bin/gncdu

Linux (amd64)

    wget -O /usr/local/bin/gncdu https://github.com/bastengao/gncdu/releases/download/v0.7.0/gncdu-linux-amd64 && chmod +x /usr/local/bin/gncdu

Or download executable file from [releases](https://github.com/bastengao/gncdu/releases) page.

### Install from source

    go install github.com/bastengao/gncdu

## Usage

    gncdu [path]

![screenshot](http://bastengao.com/images/others/gncdu-screenshot-v0.7.0.png)
