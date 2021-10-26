package lastsharedmethod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0a033e412571e1945babcf034dd8ffd565cf13cf8f5ce54b0b7cf782ce0210a7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange"
    i265e40e0c2e17ab364fcac1c3d78056922e09bd40345d83dc564a6b9831d29a5 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/mobileappcontentfile"
    i405d48f120f93a48e9a42bdca4c47228601b0a88f5e167b2a0498fd771ca2f4c "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/managedappprotection"
    i49c48c64c6a086f62ee1f68833237d81b389acf6466a8dc719e37e92335d9146 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/windowsinformationprotection"
    i565ca0b1b9ad912766fbf0e58273b9b820797d53acf7b1a52b069c653ed89642 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/calendarsharingmessage"
    i6cc0aa29fc3d62fe6f6b63130592b1d29cd1d8fb11958714c2e1742bca9aebc4 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/printdocument"
    i7683100b0875413d2b2c4dcf5d11783cac950aa58c66828826f8aece48cbf65a "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrangefill"
    ia9cce8d9fe15f9c6483c6dc727188f8dd88d1e7ba94f24973a7777efb1d5f9e9 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/targetedmanagedappprotection"
    ib3f96e09937f4ccaf5d8a6c187f8e47d8d4c861bf85934c288824052c89ddf6b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/printjob"
    ibf3077bbc251729ad3fc7d0b16ffb0f2f39a470f7076ba6e5686549a2f39d417 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrangesort"
    ic6098b99dd64267b445f65f733d45bab2c4faf563347d66dd631305db1eead87 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrangeformat"
    ic8c843b3bce9c85560ebc9e7432a301170373600b571a8b38f8cc1b520ed2eb7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/schedulechangerequest"
    id0b8429437699a0aeb4d26b39f67b7b2828ab2d3b5b1b04736cb459af67f6fea "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrangeview"
    ieea8499a6b9140ef511ce921340edd0a6ffa314d2dc7545baba72a6a52627bd1 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/ref"
)

// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod
type LastSharedMethodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get lastSharedMethod from me
type LastSharedMethodRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *LastSharedMethodRequestBuilder) CalendarSharingMessage()(*i565ca0b1b9ad912766fbf0e58273b9b820797d53acf7b1a52b069c653ed89642.CalendarSharingMessageRequestBuilder) {
    return i565ca0b1b9ad912766fbf0e58273b9b820797d53acf7b1a52b069c653ed89642.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new LastSharedMethodRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewLastSharedMethodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LastSharedMethodRequestBuilder) {
    m := &LastSharedMethodRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/lastSharedMethod{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new LastSharedMethodRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewLastSharedMethodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LastSharedMethodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLastSharedMethodRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lastSharedMethod from me
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *LastSharedMethodRequestBuilder) CreateGetRequestInformation(q func (value *LastSharedMethodRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(LastSharedMethodRequestBuilderGetQueryParameters)
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
// Get lastSharedMethod from me
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *LastSharedMethodRequestBuilder) Get(q func (value *LastSharedMethodRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEntity() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity), nil
}
func (m *LastSharedMethodRequestBuilder) ManagedAppProtection()(*i405d48f120f93a48e9a42bdca4c47228601b0a88f5e167b2a0498fd771ca2f4c.ManagedAppProtectionRequestBuilder) {
    return i405d48f120f93a48e9a42bdca4c47228601b0a88f5e167b2a0498fd771ca2f4c.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) MobileAppContentFile()(*i265e40e0c2e17ab364fcac1c3d78056922e09bd40345d83dc564a6b9831d29a5.MobileAppContentFileRequestBuilder) {
    return i265e40e0c2e17ab364fcac1c3d78056922e09bd40345d83dc564a6b9831d29a5.NewMobileAppContentFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) PrintDocument()(*i6cc0aa29fc3d62fe6f6b63130592b1d29cd1d8fb11958714c2e1742bca9aebc4.PrintDocumentRequestBuilder) {
    return i6cc0aa29fc3d62fe6f6b63130592b1d29cd1d8fb11958714c2e1742bca9aebc4.NewPrintDocumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) PrintJob()(*ib3f96e09937f4ccaf5d8a6c187f8e47d8d4c861bf85934c288824052c89ddf6b.PrintJobRequestBuilder) {
    return ib3f96e09937f4ccaf5d8a6c187f8e47d8d4c861bf85934c288824052c89ddf6b.NewPrintJobRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) Ref()(*ieea8499a6b9140ef511ce921340edd0a6ffa314d2dc7545baba72a6a52627bd1.RefRequestBuilder) {
    return ieea8499a6b9140ef511ce921340edd0a6ffa314d2dc7545baba72a6a52627bd1.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) ScheduleChangeRequest()(*ic8c843b3bce9c85560ebc9e7432a301170373600b571a8b38f8cc1b520ed2eb7.ScheduleChangeRequestRequestBuilder) {
    return ic8c843b3bce9c85560ebc9e7432a301170373600b571a8b38f8cc1b520ed2eb7.NewScheduleChangeRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) TargetedManagedAppProtection()(*ia9cce8d9fe15f9c6483c6dc727188f8dd88d1e7ba94f24973a7777efb1d5f9e9.TargetedManagedAppProtectionRequestBuilder) {
    return ia9cce8d9fe15f9c6483c6dc727188f8dd88d1e7ba94f24973a7777efb1d5f9e9.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WindowsInformationProtection()(*i49c48c64c6a086f62ee1f68833237d81b389acf6466a8dc719e37e92335d9146.WindowsInformationProtectionRequestBuilder) {
    return i49c48c64c6a086f62ee1f68833237d81b389acf6466a8dc719e37e92335d9146.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRange()(*i0a033e412571e1945babcf034dd8ffd565cf13cf8f5ce54b0b7cf782ce0210a7.WorkbookRangeRequestBuilder) {
    return i0a033e412571e1945babcf034dd8ffd565cf13cf8f5ce54b0b7cf782ce0210a7.NewWorkbookRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRangeFill()(*i7683100b0875413d2b2c4dcf5d11783cac950aa58c66828826f8aece48cbf65a.WorkbookRangeFillRequestBuilder) {
    return i7683100b0875413d2b2c4dcf5d11783cac950aa58c66828826f8aece48cbf65a.NewWorkbookRangeFillRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRangeFormat()(*ic6098b99dd64267b445f65f733d45bab2c4faf563347d66dd631305db1eead87.WorkbookRangeFormatRequestBuilder) {
    return ic6098b99dd64267b445f65f733d45bab2c4faf563347d66dd631305db1eead87.NewWorkbookRangeFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRangeSort()(*ibf3077bbc251729ad3fc7d0b16ffb0f2f39a470f7076ba6e5686549a2f39d417.WorkbookRangeSortRequestBuilder) {
    return ibf3077bbc251729ad3fc7d0b16ffb0f2f39a470f7076ba6e5686549a2f39d417.NewWorkbookRangeSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRangeView()(*id0b8429437699a0aeb4d26b39f67b7b2828ab2d3b5b1b04736cb459af67f6fea.WorkbookRangeViewRequestBuilder) {
    return id0b8429437699a0aeb4d26b39f67b7b2828ab2d3b5b1b04736cb459af67f6fea.NewWorkbookRangeViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
