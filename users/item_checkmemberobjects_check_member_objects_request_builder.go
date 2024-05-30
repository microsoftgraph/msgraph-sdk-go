package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder provides operations to call the checkMemberObjects method.
type ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal instantiates a new ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder and sets the default values.
func NewItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    m := &ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/checkMemberObjects", pathParameters),
    }
    return m
}
// NewItemCheckmemberobjectsCheckMemberObjectsRequestBuilder instantiates a new ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder and sets the default values.
func NewItemCheckmemberobjectsCheckMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkMemberObjects
// Deprecated: This method is obsolete. Use PostAsCheckMemberObjectsPostResponse instead.
// returns a ItemCheckmemberobjectsCheckMemberObjectsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) Post(ctx context.Context, body ItemCheckmemberobjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration)(ItemCheckmemberobjectsCheckMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCheckmemberobjectsCheckMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCheckmemberobjectsCheckMemberObjectsResponseable), nil
}
// PostAsCheckMemberObjectsPostResponse invoke action checkMemberObjects
// returns a ItemCheckmemberobjectsCheckMemberObjectsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) PostAsCheckMemberObjectsPostResponse(ctx context.Context, body ItemCheckmemberobjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration)(ItemCheckmemberobjectsCheckMemberObjectsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCheckmemberobjectsCheckMemberObjectsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCheckmemberobjectsCheckMemberObjectsPostResponseable), nil
}
// ToPostRequestInformation invoke action checkMemberObjects
// returns a *RequestInformation when successful
func (m *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCheckmemberobjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder when successful
func (m *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) WithUrl(rawUrl string)(*ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    return NewItemCheckmemberobjectsCheckMemberObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
