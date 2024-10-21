package submission

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/endpoint"
	"github.com/gin-gonic/gin"
)

func Api(service SubmissionService) {
	r := endpoint.GetEngine()

	r.GET("/submission/:id", func(c *gin.Context) {
		id := c.Param("id")
		submission, err := service.GetById(c, id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, submission)
	})

	r.POST("/submission", func(c *gin.Context) {
		submission := &domain.Submission{}
		err := c.BindJSON(submission)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = service.Create(c, submission)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, submission)
	})

	r.GET("/submission/list/:id", func(c *gin.Context) {
		id := c.Param("id")
		submissions, err := service.List(c, id, &domain.ListOpts{})
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, submissions)
	})

	r.PUT("/submission", func(c *gin.Context) {
		submission := &domain.Submission{}
		err := c.BindJSON(submission)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = service.Update(c, submission)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, submission)
	})
}
