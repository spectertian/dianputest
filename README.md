# dianputest

Golang 商铺管理后台 REST API，基于 **Gin + GORM + MySQL + JWT** 实现。

---

## 🚀 项目简介

提供完整的商铺管理后台接口，包含用户注册登录认证（JWT）和商铺增删改查功能。

---

## 📦 依赖安装

```bash
go mod tidy
```

---

## ⚙️ 数据库配置

默认数据库连接信息（可通过环境变量覆盖）：

| 环境变量 | 默认值 | 说明 |
|----------|--------|------|
| `DSN` | `root:123456@tcp(127.0.0.1:3306)/dianpu?charset=utf8mb4&parseTime=True&loc=Local` | MySQL 连接串 |
| `JWT_SECRET` | `dianpu-secret-key` | JWT 签名密钥 |
| `PORT` | `:8080` | 监听端口 |

启动前需先创建数据库：

```sql
CREATE DATABASE dianpu CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

---

## ▶️ 启动方式

```bash
go run main.go
```

---

## 📡 接口文档

### 🔓 公开接口（无需 Token）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/register` | 用户注册 |
| POST | `/api/login` | 用户登录，返回 JWT Token |

### 🔒 认证接口（Header: `Authorization: Bearer <token>`）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/logout` | 登出（客户端清除 Token） |
| GET | `/api/shops` | 商铺列表（支持 `page`、`page_size`、`name` 查询参数） |
| GET | `/api/shops/:id` | 获取单个商铺详情 |
| POST | `/api/shops` | 新增商铺 |
| PUT | `/api/shops/:id` | 编辑商铺 |
| DELETE | `/api/shops/:id` | 删除商铺（软删除） |

### 请求/响应示例

#### POST /api/login
```json
// 请求
{ "username": "admin", "password": "123456" }

// 响应
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": { "id": 1, "username": "admin", "nickname": "管理员" }
  }
}
```

#### GET /api/shops?page=1&page_size=10&name=咖啡
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [...],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

#### POST /api/shops
```json
// 请求
{
  "name": "星巴克",
  "image": "https://example.com/img.jpg",
  "address": "北京市朝阳区",
  "phone": "010-12345678",
  "description": "精品咖啡",
  "status": 1
}

// 响应
{ "code": 200, "message": "创建成功", "data": { ... } }
```

---

## 📁 目录结构

```
dianputest/
├── main.go                      # 入口文件
├── go.mod / go.sum
├── config/config.go             # 配置（数据库、JWT secret、端口）
├── database/db.go               # GORM 数据库初始化
├── model/
│   ├── user.go                  # 用户模型
│   └── shop.go                  # 商铺模型
├── middleware/auth.go           # JWT 认证中间件
├── controller/
│   ├── auth_controller.go       # 登录/注册/登出控制器
│   └── shop_controller.go       # 商铺 CRUD 控制器
├── service/
│   ├── auth_service.go          # 登录业务逻辑
│   └── shop_service.go          # 商铺业务逻辑
├── router/router.go             # 路由注册
└── utils/
    ├── jwt.go                   # JWT 生成与解析
    └── response.go              # 统一响应格式
```