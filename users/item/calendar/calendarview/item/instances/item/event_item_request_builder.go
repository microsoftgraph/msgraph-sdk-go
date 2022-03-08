package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0a1772f845f9085861aa090a493458a98894090017e54a50f4c5b88de6644f90 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/singlevalueextendedproperties"
    i12642507ce20354eb16e1d03db64db137eee76ee94c176fcde74028377188b9f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/tentativelyaccept"
    i267fa55853fc556031a9e8954e6d6ca7189170e57c14c5e9becdde2fc1a25509 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/forward"
    i2db7a98d47f916ba53c004f1824aef73f2e99fd54aa97c680370fe090f7e39ad "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/extensions"
    i32a121e8104eb28a076666d209446bbafc37ad8d98c7a39c5f3bec7e5b981a17 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/decline"
    i4aff33fe86c5b97a31c51b5cc8f9b25789985c6f10f08b5f1b8ef661e90c5631 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/attachments"
    i51e9808b3bbcc42c101ad287b0e4d35e9d3d6e9e98cdb407094e22377d4a5c60 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/dismissreminder"
    i5750fd2e6ce08fd287cd06791c29fcb169a7c93950f74de254caae9571cf1627 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/calendar"
    i611d094d48a3d02bdba0eb293c70d1357016e16620d57617b01667f9128608f9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/multivalueextendedproperties"
    i74a6f8db4dc95067036c531df6f30ff434c1c7836b5225ed2fac1b9c6f36627b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/snoozereminder"
    i772f4b3da338df2e528317bf41b414aae3b767a697ec743d0992e790dffc7192 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/accept"
    ibda755edec3f52d7eb9620d5e348523a4620dd11e6f74a5359583a040965cd29 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/cancel"
    i7e15d65fb8896ad37f5d5bafc592ec1b831d58049ce192cc0743eabb2e44c602 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/multivalueextendedproperties/item"
    i854847a82f4ba11c9ef275cfd1248e266e65cad2a886be229151764e47b5704e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/attachments/item"
    ic03d56835c6ce122a368ce4889338baf6dc5aea7daa3939cc717f452afcdfba8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/extensions/item"
    ica43d1587570c8b62a74eb2fa6ad4efe2a21fa8bb7d34c1336464e09ef54f656 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
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
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
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
func (m *EventItemRequestBuilder) Accept()(*i772f4b3da338df2e528317bf41b414aae3b767a697ec743d0992e790dffc7192.AcceptRequestBuilder) {
    return i772f4b3da338df2e528317bf41b414aae3b767a697ec743d0992e790dffc7192.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i4aff33fe86c5b97a31c51b5cc8f9b25789985c6f10f08b5f1b8ef661e90c5631.AttachmentsRequestBuilder) {
    return i4aff33fe86c5b97a31c51b5cc8f9b25789985c6f10f08b5f1b8ef661e90c5631.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i854847a82f4ba11c9ef275cfd1248e266e65cad2a886be229151764e47b5704e.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i854847a82f4ba11c9ef275cfd1248e266e65cad2a886be229151764e47b5704e.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i5750fd2e6ce08fd287cd06791c29fcb169a7c93950f74de254caae9571cf1627.CalendarRequestBuilder) {
    return i5750fd2e6ce08fd287cd06791c29fcb169a7c93950f74de254caae9571cf1627.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ibda755edec3f52d7eb9620d5e348523a4620dd11e6f74a5359583a040965cd29.CancelRequestBuilder) {
    return ibda755edec3f52d7eb9620d5e348523a4620dd11e6f74a5359583a040965cd29.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/calendarView/{event_id}/instances/{event_id1}{?select}";
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
// CreateDeleteRequestInformation delete navigation property instances for users
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property instances in users
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
func (m *EventItemRequestBuilder) Decline()(*i32a121e8104eb28a076666d209446bbafc37ad8d98c7a39c5f3bec7e5b981a17.DeclineRequestBuilder) {
    return i32a121e8104eb28a076666d209446bbafc37ad8d98c7a39c5f3bec7e5b981a17.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i51e9808b3bbcc42c101ad287b0e4d35e9d3d6e9e98cdb407094e22377d4a5c60.DismissReminderRequestBuilder) {
    return i51e9808b3bbcc42c101ad287b0e4d35e9d3d6e9e98cdb407094e22377d4a5c60.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i2db7a98d47f916ba53c004f1824aef73f2e99fd54aa97c680370fe090f7e39ad.ExtensionsRequestBuilder) {
    return i2db7a98d47f916ba53c004f1824aef73f2e99fd54aa97c680370fe090f7e39ad.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic03d56835c6ce122a368ce4889338baf6dc5aea7daa3939cc717f452afcdfba8.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ic03d56835c6ce122a368ce4889338baf6dc5aea7daa3939cc717f452afcdfba8.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i267fa55853fc556031a9e8954e6d6ca7189170e57c14c5e9becdde2fc1a25509.ForwardRequestBuilder) {
    return i267fa55853fc556031a9e8954e6d6ca7189170e57c14c5e9becdde2fc1a25509.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i611d094d48a3d02bdba0eb293c70d1357016e16620d57617b01667f9128608f9.MultiValueExtendedPropertiesRequestBuilder) {
    return i611d094d48a3d02bdba0eb293c70d1357016e16620d57617b01667f9128608f9.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7e15d65fb8896ad37f5d5bafc592ec1b831d58049ce192cc0743eabb2e44c602.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i7e15d65fb8896ad37f5d5bafc592ec1b831d58049ce192cc0743eabb2e44c602.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i0a1772f845f9085861aa090a493458a98894090017e54a50f4c5b88de6644f90.SingleValueExtendedPropertiesRequestBuilder) {
    return i0a1772f845f9085861aa090a493458a98894090017e54a50f4c5b88de6644f90.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ica43d1587570c8b62a74eb2fa6ad4efe2a21fa8bb7d34c1336464e09ef54f656.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ica43d1587570c8b62a74eb2fa6ad4efe2a21fa8bb7d34c1336464e09ef54f656.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i74a6f8db4dc95067036c531df6f30ff434c1c7836b5225ed2fac1b9c6f36627b.SnoozeReminderRequestBuilder) {
    return i74a6f8db4dc95067036c531df6f30ff434c1c7836b5225ed2fac1b9c6f36627b.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i12642507ce20354eb16e1d03db64db137eee76ee94c176fcde74028377188b9f.TentativelyAcceptRequestBuilder) {
    return i12642507ce20354eb16e1d03db64db137eee76ee94c176fcde74028377188b9f.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
