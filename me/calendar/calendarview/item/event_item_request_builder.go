package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
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

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type EventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventItemRequestBuilder) Accept()(*i7ded4b752eb3655e749d50f9eafabc0cbea03fdebdfb187b26ddd1566a7c4236.AcceptRequestBuilder) {
    return i7ded4b752eb3655e749d50f9eafabc0cbea03fdebdfb187b26ddd1566a7c4236.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i719ba32a86c43f7802b551a6d9320eb0b3e90c285148770a88c044f42b79b024.AttachmentsRequestBuilder) {
    return i719ba32a86c43f7802b551a6d9320eb0b3e90c285148770a88c044f42b79b024.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i55f8e56a76f6d157ef3775e097b2a4e244e70c88edea43c5a8ba28bf62638a17.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i55f8e56a76f6d157ef3775e097b2a4e244e70c88edea43c5a8ba28bf62638a17.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ia1eb05bd3c4f8481e8d0ac2eb5ff47aa090fb98c35c567df6e13064f70d0b3d0.CalendarRequestBuilder) {
    return ia1eb05bd3c4f8481e8d0ac2eb5ff47aa090fb98c35c567df6e13064f70d0b3d0.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i4b3ac3d77873723261c8c94dc247eb0c4e33b3adc4a4b2528b3a8fc513f95250.CancelRequestBuilder) {
    return i4b3ac3d77873723261c8c94dc247eb0c4e33b3adc4a4b2528b3a8fc513f95250.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calendarView for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calendarView in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventItemRequestBuilder) Decline()(*iffae5ee2a2d9a81008b555644a80fc76abfbbd265c837cecf1568216e5db2972.DeclineRequestBuilder) {
    return iffae5ee2a2d9a81008b555644a80fc76abfbbd265c837cecf1568216e5db2972.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*i112b1ed50d98e06e2afc6c27777a4413e4aec03373c2c045d8fab1d827230449.DismissReminderRequestBuilder) {
    return i112b1ed50d98e06e2afc6c27777a4413e4aec03373c2c045d8fab1d827230449.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*iece4056a8d35809aba6e4ff7ae53718558f0501f3bf175b276feb26afb59a056.ExtensionsRequestBuilder) {
    return iece4056a8d35809aba6e4ff7ae53718558f0501f3bf175b276feb26afb59a056.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i536a7b87db2cf3b5a9fc821e79fa11d85cd220fa19be5513d62932141a446474.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i536a7b87db2cf3b5a9fc821e79fa11d85cd220fa19be5513d62932141a446474.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i965695020684e888087236e60e5c0f7971367bc349f4a0088bd1af20ca363f05.ForwardRequestBuilder) {
    return i965695020684e888087236e60e5c0f7971367bc349f4a0088bd1af20ca363f05.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable), nil
}
func (m *EventItemRequestBuilder) Instances()(*i651ebf3ec997ed2debe308c2e0ed3a208f869f2a9c82390f706b6456e906e255.InstancesRequestBuilder) {
    return i651ebf3ec997ed2debe308c2e0ed3a208f869f2a9c82390f706b6456e906e255.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i0df45b242d0b94600d9cbfc3f4e57b9acc807e751ed7030c7c5d4fb2846927eb.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i0df45b242d0b94600d9cbfc3f4e57b9acc807e751ed7030c7c5d4fb2846927eb.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i4429fd37a56c4f85306e5b908f5976bb1569387d88606fdbc734b5b5afff1a5e.MultiValueExtendedPropertiesRequestBuilder) {
    return i4429fd37a56c4f85306e5b908f5976bb1569387d88606fdbc734b5b5afff1a5e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i20495d21133b26c7848a42901db05d39645c739c0ef4dfc27137db115c7ff68a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i20495d21133b26c7848a42901db05d39645c739c0ef4dfc27137db115c7ff68a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia39bb4166c4ef2deee03aec7af70c9763aea4f1fc2b921860014e829fd638086.SingleValueExtendedPropertiesRequestBuilder) {
    return ia39bb4166c4ef2deee03aec7af70c9763aea4f1fc2b921860014e829fd638086.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i1af302b703d5ff4044ea6d579861399d521115e720901fd67e398c46ac2c4dde.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i1af302b703d5ff4044ea6d579861399d521115e720901fd67e398c46ac2c4dde.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i3c35eb67b420d24b8b792bd486949e59c9d8b4b7fb3f16f018266a3064470c46.SnoozeReminderRequestBuilder) {
    return i3c35eb67b420d24b8b792bd486949e59c9d8b4b7fb3f16f018266a3064470c46.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i1bcbde3648b969716e475c91cdb32d8cdf62c76e85838d834083a236f2d38333.TentativelyAcceptRequestBuilder) {
    return i1bcbde3648b969716e475c91cdb32d8cdf62c76e85838d834083a236f2d38333.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
