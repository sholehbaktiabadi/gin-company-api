package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sholehbaktiabadi/go-api/helper"
	"github.com/sholehbaktiabadi/go-api/service"
)

//authorize jwt validate
func AuthorizeJwt(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed prosess request", "no token provided", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[user_id]:", claims["user_id"])
			log.Println("Claims[issuer]:", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is Not Valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}
