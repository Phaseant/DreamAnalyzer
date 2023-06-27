package handler

import (
	"log"
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

	text, err := h.services.NewRequest(input.Text, input.Lang)
	if err != nil {
		log.Println("Error in sendPrompt: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.OutputModel{Text: text, Error: ""})
}
