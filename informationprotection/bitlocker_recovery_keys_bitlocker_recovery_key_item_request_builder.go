package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
type BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a bitlockerRecoveryKey object.  By default, this operation does not return the **key** property that represents the actual recovery key. To include the **key** property in the response, use the `$select` OData query parameter. Including the `$select` query parameter triggers an Azure AD audit of the operation and generates an audit log. For more information on audit logs for bitlocker recovery keys, see the KeyManagement category filter of Azure AD audit logs.
type BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetQueryParameters
}
// NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderInternal instantiates a new BitlockerRecoveryKeyItemRequestBuilder and sets the default values.
func NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) {
    m := &BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/bitlocker/recoveryKeys/{bitlockerRecoveryKey%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder instantiates a new BitlockerRecoveryKeyItemRequestBuilder and sets the default values.
func NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve the properties and relationships of a bitlockerRecoveryKey object.  By default, this operation does not return the **key** property that represents the actual recovery key. To include the **key** property in the response, use the `$select` OData query parameter. Including the `$select` query parameter triggers an Azure AD audit of the operation and generates an audit log. For more information on audit logs for bitlocker recovery keys, see the KeyManagement category filter of Azure AD audit logs.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/bitlockerrecoverykey-get?view=graph-rest-1.0
func (m *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BitlockerRecoveryKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation retrieve the properties and relationships of a bitlockerRecoveryKey object.  By default, this operation does not return the **key** property that represents the actual recovery key. To include the **key** property in the response, use the `$select` OData query parameter. Including the `$select` query parameter triggers an Azure AD audit of the operation and generates an audit log. For more information on audit logs for bitlocker recovery keys, see the KeyManagement category filter of Azure AD audit logs.
func (m *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
