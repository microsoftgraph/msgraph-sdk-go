package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedchatsItemUndodeleteUndoDeleteRequestBuilder provides operations to call the undoDelete method.
type DeletedchatsItemUndodeleteUndoDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedchatsItemUndodeleteUndoDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedchatsItemUndodeleteUndoDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedchatsItemUndodeleteUndoDeleteRequestBuilderInternal instantiates a new DeletedchatsItemUndodeleteUndoDeleteRequestBuilder and sets the default values.
func NewDeletedchatsItemUndodeleteUndoDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedchatsItemUndodeleteUndoDeleteRequestBuilder) {
    m := &DeletedchatsItemUndodeleteUndoDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedChats/{deletedChat%2Did}/undoDelete", pathParameters),
    }
    return m
}
// NewDeletedchatsItemUndodeleteUndoDeleteRequestBuilder instantiates a new DeletedchatsItemUndodeleteUndoDeleteRequestBuilder and sets the default values.
func NewDeletedchatsItemUndodeleteUndoDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedchatsItemUndodeleteUndoDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedchatsItemUndodeleteUndoDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a  deletedChat to an active chat.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/deletedchat-undodelete?view=graph-rest-1.0
func (m *DeletedchatsItemUndodeleteUndoDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *DeletedchatsItemUndodeleteUndoDeleteRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation restore a  deletedChat to an active chat.
// returns a *RequestInformation when successful
func (m *DeletedchatsItemUndodeleteUndoDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeletedchatsItemUndodeleteUndoDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeletedchatsItemUndodeleteUndoDeleteRequestBuilder when successful
func (m *DeletedchatsItemUndodeleteUndoDeleteRequestBuilder) WithUrl(rawUrl string)(*DeletedchatsItemUndodeleteUndoDeleteRequestBuilder) {
    return NewDeletedchatsItemUndodeleteUndoDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
