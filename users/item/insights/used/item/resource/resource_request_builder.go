package resource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i09507481e6d2771384a602969a9296dd81422a9aa6cbf3792949c2330818fe27 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/managedappprotection"
    i0ae9916e8f3c05c45ee170b46064159741b4d56af279d3cdeaceecb471451564 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/windowsinformationprotection"
    i16d48cc90853d8bea741488599b7c134badc2d679834c0b865a98f9d00de6520 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/mobileappcontentfile"
    i20f9fc378c0b3dd78f5071ecb3d0a1caeaf079d864776908308c076a648c96c7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/caseexportoperation"
    i28f0c52b8732c8205239fd962478b9e6f1b680c90da1db40cb50abf1f792eb1f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/ref"
    i29bc9fc3241164c6d3dfa03b83c3e5a0bf3168ef35052eff3fb6c35091fd8c66 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrangesort"
    i2c6aa2662728e6c096102f9113d7d1d2bdfc5dfa15816aee10a61a92677a931a "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange"
    i61dbf79e1fa3b509f7f5fa59f5b8db7e29db1a49ca859fe03f3fa415c79366b7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrangeview"
    i97108c709e4028037f4b310dfb83bf80bbbe55de04a9237452f7befda3f1d520 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/targetedmanagedappprotection"
    i9cb53eb794e87ffa23e16cb3627a7a85def0294e7203cc94c3b1d4745fe64b89 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/printjob"
    ia33b6428b780c7938e0f10dc4cc5d7619f525dad067fa9aab5e116580b6eedf2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/schedulechangerequest"
    ic2305ade8f5cf99ed4fad3b7d3dfcf16f67b6e3a1bcbb312c66f9192cf64b07f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/printdocument"
    ieaff7197006ac121e523da2ff04f1f6c750d2d2615637b2eea5cf1c617e4bd13 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrangeformat"
    if00eeab0c353149d2481dfc22f7e6d8156f40941ab73d3473a7d904f4d56bef3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrangefill"
    ifadb9c06695851288d7f222b713a8b88b32e45403b46d20966111b125faacac6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/calendarsharingmessage"
)

// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource
type ResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
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
// Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
type ResourceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *ResourceRequestBuilder) CalendarSharingMessage()(*ifadb9c06695851288d7f222b713a8b88b32e45403b46d20966111b125faacac6.CalendarSharingMessageRequestBuilder) {
    return ifadb9c06695851288d7f222b713a8b88b32e45403b46d20966111b125faacac6.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) CaseExportOperation()(*i20f9fc378c0b3dd78f5071ecb3d0a1caeaf079d864776908308c076a648c96c7.CaseExportOperationRequestBuilder) {
    return i20f9fc378c0b3dd78f5071ecb3d0a1caeaf079d864776908308c076a648c96c7.NewCaseExportOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ResourceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceRequestBuilder) {
    m := &ResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/used/{usedInsight_id}/resource{?select,expand}";
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
//  - options : Options for the request
func (m *ResourceRequestBuilder) CreateGetRequestInformation(options *ResourceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
// Parameters:
//  - options : Options for the request
func (m *ResourceRequestBuilder) Get(options *ResourceRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEntity() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity), nil
}
func (m *ResourceRequestBuilder) ManagedAppProtection()(*i09507481e6d2771384a602969a9296dd81422a9aa6cbf3792949c2330818fe27.ManagedAppProtectionRequestBuilder) {
    return i09507481e6d2771384a602969a9296dd81422a9aa6cbf3792949c2330818fe27.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) MobileAppContentFile()(*i16d48cc90853d8bea741488599b7c134badc2d679834c0b865a98f9d00de6520.MobileAppContentFileRequestBuilder) {
    return i16d48cc90853d8bea741488599b7c134badc2d679834c0b865a98f9d00de6520.NewMobileAppContentFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintDocument()(*ic2305ade8f5cf99ed4fad3b7d3dfcf16f67b6e3a1bcbb312c66f9192cf64b07f.PrintDocumentRequestBuilder) {
    return ic2305ade8f5cf99ed4fad3b7d3dfcf16f67b6e3a1bcbb312c66f9192cf64b07f.NewPrintDocumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintJob()(*i9cb53eb794e87ffa23e16cb3627a7a85def0294e7203cc94c3b1d4745fe64b89.PrintJobRequestBuilder) {
    return i9cb53eb794e87ffa23e16cb3627a7a85def0294e7203cc94c3b1d4745fe64b89.NewPrintJobRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) Ref()(*i28f0c52b8732c8205239fd962478b9e6f1b680c90da1db40cb50abf1f792eb1f.RefRequestBuilder) {
    return i28f0c52b8732c8205239fd962478b9e6f1b680c90da1db40cb50abf1f792eb1f.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) ScheduleChangeRequest()(*ia33b6428b780c7938e0f10dc4cc5d7619f525dad067fa9aab5e116580b6eedf2.ScheduleChangeRequestRequestBuilder) {
    return ia33b6428b780c7938e0f10dc4cc5d7619f525dad067fa9aab5e116580b6eedf2.NewScheduleChangeRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) TargetedManagedAppProtection()(*i97108c709e4028037f4b310dfb83bf80bbbe55de04a9237452f7befda3f1d520.TargetedManagedAppProtectionRequestBuilder) {
    return i97108c709e4028037f4b310dfb83bf80bbbe55de04a9237452f7befda3f1d520.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WindowsInformationProtection()(*i0ae9916e8f3c05c45ee170b46064159741b4d56af279d3cdeaceecb471451564.WindowsInformationProtectionRequestBuilder) {
    return i0ae9916e8f3c05c45ee170b46064159741b4d56af279d3cdeaceecb471451564.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRange()(*i2c6aa2662728e6c096102f9113d7d1d2bdfc5dfa15816aee10a61a92677a931a.WorkbookRangeRequestBuilder) {
    return i2c6aa2662728e6c096102f9113d7d1d2bdfc5dfa15816aee10a61a92677a931a.NewWorkbookRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFill()(*if00eeab0c353149d2481dfc22f7e6d8156f40941ab73d3473a7d904f4d56bef3.WorkbookRangeFillRequestBuilder) {
    return if00eeab0c353149d2481dfc22f7e6d8156f40941ab73d3473a7d904f4d56bef3.NewWorkbookRangeFillRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFormat()(*ieaff7197006ac121e523da2ff04f1f6c750d2d2615637b2eea5cf1c617e4bd13.WorkbookRangeFormatRequestBuilder) {
    return ieaff7197006ac121e523da2ff04f1f6c750d2d2615637b2eea5cf1c617e4bd13.NewWorkbookRangeFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeSort()(*i29bc9fc3241164c6d3dfa03b83c3e5a0bf3168ef35052eff3fb6c35091fd8c66.WorkbookRangeSortRequestBuilder) {
    return i29bc9fc3241164c6d3dfa03b83c3e5a0bf3168ef35052eff3fb6c35091fd8c66.NewWorkbookRangeSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeView()(*i61dbf79e1fa3b509f7f5fa59f5b8db7e29db1a49ca859fe03f3fa415c79366b7.WorkbookRangeViewRequestBuilder) {
    return i61dbf79e1fa3b509f7f5fa59f5b8db7e29db1a49ca859fe03f3fa415c79366b7.NewWorkbookRangeViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
