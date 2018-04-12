package midleware

import (
	/*"net/http"
	"os"*/
	"time"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"../../../../../../utilities"
	"../../../../../../dao/factory"
	"../../../../../../models"
	"database/sql"
	"strconv"
)

const SECRECT_KEY  = "xnyAqBz3ERocwP1lxi12KUyx0zJ76zwxTicap3QWd59Hwyd19Tq3LIAMInCHkLq.775806DFA54B06C277CCF4D355F42F831D7D9226BB247B6A5163EE4A505AD0E1"

func BuildMidleWare() *jwt.GinJWTMiddleware {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "strangebug.com",
		Key:        []byte(SECRECT_KEY),
		Timeout:    time.Hour*24,
		MaxRefresh: time.Hour,

		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			result,_ := FindUsuario(userId,password)
			if result.IdUsuarios > 0 {
				return strconv.Itoa(result.IdUsuarios), true
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
			result,_ := FindUsuario(userID,"")

			return map[string]interface{}{
				"dependencia" : result,
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

	/*r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}*/

	return authMiddleware
}


func Authentication(con *sql.DB) gin.HandlerFunc {
	query := "SELECT user_id FROM users WHERE token = ?"

	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Token")
		var user_id int
		err := con.QueryRow(query, tokenHeader ).Scan(&user_id)
		if err != nil {
			// UNAUTHORISED ERROR ERROR ERROR
		}
		c.Set("UserID", user_id)
		c.Next()

	}
}

func FindUsuario(id string,pass string) (models.Usuarios,error){
	config, err := utilities.GetConfiguration()
	if err != nil {
		panic(err)
	}

	type rque struct {
		idUsuario string
		password  string
	}

	/*var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	ss := new(rque)
	error := c.Bind(&ss)
fmt.Println("error : ",error)*/

	radicacionDao := factory.FactoryDaoUsuario(config.Engine)
	radicaResult, errror := radicacionDao.FindUsuario(id, pass)
	if errror != nil {
		panic(errror)
		return radicaResult,errror
	}

return radicaResult,nil
	//c.JSON(200, radicaResult)
}