package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2aa695d2031f70f5d124f226c0bb7d90cba2b281788a0e5f062cfd23464f568e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/dismissreminder"
    i32d20a3dbf1acc1aeb627a8d900b6fc7fc488631ad29c0fe0952e66026ab99f2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/forward"
    i33da0e0ab9fc76274bc798c0179ec96b44cd3ca8ae698eb71bb470d7d3a64266 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/singlevalueextendedproperties"
    i4aa0979a3421851170b7f9c219354310292ae88cecddd0c34707076ae4c39872 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/snoozereminder"
    i50529b46cb9815e1ee6989b50a5e7c9d65cdbb639fbeeb17a1745097ca1de34a "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/multivalueextendedproperties"
    i626438065d3e2e4196a8edc03c59cccec617b67714a70388b865384f45bf7d54 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/extensions"
    i76b4b348bf50e98fa2fe202aa8497f6cc42d90db777ae33eef2bad180ba1997e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/decline"
    i77ccc0b59e6aff950dfcfe7893300dcac0af0f4c6a63312c005fcd867b4af5c4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/cancel"
    i8409d063ec88dfe68141ac2e114f2ce751bd923950f2b07ebd6d6538402649d7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/attachments"
    ibbc9746075ab51b54d510dd1431f7223ed53e3c1ead9dbfbf8d4b1c10db05783 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/tentativelyaccept"
    ic2ec03aae4a44329c27dd64dfc4e4a30cf275384f3242568f08e627453d33cf9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/calendar"
    ie4edeb5b8615215dcf77c191289d9633b816d905d2859e386e9bbeaa680a7924 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/accept"
    i14989f66f870841793770a7d9450c5e1f60ce6f731702c9e59616eb27ef280ec "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/extensions/item"
    i44fd1ffc04685be592ee41c9e9af8f5115b183df895b51a20476f788d5cdcff2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/attachments/item"
    i62922412829845e613386fdd7e95855dafb210457259ca4401e86b26f35228bf "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/multivalueextendedproperties/item"
    icbe569d9ab31fcdf58f7ea64bcbe8538e08f365719a1773e15b606d4c87a7a33 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*ie4edeb5b8615215dcf77c191289d9633b816d905d2859e386e9bbeaa680a7924.AcceptRequestBuilder) {
    return ie4edeb5b8615215dcf77c191289d9633b816d905d2859e386e9bbeaa680a7924.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i8409d063ec88dfe68141ac2e114f2ce751bd923950f2b07ebd6d6538402649d7.AttachmentsRequestBuilder) {
    return i8409d063ec88dfe68141ac2e114f2ce751bd923950f2b07ebd6d6538402649d7.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i44fd1ffc04685be592ee41c9e9af8f5115b183df895b51a20476f788d5cdcff2.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i44fd1ffc04685be592ee41c9e9af8f5115b183df895b51a20476f788d5cdcff2.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ic2ec03aae4a44329c27dd64dfc4e4a30cf275384f3242568f08e627453d33cf9.CalendarRequestBuilder) {
    return ic2ec03aae4a44329c27dd64dfc4e4a30cf275384f3242568f08e627453d33cf9.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i77ccc0b59e6aff950dfcfe7893300dcac0af0f4c6a63312c005fcd867b4af5c4.CancelRequestBuilder) {
    return i77ccc0b59e6aff950dfcfe7893300dcac0af0f4c6a63312c005fcd867b4af5c4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/events/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i76b4b348bf50e98fa2fe202aa8497f6cc42d90db777ae33eef2bad180ba1997e.DeclineRequestBuilder) {
    return i76b4b348bf50e98fa2fe202aa8497f6cc42d90db777ae33eef2bad180ba1997e.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*i2aa695d2031f70f5d124f226c0bb7d90cba2b281788a0e5f062cfd23464f568e.DismissReminderRequestBuilder) {
    return i2aa695d2031f70f5d124f226c0bb7d90cba2b281788a0e5f062cfd23464f568e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i626438065d3e2e4196a8edc03c59cccec617b67714a70388b865384f45bf7d54.ExtensionsRequestBuilder) {
    return i626438065d3e2e4196a8edc03c59cccec617b67714a70388b865384f45bf7d54.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i14989f66f870841793770a7d9450c5e1f60ce6f731702c9e59616eb27ef280ec.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i14989f66f870841793770a7d9450c5e1f60ce6f731702c9e59616eb27ef280ec.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i32d20a3dbf1acc1aeb627a8d900b6fc7fc488631ad29c0fe0952e66026ab99f2.ForwardRequestBuilder) {
    return i32d20a3dbf1acc1aeb627a8d900b6fc7fc488631ad29c0fe0952e66026ab99f2.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable), nil
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i50529b46cb9815e1ee6989b50a5e7c9d65cdbb639fbeeb17a1745097ca1de34a.MultiValueExtendedPropertiesRequestBuilder) {
    return i50529b46cb9815e1ee6989b50a5e7c9d65cdbb639fbeeb17a1745097ca1de34a.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i62922412829845e613386fdd7e95855dafb210457259ca4401e86b26f35228bf.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i62922412829845e613386fdd7e95855dafb210457259ca4401e86b26f35228bf.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i33da0e0ab9fc76274bc798c0179ec96b44cd3ca8ae698eb71bb470d7d3a64266.SingleValueExtendedPropertiesRequestBuilder) {
    return i33da0e0ab9fc76274bc798c0179ec96b44cd3ca8ae698eb71bb470d7d3a64266.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*icbe569d9ab31fcdf58f7ea64bcbe8538e08f365719a1773e15b606d4c87a7a33.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return icbe569d9ab31fcdf58f7ea64bcbe8538e08f365719a1773e15b606d4c87a7a33.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i4aa0979a3421851170b7f9c219354310292ae88cecddd0c34707076ae4c39872.SnoozeReminderRequestBuilder) {
    return i4aa0979a3421851170b7f9c219354310292ae88cecddd0c34707076ae4c39872.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibbc9746075ab51b54d510dd1431f7223ed53e3c1ead9dbfbf8d4b1c10db05783.TentativelyAcceptRequestBuilder) {
    return ibbc9746075ab51b54d510dd1431f7223ed53e3c1ead9dbfbf8d4b1c10db05783.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
