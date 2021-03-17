package main

import (
	"flyalien"
	"log"
	"net/http"
	"time"
)

func onlyForV2() flyalien.HandlerFunc {
	return func(c *flyalien.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := flyalien.New()
	r.Use(flyalien.Logger()) // global midlleware

	r.GET("/", func(c *flyalien.Context) {
		c.HTML(http.StatusOK, "<h1>Hello flyalien</h1>")
	})
	r.GET("/hello", func(c *flyalien.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *flyalien.Context) {
		c.JSON(http.StatusOK, flyalien.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.GET("/hello/:name", func(c *flyalien.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *flyalien.Context) {
		c.JSON(http.StatusOK, flyalien.H{"filepath": c.Param("filepath")})
	})

	v1 := r.Group("/v1")
	v1.GET("/hello", func(c *flyalien.Context) {
		c.HTML(http.StatusOK, "<h1>Hello v1</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *flyalien.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
