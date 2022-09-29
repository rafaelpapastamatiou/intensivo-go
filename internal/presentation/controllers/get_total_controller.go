package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelpapastamatiou/imersao-go/internal/application/usecases"
)

type GetTotalController struct {
	getTotalUseCase usecases.GetTotalUseCase
}

func NewGetTotalController(getTotalUseCase usecases.GetTotalUseCase) *GetTotalController {
	return &GetTotalController{getTotalUseCase}
}

func (c *GetTotalController) Handle(w http.ResponseWriter, r *http.Request) {
	output, err := c.getTotalUseCase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(output)
}
