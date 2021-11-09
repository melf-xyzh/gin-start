# gin-start

### 介绍
一个Gin框架的项目模板

### 软件架构
> 语言：Golang
>
> Web框架：Gin
>
> ORM：GORM 

### 项目目录

```
├─config			配置文件
├─controller		Controller层
├─middleware		中间件
├─models			模型实体
├─pkg				工具包
├─prd				项目文档，ER图
├─routers			路由相关
├─service			Service层
├─static			静态文件
├─templates			Go模板文件
│  └─index
├─upload			文件上传目录
└─views				Views层
│ .gitignore        git
│ go.mod            go mod  
| main.exe          Windows执行文件
│ main.go           main入口
│ README.md    
```

### 启动

```bash
go run main.go
```

### 打包部署

```
go build -o main.exe main.go
```

### 更新日志

### v 0.0.1

- 初始化项目
- 配置数据库连接池
- 配置路由

### 备注

> 暂无
