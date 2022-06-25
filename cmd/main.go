package main

import (
	"github.com/PedroMiguel7/GO-ANGULAR-POSTGRE/pkg/people"
    "github.com/PedroMiguel7/GO-ANGULAR-POSTGRE/pkg/teams"
    "github.com/PedroMiguel7/GO-ANGULAR-POSTGRE/pkg/projects"
    "github.com/PedroMiguel7/GO-ANGULAR-POSTGRE/pkg/tasks"
	
	"github.com/PedroMiguel7/GO-ANGULAR-POSTGRE/pkg/common/db"
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

    
    pessoas.RegisterRoutes(r, h)
    projetos.RegisterRoutes(r, h)
    equipes.RegisterRoutes(r, h)
    tasks.RegisterRoutes(r, h)
    // register more routes here

    r.Run(port)
}