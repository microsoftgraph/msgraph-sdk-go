package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2541624cc2b7c2f09531c622ffbb42f4a195d101ad0ad14960f1f7e657102dbc "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/extensions"
    i2a0da227fb026925fc8b27a16ecbf8cc55f0b9a76e5549734baa76307a0bb83e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/multivalueextendedproperties"
    i3d8ae6cb74af0972786738f46478813a75154a5fe06967bf203a2e0918eabee5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/dismissreminder"
    i4ea184a6331e61d7549b8b34404e661ecdeb5805dfd8c57e0c1b3b015dffc03c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/attachments"
    i5a19924c15817bb8517dee2b633e6419ee084968aef5fcaf92f58cd7172a31ed "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/decline"
    i73cc31439204381dc3d6dc91ac06813562311395172df52ccc39264157757675 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/snoozereminder"
    i76172f12b34f51d0a11592d73ad0199c16d4ef8ebf8706eb0c1a27110f6d1f45 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/accept"
    i8432ac922f3623c62732b850934ae00f8014ebaabc1bb287cef0cef6433702f1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    i8631ad4f49ade0f7640c89e9d89155294a2c76a8fb7649c43e851a03b4311aee "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/forward"
    i8a87a0c791b5a48fd3b0ec426dab5c995b07ba2ee236b6f6b8a734cc5bf6c13d "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/cancel"
    i8d260fb96002fb446c710037c6bf1fb37f6ab3fbde3346f2af2a637e926d6baf "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances"
    ia4474e31f9f6f53023bfbed80dd5f789939a7bbae9e7a395838fa9a4bae43623 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/calendar"
    ie4908e0e67bcb1af8d1ee415aaefee36f4683dd9af729a4ffdc661090e90eb73 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/tentativelyaccept"
    i0a374527704358b2f886240ad18b1b8556ada3a535e6ffcd229d15b9decf0a3e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/attachments/item"
    i2c1f1efe93a31892c62f2d27507ea8bc2b85dd0749dc109db8a9be7bc576746c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    i361917cefafc439600ed28778b8652cd53d5430c792eb398c7d41102efe1a81e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/extensions/item"
    ib6f4ea864b364faa0da40819efa0bbe6984fece1d20e43e11a87d621c451a0b0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item"
    id28d63750d1254314b44fe7205991f1cd9bc0de925b89d62436af5889d20a03a "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder builds and executes requests for operations under \users\{user-id}\calendars\{calendar-id}\calendarView\{event-id}
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
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventItemRequestBuilder) Accept()(*i76172f12b34f51d0a11592d73ad0199c16d4ef8ebf8706eb0c1a27110f6d1f45.AcceptRequestBuilder) {
    return i76172f12b34f51d0a11592d73ad0199c16d4ef8ebf8706eb0c1a27110f6d1f45.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i4ea184a6331e61d7549b8b34404e661ecdeb5805dfd8c57e0c1b3b015dffc03c.AttachmentsRequestBuilder) {
    return i4ea184a6331e61d7549b8b34404e661ecdeb5805dfd8c57e0c1b3b015dffc03c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i0a374527704358b2f886240ad18b1b8556ada3a535e6ffcd229d15b9decf0a3e.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i0a374527704358b2f886240ad18b1b8556ada3a535e6ffcd229d15b9decf0a3e.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ia4474e31f9f6f53023bfbed80dd5f789939a7bbae9e7a395838fa9a4bae43623.CalendarRequestBuilder) {
    return ia4474e31f9f6f53023bfbed80dd5f789939a7bbae9e7a395838fa9a4bae43623.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i8a87a0c791b5a48fd3b0ec426dab5c995b07ba2ee236b6f6b8a734cc5bf6c13d.CancelRequestBuilder) {
    return i8a87a0c791b5a48fd3b0ec426dab5c995b07ba2ee236b6f6b8a734cc5bf6c13d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}/calendarView/{event_id}{?startDateTime,endDateTime,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Decline()(*i5a19924c15817bb8517dee2b633e6419ee084968aef5fcaf92f58cd7172a31ed.DeclineRequestBuilder) {
    return i5a19924c15817bb8517dee2b633e6419ee084968aef5fcaf92f58cd7172a31ed.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*i3d8ae6cb74af0972786738f46478813a75154a5fe06967bf203a2e0918eabee5.DismissReminderRequestBuilder) {
    return i3d8ae6cb74af0972786738f46478813a75154a5fe06967bf203a2e0918eabee5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i2541624cc2b7c2f09531c622ffbb42f4a195d101ad0ad14960f1f7e657102dbc.ExtensionsRequestBuilder) {
    return i2541624cc2b7c2f09531c622ffbb42f4a195d101ad0ad14960f1f7e657102dbc.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i361917cefafc439600ed28778b8652cd53d5430c792eb398c7d41102efe1a81e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i361917cefafc439600ed28778b8652cd53d5430c792eb398c7d41102efe1a81e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i8631ad4f49ade0f7640c89e9d89155294a2c76a8fb7649c43e851a03b4311aee.ForwardRequestBuilder) {
    return i8631ad4f49ade0f7640c89e9d89155294a2c76a8fb7649c43e851a03b4311aee.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event), nil
}
func (m *EventItemRequestBuilder) Instances()(*i8d260fb96002fb446c710037c6bf1fb37f6ab3fbde3346f2af2a637e926d6baf.InstancesRequestBuilder) {
    return i8d260fb96002fb446c710037c6bf1fb37f6ab3fbde3346f2af2a637e926d6baf.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ib6f4ea864b364faa0da40819efa0bbe6984fece1d20e43e11a87d621c451a0b0.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ib6f4ea864b364faa0da40819efa0bbe6984fece1d20e43e11a87d621c451a0b0.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i2a0da227fb026925fc8b27a16ecbf8cc55f0b9a76e5549734baa76307a0bb83e.MultiValueExtendedPropertiesRequestBuilder) {
    return i2a0da227fb026925fc8b27a16ecbf8cc55f0b9a76e5549734baa76307a0bb83e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2c1f1efe93a31892c62f2d27507ea8bc2b85dd0749dc109db8a9be7bc576746c.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i2c1f1efe93a31892c62f2d27507ea8bc2b85dd0749dc109db8a9be7bc576746c.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i8432ac922f3623c62732b850934ae00f8014ebaabc1bb287cef0cef6433702f1.SingleValueExtendedPropertiesRequestBuilder) {
    return i8432ac922f3623c62732b850934ae00f8014ebaabc1bb287cef0cef6433702f1.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id28d63750d1254314b44fe7205991f1cd9bc0de925b89d62436af5889d20a03a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return id28d63750d1254314b44fe7205991f1cd9bc0de925b89d62436af5889d20a03a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i73cc31439204381dc3d6dc91ac06813562311395172df52ccc39264157757675.SnoozeReminderRequestBuilder) {
    return i73cc31439204381dc3d6dc91ac06813562311395172df52ccc39264157757675.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ie4908e0e67bcb1af8d1ee415aaefee36f4683dd9af729a4ffdc661090e90eb73.TentativelyAcceptRequestBuilder) {
    return ie4908e0e67bcb1af8d1ee415aaefee36f4683dd9af729a4ffdc661090e90eb73.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
