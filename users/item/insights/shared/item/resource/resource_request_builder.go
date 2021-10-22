package resource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0760a086d672db9e1f9badd4568fe68a0f937e8b9800b0a28bdec7daa64d5ccd "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/ref"
    i34d062f7c1f9b651ed7e321f7b6727ce6d621992085db8d45c80fbc916188dcc "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/managedappprotection"
    i5f98868acbc862431e923374a6ba8e936fb3eec70eed687c4ab204b8f1263b09 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/printdocument"
    i70f104074c4620fcf4c3f78aeafd5d45b974329cb2b41fbd3bb6e04d67f12985 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/targetedmanagedappprotection"
    i80e2ebfc41b0625dcbf2373eb1a53b5d52caccb377f6583381cf8f337986f05b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrangeview"
    i852cb7dda78fe8aa1adb8c09e92988de16d67c84ffb56b0e80172e3123b4a7e5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrangefill"
    i89ca78a87910a4d326eebc0ff244cb6c019d297c07a94b543d3173c95db7ecd2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/calendarsharingmessage"
    i90cf4f842524a7ba0d541f2cbf2b82801f174573af43720ebaec35e9e485ff4c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/mobileappcontentfile"
    ic993f2af526ca3aaa2c9130990be0e147be1fe3f7bbbe91252074ffdeaae61e5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrangesort"
    icaa18da9fb901891c3f568bd8f939aba235669e5282c8760a7b77dd52a05aeee "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/schedulechangerequest"
    ida50a4c71903ec124c290d1aef807b13713b9abe73d81f9f9ea03776fd6cff06 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/printjob"
    ie01efc200274d25bd47eff13ceb7c60cf00d2c124bce192aa2e6c0a30f1f0cb6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrangeformat"
    ie833ef9de6455e7dfa16db8692a2538511702bbbf4e3dce6dafe54ad72f7a1fd "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange"
    iea260140d987ad9e7e464b0abbd7f734a5df4ffb5180523a47049b434461809b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/windowsinformationprotection"
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
func (m *ResourceRequestBuilder) CalendarSharingMessage()(*i89ca78a87910a4d326eebc0ff244cb6c019d297c07a94b543d3173c95db7ecd2.CalendarSharingMessageRequestBuilder) {
    return i89ca78a87910a4d326eebc0ff244cb6c019d297c07a94b543d3173c95db7ecd2.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceRequestBuilder) {
    m := &ResourceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/resource{?select,expand}";
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
func (m *ResourceRequestBuilder) ManagedAppProtection()(*i34d062f7c1f9b651ed7e321f7b6727ce6d621992085db8d45c80fbc916188dcc.ManagedAppProtectionRequestBuilder) {
    return i34d062f7c1f9b651ed7e321f7b6727ce6d621992085db8d45c80fbc916188dcc.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) MobileAppContentFile()(*i90cf4f842524a7ba0d541f2cbf2b82801f174573af43720ebaec35e9e485ff4c.MobileAppContentFileRequestBuilder) {
    return i90cf4f842524a7ba0d541f2cbf2b82801f174573af43720ebaec35e9e485ff4c.NewMobileAppContentFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintDocument()(*i5f98868acbc862431e923374a6ba8e936fb3eec70eed687c4ab204b8f1263b09.PrintDocumentRequestBuilder) {
    return i5f98868acbc862431e923374a6ba8e936fb3eec70eed687c4ab204b8f1263b09.NewPrintDocumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) PrintJob()(*ida50a4c71903ec124c290d1aef807b13713b9abe73d81f9f9ea03776fd6cff06.PrintJobRequestBuilder) {
    return ida50a4c71903ec124c290d1aef807b13713b9abe73d81f9f9ea03776fd6cff06.NewPrintJobRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) Ref()(*i0760a086d672db9e1f9badd4568fe68a0f937e8b9800b0a28bdec7daa64d5ccd.RefRequestBuilder) {
    return i0760a086d672db9e1f9badd4568fe68a0f937e8b9800b0a28bdec7daa64d5ccd.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) ScheduleChangeRequest()(*icaa18da9fb901891c3f568bd8f939aba235669e5282c8760a7b77dd52a05aeee.ScheduleChangeRequestRequestBuilder) {
    return icaa18da9fb901891c3f568bd8f939aba235669e5282c8760a7b77dd52a05aeee.NewScheduleChangeRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) TargetedManagedAppProtection()(*i70f104074c4620fcf4c3f78aeafd5d45b974329cb2b41fbd3bb6e04d67f12985.TargetedManagedAppProtectionRequestBuilder) {
    return i70f104074c4620fcf4c3f78aeafd5d45b974329cb2b41fbd3bb6e04d67f12985.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WindowsInformationProtection()(*iea260140d987ad9e7e464b0abbd7f734a5df4ffb5180523a47049b434461809b.WindowsInformationProtectionRequestBuilder) {
    return iea260140d987ad9e7e464b0abbd7f734a5df4ffb5180523a47049b434461809b.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRange()(*ie833ef9de6455e7dfa16db8692a2538511702bbbf4e3dce6dafe54ad72f7a1fd.WorkbookRangeRequestBuilder) {
    return ie833ef9de6455e7dfa16db8692a2538511702bbbf4e3dce6dafe54ad72f7a1fd.NewWorkbookRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFill()(*i852cb7dda78fe8aa1adb8c09e92988de16d67c84ffb56b0e80172e3123b4a7e5.WorkbookRangeFillRequestBuilder) {
    return i852cb7dda78fe8aa1adb8c09e92988de16d67c84ffb56b0e80172e3123b4a7e5.NewWorkbookRangeFillRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeFormat()(*ie01efc200274d25bd47eff13ceb7c60cf00d2c124bce192aa2e6c0a30f1f0cb6.WorkbookRangeFormatRequestBuilder) {
    return ie01efc200274d25bd47eff13ceb7c60cf00d2c124bce192aa2e6c0a30f1f0cb6.NewWorkbookRangeFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeSort()(*ic993f2af526ca3aaa2c9130990be0e147be1fe3f7bbbe91252074ffdeaae61e5.WorkbookRangeSortRequestBuilder) {
    return ic993f2af526ca3aaa2c9130990be0e147be1fe3f7bbbe91252074ffdeaae61e5.NewWorkbookRangeSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceRequestBuilder) WorkbookRangeView()(*i80e2ebfc41b0625dcbf2373eb1a53b5d52caccb377f6583381cf8f337986f05b.WorkbookRangeViewRequestBuilder) {
    return i80e2ebfc41b0625dcbf2373eb1a53b5d52caccb377f6583381cf8f337986f05b.NewWorkbookRangeViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
