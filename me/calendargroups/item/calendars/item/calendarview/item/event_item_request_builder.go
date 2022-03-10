package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i034f075a9742087cc62fa2fa3c2d823b0c34579b549387b0217d091d4a084aed "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    i0d1833fafe91c99d46532bdcc396354079ec6601492e051fcf683e98f02b4aee "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/decline"
    i322baabcc84e3997da553dda536125328af71f44c1c87052ffeafb14a6877592 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/cancel"
    i42aeafc43ff4f3886ddd2c06a4c17d4ba4cc3b7fd3216d6c353debc41b3d4d14 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/calendar"
    i554abe6eb0940df39daf173fba7aa2976e43e0e1cde74ba38fb1ad212c524b1e "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/dismissreminder"
    i7e9c02a874a50f511f5d63885525677f008606e724bca4866861a307c5f4886a "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/extensions"
    i85e1eadeddf7ec2c307967b39d84d66cd802ef3edd7cea67e3fcdd2e9efae123 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/attachments"
    i8bd23f6c7e0c42bea79bd661447dc93d8e0ada59e8bba7ad08a5783f58f6ad68 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties"
    i94dcdd3851da0f02fa640225ba331bea6617b62727e0061c6bb8e9ab386b8629 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/accept"
    ibcb3c8ba4cd7beda528fefdc4942827a7ca6e5c6b961928da585dfc095291fb4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/forward"
    ic66f697cfef967cd4effb4c9d92b29e287a4def44f9b2aea82d334789318299f "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/snoozereminder"
    ie4a846e389b081fa82f31b8581ae8dfed0401f4a81a686e4cc965b82d9606851 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/tentativelyaccept"
    ieb9e6dcb9bcea88f684ee7112127bd4d1df6deb986fab653d792f9b1389cc795 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances"
    i4e4fe41f8ec4b0f84b1c21baa5931909bd57fcb9768d3be13b17583642944ef8 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    i62e0076ba3224fb44507e4ef678645cbc8c53f273bb42881cc65121616fbd8f0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/attachments/item"
    ie57987dfe2c89ca06fcc0fc5f474e14183f8c3cec8825b4b733c1da6a5d7ea6a "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/extensions/item"
    if26639b4bc627d0cd7049c7f814bf3f0458856e94fc4fce0c5d95011495f7257 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item"
    if7192229b65a228ea27ede8842c1b7d3638ac8120f07a51c43c8174a1920dd14 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
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
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Accept()(*i94dcdd3851da0f02fa640225ba331bea6617b62727e0061c6bb8e9ab386b8629.AcceptRequestBuilder) {
    return i94dcdd3851da0f02fa640225ba331bea6617b62727e0061c6bb8e9ab386b8629.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i85e1eadeddf7ec2c307967b39d84d66cd802ef3edd7cea67e3fcdd2e9efae123.AttachmentsRequestBuilder) {
    return i85e1eadeddf7ec2c307967b39d84d66cd802ef3edd7cea67e3fcdd2e9efae123.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i62e0076ba3224fb44507e4ef678645cbc8c53f273bb42881cc65121616fbd8f0.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i62e0076ba3224fb44507e4ef678645cbc8c53f273bb42881cc65121616fbd8f0.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i42aeafc43ff4f3886ddd2c06a4c17d4ba4cc3b7fd3216d6c353debc41b3d4d14.CalendarRequestBuilder) {
    return i42aeafc43ff4f3886ddd2c06a4c17d4ba4cc3b7fd3216d6c353debc41b3d4d14.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i322baabcc84e3997da553dda536125328af71f44c1c87052ffeafb14a6877592.CancelRequestBuilder) {
    return i322baabcc84e3997da553dda536125328af71f44c1c87052ffeafb14a6877592.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/calendarView/{event_id}{?select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for me
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation update the navigation property calendarView in me
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
func (m *EventItemRequestBuilder) Decline()(*i0d1833fafe91c99d46532bdcc396354079ec6601492e051fcf683e98f02b4aee.DeclineRequestBuilder) {
    return i0d1833fafe91c99d46532bdcc396354079ec6601492e051fcf683e98f02b4aee.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i554abe6eb0940df39daf173fba7aa2976e43e0e1cde74ba38fb1ad212c524b1e.DismissReminderRequestBuilder) {
    return i554abe6eb0940df39daf173fba7aa2976e43e0e1cde74ba38fb1ad212c524b1e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i7e9c02a874a50f511f5d63885525677f008606e724bca4866861a307c5f4886a.ExtensionsRequestBuilder) {
    return i7e9c02a874a50f511f5d63885525677f008606e724bca4866861a307c5f4886a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ie57987dfe2c89ca06fcc0fc5f474e14183f8c3cec8825b4b733c1da6a5d7ea6a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ie57987dfe2c89ca06fcc0fc5f474e14183f8c3cec8825b4b733c1da6a5d7ea6a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*ibcb3c8ba4cd7beda528fefdc4942827a7ca6e5c6b961928da585dfc095291fb4.ForwardRequestBuilder) {
    return ibcb3c8ba4cd7beda528fefdc4942827a7ca6e5c6b961928da585dfc095291fb4.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*ieb9e6dcb9bcea88f684ee7112127bd4d1df6deb986fab653d792f9b1389cc795.InstancesRequestBuilder) {
    return ieb9e6dcb9bcea88f684ee7112127bd4d1df6deb986fab653d792f9b1389cc795.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*if26639b4bc627d0cd7049c7f814bf3f0458856e94fc4fce0c5d95011495f7257.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return if26639b4bc627d0cd7049c7f814bf3f0458856e94fc4fce0c5d95011495f7257.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i8bd23f6c7e0c42bea79bd661447dc93d8e0ada59e8bba7ad08a5783f58f6ad68.MultiValueExtendedPropertiesRequestBuilder) {
    return i8bd23f6c7e0c42bea79bd661447dc93d8e0ada59e8bba7ad08a5783f58f6ad68.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4e4fe41f8ec4b0f84b1c21baa5931909bd57fcb9768d3be13b17583642944ef8.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i4e4fe41f8ec4b0f84b1c21baa5931909bd57fcb9768d3be13b17583642944ef8.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i034f075a9742087cc62fa2fa3c2d823b0c34579b549387b0217d091d4a084aed.SingleValueExtendedPropertiesRequestBuilder) {
    return i034f075a9742087cc62fa2fa3c2d823b0c34579b549387b0217d091d4a084aed.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if7192229b65a228ea27ede8842c1b7d3638ac8120f07a51c43c8174a1920dd14.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if7192229b65a228ea27ede8842c1b7d3638ac8120f07a51c43c8174a1920dd14.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ic66f697cfef967cd4effb4c9d92b29e287a4def44f9b2aea82d334789318299f.SnoozeReminderRequestBuilder) {
    return ic66f697cfef967cd4effb4c9d92b29e287a4def44f9b2aea82d334789318299f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ie4a846e389b081fa82f31b8581ae8dfed0401f4a81a686e4cc965b82d9606851.TentativelyAcceptRequestBuilder) {
    return ie4a846e389b081fa82f31b8581ae8dfed0401f4a81a686e4cc965b82d9606851.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
