package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder provides operations to manage the appScope property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
type DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderGetQueryParameters read-only property with details of the app-specific scope when the role eligibility is scoped to an app. Nullable. Supports $expand.
type DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderGetQueryParameters
}
// NewDirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderInternal instantiates a new DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder) {
    m := &DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}/appScope{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder instantiates a new DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only property with details of the app-specific scope when the role eligibility is scoped to an app. Nullable. Supports $expand.
// returns a AppScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppScopeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppScopeable), nil
}
// ToGetRequestInformation read-only property with details of the app-specific scope when the role eligibility is scoped to an app. Nullable. Supports $expand.
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
