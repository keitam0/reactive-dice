package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keita7i/reactive-dice/dice"
)

type DiceHandler struct {
	Dice dice.Dice
}

func (h DiceHandler) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, DiceResponse{
		Name:  h.Dice.Name(),
		Faces: h.Dice.Ref(),
	})
}

func (h DiceHandler) Post(ctx *gin.Context) {
	h.Dice.Roll()
	ctx.Status(http.StatusNoContent)
}
