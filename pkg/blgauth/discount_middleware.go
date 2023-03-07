package blgauth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func BrandAuthRequired() gin.HandlerFunc {
	brands := map[int]string{
		1: "brand1",
		2: "brand2",
		3: "brand3",
		4: "brand4",
	}

	return func(c *gin.Context) {
		brandId, err := strconv.Atoi(c.GetHeader("Blg-Brand-Id"))
		if err != nil {
			respondWithError(c, http.StatusPreconditionRequired, "Brand ID and Name required")
			return
		}
		brandName := c.GetHeader("Blg-Brand-Name")

		val, ok := brands[brandId]
		if !ok {
			respondWithError(c, http.StatusPreconditionRequired, "Brand ID and Name required")
			return
		} else if val != brandName {
			respondWithError(c, http.StatusPreconditionRequired, "Brand ID and Name required")
			return
		}

		if brandId == 0 || brandName == "" {
			respondWithError(c, http.StatusPreconditionRequired, "Brand ID and Name required")
			return
		}

		c.Next()
	}
}
