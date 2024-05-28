package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementMessagesMarkreadMarkReadRequestBuilder provides operations to call the markRead method.
type ServiceannouncementMessagesMarkreadMarkReadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementMessagesMarkreadMarkReadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesMarkreadMarkReadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementMessagesMarkreadMarkReadRequestBuilderInternal instantiates a new ServiceannouncementMessagesMarkreadMarkReadRequestBuilder and sets the default values.
func NewServiceannouncementMessagesMarkreadMarkReadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesMarkreadMarkReadRequestBuilder) {
    m := &ServiceannouncementMessagesMarkreadMarkReadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages/markRead", pathParameters),
    }
    return m
}
// NewServiceannouncementMessagesMarkreadMarkReadRequestBuilder instantiates a new ServiceannouncementMessagesMarkreadMarkReadRequestBuilder and sets the default values.
func NewServiceannouncementMessagesMarkreadMarkReadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesMarkreadMarkReadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementMessagesMarkreadMarkReadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark a list of serviceUpdateMessages as read for the signed in user.
// Deprecated: This method is obsolete. Use PostAsMarkReadPostResponse instead.
// returns a ServiceannouncementMessagesMarkreadMarkReadResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-markread?view=graph-rest-1.0
func (m *ServiceannouncementMessagesMarkreadMarkReadRequestBuilder) Post(ctx context.Context, body ServiceannouncementMessagesMarkreadMarkReadPostRequestBodyable, requestConfiguration *ServiceannouncementMessagesMarkreadMarkReadRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesMarkreadMarkReadResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesMarkreadMarkReadResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesMarkreadMarkReadResponseable), nil
}
// PostAsMarkReadPostResponse mark a list of serviceUpdateMessages as read for the signed in user.
// returns a ServiceannouncementMessagesMarkreadMarkReadPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-markread?view=graph-rest-1.0
func (m *ServiceannouncementMessagesMarkreadMarkReadRequestBuilder) PostAsMarkReadPostResponse(ctx context.Context, body ServiceannouncementMessagesMarkreadMarkReadPostRequestBodyable, requestConfiguration *ServiceannouncementMessagesMarkreadMarkReadRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesMarkreadMarkReadPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesMarkreadMarkReadPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesMarkreadMarkReadPostResponseable), nil
}
// ToPostRequestInformation mark a list of serviceUpdateMessages as read for the signed in user.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesMarkreadMarkReadRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceannouncementMessagesMarkreadMarkReadPostRequestBodyable, requestConfiguration *ServiceannouncementMessagesMarkreadMarkReadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementMessagesMarkreadMarkReadRequestBuilder when successful
func (m *ServiceannouncementMessagesMarkreadMarkReadRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementMessagesMarkreadMarkReadRequestBuilder) {
    return NewServiceannouncementMessagesMarkreadMarkReadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
