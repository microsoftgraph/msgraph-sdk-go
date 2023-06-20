package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilder provides operations to call the removeGroup method.
type ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilderInternal instantiates a new RemoveGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilder) {
    m := &ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/groupLifecyclePolicies/{groupLifecyclePolicy%2Did}/removeGroup", pathParameters),
    }
    return m
}
// NewItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilder instantiates a new RemoveGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post removes a group from a lifecycle policy.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/grouplifecyclepolicy-removegroup?view=graph-rest-1.0
func (m *ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilder) Post(ctx context.Context, body ItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilderPostRequestConfiguration)(ItemGroupLifecyclePoliciesItemRemoveGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGroupLifecyclePoliciesItemRemoveGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGroupLifecyclePoliciesItemRemoveGroupResponseable), nil
}
// ToPostRequestInformation removes a group from a lifecycle policy.
func (m *ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesItemRemoveGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
