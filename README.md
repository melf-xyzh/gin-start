# gin-start

### 介绍
一个Gin框架的项目模板

### 软件架构
> 语言：Golang
>
> Web框架：Gin
>
> ORM：GORM 
>
> 数据库：MySQL
>
> 缓存：Redis

### 项目目录

```
├─config			配置文件
├─connection		数据库连接池
├─controller		Controller层
├─middleware		中间件
├─models			模型实体
├─pkg				工具包
├─prd				项目文档，ER图
├─routers			路由相关
│  └─urlmap			url，tmpl命名映射
├─service			Service层
├─static			静态文件
├─templates			Go模板文件
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

### V 0.0.3

- 独立出数据库连接池、Redis连接池
- 引入Session，实现身份验证中间件
- 实现登录界面

### V 0.0.2

- 修改配置文件方式为Json

### v 0.0.1

- 初始化项目
- 配置数据库连接池
- 配置路由

### 备注

> 暂无
