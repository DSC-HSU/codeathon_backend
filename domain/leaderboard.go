package domain

import "github.com/ServiceWeaver/weaver"

type Leaderboard struct {
	weaver.AutoMarshal
	CId          string `json:"c_id"`
	UId          string `json:"u_id"`
	Score        int    `json:"score"`
	SubmittedAt  string `json:"submitted_at"`
	LatestUpdate string `json:"latest_update"`
	RankScore    int    `json:"rank_score"`
}
