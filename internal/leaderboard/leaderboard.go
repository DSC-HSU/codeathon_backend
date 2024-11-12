package leaderboard

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type Leaderboard struct {
	weaver.AutoMarshal
	EndPage int                  `json:"end_page"`
	Data    []*domain.Submission `json:"data"`
}

type GlobalLeaderboard struct {
	weaver.AutoMarshal
	EndPage int                         `json:"end_page"`
	Data    []*domain.GlobalLeaderboard `json:"data"`
}

type LeaderboardService interface {
	GetGlobal(ctx context.Context, listOpts *domain.ListOpts) (*GlobalLeaderboard, error)
	GetByCId(ctx context.Context, cId string, listOpts *domain.ListOpts) (*Leaderboard, error)
	Recalculate(ctx context.Context, cId string) error
}

type leaderboardService struct {
	weaver.Implements[LeaderboardService]
}

func (l leaderboardService) GetGlobal(ctx context.Context, listOpts *domain.ListOpts) (*GlobalLeaderboard, error) {
	str := supabase.Client.Rpc("get_global_leaderboard", "", map[string]interface{}{"offset_value": listOpts.Offset, "limit_value": listOpts.Limit})
	fmt.Println(str)
	var leaderboard []*GlobalLeaderboard
	err := json.Unmarshal([]byte(str), &leaderboard)
	if err != nil {
		return nil, err
	}
	return leaderboard[0], nil
}

func (l leaderboardService) GetByCId(ctx context.Context, cId string, listOpts *domain.ListOpts) (*Leaderboard, error) {
	cidUUID, err := uuid.Parse(cId)
	if err != nil {
		return nil, err
	}

	str := supabase.Client.Rpc("get_leaderboard", "", map[string]interface{}{
		"cid":          cidUUID,
		"offset_value": listOpts.Offset,
		"limit_value":  listOpts.Limit,
	})
	fmt.Println(str)
	var leaderboard []*Leaderboard
	err = json.Unmarshal([]byte(str), &leaderboard)
	if err != nil {
		return nil, err
	}
	return leaderboard[0], nil
}

func (l leaderboardService) Recalculate(ctx context.Context, cId string) error {
	// convert to uuid
	cIdUUID, err := uuid.Parse(cId)
	if err != nil {
		return err
	}
	_ = supabase.Client.Rpc("recalculate_rank_score", "", map[string]interface{}{"cid": cIdUUID})
	return nil
}
