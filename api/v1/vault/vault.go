package vault

import (
	"net/http"

	"github.com/gin-gonic/gin"
	p "github.com/matheusrbarbosa/gofin/api/v1/utils/pagination"
	"github.com/matheusrbarbosa/gofin/application/handlers"
)

func handleGetUserVaults(ctx *gin.Context) {
	handler := handlers.VaultHandler()

	response, err := handler.GetAll(p.GetPagination(ctx))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
