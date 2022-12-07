package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	c := gin.Default()

	Run(c)
}

func Run(c *gin.Engine){

	Router(c)

	if err := c.Run("127.0.0.1:8000"); err != nil {
		fmt.Println(err.Error())
	}
}