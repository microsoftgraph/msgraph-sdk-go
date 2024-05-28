package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder provides operations to manage the activatedUsing property of the microsoft.graph.privilegedAccessGroupAssignmentSchedule entity.
type PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderGetQueryParameters when the request activates an ownership or membership assignment in PIM for groups, this object represents the eligibility relationship. Otherwise, it's null. Supports $expand.
type PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderGetQueryParameters
}
// NewPrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentSchedules/{privilegedAccessGroupAssignmentSchedule%2Did}/activatedUsing{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderInternal(urlParams, requestAdapter)
}
// Get when the request activates an ownership or membership assignment in PIM for groups, this object represents the eligibility relationship. Otherwise, it's null. Supports $expand.
// returns a PrivilegedAccessGroupEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable), nil
}
// ToGetRequestInformation when the request activates an ownership or membership assignment in PIM for groups, this object represents the eligibility relationship. Otherwise, it's null. Supports $expand.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
