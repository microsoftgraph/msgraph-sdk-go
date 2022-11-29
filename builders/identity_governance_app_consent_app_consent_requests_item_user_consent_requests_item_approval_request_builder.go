package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder provides operations to manage the approval property of the microsoft.graph.userConsentRequest entity.
type IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderGetQueryParameters approval decisions associated with a request.
type IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderGetQueryParameters
}
// IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderInternal instantiates a new ApprovalRequestBuilder and sets the default values.
func NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) {
    m := &IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/appConsent/appConsentRequests/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}/approval{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder instantiates a new ApprovalRequestBuilder and sets the default values.
func NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property approval for identityGovernance
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation approval decisions associated with a request.
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property approval in identityGovernance
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property approval for identityGovernance
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get approval decisions associated with a request.
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApprovalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable), nil
}
// Patch update the navigation property approval in identityGovernance
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApprovalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable), nil
}
// Stages provides operations to manage the stages property of the microsoft.graph.approval entity.
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) Stages()(*IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesRequestBuilder) {
    return NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById provides operations to manage the stages property of the microsoft.graph.approval entity.
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder) StagesById(id string)(*IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approvalStage%2Did"] = id
    }
    return NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
