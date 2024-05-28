package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderGetQueryParameters the user experience analytics work from anywhere model performance
type UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) ByUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId(userExperienceAnalyticsWorkFromAnywhereModelPerformanceId string)(*UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsWorkFromAnywhereModelPerformanceId != "" {
        urlTplParams["userExperienceAnalyticsWorkFromAnywhereModelPerformance%2Did"] = userExperienceAnalyticsWorkFromAnywhereModelPerformanceId
    }
    return NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderInternal instantiates a new UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) {
    m := &UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder instantiates a new UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsworkfromanywheremodelperformanceCountRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) Count()(*UserexperienceanalyticsworkfromanywheremodelperformanceCountRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremodelperformanceCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the user experience analytics work from anywhere model performance
// returns a UserExperienceAnalyticsWorkFromAnywhereModelPerformanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsWorkFromAnywhereModelPerformance for deviceManagement
// returns a UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable), nil
}
// ToGetRequestInformation the user experience analytics work from anywhere model performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsWorkFromAnywhereModelPerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
