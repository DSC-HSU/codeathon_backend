package submission

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/endpoint"
	"github.com/gin-gonic/gin"
)

func Api(service SubmissionService) {
	r := endpoint.GetEngine()

	r.GET("/submission", func(c *gin.Context) {

		//get challenge id and user id from query
		challengeId := c.Query("challenge_id")
		userId := c.Query("user_id")

		submission, err := service.GetByChallengeIdAndUserId(c, userId, challengeId)
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
