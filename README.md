# 抖音极简版

抖音极简版是 2022 年字节跳动青训营的结营项目。

## 技术栈

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://github.com/go-gorm/gorm)
- [MySQL](https://www.mysql.com/)
- [JWT](https://github.com/golang-jwt/jwt)
- [乐观锁](https://github.com/go-gorm/optimisticlock)
- [FFmpeg](https://ffmpeg.org/)

## 文件结构

```text      
│  config.yaml        ──── 配置文件
│
├─app
│  │  main.go         ──── 项目入口
│  │  router.go       ──── API路由配置
│  │
│  ├─api              ──── [API层]
│  │    responses.go  ──── 响应结构体
│  │    user.go       ──── 注册、登录、用户信息
│  │    relation.go   ──── 关注、粉丝
│  │    feed.go       ──── 视频流
│  │    publish.go    ──── 视频投稿
│  │    comment.go    ──── 视频评论
│  │    favorite.go   ──── 视频点赞
│  │
│  ├─config           ──── [项目配置]
│  │
│  ├─dao              ──── [持久化层]
│  │    db.go         ──── 数据库初始化
│  │
│  ├─errs             ──── [错误管理]
│  │
│  └─service          ──── [服务层]
│
├─pkg                 ──── [公共依赖]
│
├─test                ──── [测试相关]
│
└─docs                ──── [项目文档]
```

## 运行

项目配置文件为项目根目录的 [config.yaml](./config.yaml) 文件。启动项目前需做以下准备工作：

- 安装 Go
- 安装 [FFmpeg](https://ffmpeg.org/download.html)
- 启动 MySQL 8.0 并确保有一个名为 `douyin_db` 的数据库
- 在 [config.yaml](./config.yaml) 文件中配置好数据库的地址和账号密码

然后执行此命令即可运行：

```shell
go run ./app
```

若想编译可执行文件，请执行命令：

```shell
go build -o ./bin/app ./app
```

这会将可执行文件保存到 `./bin` 目录。

## 测试

测试环境的配置文件为 [test/config.yaml](./test/config.yaml) ，测试前请做好配置工作。

## 项目文档

- [开发环境配置](./docs/development-setup.md)
- [Git 使用规范](./docs/git-standard.md)

## 问题

由于本项目的 API 设计是训练营安排的且不可修改，存在的以下问题我们不考虑解决：

- 注册和登录时，通过 URL 参数传密码存在安全问题
- 通过 URL 参数传鉴权 Token 存在安全问题
- 部分列表查询没有分页
