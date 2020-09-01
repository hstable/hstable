package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/stevenroose/gonfig"
	"log"
	"net/http"
	"os"
	"service/crawler"
	"service/handler"
	"service/model"
	"time"
)

//func main() {
//	//r := gin.Default()
//	//r.POST("/course", handler.Get_course)
//	//r.Run()
//	err := crawler.Log_in("20S051030", "liguojian")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("---------------------------------------------------------")
//	crawler.CrawlerCourse()
//}

var identityKey = handler.IdentityKey

//func helloHandler(c *gin.Context) {
//	claims := jwt.ExtractClaims(c)
//	user, _ := c.Get(identityKey)
//	c.JSON(200, gin.H{
//		"userID":   claims[identityKey],
//		"userName": user.(*User).UserName,
//		"text":     "Hello World.",
//	})
//}

var config struct {
	Key       string `desc:"key of the thing"`
}

func init() {
	gonfig.Load(&config, gonfig.Conf{
		//ConfigFileVariable: "config", // enables passing --configfile myfile.conf
		FileDisable: true,
		FileDefaultFilename: "",
		// The default decoder will try TOML, YAML and JSON.
		//FileDecoder: gonfig.DecoderTOML,
		EnvDisable: false,
		EnvPrefix: "HSTABLE_",
	})
	//fmt.Println("config: " + config.Key)
}

func main() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(config.Key),
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
			// Login
			err := crawler.Log_in(userAccount, password)
			if (err == nil) { // login success
				// crawler and store data
				course_data := crawler.CrawlerCourse()
				crawler.StoreData(course_data)
				return &model.GotCourse{
					Account: userAccount,
					Course:  "xxx",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.GotCourse); ok{
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

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.PUT("/course", handler.GetCourseByJW)
		auth.GET("/course", handler.GetCourseByDB)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
