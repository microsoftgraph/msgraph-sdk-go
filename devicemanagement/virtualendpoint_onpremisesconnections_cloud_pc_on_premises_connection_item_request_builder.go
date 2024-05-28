package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters read the properties and relationships of the cloudPcOnPremisesConnection object.
type VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters
}
// VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal instantiates a new VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewVirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    m := &VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder instantiates a new VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewVirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a specific cloudPcOnPremisesConnection object. When you delete an Azure network connection, permissions to the service are removed from the specified Azure resources. You cannot delete an Azure network connection when it's in use, as indicated by the inUse property.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpconpremisesconnection-delete?view=graph-rest-1.0
func (m *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of the cloudPcOnPremisesConnection object.
// returns a CloudPcOnPremisesConnectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpconpremisesconnection-get?view=graph-rest-1.0
func (m *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcOnPremisesConnectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcOnPremisesConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcOnPremisesConnectionable), nil
}
// Patch update the properties of a cloudPcOnPremisesConnection object.
// returns a CloudPcOnPremisesConnectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpconpremisesconnection-update?view=graph-rest-1.0
func (m *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcOnPremisesConnectionable, requestConfiguration *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcOnPremisesConnectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcOnPremisesConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcOnPremisesConnectionable), nil
}
// RunHealthChecks provides operations to call the runHealthChecks method.
// returns a *VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder when successful
func (m *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) RunHealthChecks()(*VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder) {
    return NewVirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a specific cloudPcOnPremisesConnection object. When you delete an Azure network connection, permissions to the service are removed from the specified Azure resources. You cannot delete an Azure network connection when it's in use, as indicated by the inUse property.
// returns a *RequestInformation when successful
func (m *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of the cloudPcOnPremisesConnection object.
// returns a *RequestInformation when successful
func (m *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a cloudPcOnPremisesConnection object.
// returns a *RequestInformation when successful
func (m *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcOnPremisesConnectionable, requestConfiguration *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder when successful
func (m *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    return NewVirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
