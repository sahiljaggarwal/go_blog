package middlewares

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func ValidateDTO(dtoType interface{}) gin.HandlerFunc {
  return func(c *gin.Context) {
    dtoTypeInstance := reflect.TypeOf(dtoType).Elem()
    dtoInstance := reflect.New(dtoTypeInstance).Interface()

    if err := c.ShouldBindJSON(dtoInstance); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: " + err.Error()})
      c.Abort()
      return
    }

    // Log bound data for debugging
    log.Printf("DTO after binding: %+v", dtoInstance)

    if validator, ok := dtoInstance.(interface{ Validate() error }); ok {
      if err := validator.Validate(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        c.Abort()
        return
      }
    }

    c.Set("dto", dtoInstance)
    c.Next()
  }
}
