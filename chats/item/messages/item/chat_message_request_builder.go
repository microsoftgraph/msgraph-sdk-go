package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4dd20ef2f03eb3ac5a88d1d335d2207bad29d018608ebdcff2629ce7beb73cba "github.com/microsoftgraph/msgraph-sdk-go/chats/item/messages/item/replies"
    i8872343eb24a7e3df178dcacc3a076ed4cd290431032b15da06fbae884e88eeb "github.com/microsoftgraph/msgraph-sdk-go/chats/item/messages/item/hostedcontents"
    i5a9c11331b15ee047ffec77dd1c0041696766415bb61db786849e056a16de016 "github.com/microsoftgraph/msgraph-sdk-go/chats/item/messages/item/hostedcontents/item"
    i87168d76d278ba19d5bf4864ada7dbaff8c9afe027072ddeb1dec7c4368fda6f "github.com/microsoftgraph/msgraph-sdk-go/chats/item/messages/item/replies/item"
)

// Builds and executes requests for operations under \chats\{chat-id}\messages\{chatMessage-id}
type ChatMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ChatMessageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ChatMessageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChatMessageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// A collection of all the messages in the chat. Nullable.
type ChatMessageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ChatMessageRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ChatMessage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ChatMessageRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewChatMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatMessageRequestBuilder) {
    m := &ChatMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/chats/{chat_id}/messages/{chatMessage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ChatMessageRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewChatMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatMessageRequestBuilderInternal(urlParams, requestAdapter)
}
// A collection of all the messages in the chat. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ChatMessageRequestBuilder) CreateDeleteRequestInformation(options *ChatMessageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// A collection of all the messages in the chat. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ChatMessageRequestBuilder) CreateGetRequestInformation(options *ChatMessageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// A collection of all the messages in the chat. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ChatMessageRequestBuilder) CreatePatchRequestInformation(options *ChatMessageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// A collection of all the messages in the chat. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ChatMessageRequestBuilder) Delete(options *ChatMessageRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// A collection of all the messages in the chat. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ChatMessageRequestBuilder) Get(options *ChatMessageRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ChatMessage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewChatMessage() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ChatMessage), nil
}
func (m *ChatMessageRequestBuilder) HostedContents()(*i8872343eb24a7e3df178dcacc3a076ed4cd290431032b15da06fbae884e88eeb.HostedContentsRequestBuilder) {
    return i8872343eb24a7e3df178dcacc3a076ed4cd290431032b15da06fbae884e88eeb.NewHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.messages.item.hostedContents.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ChatMessageRequestBuilder) HostedContentsById(id string)(*i5a9c11331b15ee047ffec77dd1c0041696766415bb61db786849e056a16de016.ChatMessageHostedContentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent_id"] = id
    }
    return i5a9c11331b15ee047ffec77dd1c0041696766415bb61db786849e056a16de016.NewChatMessageHostedContentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// A collection of all the messages in the chat. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ChatMessageRequestBuilder) Patch(options *ChatMessageRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChatMessageRequestBuilder) Replies()(*i4dd20ef2f03eb3ac5a88d1d335d2207bad29d018608ebdcff2629ce7beb73cba.RepliesRequestBuilder) {
    return i4dd20ef2f03eb3ac5a88d1d335d2207bad29d018608ebdcff2629ce7beb73cba.NewRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.messages.item.replies.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ChatMessageRequestBuilder) RepliesById(id string)(*i87168d76d278ba19d5bf4864ada7dbaff8c9afe027072ddeb1dec7c4368fda6f.ChatMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id1"] = id
    }
    return i87168d76d278ba19d5bf4864ada7dbaff8c9afe027072ddeb1dec7c4368fda6f.NewChatMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
