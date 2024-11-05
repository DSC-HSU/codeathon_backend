package leaderboard

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/endpoint"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Api(service LeaderboardService) {
	r := endpoint.GetEngine()

	// get global leaderboard
	r.GET("/leaderboard", func(c *gin.Context) {
		limit := 10
		offset := 0

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

		list, err := service.GetGlobal(c, &domain.ListOpts{Offset: offset, Limit: limit})
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, list)
	})

	// get leaderboard by challenge id
	r.GET("/leaderboard/:cid", func(c *gin.Context) {
		// query params
		limit := 10
		offset := 0

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

		cid := c.Param("cid")
		list, err := service.GetByCId(c, cid, &domain.ListOpts{Offset: offset, Limit: limit})
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, list)
	})

	// recalculate leaderboard
	r.POST("/leaderboard/:cid/recalculate", func(c *gin.Context) {
		cid := c.Param("cid")
		err := service.Recalculate(c, cid)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "leaderboard recalculated",
		})
	})

}
