package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder provides operations to call the copyToDefaultContentLocation method.
type ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderInternal instantiates a new ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) {
    m := &ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/copyToDefaultContentLocation", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder instantiates a new ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post copy a file to a default content location in a content type. The file can then be added as a default file or template via a POST operation.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-copytodefaultcontentlocation?view=graph-rest-1.0
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) Post(ctx context.Context, body ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation copy a file to a default content location in a content type. The file can then be added as a default file or template via a POST operation.
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
