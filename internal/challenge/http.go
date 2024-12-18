package challenge

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/endpoint"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strconv"
)

func Api(service ChallengeService) {
	r := endpoint.GetEngine()

	r.GET("/challenge/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Check if 'id' is a valid UUID
		if _, err := uuid.Parse(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid ID format, expected UUID",
			})
			return
		}

		// Attempt to get the challenge by ID
		challenge, err := service.GetById(c, id)
		if err != nil || challenge == nil {
			// If there's an error or the result is nil, assume not found
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Challenge not found",
			})
			return
		}

		// If no error, return the challenge
		c.JSON(http.StatusOK, challenge)
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
		challenge.Id = uuid.New()
		err = service.Create(c, challenge)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, challenge)
	})

	r.GET("/challenges", func(c *gin.Context) {
		// Get limit and offset from query parameters
		var limit, offset int
		var err error

		if l := c.Query("limit"); l != "" {
			if limit, err = strconv.Atoi(l); err != nil {
				c.JSON(400, gin.H{
					"message": "Invalid limit value",
				})
				return
			}
		} else {
			c.JSON(400, gin.H{
				"message": "Limit is required",
			})
			return
		}

		if o := c.Query("offset"); o != "" {
			if offset, err = strconv.Atoi(o); err != nil {
				c.JSON(400, gin.H{
					"message": "Invalid offset value",
				})
				return
			}
		} else {
			c.JSON(400, gin.H{
				"message": "Offset is required",
			})
			return
		}

		challenges, err := service.List(c, &domain.ListOpts{Offset: offset, Limit: limit})
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

		// Check if the challenge exists
		existingChallenge, err := service.GetById(c, challenge.Id.String())
		if err != nil || existingChallenge == nil {
			c.JSON(404, gin.H{
				"message": "Challenge not found",
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

		// Check if the challenge exists
		_, err := service.GetById(c, id)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Challenge not found",
			})
			return
		}

		err = service.Delete(c, id)
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

		// Call the scoring function with the submission and the script content
		result, err := service.Scoring(c, id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Return the result as a JSON response
		c.JSON(200, result)
	})

	// run script
	r.POST("/challenge/run-script", func(c *gin.Context) {
		// receive the script file
		file, err := c.FormFile("script")
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

		// Read the content of the file into a byte slice
		fileBytes, err := io.ReadAll(fileContent)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to read file: " + err.Error(),
			})
			return
		}

		// Call the scoring function with the submission and the script content
		result, err := service.RunScript(c, fileBytes)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Return the result as a JSON response
		c.JSON(200, result)
	})

	// upload eval script
	r.POST("/challenge/:id/eval-script", func(c *gin.Context) {
		id := c.Param("id")

		// Check if 'id' is a valid UUID
		if _, err := uuid.Parse(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid ID format, expected UUID",
			})
			return
		}

		// Extract the file from the request
		file, err := c.FormFile("eval_script")
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

		// Read the content of the file into a byte slice
		fileBytes, err := io.ReadAll(fileContent)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to read file: " + err.Error(),
			})
			return
		}

		// Upload the file to the challenge
		err = service.UploadEvalScript(c, id, fileBytes)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to upload file: " + err.Error(),
			})
			return
		}

		// Return the URL of the uploaded file
		c.JSON(200, gin.H{
			"message": "File uploaded successfully",
		})

	})

	// upload with many input files
	r.POST("/challenge/:id/input-files", func(c *gin.Context) {
		id := c.Param("id")

		// Check if 'id' is a valid UUID
		if _, err := uuid.Parse(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid ID format, expected UUID",
			})
			return
		}

		// Extract the files from the request
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to retrieve files: " + err.Error(),
			})
			return
		}

		// Get the files from the form
		files := form.File["files"]

		// Check if there are any files
		if len(files) == 0 {
			c.JSON(400, gin.H{
				"message": "No files found",
			})
			return
		}

		// Iterate over the files
		var fileContents [][]byte
		for _, file := range files {
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

			// Append the file content to the slice
			fileContents = append(fileContents, fileBytes)
		}

		// Upload the files to the challenge
		err = service.UploadInputFiles(c, id, fileContents)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to upload files: " + err.Error(),
			})
			return
		}

		// Return the URLs of the uploaded files
		c.JSON(200, gin.H{
			"message": "Files uploaded successfully",
		})
	})

}
