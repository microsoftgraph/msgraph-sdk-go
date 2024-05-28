package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BitlockerRecoverykeysRecoveryKeysRequestBuilder provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
type BitlockerRecoverykeysRecoveryKeysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BitlockerRecoverykeysRecoveryKeysRequestBuilderGetQueryParameters get a list of the bitlockerRecoveryKey objects and their properties.  This operation does not return the key property. For information about how to read the key property, see Get bitlockerRecoveryKey.
type BitlockerRecoverykeysRecoveryKeysRequestBuilderGetQueryParameters struct {
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
// BitlockerRecoverykeysRecoveryKeysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BitlockerRecoverykeysRecoveryKeysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BitlockerRecoverykeysRecoveryKeysRequestBuilderGetQueryParameters
}
// ByBitlockerRecoveryKeyId provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
// returns a *BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder when successful
func (m *BitlockerRecoverykeysRecoveryKeysRequestBuilder) ByBitlockerRecoveryKeyId(bitlockerRecoveryKeyId string)(*BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if bitlockerRecoveryKeyId != "" {
        urlTplParams["bitlockerRecoveryKey%2Did"] = bitlockerRecoveryKeyId
    }
    return NewBitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBitlockerRecoverykeysRecoveryKeysRequestBuilderInternal instantiates a new BitlockerRecoverykeysRecoveryKeysRequestBuilder and sets the default values.
func NewBitlockerRecoverykeysRecoveryKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoverykeysRecoveryKeysRequestBuilder) {
    m := &BitlockerRecoverykeysRecoveryKeysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/bitlocker/recoveryKeys{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBitlockerRecoverykeysRecoveryKeysRequestBuilder instantiates a new BitlockerRecoverykeysRecoveryKeysRequestBuilder and sets the default values.
func NewBitlockerRecoverykeysRecoveryKeysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoverykeysRecoveryKeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBitlockerRecoverykeysRecoveryKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BitlockerRecoverykeysCountRequestBuilder when successful
func (m *BitlockerRecoverykeysRecoveryKeysRequestBuilder) Count()(*BitlockerRecoverykeysCountRequestBuilder) {
    return NewBitlockerRecoverykeysCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the bitlockerRecoveryKey objects and their properties.  This operation does not return the key property. For information about how to read the key property, see Get bitlockerRecoveryKey.
// returns a BitlockerRecoveryKeyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bitlocker-list-recoverykeys?view=graph-rest-1.0
func (m *BitlockerRecoverykeysRecoveryKeysRequestBuilder) Get(ctx context.Context, requestConfiguration *BitlockerRecoverykeysRecoveryKeysRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BitlockerRecoveryKeyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBitlockerRecoveryKeyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BitlockerRecoveryKeyCollectionResponseable), nil
}
// ToGetRequestInformation get a list of the bitlockerRecoveryKey objects and their properties.  This operation does not return the key property. For information about how to read the key property, see Get bitlockerRecoveryKey.
// returns a *RequestInformation when successful
func (m *BitlockerRecoverykeysRecoveryKeysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BitlockerRecoverykeysRecoveryKeysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BitlockerRecoverykeysRecoveryKeysRequestBuilder when successful
func (m *BitlockerRecoverykeysRecoveryKeysRequestBuilder) WithUrl(rawUrl string)(*BitlockerRecoverykeysRecoveryKeysRequestBuilder) {
    return NewBitlockerRecoverykeysRecoveryKeysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
