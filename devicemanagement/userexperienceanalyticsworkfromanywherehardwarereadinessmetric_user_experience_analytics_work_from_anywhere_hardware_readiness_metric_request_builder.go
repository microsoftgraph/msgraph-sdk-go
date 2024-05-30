package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder provides operations to manage the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetQueryParameters user experience analytics work from anywhere hardware readiness metrics.
type UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal instantiates a new UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    m := &UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder instantiates a new UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics work from anywhere hardware readiness metrics.
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable), nil
}
// Patch update the navigation property userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric in deviceManagement
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable, requestConfiguration *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics work from anywhere hardware readiness metrics.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable, requestConfiguration *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
