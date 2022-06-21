package main

import (
	"github.com/PedroMiguel7/go-gin-api-medium/pkg/people"
    "github.com/PedroMiguel7/go-gin-api-medium/pkg/teams"
    "github.com/PedroMiguel7/go-gin-api-medium/pkg/projects"
    "github.com/PedroMiguel7/go-gin-api-medium/pkg/tasks"
	
	"github.com/PedroMiguel7/go-gin-api-medium/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

    
    people.RegisterRoutes(r, h)
    projects.RegisterRoutes(r, h)
    teams.RegisterRoutes(r, h)
    // register more routes here

    r.Run(port)
}