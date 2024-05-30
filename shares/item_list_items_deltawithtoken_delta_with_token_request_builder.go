package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder provides operations to call the delta method.
type ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters invoke function delta
type ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters struct {
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
// ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters
}
// NewItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderInternal instantiates a new ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, token *string)(*ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder) {
    m := &ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shares/{sharedDriveItem%2Did}/list/items/delta(token='{token}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if token != nil {
        m.BaseRequestBuilder.PathParameters["token"] = *token
    }
    return m
}
// NewItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder instantiates a new ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function delta
// Deprecated: This method is obsolete. Use GetAsDeltaWithTokenGetResponse instead.
// returns a ItemListItemsDeltawithtokenDeltaWithTokenResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(ItemListItemsDeltawithtokenDeltaWithTokenResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListItemsDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListItemsDeltawithtokenDeltaWithTokenResponseable), nil
}
// GetAsDeltaWithTokenGetResponse invoke function delta
// returns a ItemListItemsDeltawithtokenDeltaWithTokenGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder) GetAsDeltaWithTokenGetResponse(ctx context.Context, requestConfiguration *ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(ItemListItemsDeltawithtokenDeltaWithTokenGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListItemsDeltawithtokenDeltaWithTokenGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListItemsDeltawithtokenDeltaWithTokenGetResponseable), nil
}
// ToGetRequestInformation invoke function delta
// returns a *RequestInformation when successful
func (m *ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder when successful
func (m *ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder) WithUrl(rawUrl string)(*ItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder) {
    return NewItemListItemsDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
