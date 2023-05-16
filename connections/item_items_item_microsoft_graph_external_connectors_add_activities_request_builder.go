package connections

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder provides operations to call the addActivities method.
type ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderInternal instantiates a new MicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder and sets the default values.
func NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) {
    m := &ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/connections/{externalConnection%2Did}/items/{externalItem%2Did}/microsoft.graph.externalConnectors.addActivities", pathParameters),
    }
    return m
}
// NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder instantiates a new MicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder and sets the default values.
func NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addActivities
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) Post(ctx context.Context, body ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyable, requestConfiguration *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration)(ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable), nil
}
// ToPostRequestInformation invoke action addActivities
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyable, requestConfiguration *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
