package main

import (
    "github.com/gin-gonic/gin"
    "github.com/PedroMIguel7/API-DESAFIO-BD/pkg/common/models"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    db.Init(dbUrl)

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "port": port,
            "dbUrl": dbUrl,
        })
    })

    r.Run(port)
}