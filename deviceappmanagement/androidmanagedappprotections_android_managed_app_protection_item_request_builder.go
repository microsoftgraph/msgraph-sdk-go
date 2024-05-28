package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder provides operations to manage the androidManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderGetQueryParameters read properties and relationships of the androidManagedAppProtection object.
type AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderGetQueryParameters
}
// AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.androidManagedAppProtection entity.
// returns a *AndroidmanagedappprotectionsItemAppsRequestBuilder when successful
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) Apps()(*AndroidmanagedappprotectionsItemAppsRequestBuilder) {
    return NewAndroidmanagedappprotectionsItemAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.targetedManagedAppProtection entity.
// returns a *AndroidmanagedappprotectionsItemAssignmentsRequestBuilder when successful
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) Assignments()(*AndroidmanagedappprotectionsItemAssignmentsRequestBuilder) {
    return NewAndroidmanagedappprotectionsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderInternal instantiates a new AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder and sets the default values.
func NewAndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) {
    m := &AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder instantiates a new AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder and sets the default values.
func NewAndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a androidManagedAppProtection.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-androidmanagedappprotection-delete?view=graph-rest-1.0
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploymentSummary provides operations to manage the deploymentSummary property of the microsoft.graph.androidManagedAppProtection entity.
// returns a *AndroidmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder when successful
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) DeploymentSummary()(*AndroidmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    return NewAndroidmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the androidManagedAppProtection object.
// returns a AndroidManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-androidmanagedappprotection-get?view=graph-rest-1.0
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AndroidManagedAppProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAndroidManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AndroidManagedAppProtectionable), nil
}
// Patch update the properties of a androidManagedAppProtection object.
// returns a AndroidManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-androidmanagedappprotection-update?view=graph-rest-1.0
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AndroidManagedAppProtectionable, requestConfiguration *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AndroidManagedAppProtectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAndroidManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AndroidManagedAppProtectionable), nil
}
// ToDeleteRequestInformation deletes a androidManagedAppProtection.
// returns a *RequestInformation when successful
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the androidManagedAppProtection object.
// returns a *RequestInformation when successful
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a androidManagedAppProtection object.
// returns a *RequestInformation when successful
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AndroidManagedAppProtectionable, requestConfiguration *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder when successful
func (m *AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) WithUrl(rawUrl string)(*AndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder) {
    return NewAndroidmanagedappprotectionsAndroidManagedAppProtectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
