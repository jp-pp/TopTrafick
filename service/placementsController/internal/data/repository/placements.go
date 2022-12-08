package repository

import (
	"TopTrafficTest/service/placementsController/internal/errors"
	"TopTrafficTest/tools/entity/bid"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type placements struct {
	ConnParameters
}

type ConnParameters interface {
	URL() string
}

func (p *placements) Placements(ctx context.Context, request *bid.Request) (response *bid.Response, err error) {

	var (
		client 		*http.Client
		resp		*http.Response
		jsonBytes	[]byte
	)

	jsonBytes, err = json.Marshal(&request)
	if err != nil {
		err = fmt.Errorf("%v: %v", errors.PlacementsRequestIncorrectParametersError.Error(), err.Error())
		return
	}

	resp, err = client.Post(p.URL(), "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		err = fmt.Errorf("%v: %v", errors.PlacementsRequestError.Error(), err.Error())
		return
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			err = fmt.Errorf("%v: %v", errors.UnknownError.Error(), err.Error())
			return
		}
	}()

	jsonBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("%v: %v", errors.PlacementsResponseIncorrectParametersError.Error(), err.Error())
		return
	}

	err = json.Unmarshal(jsonBytes, &response)
	if err != nil {
		err = fmt.Errorf("%v: %v", errors.PlacementsResponseIncorrectParametersError.Error(), err.Error())
		return
	}

	return
}
