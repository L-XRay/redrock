package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type All struct {
	ID int `json:"id"`
	Name string  `json:"title"`
	State bool  `json:"status"`
}

var(
	db *gorm.DB
)
func MySQL()(err error){
	flag:="root:2020215059@tcp(127.0.0.1:3306)/remember?charset=utf8mb4&parseTime=True&loc=Local"
	db,err= gorm.Open("mysql",flag)
	if err!=nil{
		return
	}
	return db.DB().Ping()
}
func main() {
	err:=MySQL()
	if err!=nil{
		panic(err)
	}
	defer db.Close() //程序退出关闭数据库

	router := gin.Default()

	db.AutoMigrate(&All{})//自动迁移

	router.Static("/static","static")//静态文件寻址

	router.LoadHTMLGlob("templates/*")//模板index寻址

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)

	})

	apiGroup := router.Group("v1")
	{
		apiGroup.POST("/todo", func(c *gin.Context) {
			var todo All
			c.BindJSON(&todo)
			err=db.Create(&todo).Error
			if err!=nil{
				return
			}else{
				c.JSON(http.StatusOK,todo)
			}
		})//添加
		apiGroup.GET("/todo",func(c *gin.Context){
			var todolist []All
			err=db.Find(&todolist).Error
			if err!=nil{
				return
			}else{
				c.JSON(http.StatusOK,todolist)
		    }
		})//查看
		apiGroup.PUT("/todo/:id",func(c *gin.Context){
			var all All
			id,_:=c.Params.Get("id")
			db.Where("id=?",id).First(&all)
			c.BindJSON(&all)
			db.Save(&all)
		})//修改
		apiGroup.DELETE("/todo/:id",func(c *gin.Context){
			id,_:=c.Params.Get("id")
			db.Where("id=?",id).Delete(All{})
		})//删除
	}
	router.Run(":0124")
}