package main

import (
	"codeathon.runwayclub.dev/internal/challenge"
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/endpoint"
	"codeathon.runwayclub.dev/internal/leaderboard"
	"codeathon.runwayclub.dev/internal/profile"
	"codeathon.runwayclub.dev/internal/submission"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"log"
)

func main() {
	// read conf
	err := conf.ReadConfig("./env/config.json")
	if err != nil {
		panic(err)
	}
	// print config
	log.Printf("config loaded: %v", conf.Config)
	supabase.Init()
	log.Printf("supabase client initialized")

	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}

}

type app struct {
	weaver.Implements[weaver.Main]
	profileService     weaver.Ref[profile.ProfileService]
	challengeService   weaver.Ref[challenge.ChallengeService]
	submissionService  weaver.Ref[submission.SubmissionService]
	leaderboardService weaver.Ref[leaderboard.LeaderboardService]
	listener           weaver.Listener
}

func serve(ctx context.Context, app *app) error {
	log.Printf("serving on %s", app.listener.Addr())

	r := endpoint.GetEngine()
	// Add CORS
	r.Use(cors.Default())

	// Add profile endpoint
	profile.Api(app.profileService.Get())
	// Add challenge endpoint
	challenge.Api(app.challengeService.Get())
	// Add submission endpoint
	submission.Api(app.submissionService.Get())
	// Add leaderboard endpoint
	leaderboard.Api(app.leaderboardService.Get())

	return r.RunListener(app.listener)
}
