package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder provides operations to call the checkMemberGroups method.
type DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilderInternal instantiates a new DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder and sets the default values.
func NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder) {
    m := &DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/deletedItems/{directoryObject%2Did}/checkMemberGroups", pathParameters),
    }
    return m
}
// NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder instantiates a new DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder and sets the default values.
func NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post check for membership in a specified list of group IDs, and return from that list those groups (identified by IDs) of which the specified user, group, service principal, organizational contact, device, or directory object is a member. This function is transitive. You can check up to a maximum of 20 groups per request. This function supports all groups provisioned in Microsoft Entra ID. Because Microsoft 365 groups cannot contain other groups, membership in a Microsoft 365 group is always direct.
// Deprecated: This method is obsolete. Use PostAsCheckMemberGroupsPostResponse instead.
// returns a DeleteditemsItemCheckmembergroupsCheckMemberGroupsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-checkmembergroups?view=graph-rest-1.0
func (m *DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder) Post(ctx context.Context, body DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostRequestBodyable, requestConfiguration *DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilderPostRequestConfiguration)(DeleteditemsItemCheckmembergroupsCheckMemberGroupsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsItemCheckmembergroupsCheckMemberGroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsItemCheckmembergroupsCheckMemberGroupsResponseable), nil
}
// PostAsCheckMemberGroupsPostResponse check for membership in a specified list of group IDs, and return from that list those groups (identified by IDs) of which the specified user, group, service principal, organizational contact, device, or directory object is a member. This function is transitive. You can check up to a maximum of 20 groups per request. This function supports all groups provisioned in Microsoft Entra ID. Because Microsoft 365 groups cannot contain other groups, membership in a Microsoft 365 group is always direct.
// returns a DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-checkmembergroups?view=graph-rest-1.0
func (m *DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder) PostAsCheckMemberGroupsPostResponse(ctx context.Context, body DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostRequestBodyable, requestConfiguration *DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilderPostRequestConfiguration)(DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponseable), nil
}
// ToPostRequestInformation check for membership in a specified list of group IDs, and return from that list those groups (identified by IDs) of which the specified user, group, service principal, organizational contact, device, or directory object is a member. This function is transitive. You can check up to a maximum of 20 groups per request. This function supports all groups provisioned in Microsoft Entra ID. Because Microsoft 365 groups cannot contain other groups, membership in a Microsoft 365 group is always direct.
// returns a *RequestInformation when successful
func (m *DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostRequestBodyable, requestConfiguration *DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder when successful
func (m *DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder) WithUrl(rawUrl string)(*DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder) {
    return NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
