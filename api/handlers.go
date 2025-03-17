package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func forcastnext15days(c *gin.Context){
	
	location := c.Param("location")

	data, err := getForcastNext15Days(location)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, data)

}