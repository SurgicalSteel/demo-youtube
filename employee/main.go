package main

import (
	"github.com/gin-gonic/gin"
)

const (
	ginMode = "release"
)

func main() {
	gin.SetMode(ginMode)

	defaultEngine := gin.Default()

	apiV1 := defaultEngine.Group("/api/v1")
	{
		apiV1.GET("/hello-world", handleHelloWorld)
		apiV1.POST("/employee", handleInsertEmployee)
		apiV1.GET("/employees", handleGetAllEmployees)
		apiV1.GET("/employee/:id", handleGetEmployeeByID)
	}

	defaultEngine.Run(":9090")
	//localhost:9090/api/v1/hello-world
}

func handleHelloWorld(c *gin.Context) {
	c.Header("Content-Type", "text")

	message := "Hello World!"
	rawByte := []byte(message)

	c.Writer.Write(rawByte)
}
