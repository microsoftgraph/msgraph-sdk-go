package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemInstalledappsItemUpgradeRequestBuilder provides operations to call the upgrade method.
type ItemInstalledappsItemUpgradeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInstalledappsItemUpgradeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstalledappsItemUpgradeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInstalledappsItemUpgradeRequestBuilderInternal instantiates a new ItemInstalledappsItemUpgradeRequestBuilder and sets the default values.
func NewItemInstalledappsItemUpgradeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstalledappsItemUpgradeRequestBuilder) {
    m := &ItemInstalledappsItemUpgradeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/installedApps/{teamsAppInstallation%2Did}/upgrade", pathParameters),
    }
    return m
}
// NewItemInstalledappsItemUpgradeRequestBuilder instantiates a new ItemInstalledappsItemUpgradeRequestBuilder and sets the default values.
func NewItemInstalledappsItemUpgradeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstalledappsItemUpgradeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstalledappsItemUpgradeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upgrade an app installation within a chat.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chat-teamsappinstallation-upgrade?view=graph-rest-1.0
func (m *ItemInstalledappsItemUpgradeRequestBuilder) Post(ctx context.Context, body ItemInstalledappsItemUpgradePostRequestBodyable, requestConfiguration *ItemInstalledappsItemUpgradeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation upgrade an app installation within a chat.
// returns a *RequestInformation when successful
func (m *ItemInstalledappsItemUpgradeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemInstalledappsItemUpgradePostRequestBodyable, requestConfiguration *ItemInstalledappsItemUpgradeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInstalledappsItemUpgradeRequestBuilder when successful
func (m *ItemInstalledappsItemUpgradeRequestBuilder) WithUrl(rawUrl string)(*ItemInstalledappsItemUpgradeRequestBuilder) {
    return NewItemInstalledappsItemUpgradeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
