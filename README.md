# iris-demo

## 描述
iris + xorm 框架 demo

## 版本
go 1.13

## 结构
```cassandraql
conf  配置文件
controllers  控制器 接受参数 api的入口
datasource 数据库配置
models  结构体模型
repo 数据库的操作
route  注册路由
service 业务逻辑代码
utils  工具类
config.json 配置文件的映射
go.mod mod文件
main.go 主程序入口
```

## 部署
```cassandraql
go build iris-demo
```

## 打包
```cassandraql
# mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
# linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
# windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

## 历史

### 2019-09-19
1. 增加user 增删改查接口
```
[DBUG] 2019/09/19 13:20 DELETE: /user -> controllers.UserController.Delete() and 2 more
[DBUG] 2019/09/19 13:20 GET: /user/all -> controllers.UserController.GetAll() and 2 more
[DBUG] 2019/09/19 13:20 GET: /user -> controllers.UserController.Get() and 2 more
[DBUG] 2019/09/19 13:20 POST: /user/edit -> controllers.UserController.PostEdit() and 2 more
[DBUG] 2019/09/19 13:20 POST: /user/save -> controllers.UserController.PostSave() and 2 more
```


### 2019-09-18
1. 代码上传
2. 固定目录结构
3. 增加ymal配置文件
4. 增加数据库工具类
5. 增加models 实体