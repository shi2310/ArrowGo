ArrowGo 脚手架

# 框架说明

golang+gomod+gin+gorm+go-jwt+gin-swagger

---

# 使用说明

## 生成 swagger 文档

```
swag init
```

如果使用的 models 中模型作为数据传输对象（DTO），因为其中包含 gorm.Model，会造成 swag 无法识别"cannot find type definition: gorm.Model"，此时需要执行

```
swag init --parseDependency --parseInternal
```

但是上述解析需要大量时间，不建议使用

## 运行项目

```
go run main.go
```

## Docker 部署

1. win10 交叉编译 linux

```
$env:GOOS="linux"
$env:GOARCH="amd64"
go build
```

2. 根据 Dockerfile 生成镜像

```
docker build --rm -t arrowgo:v1 .
```

3. 生成容器

```
docker run -p 48066:8044 --name arrow_api -d --restart=always -v ${PWD}/upload:/arrowgo/upload -v /etc/localtime:/etc/localtime arrowgo:v1
```
