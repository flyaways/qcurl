#!/usr/bin/env bash

go build

# docker run -d --name srs3 -p 9935:1935 -p 9985:1985 -p 9080:8080 registry.cn-hangzhou.aliyuncs.com/ossrs/srs:3

# nohup ffmpeg -stream_loop -1 -re -i $HOME/minio/video/_Twelve_\ Ep.\ 1\ of\ 7.mp4 -c copy -f flv "rtmp://10.64.7.106:9935/quic.video.io/live/surfing" &
