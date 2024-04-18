package helper

import "github.com/gin-gonic/gin"

func ErrorStatusCode(code int, c *gin.Context, err error, message string) {
	if err != nil {
		c.JSON(code, gin.H{
			"message": message,
		})
	}
	return
}
