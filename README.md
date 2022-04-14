# rpi-babycam

Turning raspberry pi into a babycam

## Preparation

### Install docker

    sudo apt-get update && sudo apt-get upgrade
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
    sudo usermod -aG docker pi

### Go

    wget https://go.dev/dl/go1.18.1.linux-armv6l.tar.gz
    sudo tar -C /usr/local -xzf go1.18.1.linux-armv6l.tar.gz
    rm go1.18.1.linux-armv6l.tar.gz

    nano ~/.profile

> PATH=$PATH:/usr/local/go/bin

> GOPATH=$HOME/go

    source ~/.profile