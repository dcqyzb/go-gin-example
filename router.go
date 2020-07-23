package main

import (
	. "ginDemo/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)
	router.POST("/AddPerson", AddPersonApi)  //添加数据http://localhost:8080/AddPerson
	router.GET("/persons", GetPersonsApi)    //查询所有数据http://localhost:8080/persons
	router.GET("/person/:id", GetPersonById) //查询单条记录根据id http://localhost:8080/person/2
	return router
}
