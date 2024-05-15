package main

import (
	"ina-gin-crud/config"
	"ina-gin-crud/routes"

	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
)

func initOAuth2Manager() *manage.Manager {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientStore.Set("client_id", &models.Client{
		ID:     "client_id",
		Secret: "client_secret",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	manager.MapAccessGenerate(generates.NewAccessGenerate())

	return manager
}

func main() {
	config.ConnectDatabase()
	manager := initOAuth2Manager()
	ginserver.InitServer(manager)

	r := routes.SetupRouter()

	// Setup endpoint untuk mendapatkan token
	r.POST("/token", ginserver.HandleTokenRequest)

	// Integrating OAuth2 middleware with the existing router for tasks
	taskGroup := r.Group("/tasks")
	taskGroup.Use(ginserver.HandleTokenVerify())
	{
		// Here you would setup your task routes within this group,
		// assuming these are defined elsewhere in your routes module.
	}

	r.Run(":8080")
}
