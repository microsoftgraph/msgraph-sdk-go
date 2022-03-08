package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0a0fc3d004d78576d043eb521ad26f53fe8214afd2b3ebe973fe93e60f680f71 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/attachments"
    i16240d0de448bc4c6216c08277550221d94bc0f48ba2accd218749ad7c953187 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/extensions"
    i3c7b0004be1fa25f87efde7027fd8e976d9cf7bf986843ed38a31557bf5e7a95 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/accept"
    i49c9c1594c0f4696377dfecc53ccf788c8e8b710f8313547cbcd0af990f5bbf6 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/cancel"
    i68b73555ae5e3f902578ad78535a4acc022d87da89f72cc819d660c671ed5ece "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/tentativelyaccept"
    i8bfe1260d815ef118a4accfe78279062de1c84617b6bf2fd2fcbee8aaa6e4a3e "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/multivalueextendedproperties"
    ib177b4b7402e2bd1b4c939d9165f6cc6ad4c888362c100157bbf3567e762c3a8 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/singlevalueextendedproperties"
    ib7d6bec077bf9e6ed260a5dcae2c2e771c594d05995ed7713b4f6d55c140e8e9 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/snoozereminder"
    ic273ea006cf9b2cbb523ae048eff477ad9414b6846e6486ee9ecb96844555f5e "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/forward"
    idc8f27b80133ba40f0ebae7ae73b0927fa97d40f63c920f7718b0b1bb3646f33 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/decline"
    idcb9651181c4cfa39068cf83f36f38021b6d83f2785c96c836454a8066720a49 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/dismissreminder"
    if25488a44284a9c63a32442f9d9a88bb047100fbcbc4cbe3daa8228fb39d90d7 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/calendar"
    i684e98bc756b097dd50799c2b776635f96b124dbc95dd1f25f6d308ac7ecfa6f "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/attachments/item"
    ia70c4b879b44928811313bb8380c609fb4f17a2429974bfc650c78042dd445a2 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/extensions/item"
    ib4ab2585af40e6299bd85e8d11fe39d4c9e2950cba8aaab47b74a63390ade002 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/singlevalueextendedproperties/item"
    ifb082cc9281770b0f14bb66061bb2f300737211dde2040e7b7d2bb48d23b51ae "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i3c7b0004be1fa25f87efde7027fd8e976d9cf7bf986843ed38a31557bf5e7a95.AcceptRequestBuilder) {
    return i3c7b0004be1fa25f87efde7027fd8e976d9cf7bf986843ed38a31557bf5e7a95.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i0a0fc3d004d78576d043eb521ad26f53fe8214afd2b3ebe973fe93e60f680f71.AttachmentsRequestBuilder) {
    return i0a0fc3d004d78576d043eb521ad26f53fe8214afd2b3ebe973fe93e60f680f71.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i684e98bc756b097dd50799c2b776635f96b124dbc95dd1f25f6d308ac7ecfa6f.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i684e98bc756b097dd50799c2b776635f96b124dbc95dd1f25f6d308ac7ecfa6f.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*if25488a44284a9c63a32442f9d9a88bb047100fbcbc4cbe3daa8228fb39d90d7.CalendarRequestBuilder) {
    return if25488a44284a9c63a32442f9d9a88bb047100fbcbc4cbe3daa8228fb39d90d7.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i49c9c1594c0f4696377dfecc53ccf788c8e8b710f8313547cbcd0af990f5bbf6.CancelRequestBuilder) {
    return i49c9c1594c0f4696377dfecc53ccf788c8e8b710f8313547cbcd0af990f5bbf6.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*idc8f27b80133ba40f0ebae7ae73b0927fa97d40f63c920f7718b0b1bb3646f33.DeclineRequestBuilder) {
    return idc8f27b80133ba40f0ebae7ae73b0927fa97d40f63c920f7718b0b1bb3646f33.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*idcb9651181c4cfa39068cf83f36f38021b6d83f2785c96c836454a8066720a49.DismissReminderRequestBuilder) {
    return idcb9651181c4cfa39068cf83f36f38021b6d83f2785c96c836454a8066720a49.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i16240d0de448bc4c6216c08277550221d94bc0f48ba2accd218749ad7c953187.ExtensionsRequestBuilder) {
    return i16240d0de448bc4c6216c08277550221d94bc0f48ba2accd218749ad7c953187.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ia70c4b879b44928811313bb8380c609fb4f17a2429974bfc650c78042dd445a2.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ia70c4b879b44928811313bb8380c609fb4f17a2429974bfc650c78042dd445a2.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*ic273ea006cf9b2cbb523ae048eff477ad9414b6846e6486ee9ecb96844555f5e.ForwardRequestBuilder) {
    return ic273ea006cf9b2cbb523ae048eff477ad9414b6846e6486ee9ecb96844555f5e.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable), nil
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i8bfe1260d815ef118a4accfe78279062de1c84617b6bf2fd2fcbee8aaa6e4a3e.MultiValueExtendedPropertiesRequestBuilder) {
    return i8bfe1260d815ef118a4accfe78279062de1c84617b6bf2fd2fcbee8aaa6e4a3e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ifb082cc9281770b0f14bb66061bb2f300737211dde2040e7b7d2bb48d23b51ae.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ifb082cc9281770b0f14bb66061bb2f300737211dde2040e7b7d2bb48d23b51ae.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib177b4b7402e2bd1b4c939d9165f6cc6ad4c888362c100157bbf3567e762c3a8.SingleValueExtendedPropertiesRequestBuilder) {
    return ib177b4b7402e2bd1b4c939d9165f6cc6ad4c888362c100157bbf3567e762c3a8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib4ab2585af40e6299bd85e8d11fe39d4c9e2950cba8aaab47b74a63390ade002.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib4ab2585af40e6299bd85e8d11fe39d4c9e2950cba8aaab47b74a63390ade002.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ib7d6bec077bf9e6ed260a5dcae2c2e771c594d05995ed7713b4f6d55c140e8e9.SnoozeReminderRequestBuilder) {
    return ib7d6bec077bf9e6ed260a5dcae2c2e771c594d05995ed7713b4f6d55c140e8e9.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i68b73555ae5e3f902578ad78535a4acc022d87da89f72cc819d660c671ed5ece.TentativelyAcceptRequestBuilder) {
    return i68b73555ae5e3f902578ad78535a4acc022d87da89f72cc819d660c671ed5ece.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
