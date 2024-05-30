package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetQueryParameters user experience analytics appHealth Model Performance
type UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal instantiates a new UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    m := &UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/{userExperienceAnalyticsAppHealthDeviceModelPerformance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder instantiates a new UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Model Performance
// returns a UserExperienceAnalyticsAppHealthDeviceModelPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance in deviceManagement
// returns a UserExperienceAnalyticsAppHealthDeviceModelPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, requestConfiguration *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics appHealth Model Performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, requestConfiguration *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    return NewUserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
