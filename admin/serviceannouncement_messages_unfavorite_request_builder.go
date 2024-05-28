package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementMessagesUnfavoriteRequestBuilder provides operations to call the unfavorite method.
type ServiceannouncementMessagesUnfavoriteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementMessagesUnfavoriteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesUnfavoriteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementMessagesUnfavoriteRequestBuilderInternal instantiates a new ServiceannouncementMessagesUnfavoriteRequestBuilder and sets the default values.
func NewServiceannouncementMessagesUnfavoriteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesUnfavoriteRequestBuilder) {
    m := &ServiceannouncementMessagesUnfavoriteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages/unfavorite", pathParameters),
    }
    return m
}
// NewServiceannouncementMessagesUnfavoriteRequestBuilder instantiates a new ServiceannouncementMessagesUnfavoriteRequestBuilder and sets the default values.
func NewServiceannouncementMessagesUnfavoriteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesUnfavoriteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementMessagesUnfavoriteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove the favorite status of serviceUpdateMessages for the signed in user.
// Deprecated: This method is obsolete. Use PostAsUnfavoritePostResponse instead.
// returns a ServiceannouncementMessagesUnfavoriteResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-unfavorite?view=graph-rest-1.0
func (m *ServiceannouncementMessagesUnfavoriteRequestBuilder) Post(ctx context.Context, body ServiceannouncementMessagesUnfavoritePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesUnfavoriteRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesUnfavoriteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesUnfavoriteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesUnfavoriteResponseable), nil
}
// PostAsUnfavoritePostResponse remove the favorite status of serviceUpdateMessages for the signed in user.
// returns a ServiceannouncementMessagesUnfavoritePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-unfavorite?view=graph-rest-1.0
func (m *ServiceannouncementMessagesUnfavoriteRequestBuilder) PostAsUnfavoritePostResponse(ctx context.Context, body ServiceannouncementMessagesUnfavoritePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesUnfavoriteRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesUnfavoritePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesUnfavoritePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesUnfavoritePostResponseable), nil
}
// ToPostRequestInformation remove the favorite status of serviceUpdateMessages for the signed in user.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesUnfavoriteRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceannouncementMessagesUnfavoritePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesUnfavoriteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementMessagesUnfavoriteRequestBuilder when successful
func (m *ServiceannouncementMessagesUnfavoriteRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementMessagesUnfavoriteRequestBuilder) {
    return NewServiceannouncementMessagesUnfavoriteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
