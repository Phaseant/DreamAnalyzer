package handler

import (
	"net/http"

	"github.com/Phaseant/DreamAnalyzer/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) sendPrompt(c *gin.Context) {
	var input model.InputModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	text, err := h.services.NewRequest(input.Text, BindLang(input.Lang))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.OutputModel{Text: text, Error: ""})
}

func BindLang(lang string) int {
	switch lang {
	case "english", "English":
		return model.English
	case "russian", "Russian":
		return model.Russian
	default:
		return -1
	}
}
