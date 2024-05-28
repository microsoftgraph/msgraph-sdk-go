package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder provides operations to call the copy method.
type ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilderInternal instantiates a new ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder) {
    m := &ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/copy", pathParameters),
    }
    return m
}
// NewItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder instantiates a new ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post copy a message to a folder within the user's mailbox.
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/message-copy?view=graph-rest-1.0
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder) Post(ctx context.Context, body ItemMailfoldersItemChildfoldersItemMessagesItemCopyPostRequestBodyable, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable), nil
}
// ToPostRequestInformation copy a message to a folder within the user's mailbox.
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailfoldersItemChildfoldersItemMessagesItemCopyPostRequestBodyable, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder) WithUrl(rawUrl string)(*ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
