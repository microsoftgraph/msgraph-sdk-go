package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder provides operations to call the clearUserPreferredPresence method.
type PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderInternal instantiates a new PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder and sets the default values.
func NewPresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) {
    m := &PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/presences/{presence%2Did}/clearUserPreferredPresence", pathParameters),
    }
    return m
}
// NewPresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder instantiates a new PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder and sets the default values.
func NewPresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the preferred availability and activity status for a user.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/presence-clearuserpreferredpresence?view=graph-rest-1.0
func (m *PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) Post(ctx context.Context, requestConfiguration *PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderPostRequestConfiguration)(error) {
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
func (m *PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder when successful
func (m *PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) WithUrl(rawUrl string)(*PresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder) {
    return NewPresencesItemClearuserpreferredpresenceClearUserPreferredPresenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
