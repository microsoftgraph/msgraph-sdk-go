package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*ib629ac522c6923ff60373a8f5ede6f563ec8d71c473a242f88bb8f168c7b42aa.AcceptRequestBuilder) {
    return ib629ac522c6923ff60373a8f5ede6f563ec8d71c473a242f88bb8f168c7b42aa.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
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
        urlTplParams["attachment%2Did"] = id
    }
    return i64a5c9c18c69228ae31df0fc07823e1bb8edb2fa5ce62aba30e6a8611d0b5e81.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*id77a3ac83496e8a3a7dee82926d1945f7859b14bd91428167f89a3ba2b8a2beb.CalendarRequestBuilder) {
    return id77a3ac83496e8a3a7dee82926d1945f7859b14bd91428167f89a3ba2b8a2beb.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i6d64c7e1cca7ad8fa577ce030dfa985065e9ffdd2d1adff55f625ecfe68d809a.CancelRequestBuilder) {
    return i6d64c7e1cca7ad8fa577ce030dfa985065e9ffdd2d1adff55f625ecfe68d809a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property instances for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property instances in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*id3f2160437522c02a65bb24b9f83a04ea304a5bb480fedccb76f6f9e05681509.DeclineRequestBuilder) {
    return id3f2160437522c02a65bb24b9f83a04ea304a5bb480fedccb76f6f9e05681509.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*if187a802e3c892f44d6587a829fb53df9b27881331adc50d52ee57daed93d19e.DismissReminderRequestBuilder) {
    return if187a802e3c892f44d6587a829fb53df9b27881331adc50d52ee57daed93d19e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
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
        urlTplParams["extension%2Did"] = id
    }
    return i5cd06f5898e114f5d00c935bf07bb821eaef172311dc3e9f99c1b10e313752c9.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i0602a4ae10e624862125c53d51d3d51c9dc5a37113946633bed40f17fb781f7b.ForwardRequestBuilder) {
    return i0602a4ae10e624862125c53d51d3d51c9dc5a37113946633bed40f17fb781f7b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
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
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i58612f00344d8937539df4d0589e1ede29951a228bc6b43182b3bfd8475277ce.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
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
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ia1c816b1919a9975c9f98367893a6c5b96b6b8a8ec392fd295e4a0046cc3f4d9.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i930cb1687a03e45ac860635c11939e503974b5243752ace0ffaecd6e51d9dd43.SnoozeReminderRequestBuilder) {
    return i930cb1687a03e45ac860635c11939e503974b5243752ace0ffaecd6e51d9dd43.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i7634593a5cb635e59e07827b667d70de01d548295db9c7af82abd4b58e164d56.TentativelyAcceptRequestBuilder) {
    return i7634593a5cb635e59e07827b667d70de01d548295db9c7af82abd4b58e164d56.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
