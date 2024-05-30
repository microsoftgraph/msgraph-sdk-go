package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemFindmeetingtimesFindMeetingTimesRequestBuilder provides operations to call the findMeetingTimes method.
type ItemFindmeetingtimesFindMeetingTimesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFindmeetingtimesFindMeetingTimesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFindmeetingtimesFindMeetingTimesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemFindmeetingtimesFindMeetingTimesRequestBuilderInternal instantiates a new ItemFindmeetingtimesFindMeetingTimesRequestBuilder and sets the default values.
func NewItemFindmeetingtimesFindMeetingTimesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFindmeetingtimesFindMeetingTimesRequestBuilder) {
    m := &ItemFindmeetingtimesFindMeetingTimesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/findMeetingTimes", pathParameters),
    }
    return m
}
// NewItemFindmeetingtimesFindMeetingTimesRequestBuilder instantiates a new ItemFindmeetingtimesFindMeetingTimesRequestBuilder and sets the default values.
func NewItemFindmeetingtimesFindMeetingTimesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFindmeetingtimesFindMeetingTimesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFindmeetingtimesFindMeetingTimesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post suggest meeting times and locations based on organizer and attendee availability, and time or location constraints specified as parameters. If findMeetingTimes cannot return any meeting suggestions, the response would indicate a reason in the emptySuggestionsReason property.Based on this value, you can better adjust the parameters and call findMeetingTimes again. The algorithm used to suggest meeting times and locations undergoes fine-tuning from time to time. In scenarios like test environments where the input parameters and calendar data remain static, expect that the suggested results may differ over time.
// returns a MeetingTimeSuggestionsResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-findmeetingtimes?view=graph-rest-1.0
func (m *ItemFindmeetingtimesFindMeetingTimesRequestBuilder) Post(ctx context.Context, body ItemFindmeetingtimesFindMeetingTimesPostRequestBodyable, requestConfiguration *ItemFindmeetingtimesFindMeetingTimesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingTimeSuggestionsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMeetingTimeSuggestionsResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingTimeSuggestionsResultable), nil
}
// ToPostRequestInformation suggest meeting times and locations based on organizer and attendee availability, and time or location constraints specified as parameters. If findMeetingTimes cannot return any meeting suggestions, the response would indicate a reason in the emptySuggestionsReason property.Based on this value, you can better adjust the parameters and call findMeetingTimes again. The algorithm used to suggest meeting times and locations undergoes fine-tuning from time to time. In scenarios like test environments where the input parameters and calendar data remain static, expect that the suggested results may differ over time.
// returns a *RequestInformation when successful
func (m *ItemFindmeetingtimesFindMeetingTimesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemFindmeetingtimesFindMeetingTimesPostRequestBodyable, requestConfiguration *ItemFindmeetingtimesFindMeetingTimesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFindmeetingtimesFindMeetingTimesRequestBuilder when successful
func (m *ItemFindmeetingtimesFindMeetingTimesRequestBuilder) WithUrl(rawUrl string)(*ItemFindmeetingtimesFindMeetingTimesRequestBuilder) {
    return NewItemFindmeetingtimesFindMeetingTimesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
