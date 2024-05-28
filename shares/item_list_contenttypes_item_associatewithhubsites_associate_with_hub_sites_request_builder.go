package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder provides operations to call the associateWithHubSites method.
type ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal instantiates a new ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    m := &ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shares/{sharedDriveItem%2Did}/list/contentTypes/{contentType%2Did}/associateWithHubSites", pathParameters),
    }
    return m
}
// NewItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder instantiates a new ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post associate a published content type present in a content type hub with a list of hub sites.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-associatewithhubsites?view=graph-rest-1.0
func (m *ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) Post(ctx context.Context, body ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder when successful
func (m *ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) WithUrl(rawUrl string)(*ItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    return NewItemListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
