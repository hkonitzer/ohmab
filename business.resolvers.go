package OHMAB

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"

	"github.com/google/uuid"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/user"
	"hynie.de/ohmab/internal/pkg/log"
	"hynie.de/ohmab/internal/pkg/privacy"
)

// CreateBusiness is the resolver for the createBusiness field.
func (r *mutationResolver) CreateBusiness(ctx context.Context, input ent.CreateBusinessInput) (*ent.Business, error) {
	logger := log.GetLoggerInstance()
	uv := privacy.FromContext(ctx)
	uid, err := uuid.Parse(uv.GetUserID())
	if err != nil {
		logger.Err(err).Msgf("Error parsing UUID from UserViewer: %v", uv.GetUserID())
	}
	u, err := r.client.User.Query().Where(user.IDEQ(uid)).Only(ctx)
	if err != nil {
		logger.Err(err).Msgf("Error querying user with ID: '%v'", uid)
		return nil, err
	}
	input.UserIDs = append(input.UserIDs, u.ID)
	b, err := r.client.Business.Create().SetInput(input).Save(ctx)
	if err != nil {
		logger.Err(err).Msgf("Error creating business with Name1: '%v'", input.Name1)
		return nil, err
	}
	return b, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
