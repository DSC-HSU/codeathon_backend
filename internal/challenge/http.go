package challenge

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/endpoint"
	"github.com/gin-gonic/gin"
	"io"
)

func Api(service ChallengeService) {
	r := endpoint.GetEngine()

	r.GET("/challenge/:id", func(c *gin.Context) {
		id := c.Param("id")
		challenge, err := service.GetById(c, id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, challenge)
	})

	r.POST("/challenge", func(c *gin.Context) {
		challenge := &domain.Challenge{}
		err := c.BindJSON(challenge)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = service.Create(c, challenge)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, challenge)
	})

	r.GET("/challenge/list", func(c *gin.Context) {
		challenges, err := service.List(c, &domain.ListOpts{})
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, challenges)
	})

	r.PUT("/challenge", func(c *gin.Context) {
		challenge := &domain.Challenge{}
		err := c.BindJSON(challenge)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = service.Update(c, challenge)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, challenge)
	})

	r.DELETE("/challenge/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := service.Delete(c, id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "deleted",
		})
	})

	r.POST("/challenge/scoring/:id", func(c *gin.Context) {

		id := c.Param("id")
		// Extract the file from the request
		file, err := c.FormFile("scriptFile")
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to retrieve file: " + err.Error(),
			})
			return
		}

		// Open the uploaded file
		fileContent, err := file.Open()
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to open file: " + err.Error(),
			})
			return
		}
		defer fileContent.Close()

		// Read the content of the file into a string using io.ReadAll
		fileBytes, err := io.ReadAll(fileContent)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to read file: " + err.Error(),
			})
			return
		}
		scriptContent := string(fileBytes)

		submission := &domain.Submission{}
		err = c.BindJSON(submission)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Call the scoring function with the submission and the script content
		result, err := service.Scoring(c, submission, scriptContent, id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Return the result as a JSON response
		c.JSON(200, result)
	})
}
