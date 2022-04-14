package revokesigninsessions

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RevokeSignInSessionsRequestBuilder provides operations to call the revokeSignInSessions method.
type RevokeSignInSessionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RevokeSignInSessionsRequestBuilderPostOptions options for Post
type RevokeSignInSessionsRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewRevokeSignInSessionsRequestBuilderInternal instantiates a new RevokeSignInSessionsRequestBuilder and sets the default values.
func NewRevokeSignInSessionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RevokeSignInSessionsRequestBuilder) {
    m := &RevokeSignInSessionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.revokeSignInSessions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRevokeSignInSessionsRequestBuilder instantiates a new RevokeSignInSessionsRequestBuilder and sets the default values.
func NewRevokeSignInSessionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RevokeSignInSessionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRevokeSignInSessionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action revokeSignInSessions
func (m *RevokeSignInSessionsRequestBuilder) CreatePostRequestInformation(options *RevokeSignInSessionsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action revokeSignInSessions
func (m *RevokeSignInSessionsRequestBuilder) Post(options *RevokeSignInSessionsRequestBuilderPostOptions)(RevokeSignInSessionsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateRevokeSignInSessionsResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(RevokeSignInSessionsResponseable), nil
}
