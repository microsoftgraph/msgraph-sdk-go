package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder provides operations to manage the defaultManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderGetQueryParameters read properties and relationships of the defaultManagedAppProtection object.
type DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderGetQueryParameters
}
// DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.defaultManagedAppProtection entity.
// returns a *DefaultmanagedappprotectionsItemAppsRequestBuilder when successful
func (m *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) Apps()(*DefaultmanagedappprotectionsItemAppsRequestBuilder) {
    return NewDefaultmanagedappprotectionsItemAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderInternal instantiates a new DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder and sets the default values.
func NewDefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) {
    m := &DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder instantiates a new DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder and sets the default values.
func NewDefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a defaultManagedAppProtection.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-defaultmanagedappprotection-delete?view=graph-rest-1.0
func (m *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploymentSummary provides operations to manage the deploymentSummary property of the microsoft.graph.defaultManagedAppProtection entity.
// returns a *DefaultmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder when successful
func (m *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) DeploymentSummary()(*DefaultmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    return NewDefaultmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the defaultManagedAppProtection object.
// returns a DefaultManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-defaultmanagedappprotection-get?view=graph-rest-1.0
func (m *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDefaultManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable), nil
}
// Patch update the properties of a defaultManagedAppProtection object.
// returns a DefaultManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-defaultmanagedappprotection-update?view=graph-rest-1.0
func (m *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable, requestConfiguration *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDefaultManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable), nil
}
// ToDeleteRequestInformation deletes a defaultManagedAppProtection.
// returns a *RequestInformation when successful
func (m *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the defaultManagedAppProtection object.
// returns a *RequestInformation when successful
func (m *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a defaultManagedAppProtection object.
// returns a *RequestInformation when successful
func (m *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable, requestConfiguration *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder when successful
func (m *DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) WithUrl(rawUrl string)(*DefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder) {
    return NewDefaultmanagedappprotectionsDefaultManagedAppProtectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
