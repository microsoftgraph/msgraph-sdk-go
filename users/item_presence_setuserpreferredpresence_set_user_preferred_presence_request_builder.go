package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder provides operations to call the setUserPreferredPresence method.
type ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilderInternal instantiates a new ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder and sets the default values.
func NewItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder) {
    m := &ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/presence/setUserPreferredPresence", pathParameters),
    }
    return m
}
// NewItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder instantiates a new ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder and sets the default values.
func NewItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the preferred availability and activity status for a user. If the preferred presence of a user is set, the user's presence shows as the preferred status. Preferred presence takes effect only when at least one presence session exists for the user. Otherwise, the user's presence shows as Offline. A presence session is created as a result of a successful setPresence operation, or if the user is signed in on a Microsoft Teams client. For more details, see presence sessions and time-out and expiration.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/presence-setuserpreferredpresence?view=graph-rest-1.0
func (m *ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder) Post(ctx context.Context, body ItemPresenceSetuserpreferredpresenceSetUserPreferredPresencePostRequestBodyable, requestConfiguration *ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation set the preferred availability and activity status for a user. If the preferred presence of a user is set, the user's presence shows as the preferred status. Preferred presence takes effect only when at least one presence session exists for the user. Otherwise, the user's presence shows as Offline. A presence session is created as a result of a successful setPresence operation, or if the user is signed in on a Microsoft Teams client. For more details, see presence sessions and time-out and expiration.
// returns a *RequestInformation when successful
func (m *ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPresenceSetuserpreferredpresenceSetUserPreferredPresencePostRequestBodyable, requestConfiguration *ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder when successful
func (m *ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder) WithUrl(rawUrl string)(*ItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder) {
    return NewItemPresenceSetuserpreferredpresenceSetUserPreferredPresenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
