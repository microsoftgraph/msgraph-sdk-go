package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PresencesItemClearpresenceClearPresenceRequestBuilder provides operations to call the clearPresence method.
type PresencesItemClearpresenceClearPresenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PresencesItemClearpresenceClearPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresencesItemClearpresenceClearPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresencesItemClearpresenceClearPresenceRequestBuilderInternal instantiates a new PresencesItemClearpresenceClearPresenceRequestBuilder and sets the default values.
func NewPresencesItemClearpresenceClearPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemClearpresenceClearPresenceRequestBuilder) {
    m := &PresencesItemClearpresenceClearPresenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/presences/{presence%2Did}/clearPresence", pathParameters),
    }
    return m
}
// NewPresencesItemClearpresenceClearPresenceRequestBuilder instantiates a new PresencesItemClearpresenceClearPresenceRequestBuilder and sets the default values.
func NewPresencesItemClearpresenceClearPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemClearpresenceClearPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresencesItemClearpresenceClearPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the application's presence session for a user. If it is the user's only presence session, the user's presence will change to Offline/Offline. For details about presences sessions, see presence: setPresence.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/presence-clearpresence?view=graph-rest-1.0
func (m *PresencesItemClearpresenceClearPresenceRequestBuilder) Post(ctx context.Context, body PresencesItemClearpresenceClearPresencePostRequestBodyable, requestConfiguration *PresencesItemClearpresenceClearPresenceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation clear the application's presence session for a user. If it is the user's only presence session, the user's presence will change to Offline/Offline. For details about presences sessions, see presence: setPresence.
// returns a *RequestInformation when successful
func (m *PresencesItemClearpresenceClearPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body PresencesItemClearpresenceClearPresencePostRequestBodyable, requestConfiguration *PresencesItemClearpresenceClearPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PresencesItemClearpresenceClearPresenceRequestBuilder when successful
func (m *PresencesItemClearpresenceClearPresenceRequestBuilder) WithUrl(rawUrl string)(*PresencesItemClearpresenceClearPresenceRequestBuilder) {
    return NewPresencesItemClearpresenceClearPresenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
