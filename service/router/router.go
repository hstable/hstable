package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"service/config"
	"service/crawler"
	"service/handler"
	"service/model"
	"time"
)

var identityKey = handler.IdentityKey

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	if gin.Mode() == gin.ReleaseMode {
		corsConfig.AllowOrigins = []string{
			"hstable.cn", "www.hstable.cn",
		}
	} else {
		corsConfig.AllowAllOrigins = true
	}
	r.Use(cors.New(corsConfig))
	r.Use(gin.Recovery())
	g := r.Group("/api")
	conf := config.GetConfig()

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(conf.Key),
		Timeout:     time.Hour * 24 * 30,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.GotCourse); ok {
				return jwt.MapClaims{
					identityKey: v.Account,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.GotCourse{
				Account: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals model.LoginData
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userAccount := loginVals.Account
			password := loginVals.Password
			// Login crawler and store data
			_, err := crawler.Log_in(userAccount, password)
			if err == nil { // login success
				return &model.GotCourse{
					Account: userAccount,
					Course:  "xxx",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.GotCourse); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	g.POST("/login", authMiddleware.LoginHandler)

	//r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
	//	claims := jwt.ExtractClaims(c)
	//	log.Printf("NoRoute claims: %#v\n", claims)
	//	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	//})

	auth := g.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.PUT("/course", handler.GetCourseByJW)
		auth.GET("/course", handler.GetCourseByDB)
	}
	return r
}
