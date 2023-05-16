package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder provides operations to call the addActivities method.
type ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderInternal instantiates a new MicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder and sets the default values.
func NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) {
    m := &ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/connections/{externalConnection%2Did}/items/{externalItem%2Did}/microsoft.graph.externalConnectors.addActivities", pathParameters),
    }
    return m
}
// NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder instantiates a new MicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder and sets the default values.
func NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addActivities
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) Post(ctx context.Context, body ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyable, requestConfiguration *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration)(ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable), nil
}
// ToPostRequestInformation invoke action addActivities
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyable, requestConfiguration *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
