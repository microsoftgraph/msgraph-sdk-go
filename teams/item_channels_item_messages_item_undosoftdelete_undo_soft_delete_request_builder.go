package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder provides operations to call the undoSoftDelete method.
type ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderInternal instantiates a new ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) {
    m := &ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/undoSoftDelete", pathParameters),
    }
    return m
}
// NewItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder instantiates a new ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post undo soft deletion of a single chatMessage or a chat message reply in a channel or a chat.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-undosoftdelete?view=graph-rest-1.0
func (m *ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation undo soft deletion of a single chatMessage or a chat message reply in a channel or a chat.
// returns a *RequestInformation when successful
func (m *ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder when successful
func (m *ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) WithUrl(rawUrl string)(*ItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) {
    return NewItemChannelsItemMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
