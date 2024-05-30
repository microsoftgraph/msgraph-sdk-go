package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder provides operations to call the addCopy method.
type ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilderInternal instantiates a new ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder) {
    m := &ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/addCopy", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder instantiates a new ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add a copy of a content type from a site to a list.
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-addcopy?view=graph-rest-1.0
func (m *ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder) Post(ctx context.Context, body ItemSitesItemListsItemContenttypesAddcopyAddCopyPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// ToPostRequestInformation add a copy of a content type from a site to a list.
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemListsItemContenttypesAddcopyAddCopyPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
