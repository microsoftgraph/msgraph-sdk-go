package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder provides operations to manage the stages property of the microsoft.graph.approval entity.
type AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderGetQueryParameters in Azure AD entitlement management, retrieve the properties of an approvalStage object. An approval stage is contained within an approval object.
type AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderGetQueryParameters
}
// AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderInternal instantiates a new ApprovalStageItemRequestBuilder and sets the default values.
func NewAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder) {
    m := &AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/appConsent/appConsentRequests/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}/approval/stages/{approvalStage%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder instantiates a new ApprovalStageItemRequestBuilder and sets the default values.
func NewAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property stages for identityGovernance
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get in Azure AD entitlement management, retrieve the properties of an approvalStage object. An approval stage is contained within an approval object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/approvalstage-get?view=graph-rest-1.0
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// Patch in Azure AD entitlement management, approve or deny an approvalStage object in an approval.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/approvalstage-update?view=graph-rest-1.0
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation in Azure AD entitlement management, retrieve the properties of an approvalStage object. An approval stage is contained within an approval object.
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation in Azure AD entitlement management, approve or deny an approvalStage object in an approval.
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApprovalStageable, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStagesApprovalStageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
