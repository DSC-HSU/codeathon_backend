package endpoint

import "github.com/gin-gonic/gin"

var mainEngine *gin.Engine

func GetEngine() *gin.Engine {
	if mainEngine == nil {
		mainEngine = gin.Default()
	}
	return mainEngine
}
