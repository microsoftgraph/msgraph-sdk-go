package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder builds and executes requests for operations under \identityGovernance\privilegedAccess\group\assignmentScheduleInstances\{privilegedAccessGroupAssignmentScheduleInstance-id}\group\serviceProvisioningErrors
type PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetQueryParameters errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a group object .  Supports $filter (eq, not, for isResolved and serviceInstance).
type PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetQueryParameters struct {
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
// PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetQueryParameters
}
// NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/{privilegedAccessGroupAssignmentScheduleInstance%2Did}/group/serviceProvisioningErrors{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsCountRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) Count()(*PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsCountRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a group object .  Supports $filter (eq, not, for isResolved and serviceInstance).
// returns a ServiceProvisioningErrorCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceProvisioningErrorCollectionResponseable, error) {
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
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
