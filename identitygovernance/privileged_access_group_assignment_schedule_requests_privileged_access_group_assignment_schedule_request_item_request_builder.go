package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder provides operations to manage the assignmentScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetQueryParameters get assignmentScheduleRequests from identityGovernance
type PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest entity.
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*PrivilegedAccessGroupAssignmentScheduleRequestsItemActivatedUsingRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsItemActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Cancel()(*PrivilegedAccessGroupAssignmentScheduleRequestsItemCancelRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new PrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) {
    m := &PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/{privilegedAccessGroupAssignmentScheduleRequest%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder instantiates a new PrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentScheduleRequests for identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get assignmentScheduleRequests from identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest entity.
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Group()(*PrivilegedAccessGroupAssignmentScheduleRequestsItemGroupRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property assignmentScheduleRequests in identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest entity.
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Principal()(*PrivilegedAccessGroupAssignmentScheduleRequestsItemPrincipalRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest entity.
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*PrivilegedAccessGroupAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsItemTargetScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignmentScheduleRequests for identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get assignmentScheduleRequests from identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property assignmentScheduleRequests in identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
