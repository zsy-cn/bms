要在`.proto`文件里使用import导入其他`.proto`文件不仅`protoc`编译命令要修改, 在vscode里也要加上相应配置, 不然代码提示总被`.proto`的错误阻挡, 干脆直接把`.proto`文件后缀去掉, 这样编译没问题, 代码提示也不存在, 只是少了代码高亮而已.