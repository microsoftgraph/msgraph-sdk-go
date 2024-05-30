package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder provides operations to manage the customExtension property of the microsoft.graph.customExtensionStageSetting entity.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters indicates the custom workflow extension that will be executed at this stage. Nullable. Supports $expand.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionStageSettings/{customExtensionStageSetting%2Did}/customExtension{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get indicates the custom workflow extension that will be executed at this stage. Nullable. Supports $expand.
// returns a CustomCalloutExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCustomCalloutExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable), nil
}
// ToGetRequestInformation indicates the custom workflow extension that will be executed at this stage. Nullable. Supports $expand.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
