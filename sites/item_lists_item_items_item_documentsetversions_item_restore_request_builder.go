package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder provides operations to call the restore method.
type ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderInternal instantiates a new ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder and sets the default values.
func NewItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) {
    m := &ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/documentSetVersions/{documentSetVersion%2Did}/restore", pathParameters),
    }
    return m
}
// NewItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder instantiates a new ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder and sets the default values.
func NewItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a document set version.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/documentsetversion-restore?view=graph-rest-1.0
func (m *ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation restore a document set version.
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder when successful
func (m *ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) {
    return NewItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
