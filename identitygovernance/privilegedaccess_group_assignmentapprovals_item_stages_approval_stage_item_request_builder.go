package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder provides operations to manage the stages property of the microsoft.graph.approval entity.
type PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderGetQueryParameters a collection of stages in the approval decision.
type PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentApprovals/{approval%2Did}/stages/{approvalStage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property stages for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of stages in the approval decision.
// returns a ApprovalStageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApprovalStageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable), nil
}
// Patch update the navigation property stages in identityGovernance
// returns a ApprovalStageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable, requestConfiguration *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApprovalStageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable), nil
}
// ToDeleteRequestInformation delete navigation property stages for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of stages in the approval decision.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property stages in identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable, requestConfiguration *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentapprovalsItemStagesApprovalStageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
