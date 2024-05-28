package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version Device Id
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal instantiates a new UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) {
    m := &UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId/{userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder instantiates a new UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Application Performance by App Version Device Id
// returns a UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId in deviceManagement
// returns a UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics appHealth Application Performance by App Version Device Id
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
