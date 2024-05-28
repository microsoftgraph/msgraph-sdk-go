package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetQueryParameters user experience analytics device Startup Process Performance
type UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal instantiates a new UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    m := &UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/{userExperienceAnalyticsDeviceStartupProcessPerformance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder instantiates a new UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsDeviceStartupProcessPerformance for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics device Startup Process Performance
// returns a UserExperienceAnalyticsDeviceStartupProcessPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsDeviceStartupProcessPerformance in deviceManagement
// returns a UserExperienceAnalyticsDeviceStartupProcessPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsDeviceStartupProcessPerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics device Startup Process Performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsDeviceStartupProcessPerformance in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
