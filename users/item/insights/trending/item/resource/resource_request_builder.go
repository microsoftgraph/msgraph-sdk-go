package resource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0508fb7692a84eb80a47399d8aab5898f05ca9cbebe2711b96db985c469e4d27 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/schedulechangerequest"
    i0d83af0c5ec4837536fd038b8172600dc37c88dfeeca7555f68e2ddb4a3085c2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/ref"
    i19f94efb6fc20640201c3ee04bf1383a7e135ee2fbfbf01d1b2adc8928875109 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/managedappprotection"
    i288642a73d5b1227925c09a9339283a10a06760ade2e86acd6a90645109b7351 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/printjob"
    i649d6110c4c93db4931e16eec288dd71976580b8ac61f92eecdbeff992c42eaf "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrangeview"
    i7d67d73a2580465483570d6bb36a28c49e06d5876f7b457528576eef2f263c95 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrangesort"
    i7d79a1ebb4b2030cbe3b093a4f6761343ee07964649c4ad023bb26bbcdc66ea4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/windowsinformationprotection"
    i906d4f14161c03dedaa24baf41081af9da98eff2b35dcf115f03f60d701a350c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/printdocument"
    i919e2f7c61a582a751e71e7f9d8c3fc3c17163411877b465756860f8c36a89c6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/calendarsharingmessage"
    iafd5ab9d6c238e9eadcc9f77d5bd940a2f23f4035145e85155ed355eb04708a9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange"
    ibceebb108ca739619379bc5738b0a7c30356db34ba8ab5eb0a44723691af2c9d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/mobileappcontentfile"
    ic1ec7a08e927a127fa7f7b342a9970799a3202f31b0a2d9fb38c3500bd905557 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrangeformat"
    id5a13ac26a3dd37d1067017bd48306686f285bc748329caaec5307240144b1ee "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/targetedmanagedappprotection"
    ide90b82c9976bd9a2bcbeb3e59519101171b7c3b844b3c76d6a77c8715537d8d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrangefill"
)

// Builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource
type ResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Used for navigating to the trending document.
type ResourceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *ResourceRequestBuilder) CalendarSharingMessage()(*i919e2f7c61a582a751e71e7f9d8c3fc3c17163411877b465756860f8c36a89c6.CalendarSharingMessageRequestBuilder) {
    return i919e2f7c61a582a751e71e7f9d8c3fc3c17163411877b465756860f8c36a89c6.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ResourceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceRequestBuilder) {
    m := &ResourceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/trending/{trending_id}/resource{?select,expand}";
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
// Used for navigating to the trending document.
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
// Used for navigating to the trending document.
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
func (m *ResourceRequestBuilder) ManagedAppProtection()(*i19f94efb6fc20640201c3ee04bf1383a7e135ee2fbfbf01d1b2adc8928875109.ManagedAppProtectionRequestBuilder) {
    return i19f94efb6fc20640201c3ee04bf1383a7e135ee2fbfbf01d1b2adc8928875109.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) MobileAppContentFile()(*ibceebb108ca739619379bc5738b0a7c30356db34ba8ab5eb0a44723691af2c9d.MobileAppContentFileRequestBuilder) {
    return ibceebb108ca739619379bc5738b0a7c30356db34ba8ab5eb0a44723691af2c9d.NewMobileAppContentFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintDocument()(*i906d4f14161c03dedaa24baf41081af9da98eff2b35dcf115f03f60d701a350c.PrintDocumentRequestBuilder) {
    return i906d4f14161c03dedaa24baf41081af9da98eff2b35dcf115f03f60d701a350c.NewPrintDocumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintJob()(*i288642a73d5b1227925c09a9339283a10a06760ade2e86acd6a90645109b7351.PrintJobRequestBuilder) {
    return i288642a73d5b1227925c09a9339283a10a06760ade2e86acd6a90645109b7351.NewPrintJobRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) Ref()(*i0d83af0c5ec4837536fd038b8172600dc37c88dfeeca7555f68e2ddb4a3085c2.RefRequestBuilder) {
    return i0d83af0c5ec4837536fd038b8172600dc37c88dfeeca7555f68e2ddb4a3085c2.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) ScheduleChangeRequest()(*i0508fb7692a84eb80a47399d8aab5898f05ca9cbebe2711b96db985c469e4d27.ScheduleChangeRequestRequestBuilder) {
    return i0508fb7692a84eb80a47399d8aab5898f05ca9cbebe2711b96db985c469e4d27.NewScheduleChangeRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) TargetedManagedAppProtection()(*id5a13ac26a3dd37d1067017bd48306686f285bc748329caaec5307240144b1ee.TargetedManagedAppProtectionRequestBuilder) {
    return id5a13ac26a3dd37d1067017bd48306686f285bc748329caaec5307240144b1ee.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WindowsInformationProtection()(*i7d79a1ebb4b2030cbe3b093a4f6761343ee07964649c4ad023bb26bbcdc66ea4.WindowsInformationProtectionRequestBuilder) {
    return i7d79a1ebb4b2030cbe3b093a4f6761343ee07964649c4ad023bb26bbcdc66ea4.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRange()(*iafd5ab9d6c238e9eadcc9f77d5bd940a2f23f4035145e85155ed355eb04708a9.WorkbookRangeRequestBuilder) {
    return iafd5ab9d6c238e9eadcc9f77d5bd940a2f23f4035145e85155ed355eb04708a9.NewWorkbookRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFill()(*ide90b82c9976bd9a2bcbeb3e59519101171b7c3b844b3c76d6a77c8715537d8d.WorkbookRangeFillRequestBuilder) {
    return ide90b82c9976bd9a2bcbeb3e59519101171b7c3b844b3c76d6a77c8715537d8d.NewWorkbookRangeFillRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFormat()(*ic1ec7a08e927a127fa7f7b342a9970799a3202f31b0a2d9fb38c3500bd905557.WorkbookRangeFormatRequestBuilder) {
    return ic1ec7a08e927a127fa7f7b342a9970799a3202f31b0a2d9fb38c3500bd905557.NewWorkbookRangeFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeSort()(*i7d67d73a2580465483570d6bb36a28c49e06d5876f7b457528576eef2f263c95.WorkbookRangeSortRequestBuilder) {
    return i7d67d73a2580465483570d6bb36a28c49e06d5876f7b457528576eef2f263c95.NewWorkbookRangeSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeView()(*i649d6110c4c93db4931e16eec288dd71976580b8ac61f92eecdbeff992c42eaf.WorkbookRangeViewRequestBuilder) {
    return i649d6110c4c93db4931e16eec288dd71976580b8ac61f92eecdbeff992c42eaf.NewWorkbookRangeViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
