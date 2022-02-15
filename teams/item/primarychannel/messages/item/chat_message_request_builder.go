package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i571bf3023263021405f4655cf8847eced3afd54194b2424b5494f709a2d60d55 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/messages/item/replies"
    id0e73af84ecb90a9c8c83df652b8a59552d9cd05e0f8c4809ff54a88590691de "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/messages/item/hostedcontents"
    ia90c0139ab7dee9aa1759297d60868917e86090a24094d3f9381ad741e147bc1 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/messages/item/hostedcontents/item"
    iac6cc4119abbf87cd78b8b30a2e5d3c0e8f11efb4c6cdffd9c664bb37f24623e "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/messages/item/replies/item"
)

// ChatMessageRequestBuilder builds and executes requests for operations under \teams\{team-id}\primaryChannel\messages\{chatMessage-id}
type ChatMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChatMessageRequestBuilderDeleteOptions options for Delete
type ChatMessageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatMessageRequestBuilderGetOptions options for Get
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
// ChatMessageRequestBuilderGetQueryParameters a collection of all the messages in the channel. A navigation property. Nullable.
type ChatMessageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ChatMessageRequestBuilderPatchOptions options for Patch
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
// NewChatMessageRequestBuilderInternal instantiates a new ChatMessageRequestBuilder and sets the default values.
func NewChatMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatMessageRequestBuilder) {
    m := &ChatMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}/primaryChannel/messages/{chatMessage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatMessageRequestBuilder instantiates a new ChatMessageRequestBuilder and sets the default values.
func NewChatMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatMessageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a collection of all the messages in the channel. A navigation property. Nullable.
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
// CreateGetRequestInformation a collection of all the messages in the channel. A navigation property. Nullable.
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
// CreatePatchRequestInformation a collection of all the messages in the channel. A navigation property. Nullable.
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
// Delete a collection of all the messages in the channel. A navigation property. Nullable.
func (m *ChatMessageRequestBuilder) Delete(options *ChatMessageRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of all the messages in the channel. A navigation property. Nullable.
func (m *ChatMessageRequestBuilder) Get(options *ChatMessageRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ChatMessage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewChatMessage() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ChatMessage), nil
}
func (m *ChatMessageRequestBuilder) HostedContents()(*id0e73af84ecb90a9c8c83df652b8a59552d9cd05e0f8c4809ff54a88590691de.HostedContentsRequestBuilder) {
    return id0e73af84ecb90a9c8c83df652b8a59552d9cd05e0f8c4809ff54a88590691de.NewHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostedContentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.primaryChannel.messages.item.hostedContents.item collection
func (m *ChatMessageRequestBuilder) HostedContentsById(id string)(*ia90c0139ab7dee9aa1759297d60868917e86090a24094d3f9381ad741e147bc1.ChatMessageHostedContentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent_id"] = id
    }
    return ia90c0139ab7dee9aa1759297d60868917e86090a24094d3f9381ad741e147bc1.NewChatMessageHostedContentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch a collection of all the messages in the channel. A navigation property. Nullable.
func (m *ChatMessageRequestBuilder) Patch(options *ChatMessageRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChatMessageRequestBuilder) Replies()(*i571bf3023263021405f4655cf8847eced3afd54194b2424b5494f709a2d60d55.RepliesRequestBuilder) {
    return i571bf3023263021405f4655cf8847eced3afd54194b2424b5494f709a2d60d55.NewRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RepliesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.primaryChannel.messages.item.replies.item collection
func (m *ChatMessageRequestBuilder) RepliesById(id string)(*iac6cc4119abbf87cd78b8b30a2e5d3c0e8f11efb4c6cdffd9c664bb37f24623e.ChatMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id1"] = id
    }
    return iac6cc4119abbf87cd78b8b30a2e5d3c0e8f11efb4c6cdffd9c664bb37f24623e.NewChatMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
