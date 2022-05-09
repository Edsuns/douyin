# 抖音极简版

抖音极简版是 2022 年字节跳动青训营的结营项目。

## 技术栈

- GO
- Gin
- GORM
- MySQL
- JWT

## 项目文件介绍

```text
├─app
│  │  main.go         ──── 项目入口
│  │  router.go       ──── API路由配置
│  │
│  ├─api              ──── [API层]
│  │    responses.go  ──── 响应结构体
│  │
│  ├─config           ──── [项目配置]
│  │
│  ├─dao              ──── [持久化层]
│  │    db.go         ──── 数据库初始化
│  │    user.go       ──── 注册、登录、查询用户信息
│  │
│  ├─errs             ──── [错误管理]
│  │
│  └─service          ──── [服务层]
│
├─pkg                 ──── [公共依赖]
│
└─test                ──── [测试相关]
```

## 配置与运行

项目配置文件为项目根目录的 [config.yaml](./config.yaml) 文件。启动项目前需做以下准备工作：

- 启动 MySQL 8.0 并确保有一个名为 `douyin_db` 的数据库
- 在 [config.yaml](./config.yaml) 文件中配置好数据库的地址和账号密码
