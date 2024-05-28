package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderGetQueryParameters user experience analytics device Startup Processes
type UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderInternal instantiates a new UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) {
    m := &UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses/{userExperienceAnalyticsDeviceStartupProcess%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder instantiates a new UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsDeviceStartupProcesses for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics device Startup Processes
// returns a UserExperienceAnalyticsDeviceStartupProcessable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessable), nil
}
// Patch update the navigation property userExperienceAnalyticsDeviceStartupProcesses in deviceManagement
// returns a UserExperienceAnalyticsDeviceStartupProcessable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessable, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsDeviceStartupProcesses for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics device Startup Processes
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsDeviceStartupProcesses in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessable, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder when successful
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
