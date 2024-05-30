package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version details
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) ByUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId(userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId string)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails%2Did"] = userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId
    }
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderInternal instantiates a new UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) {
    m := &UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder instantiates a new UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsCountRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) Count()(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsCountRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics appHealth Application Performance by App Version details
// returns a UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails for deviceManagement
// returns a UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable), nil
}
// ToGetRequestInformation user experience analytics appHealth Application Performance by App Version details
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
