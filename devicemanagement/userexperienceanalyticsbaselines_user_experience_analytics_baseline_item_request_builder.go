package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetQueryParameters user experience analytics baselines
type UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppHealthMetrics provides operations to manage the appHealthMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) AppHealthMetrics()(*UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BatteryHealthMetrics provides operations to manage the batteryHealthMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserexperienceanalyticsbaselinesItemBatteryhealthmetricsBatteryHealthMetricsRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) BatteryHealthMetrics()(*UserexperienceanalyticsbaselinesItemBatteryhealthmetricsBatteryHealthMetricsRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesItemBatteryhealthmetricsBatteryHealthMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BestPracticesMetrics provides operations to manage the bestPracticesMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserexperienceanalyticsbaselinesItemBestpracticesmetricsBestPracticesMetricsRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) BestPracticesMetrics()(*UserexperienceanalyticsbaselinesItemBestpracticesmetricsBestPracticesMetricsRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesItemBestpracticesmetricsBestPracticesMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal instantiates a new UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    m := &UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder instantiates a new UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsBaselines for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// DeviceBootPerformanceMetrics provides operations to manage the deviceBootPerformanceMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserexperienceanalyticsbaselinesItemDevicebootperformancemetricsDeviceBootPerformanceMetricsRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) DeviceBootPerformanceMetrics()(*UserexperienceanalyticsbaselinesItemDevicebootperformancemetricsDeviceBootPerformanceMetricsRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesItemDevicebootperformancemetricsDeviceBootPerformanceMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics baselines
// returns a UserExperienceAnalyticsBaselineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable), nil
}
// Patch update the navigation property userExperienceAnalyticsBaselines in deviceManagement
// returns a UserExperienceAnalyticsBaselineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable), nil
}
// RebootAnalyticsMetrics provides operations to manage the rebootAnalyticsMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserexperienceanalyticsbaselinesItemRebootanalyticsmetricsRebootAnalyticsMetricsRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) RebootAnalyticsMetrics()(*UserexperienceanalyticsbaselinesItemRebootanalyticsmetricsRebootAnalyticsMetricsRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesItemRebootanalyticsmetricsRebootAnalyticsMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourcePerformanceMetrics provides operations to manage the resourcePerformanceMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserexperienceanalyticsbaselinesItemResourceperformancemetricsResourcePerformanceMetricsRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) ResourcePerformanceMetrics()(*UserexperienceanalyticsbaselinesItemResourceperformancemetricsResourcePerformanceMetricsRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesItemResourceperformancemetricsResourcePerformanceMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsBaselines for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics baselines
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsBaselines in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// WorkFromAnywhereMetrics provides operations to manage the workFromAnywhereMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserexperienceanalyticsbaselinesItemWorkfromanywheremetricsWorkFromAnywhereMetricsRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) WorkFromAnywhereMetrics()(*UserexperienceanalyticsbaselinesItemWorkfromanywheremetricsWorkFromAnywhereMetricsRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesItemWorkfromanywheremetricsWorkFromAnywhereMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
