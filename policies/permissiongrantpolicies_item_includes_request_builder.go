package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PermissiongrantpoliciesItemIncludesRequestBuilder provides operations to manage the includes property of the microsoft.graph.permissionGrantPolicy entity.
type PermissiongrantpoliciesItemIncludesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissiongrantpoliciesItemIncludesRequestBuilderGetQueryParameters retrieve the condition sets which are *included* in a permissionGrantPolicy.
type PermissiongrantpoliciesItemIncludesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PermissiongrantpoliciesItemIncludesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissiongrantpoliciesItemIncludesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissiongrantpoliciesItemIncludesRequestBuilderGetQueryParameters
}
// PermissiongrantpoliciesItemIncludesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissiongrantpoliciesItemIncludesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPermissionGrantConditionSetId provides operations to manage the includes property of the microsoft.graph.permissionGrantPolicy entity.
// returns a *PermissiongrantpoliciesItemIncludesPermissionGrantConditionSetItemRequestBuilder when successful
func (m *PermissiongrantpoliciesItemIncludesRequestBuilder) ByPermissionGrantConditionSetId(permissionGrantConditionSetId string)(*PermissiongrantpoliciesItemIncludesPermissionGrantConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if permissionGrantConditionSetId != "" {
        urlTplParams["permissionGrantConditionSet%2Did"] = permissionGrantConditionSetId
    }
    return NewPermissiongrantpoliciesItemIncludesPermissionGrantConditionSetItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissiongrantpoliciesItemIncludesRequestBuilderInternal instantiates a new PermissiongrantpoliciesItemIncludesRequestBuilder and sets the default values.
func NewPermissiongrantpoliciesItemIncludesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissiongrantpoliciesItemIncludesRequestBuilder) {
    m := &PermissiongrantpoliciesItemIncludesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy%2Did}/includes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPermissiongrantpoliciesItemIncludesRequestBuilder instantiates a new PermissiongrantpoliciesItemIncludesRequestBuilder and sets the default values.
func NewPermissiongrantpoliciesItemIncludesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissiongrantpoliciesItemIncludesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissiongrantpoliciesItemIncludesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PermissiongrantpoliciesItemIncludesCountRequestBuilder when successful
func (m *PermissiongrantpoliciesItemIncludesRequestBuilder) Count()(*PermissiongrantpoliciesItemIncludesCountRequestBuilder) {
    return NewPermissiongrantpoliciesItemIncludesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the condition sets which are *included* in a permissionGrantPolicy.
// returns a PermissionGrantConditionSetCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissiongrantpolicy-list-includes?view=graph-rest-1.0
func (m *PermissiongrantpoliciesItemIncludesRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissiongrantpoliciesItemIncludesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantConditionSetCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetCollectionResponseable), nil
}
// Post add conditions under which a permission grant event is *included* in a permission grant policy. You do this by adding a permissionGrantConditionSet to the includes collection of a  permissionGrantPolicy.
// returns a PermissionGrantConditionSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissiongrantpolicy-post-includes?view=graph-rest-1.0
func (m *PermissiongrantpoliciesItemIncludesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissiongrantpoliciesItemIncludesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation retrieve the condition sets which are *included* in a permissionGrantPolicy.
// returns a *RequestInformation when successful
func (m *PermissiongrantpoliciesItemIncludesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissiongrantpoliciesItemIncludesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add conditions under which a permission grant event is *included* in a permission grant policy. You do this by adding a permissionGrantConditionSet to the includes collection of a  permissionGrantPolicy.
// returns a *RequestInformation when successful
func (m *PermissiongrantpoliciesItemIncludesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissiongrantpoliciesItemIncludesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *PermissiongrantpoliciesItemIncludesRequestBuilder when successful
func (m *PermissiongrantpoliciesItemIncludesRequestBuilder) WithUrl(rawUrl string)(*PermissiongrantpoliciesItemIncludesRequestBuilder) {
    return NewPermissiongrantpoliciesItemIncludesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
