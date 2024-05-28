package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetpresencesbyuseridGetPresencesByUserIdRequestBuilder provides operations to call the getPresencesByUserId method.
type GetpresencesbyuseridGetPresencesByUserIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetpresencesbyuseridGetPresencesByUserIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetpresencesbyuseridGetPresencesByUserIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetpresencesbyuseridGetPresencesByUserIdRequestBuilderInternal instantiates a new GetpresencesbyuseridGetPresencesByUserIdRequestBuilder and sets the default values.
func NewGetpresencesbyuseridGetPresencesByUserIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetpresencesbyuseridGetPresencesByUserIdRequestBuilder) {
    m := &GetpresencesbyuseridGetPresencesByUserIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/getPresencesByUserId", pathParameters),
    }
    return m
}
// NewGetpresencesbyuseridGetPresencesByUserIdRequestBuilder instantiates a new GetpresencesbyuseridGetPresencesByUserIdRequestBuilder and sets the default values.
func NewGetpresencesbyuseridGetPresencesByUserIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetpresencesbyuseridGetPresencesByUserIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetpresencesbyuseridGetPresencesByUserIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the presence information for multiple users.
// Deprecated: This method is obsolete. Use PostAsGetPresencesByUserIdPostResponse instead.
// returns a GetpresencesbyuseridGetPresencesByUserIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudcommunications-getpresencesbyuserid?view=graph-rest-1.0
func (m *GetpresencesbyuseridGetPresencesByUserIdRequestBuilder) Post(ctx context.Context, body GetpresencesbyuseridGetPresencesByUserIdPostRequestBodyable, requestConfiguration *GetpresencesbyuseridGetPresencesByUserIdRequestBuilderPostRequestConfiguration)(GetpresencesbyuseridGetPresencesByUserIdResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetpresencesbyuseridGetPresencesByUserIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetpresencesbyuseridGetPresencesByUserIdResponseable), nil
}
// PostAsGetPresencesByUserIdPostResponse get the presence information for multiple users.
// returns a GetpresencesbyuseridGetPresencesByUserIdPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudcommunications-getpresencesbyuserid?view=graph-rest-1.0
func (m *GetpresencesbyuseridGetPresencesByUserIdRequestBuilder) PostAsGetPresencesByUserIdPostResponse(ctx context.Context, body GetpresencesbyuseridGetPresencesByUserIdPostRequestBodyable, requestConfiguration *GetpresencesbyuseridGetPresencesByUserIdRequestBuilderPostRequestConfiguration)(GetpresencesbyuseridGetPresencesByUserIdPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetpresencesbyuseridGetPresencesByUserIdPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetpresencesbyuseridGetPresencesByUserIdPostResponseable), nil
}
// ToPostRequestInformation get the presence information for multiple users.
// returns a *RequestInformation when successful
func (m *GetpresencesbyuseridGetPresencesByUserIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body GetpresencesbyuseridGetPresencesByUserIdPostRequestBodyable, requestConfiguration *GetpresencesbyuseridGetPresencesByUserIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetpresencesbyuseridGetPresencesByUserIdRequestBuilder when successful
func (m *GetpresencesbyuseridGetPresencesByUserIdRequestBuilder) WithUrl(rawUrl string)(*GetpresencesbyuseridGetPresencesByUserIdRequestBuilder) {
    return NewGetpresencesbyuseridGetPresencesByUserIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
