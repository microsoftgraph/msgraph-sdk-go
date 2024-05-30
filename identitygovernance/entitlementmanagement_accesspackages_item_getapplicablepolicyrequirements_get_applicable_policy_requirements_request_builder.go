package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder provides operations to call the getApplicablePolicyRequirements method.
type EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/getApplicablePolicyRequirements", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for.
// Deprecated: This method is obsolete. Use PostAsGetApplicablePolicyRequirementsPostResponse instead.
// returns a EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-getapplicablepolicyrequirements?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsResponseable), nil
}
// PostAsGetApplicablePolicyRequirementsPostResponse in Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for.
// returns a EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-getapplicablepolicyrequirements?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder) PostAsGetApplicablePolicyRequirementsPostResponse(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsPostResponseable), nil
}
// ToPostRequestInformation in Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
