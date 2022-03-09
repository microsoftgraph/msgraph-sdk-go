package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i1398a2b6b7b6fa6075615f0b696c5d5d393449822e80a67ff6e1f1b006570589 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/cancel"
    i23879bc3fa3602cbaf93cee98173f588cf737baf3fa364fa5d532b4a7ccffb6f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/instances"
    i4b518553c867ab557e4a9014d64bff52b4a5614abc228a534b622097747f58c4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/extensions"
    i855851d197827c1ee6475b539166033ff1c226b94affc6b7876227a822cefe56 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/singlevalueextendedproperties"
    i91e19aea3690b9482c589e1b4fffc198beb93c8882e7ad5dc30a2c61b975241d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/forward"
    i967611b2c88418ecd27c9ba55d8e1a0fb980af96f161a7941b0122d0af98e20b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/multivalueextendedproperties"
    i973bdbcda3dd1f6f25a1482be6ac8c9929a425f352242a5f72ad56ed2695a133 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/dismissreminder"
    i9d4cd76f12610ad565059847d3a79b13115ec8cc1967d272b7f9e33b62908d70 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/decline"
    ia9eabac6ecd410545dc4495208341bd846651d6673db813cabee1323b7ccebb1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/attachments"
    iaf6ca7f88694008763f7c64dc5f9e9a4696d364070e649a47f06c9cbb9720f79 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/accept"
    ic09438710245b2671bf04abe5db2656009ca7799d1d80105aa23a063cf2f4e40 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/tentativelyaccept"
    idaa6a05522f747ce1a1a1c6a68f0551d5b3d05627b4ce5ddef9c85b80d4684b2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar"
    ie274417b453065e3b9c1c84afef158b9149fe29aa54c90fc8fe2a8361a6a6880 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/snoozereminder"
    i3924f06647673effb94eed5dd5d2d7fe2b5b2f58ac82d0f02eab48c1219f1c07 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/multivalueextendedproperties/item"
    i5042ede149a551ff0f8fb6fd788882bbcfd795931aefebb153289527246e5e4d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/instances/item"
    i68e56f812a05f1c50a7fc8625d20c43951316b2122068afdbe3b9ec586bd6fe2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/extensions/item"
    i9fcaa4fe7aa3132d8c368ba6abc3e2992d0ff4b8a67614921b91e92aab8cab92 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/singlevalueextendedproperties/item"
    ia3fe1bf9b023096d20271ede0e7f809b9ea0d5dc59475200a395629f1ecb917c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.group entity.
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
// EventItemRequestBuilderGetQueryParameters the group's calendar events.
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
func (m *EventItemRequestBuilder) Accept()(*iaf6ca7f88694008763f7c64dc5f9e9a4696d364070e649a47f06c9cbb9720f79.AcceptRequestBuilder) {
    return iaf6ca7f88694008763f7c64dc5f9e9a4696d364070e649a47f06c9cbb9720f79.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ia9eabac6ecd410545dc4495208341bd846651d6673db813cabee1323b7ccebb1.AttachmentsRequestBuilder) {
    return ia9eabac6ecd410545dc4495208341bd846651d6673db813cabee1323b7ccebb1.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia3fe1bf9b023096d20271ede0e7f809b9ea0d5dc59475200a395629f1ecb917c.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia3fe1bf9b023096d20271ede0e7f809b9ea0d5dc59475200a395629f1ecb917c.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*idaa6a05522f747ce1a1a1c6a68f0551d5b3d05627b4ce5ddef9c85b80d4684b2.CalendarRequestBuilder) {
    return idaa6a05522f747ce1a1a1c6a68f0551d5b3d05627b4ce5ddef9c85b80d4684b2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i1398a2b6b7b6fa6075615f0b696c5d5d393449822e80a67ff6e1f1b006570589.CancelRequestBuilder) {
    return i1398a2b6b7b6fa6075615f0b696c5d5d393449822e80a67ff6e1f1b006570589.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation delete navigation property events for groups
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
// CreateGetRequestInformation the group's calendar events.
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
// CreatePatchRequestInformation update the navigation property events in groups
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
func (m *EventItemRequestBuilder) Decline()(*i9d4cd76f12610ad565059847d3a79b13115ec8cc1967d272b7f9e33b62908d70.DeclineRequestBuilder) {
    return i9d4cd76f12610ad565059847d3a79b13115ec8cc1967d272b7f9e33b62908d70.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*i973bdbcda3dd1f6f25a1482be6ac8c9929a425f352242a5f72ad56ed2695a133.DismissReminderRequestBuilder) {
    return i973bdbcda3dd1f6f25a1482be6ac8c9929a425f352242a5f72ad56ed2695a133.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i4b518553c867ab557e4a9014d64bff52b4a5614abc228a534b622097747f58c4.ExtensionsRequestBuilder) {
    return i4b518553c867ab557e4a9014d64bff52b4a5614abc228a534b622097747f58c4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i68e56f812a05f1c50a7fc8625d20c43951316b2122068afdbe3b9ec586bd6fe2.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i68e56f812a05f1c50a7fc8625d20c43951316b2122068afdbe3b9ec586bd6fe2.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i91e19aea3690b9482c589e1b4fffc198beb93c8882e7ad5dc30a2c61b975241d.ForwardRequestBuilder) {
    return i91e19aea3690b9482c589e1b4fffc198beb93c8882e7ad5dc30a2c61b975241d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the group's calendar events.
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
func (m *EventItemRequestBuilder) Instances()(*i23879bc3fa3602cbaf93cee98173f588cf737baf3fa364fa5d532b4a7ccffb6f.InstancesRequestBuilder) {
    return i23879bc3fa3602cbaf93cee98173f588cf737baf3fa364fa5d532b4a7ccffb6f.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i5042ede149a551ff0f8fb6fd788882bbcfd795931aefebb153289527246e5e4d.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i5042ede149a551ff0f8fb6fd788882bbcfd795931aefebb153289527246e5e4d.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i967611b2c88418ecd27c9ba55d8e1a0fb980af96f161a7941b0122d0af98e20b.MultiValueExtendedPropertiesRequestBuilder) {
    return i967611b2c88418ecd27c9ba55d8e1a0fb980af96f161a7941b0122d0af98e20b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i3924f06647673effb94eed5dd5d2d7fe2b5b2f58ac82d0f02eab48c1219f1c07.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i3924f06647673effb94eed5dd5d2d7fe2b5b2f58ac82d0f02eab48c1219f1c07.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i855851d197827c1ee6475b539166033ff1c226b94affc6b7876227a822cefe56.SingleValueExtendedPropertiesRequestBuilder) {
    return i855851d197827c1ee6475b539166033ff1c226b94affc6b7876227a822cefe56.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9fcaa4fe7aa3132d8c368ba6abc3e2992d0ff4b8a67614921b91e92aab8cab92.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i9fcaa4fe7aa3132d8c368ba6abc3e2992d0ff4b8a67614921b91e92aab8cab92.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie274417b453065e3b9c1c84afef158b9149fe29aa54c90fc8fe2a8361a6a6880.SnoozeReminderRequestBuilder) {
    return ie274417b453065e3b9c1c84afef158b9149fe29aa54c90fc8fe2a8361a6a6880.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ic09438710245b2671bf04abe5db2656009ca7799d1d80105aa23a063cf2f4e40.TentativelyAcceptRequestBuilder) {
    return ic09438710245b2671bf04abe5db2656009ca7799d1d80105aa23a063cf2f4e40.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
