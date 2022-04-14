# rpi-babycam

Turning raspberry pi into a babycam

## Preparation

### Install docker

    sudo apt-get update && sudo apt-get upgrade
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
    sudo usermod -aG docker pi

### docker-compose
    sudo curl -L "https://github.com/docker/compose/releases/download/v2.4.1/docker-compose-linux-armv6" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose

### Go

    wget https://go.dev/dl/go1.18.1.linux-armv6l.tar.gz
    sudo tar -C /usr/local -xzf go1.18.1.linux-armv6l.tar.gz
    rm go1.18.1.linux-armv6l.tar.gz

    nano ~/.profile

> PATH=$PATH:/usr/local/go/bin

> GOPATH=$HOME/go

    source ~/.profile

## GStreamer

sudo docker run -d --name gstreamer-videotestsrc pavelis/gstreamer:latest bash -c 'gst-launch-1.0 -v v4l2src device=/dev/video0 ! video/x-raw,width=640,height=480,framerate=30/1 ! x264enc tune=zerolatency bitrate=500 speed-preset=superfast ! rtph264pay ! application/x-rtp, media=video, encoding-name="H264", payload=96 ! udpsink port=5004 host=janus.teleop.be.continental.cloud'