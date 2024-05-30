package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder provides operations to call the associateWithHubSites method.
type ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal instantiates a new ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    m := &ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/associateWithHubSites", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder instantiates a new ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post associate a published content type present in a content type hub with a list of hub sites.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-associatewithhubsites?view=graph-rest-1.0
func (m *ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) Post(ctx context.Context, body ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation associate a published content type present in a content type hub with a list of hub sites.
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
