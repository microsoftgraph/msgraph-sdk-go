package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version Device Id
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) ByUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId(userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId string)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId%2Did"] = userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId
    }
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal instantiates a new UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    m := &UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder instantiates a new UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidCountRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) Count()(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidCountRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics appHealth Application Performance by App Version Device Id
// returns a UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId for deviceManagement
// returns a UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation user experience analytics appHealth Application Performance by App Version Device Id
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
