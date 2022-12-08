package http

import (
	"TopTrafficTest/service/placementsController/internal/domain/entity/placements"
	"TopTrafficTest/service/placementsController/internal/errors"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type placementsController struct {
	PlacementsUseCase
}

func NewPlacementsController(placementsUseCase PlacementsUseCase) *placementsController {
	return &placementsController{PlacementsUseCase: placementsUseCase}
}

type PlacementsUseCase interface {
	Placements(ctx context.Context, request *placements.Request) (response	*placements.Response, err error)
}

func (c *placementsController) Placements(w http.ResponseWriter, r *http.Request) {

	var (
		err					error
		placementsRequest 	*placements.Request
		placementResponse	*placements.Response
		body 				[]byte
	)


	defer func() {
		err = r.Body.Close()
		if err != nil {
			http.Error(w, errors.UnknownError.Error(), http.StatusInternalServerError)
		}
	}()

	switch r.Method {
	case http.MethodPost:

		body, err = io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, errors.PlacementsRequestIncorrectParametersError.Error(), http.StatusBadRequest)
		}

		err = json.Unmarshal(body, &placementsRequest)
		if err != nil {
			http.Error(w, errors.PlacementsRequestIncorrectParametersError.Error(), http.StatusBadRequest)
		}

		placementResponse, err = c.PlacementsUseCase.Placements(r.Context(), placementsRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(placementResponse)
		if err != nil {
			http.Error(w, errors.UnknownError.Error(), http.StatusInternalServerError)
		}

	default:
		http.Error(w, errors.MethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}

}