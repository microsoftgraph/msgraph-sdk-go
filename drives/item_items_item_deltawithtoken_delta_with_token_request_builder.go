package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder provides operations to call the delta method.
type ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters invoke function delta
type ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters struct {
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
// ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters
}
// NewItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderInternal instantiates a new ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, token *string)(*ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) {
    m := &ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/delta(token='{token}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if token != nil {
        m.BaseRequestBuilder.PathParameters["token"] = *token
    }
    return m
}
// NewItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder instantiates a new ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function delta
// Deprecated: This method is obsolete. Use GetAsDeltaWithTokenGetResponse instead.
// returns a ItemItemsItemDeltawithtokenDeltaWithTokenResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(ItemItemsItemDeltawithtokenDeltaWithTokenResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemDeltawithtokenDeltaWithTokenResponseable), nil
}
// GetAsDeltaWithTokenGetResponse invoke function delta
// returns a ItemItemsItemDeltawithtokenDeltaWithTokenGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) GetAsDeltaWithTokenGetResponse(ctx context.Context, requestConfiguration *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(ItemItemsItemDeltawithtokenDeltaWithTokenGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemDeltawithtokenDeltaWithTokenGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemDeltawithtokenDeltaWithTokenGetResponseable), nil
}
// ToGetRequestInformation invoke function delta
// returns a *RequestInformation when successful
func (m *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder when successful
func (m *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) {
    return NewItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
