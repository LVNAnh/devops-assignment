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
## **Jenkinsfile**

1. **Môi Trường (Environment Variables)**  
   - Sử dụng biến môi trường DOCKER_IMAGE để định nghĩa tên image cho Backend.
   - Sử dụng PROD_SERVER để cấu hình địa chỉ IP cho production của dự án.

2. **Các Giai Đoạn (Stages)**  
   - **Clone Repository:** Sao chép source code từ nhánh `master` trên GitHub.  
   - **Build Docker Images:** Xây dựng image Docker cho Backend.  
   - **Push Docker Images:** Đẩy image Docker lên DockerHub.  
   - **Deploy to DEV:** Triển khai Backend trên môi trường phát triển (DEV) bằng Docker.  

3. **Quản Lý Container và Mạng**  
   - Tạo mạng Docker (`dev`) nếu chưa tồn tại.  
   - Dừng và khởi chạy lại container Backend.  

5. **Thông Báo Trạng Thái Build**  
   - Sử dụng API Telegram để gửi thông báo kết quả xây dựng backend (thành công/thất bại).  

6. **Dọn Dẹp Workspace**  
   - Luôn dọn dẹp workspace sau mỗi pipeline (`cleanWs`).  

---  

## **Cách xây dựng dự án**

###  **1. Chạy với Docker Compose**

1. **Clone dự án:**
   ```bash
   git clone https://github.com/LVNAnh/devops-assignment.git
   cd task-manager
   ```

2. **Chạy Docker Compose:**
   ```bash
   docker-compose up --build
   ```

3. **Truy cập ứng dụng:**
   - Backend: (http://localhost:3005)  

---

