package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1233ad0ba006ab96918402fc9449bdbea95cc2732eac86365bde41c84b357d21 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/singlevalueextendedproperties"
    i133098166421eafcbb7d41e4a68f3d29757bcf6688e847c4a7808bf543993aef "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/tentativelyaccept"
    i2ef9251ae81121975f3c9a073ec23b2469d36ed1973d4d03af6e70a2f6fc3c4d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/multivalueextendedproperties"
    i32caceec19579e592d7312d6a514ad4bfa9807198b6c2414606a3703116e102f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/snoozereminder"
    i42df08a14e97e22bb26eddf9cbaa35aa7bf02405a8bbc84fe4fe5b4203cc2ff8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances"
    i53ba2fe10fcac93a89e1ab826cf6d7b8e1e1069ce7a61b20bc3ce310e26333ab "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/accept"
    i6a59555a6a577b7a15b03f87f857fa23cfe6af2eace90222104dbcdc30a0d24d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/attachments"
    i6bf2dbc96674e29a2505ab70f10b8f1735eb7fe1b7640bf2d65e16312c0a071d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/calendar"
    i8c3a0045c4c0ac4da728d4aa37dbba4e6df99e52c194414aedb749b76e669fca "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/forward"
    i921d4dffa29707e7b15bb069fe546aa8d25d138f815d39cdd30ed48b40c27465 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/extensions"
    iac711011c7c99048797a7e80afa2fb7cb707667348b6c9c773d351f73e0a403f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/cancel"
    ie6ebec3d6348227c40f20857d0105bccb0382669c86a08c61932f7f04269ae7a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/dismissreminder"
    ie8cb6c49b7e9e2ac0ad0271442195855cb47449992c78270fe676de14c335081 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/decline"
    i00430d12f816c0c6997639410f009cc5f413c588c399506a408cafb1d4a794ec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/multivalueextendedproperties/item"
    i1ec50c12270bd7bd6ab657694752ebec839c5bc0a607be433717875d65bff30e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/extensions/item"
    i9be5b8dcebdfca6e7ace3db98d602d80c97a8f974e4125c75bacc246c066e21e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/singlevalueextendedproperties/item"
    i9fe44f8859760d86bd88baff7a59fc2a46f6264caa52b2adf30c8d7f4ae24672 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/attachments/item"
    id60fc1ecd13479f120a7edbab2577fd0b8fff9ebc834f6c7d280345dd5f73d2e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item"
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
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
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
func (m *EventItemRequestBuilder) Accept()(*i53ba2fe10fcac93a89e1ab826cf6d7b8e1e1069ce7a61b20bc3ce310e26333ab.AcceptRequestBuilder) {
    return i53ba2fe10fcac93a89e1ab826cf6d7b8e1e1069ce7a61b20bc3ce310e26333ab.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i6a59555a6a577b7a15b03f87f857fa23cfe6af2eace90222104dbcdc30a0d24d.AttachmentsRequestBuilder) {
    return i6a59555a6a577b7a15b03f87f857fa23cfe6af2eace90222104dbcdc30a0d24d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i9fe44f8859760d86bd88baff7a59fc2a46f6264caa52b2adf30c8d7f4ae24672.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i9fe44f8859760d86bd88baff7a59fc2a46f6264caa52b2adf30c8d7f4ae24672.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i6bf2dbc96674e29a2505ab70f10b8f1735eb7fe1b7640bf2d65e16312c0a071d.CalendarRequestBuilder) {
    return i6bf2dbc96674e29a2505ab70f10b8f1735eb7fe1b7640bf2d65e16312c0a071d.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*iac711011c7c99048797a7e80afa2fb7cb707667348b6c9c773d351f73e0a403f.CancelRequestBuilder) {
    return iac711011c7c99048797a7e80afa2fb7cb707667348b6c9c773d351f73e0a403f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for groups
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
// CreatePatchRequestInformation update the navigation property calendarView in groups
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
func (m *EventItemRequestBuilder) Decline()(*ie8cb6c49b7e9e2ac0ad0271442195855cb47449992c78270fe676de14c335081.DeclineRequestBuilder) {
    return ie8cb6c49b7e9e2ac0ad0271442195855cb47449992c78270fe676de14c335081.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*ie6ebec3d6348227c40f20857d0105bccb0382669c86a08c61932f7f04269ae7a.DismissReminderRequestBuilder) {
    return ie6ebec3d6348227c40f20857d0105bccb0382669c86a08c61932f7f04269ae7a.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i921d4dffa29707e7b15bb069fe546aa8d25d138f815d39cdd30ed48b40c27465.ExtensionsRequestBuilder) {
    return i921d4dffa29707e7b15bb069fe546aa8d25d138f815d39cdd30ed48b40c27465.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i1ec50c12270bd7bd6ab657694752ebec839c5bc0a607be433717875d65bff30e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i1ec50c12270bd7bd6ab657694752ebec839c5bc0a607be433717875d65bff30e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i8c3a0045c4c0ac4da728d4aa37dbba4e6df99e52c194414aedb749b76e669fca.ForwardRequestBuilder) {
    return i8c3a0045c4c0ac4da728d4aa37dbba4e6df99e52c194414aedb749b76e669fca.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*i42df08a14e97e22bb26eddf9cbaa35aa7bf02405a8bbc84fe4fe5b4203cc2ff8.InstancesRequestBuilder) {
    return i42df08a14e97e22bb26eddf9cbaa35aa7bf02405a8bbc84fe4fe5b4203cc2ff8.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*id60fc1ecd13479f120a7edbab2577fd0b8fff9ebc834f6c7d280345dd5f73d2e.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return id60fc1ecd13479f120a7edbab2577fd0b8fff9ebc834f6c7d280345dd5f73d2e.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i2ef9251ae81121975f3c9a073ec23b2469d36ed1973d4d03af6e70a2f6fc3c4d.MultiValueExtendedPropertiesRequestBuilder) {
    return i2ef9251ae81121975f3c9a073ec23b2469d36ed1973d4d03af6e70a2f6fc3c4d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i00430d12f816c0c6997639410f009cc5f413c588c399506a408cafb1d4a794ec.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i00430d12f816c0c6997639410f009cc5f413c588c399506a408cafb1d4a794ec.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i1233ad0ba006ab96918402fc9449bdbea95cc2732eac86365bde41c84b357d21.SingleValueExtendedPropertiesRequestBuilder) {
    return i1233ad0ba006ab96918402fc9449bdbea95cc2732eac86365bde41c84b357d21.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9be5b8dcebdfca6e7ace3db98d602d80c97a8f974e4125c75bacc246c066e21e.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i9be5b8dcebdfca6e7ace3db98d602d80c97a8f974e4125c75bacc246c066e21e.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i32caceec19579e592d7312d6a514ad4bfa9807198b6c2414606a3703116e102f.SnoozeReminderRequestBuilder) {
    return i32caceec19579e592d7312d6a514ad4bfa9807198b6c2414606a3703116e102f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i133098166421eafcbb7d41e4a68f3d29757bcf6688e847c4a7808bf543993aef.TentativelyAcceptRequestBuilder) {
    return i133098166421eafcbb7d41e4a68f3d29757bcf6688e847c4a7808bf543993aef.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
