package getnotebookfromweburl

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// GetNotebookFromWebUrlRequestBuilder provides operations to call the getNotebookFromWebUrl method.
type GetNotebookFromWebUrlRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetNotebookFromWebUrlRequestBuilderPostOptions options for Post
type GetNotebookFromWebUrlRequestBuilderPostOptions struct {
    // 
    Body GetNotebookFromWebUrlRequestBodyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetNotebookFromWebUrlRequestBuilderInternal instantiates a new GetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewGetNotebookFromWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetNotebookFromWebUrlRequestBuilder) {
    m := &GetNotebookFromWebUrlRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onenote/notebooks/microsoft.graph.getNotebookFromWebUrl";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetNotebookFromWebUrlRequestBuilder instantiates a new GetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewGetNotebookFromWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetNotebookFromWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetNotebookFromWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getNotebookFromWebUrl
func (m *GetNotebookFromWebUrlRequestBuilder) CreatePostRequestInformation(options *GetNotebookFromWebUrlRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getNotebookFromWebUrl
func (m *GetNotebookFromWebUrlRequestBuilder) Post(options *GetNotebookFromWebUrlRequestBuilderPostOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CopyNotebookModelable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCopyNotebookModelFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CopyNotebookModelable), nil
}
