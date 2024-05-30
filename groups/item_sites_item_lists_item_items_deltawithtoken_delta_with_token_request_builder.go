package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder provides operations to call the delta method.
type ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters invoke function delta
type ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
// ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters
}
// NewItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderInternal instantiates a new ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, token *string)(*ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) {
    m := &ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/items/delta(token='{token}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if token != nil {
        m.BaseRequestBuilder.PathParameters["token"] = *token
    }
    return m
}
// NewItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder instantiates a new ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function delta
// Deprecated: This method is obsolete. Use GetAsDeltaWithTokenGetResponse instead.
// returns a ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenResponseable), nil
}
// GetAsDeltaWithTokenGetResponse invoke function delta
// returns a ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) GetAsDeltaWithTokenGetResponse(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseable), nil
}
// ToGetRequestInformation invoke function delta
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder when successful
func (m *ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) {
    return NewItemSitesItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
