package main

import (
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/endpoint"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// read conf
	err := conf.ReadConfig("./env/config.json")
	if err != nil {
		panic(err)
	}
	// print config
	log.Printf("config loaded: %v", conf.Config)
	supabase.Init()
	log.Printf("supabase client initialized")

	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}

}

type app struct {
	weaver.Implements[weaver.Main]
	listener weaver.Listener
}

func serve(ctx context.Context, app *app) error {
	log.Printf("serving on %s", app.listener.Addr())

	r := endpoint.GetEngine()
	// Add CORS
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r.RunListener(app.listener)
}
