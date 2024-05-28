package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder provides operations to call the sendActivityNotification method.
type ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilderInternal instantiates a new ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder and sets the default values.
func NewItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder) {
    m := &ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/teamwork/sendActivityNotification", pathParameters),
    }
    return m
}
// NewItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder instantiates a new ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder and sets the default values.
func NewItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an activity feed notification to a user. For more information, see sending Teams activity notifications.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/userteamwork-sendactivitynotification?view=graph-rest-1.0
func (m *ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder) Post(ctx context.Context, body ItemTeamworkSendactivitynotificationSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send an activity feed notification to a user. For more information, see sending Teams activity notifications.
// returns a *RequestInformation when successful
func (m *ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamworkSendactivitynotificationSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder when successful
func (m *ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder) WithUrl(rawUrl string)(*ItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder) {
    return NewItemTeamworkSendactivitynotificationSendActivityNotificationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
