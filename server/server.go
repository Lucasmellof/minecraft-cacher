package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasmellof/mojangcache/helpers"
	"github.com/lucasmellof/mojangcache/model"
	"io/ioutil"
	"log"
)

var VERSION = "0.0.1"

func InitServer() {
	log.Printf("Starting MojangCacher developed by Lucasmellof.")
	log.Printf("https://github.com/Lucasmellof")
	log.Printf("Version: %s", VERSION)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "MojangCacher",
			"version": VERSION,
			"author":  "Lucasmellof",
			"url":     "https://github.com/lucasmellof/",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/*
		profiles
		POST /profiles
		Request: ["LucasTSu"]
		Response: [
			{
				"id": "5f0cd06ed0e243a3a79e32feae4b6648",
				"name": "LucasTSu"
			}
		]
	*/
	r.POST("/profiles", func(c *gin.Context) {
		resp, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)
		}
		writeJsonResponse(c, UsernamesToUuids(resp))
	})

	r.GET("/uuid/:username", func(c *gin.Context) {
		username := c.Params.ByName("username")
		writeJsonResponse(c, UsernameToUuid(username))
	})

	r.GET("/username/:uuid", func(c *gin.Context) {
		uuid := c.Params.ByName("uuid")
		writeJsonResponse(c, UuidToUsername(uuid))
	})

	r.GET("/profile/:uuid", func(c *gin.Context) {
		uuid := c.Params.ByName("uuid")
		unsigned, exist := c.GetQuery("unsigned")
		if !exist {
			unsigned = "true"
		}
		writeJsonResponse(c, UuidToProfile(uuid, unsigned))
	})

	addr := ":" + helpers.GetEnv("PORT", "3001")
	log.Printf("Starting server on %s", addr)
	log.Printf("Server started.")
	r.Run(addr)
}

func writeJsonResponse(c *gin.Context, response model.MojangResponse) {
	c.Header("Content-Type", "application/json")
	c.String(response.Code, response.Json)
}
