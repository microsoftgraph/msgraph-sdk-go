package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder provides operations to manage the userExperienceAnalyticsOverview property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderGetQueryParameters user experience analytics overview
type UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderInternal instantiates a new UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder and sets the default values.
func NewUserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) {
    m := &UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsOverview{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder instantiates a new UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder and sets the default values.
func NewUserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsOverview for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics overview
// returns a UserExperienceAnalyticsOverviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsOverviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsOverviewable), nil
}
// Patch update the navigation property userExperienceAnalyticsOverview in deviceManagement
// returns a UserExperienceAnalyticsOverviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsOverviewable, requestConfiguration *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsOverviewable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsOverviewable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsOverview for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics overview
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsOverview in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsOverviewable, requestConfiguration *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder when successful
func (m *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) {
    return NewUserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
