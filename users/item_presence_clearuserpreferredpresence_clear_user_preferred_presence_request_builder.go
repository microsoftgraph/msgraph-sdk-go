package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder provides operations to call the clearUserPreferredPresence method.
type ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderInternal instantiates a new ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder and sets the default values.
func NewItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) {
    m := &ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/presence/clearUserPreferredPresence", pathParameters),
    }
    return m
}
// NewItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder instantiates a new ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder and sets the default values.
func NewItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the preferred availability and activity status for a user.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/presence-clearuserpreferredpresence?view=graph-rest-1.0
func (m *ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation clear the preferred availability and activity status for a user.
// returns a *RequestInformation when successful
func (m *ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder when successful
func (m *ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) WithUrl(rawUrl string)(*ItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) {
    return NewItemPresenceClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
