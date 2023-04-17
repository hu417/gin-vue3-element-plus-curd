package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 表字段映射结构体
type List struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20);not null" json:"name" binding:"required"`
	Status string `gorm:"type:varchar(20);not null" json:"status" binding:"required"`
	Phone  string `gorm:"type:varchar(20);not null" json:"phone" binding:"required"`
	Email  string `gorm:"type:varchar(20);not null" json:"email" binding:"required"`
	City   string `gorm:"type:varchar(20);not null" json:"city" binding:"required"`
	// gorm: 字段类型; json数据对应字段; binding请求绑定参数
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(10.0.0.91:3306)/curd-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 禁止表名变成复数(s)
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("数据库连接错误,err= %s", err)
	}

	// 连接池
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Print(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 自动迁移创建表
	db.AutoMigrate(&List{})

	// gin请求
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// user接口
	userGroup := r.Group("user")
	{
		// 增
		userGroup.POST("/add", func(c *gin.Context) {
			var data List
			err := c.ShouldBindJSON(&data)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    500,
					"message": fmt.Sprintf("添加失败, err => %s", err.Error()),
				})
				return
			} else {
				num := db.Create(&data).RowsAffected
				fmt.Printf("影响的条数: %d", num)
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "添加成功",
				})
			}
		})

		// 查  //
		// 条件查询
		userGroup.GET("/list/:name", func(c *gin.Context) {
			var data []List
			name := c.Param("name")
			db.Where("name = ?", name).Find(&data)
			if len(data) == 0 {
				c.JSON(http.StatusOK, gin.H{
					"code":    400,
					"message": fmt.Sprintf("查询失败, %s 不存在", name),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "查询成功",
					"data":    data,
				})
			}
		})
		// 全部查询
		userGroup.GET("/list", func(c *gin.Context) {
			var data []List
			// 分页查询
			pageSize, _ := strconv.Atoi(c.Query("pageSize"))
			pageNum, _ := strconv.Atoi(c.Query("pageNum"))
			// 总数
			var total int64
			// 判断参数值
			if pageNum == 0 {
				pageNum = -1
			}
			if pageSize == 0 {
				pageSize = -1
			}

			// 偏移数据
			offsetVal := (pageNum - 1) * pageSize
			if pageSize == -1 && pageNum == -1 {
				offsetVal = -1
			}
			db.Model(&data).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&data)
			if len(data) == 0 {
				c.JSON(http.StatusOK, gin.H{
					"code":    400,
					"message": "查询失败, 没有数据",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "查询成功",
					"data": gin.H{
						"list":     data,
						"total":    total,
						"pageSize": pageSize,
						"pageNum":  pageNum,
					},
				})
			}
		})

		// 改
		userGroup.PUT("/update/:id", func(c *gin.Context) {
			var data List
			// 绑定uri参数id
			id := c.Param("id")

			// 判断id是否存在
			db.Select("id").Where("id = ?", id).Find(&data)
			if data.ID == 0 {
				c.JSON(http.StatusOK, gin.H{
					"code":    400,
					"message": fmt.Sprintf("id = %s 没找到", id),
				})
			} else {
				// 更新数据
				err := c.ShouldBindJSON(&data)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"code":    400,
						"message": fmt.Sprintf("参数有误,err = %s", err),
					})
				} else {
					db.Where("id = ?", id).Updates(&data)
					c.JSON(http.StatusOK, gin.H{
						"code":    200,
						"message": "更新成功",
					})
				}
			}
		})
		// 删
		userGroup.DELETE("/delete/:id", func(c *gin.Context) {
			var data []List
			// 绑定uri参数id
			id := c.Param("id")

			// 判断id是否存在
			db.Where("id = ?", id).Find(&data)
			if len(data) == 0 {
				c.JSON(http.StatusOK, gin.H{
					"code":    400,
					"message": fmt.Sprintf("id = %s 没找到", id),
				})
			} else {
				// 删除数据
				db.Where("id = ?", id).Delete(&data)
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "删除成功",
				})
			}
		})
	}

	r.Run(":8090")

}
