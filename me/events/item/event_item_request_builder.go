package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1426802efbd0f5039b4ed42db7c3a88b7863b74e5a46d0d4bc134d1e85fbd90c "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/accept"
    i26b1c7cc63b87d2dd7618405741954826378e5687e20fdad5c9aa35ed448f8ff "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/decline"
    i2badd6557828e3ec37a7444fdf50ab5f7b8fdb7ae4b2a7ee749da38e2e461018 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/dismissreminder"
    i4209b20d5ab4261e9fa52d2682ea4586281037bc7cc1dfad2dcf1191b7d6dbc2 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar"
    i44038414241e1d1ac04a807191ed07589543d0a83a4e70fb548ebd312bf47055 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/attachments"
    i54078d1892815baae8a8fd0cece8461711b63ba916021bd5ced1c30195a1a698 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/multivalueextendedproperties"
    i5726cc676e11fe9fa90959a9a58ddf2f16c34f2a86c819a81f6f2ce4937b710a "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/snoozereminder"
    i774cf8a56cdf6984d443e1380e0d0dcd2ca8d118dde0425d0b60e04c5f1c436c "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/tentativelyaccept"
    i95058b33c206d566c03da284448a36f4ab33704d456ac0a7ebb2cd7b6cd2fb0f "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/extensions"
    ib59086c56b738f1f0717fb1ceaf00dffeac922ef570d9103a1ff449f48fca45d "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances"
    id5711ae3f0bfb22fb24dc70cb4be78fa08cb4a509118586258c0f33adb016cf2 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/singlevalueextendedproperties"
    idcc656e3c5d595c7e17926a2796ec2029a16b9df71a54b6dd2ced6343b191c8d "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/forward"
    ie36ea007a91599e24d84e7f363e1abf4b6e65cf46bfa7af40b5e30187cfa1602 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/cancel"
    i2a46836ccae42574e0c03c77e8963f70e5d011a92ca307f636d258be89ac1916 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/multivalueextendedproperties/item"
    i4cdf65b58617609797304bf46ce9451a705dd1606550f29e20aa67daf8236b42 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item"
    i8414fee0ad5be4f22aa3ea04ac1422e9124533bafff06b8fcddc21e0c7ca538a "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/extensions/item"
    ia9d78ab81a6804e22ae5ba33ac5f181e55a589f99d2e4208a5aa8aa0856893e1 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/attachments/item"
    ic43a8f89b11e7c86745f7bafd213b5d40252e1081389e9965812ca0da060a016 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.user entity.
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
// EventItemRequestBuilderGetQueryParameters the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Accept()(*i1426802efbd0f5039b4ed42db7c3a88b7863b74e5a46d0d4bc134d1e85fbd90c.AcceptRequestBuilder) {
    return i1426802efbd0f5039b4ed42db7c3a88b7863b74e5a46d0d4bc134d1e85fbd90c.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i44038414241e1d1ac04a807191ed07589543d0a83a4e70fb548ebd312bf47055.AttachmentsRequestBuilder) {
    return i44038414241e1d1ac04a807191ed07589543d0a83a4e70fb548ebd312bf47055.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia9d78ab81a6804e22ae5ba33ac5f181e55a589f99d2e4208a5aa8aa0856893e1.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ia9d78ab81a6804e22ae5ba33ac5f181e55a589f99d2e4208a5aa8aa0856893e1.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i4209b20d5ab4261e9fa52d2682ea4586281037bc7cc1dfad2dcf1191b7d6dbc2.CalendarRequestBuilder) {
    return i4209b20d5ab4261e9fa52d2682ea4586281037bc7cc1dfad2dcf1191b7d6dbc2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ie36ea007a91599e24d84e7f363e1abf4b6e65cf46bfa7af40b5e30187cfa1602.CancelRequestBuilder) {
    return ie36ea007a91599e24d84e7f363e1abf4b6e65cf46bfa7af40b5e30187cfa1602.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event%2Did}{?%24select}";
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
// CreateGetRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Decline()(*i26b1c7cc63b87d2dd7618405741954826378e5687e20fdad5c9aa35ed448f8ff.DeclineRequestBuilder) {
    return i26b1c7cc63b87d2dd7618405741954826378e5687e20fdad5c9aa35ed448f8ff.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i2badd6557828e3ec37a7444fdf50ab5f7b8fdb7ae4b2a7ee749da38e2e461018.DismissReminderRequestBuilder) {
    return i2badd6557828e3ec37a7444fdf50ab5f7b8fdb7ae4b2a7ee749da38e2e461018.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i95058b33c206d566c03da284448a36f4ab33704d456ac0a7ebb2cd7b6cd2fb0f.ExtensionsRequestBuilder) {
    return i95058b33c206d566c03da284448a36f4ab33704d456ac0a7ebb2cd7b6cd2fb0f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8414fee0ad5be4f22aa3ea04ac1422e9124533bafff06b8fcddc21e0c7ca538a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8414fee0ad5be4f22aa3ea04ac1422e9124533bafff06b8fcddc21e0c7ca538a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*idcc656e3c5d595c7e17926a2796ec2029a16b9df71a54b6dd2ced6343b191c8d.ForwardRequestBuilder) {
    return idcc656e3c5d595c7e17926a2796ec2029a16b9df71a54b6dd2ced6343b191c8d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Instances()(*ib59086c56b738f1f0717fb1ceaf00dffeac922ef570d9103a1ff449f48fca45d.InstancesRequestBuilder) {
    return ib59086c56b738f1f0717fb1ceaf00dffeac922ef570d9103a1ff449f48fca45d.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i4cdf65b58617609797304bf46ce9451a705dd1606550f29e20aa67daf8236b42.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i4cdf65b58617609797304bf46ce9451a705dd1606550f29e20aa67daf8236b42.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i54078d1892815baae8a8fd0cece8461711b63ba916021bd5ced1c30195a1a698.MultiValueExtendedPropertiesRequestBuilder) {
    return i54078d1892815baae8a8fd0cece8461711b63ba916021bd5ced1c30195a1a698.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2a46836ccae42574e0c03c77e8963f70e5d011a92ca307f636d258be89ac1916.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i2a46836ccae42574e0c03c77e8963f70e5d011a92ca307f636d258be89ac1916.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*id5711ae3f0bfb22fb24dc70cb4be78fa08cb4a509118586258c0f33adb016cf2.SingleValueExtendedPropertiesRequestBuilder) {
    return id5711ae3f0bfb22fb24dc70cb4be78fa08cb4a509118586258c0f33adb016cf2.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic43a8f89b11e7c86745f7bafd213b5d40252e1081389e9965812ca0da060a016.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ic43a8f89b11e7c86745f7bafd213b5d40252e1081389e9965812ca0da060a016.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i5726cc676e11fe9fa90959a9a58ddf2f16c34f2a86c819a81f6f2ce4937b710a.SnoozeReminderRequestBuilder) {
    return i5726cc676e11fe9fa90959a9a58ddf2f16c34f2a86c819a81f6f2ce4937b710a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i774cf8a56cdf6984d443e1380e0d0dcd2ca8d118dde0425d0b60e04c5f1c436c.TentativelyAcceptRequestBuilder) {
    return i774cf8a56cdf6984d443e1380e0d0dcd2ca8d118dde0425d0b60e04c5f1c436c.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
