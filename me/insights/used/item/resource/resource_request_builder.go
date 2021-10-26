package resource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i107d9ae42c7ac65b388a10b884476850058b8219d2aa81890b49ded2bee00b96 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/windowsinformationprotection"
    i1429d6e2a4ee92153e765d74efc3a1698c621c7bff18d626d6323c8fe207d52b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/printdocument"
    i1fb984c261149e510758d4f79d6a05df1668b3f93b94f0ca606302d3f5cfbba6 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/managedappprotection"
    i3106f034381a839ffaa7ed14980bc96ae3ffda145a8feef3095a5644a025d5f5 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/targetedmanagedappprotection"
    i380f882243d3b01ee5998bc5277bb5b8cd51f88b6780f54ca772b6438de2c907 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrangefill"
    i3899a9ba0baa27c5fe68a34b5ecfbca54601b5377612320d3d830977db640be1 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/mobileappcontentfile"
    i53fc521cad616e0338103f1a3d3f350146b47b1d1122b2cc6d25f66fc30305d6 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrangeview"
    i5734231ccae3cd86a3771bee7d36ee043b680174b0719239d8a7c93de272eff4 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/printjob"
    i62b687cd6b7e485c303b6bb216790b11502ddc55e296673f7f2a481a70fc3dea "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/calendarsharingmessage"
    i787d55c8356360bab8bccd716ddd3590c7fb9673213ab43e79a1df1cf4c3fa64 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrangeformat"
    ib2396b660c602ccf38644069b153c5cbd45ad9ae18750ad44428e4d46c0288e9 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/schedulechangerequest"
    icef874203b1f8f748886040eb3e5ad25da1292646996c1141cef7046c9125fac "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange"
    idc093d52f6623e85aaab6944a7e92f6691e8582f87e4683821b7ad6d35c9a616 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/ref"
    idd14bab8677d30ce63eaddea348931c8eb5a02e1d44c4925634f485d79e52037 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrangesort"
)

// Builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource
type ResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
type ResourceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *ResourceRequestBuilder) CalendarSharingMessage()(*i62b687cd6b7e485c303b6bb216790b11502ddc55e296673f7f2a481a70fc3dea.CalendarSharingMessageRequestBuilder) {
    return i62b687cd6b7e485c303b6bb216790b11502ddc55e296673f7f2a481a70fc3dea.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ResourceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceRequestBuilder) {
    m := &ResourceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/used/{usedInsight_id}/resource{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ResourceRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ResourceRequestBuilder) CreateGetRequestInformation(q func (value *ResourceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ResourceRequestBuilderGetQueryParameters)
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
// Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ResourceRequestBuilder) Get(q func (value *ResourceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity, error) {
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
func (m *ResourceRequestBuilder) ManagedAppProtection()(*i1fb984c261149e510758d4f79d6a05df1668b3f93b94f0ca606302d3f5cfbba6.ManagedAppProtectionRequestBuilder) {
    return i1fb984c261149e510758d4f79d6a05df1668b3f93b94f0ca606302d3f5cfbba6.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) MobileAppContentFile()(*i3899a9ba0baa27c5fe68a34b5ecfbca54601b5377612320d3d830977db640be1.MobileAppContentFileRequestBuilder) {
    return i3899a9ba0baa27c5fe68a34b5ecfbca54601b5377612320d3d830977db640be1.NewMobileAppContentFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintDocument()(*i1429d6e2a4ee92153e765d74efc3a1698c621c7bff18d626d6323c8fe207d52b.PrintDocumentRequestBuilder) {
    return i1429d6e2a4ee92153e765d74efc3a1698c621c7bff18d626d6323c8fe207d52b.NewPrintDocumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintJob()(*i5734231ccae3cd86a3771bee7d36ee043b680174b0719239d8a7c93de272eff4.PrintJobRequestBuilder) {
    return i5734231ccae3cd86a3771bee7d36ee043b680174b0719239d8a7c93de272eff4.NewPrintJobRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) Ref()(*idc093d52f6623e85aaab6944a7e92f6691e8582f87e4683821b7ad6d35c9a616.RefRequestBuilder) {
    return idc093d52f6623e85aaab6944a7e92f6691e8582f87e4683821b7ad6d35c9a616.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) ScheduleChangeRequest()(*ib2396b660c602ccf38644069b153c5cbd45ad9ae18750ad44428e4d46c0288e9.ScheduleChangeRequestRequestBuilder) {
    return ib2396b660c602ccf38644069b153c5cbd45ad9ae18750ad44428e4d46c0288e9.NewScheduleChangeRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) TargetedManagedAppProtection()(*i3106f034381a839ffaa7ed14980bc96ae3ffda145a8feef3095a5644a025d5f5.TargetedManagedAppProtectionRequestBuilder) {
    return i3106f034381a839ffaa7ed14980bc96ae3ffda145a8feef3095a5644a025d5f5.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WindowsInformationProtection()(*i107d9ae42c7ac65b388a10b884476850058b8219d2aa81890b49ded2bee00b96.WindowsInformationProtectionRequestBuilder) {
    return i107d9ae42c7ac65b388a10b884476850058b8219d2aa81890b49ded2bee00b96.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRange()(*icef874203b1f8f748886040eb3e5ad25da1292646996c1141cef7046c9125fac.WorkbookRangeRequestBuilder) {
    return icef874203b1f8f748886040eb3e5ad25da1292646996c1141cef7046c9125fac.NewWorkbookRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFill()(*i380f882243d3b01ee5998bc5277bb5b8cd51f88b6780f54ca772b6438de2c907.WorkbookRangeFillRequestBuilder) {
    return i380f882243d3b01ee5998bc5277bb5b8cd51f88b6780f54ca772b6438de2c907.NewWorkbookRangeFillRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFormat()(*i787d55c8356360bab8bccd716ddd3590c7fb9673213ab43e79a1df1cf4c3fa64.WorkbookRangeFormatRequestBuilder) {
    return i787d55c8356360bab8bccd716ddd3590c7fb9673213ab43e79a1df1cf4c3fa64.NewWorkbookRangeFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeSort()(*idd14bab8677d30ce63eaddea348931c8eb5a02e1d44c4925634f485d79e52037.WorkbookRangeSortRequestBuilder) {
    return idd14bab8677d30ce63eaddea348931c8eb5a02e1d44c4925634f485d79e52037.NewWorkbookRangeSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeView()(*i53fc521cad616e0338103f1a3d3f350146b47b1d1122b2cc6d25f66fc30305d6.WorkbookRangeViewRequestBuilder) {
    return i53fc521cad616e0338103f1a3d3f350146b47b1d1122b2cc6d25f66fc30305d6.NewWorkbookRangeViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
