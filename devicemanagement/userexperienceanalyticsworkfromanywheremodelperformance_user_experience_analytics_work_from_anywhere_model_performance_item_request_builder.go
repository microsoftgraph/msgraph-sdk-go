package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetQueryParameters the user experience analytics work from anywhere model performance
type UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal instantiates a new UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) {
    m := &UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance/{userExperienceAnalyticsWorkFromAnywhereModelPerformance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder instantiates a new UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsWorkFromAnywhereModelPerformance for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the user experience analytics work from anywhere model performance
// returns a UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property userExperienceAnalyticsWorkFromAnywhereModelPerformance in deviceManagement
// returns a UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsWorkFromAnywhereModelPerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the user experience analytics work from anywhere model performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsWorkFromAnywhereModelPerformance in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, requestConfiguration *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
