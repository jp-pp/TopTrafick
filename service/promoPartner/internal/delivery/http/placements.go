package http

import (
	"TopTrafficTest/service/promoPartner/internal/errors"
	"TopTrafficTest/tools/entity/bid"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type placements struct {
	errorLogger 			*log.Logger
	placementsRepository 	PlacementsRepository
}

func NewPlacements(errorLogger *log.Logger, placementsRepository PlacementsRepository) *placements {

	return &placements{errorLogger: errorLogger, placementsRepository: placementsRepository}
}

type PlacementsRepository interface {
	Placements(_ context.Context, request *bid.Request) (response *bid.Response, err error)
}

func (p *placements) Placements(w http.ResponseWriter, r *http.Request) {

	var (
		err					error
		placementsRequest 	*bid.Request
		placementResponse	*bid.Response
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

		placementResponse, err = p.placementsRepository.Placements(r.Context(), placementsRequest)

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

	return
}
