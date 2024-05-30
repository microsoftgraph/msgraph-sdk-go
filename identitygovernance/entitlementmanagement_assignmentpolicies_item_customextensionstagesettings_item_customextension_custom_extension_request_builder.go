package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder provides operations to manage the customExtension property of the microsoft.graph.customExtensionStageSetting entity.
type EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters indicates the custom workflow extension that will be executed at this stage. Nullable. Supports $expand.
type EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    m := &EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionStageSettings/{customExtensionStageSetting%2Did}/customExtension{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder instantiates a new EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get indicates the custom workflow extension that will be executed at this stage. Nullable. Supports $expand.
// returns a CustomCalloutExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, error) {
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
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder when successful
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    return NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
