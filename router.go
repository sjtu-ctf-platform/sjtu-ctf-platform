package main

import "github.com/gin-gonic/gin"

func Router(c *gin.Engine) {
	user := c.Group("/user")
	userHandler(user)

	problem := c.Group("/problem")
	problemHandler(problem)

	sandbox := c.Group("/sandbox")
	sandboxHandler(sandbox)

	file := c.Group("/file")
	fileHandler(file)

	score := c.Group("/score")
	scoreHandler(score)

	manage := c.Group("/manage")
	manageHandler(manage)
}

func userHandler(r *gin.RouterGroup) {
	r.POST("/login")
	r.POST("/register")

	profile := r.Group("/profile")

	profile.GET("/:userID")   // basic information
	profile.GET("/:userID/avatar") // user's avatar
	profile.POST("/:userID")
	profile.POST("/:userID/avatar")

	profile.GET("/finishedProblem") // user's submit records(only problems got scores)
}

func problemHandler(r *gin.RouterGroup) {

}

func manageHandler(r *gin.RouterGroup) {

}

func fileHandler(r *gin.RouterGroup) {

}

func sandboxHandler(r *gin.RouterGroup) {

}

func scoreHandler(r *gin.RouterGroup) {
	
}

