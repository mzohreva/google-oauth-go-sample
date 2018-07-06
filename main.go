package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte(RandToken(64)))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("goquestsession", store))
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", IndexHandler)
	router.POST("/login", LoginHandler)
	router.GET("/auth", AuthHandler)
	router.GET("/logout", LogoutHandler)

	authorized := router.Group("/battle")
	authorized.Use(AuthorizeRequest())
	{
		authorized.GET("/field", FieldHandler)
	}

	router.Run("127.0.0.1:9090")
}
