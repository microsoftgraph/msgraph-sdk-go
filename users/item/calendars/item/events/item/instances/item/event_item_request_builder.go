package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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
func (m *EventItemRequestBuilder) Accept()(*i34daf211b48600b2216e8e790110de9faa1ab5847fa19b5e410cd0ac6e877463.AcceptRequestBuilder) {
    return i34daf211b48600b2216e8e790110de9faa1ab5847fa19b5e410cd0ac6e877463.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
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
        urlTplParams["attachment%2Did"] = id
    }
    return ifb315e64f8f35196a92059a0850e587292399235989be2bc2575c78ab454764c.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i3b6422551344864f7f102c023b805905ad53a35a7a02e3fff045d0495d483794.CalendarRequestBuilder) {
    return i3b6422551344864f7f102c023b805905ad53a35a7a02e3fff045d0495d483794.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ib117e40a9ddce66f903053bf2be0f629e7291cc3d8b8f828c2f792076d29fde7.CancelRequestBuilder) {
    return ib117e40a9ddce66f903053bf2be0f629e7291cc3d8b8f828c2f792076d29fde7.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property instances for users
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
// CreatePatchRequestInformation update the navigation property instances in users
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
func (m *EventItemRequestBuilder) Decline()(*idd6eb09d6fc9cdb240da415ff91194979a0f6eb3834f18dcd6c6f880dbb26eec.DeclineRequestBuilder) {
    return idd6eb09d6fc9cdb240da415ff91194979a0f6eb3834f18dcd6c6f880dbb26eec.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*ic127dae7bd0869437349a4ff127c98df169e7d0ef18228fabc35a3c4d375b734.DismissReminderRequestBuilder) {
    return ic127dae7bd0869437349a4ff127c98df169e7d0ef18228fabc35a3c4d375b734.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
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
        urlTplParams["extension%2Did"] = id
    }
    return iea3b8e4e50ad8777b837123fd316901b0e47752fc4fd6f0e4ab08cb0e2f5bdba.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i21350dd9c3ebf7c7980d92748cfad6c60c4f64d412a1db5393e484871a350e53.ForwardRequestBuilder) {
    return i21350dd9c3ebf7c7980d92748cfad6c60c4f64d412a1db5393e484871a350e53.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ia12eeadab8fe41df793fa650415eef25a61e411f7a1eda0a67499e94cf7d9a1f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
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
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return if5b90af817e93df938ce09ae61c32eed8b2657669792ce7af99e3a4bc95aebb4.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i8fdfef149ef003975eb75cecda7169ef00e36c589d1598046deaffe609a0379a.SnoozeReminderRequestBuilder) {
    return i8fdfef149ef003975eb75cecda7169ef00e36c589d1598046deaffe609a0379a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i975df9cc8ff0e41089db66bb1a97dabd9a6aa59354310dd92f3ccda91f18ff63.TentativelyAcceptRequestBuilder) {
    return i975df9cc8ff0e41089db66bb1a97dabd9a6aa59354310dd92f3ccda91f18ff63.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
