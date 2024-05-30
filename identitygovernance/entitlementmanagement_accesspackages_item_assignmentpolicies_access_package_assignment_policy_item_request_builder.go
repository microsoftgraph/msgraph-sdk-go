package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder provides operations to manage the assignmentPolicies property of the microsoft.graph.accessPackage entity.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemAccesspackageAccessPackageRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) AccessPackage()(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemAccesspackageAccessPackageRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemAccesspackageAccessPackageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Catalog provides operations to manage the catalog property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Catalog()(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomExtensionStageSettings provides operations to manage the customExtensionStageSettings property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CustomExtensionStageSettings()(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property assignmentPolicies for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Nullable. Supports $expand.
// returns a AccessPackageAssignmentPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, error) {
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
// Patch update the navigation property assignmentPolicies in identityGovernance
// returns a AccessPackageAssignmentPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemQuestionsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Questions()(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemQuestionsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemQuestionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignmentPolicies for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. Supports $expand.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignmentPolicies in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesAccessPackageAssignmentPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
