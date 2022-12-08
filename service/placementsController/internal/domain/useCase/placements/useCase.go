package placements

import (
	placementsEntity "TopTrafficTest/service/placementsController/internal/domain/entity/placements"
	"TopTrafficTest/tools/entity/bid"
	"context"
	"math"
	"time"
)

type placements struct {
	RepositoryList []Repository
}

func NewPlacements(repositoryList []Repository) *placements {
	return &placements{RepositoryList: repositoryList}
}

type Repository interface {
	Placements(ctx context.Context, request *bid.Request) (response *bid.Response, err error)
}

func (p *placements) Placements(ctx context.Context, request *placementsEntity.Request) (response *placementsEntity.Response, err error) {

	var (
		bidReq *bid.Request
	)


	bidReq = &bid.Request{
		Id:  request.Id,
		Context: struct {
			Ip        string `json:"ip"`
			UserAgent string `json:"user_agent"`
		}{
			Ip: request.Context.Ip,
			UserAgent: request.Context.UserAgent,
		},
	}

	bidReq.Imp = make([]struct {
		Id        uint `json:"id"`
		Minwidth  uint `json:"minwidth"`
		Minheight uint `json:"minheight"`
	}, len(request.Tiles))

	for i := 0; i < len(request.Tiles); i++ {
		bidReq.Imp[i] = struct {
			Id        uint `json:"id"`
			Minwidth  uint `json:"minwidth"`
			Minheight uint `json:"minheight"`
		}{
			Id: request.Tiles[i].Id,
			Minwidth: request.Tiles[i].Width,
			Minheight: uint(math.Floor(float64(request.Tiles[i].Width) * request.Tiles[i].Ratio)),
		}
	}

	ctxwt, cancel := context.WithTimeout(ctx, time.Millisecond*200)

	for i := 0; i < len(p.RepositoryList); i++ {
		go p.RepositoryList[i].Placements(ctxwt, bidReq)
	}



	return
}