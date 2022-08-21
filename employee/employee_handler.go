package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Sex         string `json:"sex"`
	Nationality string `json:"nationality"`
}

var employeeList []Employee

func handleInsertEmployee(c *gin.Context) {
	var id int = 0
	if len(employeeList) == 0 {
		id = 1
	} else {
		for _, employee := range employeeList {
			if employee.ID > id {
				id = employee.ID
			}
		}
		id += 1
	}

	var newEmployeeData Employee

	err := c.BindJSON(&newEmployeeData)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error Getting Request Body")
		return
	}

	newEmployeeData.ID = id
	employeeList = append(employeeList, newEmployeeData)

	log.Printf("Employee List : %+v\n", employeeList)

	c.Header("Content-Type", "text")
	c.JSON(http.StatusOK, "sukses cuy!")
}

func handleGetAllEmployees(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, employeeList)
}

func handleGetEmployeeByID(c *gin.Context) {
	rawID := c.Param("id")

	id, err := strconv.Atoi(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error Getting ID")
		return
	}

	var isFound bool = false
	var result Employee

	for _, employee := range employeeList {
		if employee.ID == id {
			isFound = true
			result = employee
			break
		}
	}

	if !isFound {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Employee with ID %d is not found", id))
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, result)

}
