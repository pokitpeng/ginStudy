# ginStudy

## 目录结构
```go
├── common                      //公共包
│   ├── database.go             
│   └── jwt.go                  
├── config                      //配置文件
│   └── application.yml
├── controller                  //控制器，主业务
│   └── UserController.go
├── docs                        //swagger生成的文档,swag init生成
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── dto                         //数据对象
│   └── user_dto.go
├── go.mod
├── go.sum
├── http-test                   //vscode插件 http测试文件
│   └── test-request.http
├── logger                      //日志模块
│   ├── logger.go
│   └── logger_test.go
├── main.go                     //主程序   
├── middleware                  
│   └── AuthMiddleware.go
├── model                       //gorm模型
│   └── user.go
├── README.md
├── response                    //响应封装
│   └── response.go
├── router                      //路由信息
│   └── router.go
└── util                        //工具包 
    └── util.go
```

## 运行
若没有docs文件夹需要使用swag生成接口文档: `swag init`

```bash
go build && ./ginStudy
```