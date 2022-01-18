package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i37c730f3a776c4861a3ee114fea1786dc2f8694418aa9a11053db984285dbb1c "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/item/attachments"
    ibfcfb3a0d8e4cf027a88191b2472bea77c62dbd34325a76ccc824f3b2f88318f "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/item/attachmentsarchive"
    id240fb98c68f157c887c62d55372882e4df327f7de3afa3a7c9d7b6ba3baefa6 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/item/attachments/item"
)

// ServiceUpdateMessageRequestBuilder builds and executes requests for operations under \admin\serviceAnnouncement\messages\{serviceUpdateMessage-id}
type ServiceUpdateMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ServiceUpdateMessageRequestBuilderDeleteOptions options for Delete
type ServiceUpdateMessageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServiceUpdateMessageRequestBuilderGetOptions options for Get
type ServiceUpdateMessageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServiceUpdateMessageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServiceUpdateMessageRequestBuilderGetQueryParameters a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
type ServiceUpdateMessageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ServiceUpdateMessageRequestBuilderPatchOptions options for Patch
type ServiceUpdateMessageRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceUpdateMessage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ServiceUpdateMessageRequestBuilder) Attachments()(*i37c730f3a776c4861a3ee114fea1786dc2f8694418aa9a11053db984285dbb1c.AttachmentsRequestBuilder) {
    return i37c730f3a776c4861a3ee114fea1786dc2f8694418aa9a11053db984285dbb1c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServiceUpdateMessageRequestBuilder) AttachmentsArchive()(*ibfcfb3a0d8e4cf027a88191b2472bea77c62dbd34325a76ccc824f3b2f88318f.AttachmentsArchiveRequestBuilder) {
    return ibfcfb3a0d8e4cf027a88191b2472bea77c62dbd34325a76ccc824f3b2f88318f.NewAttachmentsArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.messages.item.attachments.item collection
func (m *ServiceUpdateMessageRequestBuilder) AttachmentsById(id string)(*id240fb98c68f157c887c62d55372882e4df327f7de3afa3a7c9d7b6ba3baefa6.ServiceAnnouncementAttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceAnnouncementAttachment_id"] = id
    }
    return id240fb98c68f157c887c62d55372882e4df327f7de3afa3a7c9d7b6ba3baefa6.NewServiceAnnouncementAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewServiceUpdateMessageRequestBuilderInternal instantiates a new ServiceUpdateMessageRequestBuilder and sets the default values.
func NewServiceUpdateMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceUpdateMessageRequestBuilder) {
    m := &ServiceUpdateMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages/{serviceUpdateMessage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServiceUpdateMessageRequestBuilder instantiates a new ServiceUpdateMessageRequestBuilder and sets the default values.
func NewServiceUpdateMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceUpdateMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceUpdateMessageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) CreateDeleteRequestInformation(options *ServiceUpdateMessageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) CreateGetRequestInformation(options *ServiceUpdateMessageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) CreatePatchRequestInformation(options *ServiceUpdateMessageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) Delete(options *ServiceUpdateMessageRequestBuilderDeleteOptions)(error) {
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
// Get a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) Get(options *ServiceUpdateMessageRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceUpdateMessage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewServiceUpdateMessage() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceUpdateMessage), nil
}
// Patch a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) Patch(options *ServiceUpdateMessageRequestBuilderPatchOptions)(error) {
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
