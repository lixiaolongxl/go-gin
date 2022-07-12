package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"time"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
func AccessLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		str := fmt.Sprintf("url=%s, status=%d, resp=%s \n", c.Request.URL, c.Writer.Status(), blw.body.String())
		fmt.Fprint(gin.DefaultWriter, "<====================================================================>\n")
		fmt.Fprint(gin.DefaultWriter, str)
	}
}

//func MiddleWare(c *gin.Context) {
//	t := time.Now()
//	fmt.Println("中间件开始执行了")
//	// 设置变量到Context的key中，可以通过Get()取
//	c.Set("request", "中间件")
//	// 执行函数
//	c.Next()
//	// 中间件执行完后续的一些事情
//	status := c.Writer.Status()
//	fmt.Println("中间件执行完毕", status)
//	t2 := time.Since(t)
//	fmt.Println("time:", t2)
//}

func MiddleWare(doChek bool) gin.HandlerFunc {
	//可以做一些查询数据库逻辑等
	return func(c *gin.Context) {
		if doChek {
			t := time.Now()
			fmt.Println("中间件开始执行了")
			// 设置变量到Context的key中，可以通过Get()取
			c.Set("request", "中间件")
			c.Next()
			status := c.Writer.Status()
			fmt.Println("中间件执行完毕", status)
			t2 := time.Since(t)
			fmt.Println("time:", t2)

			//fmt.Fprint(gin.DefaultWriter, &c.Writer)
		} else {
			c.Next()
		}

	}
}
func m2(c *gin.Context) {
	//
	//c.Next()  // 执行后续函数
	//c.Abort() //  停止执行后序函数

	fmt.Println("time:", "mmmmm in ")
}

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key;column:beast_id"`
	Name     string
	Age      int64
}

//修改表名s
func (Animal) TableName() string {
	return "profile"
}
func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	dsn := "root:root@tcp(127.0.0.1:3306)/orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移
	db.AutoMigrate(&UserInfo{}, &User{}, &Animal{})

	//创建记录

	//u1 := UserInfo{1, "七米", "男", "篮球"}
	//u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	//db.Debug().Create(&u1)
	//db.Create(&u2)
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 1.创建路由
	r := gin.Default()
	//r := gin.New()
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	//}

	r.Use(MiddleWare(true))

	r.Use(AccessLogHandler())
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		//c.String(http.StatusOK, "hello World!")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello word",
		})
	})
	type Data struct {
		Name string `json:"name"`
		Age  int
		Sex  string
	}
	r.GET("index", func(context *gin.Context) {
		context.JSON(http.StatusOK, Data{
			Name: "李小龙",
			Age:  18,
			Sex:  "男",
		})
	})
	//r.POST("post", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"MD": "post",
	//	})
	//})
	//GET  接收参数
	r.GET("web", func(context *gin.Context) {
		//name := context.Query("name")
		//name := context.DefaultQuery("name", "李小龙")

		a := context.QueryMap("a")
		name, ok := context.GetQuery("name")
		if !ok {
			name = "李小龙"

		}
		context.JSON(http.StatusOK, gin.H{
			"message": name,
			"a":       a,
		})
	})
	//post 参数
	r.POST("post", func(context *gin.Context) {

		//name := context.PostForm("name")
		name, _ := context.GetPostForm("name")
		//if !ok {
		//
		//}
		//name := context.DefaultPostForm("name", "234")
		age := context.PostForm("age")
		context.JSON(http.StatusOK, gin.H{
			"message": 200,
			"name":    name,
			"age":     age,
		})
	})

	r.POST("json/:age", func(context *gin.Context) {

		age := context.Param("age")
		json := make(map[string]interface{}) //注意该结构接受的内容
		context.BindJSON(&json)
		context.JSON(200, gin.H{
			"name": json["name"],
			"age":  age,
		})
	})

	//path
	r.GET("path/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	type UserInfo struct {
		UserName string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}
	r.GET("user", func(context *gin.Context) {
		var U UserInfo
		err := context.BindQuery(&U)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			//fmt.Printf("%#v\n", U)
			context.JSON(200, gin.H{
				"Password": U.Password,
				"UserName": U.UserName,
			})
		}
	})

	r.POST("user", func(context *gin.Context) {
		var U UserInfo
		err := context.ShouldBind(&U) //接收post 参数的
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

		} else {
			fmt.Printf("%#v\n", U)
			context.JSON(200, gin.H{
				"message": U,
			})
		}
	})
	//重定向
	r.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.sogo.com")
	})
	//路由重定向
	r.GET("/a", func(context *gin.Context) {
		context.Request.URL.Path = "/index" //修改url
		r.HandleContext(context)            //继续后续处理
	})

	//路由组
	//shopGroup := r.Group("/shop", m2)
	shopGroup := r.Group("/shop")
	shopGroup.Use(m2)
	{
		shopGroup.GET("index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "index",
			})
			//go 中必须使用context 的copy
			//go funcxx(context.Copy())
		})
		shopGroup.GET("xx", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "xx",
			})
		})
		shopGroup.GET("oo", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "oo",
			})
		})
	}

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "path not fount",
		})
	})
	//中间建

	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
