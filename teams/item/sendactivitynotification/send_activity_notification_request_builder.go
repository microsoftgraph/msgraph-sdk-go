package sendactivitynotification

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SendActivityNotificationRequestBuilder provides operations to call the sendActivityNotification method.
type SendActivityNotificationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SendActivityNotificationRequestBuilderPostOptions options for Post
type SendActivityNotificationRequestBuilderPostOptions struct {
    // 
    Body SendActivityNotificationRequestBodyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewSendActivityNotificationRequestBuilderInternal instantiates a new SendActivityNotificationRequestBuilder and sets the default values.
func NewSendActivityNotificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SendActivityNotificationRequestBuilder) {
    m := &SendActivityNotificationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}/microsoft.graph.sendActivityNotification";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSendActivityNotificationRequestBuilder instantiates a new SendActivityNotificationRequestBuilder and sets the default values.
func NewSendActivityNotificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SendActivityNotificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSendActivityNotificationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action sendActivityNotification
func (m *SendActivityNotificationRequestBuilder) CreatePostRequestInformation(options *SendActivityNotificationRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action sendActivityNotification
func (m *SendActivityNotificationRequestBuilder) Post(options *SendActivityNotificationRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
