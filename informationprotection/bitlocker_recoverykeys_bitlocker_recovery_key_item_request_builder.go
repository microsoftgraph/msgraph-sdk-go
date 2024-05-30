package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
type BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a bitlockerRecoveryKey object.  By default, this operation doesn't return the key property that represents the actual recovery key. To include the key property in the response, use the $select OData query parameter. Including the $select query parameter triggers a Microsoft Entra audit of the operation and generates an audit log. For more information on audit logs for bitlocker recovery keys, see the KeyManagement category filter of Microsoft Entra audit logs.
type BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderGetQueryParameters
}
// NewBitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderInternal instantiates a new BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder and sets the default values.
func NewBitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder) {
    m := &BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/bitlocker/recoveryKeys/{bitlockerRecoveryKey%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder instantiates a new BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder and sets the default values.
func NewBitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve the properties and relationships of a bitlockerRecoveryKey object.  By default, this operation doesn't return the key property that represents the actual recovery key. To include the key property in the response, use the $select OData query parameter. Including the $select query parameter triggers a Microsoft Entra audit of the operation and generates an audit log. For more information on audit logs for bitlocker recovery keys, see the KeyManagement category filter of Microsoft Entra audit logs.
// returns a BitlockerRecoveryKeyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bitlockerrecoverykey-get?view=graph-rest-1.0
func (m *BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BitlockerRecoveryKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBitlockerRecoveryKeyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BitlockerRecoveryKeyable), nil
}
// ToGetRequestInformation retrieve the properties and relationships of a bitlockerRecoveryKey object.  By default, this operation doesn't return the key property that represents the actual recovery key. To include the key property in the response, use the $select OData query parameter. Including the $select query parameter triggers a Microsoft Entra audit of the operation and generates an audit log. For more information on audit logs for bitlocker recovery keys, see the KeyManagement category filter of Microsoft Entra audit logs.
// returns a *RequestInformation when successful
func (m *BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder when successful
func (m *BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder) WithUrl(rawUrl string)(*BitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder) {
    return NewBitlockerRecoverykeysBitlockerRecoveryKeyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
