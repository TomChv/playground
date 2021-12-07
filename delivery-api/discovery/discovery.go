package discovery

import (
	"context"
	"fmt"
)

// UrlParamResponse Response content
type UrlParamResponse struct {
	Message string `json:"message"`
}

// UrlParam A simple route to understand how retrieve a parameter from URL
//encore:api public method=GET path=/url/:param
func UrlParam(ctx context.Context, param string) (*UrlParamResponse, error) {
	return &UrlParamResponse{fmt.Sprintf("Received a query with '%s' as url parameter", param)}, nil
}

// QueryParamData A simple DTO with foo and bar
type QueryParamData struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

// QueryParamResponse Returned result
type QueryParamResponse struct {
	Result string `json:"result"`
}

// QueryParam A simple route to understand how retrieve query parameters
//encore:api public method=GET path=/query
func QueryParam(_ context.Context, params *QueryParamData) (*QueryParamResponse, error) {
	return &QueryParamResponse{
		fmt.Sprintf("Received a query with param Foo='%s' and param Bar='%d'", params.Foo, params.Bar),
	}, nil
}

// BodyDataDTO A simple DTO with a title and a content
type BodyDataDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// BodyDataResponse Response content
type BodyDataResponse struct {
	Message string `json:"message"`
}

// BodyData A simple route to understand how retrieve a body with Encore resolver
//encore:api public method=POST path=/body
func BodyData(_ context.Context, body *BodyDataDTO) (*BodyDataResponse, error) {
	return &BodyDataResponse{
		fmt.Sprintf("Received a body with title '%s' and content '%s'", body.Title, body.Content),
	}, nil
}
