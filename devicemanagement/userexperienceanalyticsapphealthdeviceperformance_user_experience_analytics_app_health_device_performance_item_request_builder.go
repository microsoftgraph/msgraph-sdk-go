package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderGetQueryParameters user experience analytics appHealth Device Performance
type UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderInternal instantiates a new UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) {
    m := &UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance/{userExperienceAnalyticsAppHealthDevicePerformance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder instantiates a new UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthDevicePerformance for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Device Performance
// returns a UserExperienceAnalyticsAppHealthDevicePerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDevicePerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthDevicePerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDevicePerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthDevicePerformance in deviceManagement
// returns a UserExperienceAnalyticsAppHealthDevicePerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDevicePerformanceable, requestConfiguration *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDevicePerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthDevicePerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDevicePerformanceable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthDevicePerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics appHealth Device Performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthDevicePerformance in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDevicePerformanceable, requestConfiguration *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) {
    return NewUserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
