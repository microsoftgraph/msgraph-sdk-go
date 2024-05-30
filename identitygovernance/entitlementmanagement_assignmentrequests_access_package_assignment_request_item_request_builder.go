package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder provides operations to manage the assignmentRequests property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters in Microsoft Entra entitlement management, retrieve the properties and relationships of an  accessPackageAssignmentRequest object.
type EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentRequest entity.
// returns a *EntitlementmanagementAssignmentrequestsItemAccesspackageAccessPackageRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) AccessPackage()(*EntitlementmanagementAssignmentrequestsItemAccesspackageAccessPackageRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemAccesspackageAccessPackageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignment provides operations to manage the assignment property of the microsoft.graph.accessPackageAssignmentRequest entity.
// returns a *EntitlementmanagementAssignmentrequestsItemAssignmentRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) Assignment()(*EntitlementmanagementAssignmentrequestsItemAssignmentRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *EntitlementmanagementAssignmentrequestsItemCancelRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) Cancel()(*EntitlementmanagementAssignmentrequestsItemCancelRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    m := &EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentRequests/{accessPackageAssignmentRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder instantiates a new EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an accessPackageAssignmentRequest object. This request can be made to remove a denied or completed request.  You cannot delete an access package assignment request if it has any accessPackageAssignment objects.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentrequest-delete?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get in Microsoft Entra entitlement management, retrieve the properties and relationships of an  accessPackageAssignmentRequest object.
// returns a AccessPackageAssignmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentrequest-get?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable), nil
}
// Patch update the navigation property assignmentRequests in identityGovernance
// returns a AccessPackageAssignmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable, requestConfiguration *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable), nil
}
// Reprocess provides operations to call the reprocess method.
// returns a *EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) Reprocess()(*EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemReprocessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Requestor provides operations to manage the requestor property of the microsoft.graph.accessPackageAssignmentRequest entity.
// returns a *EntitlementmanagementAssignmentrequestsItemRequestorRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) Requestor()(*EntitlementmanagementAssignmentrequestsItemRequestorRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemRequestorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resume provides operations to call the resume method.
// returns a *EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) Resume()(*EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemResumeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete an accessPackageAssignmentRequest object. This request can be made to remove a denied or completed request.  You cannot delete an access package assignment request if it has any accessPackageAssignment objects.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation in Microsoft Entra entitlement management, retrieve the properties and relationships of an  accessPackageAssignmentRequest object.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignmentRequests in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable, requestConfiguration *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsAccessPackageAssignmentRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
