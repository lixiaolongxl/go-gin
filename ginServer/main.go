package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建路由
	r := gin.Default()
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
			fmt.Printf("%#v\n", U)
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
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "index",
			})
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
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
