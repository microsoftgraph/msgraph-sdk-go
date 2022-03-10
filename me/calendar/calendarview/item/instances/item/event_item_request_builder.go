package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i2981f58c7f06c0386f9a4e727317aac9d4c17f231776484a8690114b12a85e5e "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/dismissreminder"
    i3317acf916015a6da98992a383137812a5961e5c81e803d22522b2484c48033d "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/forward"
    i38436f0a4e92197268eee818a3e5019cb401034642c8e29cda6af642314b18ce "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/decline"
    i3a260f80b9a8a1f84911d79d4cdcf7bc1d18069773a7f6200a5f6a416868ab04 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/extensions"
    i3e7e30189bb9604f17498df3ab4b1e61dc383ec2de779cbb344374170c352485 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/snoozereminder"
    i5e3379b62cbc68e92960e6e91cd5dc54475bf7d326ab6332654cc7d7fdebca1e "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/cancel"
    i73a38c0e9446fa2d7f78ea0d1b6d29befa145927b44439b28f9d2f22a32bcff7 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/attachments"
    i857d339808a1eb72877ecb77512cc129ce7eb1d489c756636e0d8640404d5eac "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/multivalueextendedproperties"
    id2c4b5c449138e176f6eb07afd0ed786b9fec3efba8c9c0dc5c7128b53d7043a "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/accept"
    ie01b50d6f430eb576e8d08605a1c6be014f5239a0ad0a26f95521f85fb8de7e5 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/calendar"
    if3b11f993193c9a0f6e47f8102337e61cb11c136e848eac16c9fa4a1805915dc "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/singlevalueextendedproperties"
    if5e5d38f1d503ab7b83c4d2fdf8c74d829bd0dd8de7f8a8cdd7b612a99a2b9a2 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/tentativelyaccept"
    i07201606642086d089115851b2b58182e0b540badcccd5a11c61214c00edbf81 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/extensions/item"
    i3e804013a81aecff2c19c493c5d680862564cbd38a48d47aeac83d3712b0f930 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/singlevalueextendedproperties/item"
    ie56b9658d8470399c10c76bcf6ba4e2c225096cc3580b7ae93778faa0c4c3153 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/multivalueextendedproperties/item"
    if4a028fefbf5e7f4c6969694da62a99987871155bdaae6ae5f256353567dba98 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item/instances/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*id2c4b5c449138e176f6eb07afd0ed786b9fec3efba8c9c0dc5c7128b53d7043a.AcceptRequestBuilder) {
    return id2c4b5c449138e176f6eb07afd0ed786b9fec3efba8c9c0dc5c7128b53d7043a.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i73a38c0e9446fa2d7f78ea0d1b6d29befa145927b44439b28f9d2f22a32bcff7.AttachmentsRequestBuilder) {
    return i73a38c0e9446fa2d7f78ea0d1b6d29befa145927b44439b28f9d2f22a32bcff7.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*if4a028fefbf5e7f4c6969694da62a99987871155bdaae6ae5f256353567dba98.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return if4a028fefbf5e7f4c6969694da62a99987871155bdaae6ae5f256353567dba98.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ie01b50d6f430eb576e8d08605a1c6be014f5239a0ad0a26f95521f85fb8de7e5.CalendarRequestBuilder) {
    return ie01b50d6f430eb576e8d08605a1c6be014f5239a0ad0a26f95521f85fb8de7e5.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i5e3379b62cbc68e92960e6e91cd5dc54475bf7d326ab6332654cc7d7fdebca1e.CancelRequestBuilder) {
    return i5e3379b62cbc68e92960e6e91cd5dc54475bf7d326ab6332654cc7d7fdebca1e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event_id}/instances/{event_id1}{?select}";
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
// CreateDeleteRequestInformation delete navigation property instances for me
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
// CreatePatchRequestInformation update the navigation property instances in me
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
func (m *EventItemRequestBuilder) Decline()(*i38436f0a4e92197268eee818a3e5019cb401034642c8e29cda6af642314b18ce.DeclineRequestBuilder) {
    return i38436f0a4e92197268eee818a3e5019cb401034642c8e29cda6af642314b18ce.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i2981f58c7f06c0386f9a4e727317aac9d4c17f231776484a8690114b12a85e5e.DismissReminderRequestBuilder) {
    return i2981f58c7f06c0386f9a4e727317aac9d4c17f231776484a8690114b12a85e5e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i3a260f80b9a8a1f84911d79d4cdcf7bc1d18069773a7f6200a5f6a416868ab04.ExtensionsRequestBuilder) {
    return i3a260f80b9a8a1f84911d79d4cdcf7bc1d18069773a7f6200a5f6a416868ab04.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i07201606642086d089115851b2b58182e0b540badcccd5a11c61214c00edbf81.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i07201606642086d089115851b2b58182e0b540badcccd5a11c61214c00edbf81.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i3317acf916015a6da98992a383137812a5961e5c81e803d22522b2484c48033d.ForwardRequestBuilder) {
    return i3317acf916015a6da98992a383137812a5961e5c81e803d22522b2484c48033d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i857d339808a1eb72877ecb77512cc129ce7eb1d489c756636e0d8640404d5eac.MultiValueExtendedPropertiesRequestBuilder) {
    return i857d339808a1eb72877ecb77512cc129ce7eb1d489c756636e0d8640404d5eac.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie56b9658d8470399c10c76bcf6ba4e2c225096cc3580b7ae93778faa0c4c3153.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ie56b9658d8470399c10c76bcf6ba4e2c225096cc3580b7ae93778faa0c4c3153.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*if3b11f993193c9a0f6e47f8102337e61cb11c136e848eac16c9fa4a1805915dc.SingleValueExtendedPropertiesRequestBuilder) {
    return if3b11f993193c9a0f6e47f8102337e61cb11c136e848eac16c9fa4a1805915dc.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i3e804013a81aecff2c19c493c5d680862564cbd38a48d47aeac83d3712b0f930.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i3e804013a81aecff2c19c493c5d680862564cbd38a48d47aeac83d3712b0f930.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i3e7e30189bb9604f17498df3ab4b1e61dc383ec2de779cbb344374170c352485.SnoozeReminderRequestBuilder) {
    return i3e7e30189bb9604f17498df3ab4b1e61dc383ec2de779cbb344374170c352485.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*if5e5d38f1d503ab7b83c4d2fdf8c74d829bd0dd8de7f8a8cdd7b612a99a2b9a2.TentativelyAcceptRequestBuilder) {
    return if5e5d38f1d503ab7b83c4d2fdf8c74d829bd0dd8de7f8a8cdd7b612a99a2b9a2.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
