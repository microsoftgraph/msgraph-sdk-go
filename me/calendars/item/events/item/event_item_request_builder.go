package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i09aae1c2c441b5461fcc90aa35dcd1c4c2ca0552f429aa68c3db1c56bf0d2153 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/cancel"
    i27c7643dce2bb6c7a69be2cdf3257e7c4a990cf16f691aac9ce3596d45aad48f "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/forward"
    i383bbe46697a339b06b5bafed4a3590534909db04af55c794a8faec30e41ca1a "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/snoozereminder"
    i613d576972f3d8f4d501ac839ed6758ffca96792d87fe3cdc16bfbbcd999e339 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/dismissreminder"
    i72cf340f081c108008d2cadd6aaa88898cb9a7ec52a019ac18c447c5fc9e3cc4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/singlevalueextendedproperties"
    i882ced2bf5312be86e2c2ae6ea20a683cd710fdb806ed0768072570524785b86 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/decline"
    i8e5f4a41163bc4194e03167ef61b98c0e69797a47d44c62adb54062836a1bbe6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/multivalueextendedproperties"
    i9b6b878aac98fae97c8910a9408f847174ad681c9aa60585e78ec39998ab2778 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/attachments"
    iac14f0c5ccb2a5f3b9bafd14cee7eed24b0f2d2d6879f08fbd6269784857e360 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances"
    ib838f0df0ca92ac8dcade427c8a94d2f558237dc9cb00aa645b3cec1b9614bee "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/accept"
    ibd576f1728c4dc9fbe5962a0ee16ce5f68c7cfab3b2852ddf9c17f07ed387504 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/calendar"
    idb076a8a69d2a57d5dcdc849f6cbb856c0d28a0d737152e3dd6285793373a285 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/tentativelyaccept"
    idb61d479748362a6e089afd5c0b858f09487d50b611d7917271603f399cddcc3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/extensions"
    i072f4a99f49d2b8f6761bf43cfd5b2b95f06274c07ba44d9a737e87df17d7886 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/attachments/item"
    i157450aaee9c55248f9ba5ffbbf66bf16cc3d8d3d2bd8c9cfd4a765ad8bae272 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/multivalueextendedproperties/item"
    i4b79cd0561a1ab0685360c4682018160dedd8ecad4d5fd7dd32992b4017c80fd "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item"
    i5b791a9e8a95f008bb13124c648ebdef495fb58e993ff30cf6006693ece7f80a "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/singlevalueextendedproperties/item"
    ifd72b733fd36497510451fbda93ed7a025d727b889c80b685dc6fcdf606233c6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*ib838f0df0ca92ac8dcade427c8a94d2f558237dc9cb00aa645b3cec1b9614bee.AcceptRequestBuilder) {
    return ib838f0df0ca92ac8dcade427c8a94d2f558237dc9cb00aa645b3cec1b9614bee.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i9b6b878aac98fae97c8910a9408f847174ad681c9aa60585e78ec39998ab2778.AttachmentsRequestBuilder) {
    return i9b6b878aac98fae97c8910a9408f847174ad681c9aa60585e78ec39998ab2778.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i072f4a99f49d2b8f6761bf43cfd5b2b95f06274c07ba44d9a737e87df17d7886.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i072f4a99f49d2b8f6761bf43cfd5b2b95f06274c07ba44d9a737e87df17d7886.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ibd576f1728c4dc9fbe5962a0ee16ce5f68c7cfab3b2852ddf9c17f07ed387504.CalendarRequestBuilder) {
    return ibd576f1728c4dc9fbe5962a0ee16ce5f68c7cfab3b2852ddf9c17f07ed387504.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i09aae1c2c441b5461fcc90aa35dcd1c4c2ca0552f429aa68c3db1c56bf0d2153.CancelRequestBuilder) {
    return i09aae1c2c441b5461fcc90aa35dcd1c4c2ca0552f429aa68c3db1c56bf0d2153.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation delete navigation property events for me
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
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation update the navigation property events in me
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
func (m *EventItemRequestBuilder) Decline()(*i882ced2bf5312be86e2c2ae6ea20a683cd710fdb806ed0768072570524785b86.DeclineRequestBuilder) {
    return i882ced2bf5312be86e2c2ae6ea20a683cd710fdb806ed0768072570524785b86.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i613d576972f3d8f4d501ac839ed6758ffca96792d87fe3cdc16bfbbcd999e339.DismissReminderRequestBuilder) {
    return i613d576972f3d8f4d501ac839ed6758ffca96792d87fe3cdc16bfbbcd999e339.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*idb61d479748362a6e089afd5c0b858f09487d50b611d7917271603f399cddcc3.ExtensionsRequestBuilder) {
    return idb61d479748362a6e089afd5c0b858f09487d50b611d7917271603f399cddcc3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ifd72b733fd36497510451fbda93ed7a025d727b889c80b685dc6fcdf606233c6.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ifd72b733fd36497510451fbda93ed7a025d727b889c80b685dc6fcdf606233c6.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i27c7643dce2bb6c7a69be2cdf3257e7c4a990cf16f691aac9ce3596d45aad48f.ForwardRequestBuilder) {
    return i27c7643dce2bb6c7a69be2cdf3257e7c4a990cf16f691aac9ce3596d45aad48f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*iac14f0c5ccb2a5f3b9bafd14cee7eed24b0f2d2d6879f08fbd6269784857e360.InstancesRequestBuilder) {
    return iac14f0c5ccb2a5f3b9bafd14cee7eed24b0f2d2d6879f08fbd6269784857e360.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i4b79cd0561a1ab0685360c4682018160dedd8ecad4d5fd7dd32992b4017c80fd.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i4b79cd0561a1ab0685360c4682018160dedd8ecad4d5fd7dd32992b4017c80fd.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i8e5f4a41163bc4194e03167ef61b98c0e69797a47d44c62adb54062836a1bbe6.MultiValueExtendedPropertiesRequestBuilder) {
    return i8e5f4a41163bc4194e03167ef61b98c0e69797a47d44c62adb54062836a1bbe6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i157450aaee9c55248f9ba5ffbbf66bf16cc3d8d3d2bd8c9cfd4a765ad8bae272.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i157450aaee9c55248f9ba5ffbbf66bf16cc3d8d3d2bd8c9cfd4a765ad8bae272.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i72cf340f081c108008d2cadd6aaa88898cb9a7ec52a019ac18c447c5fc9e3cc4.SingleValueExtendedPropertiesRequestBuilder) {
    return i72cf340f081c108008d2cadd6aaa88898cb9a7ec52a019ac18c447c5fc9e3cc4.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5b791a9e8a95f008bb13124c648ebdef495fb58e993ff30cf6006693ece7f80a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i5b791a9e8a95f008bb13124c648ebdef495fb58e993ff30cf6006693ece7f80a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i383bbe46697a339b06b5bafed4a3590534909db04af55c794a8faec30e41ca1a.SnoozeReminderRequestBuilder) {
    return i383bbe46697a339b06b5bafed4a3590534909db04af55c794a8faec30e41ca1a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*idb076a8a69d2a57d5dcdc849f6cbb856c0d28a0d737152e3dd6285793373a285.TentativelyAcceptRequestBuilder) {
    return idb076a8a69d2a57d5dcdc849f6cbb856c0d28a0d737152e3dd6285793373a285.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
