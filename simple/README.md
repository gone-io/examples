# examples/simple

## 本demo依赖了redis，需要正确配置redis才能正常运行

## 在当前目录提供了一个docker-compose，可以方便在docker中运行一个redis服务(如果你本地安装了docker)，命令如下：

```shell
docker-comopse up -d
```

## 使用docker-compose提供的redis运行

```shell
# 启动docker-compose中的redis服务
docker-compose up -d

#运行代码
go run main.go

#定制docker-compose中的redis服务
docker-compose down
```