package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
type PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderGetQueryParameters retrieve a single permissionGrantPolicy object.
type PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderGetQueryParameters
}
// PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderInternal instantiates a new PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) {
    m := &PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder instantiates a new PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a permissionGrantPolicy object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissiongrantpolicy-delete?view=graph-rest-1.0
func (m *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Excludes provides operations to manage the excludes property of the microsoft.graph.permissionGrantPolicy entity.
// returns a *PermissiongrantpoliciesItemExcludesRequestBuilder when successful
func (m *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) Excludes()(*PermissiongrantpoliciesItemExcludesRequestBuilder) {
    return NewPermissiongrantpoliciesItemExcludesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a single permissionGrantPolicy object.
// returns a PermissionGrantPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissiongrantpolicy-get?view=graph-rest-1.0
func (m *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable), nil
}
// Includes provides operations to manage the includes property of the microsoft.graph.permissionGrantPolicy entity.
// returns a *PermissiongrantpoliciesItemIncludesRequestBuilder when successful
func (m *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) Includes()(*PermissiongrantpoliciesItemIncludesRequestBuilder) {
    return NewPermissiongrantpoliciesItemIncludesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update properties of a  permissionGrantPolicy.
// returns a PermissionGrantPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissiongrantpolicy-update?view=graph-rest-1.0
func (m *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable, requestConfiguration *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable), nil
}
// ToDeleteRequestInformation delete a permissionGrantPolicy object.
// returns a *RequestInformation when successful
func (m *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a single permissionGrantPolicy object.
// returns a *RequestInformation when successful
func (m *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update properties of a  permissionGrantPolicy.
// returns a *RequestInformation when successful
func (m *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable, requestConfiguration *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder when successful
func (m *PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) WithUrl(rawUrl string)(*PermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder) {
    return NewPermissiongrantpoliciesPermissionGrantPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
