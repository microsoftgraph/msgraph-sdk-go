package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemInstalledappsItemUpgradeRequestBuilder provides operations to call the upgrade method.
type ItemChatsItemInstalledappsItemUpgradeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemInstalledappsItemUpgradeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemInstalledappsItemUpgradeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemInstalledappsItemUpgradeRequestBuilderInternal instantiates a new ItemChatsItemInstalledappsItemUpgradeRequestBuilder and sets the default values.
func NewItemChatsItemInstalledappsItemUpgradeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemInstalledappsItemUpgradeRequestBuilder) {
    m := &ItemChatsItemInstalledappsItemUpgradeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/installedApps/{teamsAppInstallation%2Did}/upgrade", pathParameters),
    }
    return m
}
// NewItemChatsItemInstalledappsItemUpgradeRequestBuilder instantiates a new ItemChatsItemInstalledappsItemUpgradeRequestBuilder and sets the default values.
func NewItemChatsItemInstalledappsItemUpgradeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemInstalledappsItemUpgradeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemInstalledappsItemUpgradeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upgrade an app installation within a chat.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chat-teamsappinstallation-upgrade?view=graph-rest-1.0
func (m *ItemChatsItemInstalledappsItemUpgradeRequestBuilder) Post(ctx context.Context, body ItemChatsItemInstalledappsItemUpgradePostRequestBodyable, requestConfiguration *ItemChatsItemInstalledappsItemUpgradeRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemChatsItemInstalledappsItemUpgradeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemInstalledappsItemUpgradePostRequestBodyable, requestConfiguration *ItemChatsItemInstalledappsItemUpgradeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChatsItemInstalledappsItemUpgradeRequestBuilder when successful
func (m *ItemChatsItemInstalledappsItemUpgradeRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemInstalledappsItemUpgradeRequestBuilder) {
    return NewItemChatsItemInstalledappsItemUpgradeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
