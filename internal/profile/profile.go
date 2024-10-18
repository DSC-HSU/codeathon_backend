package profile

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"encoding/json"
	"github.com/ServiceWeaver/weaver"
)

type ProfileService interface {
	GetById(ctx context.Context, id string) (*domain.Profile, error)
	List(ctx context.Context, opts *domain.ListOpts) ([]*domain.Profile, error)
	Update(ctx context.Context, profile *domain.Profile) error
	Delete(ctx context.Context, id string) error
}

type profileService struct {
	weaver.Implements[ProfileService]
}

func (p profileService) GetById(ctx context.Context, id string) (*domain.Profile, error) {
	data, _, err := supabase.Client.From("profiles").Select("*", "", false).Eq("id", id).Single().Execute()
	if err != nil {
		return nil, err
	}
	profile := &domain.Profile{}
	err = json.Unmarshal(data, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (p profileService) List(ctx context.Context, opts *domain.ListOpts) ([]*domain.Profile, error) {
	data, _, err := supabase.Client.From("profiles").Select("*", "", false).Range(opts.Offset, opts.Offset+opts.Limit, "").Execute()
	if err != nil {
		return nil, err
	}
	var profiles []*domain.Profile
	err = json.Unmarshal(data, &profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (p profileService) Update(ctx context.Context, profile *domain.Profile) error {
	_, err := p.GetById(ctx, profile.Id)
	if err != nil {
		return err
	}

	_, _, err = supabase.Client.From("profiles").Update(profile, "", "").Eq("id", profile.Id).Execute()
	return err
}

func (p profileService) Delete(ctx context.Context, id string) error {
	_, err := p.GetById(ctx, id)
	if err != nil {
		return err
	}
	_, _, err = supabase.Client.From("profiles").Delete("", "").Eq("id", id).Execute()
	return err
}
