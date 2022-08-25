package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleStaticHelloWorld(c *gin.Context) {
	c.HTML(http.StatusOK, "hello-world.tmpl", nil)
}

func handleDynamicHello(c *gin.Context) {
	name := c.Param("name")
	c.HTML(http.StatusOK, "dynamic-hello.tmpl", gin.H{
		"name": name,
	})
}

func handlePageAllEmployees(c *gin.Context) {
	c.HTML(http.StatusOK, "employee-list.tmpl", gin.H{
		"employeeList": employeeList,
	})
}
