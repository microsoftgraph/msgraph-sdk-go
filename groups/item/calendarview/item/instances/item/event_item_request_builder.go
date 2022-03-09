package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i3cc34074d4799ae365015afe8a944936199187ea89a89982620bc4d5270ee20a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/accept"
    i4acb1eeebf60790e92c699b157738449aea78a952be2463ce026b8040d7e5789 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/multivalueextendedproperties"
    i578c3c432a4af3223ca10fc3355e1e507450b8186f1aafc298ff748751814861 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/attachments"
    i6974512b0df45bd6478c5bd76ba677d7332b9275200ed141bb1c82d44ac9d26f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/singlevalueextendedproperties"
    i756e61eebdd85c15621a1de3592845102e21d2510bcf4f47c8b6f633b3d7cc8e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/calendar"
    i7758ded31927896a0324ac6ba89d90109741d24d1a194a275042f6f228cdc4be "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/dismissreminder"
    i849f6cd4968ad54bff7c4acc3d2c814119ed8035747edd7da13892771dd1f6cf "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/forward"
    i8c1de96d676c31bb3baa4d5578220df8cc694fcb1499fa3ced34cc06d55588c9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/snoozereminder"
    ib55ccb1e43bdb65ed253495385ecec98286a535ace3067311f181b27a9a373d7 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/extensions"
    ib5a7878855ffd0e4fbdf84af0ae8b0f0a7974f30f055f4ced1d57fdb861e0fa5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/tentativelyaccept"
    ib9699fbc3be709b24c33827610ccf455d295816e912a53e123836e2c220e6877 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/decline"
    ic8a6a36a3fed4f835a1259b874a6c34cb4c966ccbabfa5b23cfa9244dbbb2151 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/cancel"
    i006d6330e3c4b5ce635e3dd7b2de257d6731702cbef84980ddf3887ec16eb40e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    i538b156a8aee4235e0fe1fbffcecd293c195c55b10dc0e4aaa92e8a97447ea26 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/attachments/item"
    iabde4e47b53f2370aeb27d4c48fa6dd8b902a360e08ca7a4f145a8b7d3e0fd2e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/extensions/item"
    ifd4e95066ee67088ba4bf159e28dd71f75754bc7edc71bf65b717d7f84239541 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i3cc34074d4799ae365015afe8a944936199187ea89a89982620bc4d5270ee20a.AcceptRequestBuilder) {
    return i3cc34074d4799ae365015afe8a944936199187ea89a89982620bc4d5270ee20a.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i578c3c432a4af3223ca10fc3355e1e507450b8186f1aafc298ff748751814861.AttachmentsRequestBuilder) {
    return i578c3c432a4af3223ca10fc3355e1e507450b8186f1aafc298ff748751814861.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i538b156a8aee4235e0fe1fbffcecd293c195c55b10dc0e4aaa92e8a97447ea26.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i538b156a8aee4235e0fe1fbffcecd293c195c55b10dc0e4aaa92e8a97447ea26.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i756e61eebdd85c15621a1de3592845102e21d2510bcf4f47c8b6f633b3d7cc8e.CalendarRequestBuilder) {
    return i756e61eebdd85c15621a1de3592845102e21d2510bcf4f47c8b6f633b3d7cc8e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ic8a6a36a3fed4f835a1259b874a6c34cb4c966ccbabfa5b23cfa9244dbbb2151.CancelRequestBuilder) {
    return ic8a6a36a3fed4f835a1259b874a6c34cb4c966ccbabfa5b23cfa9244dbbb2151.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendarView/{event_id}/instances/{event_id1}{?select}";
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
// CreateDeleteRequestInformation delete navigation property instances for groups
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
// CreatePatchRequestInformation update the navigation property instances in groups
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
func (m *EventItemRequestBuilder) Decline()(*ib9699fbc3be709b24c33827610ccf455d295816e912a53e123836e2c220e6877.DeclineRequestBuilder) {
    return ib9699fbc3be709b24c33827610ccf455d295816e912a53e123836e2c220e6877.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*i7758ded31927896a0324ac6ba89d90109741d24d1a194a275042f6f228cdc4be.DismissReminderRequestBuilder) {
    return i7758ded31927896a0324ac6ba89d90109741d24d1a194a275042f6f228cdc4be.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*ib55ccb1e43bdb65ed253495385ecec98286a535ace3067311f181b27a9a373d7.ExtensionsRequestBuilder) {
    return ib55ccb1e43bdb65ed253495385ecec98286a535ace3067311f181b27a9a373d7.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iabde4e47b53f2370aeb27d4c48fa6dd8b902a360e08ca7a4f145a8b7d3e0fd2e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iabde4e47b53f2370aeb27d4c48fa6dd8b902a360e08ca7a4f145a8b7d3e0fd2e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i849f6cd4968ad54bff7c4acc3d2c814119ed8035747edd7da13892771dd1f6cf.ForwardRequestBuilder) {
    return i849f6cd4968ad54bff7c4acc3d2c814119ed8035747edd7da13892771dd1f6cf.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i4acb1eeebf60790e92c699b157738449aea78a952be2463ce026b8040d7e5789.MultiValueExtendedPropertiesRequestBuilder) {
    return i4acb1eeebf60790e92c699b157738449aea78a952be2463ce026b8040d7e5789.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ifd4e95066ee67088ba4bf159e28dd71f75754bc7edc71bf65b717d7f84239541.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ifd4e95066ee67088ba4bf159e28dd71f75754bc7edc71bf65b717d7f84239541.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i6974512b0df45bd6478c5bd76ba677d7332b9275200ed141bb1c82d44ac9d26f.SingleValueExtendedPropertiesRequestBuilder) {
    return i6974512b0df45bd6478c5bd76ba677d7332b9275200ed141bb1c82d44ac9d26f.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i006d6330e3c4b5ce635e3dd7b2de257d6731702cbef84980ddf3887ec16eb40e.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i006d6330e3c4b5ce635e3dd7b2de257d6731702cbef84980ddf3887ec16eb40e.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i8c1de96d676c31bb3baa4d5578220df8cc694fcb1499fa3ced34cc06d55588c9.SnoozeReminderRequestBuilder) {
    return i8c1de96d676c31bb3baa4d5578220df8cc694fcb1499fa3ced34cc06d55588c9.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ib5a7878855ffd0e4fbdf84af0ae8b0f0a7974f30f055f4ced1d57fdb861e0fa5.TentativelyAcceptRequestBuilder) {
    return ib5a7878855ffd0e4fbdf84af0ae8b0f0a7974f30f055f4ced1d57fdb861e0fa5.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
