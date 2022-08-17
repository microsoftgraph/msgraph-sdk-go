package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1a3b6643861b0a0caa0ff1eb68e5c488a43aff0aa1f3a06093f6e226fa2b917d "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/dismissreminder"
    i32c37b1c04b5be6ff9366461bd21b5fde755e861ca804c535263e878e67f6c27 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/forward"
    i36d3a71fa12eb21299e38ecfe8e5c1c9858e945064a266a1dc3886953114eece "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances"
    i40517a193f3b25c5466a4fb68f9a11b390590bed4b80557218b6fa94bba2bb03 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/multivalueextendedproperties"
    i5823ac8df2710b8a2e2cffa605921bb14b14de21e070c27b7581b8e8a5ba871e "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/cancel"
    i70149e4bcbdc49f55d93a41203bb62f881cf96a4136c6be6de44d5750c55cbec "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/tentativelyaccept"
    i735b02580c067f7e9c04d97f0df8b9805febcc2badf5811dc6135a7bcbd36880 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/decline"
    i73fdb8c241b8badf777c9058cb4d1f11e3c5329111efcaf2510b52537daa9874 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/snoozereminder"
    i7bd7ba2ce48d243a859e20ef25790047a8675315ee1da0126957422798414ec0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/accept"
    i91c0a715e42a5cbf02de3870dd658b8c3716027385e23055057583d5a1b3597e "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/extensions"
    ib5ca28c6c72c1db8b6f37b269c2876fab903b374b0d30cc8404a167730803940 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/attachments"
    ibc4cc90649e04bb4574982b4a76356636e2914d984eb4131c9ea6e287f13ebdb "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/singlevalueextendedproperties"
    idd1ffd3d2adc7fab21de995ec4b6949ca9bb47b13c5e7cc18129c1449dc2651a "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar"
    i58869d35545fe6338f6dd5b999816eceba5ff2ae808d4a6636dd981dc0025fa3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/multivalueextendedproperties/item"
    i95164ba391501ae53794b155baaed0825128639bb413adda940b803c1b5ae14a "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item"
    ief46ce2c42a9faa7186bd9c6c9a601dc9f0722341a69677d6743c991650b60fd "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/attachments/item"
    if6d2f1f347111fdd7f5b334372ce95e9ea9e20681fbc3bd7a3a39fdb53b19ca4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/extensions/item"
    ifccd98b65865537ee9ddb082f187a02b159c058c0cdbeb4ffa28207f93bca168 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/singlevalueextendedproperties/item"
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
// EventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EventItemRequestBuilderGetQueryParameters the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Accept()(*i7bd7ba2ce48d243a859e20ef25790047a8675315ee1da0126957422798414ec0.AcceptRequestBuilder) {
    return i7bd7ba2ce48d243a859e20ef25790047a8675315ee1da0126957422798414ec0.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ib5ca28c6c72c1db8b6f37b269c2876fab903b374b0d30cc8404a167730803940.AttachmentsRequestBuilder) {
    return ib5ca28c6c72c1db8b6f37b269c2876fab903b374b0d30cc8404a167730803940.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ief46ce2c42a9faa7186bd9c6c9a601dc9f0722341a69677d6743c991650b60fd.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ief46ce2c42a9faa7186bd9c6c9a601dc9f0722341a69677d6743c991650b60fd.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*idd1ffd3d2adc7fab21de995ec4b6949ca9bb47b13c5e7cc18129c1449dc2651a.CalendarRequestBuilder) {
    return idd1ffd3d2adc7fab21de995ec4b6949ca9bb47b13c5e7cc18129c1449dc2651a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i5823ac8df2710b8a2e2cffa605921bb14b14de21e070c27b7581b8e8a5ba871e.CancelRequestBuilder) {
    return i5823ac8df2710b8a2e2cffa605921bb14b14de21e070c27b7581b8e8a5ba871e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/events/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property events for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property events for users
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
// CreateGetRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property events in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property events in users
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
func (m *EventItemRequestBuilder) Decline()(*i735b02580c067f7e9c04d97f0df8b9805febcc2badf5811dc6135a7bcbd36880.DeclineRequestBuilder) {
    return i735b02580c067f7e9c04d97f0df8b9805febcc2badf5811dc6135a7bcbd36880.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property events for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i1a3b6643861b0a0caa0ff1eb68e5c488a43aff0aa1f3a06093f6e226fa2b917d.DismissReminderRequestBuilder) {
    return i1a3b6643861b0a0caa0ff1eb68e5c488a43aff0aa1f3a06093f6e226fa2b917d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i91c0a715e42a5cbf02de3870dd658b8c3716027385e23055057583d5a1b3597e.ExtensionsRequestBuilder) {
    return i91c0a715e42a5cbf02de3870dd658b8c3716027385e23055057583d5a1b3597e.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*if6d2f1f347111fdd7f5b334372ce95e9ea9e20681fbc3bd7a3a39fdb53b19ca4.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return if6d2f1f347111fdd7f5b334372ce95e9ea9e20681fbc3bd7a3a39fdb53b19ca4.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i32c37b1c04b5be6ff9366461bd21b5fde755e861ca804c535263e878e67f6c27.ForwardRequestBuilder) {
    return i32c37b1c04b5be6ff9366461bd21b5fde755e861ca804c535263e878e67f6c27.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Instances()(*i36d3a71fa12eb21299e38ecfe8e5c1c9858e945064a266a1dc3886953114eece.InstancesRequestBuilder) {
    return i36d3a71fa12eb21299e38ecfe8e5c1c9858e945064a266a1dc3886953114eece.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i95164ba391501ae53794b155baaed0825128639bb413adda940b803c1b5ae14a.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i95164ba391501ae53794b155baaed0825128639bb413adda940b803c1b5ae14a.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i40517a193f3b25c5466a4fb68f9a11b390590bed4b80557218b6fa94bba2bb03.MultiValueExtendedPropertiesRequestBuilder) {
    return i40517a193f3b25c5466a4fb68f9a11b390590bed4b80557218b6fa94bba2bb03.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i58869d35545fe6338f6dd5b999816eceba5ff2ae808d4a6636dd981dc0025fa3.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i58869d35545fe6338f6dd5b999816eceba5ff2ae808d4a6636dd981dc0025fa3.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in users
func (m *EventItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property events in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ibc4cc90649e04bb4574982b4a76356636e2914d984eb4131c9ea6e287f13ebdb.SingleValueExtendedPropertiesRequestBuilder) {
    return ibc4cc90649e04bb4574982b4a76356636e2914d984eb4131c9ea6e287f13ebdb.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifccd98b65865537ee9ddb082f187a02b159c058c0cdbeb4ffa28207f93bca168.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ifccd98b65865537ee9ddb082f187a02b159c058c0cdbeb4ffa28207f93bca168.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i73fdb8c241b8badf777c9058cb4d1f11e3c5329111efcaf2510b52537daa9874.SnoozeReminderRequestBuilder) {
    return i73fdb8c241b8badf777c9058cb4d1f11e3c5329111efcaf2510b52537daa9874.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i70149e4bcbdc49f55d93a41203bb62f881cf96a4136c6be6de44d5750c55cbec.TentativelyAcceptRequestBuilder) {
    return i70149e4bcbdc49f55d93a41203bb62f881cf96a4136c6be6de44d5750c55cbec.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
