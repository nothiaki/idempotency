package orders

import (
	"github/nothiaki/idempotency/internal/models"
	"github/nothiaki/idempotency/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
  idempotencyId, err := ctx.Request.Header["Idempotency-Id"]
  if err == false {
    ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Error when get payload" })
    return
  }

  for _, order := range database.Data {
    if order.IdempotencyId == idempotencyId[0] {
      ctx.JSON(http.StatusCreated, order)
      return
    }
  }

  var newOrder models.Order

  if err := ctx.BindJSON(&newOrder); err != nil {
    ctx.JSON(http.StatusBadRequest, newOrder)
  }

  newOrder.Id = len(database.Data)
  newOrder.IdempotencyId = idempotencyId[0]

  database.Data = append(database.Data, newOrder)

  ctx.JSON(http.StatusOK, newOrder)
  return
}
