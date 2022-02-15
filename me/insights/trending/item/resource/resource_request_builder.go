package resource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i26d7cebd32d3e662f5ac28b3768366f7d8d295776ce25bd75f83bd6f81cd6cd0 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/printdocument"
    i300e69c215fef08ebcfcfe52daf43625724dc4e851ff683e972488cc18531540 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/printjob"
    i339d6ee32082bf4a240b0c9283afe3c8cc92c20c873f30a27dc1f0c814dcdd30 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange"
    i4051f44cf261524787095cbd29679f142c1f52e5b0244f3a8effa077ea5868a3 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/ref"
    i6400b51d214acd487fff63835b7fb2c6c68873cc4f1e010e14c4c0728c92b225 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/managedappprotection"
    i6893fc0a3442981428ad0f00ad592143ba65d4895acac83b9bacd362618f7f03 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrangefill"
    i84f8f5a64514eebc8bfcd77e0aed520ff21eec6aa53873f1894bd18b5d9aedd0 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/calendarsharingmessage"
    i89966168170997d2aed24b007473bf974b970bb390f147b3b055f424e2338841 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/schedulechangerequest"
    i93b4858e144352162776d0b982472c140438e266bc138cb7bd1deacedcdc6c06 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrangesort"
    iaf1bbc68b495ef6c434777ba234a41b7129a04adaed4c5d34c058d9fdca5eb2f "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/windowsinformationprotection"
    ic09e25ea07b805ea73e40360cd847c88a966af5d477ece74aeb6fffa36cb13e7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/targetedmanagedappprotection"
    ic62b66384099a7e9134470de00370e4c1df58b0e6c5476f324a976f8dfff968e "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/mobileappcontentfile"
    ifbdb9332e5ed4f21f4492c59da440ab5a17b5c7f1ecbe0d619516b6a0dd24a77 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrangeview"
    ifdcba498de87955cbc3a2bf83de155c2ee78f3458c9bd2f2b2baf5161156dd23 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrangeformat"
)

// ResourceRequestBuilder builds and executes requests for operations under \me\insights\trending\{trending-id}\resource
type ResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ResourceRequestBuilderGetOptions options for Get
type ResourceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ResourceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ResourceRequestBuilderGetQueryParameters used for navigating to the trending document.
type ResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *ResourceRequestBuilder) CalendarSharingMessage()(*i84f8f5a64514eebc8bfcd77e0aed520ff21eec6aa53873f1894bd18b5d9aedd0.CalendarSharingMessageRequestBuilder) {
    return i84f8f5a64514eebc8bfcd77e0aed520ff21eec6aa53873f1894bd18b5d9aedd0.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewResourceRequestBuilderInternal instantiates a new ResourceRequestBuilder and sets the default values.
func NewResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceRequestBuilder) {
    m := &ResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/trending/{trending_id}/resource{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResourceRequestBuilder instantiates a new ResourceRequestBuilder and sets the default values.
func NewResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation used for navigating to the trending document.
func (m *ResourceRequestBuilder) CreateGetRequestInformation(options *ResourceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get used for navigating to the trending document.
func (m *ResourceRequestBuilder) Get(options *ResourceRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEntity() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity), nil
}
func (m *ResourceRequestBuilder) ManagedAppProtection()(*i6400b51d214acd487fff63835b7fb2c6c68873cc4f1e010e14c4c0728c92b225.ManagedAppProtectionRequestBuilder) {
    return i6400b51d214acd487fff63835b7fb2c6c68873cc4f1e010e14c4c0728c92b225.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) MobileAppContentFile()(*ic62b66384099a7e9134470de00370e4c1df58b0e6c5476f324a976f8dfff968e.MobileAppContentFileRequestBuilder) {
    return ic62b66384099a7e9134470de00370e4c1df58b0e6c5476f324a976f8dfff968e.NewMobileAppContentFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintDocument()(*i26d7cebd32d3e662f5ac28b3768366f7d8d295776ce25bd75f83bd6f81cd6cd0.PrintDocumentRequestBuilder) {
    return i26d7cebd32d3e662f5ac28b3768366f7d8d295776ce25bd75f83bd6f81cd6cd0.NewPrintDocumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintJob()(*i300e69c215fef08ebcfcfe52daf43625724dc4e851ff683e972488cc18531540.PrintJobRequestBuilder) {
    return i300e69c215fef08ebcfcfe52daf43625724dc4e851ff683e972488cc18531540.NewPrintJobRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) Ref()(*i4051f44cf261524787095cbd29679f142c1f52e5b0244f3a8effa077ea5868a3.RefRequestBuilder) {
    return i4051f44cf261524787095cbd29679f142c1f52e5b0244f3a8effa077ea5868a3.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) ScheduleChangeRequest()(*i89966168170997d2aed24b007473bf974b970bb390f147b3b055f424e2338841.ScheduleChangeRequestRequestBuilder) {
    return i89966168170997d2aed24b007473bf974b970bb390f147b3b055f424e2338841.NewScheduleChangeRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) TargetedManagedAppProtection()(*ic09e25ea07b805ea73e40360cd847c88a966af5d477ece74aeb6fffa36cb13e7.TargetedManagedAppProtectionRequestBuilder) {
    return ic09e25ea07b805ea73e40360cd847c88a966af5d477ece74aeb6fffa36cb13e7.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WindowsInformationProtection()(*iaf1bbc68b495ef6c434777ba234a41b7129a04adaed4c5d34c058d9fdca5eb2f.WindowsInformationProtectionRequestBuilder) {
    return iaf1bbc68b495ef6c434777ba234a41b7129a04adaed4c5d34c058d9fdca5eb2f.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRange()(*i339d6ee32082bf4a240b0c9283afe3c8cc92c20c873f30a27dc1f0c814dcdd30.WorkbookRangeRequestBuilder) {
    return i339d6ee32082bf4a240b0c9283afe3c8cc92c20c873f30a27dc1f0c814dcdd30.NewWorkbookRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFill()(*i6893fc0a3442981428ad0f00ad592143ba65d4895acac83b9bacd362618f7f03.WorkbookRangeFillRequestBuilder) {
    return i6893fc0a3442981428ad0f00ad592143ba65d4895acac83b9bacd362618f7f03.NewWorkbookRangeFillRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFormat()(*ifdcba498de87955cbc3a2bf83de155c2ee78f3458c9bd2f2b2baf5161156dd23.WorkbookRangeFormatRequestBuilder) {
    return ifdcba498de87955cbc3a2bf83de155c2ee78f3458c9bd2f2b2baf5161156dd23.NewWorkbookRangeFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeSort()(*i93b4858e144352162776d0b982472c140438e266bc138cb7bd1deacedcdc6c06.WorkbookRangeSortRequestBuilder) {
    return i93b4858e144352162776d0b982472c140438e266bc138cb7bd1deacedcdc6c06.NewWorkbookRangeSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeView()(*ifbdb9332e5ed4f21f4492c59da440ab5a17b5c7f1ecbe0d619516b6a0dd24a77.WorkbookRangeViewRequestBuilder) {
    return ifbdb9332e5ed4f21f4492c59da440ab5a17b5c7f1ecbe0d619516b6a0dd24a77.NewWorkbookRangeViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
