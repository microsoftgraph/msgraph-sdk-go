package serviceannouncement

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5a1d633dc34eb5609de7c5475bea4c4ef223726f894d6018982fd790c70858ff "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages"
    i64b9866ddf18d005e45ccbc515fa8deeeb01d6d4ad4b07bdd67fd26595e01984 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/issues"
    i993eabb4740237d5836875ca4abee991a6d6f1070cef12f2f5cc6ca062b03f47 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/healthoverviews"
    i05cb520997f5472376a5296872f494ab5bc6678ca943b0a20d8a0a735c91b06c "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/healthoverviews/item"
    i2b040537addd757f1ad668fb2339fe6e9508193ac05e02dec09691e60f070df3 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/issues/item"
    ib231b71ad312efbdfea743bca259e9f04366799c7f96005b7055ab09c71f271c "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/item"
)

// Builds and executes requests for operations under \admin\serviceAnnouncement
type ServiceAnnouncementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// A container for service communications resources. Read-only.
type ServiceAnnouncementRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new ServiceAnnouncementRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewServiceAnnouncementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    m := &ServiceAnnouncementRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/admin/serviceAnnouncement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ServiceAnnouncementRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewServiceAnnouncementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementRequestBuilderInternal(urlParams, requestAdapter)
}
// A container for service communications resources. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *ServiceAnnouncementRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// A container for service communications resources. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ServiceAnnouncementRequestBuilder) CreateGetRequestInformation(q func (value *ServiceAnnouncementRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ServiceAnnouncementRequestBuilderGetQueryParameters)
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
// A container for service communications resources. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ServiceAnnouncementRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceAnnouncement, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// A container for service communications resources. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ServiceAnnouncementRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// A container for service communications resources. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ServiceAnnouncementRequestBuilder) Get(q func (value *ServiceAnnouncementRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceAnnouncement, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewServiceAnnouncement() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceAnnouncement), nil
}
func (m *ServiceAnnouncementRequestBuilder) HealthOverviews()(*i993eabb4740237d5836875ca4abee991a6d6f1070cef12f2f5cc6ca062b03f47.HealthOverviewsRequestBuilder) {
    return i993eabb4740237d5836875ca4abee991a6d6f1070cef12f2f5cc6ca062b03f47.NewHealthOverviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.healthOverviews.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ServiceAnnouncementRequestBuilder) HealthOverviewsById(id string)(*i05cb520997f5472376a5296872f494ab5bc6678ca943b0a20d8a0a735c91b06c.ServiceHealthRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceHealth_id"] = id
    }
    return i05cb520997f5472376a5296872f494ab5bc6678ca943b0a20d8a0a735c91b06c.NewServiceHealthRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) Issues()(*i64b9866ddf18d005e45ccbc515fa8deeeb01d6d4ad4b07bdd67fd26595e01984.IssuesRequestBuilder) {
    return i64b9866ddf18d005e45ccbc515fa8deeeb01d6d4ad4b07bdd67fd26595e01984.NewIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.issues.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ServiceAnnouncementRequestBuilder) IssuesById(id string)(*i2b040537addd757f1ad668fb2339fe6e9508193ac05e02dec09691e60f070df3.ServiceHealthIssueRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceHealthIssue_id"] = id
    }
    return i2b040537addd757f1ad668fb2339fe6e9508193ac05e02dec09691e60f070df3.NewServiceHealthIssueRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) Messages()(*i5a1d633dc34eb5609de7c5475bea4c4ef223726f894d6018982fd790c70858ff.MessagesRequestBuilder) {
    return i5a1d633dc34eb5609de7c5475bea4c4ef223726f894d6018982fd790c70858ff.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.messages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ServiceAnnouncementRequestBuilder) MessagesById(id string)(*ib231b71ad312efbdfea743bca259e9f04366799c7f96005b7055ab09c71f271c.ServiceUpdateMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceUpdateMessage_id"] = id
    }
    return ib231b71ad312efbdfea743bca259e9f04366799c7f96005b7055ab09c71f271c.NewServiceUpdateMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// A container for service communications resources. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ServiceAnnouncementRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceAnnouncement, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
