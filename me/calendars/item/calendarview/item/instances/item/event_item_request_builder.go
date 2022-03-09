package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i0602a4ae10e624862125c53d51d3d51c9dc5a37113946633bed40f17fb781f7b "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/forward"
    i6d64c7e1cca7ad8fa577ce030dfa985065e9ffdd2d1adff55f625ecfe68d809a "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/cancel"
    i7634593a5cb635e59e07827b667d70de01d548295db9c7af82abd4b58e164d56 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/tentativelyaccept"
    i930cb1687a03e45ac860635c11939e503974b5243752ace0ffaecd6e51d9dd43 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/snoozereminder"
    i9c798380a879e24c72eaf54ac5206abd7ccf2362d4fdb06a4933a681cd45da0d "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/attachments"
    ia22ee94a64e72e6ec840e6c4eea34785ee35a3c411d147d2c521b0c9326124d5 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/multivalueextendedproperties"
    ib629ac522c6923ff60373a8f5ede6f563ec8d71c473a242f88bb8f168c7b42aa "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/accept"
    id3f2160437522c02a65bb24b9f83a04ea304a5bb480fedccb76f6f9e05681509 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/decline"
    id77a3ac83496e8a3a7dee82926d1945f7859b14bd91428167f89a3ba2b8a2beb "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/calendar"
    ie0060af3b60784ea8319dc611552e25cd5f699a6a4cc0a2e641cf66b61f40cce "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties"
    ie8345daaac85ac976d1c97b0d0bbff706cadf85464ed0fafd18021450e23dfb0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/extensions"
    if187a802e3c892f44d6587a829fb53df9b27881331adc50d52ee57daed93d19e "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/dismissreminder"
    i58612f00344d8937539df4d0589e1ede29951a228bc6b43182b3bfd8475277ce "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/multivalueextendedproperties/item"
    i5cd06f5898e114f5d00c935bf07bb821eaef172311dc3e9f99c1b10e313752c9 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/extensions/item"
    i64a5c9c18c69228ae31df0fc07823e1bb8edb2fa5ce62aba30e6a8611d0b5e81 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/attachments/item"
    ia1c816b1919a9975c9f98367893a6c5b96b6b8a8ec392fd295e4a0046cc3f4d9 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*ib629ac522c6923ff60373a8f5ede6f563ec8d71c473a242f88bb8f168c7b42aa.AcceptRequestBuilder) {
    return ib629ac522c6923ff60373a8f5ede6f563ec8d71c473a242f88bb8f168c7b42aa.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i9c798380a879e24c72eaf54ac5206abd7ccf2362d4fdb06a4933a681cd45da0d.AttachmentsRequestBuilder) {
    return i9c798380a879e24c72eaf54ac5206abd7ccf2362d4fdb06a4933a681cd45da0d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i64a5c9c18c69228ae31df0fc07823e1bb8edb2fa5ce62aba30e6a8611d0b5e81.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i64a5c9c18c69228ae31df0fc07823e1bb8edb2fa5ce62aba30e6a8611d0b5e81.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*id77a3ac83496e8a3a7dee82926d1945f7859b14bd91428167f89a3ba2b8a2beb.CalendarRequestBuilder) {
    return id77a3ac83496e8a3a7dee82926d1945f7859b14bd91428167f89a3ba2b8a2beb.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i6d64c7e1cca7ad8fa577ce030dfa985065e9ffdd2d1adff55f625ecfe68d809a.CancelRequestBuilder) {
    return i6d64c7e1cca7ad8fa577ce030dfa985065e9ffdd2d1adff55f625ecfe68d809a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/calendarView/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*id3f2160437522c02a65bb24b9f83a04ea304a5bb480fedccb76f6f9e05681509.DeclineRequestBuilder) {
    return id3f2160437522c02a65bb24b9f83a04ea304a5bb480fedccb76f6f9e05681509.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*if187a802e3c892f44d6587a829fb53df9b27881331adc50d52ee57daed93d19e.DismissReminderRequestBuilder) {
    return if187a802e3c892f44d6587a829fb53df9b27881331adc50d52ee57daed93d19e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*ie8345daaac85ac976d1c97b0d0bbff706cadf85464ed0fafd18021450e23dfb0.ExtensionsRequestBuilder) {
    return ie8345daaac85ac976d1c97b0d0bbff706cadf85464ed0fafd18021450e23dfb0.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i5cd06f5898e114f5d00c935bf07bb821eaef172311dc3e9f99c1b10e313752c9.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i5cd06f5898e114f5d00c935bf07bb821eaef172311dc3e9f99c1b10e313752c9.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i0602a4ae10e624862125c53d51d3d51c9dc5a37113946633bed40f17fb781f7b.ForwardRequestBuilder) {
    return i0602a4ae10e624862125c53d51d3d51c9dc5a37113946633bed40f17fb781f7b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ia22ee94a64e72e6ec840e6c4eea34785ee35a3c411d147d2c521b0c9326124d5.MultiValueExtendedPropertiesRequestBuilder) {
    return ia22ee94a64e72e6ec840e6c4eea34785ee35a3c411d147d2c521b0c9326124d5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i58612f00344d8937539df4d0589e1ede29951a228bc6b43182b3bfd8475277ce.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i58612f00344d8937539df4d0589e1ede29951a228bc6b43182b3bfd8475277ce.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ie0060af3b60784ea8319dc611552e25cd5f699a6a4cc0a2e641cf66b61f40cce.SingleValueExtendedPropertiesRequestBuilder) {
    return ie0060af3b60784ea8319dc611552e25cd5f699a6a4cc0a2e641cf66b61f40cce.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia1c816b1919a9975c9f98367893a6c5b96b6b8a8ec392fd295e4a0046cc3f4d9.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ia1c816b1919a9975c9f98367893a6c5b96b6b8a8ec392fd295e4a0046cc3f4d9.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i930cb1687a03e45ac860635c11939e503974b5243752ace0ffaecd6e51d9dd43.SnoozeReminderRequestBuilder) {
    return i930cb1687a03e45ac860635c11939e503974b5243752ace0ffaecd6e51d9dd43.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i7634593a5cb635e59e07827b667d70de01d548295db9c7af82abd4b58e164d56.TentativelyAcceptRequestBuilder) {
    return i7634593a5cb635e59e07827b667d70de01d548295db9c7af82abd4b58e164d56.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
