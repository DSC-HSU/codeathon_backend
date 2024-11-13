package profile

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"encoding/json"
	"github.com/ServiceWeaver/weaver"
	"strconv"
)

type ProfileService interface {
	GetById(ctx context.Context, id string) (*domain.Profile, error)
	List(ctx context.Context, accessLevel int, opts *domain.ListOpts) (*domain.ListResult[*domain.Profile], error)
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

func (p profileService) List(ctx context.Context, accessLevel int, opts *domain.ListOpts) (*domain.ListResult[*domain.Profile], error) {
	data, count, err := supabase.Client.From("profiles").
		Select("*", "exact", false).Eq("access_level", strconv.Itoa(accessLevel)).
		Range(opts.Offset, opts.Offset+opts.Limit-1, "").Execute()
	if err != nil {
		return nil, err
	}
	var profiles []*domain.Profile
	err = json.Unmarshal(data, &profiles)
	if err != nil {
		return nil, err
	}

	return &domain.ListResult[*domain.Profile]{
		TotalPage: count,
		Data:      profiles,
	}, nil
}

func (p profileService) Update(ctx context.Context, profile *domain.Profile) error {
	_, _, err := supabase.Client.From("profiles").Update(profile, "", "").Eq("id", profile.Id.String()).Execute()
	return err
}

func (p profileService) Delete(ctx context.Context, id string) error {
	_, _, err := supabase.Client.From("profiles").Delete("", "").Eq("id", id).Execute()
	return err
}
