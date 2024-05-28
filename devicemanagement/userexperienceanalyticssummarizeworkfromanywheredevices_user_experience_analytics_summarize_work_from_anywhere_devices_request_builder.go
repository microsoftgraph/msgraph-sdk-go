package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder provides operations to call the userExperienceAnalyticsSummarizeWorkFromAnywhereDevices method.
type UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal instantiates a new UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder and sets the default values.
func NewUserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    m := &UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsSummarizeWorkFromAnywhereDevices()", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder instantiates a new UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder and sets the default values.
func NewUserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function userExperienceAnalyticsSummarizeWorkFromAnywhereDevices
// returns a UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereDevicesSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable), nil
}
// ToGetRequestInformation invoke function userExperienceAnalyticsSummarizeWorkFromAnywhereDevices
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder when successful
func (m *UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    return NewUserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
