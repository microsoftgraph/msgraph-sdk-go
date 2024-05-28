package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder provides operations to call the getMemberObjects method.
type DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal instantiates a new DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder and sets the default values.
func NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    m := &DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/deletedItems/{directoryObject%2Did}/getMemberObjects", pathParameters),
    }
    return m
}
// NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder instantiates a new DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder and sets the default values.
func NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all IDs for the groups, administrative units, and directory roles that a user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive. Note: Only users and role-enabled groups can be members of directory roles.
// Deprecated: This method is obsolete. Use PostAsGetMemberObjectsPostResponse instead.
// returns a DeleteditemsItemGetmemberobjectsGetMemberObjectsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getmemberobjects?view=graph-rest-1.0
func (m *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) Post(ctx context.Context, body DeleteditemsItemGetmemberobjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration)(DeleteditemsItemGetmemberobjectsGetMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsItemGetmemberobjectsGetMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsItemGetmemberobjectsGetMemberObjectsResponseable), nil
}
// PostAsGetMemberObjectsPostResponse return all IDs for the groups, administrative units, and directory roles that a user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive. Note: Only users and role-enabled groups can be members of directory roles.
// returns a DeleteditemsItemGetmemberobjectsGetMemberObjectsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getmemberobjects?view=graph-rest-1.0
func (m *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) PostAsGetMemberObjectsPostResponse(ctx context.Context, body DeleteditemsItemGetmemberobjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration)(DeleteditemsItemGetmemberobjectsGetMemberObjectsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsItemGetmemberobjectsGetMemberObjectsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsItemGetmemberobjectsGetMemberObjectsPostResponseable), nil
}
// ToPostRequestInformation return all IDs for the groups, administrative units, and directory roles that a user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive. Note: Only users and role-enabled groups can be members of directory roles.
// returns a *RequestInformation when successful
func (m *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeleteditemsItemGetmemberobjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder when successful
func (m *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) WithUrl(rawUrl string)(*DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    return NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
