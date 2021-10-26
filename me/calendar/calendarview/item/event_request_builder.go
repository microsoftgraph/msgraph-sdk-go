package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i112b1ed50d98e06e2afc6c27777a4413e4aec03373c2c045d8fab1d827230449 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/dismissreminder"
    i1bcbde3648b969716e475c91cdb32d8cdf62c76e85838d834083a236f2d38333 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/tentativelyaccept"
    i3c35eb67b420d24b8b792bd486949e59c9d8b4b7fb3f16f018266a3064470c46 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/snoozereminder"
    i4429fd37a56c4f85306e5b908f5976bb1569387d88606fdbc734b5b5afff1a5e "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/multivalueextendedproperties"
    i4b3ac3d77873723261c8c94dc247eb0c4e33b3adc4a4b2528b3a8fc513f95250 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/cancel"
    i651ebf3ec997ed2debe308c2e0ed3a208f869f2a9c82390f706b6456e906e255 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances"
    i719ba32a86c43f7802b551a6d9320eb0b3e90c285148770a88c044f42b79b024 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/attachments"
    i7ded4b752eb3655e749d50f9eafabc0cbea03fdebdfb187b26ddd1566a7c4236 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/accept"
    i965695020684e888087236e60e5c0f7971367bc349f4a0088bd1af20ca363f05 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/forward"
    ia1eb05bd3c4f8481e8d0ac2eb5ff47aa090fb98c35c567df6e13064f70d0b3d0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/calendar"
    ia39bb4166c4ef2deee03aec7af70c9763aea4f1fc2b921860014e829fd638086 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/singlevalueextendedproperties"
    iece4056a8d35809aba6e4ff7ae53718558f0501f3bf175b276feb26afb59a056 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/extensions"
    iffae5ee2a2d9a81008b555644a80fc76abfbbd265c837cecf1568216e5db2972 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/decline"
    i0df45b242d0b94600d9cbfc3f4e57b9acc807e751ed7030c7c5d4fb2846927eb "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item"
    i1af302b703d5ff4044ea6d579861399d521115e720901fd67e398c46ac2c4dde "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/singlevalueextendedproperties/item"
    i20495d21133b26c7848a42901db05d39645c739c0ef4dfc27137db115c7ff68a "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/multivalueextendedproperties/item"
    i536a7b87db2cf3b5a9fc821e79fa11d85cd220fa19be5513d62932141a446474 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/extensions/item"
    i55f8e56a76f6d157ef3775e097b2a4e244e70c88edea43c5a8ba28bf62638a17 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/attachments/item"
)

// Builds and executes requests for operations under \me\calendar\calendarView\{event-id}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The calendar view for the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select_escaped []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
}
func (m *EventRequestBuilder) Accept()(*i7ded4b752eb3655e749d50f9eafabc0cbea03fdebdfb187b26ddd1566a7c4236.AcceptRequestBuilder) {
    return i7ded4b752eb3655e749d50f9eafabc0cbea03fdebdfb187b26ddd1566a7c4236.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i719ba32a86c43f7802b551a6d9320eb0b3e90c285148770a88c044f42b79b024.AttachmentsRequestBuilder) {
    return i719ba32a86c43f7802b551a6d9320eb0b3e90c285148770a88c044f42b79b024.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) AttachmentsById(id string)(*i55f8e56a76f6d157ef3775e097b2a4e244e70c88edea43c5a8ba28bf62638a17.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i55f8e56a76f6d157ef3775e097b2a4e244e70c88edea43c5a8ba28bf62638a17.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ia1eb05bd3c4f8481e8d0ac2eb5ff47aa090fb98c35c567df6e13064f70d0b3d0.CalendarRequestBuilder) {
    return ia1eb05bd3c4f8481e8d0ac2eb5ff47aa090fb98c35c567df6e13064f70d0b3d0.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i4b3ac3d77873723261c8c94dc247eb0c4e33b3adc4a4b2528b3a8fc513f95250.CancelRequestBuilder) {
    return i4b3ac3d77873723261c8c94dc247eb0c4e33b3adc4a4b2528b3a8fc513f95250.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *EventRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *EventRequestBuilder) CreateGetRequestInformation(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EventRequestBuilderGetQueryParameters)
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
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *EventRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*iffae5ee2a2d9a81008b555644a80fc76abfbbd265c837cecf1568216e5db2972.DeclineRequestBuilder) {
    return iffae5ee2a2d9a81008b555644a80fc76abfbbd265c837cecf1568216e5db2972.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EventRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EventRequestBuilder) DismissReminder()(*i112b1ed50d98e06e2afc6c27777a4413e4aec03373c2c045d8fab1d827230449.DismissReminderRequestBuilder) {
    return i112b1ed50d98e06e2afc6c27777a4413e4aec03373c2c045d8fab1d827230449.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*iece4056a8d35809aba6e4ff7ae53718558f0501f3bf175b276feb26afb59a056.ExtensionsRequestBuilder) {
    return iece4056a8d35809aba6e4ff7ae53718558f0501f3bf175b276feb26afb59a056.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) ExtensionsById(id string)(*i536a7b87db2cf3b5a9fc821e79fa11d85cd220fa19be5513d62932141a446474.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i536a7b87db2cf3b5a9fc821e79fa11d85cd220fa19be5513d62932141a446474.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i965695020684e888087236e60e5c0f7971367bc349f4a0088bd1af20ca363f05.ForwardRequestBuilder) {
    return i965695020684e888087236e60e5c0f7971367bc349f4a0088bd1af20ca363f05.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EventRequestBuilder) Get(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event), nil
}
func (m *EventRequestBuilder) Instances()(*i651ebf3ec997ed2debe308c2e0ed3a208f869f2a9c82390f706b6456e906e255.InstancesRequestBuilder) {
    return i651ebf3ec997ed2debe308c2e0ed3a208f869f2a9c82390f706b6456e906e255.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.instances.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) InstancesById(id string)(*i0df45b242d0b94600d9cbfc3f4e57b9acc807e751ed7030c7c5d4fb2846927eb.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i0df45b242d0b94600d9cbfc3f4e57b9acc807e751ed7030c7c5d4fb2846927eb.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i4429fd37a56c4f85306e5b908f5976bb1569387d88606fdbc734b5b5afff1a5e.MultiValueExtendedPropertiesRequestBuilder) {
    return i4429fd37a56c4f85306e5b908f5976bb1569387d88606fdbc734b5b5afff1a5e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i20495d21133b26c7848a42901db05d39645c739c0ef4dfc27137db115c7ff68a.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i20495d21133b26c7848a42901db05d39645c739c0ef4dfc27137db115c7ff68a.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EventRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ia39bb4166c4ef2deee03aec7af70c9763aea4f1fc2b921860014e829fd638086.SingleValueExtendedPropertiesRequestBuilder) {
    return ia39bb4166c4ef2deee03aec7af70c9763aea4f1fc2b921860014e829fd638086.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i1af302b703d5ff4044ea6d579861399d521115e720901fd67e398c46ac2c4dde.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i1af302b703d5ff4044ea6d579861399d521115e720901fd67e398c46ac2c4dde.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i3c35eb67b420d24b8b792bd486949e59c9d8b4b7fb3f16f018266a3064470c46.SnoozeReminderRequestBuilder) {
    return i3c35eb67b420d24b8b792bd486949e59c9d8b4b7fb3f16f018266a3064470c46.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i1bcbde3648b969716e475c91cdb32d8cdf62c76e85838d834083a236f2d38333.TentativelyAcceptRequestBuilder) {
    return i1bcbde3648b969716e475c91cdb32d8cdf62c76e85838d834083a236f2d38333.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
