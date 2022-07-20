package chats

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i58eca7b0448a35d407d9e18a625f546b474524f2b0edbbbe8c0846aac6804be8 "github.com/microsoftgraph/msgraph-sdk-go/me/chats/getallmessages"
    i6fcfe00f427d5955614f1a04fce3791bede0bc3325d70d57c29afeaef9eab1df "github.com/microsoftgraph/msgraph-sdk-go/me/chats/count"
)

// ChatsRequestBuilder provides operations to manage the chats property of the microsoft.graph.user entity.
type ChatsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatsRequestBuilderGetQueryParameters get chats from me
type ChatsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ChatsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatsRequestBuilderGetQueryParameters
}
// ChatsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatsRequestBuilderInternal instantiates a new ChatsRequestBuilder and sets the default values.
func NewChatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsRequestBuilder) {
    m := &ChatsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatsRequestBuilder instantiates a new ChatsRequestBuilder and sets the default values.
func NewChatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ChatsRequestBuilder) Count()(*i6fcfe00f427d5955614f1a04fce3791bede0bc3325d70d57c29afeaef9eab1df.CountRequestBuilder) {
    return i6fcfe00f427d5955614f1a04fce3791bede0bc3325d70d57c29afeaef9eab1df.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get chats from me
func (m *ChatsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get chats from me
func (m *ChatsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ChatsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to chats for me
func (m *ChatsRequestBuilder) CreatePostRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to chats for me
func (m *ChatsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ChatsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get chats from me
func (m *ChatsRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetAllMessages provides operations to call the getAllMessages method.
func (m *ChatsRequestBuilder) GetAllMessages()(*i58eca7b0448a35d407d9e18a625f546b474524f2b0edbbbe8c0846aac6804be8.GetAllMessagesRequestBuilder) {
    return i58eca7b0448a35d407d9e18a625f546b474524f2b0edbbbe8c0846aac6804be8.NewGetAllMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get chats from me
func (m *ChatsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ChatsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatCollectionResponseable), nil
}
// Post create new navigation property to chats for me
func (m *ChatsRequestBuilder) Post(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to chats for me
func (m *ChatsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ChatsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable), nil
}
