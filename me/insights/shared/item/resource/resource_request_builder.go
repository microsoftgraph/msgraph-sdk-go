package resource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i08933ada4c7d4bc6ea52ac88871e54f5ce4b4020fb0a55079795a6c7334ae03e "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrangeformat"
    i0f2b8dc5e4c742d6e81d625e30abd3099b488f770a51bf91ae98385cf50ec912 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrangefill"
    i10528b898c1668e849aa60de10dbf73638fac8dbc5d2399fa16cb6ad3ca86ad8 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/windowsinformationprotection"
    i2673d49fa82c20b0410d6d611a6be87918ad4909221573025dd604b084537181 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/printdocument"
    i2b6591cad5702df779835ec26cfc97beb6d20b2f57b95f923198ba932a9c547b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/schedulechangerequest"
    i6e3d8092860de35c76a3012f261944057f6cd220fe2354d05357ae0c2e1de713 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrangesort"
    i8534f03c304cb82037d3c324680d28978791ef2ffb00579723aaebd0b59c3bbb "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/mobileappcontentfile"
    i93c9183c0c22deb2e6a0c34af8661aed0696a49a0e379250354c9394af2edbd5 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/ref"
    ia09a8d69d0bb994f310babff96c81a055aeb309c4279e8d072c70b06f2b51143 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/managedappprotection"
    ia9167ab8a5620016fd4b31fe30f2fa018daf03abf3240a162b0ab30f12d1abf9 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/targetedmanagedappprotection"
    ib101b34e7159adcd137c1e833f15a82941a098e6363bda2ae169edaeb7cc77e0 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrangeview"
    ibac798766f9a218466837d454ccd671c338384367395de2980f31bb3887d6215 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/calendarsharingmessage"
    icf8494f164c6dd444b0a2bd42d132adb9e3f038817f528a84985be93b6290aa0 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange"
    ie4d5700d495bc4a2cc7061c6607509408a3cca4be4bd731eeff67486e856b6a3 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/printjob"
)

type ResourceRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ResourceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ResourceRequestBuilder) CalendarSharingMessage()(*ibac798766f9a218466837d454ccd671c338384367395de2980f31bb3887d6215.CalendarSharingMessageRequestBuilder) {
    return ibac798766f9a218466837d454ccd671c338384367395de2980f31bb3887d6215.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceRequestBuilder) {
    m := &ResourceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/resource{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceRequestBuilderInternal(urlParams, requestAdapter)
}
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
func (m *ResourceRequestBuilder) ManagedAppProtection()(*ia09a8d69d0bb994f310babff96c81a055aeb309c4279e8d072c70b06f2b51143.ManagedAppProtectionRequestBuilder) {
    return ia09a8d69d0bb994f310babff96c81a055aeb309c4279e8d072c70b06f2b51143.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) MobileAppContentFile()(*i8534f03c304cb82037d3c324680d28978791ef2ffb00579723aaebd0b59c3bbb.MobileAppContentFileRequestBuilder) {
    return i8534f03c304cb82037d3c324680d28978791ef2ffb00579723aaebd0b59c3bbb.NewMobileAppContentFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintDocument()(*i2673d49fa82c20b0410d6d611a6be87918ad4909221573025dd604b084537181.PrintDocumentRequestBuilder) {
    return i2673d49fa82c20b0410d6d611a6be87918ad4909221573025dd604b084537181.NewPrintDocumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintJob()(*ie4d5700d495bc4a2cc7061c6607509408a3cca4be4bd731eeff67486e856b6a3.PrintJobRequestBuilder) {
    return ie4d5700d495bc4a2cc7061c6607509408a3cca4be4bd731eeff67486e856b6a3.NewPrintJobRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) Ref()(*i93c9183c0c22deb2e6a0c34af8661aed0696a49a0e379250354c9394af2edbd5.RefRequestBuilder) {
    return i93c9183c0c22deb2e6a0c34af8661aed0696a49a0e379250354c9394af2edbd5.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) ScheduleChangeRequest()(*i2b6591cad5702df779835ec26cfc97beb6d20b2f57b95f923198ba932a9c547b.ScheduleChangeRequestRequestBuilder) {
    return i2b6591cad5702df779835ec26cfc97beb6d20b2f57b95f923198ba932a9c547b.NewScheduleChangeRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) TargetedManagedAppProtection()(*ia9167ab8a5620016fd4b31fe30f2fa018daf03abf3240a162b0ab30f12d1abf9.TargetedManagedAppProtectionRequestBuilder) {
    return ia9167ab8a5620016fd4b31fe30f2fa018daf03abf3240a162b0ab30f12d1abf9.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WindowsInformationProtection()(*i10528b898c1668e849aa60de10dbf73638fac8dbc5d2399fa16cb6ad3ca86ad8.WindowsInformationProtectionRequestBuilder) {
    return i10528b898c1668e849aa60de10dbf73638fac8dbc5d2399fa16cb6ad3ca86ad8.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRange()(*icf8494f164c6dd444b0a2bd42d132adb9e3f038817f528a84985be93b6290aa0.WorkbookRangeRequestBuilder) {
    return icf8494f164c6dd444b0a2bd42d132adb9e3f038817f528a84985be93b6290aa0.NewWorkbookRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFill()(*i0f2b8dc5e4c742d6e81d625e30abd3099b488f770a51bf91ae98385cf50ec912.WorkbookRangeFillRequestBuilder) {
    return i0f2b8dc5e4c742d6e81d625e30abd3099b488f770a51bf91ae98385cf50ec912.NewWorkbookRangeFillRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFormat()(*i08933ada4c7d4bc6ea52ac88871e54f5ce4b4020fb0a55079795a6c7334ae03e.WorkbookRangeFormatRequestBuilder) {
    return i08933ada4c7d4bc6ea52ac88871e54f5ce4b4020fb0a55079795a6c7334ae03e.NewWorkbookRangeFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeSort()(*i6e3d8092860de35c76a3012f261944057f6cd220fe2354d05357ae0c2e1de713.WorkbookRangeSortRequestBuilder) {
    return i6e3d8092860de35c76a3012f261944057f6cd220fe2354d05357ae0c2e1de713.NewWorkbookRangeSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeView()(*ib101b34e7159adcd137c1e833f15a82941a098e6363bda2ae169edaeb7cc77e0.WorkbookRangeViewRequestBuilder) {
    return ib101b34e7159adcd137c1e833f15a82941a098e6363bda2ae169edaeb7cc77e0.NewWorkbookRangeViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
