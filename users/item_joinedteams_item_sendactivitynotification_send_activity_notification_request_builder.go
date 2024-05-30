package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder provides operations to call the sendActivityNotification method.
type ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilderInternal instantiates a new ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder and sets the default values.
func NewItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder) {
    m := &ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/sendActivityNotification", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder instantiates a new ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder and sets the default values.
func NewItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an activity feed notification in the scope of a team. For more information about sending notifications and the requirements for doing so, seesending Teams activity notifications.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-sendactivitynotification?view=graph-rest-1.0
func (m *ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder) Post(ctx context.Context, body ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send an activity feed notification in the scope of a team. For more information about sending notifications and the requirements for doing so, seesending Teams activity notifications.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder when successful
func (m *ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder) {
    return NewItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
