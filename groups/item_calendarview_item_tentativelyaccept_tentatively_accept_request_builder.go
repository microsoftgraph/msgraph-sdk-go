package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder provides operations to call the tentativelyAccept method.
type ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderInternal instantiates a new ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder and sets the default values.
func NewItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    m := &ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/tentativelyAccept", pathParameters),
    }
    return m
}
// NewItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder instantiates a new ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder and sets the default values.
func NewItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post tentatively accept the specified event in a user calendar. If the event allows proposals for new times, on responding tentative to the event, an invitee can choose to suggest an alternative time by including the proposedNewTime parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-tentativelyaccept?view=graph-rest-1.0
func (m *ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) Post(ctx context.Context, body ItemCalendarviewItemTentativelyacceptTentativelyAcceptPostRequestBodyable, requestConfiguration *ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation tentatively accept the specified event in a user calendar. If the event allows proposals for new times, on responding tentative to the event, an invitee can choose to suggest an alternative time by including the proposedNewTime parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
// returns a *RequestInformation when successful
func (m *ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarviewItemTentativelyacceptTentativelyAcceptPostRequestBodyable, requestConfiguration *ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
