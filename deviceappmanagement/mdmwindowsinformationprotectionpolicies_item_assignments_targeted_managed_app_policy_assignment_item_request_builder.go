package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.windowsInformationProtection entity.
type MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderGetQueryParameters navigation property to list of security groups targeted for policy.
type MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderGetQueryParameters
}
// MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal instantiates a new MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    m := &MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy%2Did}/assignments/{targetedManagedAppPolicyAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder instantiates a new MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get navigation property to list of security groups targeted for policy.
// returns a TargetedManagedAppPolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppPolicyAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTargetedManagedAppPolicyAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppPolicyAssignmentable), nil
}
// Patch update the navigation property assignments in deviceAppManagement
// returns a TargetedManagedAppPolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppPolicyAssignmentable, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppPolicyAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTargetedManagedAppPolicyAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppPolicyAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property assignments for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation navigation property to list of security groups targeted for policy.
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppPolicyAssignmentable, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*MdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
