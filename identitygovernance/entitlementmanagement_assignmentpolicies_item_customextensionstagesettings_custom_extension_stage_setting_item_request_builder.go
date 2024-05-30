package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder provides operations to manage the customExtensionStageSettings property of the microsoft.graph.accessPackageAssignmentPolicy entity.
type EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderGetQueryParameters the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
type EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) {
    m := &EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionStageSettings/{customExtensionStageSetting%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder instantiates a new EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomExtension provides operations to manage the customExtension property of the microsoft.graph.customExtensionStageSetting entity.
// returns a *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder when successful
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) CustomExtension()(*EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    return NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property customExtensionStageSettings for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// returns a CustomExtensionStageSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCustomExtensionStageSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable), nil
}
// Patch update the navigation property customExtensionStageSettings in identityGovernance
// returns a CustomExtensionStageSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable, requestConfiguration *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCustomExtensionStageSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable), nil
}
// ToDeleteRequestInformation delete navigation property customExtensionStageSettings for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property customExtensionStageSettings in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable, requestConfiguration *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder when successful
func (m *EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) {
    return NewEntitlementmanagementAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
