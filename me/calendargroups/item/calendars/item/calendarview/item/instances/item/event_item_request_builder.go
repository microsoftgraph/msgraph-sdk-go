package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i068c994d6e3103d7a8c451470fd8946c9a526bf7defe804a3414bec3f058c7c0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/attachments"
    i178988a2b9571cc52545bd2a40fe9246c6b63e31f9a8adbe3f7c47021e3ebfd5 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/snoozereminder"
    i39f7b9eb2753296d89e1366bb2fa48858540dd472648ec034a0c3a194183a61c "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/cancel"
    i65f4d175fc410ed923f2cd477c6e12f9b6cfed24669da0d6147a1eea2b7353b0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/forward"
    i6d2604c2a16b79e867c6950984ac155a11a951d6935ff922fb3d80299b8aae9c "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties"
    i732564651bd295ace901245d83f7cee6eecbd49650f303b4d6f278dce8d4196c "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/decline"
    i75b85d086dff2c57d9c90524d349608e34c224e7740e4a9a09c8e3892e4d650b "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/calendar"
    ia3e61a8f98e0b4ebf997055c903954f21cee726b10079e1c431a38baa7bb1de9 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/accept"
    iaa82cd4a7429733dee7556dd9022c3f856a8e77926ee4dba3e89fb673d6c269b "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/tentativelyaccept"
    ic38a64ab747dedc47441f8456aa350055b7b0ce2810f51c3c45312d8542baab0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/dismissreminder"
    ic7e25b5dfc5b0afcb9c1a08774f134a922789857f8a916436156eac48cd979cc "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/multivalueextendedproperties"
    icf3fe84f8e6c9727dc9e6c970c133185bbdca23532e728e26e585eacbd8d2d45 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/extensions"
    i34702bf27bd058d7b5e871f07e6514ca3dd947d728fc016f03eb36f170b53074 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/attachments/item"
    iaaa7634fa7b22059d9b7249765ef9f412f206b4d530675317f19af70e389cf34 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    iac1bea5f44cb60b84d138bb975b473a16544fe8096eae5c117e8a7ac282abf8f "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/extensions/item"
    iea9fc93fa7aea2d4879fe2764aa490f2a684a91d6a20a48d909c8cfdcfe9d479 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*ia3e61a8f98e0b4ebf997055c903954f21cee726b10079e1c431a38baa7bb1de9.AcceptRequestBuilder) {
    return ia3e61a8f98e0b4ebf997055c903954f21cee726b10079e1c431a38baa7bb1de9.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i068c994d6e3103d7a8c451470fd8946c9a526bf7defe804a3414bec3f058c7c0.AttachmentsRequestBuilder) {
    return i068c994d6e3103d7a8c451470fd8946c9a526bf7defe804a3414bec3f058c7c0.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i34702bf27bd058d7b5e871f07e6514ca3dd947d728fc016f03eb36f170b53074.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i34702bf27bd058d7b5e871f07e6514ca3dd947d728fc016f03eb36f170b53074.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i75b85d086dff2c57d9c90524d349608e34c224e7740e4a9a09c8e3892e4d650b.CalendarRequestBuilder) {
    return i75b85d086dff2c57d9c90524d349608e34c224e7740e4a9a09c8e3892e4d650b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i39f7b9eb2753296d89e1366bb2fa48858540dd472648ec034a0c3a194183a61c.CancelRequestBuilder) {
    return i39f7b9eb2753296d89e1366bb2fa48858540dd472648ec034a0c3a194183a61c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/calendarView/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i732564651bd295ace901245d83f7cee6eecbd49650f303b4d6f278dce8d4196c.DeclineRequestBuilder) {
    return i732564651bd295ace901245d83f7cee6eecbd49650f303b4d6f278dce8d4196c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*ic38a64ab747dedc47441f8456aa350055b7b0ce2810f51c3c45312d8542baab0.DismissReminderRequestBuilder) {
    return ic38a64ab747dedc47441f8456aa350055b7b0ce2810f51c3c45312d8542baab0.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*icf3fe84f8e6c9727dc9e6c970c133185bbdca23532e728e26e585eacbd8d2d45.ExtensionsRequestBuilder) {
    return icf3fe84f8e6c9727dc9e6c970c133185bbdca23532e728e26e585eacbd8d2d45.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iac1bea5f44cb60b84d138bb975b473a16544fe8096eae5c117e8a7ac282abf8f.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iac1bea5f44cb60b84d138bb975b473a16544fe8096eae5c117e8a7ac282abf8f.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i65f4d175fc410ed923f2cd477c6e12f9b6cfed24669da0d6147a1eea2b7353b0.ForwardRequestBuilder) {
    return i65f4d175fc410ed923f2cd477c6e12f9b6cfed24669da0d6147a1eea2b7353b0.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic7e25b5dfc5b0afcb9c1a08774f134a922789857f8a916436156eac48cd979cc.MultiValueExtendedPropertiesRequestBuilder) {
    return ic7e25b5dfc5b0afcb9c1a08774f134a922789857f8a916436156eac48cd979cc.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iea9fc93fa7aea2d4879fe2764aa490f2a684a91d6a20a48d909c8cfdcfe9d479.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iea9fc93fa7aea2d4879fe2764aa490f2a684a91d6a20a48d909c8cfdcfe9d479.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i6d2604c2a16b79e867c6950984ac155a11a951d6935ff922fb3d80299b8aae9c.SingleValueExtendedPropertiesRequestBuilder) {
    return i6d2604c2a16b79e867c6950984ac155a11a951d6935ff922fb3d80299b8aae9c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iaaa7634fa7b22059d9b7249765ef9f412f206b4d530675317f19af70e389cf34.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return iaaa7634fa7b22059d9b7249765ef9f412f206b4d530675317f19af70e389cf34.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i178988a2b9571cc52545bd2a40fe9246c6b63e31f9a8adbe3f7c47021e3ebfd5.SnoozeReminderRequestBuilder) {
    return i178988a2b9571cc52545bd2a40fe9246c6b63e31f9a8adbe3f7c47021e3ebfd5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*iaa82cd4a7429733dee7556dd9022c3f856a8e77926ee4dba3e89fb673d6c269b.TentativelyAcceptRequestBuilder) {
    return iaa82cd4a7429733dee7556dd9022c3f856a8e77926ee4dba3e89fb673d6c269b.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
