package repository

import (
	"TopTrafficTest/service/promoPartner/internal/errors"
	"TopTrafficTest/tools/entity/bid"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type placements struct {
	file string
}

func NewPlacements(file string) *placements {
	return &placements{file: file}
}

func (p *placements) Placements(_ context.Context, request *bid.Request) (response *bid.Response, err error) {

	var (
		jsonBytes		[]byte
		requestImpMap	map[uint]*bid.RequestImp
	)

	jsonBytes, err = os.ReadFile(p.file)
	if err != nil {
		err = fmt.Errorf("%v: %v", errors.UnableToReadPlacementsError.Error(), err.Error())
		return
	}

	err = json.Unmarshal(jsonBytes, &response)
	if err != nil {
		err = fmt.Errorf("%v: %v", errors.UnmarshallingPlacementsError.Error(), err.Error())
		return
	}

	if request == nil || response == nil {
		return
	}

	if request.Id != response.Id {
		return
	}

	requestImpMap = make(map[uint]*bid.RequestImp, len(request.Imp))

	for i := 0; i < len(request.Imp); i++ {
		requestImpMap[request.Imp[i].Id] = &request.Imp[i]
	}

	var selectedImpIndex int

	for i := 0; i < len(response.Imp); i++ {

		if requestImpMap[response.Imp[i].Id] != nil {

			if requestImpMap[response.Imp[i].Id].Minwidth <= response.Imp[i].Width && requestImpMap[response.Imp[i].Id].Minheight <= response.Imp[i].Height {
				selectedResponseImp := response.Imp[i]
				response.Imp[i] = response.Imp[selectedImpIndex]
				response.Imp[selectedImpIndex] = selectedResponseImp
			}
		}

	}

	response.Imp = response.Imp[:selectedImpIndex]

	return
}
