package middleware

import (
	"demo_project/database"
	"demo_project/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	// get cookie off req
	tokenString := c.GetHeader("Authorization")

	//if err != nil {
	//	c.AbortWithStatus(http.StatusUnauthorized)
	//}
	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check nbf
		if float64(time.Now().Unix()) > claims["nbf"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// find users with token sub
		var user model.User
		database.GetInstance().GetDB().First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//attach to req
		c.Set("user", user)

		c.Next()

		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
