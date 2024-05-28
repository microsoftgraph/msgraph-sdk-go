package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance entity.
type PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderGetQueryParameters references the group that is the scope of the membership or ownership assignment through PIM for groups. Supports $expand.
type PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderGetQueryParameters
}
// NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/{privilegedAccessGroupAssignmentScheduleInstance%2Did}/group{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get references the group that is the scope of the membership or ownership assignment through PIM for groups. Supports $expand.
// returns a Groupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable), nil
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder) ServiceProvisioningErrors()(*PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation references the group that is the scope of the membership or ownership assignment through PIM for groups. Supports $expand.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
