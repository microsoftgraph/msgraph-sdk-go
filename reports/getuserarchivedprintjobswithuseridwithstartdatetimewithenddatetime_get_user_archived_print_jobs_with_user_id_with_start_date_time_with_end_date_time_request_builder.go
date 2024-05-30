package reports

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the getUserArchivedPrintJobs method.
type GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters get a list of archived print jobs for a particular user.
type GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
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
// GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewGetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewGetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getUserArchivedPrintJobs(userId='{userId}',startDateTime={startDateTime},endDateTime={endDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if userId != nil {
        m.BaseRequestBuilder.PathParameters["userId"] = *userId
    }
    return m
}
// NewGetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewGetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get get a list of archived print jobs for a particular user.
// Deprecated: This method is obsolete. Use GetAsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse instead.
// returns a GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reports-getuserarchivedprintjobs?view=graph-rest-1.0
func (m *GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseable), nil
}
// GetAsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse get a list of archived print jobs for a particular user.
// returns a GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reports-getuserarchivedprintjobs?view=graph-rest-1.0
func (m *GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) GetAsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse(ctx context.Context, requestConfiguration *GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponseable), nil
}
// ToGetRequestInformation get a list of archived print jobs for a particular user.
// returns a *RequestInformation when successful
func (m *GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) WithUrl(rawUrl string)(*GetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewGetuserarchivedprintjobswithuseridwithstartdatetimewithenddatetimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
