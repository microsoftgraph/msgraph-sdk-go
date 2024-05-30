package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder provides operations to manage the assignmentPolicies property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters in Microsoft Entra entitlement management, retrieve the properties and relationships of an accessPackageAssignmentPolicy object.
type EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// returns a *EntitlementmanagementAssignmentpoliciesItemAccesspackageAccessPackageRequestBuilder when successful
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) AccessPackage()(*EntitlementmanagementAssignmentpoliciesItemAccesspackageAccessPackageRequestBuilder) {
    return NewEntitlementmanagementAssignmentpoliciesItemAccesspackageAccessPackageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Catalog provides operations to manage the catalog property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// returns a *EntitlementmanagementAssignmentpoliciesItemCatalogRequestBuilder when successful
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Catalog()(*EntitlementmanagementAssignmentpoliciesItemCatalogRequestBuilder) {
    return NewEntitlementmanagementAssignmentpoliciesItemCatalogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    m := &EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder instantiates a new EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomExtensionStageSettings provides operations to manage the customExtensionStageSettings property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// returns a *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder when successful
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CustomExtensionStageSettings()(*EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) {
    return NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete in Microsoft Entra entitlement management, delete an accessPackageAssignmentPolicy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentpolicy-delete?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get in Microsoft Entra entitlement management, retrieve the properties and relationships of an accessPackageAssignmentPolicy object.
// returns a AccessPackageAssignmentPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentpolicy-get?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable), nil
}
// Put update an existing accessPackageAssignmentPolicy object to change one or more of its properties, such as the display name or description.
// returns a AccessPackageAssignmentPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentpolicy-update?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Put(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, requestConfiguration *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable), nil
}
// Questions provides operations to manage the questions property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// returns a *EntitlementmanagementAssignmentpoliciesItemQuestionsRequestBuilder when successful
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Questions()(*EntitlementmanagementAssignmentpoliciesItemQuestionsRequestBuilder) {
    return NewEntitlementmanagementAssignmentpoliciesItemQuestionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation in Microsoft Entra entitlement management, delete an accessPackageAssignmentPolicy.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation in Microsoft Entra entitlement management, retrieve the properties and relationships of an accessPackageAssignmentPolicy object.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation update an existing accessPackageAssignmentPolicy object to change one or more of its properties, such as the display name or description.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, requestConfiguration *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder when successful
func (m *EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    return NewEntitlementmanagementAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
