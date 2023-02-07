package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"encoding/json"
	"fmt"

	req "github.com/imroc/req/v3"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/auth"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
)

// CreateFunctionGroup is the resolver for the createFunctionGroup field.
func (r *mutationResolver) CreateFunctionGroup(ctx context.Context, input model.CreateFunctionGroupInput) (*model.FunctionGroup, error) {
	userId := auth.UserId(ctx)
	if userId == nil {
		return nil, fmt.Errorf("no valid authtoken was provided")
	}

	var responseData model.FunctionGroup
	body := models.UserIdWrapper[model.CreateFunctionGroupInput]{
		Data:   input,
		UserId: *userId,
	}
	bodyBytes, err := json.Marshal(body)

	resp, err := req.R().
		SetBody(bodyBytes).
		SetResult(&responseData).
		SetContentType("application/json").
		Post("http://localhost:10001/create")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s", resp.String())
	}
	if resp.IsSuccess() {
		return &responseData, nil
	}
	return nil, fmt.Errorf("got unexpected response, raw dump:\n%s", resp.Dump())
}

// ListEntitledGroups is the resolver for the listEntitledGroups field.
func (r *queryResolver) ListEntitledGroups(ctx context.Context) (string, error) {
	userId := auth.UserId(ctx)
	if userId == nil {
		return "", fmt.Errorf("no valid authtoken was provided")
	}

	resp, err := req.R().
		SetBody(struct {
			id string `json:"userId"`
		}{*userId}).
		Post("http://localhost:10001/list")

	if err != nil {
		return "", err
	}
	if resp.IsError() {
		return "", fmt.Errorf("%s", resp.String())
	}
	if resp.IsSuccess() {
		return "", nil
	}
	return "", fmt.Errorf("got unexpected response, raw dump:\n%s", resp.Dump())
}
