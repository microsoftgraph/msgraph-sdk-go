package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the reminderView method.
type ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters invoke function reminderView
type ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *string, startDateTime *string)(*ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/reminderView(StartDateTime='{StartDateTime}',EndDateTime='{EndDateTime}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["EndDateTime"] = *endDateTime
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["StartDateTime"] = *startDateTime
    }
    return m
}
// NewItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function reminderView
// Deprecated: This method is obsolete. Use GetAsReminderViewWithStartDateTimeWithEndDateTimeGetResponse instead.
// returns a ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-reminderview?view=graph-rest-1.0
func (m *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeResponseable), nil
}
// GetAsReminderViewWithStartDateTimeWithEndDateTimeGetResponse invoke function reminderView
// returns a ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-reminderview?view=graph-rest-1.0
func (m *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) GetAsReminderViewWithStartDateTimeWithEndDateTimeGetResponse(ctx context.Context, requestConfiguration *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeGetResponseable), nil
}
// ToGetRequestInformation invoke function reminderView
// returns a *RequestInformation when successful
func (m *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) WithUrl(rawUrl string)(*ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
