

- 项目初始化
mkdir -p go-gin-gorm-demo/ && cd go-gin-gorm-demo
go mod init go-gin-gorm-demo
 
- 安装依赖
// 设置代理
GOPROXY=https://goproxy.io,direct
// gin
go get -u github.com/gin-gonic/gin
// gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u gorm.io/driver/sqlite
// mysql
mysql -uroot -p123456 -h 10.0.0.91 -P 3306
mysql> CREATE DATABASE `curd-list` DEFAULT CHARACTER SET utf8mb4;

- 开始操作
touch main.go

- 实现CURD接口

- git操作
git clone https://github.com/hu417/gin-vue3-element-plus-curd.git
git init
git config --global user.name "***"
git config --global user.email ****@qq.com

// ssl认证关闭
git config --global http.sslVerify "false"
git config --global credential.helper manager

// 提交代码
echo "# CURD 小demo" >> README.md
git add .
git commit -m "fix: gin-vue-elementplus-curd项目
1、后端CURD的实现
" 
git branch -M main
git remote add origin https://github.com/hu417/gin-vue3-element-plus-curd.git
git branch
git tag -a v1.11 -m "版本v1.11"
git push -u origin main --tags
git log

