package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i06b377af3ae66d378617ba4960aa8e040b78e63f3ebfd980bf3a5d47930f6ecc "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/forward"
    i1a802c2e4b157587519614c1bc8be3ea7021108354647debc56f0b6d913669bd "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/calendar"
    i22b4abf5d21f46083e388fe3297f3a331d292876d79491e062ea21ab4d5d5d96 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/cancel"
    i23d3284ba64b73b3325b23aef59a544daf0104c53a03ed2b31edafd13fc4e145 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/dismissreminder"
    i2641548965eaf8952f9123f07765870567794d4a5f1f35fbb5f9b4b9179f3cda "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/singlevalueextendedproperties"
    i2914318856007078157daa88f50c0c37f7a2d7216058973bf90effba2e5d7a68 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/accept"
    i2d17d3f526e425e1407b9c1b64ad35279c0bd50d1d5815931329eb39e568cf03 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/extensions"
    i3ebed318d699e805b107ab544bb8d4e7846a7e636121dc579c5d3f23b875434b "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/tentativelyaccept"
    i82457e731323ccb1502e74a2a6dd1d66fc84027eee6d3af7a3949d1172d5ff8f "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/decline"
    i8f6b1073998fe7e072e94da48aef27ff47795e232ec787943a4bc43820018dd7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/multivalueextendedproperties"
    ib18d29dbe5c1c698693687c9c8e985db4ab1ddf50a568e2d1bfca96b9b934e87 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/attachments"
    ice7a8351756d24c105c5bc2d1f4183f8e71e711980b3646eef4d0ab27daba9f5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/snoozereminder"
    i38f54e1f5ed4752a820f6c2288d81c7117bf023cb9f0e0280f49cb1e9cb0ddfd "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/extensions/item"
    i5e8b2a3530507dc013f518d7efff9dd7ebccefa3a76f24c3968b6f462a3ffb06 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/multivalueextendedproperties/item"
    i9528e41139478696eb2cf527c6102ee1417a84b633910fb7734d66d6603dfe33 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/singlevalueextendedproperties/item"
    ia583e36d9ac073d787b4f5e0abb093224bbfbbb0acf8b3a6f1b0f1e7dfd13a55 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i2914318856007078157daa88f50c0c37f7a2d7216058973bf90effba2e5d7a68.AcceptRequestBuilder) {
    return i2914318856007078157daa88f50c0c37f7a2d7216058973bf90effba2e5d7a68.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ib18d29dbe5c1c698693687c9c8e985db4ab1ddf50a568e2d1bfca96b9b934e87.AttachmentsRequestBuilder) {
    return ib18d29dbe5c1c698693687c9c8e985db4ab1ddf50a568e2d1bfca96b9b934e87.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia583e36d9ac073d787b4f5e0abb093224bbfbbb0acf8b3a6f1b0f1e7dfd13a55.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia583e36d9ac073d787b4f5e0abb093224bbfbbb0acf8b3a6f1b0f1e7dfd13a55.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i1a802c2e4b157587519614c1bc8be3ea7021108354647debc56f0b6d913669bd.CalendarRequestBuilder) {
    return i1a802c2e4b157587519614c1bc8be3ea7021108354647debc56f0b6d913669bd.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i22b4abf5d21f46083e388fe3297f3a331d292876d79491e062ea21ab4d5d5d96.CancelRequestBuilder) {
    return i22b4abf5d21f46083e388fe3297f3a331d292876d79491e062ea21ab4d5d5d96.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/events/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i82457e731323ccb1502e74a2a6dd1d66fc84027eee6d3af7a3949d1172d5ff8f.DeclineRequestBuilder) {
    return i82457e731323ccb1502e74a2a6dd1d66fc84027eee6d3af7a3949d1172d5ff8f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i23d3284ba64b73b3325b23aef59a544daf0104c53a03ed2b31edafd13fc4e145.DismissReminderRequestBuilder) {
    return i23d3284ba64b73b3325b23aef59a544daf0104c53a03ed2b31edafd13fc4e145.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i2d17d3f526e425e1407b9c1b64ad35279c0bd50d1d5815931329eb39e568cf03.ExtensionsRequestBuilder) {
    return i2d17d3f526e425e1407b9c1b64ad35279c0bd50d1d5815931329eb39e568cf03.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i38f54e1f5ed4752a820f6c2288d81c7117bf023cb9f0e0280f49cb1e9cb0ddfd.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i38f54e1f5ed4752a820f6c2288d81c7117bf023cb9f0e0280f49cb1e9cb0ddfd.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i06b377af3ae66d378617ba4960aa8e040b78e63f3ebfd980bf3a5d47930f6ecc.ForwardRequestBuilder) {
    return i06b377af3ae66d378617ba4960aa8e040b78e63f3ebfd980bf3a5d47930f6ecc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i8f6b1073998fe7e072e94da48aef27ff47795e232ec787943a4bc43820018dd7.MultiValueExtendedPropertiesRequestBuilder) {
    return i8f6b1073998fe7e072e94da48aef27ff47795e232ec787943a4bc43820018dd7.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5e8b2a3530507dc013f518d7efff9dd7ebccefa3a76f24c3968b6f462a3ffb06.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5e8b2a3530507dc013f518d7efff9dd7ebccefa3a76f24c3968b6f462a3ffb06.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i2641548965eaf8952f9123f07765870567794d4a5f1f35fbb5f9b4b9179f3cda.SingleValueExtendedPropertiesRequestBuilder) {
    return i2641548965eaf8952f9123f07765870567794d4a5f1f35fbb5f9b4b9179f3cda.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9528e41139478696eb2cf527c6102ee1417a84b633910fb7734d66d6603dfe33.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i9528e41139478696eb2cf527c6102ee1417a84b633910fb7734d66d6603dfe33.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ice7a8351756d24c105c5bc2d1f4183f8e71e711980b3646eef4d0ab27daba9f5.SnoozeReminderRequestBuilder) {
    return ice7a8351756d24c105c5bc2d1f4183f8e71e711980b3646eef4d0ab27daba9f5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3ebed318d699e805b107ab544bb8d4e7846a7e636121dc579c5d3f23b875434b.TentativelyAcceptRequestBuilder) {
    return i3ebed318d699e805b107ab544bb8d4e7846a7e636121dc579c5d3f23b875434b.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
