package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackages\{accessPackage-id}\incompatibleGroups\{group-id}\serviceProvisioningErrors
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetQueryParameters errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a group object .  Supports $filter (eq, not, for isResolved and serviceInstance).
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups/{group%2Did}/serviceProvisioningErrors{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) Count()(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a group object .  Supports $filter (eq, not, for isResolved and serviceInstance).
// returns a ServiceProvisioningErrorCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceProvisioningErrorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceProvisioningErrorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceProvisioningErrorCollectionResponseable), nil
}
// ToGetRequestInformation errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a group object .  Supports $filter (eq, not, for isResolved and serviceInstance).
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
