package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters read properties and relationships of the targetedManagedAppConfiguration object.
type TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters
}
// TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.targetedManagedAppConfiguration entity.
// returns a *TargetedmanagedappconfigurationsItemAppsRequestBuilder when successful
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Apps()(*TargetedmanagedappconfigurationsItemAppsRequestBuilder) {
    return NewTargetedmanagedappconfigurationsItemAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assign provides operations to call the assign method.
// returns a *TargetedmanagedappconfigurationsItemAssignRequestBuilder when successful
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Assign()(*TargetedmanagedappconfigurationsItemAssignRequestBuilder) {
    return NewTargetedmanagedappconfigurationsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.targetedManagedAppConfiguration entity.
// returns a *TargetedmanagedappconfigurationsItemAssignmentsRequestBuilder when successful
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Assignments()(*TargetedmanagedappconfigurationsItemAssignmentsRequestBuilder) {
    return NewTargetedmanagedappconfigurationsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal instantiates a new TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) {
    m := &TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder instantiates a new TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a targetedManagedAppConfiguration.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-targetedmanagedappconfiguration-delete?view=graph-rest-1.0
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploymentSummary provides operations to manage the deploymentSummary property of the microsoft.graph.targetedManagedAppConfiguration entity.
// returns a *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder when successful
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) DeploymentSummary()(*TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    return NewTargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the targetedManagedAppConfiguration object.
// returns a TargetedManagedAppConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-targetedmanagedappconfiguration-get?view=graph-rest-1.0
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTargetedManagedAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable), nil
}
// Patch update the properties of a targetedManagedAppConfiguration object.
// returns a TargetedManagedAppConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-targetedmanagedappconfiguration-update?view=graph-rest-1.0
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable, requestConfiguration *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTargetedManagedAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable), nil
}
// TargetApps provides operations to call the targetApps method.
// returns a *TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder when successful
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) TargetApps()(*TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder) {
    return NewTargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a targetedManagedAppConfiguration.
// returns a *RequestInformation when successful
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the targetedManagedAppConfiguration object.
// returns a *RequestInformation when successful
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a targetedManagedAppConfiguration object.
// returns a *RequestInformation when successful
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable, requestConfiguration *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder when successful
func (m *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*TargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder) {
    return NewTargetedmanagedappconfigurationsTargetedManagedAppConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
