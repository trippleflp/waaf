package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/imroc/req/v3"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/auth"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
)

var functionGroupServiceUrl = func() string {
	url, exist := os.LookupEnv("FUNCTIONGROUP_URL")
	if !exist {
		return "http://localhost:10001"
	}
	return url
}()

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
		Post(fmt.Sprintf("%s/create", functionGroupServiceUrl))

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

// AddUserToFunctionGroup is the resolver for the addUserToFunctionGroup field.
func (r *mutationResolver) AddUserToFunctionGroup(ctx context.Context, users []*model.UserRolePairInput, functionGroupID string) (*model.FunctionGroup, error) {
	userId := auth.UserId(ctx)
	if userId == nil {
		return nil, fmt.Errorf("no valid authtoken was provided")
	}

	var responseData model.FunctionGroup
	body := models.UserIdWrapper[[]*model.UserRolePairInput]{
		Data:   users,
		UserId: *userId,
	}
	bodyBytes, err := json.Marshal(body)

	resp, err := req.R().
		SetBody(bodyBytes).
		SetResult(&responseData).
		SetContentType("application/json").
		Post(fmt.Sprintf("%s/groups/%s/addUsers", functionGroupServiceUrl, functionGroupID))

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

// RemoveUserFromFunctionGroup is the resolver for the removeUserFromFunctionGroup field.
func (r *mutationResolver) RemoveUserFromFunctionGroup(ctx context.Context, userIds []string, functionGroupID string) (*model.FunctionGroup, error) {
	callUserId := auth.UserId(ctx)
	if callUserId == nil {
		return nil, fmt.Errorf("no valid authtoken was provided")
	}

	var responseData model.FunctionGroup
	body := models.UserIdWrapper[[]string]{
		Data:   userIds,
		UserId: *callUserId,
	}
	bodyBytes, err := json.Marshal(body)

	resp, err := req.R().
		SetBody(bodyBytes).
		SetResult(&responseData).
		SetContentType("application/json").
		Post(fmt.Sprintf("%s/groups/%s/removeUsers", functionGroupServiceUrl, functionGroupID))

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

// EditUserRole is the resolver for the editUserRole field.
func (r *mutationResolver) EditUserRole(ctx context.Context, data model.UserRolePairInput, functionGroupID string) (*model.FunctionGroup, error) {
	callUserId := auth.UserId(ctx)
	if callUserId == nil {
		return nil, fmt.Errorf("no valid authtoken was provided")
	}

	var responseData model.FunctionGroup
	body := models.UserIdWrapper[model.UserRolePairInput]{
		Data:   data,
		UserId: *callUserId,
	}
	bodyBytes, err := json.Marshal(body)

	resp, err := req.R().
		SetBody(bodyBytes).
		SetResult(&responseData).
		SetContentType("application/json").
		Post(fmt.Sprintf("%s/groups/%s/editUserRole", functionGroupServiceUrl, functionGroupID))

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

// AddFunctionGroups is the resolver for the addFunctionGroups field.
func (r *mutationResolver) AddFunctionGroups(ctx context.Context, functionGroupIds []string, targetFunctionGroupID string) (*model.FunctionGroup, error) {
	userId := auth.UserId(ctx)
	if userId == nil {
		return nil, fmt.Errorf("no valid authtoken was provided")
	}

	var responseData model.FunctionGroup
	body := models.UserIdWrapper[[]string]{
		Data:   functionGroupIds,
		UserId: *userId,
	}
	bodyBytes, err := json.Marshal(body)

	resp, err := req.R().
		SetBody(bodyBytes).
		SetResult(&responseData).
		SetContentType("application/json").
		Post(fmt.Sprintf("%s/groups/%s/addFunctionGroups", functionGroupServiceUrl, targetFunctionGroupID))

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

// TriggerDeployment is the resolver for the triggerDeployment field.
func (r *mutationResolver) TriggerDeployment(ctx context.Context, functionGroupID string) (*string, error) {
	userId := auth.UserId(ctx)
	if userId == nil {
		return nil, fmt.Errorf("no valid authtoken was provided")
	}

	body := models.UserIdWrapper[any]{
		UserId: *userId,
		Data:   nil,
	}
	resp, err := req.R().
		SetBody(body).
		Post(fmt.Sprintf("%s/groups/%s/deploy", functionGroupServiceUrl, functionGroupID))

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s", resp.String())
	}
	if resp.IsSuccess() {
		successMSg := "Deployment is triggered"
		return &successMSg, nil
	}
	return nil, fmt.Errorf("got unexpected response, raw dump:\n%s", resp.Dump())
}

// ListEntitledGroups is the resolver for the listEntitledGroups field.
func (r *queryResolver) ListEntitledGroups(ctx context.Context) ([]*model.FunctionGroup, error) {
	userId := auth.UserId(ctx)
	if userId == nil {
		return nil, fmt.Errorf("no valid authtoken was provided")
	}

	var result []*model.FunctionGroup
	body := models.UserIdWrapper[any]{
		UserId: *userId,
		Data:   nil,
	}
	resp, err := req.R().
		SetBody(body).
		SetResult(&result).
		Post(fmt.Sprintf("%s/list", functionGroupServiceUrl))

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s", resp.String())
	}
	if resp.IsSuccess() {
		return result, nil
	}
	return nil, fmt.Errorf("got unexpected response, raw dump:\n%s", resp.Dump())
}

// GetFunctionGroup is the resolver for the getFunctionGroup field.
func (r *queryResolver) GetFunctionGroup(ctx context.Context, functionGroupID string) (*model.FunctionGroup, error) {
	userId := auth.UserId(ctx)
	if userId == nil {
		return nil, fmt.Errorf("no valid authtoken was provided")
	}

	var responseData model.FunctionGroup
	body := models.UserIdWrapper[any]{
		Data:   nil,
		UserId: *userId,
	}
	bodyBytes, err := json.Marshal(body)

	resp, err := req.R().
		SetBody(bodyBytes).
		SetResult(&responseData).
		SetContentType("application/json").
		Post(fmt.Sprintf("%s/groups/%s", functionGroupServiceUrl, functionGroupID))

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
