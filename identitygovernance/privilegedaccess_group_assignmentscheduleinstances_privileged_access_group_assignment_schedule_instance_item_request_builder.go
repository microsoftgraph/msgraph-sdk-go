package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder provides operations to manage the assignmentScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters read the properties and relationships of a privilegedAccessGroupAssignmentScheduleInstance object.
type PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance entity.
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) ActivatedUsing()(*PrivilegedaccessGroupAssignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/{privilegedAccessGroupAssignmentScheduleInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentScheduleInstances for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a privilegedAccessGroupAssignmentScheduleInstance object.
// returns a PrivilegedAccessGroupAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupassignmentscheduleinstance-get?view=graph-rest-1.0
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance entity.
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Group()(*PrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property assignmentScheduleInstances in identityGovernance
// returns a PrivilegedAccessGroupAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance entity.
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesItemPrincipalRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Principal()(*PrivilegedaccessGroupAssignmentscheduleinstancesItemPrincipalRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignmentScheduleInstances for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a privilegedAccessGroupAssignmentScheduleInstance object.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignmentScheduleInstances in identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
