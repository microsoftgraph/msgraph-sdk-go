package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance
type UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderInternal instantiates a new UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) {
    m := &UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance/{userExperienceAnalyticsAppHealthApplicationPerformance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder instantiates a new UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthApplicationPerformance for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Application Performance
// returns a UserExperienceAnalyticsAppHealthApplicationPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthApplicationPerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthApplicationPerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthApplicationPerformance in deviceManagement
// returns a UserExperienceAnalyticsAppHealthApplicationPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthApplicationPerformanceable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthApplicationPerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthApplicationPerformanceable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthApplicationPerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics appHealth Application Performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthApplicationPerformance in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthApplicationPerformanceable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
