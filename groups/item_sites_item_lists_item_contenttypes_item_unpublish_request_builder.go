package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder provides operations to call the unpublish method.
type ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemContenttypesItemUnpublishRequestBuilderInternal instantiates a new ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesItemUnpublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder) {
    m := &ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/unpublish", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder instantiates a new ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContenttypesItemUnpublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unpublish a contentType from a content type hub site.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-unpublish?view=graph-rest-1.0
func (m *ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation unpublish a contentType from a content type hub site.
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesItemUnpublishRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
