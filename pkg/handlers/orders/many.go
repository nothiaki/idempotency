package orders

import (
	"github/nothiaki/idempotency/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Many(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, database.Data)
}
