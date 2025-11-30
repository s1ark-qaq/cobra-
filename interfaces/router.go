package interfaces

import "github.com/gin-gonic/gin"

func Router(port string) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Run(port)
}
