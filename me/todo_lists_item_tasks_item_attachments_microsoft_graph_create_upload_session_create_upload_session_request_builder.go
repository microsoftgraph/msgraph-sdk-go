package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder provides operations to call the createUploadSession method.
type TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderInternal instantiates a new CreateUploadSessionRequestBuilder and sets the default values.
func NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder) {
    m := &TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/todo/lists/{todoTaskList%2Did}/tasks/{todoTask%2Did}/attachments/microsoft.graph.createUploadSession";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder instantiates a new CreateUploadSessionRequestBuilder and sets the default values.
func NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createUploadSession
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder) Post(ctx context.Context, body TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyable, requestConfiguration *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UploadSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUploadSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UploadSessionable), nil
}
// ToPostRequestInformation invoke action createUploadSession
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyable, requestConfiguration *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
