package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder provides operations to manage the excludes property of the microsoft.graph.permissionGrantPolicy entity.
type PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetQueryParameters condition sets that are excluded in this permission grant policy. Automatically expanded on GET.
type PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetQueryParameters
}
// PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderInternal instantiates a new PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder and sets the default values.
func NewPermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) {
    m := &PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy%2Did}/excludes/{permissionGrantConditionSet%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder instantiates a new PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder and sets the default values.
func NewPermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a permissionGrantConditionSet from the excludes collection of a permissionGrantPolicy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissiongrantpolicy-delete-excludes?view=graph-rest-1.0
func (m *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get condition sets that are excluded in this permission grant policy. Automatically expanded on GET.
// returns a PermissionGrantConditionSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantConditionSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable), nil
}
// Patch update the navigation property excludes in policies
// returns a PermissionGrantConditionSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantConditionSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable), nil
}
// ToDeleteRequestInformation deletes a permissionGrantConditionSet from the excludes collection of a permissionGrantPolicy.
// returns a *RequestInformation when successful
func (m *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation condition sets that are excluded in this permission grant policy. Automatically expanded on GET.
// returns a *RequestInformation when successful
func (m *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property excludes in policies
// returns a *RequestInformation when successful
func (m *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder when successful
func (m *PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) WithUrl(rawUrl string)(*PermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) {
    return NewPermissiongrantpoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
