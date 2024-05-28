package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder provides operations to call the createReply method.
type ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilderInternal instantiates a new ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder and sets the default values.
func NewItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder) {
    m := &ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/createReply", pathParameters),
    }
    return m
}
// NewItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder instantiates a new ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder and sets the default values.
func NewItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a draft to reply to the sender of a message in either JSON or MIME format. When using JSON format:- Specify either a comment or the body property of the message parameter. Specifying both will return an HTTP 400 Bad Request error.- If replyTo is specified in the original message, per Internet Message Format (RFC 2822), you should send the reply to the recipients in replyTo, and not the recipients in from.- You can update the draft later to add reply content to the body or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in base64 format in the request body.- Add any attachments and S/MIME properties to the MIME content. Send the draft message in a subsequent operation. Alternatively, reply to a message in a single operation.
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/message-createreply?view=graph-rest-1.0
func (m *ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder) Post(ctx context.Context, body ItemMailfoldersItemMessagesItemCreatereplyCreateReplyPostRequestBodyable, requestConfiguration *ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
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
// ToPostRequestInformation create a draft to reply to the sender of a message in either JSON or MIME format. When using JSON format:- Specify either a comment or the body property of the message parameter. Specifying both will return an HTTP 400 Bad Request error.- If replyTo is specified in the original message, per Internet Message Format (RFC 2822), you should send the reply to the recipients in replyTo, and not the recipients in from.- You can update the draft later to add reply content to the body or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in base64 format in the request body.- Add any attachments and S/MIME properties to the MIME content. Send the draft message in a subsequent operation. Alternatively, reply to a message in a single operation.
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailfoldersItemMessagesItemCreatereplyCreateReplyPostRequestBodyable, requestConfiguration *ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder when successful
func (m *ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder) WithUrl(rawUrl string)(*ItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder) {
    return NewItemMailfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
