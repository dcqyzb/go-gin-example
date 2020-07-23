package apis

import (
	"fmt"
	"ginDemo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := models.Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func GetPersonsApi(c *gin.Context) {
	p := new(models.Person)
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{"persons": persons})
}

func GetPersonById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	p := new(models.Person)
	person, err := p.GetPersonById(id)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{"person": person})
}
