#!/bin/bash
## 使用方法: 当前目录直接执行./build.sh即可
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/common
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/loraclient
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/contact
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/customer
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/group
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/device
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/device_sensor
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/core
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/manager
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/session
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/parserhub
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/device_type
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/device_model
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/manufacturer
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/parser
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/camera
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/device_management
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/message_management
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/geomagnetic
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/environ_monitor
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/sound_box
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/trashcan
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/manhole_cover
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/water_level
protoc -I $(pwd) -I $GOPATH/src --go_out=plugins=grpc:$(pwd) $(pwd)/wifi
