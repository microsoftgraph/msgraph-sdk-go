package resetunseencount

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ResetUnseenCountRequestBuilder provides operations to call the resetUnseenCount method.
type ResetUnseenCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ResetUnseenCountRequestBuilderPostOptions options for Post
type ResetUnseenCountRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewResetUnseenCountRequestBuilderInternal instantiates a new ResetUnseenCountRequestBuilder and sets the default values.
func NewResetUnseenCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResetUnseenCountRequestBuilder) {
    m := &ResetUnseenCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/microsoft.graph.resetUnseenCount";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResetUnseenCountRequestBuilder instantiates a new ResetUnseenCountRequestBuilder and sets the default values.
func NewResetUnseenCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResetUnseenCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResetUnseenCountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action resetUnseenCount
func (m *ResetUnseenCountRequestBuilder) CreatePostRequestInformation(options *ResetUnseenCountRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Post invoke action resetUnseenCount
func (m *ResetUnseenCountRequestBuilder) Post(options *ResetUnseenCountRequestBuilderPostOptions)(error) {
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
