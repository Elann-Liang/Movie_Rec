#!/bin/zsh

docker run -p 8080:5000 -d --rm -it -v `pwd`:/app go_1.2_image:latest


