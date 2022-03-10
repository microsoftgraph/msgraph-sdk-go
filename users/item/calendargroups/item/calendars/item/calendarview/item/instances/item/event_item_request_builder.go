package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i129d2b2b7386bbe656bf60028d4f5bea68f14c7529c9ce22900eb9d173d5607f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/decline"
    i5912a36466869b15f1760f5481d7f99751bf0fa1994c8956a162b1c16c782d6d "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/dismissreminder"
    i65cd74bb06865abc6771396e5525947cec5078cfca4ee99362efe3a5e866cc94 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/extensions"
    i7559f3b6e95df335c7bfc1089a4378d407a819b47827e4a3f3152d67a87136d4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/forward"
    i799e30b06602a7df5a17d8afc28cc6a8b103b8165d900f9fce26ba3de4cc0844 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/accept"
    ia1d9d5b7a2a2cfd4026a0df2895f07012da0011fb2fb585c22567a2e4f44ec89 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/tentativelyaccept"
    ia3f03df8544ee71d749c96e3fe2255384fa21dc18cdf1722846800477b7f21e6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/cancel"
    iaf8768970d08e349ba7031c622571c57023cb2a20d30187c3bb368fc6b5539a3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/snoozereminder"
    ic1edefa996b4d1844c76fb2f84a777aa0a65bfd901eccb28186e34ad7e97efc8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/calendar"
    ie6c0dac54e647e42c0d889e3dcfdc2545b660a93ff304ee74b07c315d5fb95ab "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/attachments"
    if75be8d844c5daa95fb1d84905bae6ec1f712cd6ca806ac56b050d068c17c8d7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties"
    ifaac39ed7e498fa2b4b90c7e2365897c36e7d994bfa5612dfcde2a6969648a71 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/multivalueextendedproperties"
    i20539128e96a2f01c57fe0224117c0d8ace248f888d69a0c3f769a7f1c7c9b09 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/extensions/item"
    i2a03976fdfd37c9d174c31049686743531c8cc24f38cc9e07945e7c1a7a82806 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    i7b23954f818a9b74508bc0a02417108455b8da36b94135440ff2f4b2dac36f46 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/attachments/item"
    i9025c7adc003d2d7747eef3ddca89aaeca8f0b6628378ab2b78bd5962a67b1b6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i799e30b06602a7df5a17d8afc28cc6a8b103b8165d900f9fce26ba3de4cc0844.AcceptRequestBuilder) {
    return i799e30b06602a7df5a17d8afc28cc6a8b103b8165d900f9fce26ba3de4cc0844.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ie6c0dac54e647e42c0d889e3dcfdc2545b660a93ff304ee74b07c315d5fb95ab.AttachmentsRequestBuilder) {
    return ie6c0dac54e647e42c0d889e3dcfdc2545b660a93ff304ee74b07c315d5fb95ab.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i7b23954f818a9b74508bc0a02417108455b8da36b94135440ff2f4b2dac36f46.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i7b23954f818a9b74508bc0a02417108455b8da36b94135440ff2f4b2dac36f46.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ic1edefa996b4d1844c76fb2f84a777aa0a65bfd901eccb28186e34ad7e97efc8.CalendarRequestBuilder) {
    return ic1edefa996b4d1844c76fb2f84a777aa0a65bfd901eccb28186e34ad7e97efc8.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ia3f03df8544ee71d749c96e3fe2255384fa21dc18cdf1722846800477b7f21e6.CancelRequestBuilder) {
    return ia3f03df8544ee71d749c96e3fe2255384fa21dc18cdf1722846800477b7f21e6.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/calendarView/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i129d2b2b7386bbe656bf60028d4f5bea68f14c7529c9ce22900eb9d173d5607f.DeclineRequestBuilder) {
    return i129d2b2b7386bbe656bf60028d4f5bea68f14c7529c9ce22900eb9d173d5607f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i5912a36466869b15f1760f5481d7f99751bf0fa1994c8956a162b1c16c782d6d.DismissReminderRequestBuilder) {
    return i5912a36466869b15f1760f5481d7f99751bf0fa1994c8956a162b1c16c782d6d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i65cd74bb06865abc6771396e5525947cec5078cfca4ee99362efe3a5e866cc94.ExtensionsRequestBuilder) {
    return i65cd74bb06865abc6771396e5525947cec5078cfca4ee99362efe3a5e866cc94.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i20539128e96a2f01c57fe0224117c0d8ace248f888d69a0c3f769a7f1c7c9b09.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i20539128e96a2f01c57fe0224117c0d8ace248f888d69a0c3f769a7f1c7c9b09.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i7559f3b6e95df335c7bfc1089a4378d407a819b47827e4a3f3152d67a87136d4.ForwardRequestBuilder) {
    return i7559f3b6e95df335c7bfc1089a4378d407a819b47827e4a3f3152d67a87136d4.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ifaac39ed7e498fa2b4b90c7e2365897c36e7d994bfa5612dfcde2a6969648a71.MultiValueExtendedPropertiesRequestBuilder) {
    return ifaac39ed7e498fa2b4b90c7e2365897c36e7d994bfa5612dfcde2a6969648a71.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9025c7adc003d2d7747eef3ddca89aaeca8f0b6628378ab2b78bd5962a67b1b6.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i9025c7adc003d2d7747eef3ddca89aaeca8f0b6628378ab2b78bd5962a67b1b6.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*if75be8d844c5daa95fb1d84905bae6ec1f712cd6ca806ac56b050d068c17c8d7.SingleValueExtendedPropertiesRequestBuilder) {
    return if75be8d844c5daa95fb1d84905bae6ec1f712cd6ca806ac56b050d068c17c8d7.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i2a03976fdfd37c9d174c31049686743531c8cc24f38cc9e07945e7c1a7a82806.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i2a03976fdfd37c9d174c31049686743531c8cc24f38cc9e07945e7c1a7a82806.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*iaf8768970d08e349ba7031c622571c57023cb2a20d30187c3bb368fc6b5539a3.SnoozeReminderRequestBuilder) {
    return iaf8768970d08e349ba7031c622571c57023cb2a20d30187c3bb368fc6b5539a3.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ia1d9d5b7a2a2cfd4026a0df2895f07012da0011fb2fb585c22567a2e4f44ec89.TentativelyAcceptRequestBuilder) {
    return ia1d9d5b7a2a2cfd4026a0df2895f07012da0011fb2fb585c22567a2e4f44ec89.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
