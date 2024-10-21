package quest

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"encoding/json"
	"github.com/ServiceWeaver/weaver"
)

type QuestService interface {
	Create(ctx context.Context, quest *domain.Quest) error
	GetById(ctx context.Context, id string) (*domain.Quest, error)
	List(ctx context.Context, opts *domain.ListOpts) ([]*domain.Quest, error)
	Update(ctx context.Context, quest *domain.Quest) error
	Delete(ctx context.Context, id string) error
}

type questService struct {
	weaver.Implements[QuestService]
}

func (q questService) Create(ctx context.Context, quest *domain.Quest) error {
	_, _, err := supabase.Client.From("quests").Insert(quest, false, "", "", "").Execute()
	return err
}

func (q questService) GetById(ctx context.Context, id string) (*domain.Quest, error) {
	data, _, err := supabase.Client.From("quests").Select("*", "", false).Eq("id", id).Single().Execute()
	if err != nil {
		return nil, err
	}
	quest := &domain.Quest{}
	err = json.Unmarshal(data, quest)
	if err != nil {
		return nil, err
	}
	return quest, nil
}

func (q questService) List(ctx context.Context, opts *domain.ListOpts) ([]*domain.Quest, error) {
	data, _, err := supabase.Client.From("quests").Select("*", "", false).Range(opts.Offset, opts.Offset+opts.Limit, "").Execute()
	if err != nil {
		return nil, err
	}
	var quests []*domain.Quest
	err = json.Unmarshal(data, &quests)
	if err != nil {
		return nil, err
	}
	return quests, nil
}

func (q questService) Update(ctx context.Context, quest *domain.Quest) error {
	_, err := q.GetById(ctx, quest.Id)
	if err != nil {
		return err
	}

	_, _, err = supabase.Client.From("quests").Update(quest, "", "").Eq("id", quest.Id).Execute()

	return err
}

func (q questService) Delete(ctx context.Context, id string) error {
	_, err := q.GetById(ctx, id)
	if err != nil {
		return err
	}
	_, _, err = supabase.Client.From("quests").Delete("", "").Eq("id", id).Execute()
	return err
}
