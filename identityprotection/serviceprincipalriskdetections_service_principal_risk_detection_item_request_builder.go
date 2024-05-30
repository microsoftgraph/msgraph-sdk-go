package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder provides operations to manage the servicePrincipalRiskDetections property of the microsoft.graph.identityProtectionRoot entity.
type ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderGetQueryParameters read the properties and relationships of a servicePrincipalRiskDetection object.
type ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderGetQueryParameters
}
// ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderInternal instantiates a new ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder and sets the default values.
func NewServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) {
    m := &ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/servicePrincipalRiskDetections/{servicePrincipalRiskDetection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder instantiates a new ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder and sets the default values.
func NewServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property servicePrincipalRiskDetections for identityProtection
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a servicePrincipalRiskDetection object.
// returns a ServicePrincipalRiskDetectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipalriskdetection-get?view=graph-rest-1.0
func (m *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalRiskDetectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalRiskDetectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalRiskDetectionable), nil
}
// Patch update the navigation property servicePrincipalRiskDetections in identityProtection
// returns a ServicePrincipalRiskDetectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalRiskDetectionable, requestConfiguration *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalRiskDetectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalRiskDetectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalRiskDetectionable), nil
}
// ToDeleteRequestInformation delete navigation property servicePrincipalRiskDetections for identityProtection
// returns a *RequestInformation when successful
func (m *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a servicePrincipalRiskDetection object.
// returns a *RequestInformation when successful
func (m *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property servicePrincipalRiskDetections in identityProtection
// returns a *RequestInformation when successful
func (m *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalRiskDetectionable, requestConfiguration *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder when successful
func (m *ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) WithUrl(rawUrl string)(*ServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder) {
    return NewServiceprincipalriskdetectionsServicePrincipalRiskDetectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
