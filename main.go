package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte(randomToken(64)))
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

	router.GET("/", indexHandler)
	router.POST("/login", loginHandler)
	router.GET("/auth", authHandler)
	router.GET("/logout", logoutHandler)

	authorized := router.Group("/battle")
	authorized.Use(authorizeRequest())
	{
		authorized.GET("/field", fieldHandler)
	}

	router.Run("127.0.0.1:9090")
}
