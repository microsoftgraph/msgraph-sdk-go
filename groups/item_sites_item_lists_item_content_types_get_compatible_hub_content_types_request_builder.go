package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder provides operations to call the getCompatibleHubContentTypes method.
type ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters invoke function getCompatibleHubContentTypes
type ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters
}
// NewItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderInternal instantiates a new GetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder) {
    m := &ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/getCompatibleHubContentTypes(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder instantiates a new GetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCompatibleHubContentTypes
// Deprecated: This method is obsolete. Use GetAsGetCompatibleHubContentTypesGetResponse instead.
func (m *ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesResponseable), nil
}
// GetAsGetCompatibleHubContentTypesGetResponse invoke function getCompatibleHubContentTypes
func (m *ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder) GetAsGetCompatibleHubContentTypesGetResponse(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesGetResponseable), nil
}
// ToGetRequestInformation invoke function getCompatibleHubContentTypes
func (m *ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder) {
    return NewItemSitesItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
