package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilder provides operations to call the sync method.
type ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilderInternal instantiates a new SyncRequestBuilder and sets the default values.
func NewExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilder) {
    m := &ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/exchangeConnectors/{deviceManagementExchangeConnector%2Did}/microsoft.graph.sync";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilder instantiates a new SyncRequestBuilder and sets the default values.
func NewExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action sync
func (m *ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilder) Post(ctx context.Context, body ExchangeConnectorsItemMicrosoftGraphSyncSyncPostRequestBodyable, requestConfiguration *ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action sync
func (m *ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilder) ToPostRequestInformation(ctx context.Context, body ExchangeConnectorsItemMicrosoftGraphSyncSyncPostRequestBodyable, requestConfiguration *ExchangeConnectorsItemMicrosoftGraphSyncSyncRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
