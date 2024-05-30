package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderGetQueryParameters read properties and relationships of the iosManagedAppProtection object.
type IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderGetQueryParameters
}
// IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.iosManagedAppProtection entity.
// returns a *IosmanagedappprotectionsItemAppsRequestBuilder when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) Apps()(*IosmanagedappprotectionsItemAppsRequestBuilder) {
    return NewIosmanagedappprotectionsItemAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.targetedManagedAppProtection entity.
// returns a *IosmanagedappprotectionsItemAssignmentsRequestBuilder when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) Assignments()(*IosmanagedappprotectionsItemAssignmentsRequestBuilder) {
    return NewIosmanagedappprotectionsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderInternal instantiates a new IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder and sets the default values.
func NewIosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) {
    m := &IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder instantiates a new IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder and sets the default values.
func NewIosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a iosManagedAppProtection.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-iosmanagedappprotection-delete?view=graph-rest-1.0
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploymentSummary provides operations to manage the deploymentSummary property of the microsoft.graph.iosManagedAppProtection entity.
// returns a *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) DeploymentSummary()(*IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    return NewIosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the iosManagedAppProtection object.
// returns a IosManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-iosmanagedappprotection-get?view=graph-rest-1.0
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIosManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable), nil
}
// Patch update the properties of a iosManagedAppProtection object.
// returns a IosManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-iosmanagedappprotection-update?view=graph-rest-1.0
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIosManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable), nil
}
// ToDeleteRequestInformation deletes a iosManagedAppProtection.
// returns a *RequestInformation when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the iosManagedAppProtection object.
// returns a *RequestInformation when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a iosManagedAppProtection object.
// returns a *RequestInformation when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) WithUrl(rawUrl string)(*IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) {
    return NewIosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
