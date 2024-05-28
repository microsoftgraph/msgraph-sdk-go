package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarGetscheduleGetScheduleRequestBuilder provides operations to call the getSchedule method.
type ItemCalendarGetscheduleGetScheduleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGetscheduleGetScheduleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGetscheduleGetScheduleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarGetscheduleGetScheduleRequestBuilderInternal instantiates a new ItemCalendarGetscheduleGetScheduleRequestBuilder and sets the default values.
func NewItemCalendarGetscheduleGetScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGetscheduleGetScheduleRequestBuilder) {
    m := &ItemCalendarGetscheduleGetScheduleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/getSchedule", pathParameters),
    }
    return m
}
// NewItemCalendarGetscheduleGetScheduleRequestBuilder instantiates a new ItemCalendarGetscheduleGetScheduleRequestBuilder and sets the default values.
func NewItemCalendarGetscheduleGetScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGetscheduleGetScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGetscheduleGetScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the free/busy availability information for a collection of users, distributions lists, or resources (rooms or equipment) for a specified time period.
// Deprecated: This method is obsolete. Use PostAsGetSchedulePostResponse instead.
// returns a ItemCalendarGetscheduleGetScheduleResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/calendar-getschedule?view=graph-rest-1.0
func (m *ItemCalendarGetscheduleGetScheduleRequestBuilder) Post(ctx context.Context, body ItemCalendarGetscheduleGetSchedulePostRequestBodyable, requestConfiguration *ItemCalendarGetscheduleGetScheduleRequestBuilderPostRequestConfiguration)(ItemCalendarGetscheduleGetScheduleResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarGetscheduleGetScheduleResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarGetscheduleGetScheduleResponseable), nil
}
// PostAsGetSchedulePostResponse get the free/busy availability information for a collection of users, distributions lists, or resources (rooms or equipment) for a specified time period.
// returns a ItemCalendarGetscheduleGetSchedulePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/calendar-getschedule?view=graph-rest-1.0
func (m *ItemCalendarGetscheduleGetScheduleRequestBuilder) PostAsGetSchedulePostResponse(ctx context.Context, body ItemCalendarGetscheduleGetSchedulePostRequestBodyable, requestConfiguration *ItemCalendarGetscheduleGetScheduleRequestBuilderPostRequestConfiguration)(ItemCalendarGetscheduleGetSchedulePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarGetscheduleGetSchedulePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarGetscheduleGetSchedulePostResponseable), nil
}
// ToPostRequestInformation get the free/busy availability information for a collection of users, distributions lists, or resources (rooms or equipment) for a specified time period.
// returns a *RequestInformation when successful
func (m *ItemCalendarGetscheduleGetScheduleRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarGetscheduleGetSchedulePostRequestBodyable, requestConfiguration *ItemCalendarGetscheduleGetScheduleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarGetscheduleGetScheduleRequestBuilder when successful
func (m *ItemCalendarGetscheduleGetScheduleRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarGetscheduleGetScheduleRequestBuilder) {
    return NewItemCalendarGetscheduleGetScheduleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
