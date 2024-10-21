package leaderboard

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/supabase-community/postgrest-go"
	"math"
	"strconv"
)

type Leaderboard struct {
	weaver.AutoMarshal
	TotalPage int                  `json:"total_page"`
	Data      []*domain.Submission `json:"data"`
}

type LeaderboardService interface {
	GetByCId(ctx context.Context, cId string, listOpts *domain.ListOpts) (*Leaderboard, error)
	Recalculate(ctx context.Context, cId string) error
}

type leaderboardService struct {
	weaver.Implements[LeaderboardService]
}

func (l leaderboardService) GetByCId(ctx context.Context, cId string, listOpts *domain.ListOpts) (*Leaderboard, error) {

	data := &[]*domain.Submission{}
	_, err := supabase.Client.From("submissions").Select("*", "", false).Eq("id", cId).Order("score", &postgrest.OrderOpts{Ascending: false}).Range(listOpts.Offset, listOpts.Limit+listOpts.Offset, "").ExecuteTo(data)
	if err != nil {
		return nil, err
	}
	// get total page
	totalData, _, err := supabase.Client.From("submissions").Select("count(*)", "", false).Eq("cid", cId).Execute()
	// bytes to string
	countAll, err := strconv.Atoi(string(totalData))
	if err != nil {
		return nil, err
	}

	totalPage := math.Ceil(float64(countAll) / float64(listOpts.Limit))

	return &Leaderboard{
		TotalPage: int(totalPage),
		Data:      *data,
	}, nil

}

func (l leaderboardService) Recalculate(ctx context.Context, cId string) error {
	//TODO implement me
	panic("implement me")
}
