package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder provides operations to call the isPublished method.
type ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal instantiates a new ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder and sets the default values.
func NewItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    m := &ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/isPublished()", pathParameters),
    }
    return m
}
// NewItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder instantiates a new ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder and sets the default values.
func NewItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get check the publishing status of a contentType in a content type hub site.
// Deprecated: This method is obsolete. Use GetAsIsPublishedGetResponse instead.
// returns a ItemListsItemContenttypesItemIspublishedIsPublishedResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-1.0
func (m *ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(ItemListsItemContenttypesItemIspublishedIsPublishedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListsItemContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListsItemContenttypesItemIspublishedIsPublishedResponseable), nil
}
// GetAsIsPublishedGetResponse check the publishing status of a contentType in a content type hub site.
// returns a ItemListsItemContenttypesItemIspublishedIsPublishedGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-1.0
func (m *ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) GetAsIsPublishedGetResponse(ctx context.Context, requestConfiguration *ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(ItemListsItemContenttypesItemIspublishedIsPublishedGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListsItemContenttypesItemIspublishedIsPublishedGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListsItemContenttypesItemIspublishedIsPublishedGetResponseable), nil
}
// ToGetRequestInformation check the publishing status of a contentType in a content type hub site.
// returns a *RequestInformation when successful
func (m *ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder when successful
func (m *ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    return NewItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
