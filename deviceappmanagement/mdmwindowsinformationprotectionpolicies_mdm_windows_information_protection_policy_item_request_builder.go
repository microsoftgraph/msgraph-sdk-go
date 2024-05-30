package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder provides operations to manage the mdmWindowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
type MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetQueryParameters read properties and relationships of the mdmWindowsInformationProtectionPolicy object.
type MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetQueryParameters
}
// MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsInformationProtection entity.
// returns a *MdmwindowsinformationprotectionpoliciesItemAssignmentsRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) Assignments()(*MdmwindowsinformationprotectionpoliciesItemAssignmentsRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal instantiates a new MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) {
    m := &MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder instantiates a new MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a mdmWindowsInformationProtectionPolicy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-mdmwindowsinformationprotectionpolicy-delete?view=graph-rest-1.0
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesExemptAppLockerFilesRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) ExemptAppLockerFiles()(*MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesExemptAppLockerFilesRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesExemptAppLockerFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the mdmWindowsInformationProtectionPolicy object.
// returns a MdmWindowsInformationProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-mdmwindowsinformationprotectionpolicy-get?view=graph-rest-1.0
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable), nil
}
// Patch update the properties of a mdmWindowsInformationProtectionPolicy object.
// returns a MdmWindowsInformationProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-mdmwindowsinformationprotectionpolicy-update?view=graph-rest-1.0
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable, requestConfiguration *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable), nil
}
// ProtectedAppLockerFiles provides operations to manage the protectedAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
// returns a *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) ProtectedAppLockerFiles()(*MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a mdmWindowsInformationProtectionPolicy.
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the mdmWindowsInformationProtectionPolicy object.
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a mdmWindowsInformationProtectionPolicy object.
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable, requestConfiguration *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) WithUrl(rawUrl string)(*MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
