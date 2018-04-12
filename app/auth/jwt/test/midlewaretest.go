package main

import (
	"net/http"
	"os"
	"time"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"../../../../utilities"
)

const SECRECT_KEY  = "xnyAqBz3ERocwP1lxi12KUyx0zJ76zwxTicap3QWd59Hwyd19Tq3LIAMInCHkLq.775806DFA54B06C277CCF4D355F42F831D7D9226BB247B6A5163EE4A505AD0E1"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}



func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "strange bug",
		Key:        []byte(SECRECT_KEY),
		Timeout:    time.Hour*24,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {

			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return "1988267267", true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			claims := jwt.ExtractClaims(c)
			if claims["role"] == "admin" {
				return true
			}
			return false
		},
		PayloadFunc: func(userID string) map[string]interface{} {
				return map[string]interface{}{
				"dependencia" : "algun cafe rancio",
				"role": "admin",
				//--reque / def
				"iat" : utilities.UnixNow(),
				"jti" : utilities.GetUUID(),
			}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:Authorization",
		TokenHeadName: "Bearer",//Bearer

		TimeFunc: time.Now,
	}

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	http.ListenAndServe(":"+port, r)
}