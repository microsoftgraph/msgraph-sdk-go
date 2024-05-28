package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGetmembergroupsGetMemberGroupsRequestBuilder provides operations to call the getMemberGroups method.
type ItemGetmembergroupsGetMemberGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetmembergroupsGetMemberGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetmembergroupsGetMemberGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGetmembergroupsGetMemberGroupsRequestBuilderInternal instantiates a new ItemGetmembergroupsGetMemberGroupsRequestBuilder and sets the default values.
func NewItemGetmembergroupsGetMemberGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmembergroupsGetMemberGroupsRequestBuilder) {
    m := &ItemGetmembergroupsGetMemberGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/getMemberGroups", pathParameters),
    }
    return m
}
// NewItemGetmembergroupsGetMemberGroupsRequestBuilder instantiates a new ItemGetmembergroupsGetMemberGroupsRequestBuilder and sets the default values.
func NewItemGetmembergroupsGetMemberGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmembergroupsGetMemberGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetmembergroupsGetMemberGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all the group IDs for the groups that the specified user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive. This API returns up to 11,000 group IDs. If more than 11,000 results are available, it returns a 400 Bad Request error with the Directory_ResultSizeLimitExceeded error code. As a workaround, use the List group transitive memberOf API.
// Deprecated: This method is obsolete. Use PostAsGetMemberGroupsPostResponse instead.
// returns a ItemGetmembergroupsGetMemberGroupsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getmembergroups?view=graph-rest-1.0
func (m *ItemGetmembergroupsGetMemberGroupsRequestBuilder) Post(ctx context.Context, body ItemGetmembergroupsGetMemberGroupsPostRequestBodyable, requestConfiguration *ItemGetmembergroupsGetMemberGroupsRequestBuilderPostRequestConfiguration)(ItemGetmembergroupsGetMemberGroupsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmembergroupsGetMemberGroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmembergroupsGetMemberGroupsResponseable), nil
}
// PostAsGetMemberGroupsPostResponse return all the group IDs for the groups that the specified user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive. This API returns up to 11,000 group IDs. If more than 11,000 results are available, it returns a 400 Bad Request error with the Directory_ResultSizeLimitExceeded error code. As a workaround, use the List group transitive memberOf API.
// returns a ItemGetmembergroupsGetMemberGroupsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getmembergroups?view=graph-rest-1.0
func (m *ItemGetmembergroupsGetMemberGroupsRequestBuilder) PostAsGetMemberGroupsPostResponse(ctx context.Context, body ItemGetmembergroupsGetMemberGroupsPostRequestBodyable, requestConfiguration *ItemGetmembergroupsGetMemberGroupsRequestBuilderPostRequestConfiguration)(ItemGetmembergroupsGetMemberGroupsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmembergroupsGetMemberGroupsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmembergroupsGetMemberGroupsPostResponseable), nil
}
// ToPostRequestInformation return all the group IDs for the groups that the specified user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive. This API returns up to 11,000 group IDs. If more than 11,000 results are available, it returns a 400 Bad Request error with the Directory_ResultSizeLimitExceeded error code. As a workaround, use the List group transitive memberOf API.
// returns a *RequestInformation when successful
func (m *ItemGetmembergroupsGetMemberGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemGetmembergroupsGetMemberGroupsPostRequestBodyable, requestConfiguration *ItemGetmembergroupsGetMemberGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGetmembergroupsGetMemberGroupsRequestBuilder when successful
func (m *ItemGetmembergroupsGetMemberGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemGetmembergroupsGetMemberGroupsRequestBuilder) {
    return NewItemGetmembergroupsGetMemberGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
