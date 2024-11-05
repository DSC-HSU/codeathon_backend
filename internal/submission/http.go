package submission

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/endpoint"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"strings"
)

func Api(service SubmissionService) {
	r := endpoint.GetEngine()

	r.GET("/submission", func(c *gin.Context) {
		// Get challenge ID and user ID from query parameters
		challengeId := c.Query("challengeID")
		userId := c.Query("userID")

		// Check if challengeID or userID is missing
		if challengeId == "" || userId == "" {
			c.JSON(400, gin.H{
				"message": "challengeID and userID are required",
			})
			return
		}

		submission, err := service.GetByChallengeIdAndUserId(c, userId, challengeId)
		if err != nil {
			if strings.Contains(err.Error(), "PGRST116") {
				// Respond with 404 if the specific error code is found
				c.JSON(404, gin.H{
					"message": "Submission not found",
				})
				return
			}

			// Respond with 400 for other errors
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Respond with 200 if submission is found
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
		// if the submission missing the ID, generate a new one
		if submission.Id.String() == "00000000-0000-0000-0000-000000000000" {
			submission.Id = uuid.New()
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
		// check if the submission exists
		_, err = service.GetById(c, submission.Id.String())
		if err != nil {
			fmt.Println(submission.Id.String())
			fmt.Println(err)
			// Respond with 404 if the specific error code is found
			if strings.Contains(err.Error(), "PGRST116") {
				c.JSON(404, gin.H{
					"message": "Submission not found",
				})
				return
			} else {
				// Respond with 400 for other errors
				c.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}
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

	// Upload output file for a submission
	r.POST("/submission/:id/output-file", func(c *gin.Context) {
		id := c.Param("id")
		// Check if the submission exists
		_, err := service.GetById(c, id)
		if err != nil {
			if strings.Contains(err.Error(), "PGRST116") {
				c.JSON(404, gin.H{
					"message": "Submission not found",
				})
				return
			}
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Extract the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
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

		// Read the content of the file into a byte slice
		fileBytes, err := io.ReadAll(fileContent)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to read file: " + err.Error(),
			})
			return
		}

		fileUrl, err := service.UploadOutputFile(c, id, fileBytes)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"url": fileUrl,
		})

	})

}
