package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderGetQueryParameters user experience analytics device performance
type UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsDevicePerformanceId provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) ByUserExperienceAnalyticsDevicePerformanceId(userExperienceAnalyticsDevicePerformanceId string)(*UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsDevicePerformanceId != "" {
        urlTplParams["userExperienceAnalyticsDevicePerformance%2Did"] = userExperienceAnalyticsDevicePerformanceId
    }
    return NewUserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal instantiates a new UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    m := &UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDevicePerformance{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder instantiates a new UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsdeviceperformanceCountRequestBuilder when successful
func (m *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) Count()(*UserexperienceanalyticsdeviceperformanceCountRequestBuilder) {
    return NewUserexperienceanalyticsdeviceperformanceCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics device performance
// returns a UserExperienceAnalyticsDevicePerformanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDevicePerformanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsDevicePerformanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDevicePerformanceCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsDevicePerformance for deviceManagement
// returns a UserExperienceAnalyticsDevicePerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDevicePerformanceable, requestConfiguration *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDevicePerformanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsDevicePerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDevicePerformanceable), nil
}
// SummarizeDevicePerformanceDevicesWithSummarizeBy provides operations to call the summarizeDevicePerformanceDevices method.
// returns a *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder when successful
func (m *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) SummarizeDevicePerformanceDevicesWithSummarizeBy(summarizeBy *string)(*UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) {
    return NewUserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, summarizeBy)
}
// ToGetRequestInformation user experience analytics device performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsDevicePerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDevicePerformanceable, requestConfiguration *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder when successful
func (m *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    return NewUserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
