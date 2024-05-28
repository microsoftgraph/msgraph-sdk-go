package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version details
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal instantiates a new UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    m := &UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/{userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder instantiates a new UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Application Performance by App Version details
// returns a UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails in deviceManagement
// returns a UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics appHealth Application Performance by App Version details
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, requestConfiguration *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder when successful
func (m *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
