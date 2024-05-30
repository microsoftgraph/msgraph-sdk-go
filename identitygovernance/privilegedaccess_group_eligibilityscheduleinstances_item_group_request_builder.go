package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupEligibilityScheduleInstance entity.
type PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderGetQueryParameters references the group that is the scope of the membership or ownership eligibility through PIM for groups. Supports $expand.
type PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderGetQueryParameters
}
// NewPrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderInternal instantiates a new PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder) {
    m := &PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/{privilegedAccessGroupEligibilityScheduleInstance%2Did}/group{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder instantiates a new PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get references the group that is the scope of the membership or ownership eligibility through PIM for groups. Supports $expand.
// returns a Groupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
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
// returns a *PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder) ServiceProvisioningErrors()(*PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityscheduleinstancesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation references the group that is the scope of the membership or ownership eligibility through PIM for groups. Supports $expand.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityscheduleinstancesItemGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
