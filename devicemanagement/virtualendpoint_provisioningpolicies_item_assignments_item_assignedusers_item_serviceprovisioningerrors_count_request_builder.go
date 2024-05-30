package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder provides operations to count the resources in the collection.
type VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderGetQueryParameters get the number of the resource
type VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderGetQueryParameters
}
// NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderInternal instantiates a new VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder and sets the default values.
func NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder) {
    m := &VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy%2Did}/assignments/{cloudPcProvisioningPolicyAssignment%2Did}/assignedUsers/{user%2Did}/serviceProvisioningErrors/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder instantiates a new VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder and sets the default values.
func NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder) {
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersItemServiceprovisioningerrorsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
