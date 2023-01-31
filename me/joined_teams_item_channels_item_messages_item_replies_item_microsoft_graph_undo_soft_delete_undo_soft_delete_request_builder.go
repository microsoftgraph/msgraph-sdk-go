package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilder provides operations to call the undoSoftDelete method.
type JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewJoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilderInternal instantiates a new UndoSoftDeleteRequestBuilder and sets the default values.
func NewJoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilder) {
    m := &JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/microsoft.graph.undoSoftDelete";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewJoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilder instantiates a new UndoSoftDeleteRequestBuilder and sets the default values.
func NewJoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action undoSoftDelete
func (m *JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action undoSoftDelete
func (m *JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphUndoSoftDeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
