package serviceannouncement

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5a1d633dc34eb5609de7c5475bea4c4ef223726f894d6018982fd790c70858ff "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages"
    i64b9866ddf18d005e45ccbc515fa8deeeb01d6d4ad4b07bdd67fd26595e01984 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/issues"
    i993eabb4740237d5836875ca4abee991a6d6f1070cef12f2f5cc6ca062b03f47 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/healthoverviews"
    i05cb520997f5472376a5296872f494ab5bc6678ca943b0a20d8a0a735c91b06c "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/healthoverviews/item"
    i2b040537addd757f1ad668fb2339fe6e9508193ac05e02dec09691e60f070df3 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/issues/item"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    ib231b71ad312efbdfea743bca259e9f04366799c7f96005b7055ab09c71f271c "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/item"
)

// ServiceAnnouncementRequestBuilder provides operations to manage the serviceAnnouncement property of the microsoft.graph.admin entity.
type ServiceAnnouncementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ServiceAnnouncementRequestBuilderDeleteOptions options for Delete
type ServiceAnnouncementRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServiceAnnouncementRequestBuilderGetOptions options for Get
type ServiceAnnouncementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServiceAnnouncementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServiceAnnouncementRequestBuilderGetQueryParameters a container for service communications resources. Read-only.
type ServiceAnnouncementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ServiceAnnouncementRequestBuilderPatchOptions options for Patch
type ServiceAnnouncementRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceAnnouncementable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewServiceAnnouncementRequestBuilderInternal instantiates a new ServiceAnnouncementRequestBuilder and sets the default values.
func NewServiceAnnouncementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    m := &ServiceAnnouncementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServiceAnnouncementRequestBuilder instantiates a new ServiceAnnouncementRequestBuilder and sets the default values.
func NewServiceAnnouncementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property serviceAnnouncement for admin
func (m *ServiceAnnouncementRequestBuilder) CreateDeleteRequestInformation(options *ServiceAnnouncementRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a container for service communications resources. Read-only.
func (m *ServiceAnnouncementRequestBuilder) CreateGetRequestInformation(options *ServiceAnnouncementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property serviceAnnouncement in admin
func (m *ServiceAnnouncementRequestBuilder) CreatePatchRequestInformation(options *ServiceAnnouncementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property serviceAnnouncement for admin
func (m *ServiceAnnouncementRequestBuilder) Delete(options *ServiceAnnouncementRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a container for service communications resources. Read-only.
func (m *ServiceAnnouncementRequestBuilder) Get(options *ServiceAnnouncementRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceAnnouncementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateServiceAnnouncementFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceAnnouncementable), nil
}
func (m *ServiceAnnouncementRequestBuilder) HealthOverviews()(*i993eabb4740237d5836875ca4abee991a6d6f1070cef12f2f5cc6ca062b03f47.HealthOverviewsRequestBuilder) {
    return i993eabb4740237d5836875ca4abee991a6d6f1070cef12f2f5cc6ca062b03f47.NewHealthOverviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HealthOverviewsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.healthOverviews.item collection
func (m *ServiceAnnouncementRequestBuilder) HealthOverviewsById(id string)(*i05cb520997f5472376a5296872f494ab5bc6678ca943b0a20d8a0a735c91b06c.ServiceHealthItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceHealth_id"] = id
    }
    return i05cb520997f5472376a5296872f494ab5bc6678ca943b0a20d8a0a735c91b06c.NewServiceHealthItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) Issues()(*i64b9866ddf18d005e45ccbc515fa8deeeb01d6d4ad4b07bdd67fd26595e01984.IssuesRequestBuilder) {
    return i64b9866ddf18d005e45ccbc515fa8deeeb01d6d4ad4b07bdd67fd26595e01984.NewIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IssuesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.issues.item collection
func (m *ServiceAnnouncementRequestBuilder) IssuesById(id string)(*i2b040537addd757f1ad668fb2339fe6e9508193ac05e02dec09691e60f070df3.ServiceHealthIssueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceHealthIssue_id"] = id
    }
    return i2b040537addd757f1ad668fb2339fe6e9508193ac05e02dec09691e60f070df3.NewServiceHealthIssueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) Messages()(*i5a1d633dc34eb5609de7c5475bea4c4ef223726f894d6018982fd790c70858ff.MessagesRequestBuilder) {
    return i5a1d633dc34eb5609de7c5475bea4c4ef223726f894d6018982fd790c70858ff.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.messages.item collection
func (m *ServiceAnnouncementRequestBuilder) MessagesById(id string)(*ib231b71ad312efbdfea743bca259e9f04366799c7f96005b7055ab09c71f271c.ServiceUpdateMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceUpdateMessage_id"] = id
    }
    return ib231b71ad312efbdfea743bca259e9f04366799c7f96005b7055ab09c71f271c.NewServiceUpdateMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property serviceAnnouncement in admin
func (m *ServiceAnnouncementRequestBuilder) Patch(options *ServiceAnnouncementRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
