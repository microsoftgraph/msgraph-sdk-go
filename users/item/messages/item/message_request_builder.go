package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i03bf4b34230f15d3e490ba46ab92e65b3c86befa789b18b8d0a25dec32b8ec13 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/singlevalueextendedproperties"
    i122ff3868c47eff36e63f27e4a5b59d68ad9c0ac6f599d0365b5b2a23664fef8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/attachments"
    i21a12f3061b5040c42eeb7a6e4e887142ebaa674b6c8e349ad4d27e4203799f5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/value"
    i21ac64590aeb4ea4de81e9f5c0bdb9af91b5615ed8258947f1983a17be31a12a "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/createreply"
    i288e98553df56b30f05fcd3093c61b8bf5395192fb8c9a2825f62bc5afe55a0e "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/calendarsharingmessage"
    i29ffc4815578e16ff0a35471a241efa71e25f9fce417e2ea2e64b556ae6b4084 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/copy"
    i4382614991ffa5cffd77c6a25c626b3b88a33b415d02b39342eb5f3a52cb881e "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/multivalueextendedproperties"
    i505c2e3f7a924493aa1390d23de4d48964764a251a918b6c13839c1b7c1aaf20 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/replyall"
    i77a08d89129ea8ff7ca1c841cd1e162b665795b77f30f02ee19f6b11a86fbb8e "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/forward"
    i8ad53531e48d207d730614399c20bde311ba08c3dd03b37e69bef6ba38fd79fa "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/reply"
    i8ba7bb1c08d896d87831ebedb8fe7a65e343758cbd43c2a2dcaea7a802dad0b3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/send"
    i9a54a4fefd859216588c818688d9056c0d0261d20a194bfa8cd910fe376614ea "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/extensions"
    iabf960d8fba3af22e12d1fdca449478d5dda049dece4b4177254bff526c014ae "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/createforward"
    id70c200088dfe1c9223f533b4a73c07c53ce049857b0be9d382686560b3a0a44 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/createreplyall"
    iee8102d3efbbfcfdd28ff8f8427c68562589a436cd6306dcc0ffdef0f46db71d "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/move"
    i3126e60880d45aa8eb3e50dbb7349d260566b68be6e5204b3bd5ce8027d87ae4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/multivalueextendedproperties/item"
    i4f930801d57aea145b04d038b3787d219e2ba9c79e633cdc6366b0f192b7751d "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/extensions/item"
    i569e8a41910c91a566eb55ee03be3b95dad61f13affae3df81e08c5093475576 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/attachments/item"
    iab94264dd4909f265fdecd3eca0cdeb9268a7eec19a0c90bc22d0c9cb1ce414b "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/singlevalueextendedproperties/item"
)

// Builds and executes requests for operations under \users\{user-id}\messages\{message-id}
type MessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The messages in a mailbox or folder. Read-only. Nullable.
type MessageRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
func (m *MessageRequestBuilder) Attachments()(*i122ff3868c47eff36e63f27e4a5b59d68ad9c0ac6f599d0365b5b2a23664fef8.AttachmentsRequestBuilder) {
    return i122ff3868c47eff36e63f27e4a5b59d68ad9c0ac6f599d0365b5b2a23664fef8.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.messages.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *MessageRequestBuilder) AttachmentsById(id string)(*i569e8a41910c91a566eb55ee03be3b95dad61f13affae3df81e08c5093475576.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i569e8a41910c91a566eb55ee03be3b95dad61f13affae3df81e08c5093475576.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) CalendarSharingMessage()(*i288e98553df56b30f05fcd3093c61b8bf5395192fb8c9a2825f62bc5afe55a0e.CalendarSharingMessageRequestBuilder) {
    return i288e98553df56b30f05fcd3093c61b8bf5395192fb8c9a2825f62bc5afe55a0e.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new MessageRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/messages/{message_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new MessageRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MessageRequestBuilder) Content()(*i21a12f3061b5040c42eeb7a6e4e887142ebaa674b6c8e349ad4d27e4203799f5.ContentRequestBuilder) {
    return i21a12f3061b5040c42eeb7a6e4e887142ebaa674b6c8e349ad4d27e4203799f5.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Copy()(*i29ffc4815578e16ff0a35471a241efa71e25f9fce417e2ea2e64b556ae6b4084.CopyRequestBuilder) {
    return i29ffc4815578e16ff0a35471a241efa71e25f9fce417e2ea2e64b556ae6b4084.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The messages in a mailbox or folder. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *MessageRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MessageRequestBuilder) CreateForward()(*iabf960d8fba3af22e12d1fdca449478d5dda049dece4b4177254bff526c014ae.CreateForwardRequestBuilder) {
    return iabf960d8fba3af22e12d1fdca449478d5dda049dece4b4177254bff526c014ae.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The messages in a mailbox or folder. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *MessageRequestBuilder) CreateGetRequestInformation(q func (value *MessageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(MessageRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The messages in a mailbox or folder. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *MessageRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Message, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MessageRequestBuilder) CreateReply()(*i21ac64590aeb4ea4de81e9f5c0bdb9af91b5615ed8258947f1983a17be31a12a.CreateReplyRequestBuilder) {
    return i21ac64590aeb4ea4de81e9f5c0bdb9af91b5615ed8258947f1983a17be31a12a.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) CreateReplyAll()(*id70c200088dfe1c9223f533b4a73c07c53ce049857b0be9d382686560b3a0a44.CreateReplyAllRequestBuilder) {
    return id70c200088dfe1c9223f533b4a73c07c53ce049857b0be9d382686560b3a0a44.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The messages in a mailbox or folder. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *MessageRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageRequestBuilder) Extensions()(*i9a54a4fefd859216588c818688d9056c0d0261d20a194bfa8cd910fe376614ea.ExtensionsRequestBuilder) {
    return i9a54a4fefd859216588c818688d9056c0d0261d20a194bfa8cd910fe376614ea.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.messages.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *MessageRequestBuilder) ExtensionsById(id string)(*i4f930801d57aea145b04d038b3787d219e2ba9c79e633cdc6366b0f192b7751d.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i4f930801d57aea145b04d038b3787d219e2ba9c79e633cdc6366b0f192b7751d.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Forward()(*i77a08d89129ea8ff7ca1c841cd1e162b665795b77f30f02ee19f6b11a86fbb8e.ForwardRequestBuilder) {
    return i77a08d89129ea8ff7ca1c841cd1e162b665795b77f30f02ee19f6b11a86fbb8e.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The messages in a mailbox or folder. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *MessageRequestBuilder) Get(q func (value *MessageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Message, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMessage() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Message), nil
}
func (m *MessageRequestBuilder) Move()(*iee8102d3efbbfcfdd28ff8f8427c68562589a436cd6306dcc0ffdef0f46db71d.MoveRequestBuilder) {
    return iee8102d3efbbfcfdd28ff8f8427c68562589a436cd6306dcc0ffdef0f46db71d.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*i4382614991ffa5cffd77c6a25c626b3b88a33b415d02b39342eb5f3a52cb881e.MultiValueExtendedPropertiesRequestBuilder) {
    return i4382614991ffa5cffd77c6a25c626b3b88a33b415d02b39342eb5f3a52cb881e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.messages.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i3126e60880d45aa8eb3e50dbb7349d260566b68be6e5204b3bd5ce8027d87ae4.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i3126e60880d45aa8eb3e50dbb7349d260566b68be6e5204b3bd5ce8027d87ae4.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The messages in a mailbox or folder. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *MessageRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Message, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageRequestBuilder) Reply()(*i8ad53531e48d207d730614399c20bde311ba08c3dd03b37e69bef6ba38fd79fa.ReplyRequestBuilder) {
    return i8ad53531e48d207d730614399c20bde311ba08c3dd03b37e69bef6ba38fd79fa.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) ReplyAll()(*i505c2e3f7a924493aa1390d23de4d48964764a251a918b6c13839c1b7c1aaf20.ReplyAllRequestBuilder) {
    return i505c2e3f7a924493aa1390d23de4d48964764a251a918b6c13839c1b7c1aaf20.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Send()(*i8ba7bb1c08d896d87831ebedb8fe7a65e343758cbd43c2a2dcaea7a802dad0b3.SendRequestBuilder) {
    return i8ba7bb1c08d896d87831ebedb8fe7a65e343758cbd43c2a2dcaea7a802dad0b3.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*i03bf4b34230f15d3e490ba46ab92e65b3c86befa789b18b8d0a25dec32b8ec13.SingleValueExtendedPropertiesRequestBuilder) {
    return i03bf4b34230f15d3e490ba46ab92e65b3c86befa789b18b8d0a25dec32b8ec13.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.messages.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iab94264dd4909f265fdecd3eca0cdeb9268a7eec19a0c90bc22d0c9cb1ce414b.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return iab94264dd4909f265fdecd3eca0cdeb9268a7eec19a0c90bc22d0c9cb1ce414b.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
