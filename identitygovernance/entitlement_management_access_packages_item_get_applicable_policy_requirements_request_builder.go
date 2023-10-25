package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder provides operations to call the getApplicablePolicyRequirements method.
type EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilderInternal instantiates a new GetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/getApplicablePolicyRequirements", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder instantiates a new GetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for. This API is available in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsGetApplicablePolicyRequirementsPostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-getapplicablepolicyrequirements?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponseable), nil
}
// PostAsGetApplicablePolicyRequirementsPostResponse in Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-getapplicablepolicyrequirements?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder) PostAsGetApplicablePolicyRequirementsPostResponse(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsPostResponseable), nil
}
// ToPostRequestInformation in Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for. This API is available in the following national cloud deployments.
func (m *EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
