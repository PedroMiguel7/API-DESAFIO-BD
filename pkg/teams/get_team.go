package equipes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/PedroMiguel7/GO-ANGULAR-POSTGRE/pkg/common/models"
)

func (h handler) GetTeam(c *gin.Context) {
	id := c.Param("id")

	var equipe models.Equipe

	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &equipe)
}