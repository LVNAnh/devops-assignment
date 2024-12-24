# **Web Management Project**

## Giới thiệu
---
Dự án Web Management là dự án được xây dựng dùng để quản lý các công việc cơ bản của 1 admin web như quản lý kho hay quản lý người dùng với đầy đủ các phương thức CRUD, các thành phần của dự án bao gồm:
- **Backend**: Golang (Gin Framework)
- **Database**: Cloud mongodb
- **Docker**: Triển khai web với Docker và Docker compose
- **Jenkinsfile**: Thực hiện CI/CD cho web
---
## Cấu trúc dự án

```plaintext
devops-assignment/
├── config/
│   └── database.go
├── controllers/
│   ├── product_controller.go
│   └── user_controller.go
├── models/
│   ├── product.go
│   └── user.go
├── routes/
│   ├── product_routes.go
│   └── user_routes.go
├── docker-compose.yml
├── Dockerfile
├── Jenkinsfile
├── go.mod
├── go.sum
├── main.go
└── README.md
```
---
