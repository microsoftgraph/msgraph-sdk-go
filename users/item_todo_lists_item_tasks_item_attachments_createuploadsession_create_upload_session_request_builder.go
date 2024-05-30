package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder provides operations to call the createUploadSession method.
type ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal instantiates a new ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    m := &ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/todo/lists/{todoTaskList%2Did}/tasks/{todoTask%2Did}/attachments/createUploadSession", pathParameters),
    }
    return m
}
// NewItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder instantiates a new ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an upload session to iteratively upload ranges of a file as an attachment to a todoTask. As part of the response, this action returns an upload URL that you can use in subsequent sequential PUT queries. The request headers for each PUT operation let you specify the exact range of bytes to be uploaded. This allows the transfer to be resumed, in case the network connection is dropped during the upload. The following are the steps to attach a file to a Microsoft To Do task using an upload session: For an example that describes the end-to-end attachment process, see attach files to a To Do task.
// returns a UploadSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/taskfileattachment-createuploadsession?view=graph-rest-1.0
func (m *ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) Post(ctx context.Context, body ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UploadSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUploadSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UploadSessionable), nil
}
// ToPostRequestInformation create an upload session to iteratively upload ranges of a file as an attachment to a todoTask. As part of the response, this action returns an upload URL that you can use in subsequent sequential PUT queries. The request headers for each PUT operation let you specify the exact range of bytes to be uploaded. This allows the transfer to be resumed, in case the network connection is dropped during the upload. The following are the steps to attach a file to a Microsoft To Do task using an upload session: For an example that describes the end-to-end attachment process, see attach files to a To Do task.
// returns a *RequestInformation when successful
func (m *ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) WithUrl(rawUrl string)(*ItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewItemTodoListsItemTasksItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
