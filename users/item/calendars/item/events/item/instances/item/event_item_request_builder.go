package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i21350dd9c3ebf7c7980d92748cfad6c60c4f64d412a1db5393e484871a350e53 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/forward"
    i22c1077fc354cac1adead5c46b9ab01171e20eb9531abac5b947993de3bd2bcb "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/singlevalueextendedproperties"
    i34daf211b48600b2216e8e790110de9faa1ab5847fa19b5e410cd0ac6e877463 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/accept"
    i3b6422551344864f7f102c023b805905ad53a35a7a02e3fff045d0495d483794 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/calendar"
    i576a074e6f1faa80f3aadcb2c14a61945959d8ddea9c95473b680b1dc719de06 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/attachments"
    i8fdfef149ef003975eb75cecda7169ef00e36c589d1598046deaffe609a0379a "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/snoozereminder"
    i975df9cc8ff0e41089db66bb1a97dabd9a6aa59354310dd92f3ccda91f18ff63 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/tentativelyaccept"
    ib117e40a9ddce66f903053bf2be0f629e7291cc3d8b8f828c2f792076d29fde7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/cancel"
    ic127dae7bd0869437349a4ff127c98df169e7d0ef18228fabc35a3c4d375b734 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/dismissreminder"
    idd6eb09d6fc9cdb240da415ff91194979a0f6eb3834f18dcd6c6f880dbb26eec "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/decline"
    idffa32555eb230c560496ad15ad10e85d758ff0147728b3e42ed6f6a2db875ad "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/multivalueextendedproperties"
    if905c9963286459d58d38e42f29bc0c9b0f3e1e1a72ce015ca0545fdb1f6b80b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/extensions"
    ia12eeadab8fe41df793fa650415eef25a61e411f7a1eda0a67499e94cf7d9a1f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/multivalueextendedproperties/item"
    iea3b8e4e50ad8777b837123fd316901b0e47752fc4fd6f0e4ab08cb0e2f5bdba "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/extensions/item"
    if5b90af817e93df938ce09ae61c32eed8b2657669792ce7af99e3a4bc95aebb4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/singlevalueextendedproperties/item"
    ifb315e64f8f35196a92059a0850e587292399235989be2bc2575c78ab454764c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i34daf211b48600b2216e8e790110de9faa1ab5847fa19b5e410cd0ac6e877463.AcceptRequestBuilder) {
    return i34daf211b48600b2216e8e790110de9faa1ab5847fa19b5e410cd0ac6e877463.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i576a074e6f1faa80f3aadcb2c14a61945959d8ddea9c95473b680b1dc719de06.AttachmentsRequestBuilder) {
    return i576a074e6f1faa80f3aadcb2c14a61945959d8ddea9c95473b680b1dc719de06.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ifb315e64f8f35196a92059a0850e587292399235989be2bc2575c78ab454764c.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ifb315e64f8f35196a92059a0850e587292399235989be2bc2575c78ab454764c.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i3b6422551344864f7f102c023b805905ad53a35a7a02e3fff045d0495d483794.CalendarRequestBuilder) {
    return i3b6422551344864f7f102c023b805905ad53a35a7a02e3fff045d0495d483794.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ib117e40a9ddce66f903053bf2be0f629e7291cc3d8b8f828c2f792076d29fde7.CancelRequestBuilder) {
    return ib117e40a9ddce66f903053bf2be0f629e7291cc3d8b8f828c2f792076d29fde7.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}/events/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*idd6eb09d6fc9cdb240da415ff91194979a0f6eb3834f18dcd6c6f880dbb26eec.DeclineRequestBuilder) {
    return idd6eb09d6fc9cdb240da415ff91194979a0f6eb3834f18dcd6c6f880dbb26eec.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*ic127dae7bd0869437349a4ff127c98df169e7d0ef18228fabc35a3c4d375b734.DismissReminderRequestBuilder) {
    return ic127dae7bd0869437349a4ff127c98df169e7d0ef18228fabc35a3c4d375b734.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*if905c9963286459d58d38e42f29bc0c9b0f3e1e1a72ce015ca0545fdb1f6b80b.ExtensionsRequestBuilder) {
    return if905c9963286459d58d38e42f29bc0c9b0f3e1e1a72ce015ca0545fdb1f6b80b.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iea3b8e4e50ad8777b837123fd316901b0e47752fc4fd6f0e4ab08cb0e2f5bdba.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iea3b8e4e50ad8777b837123fd316901b0e47752fc4fd6f0e4ab08cb0e2f5bdba.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i21350dd9c3ebf7c7980d92748cfad6c60c4f64d412a1db5393e484871a350e53.ForwardRequestBuilder) {
    return i21350dd9c3ebf7c7980d92748cfad6c60c4f64d412a1db5393e484871a350e53.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable), nil
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*idffa32555eb230c560496ad15ad10e85d758ff0147728b3e42ed6f6a2db875ad.MultiValueExtendedPropertiesRequestBuilder) {
    return idffa32555eb230c560496ad15ad10e85d758ff0147728b3e42ed6f6a2db875ad.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ia12eeadab8fe41df793fa650415eef25a61e411f7a1eda0a67499e94cf7d9a1f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ia12eeadab8fe41df793fa650415eef25a61e411f7a1eda0a67499e94cf7d9a1f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i22c1077fc354cac1adead5c46b9ab01171e20eb9531abac5b947993de3bd2bcb.SingleValueExtendedPropertiesRequestBuilder) {
    return i22c1077fc354cac1adead5c46b9ab01171e20eb9531abac5b947993de3bd2bcb.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if5b90af817e93df938ce09ae61c32eed8b2657669792ce7af99e3a4bc95aebb4.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if5b90af817e93df938ce09ae61c32eed8b2657669792ce7af99e3a4bc95aebb4.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i8fdfef149ef003975eb75cecda7169ef00e36c589d1598046deaffe609a0379a.SnoozeReminderRequestBuilder) {
    return i8fdfef149ef003975eb75cecda7169ef00e36c589d1598046deaffe609a0379a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i975df9cc8ff0e41089db66bb1a97dabd9a6aa59354310dd92f3ccda91f18ff63.TentativelyAcceptRequestBuilder) {
    return i975df9cc8ff0e41089db66bb1a97dabd9a6aa59354310dd92f3ccda91f18ff63.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
