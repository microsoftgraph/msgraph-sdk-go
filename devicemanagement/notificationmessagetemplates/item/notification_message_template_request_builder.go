package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    icc3fe1a1c8b964bc176be6a6d48621fc3548efaeadcb3b10e136af2f19764752 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/notificationmessagetemplates/item/localizednotificationmessages"
    if10e6b46a7479cb721a918b2370be92792065fdd20e80c33685b2f50841a32b9 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/notificationmessagetemplates/item/sendtestmessage"
    i079873d8184290e06e63c9a04520d4cbb886c99c078b4ea0a493ba22bc14901e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/notificationmessagetemplates/item/localizednotificationmessages/item"
)

// NotificationMessageTemplateRequestBuilder builds and executes requests for operations under \deviceManagement\notificationMessageTemplates\{notificationMessageTemplate-id}
type NotificationMessageTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NotificationMessageTemplateRequestBuilderDeleteOptions options for Delete
type NotificationMessageTemplateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotificationMessageTemplateRequestBuilderGetOptions options for Get
type NotificationMessageTemplateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NotificationMessageTemplateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotificationMessageTemplateRequestBuilderGetQueryParameters the Notification Message Templates.
type NotificationMessageTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NotificationMessageTemplateRequestBuilderPatchOptions options for Patch
type NotificationMessageTemplateRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NotificationMessageTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewNotificationMessageTemplateRequestBuilderInternal instantiates a new NotificationMessageTemplateRequestBuilder and sets the default values.
func NewNotificationMessageTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotificationMessageTemplateRequestBuilder) {
    m := &NotificationMessageTemplateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/notificationMessageTemplates/{notificationMessageTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNotificationMessageTemplateRequestBuilder instantiates a new NotificationMessageTemplateRequestBuilder and sets the default values.
func NewNotificationMessageTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotificationMessageTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotificationMessageTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) CreateDeleteRequestInformation(options *NotificationMessageTemplateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) CreateGetRequestInformation(options *NotificationMessageTemplateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) CreatePatchRequestInformation(options *NotificationMessageTemplateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) Delete(options *NotificationMessageTemplateRequestBuilderDeleteOptions)(error) {
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
// Get the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) Get(options *NotificationMessageTemplateRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NotificationMessageTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewNotificationMessageTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NotificationMessageTemplate), nil
}
func (m *NotificationMessageTemplateRequestBuilder) LocalizedNotificationMessages()(*icc3fe1a1c8b964bc176be6a6d48621fc3548efaeadcb3b10e136af2f19764752.LocalizedNotificationMessagesRequestBuilder) {
    return icc3fe1a1c8b964bc176be6a6d48621fc3548efaeadcb3b10e136af2f19764752.NewLocalizedNotificationMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizedNotificationMessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.notificationMessageTemplates.item.localizedNotificationMessages.item collection
func (m *NotificationMessageTemplateRequestBuilder) LocalizedNotificationMessagesById(id string)(*i079873d8184290e06e63c9a04520d4cbb886c99c078b4ea0a493ba22bc14901e.LocalizedNotificationMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["localizedNotificationMessage_id"] = id
    }
    return i079873d8184290e06e63c9a04520d4cbb886c99c078b4ea0a493ba22bc14901e.NewLocalizedNotificationMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) Patch(options *NotificationMessageTemplateRequestBuilderPatchOptions)(error) {
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
func (m *NotificationMessageTemplateRequestBuilder) SendTestMessage()(*if10e6b46a7479cb721a918b2370be92792065fdd20e80c33685b2f50841a32b9.SendTestMessageRequestBuilder) {
    return if10e6b46a7479cb721a918b2370be92792065fdd20e80c33685b2f50841a32b9.NewSendTestMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
