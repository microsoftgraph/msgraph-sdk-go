package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder provides operations to manage the windowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
type WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderGetQueryParameters read properties and relationships of the windowsInformationProtectionPolicy object.
type WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderGetQueryParameters
}
// WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsInformationProtection entity.
// returns a *WindowsinformationprotectionpoliciesItemAssignmentsRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) Assignments()(*WindowsinformationprotectionpoliciesItemAssignmentsRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderInternal instantiates a new WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) {
    m := &WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder instantiates a new WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a windowsInformationProtectionPolicy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-windowsinformationprotectionpolicy-delete?view=graph-rest-1.0
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExemptAppLockerFiles provides operations to manage the exemptAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
// returns a *WindowsinformationprotectionpoliciesItemExemptapplockerfilesExemptAppLockerFilesRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) ExemptAppLockerFiles()(*WindowsinformationprotectionpoliciesItemExemptapplockerfilesExemptAppLockerFilesRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesItemExemptapplockerfilesExemptAppLockerFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the windowsInformationProtectionPolicy object.
// returns a WindowsInformationProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-windowsinformationprotectionpolicy-get?view=graph-rest-1.0
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable), nil
}
// Patch update the properties of a windowsInformationProtectionPolicy object.
// returns a WindowsInformationProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-windowsinformationprotectionpolicy-update?view=graph-rest-1.0
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable), nil
}
// ProtectedAppLockerFiles provides operations to manage the protectedAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
// returns a *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) ProtectedAppLockerFiles()(*WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a windowsInformationProtectionPolicy.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the windowsInformationProtectionPolicy object.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a windowsInformationProtectionPolicy object.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
