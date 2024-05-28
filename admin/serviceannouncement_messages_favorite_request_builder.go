package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementMessagesFavoriteRequestBuilder provides operations to call the favorite method.
type ServiceannouncementMessagesFavoriteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementMessagesFavoriteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesFavoriteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementMessagesFavoriteRequestBuilderInternal instantiates a new ServiceannouncementMessagesFavoriteRequestBuilder and sets the default values.
func NewServiceannouncementMessagesFavoriteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesFavoriteRequestBuilder) {
    m := &ServiceannouncementMessagesFavoriteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages/favorite", pathParameters),
    }
    return m
}
// NewServiceannouncementMessagesFavoriteRequestBuilder instantiates a new ServiceannouncementMessagesFavoriteRequestBuilder and sets the default values.
func NewServiceannouncementMessagesFavoriteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesFavoriteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementMessagesFavoriteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post change the status of a list of serviceUpdateMessages to favorite for the signed in user.
// Deprecated: This method is obsolete. Use PostAsFavoritePostResponse instead.
// returns a ServiceannouncementMessagesFavoriteResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-favorite?view=graph-rest-1.0
func (m *ServiceannouncementMessagesFavoriteRequestBuilder) Post(ctx context.Context, body ServiceannouncementMessagesFavoritePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesFavoriteRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesFavoriteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesFavoriteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesFavoriteResponseable), nil
}
// PostAsFavoritePostResponse change the status of a list of serviceUpdateMessages to favorite for the signed in user.
// returns a ServiceannouncementMessagesFavoritePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-favorite?view=graph-rest-1.0
func (m *ServiceannouncementMessagesFavoriteRequestBuilder) PostAsFavoritePostResponse(ctx context.Context, body ServiceannouncementMessagesFavoritePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesFavoriteRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesFavoritePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesFavoritePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesFavoritePostResponseable), nil
}
// ToPostRequestInformation change the status of a list of serviceUpdateMessages to favorite for the signed in user.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesFavoriteRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceannouncementMessagesFavoritePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesFavoriteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementMessagesFavoriteRequestBuilder when successful
func (m *ServiceannouncementMessagesFavoriteRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementMessagesFavoriteRequestBuilder) {
    return NewServiceannouncementMessagesFavoriteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
