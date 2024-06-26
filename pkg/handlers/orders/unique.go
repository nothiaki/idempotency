package orders

import (
	"github/nothiaki/idempotency/pkg/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Unique(ctx *gin.Context) {
  idParam := ctx.Param("id")

  id, err := strconv.Atoi(idParam)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Bad request" })
    return
  }

  if id > len(database.Data) {
    ctx.JSON(http.StatusNoContent, gin.H{ "message": "Field id didn't find" })
    return
  }

  ctx.JSON(http.StatusOK, database.Data[id])
}
