package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilder provides operations to call the isPublished method.
type ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilderInternal instantiates a new IsPublishedRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilder) {
    m := &ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/{contentType%2Did}/microsoft.graph.isPublished()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilder instantiates a new IsPublishedRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function isPublished
func (m *ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilderGetRequestConfiguration)(ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponseable), nil
}
// ToGetRequestInformation invoke function isPublished
func (m *ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
