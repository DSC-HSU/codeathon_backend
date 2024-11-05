package domain

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
	"time"
)

type Leaderboard struct {
	weaver.AutoMarshal
	CId          uuid.UUID `json:"c_id"`
	UId          uuid.UUID `json:"u_id"`
	Score        int       `json:"score"`
	SubmittedAt  time.Time `json:"submitted_at"`
	LatestUpdate string    `json:"latest_update"`
	RankScore    int       `json:"rank_score"`
}

type GlobalLeaderboard struct {
	weaver.AutoMarshal
	UId   uuid.UUID `json:"user_id"`
	Score float64   `json:"global_rank_score"`
}
