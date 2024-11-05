package profile

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/endpoint"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Api(service ProfileService) {
	r := endpoint.GetEngine()
	r.GET("/profiles", func(c *gin.Context) {
		// Lấy giá trị của limit và offset từ query parameters, với giá trị mặc định là 10 và 0 nếu không có.
		limit := 10
		offset := 0

		// Kiểm tra xem người dùng có cung cấp limit và offset không
		if l := c.Query("limit"); l != "" {
			if parsedLimit, err := strconv.Atoi(l); err == nil {
				limit = parsedLimit
			}
		}
		if o := c.Query("offset"); o != "" {
			if parsedOffset, err := strconv.Atoi(o); err == nil {
				offset = parsedOffset
			}
		}

		// Gọi service với các giá trị limit và offset
		list, err := service.List(c, &domain.ListOpts{Offset: offset, Limit: limit})
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Trả về kết quả
		c.JSON(200, list)
	})

	r.GET("/profile/:id", func(c *gin.Context) {
		id := c.Param("id")
		profile, err := service.GetById(c, id)
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
		c.JSON(200, profile)
	})

	r.PUT("/profile", func(c *gin.Context) {
		profile := &domain.Profile{}
		err := c.BindJSON(profile)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Check if the profile exists
		existingProfile, err := service.GetById(c, profile.Id.String())
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
		err = service.Update(c, profile)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, existingProfile)
	})

	r.DELETE("/profile/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Check if the profile exists
		_, err := service.GetById(c, id)
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
}
