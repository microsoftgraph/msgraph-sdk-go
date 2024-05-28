package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder provides operations to manage the provisioningPolicies property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderGetQueryParameters read the properties and relationships of a cloudPcProvisioningPolicy object.
type VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderGetQueryParameters
}
// VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *VirtualendpointProvisioningpoliciesItemAssignRequestBuilder when successful
func (m *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) Assign()(*VirtualendpointProvisioningpoliciesItemAssignRequestBuilder) {
    return NewVirtualendpointProvisioningpoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.cloudPcProvisioningPolicy entity.
// returns a *VirtualendpointProvisioningpoliciesItemAssignmentsRequestBuilder when successful
func (m *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) Assignments()(*VirtualendpointProvisioningpoliciesItemAssignmentsRequestBuilder) {
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderInternal instantiates a new VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder and sets the default values.
func NewVirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) {
    m := &VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder instantiates a new VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder and sets the default values.
func NewVirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a cloudPcProvisioningPolicy object. You can’t delete a policy that’s in use.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcprovisioningpolicy-delete?view=graph-rest-1.0
func (m *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a cloudPcProvisioningPolicy object.
// returns a CloudPcProvisioningPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcprovisioningpolicy-get?view=graph-rest-1.0
func (m *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcProvisioningPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcProvisioningPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcProvisioningPolicyable), nil
}
// Patch update the properties of a cloudPcProvisioningPolicy object.
// returns a CloudPcProvisioningPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcprovisioningpolicy-update?view=graph-rest-1.0
func (m *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcProvisioningPolicyable, requestConfiguration *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcProvisioningPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcProvisioningPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcProvisioningPolicyable), nil
}
// ToDeleteRequestInformation delete a cloudPcProvisioningPolicy object. You can’t delete a policy that’s in use.
// returns a *RequestInformation when successful
func (m *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a cloudPcProvisioningPolicy object.
// returns a *RequestInformation when successful
func (m *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a cloudPcProvisioningPolicy object.
// returns a *RequestInformation when successful
func (m *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcProvisioningPolicyable, requestConfiguration *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder when successful
func (m *VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder) {
    return NewVirtualendpointProvisioningpoliciesCloudPcProvisioningPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
