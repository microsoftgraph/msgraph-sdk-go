package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i15a1cf85dcc7eeb107a61041c861c7e258520efb3c3af054aac161a329f2c15e "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/snoozereminder"
    i251db03f59cc75f48f84e7245a0dc6f3de00400380451646430f584012111760 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/tentativelyaccept"
    i29b0d509f093b40c14a151adbb10e0aab35cfe53c1d405745d4c2976d34fd389 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/decline"
    i29c1c4fc5baeb4284286c65b18c5c49cd5481668cdc91c23a63a6802353fae13 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/multivalueextendedproperties"
    i2d66a35c35d6c7f3f7ea7c5db7ef1dabea98b07af7035f0b05c692c32145e5ee "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties"
    i352926d1ca610ee05029220da93f8400faee6b385271b5f15acb9c1b3d7d74f4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/extensions"
    i3d78d5747cc294b4c322bc7aeeca2cc53380d9b4bad2965e21c621c160c79fa4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/attachments"
    i72a6e89f7ed2091ef6b627975e8427d3601c1922304a5d1f43cc3dffdda9913e "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/forward"
    i90c9ce46a5ec66dcd1b80481bc789fbb3028677cbc3c32edfae1bc6d1d519a54 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/dismissreminder"
    i95bb8b9f1540cf7b97dc6ccf9364b1acc0b031644c22f4e03491a1700a790543 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/cancel"
    ib7a7049c46e443ff9e41b4c969d1d2ce89b58a39ddbce1e602cb74d15077df88 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/calendar"
    ie17a7797e2d679fd92ef43065db9290c8683a8f64c3279969a5eb21848d09d73 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/instances"
    if0ef0ca8fe12aeb76b29f3a545e38f8de508aed0b34c5517d3799e69f33f51f5 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/accept"
    i442111c87f659072c1c2f074352a85ebd625c1e18c68b3077096aabd48a95461 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/attachments/item"
    i51b055bd1bde3fd826011eb86802dc300fac30381c9947123d8a124cf68f9c84 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/extensions/item"
    ibf36b2390d0bb0a736c630f7f1cc55af3ed6227a90b2fbb973dc5bd7ac937b7b "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties/item"
    ic6feaf2955661d48f6f8c3cf286ef82befbf22fa19c47374ebd2880710008c5d "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item"
    ica66966e5d0c3577314d4dcc18c27d0b23efe8e86f729b91df3bd4c3e6ecedb0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/multivalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters
}
// EventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*if0ef0ca8fe12aeb76b29f3a545e38f8de508aed0b34c5517d3799e69f33f51f5.AcceptRequestBuilder) {
    return if0ef0ca8fe12aeb76b29f3a545e38f8de508aed0b34c5517d3799e69f33f51f5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i3d78d5747cc294b4c322bc7aeeca2cc53380d9b4bad2965e21c621c160c79fa4.AttachmentsRequestBuilder) {
    return i3d78d5747cc294b4c322bc7aeeca2cc53380d9b4bad2965e21c621c160c79fa4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i442111c87f659072c1c2f074352a85ebd625c1e18c68b3077096aabd48a95461.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i442111c87f659072c1c2f074352a85ebd625c1e18c68b3077096aabd48a95461.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ib7a7049c46e443ff9e41b4c969d1d2ce89b58a39ddbce1e602cb74d15077df88.CalendarRequestBuilder) {
    return ib7a7049c46e443ff9e41b4c969d1d2ce89b58a39ddbce1e602cb74d15077df88.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i95bb8b9f1540cf7b97dc6ccf9364b1acc0b031644c22f4e03491a1700a790543.CancelRequestBuilder) {
    return i95bb8b9f1540cf7b97dc6ccf9364b1acc0b031644c22f4e03491a1700a790543.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}{?%24select}";
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
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property events for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property events in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property events in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i29b0d509f093b40c14a151adbb10e0aab35cfe53c1d405745d4c2976d34fd389.DeclineRequestBuilder) {
    return i29b0d509f093b40c14a151adbb10e0aab35cfe53c1d405745d4c2976d34fd389.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property events for me
func (m *EventItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i90c9ce46a5ec66dcd1b80481bc789fbb3028677cbc3c32edfae1bc6d1d519a54.DismissReminderRequestBuilder) {
    return i90c9ce46a5ec66dcd1b80481bc789fbb3028677cbc3c32edfae1bc6d1d519a54.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i352926d1ca610ee05029220da93f8400faee6b385271b5f15acb9c1b3d7d74f4.ExtensionsRequestBuilder) {
    return i352926d1ca610ee05029220da93f8400faee6b385271b5f15acb9c1b3d7d74f4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i51b055bd1bde3fd826011eb86802dc300fac30381c9947123d8a124cf68f9c84.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i51b055bd1bde3fd826011eb86802dc300fac30381c9947123d8a124cf68f9c84.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i72a6e89f7ed2091ef6b627975e8427d3601c1922304a5d1f43cc3dffdda9913e.ForwardRequestBuilder) {
    return i72a6e89f7ed2091ef6b627975e8427d3601c1922304a5d1f43cc3dffdda9913e.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable), nil
}
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*ie17a7797e2d679fd92ef43065db9290c8683a8f64c3279969a5eb21848d09d73.InstancesRequestBuilder) {
    return ie17a7797e2d679fd92ef43065db9290c8683a8f64c3279969a5eb21848d09d73.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ic6feaf2955661d48f6f8c3cf286ef82befbf22fa19c47374ebd2880710008c5d.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return ic6feaf2955661d48f6f8c3cf286ef82befbf22fa19c47374ebd2880710008c5d.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i29c1c4fc5baeb4284286c65b18c5c49cd5481668cdc91c23a63a6802353fae13.MultiValueExtendedPropertiesRequestBuilder) {
    return i29c1c4fc5baeb4284286c65b18c5c49cd5481668cdc91c23a63a6802353fae13.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ica66966e5d0c3577314d4dcc18c27d0b23efe8e86f729b91df3bd4c3e6ecedb0.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ica66966e5d0c3577314d4dcc18c27d0b23efe8e86f729b91df3bd4c3e6ecedb0.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in me
func (m *EventItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property events in me
func (m *EventItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i2d66a35c35d6c7f3f7ea7c5db7ef1dabea98b07af7035f0b05c692c32145e5ee.SingleValueExtendedPropertiesRequestBuilder) {
    return i2d66a35c35d6c7f3f7ea7c5db7ef1dabea98b07af7035f0b05c692c32145e5ee.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ibf36b2390d0bb0a736c630f7f1cc55af3ed6227a90b2fbb973dc5bd7ac937b7b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ibf36b2390d0bb0a736c630f7f1cc55af3ed6227a90b2fbb973dc5bd7ac937b7b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i15a1cf85dcc7eeb107a61041c861c7e258520efb3c3af054aac161a329f2c15e.SnoozeReminderRequestBuilder) {
    return i15a1cf85dcc7eeb107a61041c861c7e258520efb3c3af054aac161a329f2c15e.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i251db03f59cc75f48f84e7245a0dc6f3de00400380451646430f584012111760.TentativelyAcceptRequestBuilder) {
    return i251db03f59cc75f48f84e7245a0dc6f3de00400380451646430f584012111760.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
