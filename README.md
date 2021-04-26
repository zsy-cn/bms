
- [Backend Manger System(BMS)](#backend-manger-systembms)
  - [Overview](#overview)
  - [File catalog](#file-catalog)
  - [How to deploy of each package](#how-to-deploy-of-each-package)
  - [How to run tests of each package](#how-to-run-tests-of-each-package)


# Backend Manger System(BMS)

## Overview 

Most of packages here provide a layer upon database/elasticsearch. to serve as a backend API for building UI. each package have a `proto` package inside that as outer appearance of the package, normally includes data structs defined in Protocol Buffers, and API interfaces, and errors definations. and a `New` or several `NewXXXService` funcs to initialize those services. except these definiations you shouldn't need to read other codes to use it.

## File catalog

```
.
├── conf  // 配置文件目录, 包括各个服务的地址, 监听端口, redis, consul, db连接配置等信息.
├── deploy // 部署配置文件, 包括docker-compose本地部署及k8s环境部署
├── docs // 文档
├── environ_monitor
├── geomagnetic
├── loraclient // 连接lora-app-server的工具服务, 可完成添加客户时的各种资源(user, orgnization, service profile, app)初始化, 及设备增删操作.
├── manhole_cover
├── loragateway // loraWAN 网关程序
├── model // 数据库ORM模型
├── protos // protobuf声明文件
├── server // webserver, 提供对前端的API服务
│   ├── config
│   ├── controller
│   │   ├── customer
│   │   ├── devicehub
│   │   ├── devicemodel
│   │   ├── devicetype
│   │   ├── group
│   │   ├── manager
│   │   ├── manufacturer
│   │   └── parserhub
│   └── router
├── service // 各种子服务
│   ├── core
│   │   ├── base
│   │   │   ├── endpoint
│   │   │   ├── service
│   │   │   └── transport
│   │   ├── devicehub
│   │   │   ├── endpoint
│   │   │   ├── service
│   │   │   └── transport
│   │   └── parserhub
│   │       ├── endpoint
│   │       ├── service
│   │       └── transport
│   ├── device
│   │   ├── camera
│   │   └── sensor
│   ├── parser
│   │   ├── dici_weichuan_01
│   │   │   └── cmd
│   │   ├── env_yingfeichi_01
│   │   │   └── cmd
│   │   ├── sos_yingfeichi_01
│   │   │   └── cmd
│   │   └── trashcan_lierda_01
│   │       └── cmd
│   └── user
│       ├── contact
│       ├── customer
│       └── manager
├── shanghaiAPP
├── smoke
├── sos
├── util
└── water_level
```

## How to deploy of each package

Use yaml

## How to run tests of each package

Normally you can do:

Start all dependencies server of the package, normally a database, sometimes elastic search

```
$ docker-compose up
```

Start a new terminal, exports all the configuration values that points to the servers above
```
$ source dev_env
```

Automatically run tests when any file changes

```
$ modd
```