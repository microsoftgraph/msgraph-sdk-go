package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i24e24e2af499471f8e4f73b6d94d99165b1e84d2131623dcc1a383199e8f050d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/singlevalueextendedproperties"
    i2d1fe0bd4017d6a5f945a9ceefc6fde8147a060a766d7001c9b0a50c6ba6ee00 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/multivalueextendedproperties"
    i44dd09532dcb7a7d8ed2a9efa5056d3e20b725877822e9b1c509e378c8db9c90 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/snoozereminder"
    i493deeccf2b51ea98ffd5d94edc505e26b0559809a7e344b863c50a129ced7fe "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/tentativelyaccept"
    i555550f048433b0f866e19c810c2cac5a25998a168b6c16dc1a7f6a70665fd92 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar"
    i61b71d3da7a06d0f54197cc290f668d03c0ff14f20103e48cab218d31dab0913 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/extensions"
    i6c90a1bd7a4f0cece06e6b34e378ee70ca26356395f3e3815e647b33c612c4eb "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/forward"
    i6d4c88f7616ab1a1de0f125a057a6e9383f4b21373f7c695fd9c6db145fa1754 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances"
    i7a68ed9d6c87ce91e310c0b60439fa878a2f3f2ce780223e121fe57306de4cf0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/attachments"
    i89ff886574ba825b3220369ac4d9bdb3e03e16efbd810974d1a8b0c6a7d97cf8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/decline"
    i945af8cb490de0be43df42f623102f1a698bf30d1e58d7b44500a64140273ec6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/accept"
    idf3581de0ced5dd508d8f36e7490c56e109571dd437697fe9c1edb7bce3dac89 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/dismissreminder"
    ie208e488c4a77bd04b8849d39ad4a5b1c796289043cd257b42ba1a96585cd84e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/cancel"
    i243da75f0745ff07c19b77e029d0409d2627daa65cc8fff1c986fb76b2b765a9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/attachments/item"
    i64313c875409e11e8ad455ef0785e8f2d79c5d511bf6a34bdeec84f3da8c6795 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/singlevalueextendedproperties/item"
    i66a62ea9b41daca5ad35a026f8fdd67e1aa3e96483f9cb8758db312e258ae78f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/instances/item"
    i71e82e42e7ad7a048127c6e49b504d33eb9ea367867057c7a9bf3cb88ef2c551 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/extensions/item"
    ib73431f50c914fa465532d76bca0db38ab141314b5872ddb060382316b1988dd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/multivalueextendedproperties/item"
)

// Builds and executes requests for operations under \groups\{group-id}\calendarView\{event-id}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type EventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type EventRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The calendar view for the calendar. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select_escaped []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
}
// Options for Patch
type EventRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventRequestBuilder) Accept()(*i945af8cb490de0be43df42f623102f1a698bf30d1e58d7b44500a64140273ec6.AcceptRequestBuilder) {
    return i945af8cb490de0be43df42f623102f1a698bf30d1e58d7b44500a64140273ec6.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i7a68ed9d6c87ce91e310c0b60439fa878a2f3f2ce780223e121fe57306de4cf0.AttachmentsRequestBuilder) {
    return i7a68ed9d6c87ce91e310c0b60439fa878a2f3f2ce780223e121fe57306de4cf0.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendarView.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) AttachmentsById(id string)(*i243da75f0745ff07c19b77e029d0409d2627daa65cc8fff1c986fb76b2b765a9.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i243da75f0745ff07c19b77e029d0409d2627daa65cc8fff1c986fb76b2b765a9.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i555550f048433b0f866e19c810c2cac5a25998a168b6c16dc1a7f6a70665fd92.CalendarRequestBuilder) {
    return i555550f048433b0f866e19c810c2cac5a25998a168b6c16dc1a7f6a70665fd92.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*ie208e488c4a77bd04b8849d39ad4a5b1c796289043cd257b42ba1a96585cd84e.CancelRequestBuilder) {
    return ie208e488c4a77bd04b8849d39ad4a5b1c796289043cd257b42ba1a96585cd84e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendarView/{event_id}{?startDateTime,endDateTime,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// The calendar view for the calendar. Read-only.
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) CreateDeleteRequestInformation(options *EventRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The calendar view for the calendar. Read-only.
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) CreateGetRequestInformation(options *EventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
// The calendar view for the calendar. Read-only.
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) CreatePatchRequestInformation(options *EventRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*i89ff886574ba825b3220369ac4d9bdb3e03e16efbd810974d1a8b0c6a7d97cf8.DeclineRequestBuilder) {
    return i89ff886574ba825b3220369ac4d9bdb3e03e16efbd810974d1a8b0c6a7d97cf8.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Read-only.
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) Delete(options *EventRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) DismissReminder()(*idf3581de0ced5dd508d8f36e7490c56e109571dd437697fe9c1edb7bce3dac89.DismissReminderRequestBuilder) {
    return idf3581de0ced5dd508d8f36e7490c56e109571dd437697fe9c1edb7bce3dac89.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i61b71d3da7a06d0f54197cc290f668d03c0ff14f20103e48cab218d31dab0913.ExtensionsRequestBuilder) {
    return i61b71d3da7a06d0f54197cc290f668d03c0ff14f20103e48cab218d31dab0913.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendarView.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) ExtensionsById(id string)(*i71e82e42e7ad7a048127c6e49b504d33eb9ea367867057c7a9bf3cb88ef2c551.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i71e82e42e7ad7a048127c6e49b504d33eb9ea367867057c7a9bf3cb88ef2c551.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i6c90a1bd7a4f0cece06e6b34e378ee70ca26356395f3e3815e647b33c612c4eb.ForwardRequestBuilder) {
    return i6c90a1bd7a4f0cece06e6b34e378ee70ca26356395f3e3815e647b33c612c4eb.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Read-only.
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) Get(options *EventRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event), nil
}
func (m *EventRequestBuilder) Instances()(*i6d4c88f7616ab1a1de0f125a057a6e9383f4b21373f7c695fd9c6db145fa1754.InstancesRequestBuilder) {
    return i6d4c88f7616ab1a1de0f125a057a6e9383f4b21373f7c695fd9c6db145fa1754.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendarView.item.instances.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) InstancesById(id string)(*i66a62ea9b41daca5ad35a026f8fdd67e1aa3e96483f9cb8758db312e258ae78f.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i66a62ea9b41daca5ad35a026f8fdd67e1aa3e96483f9cb8758db312e258ae78f.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i2d1fe0bd4017d6a5f945a9ceefc6fde8147a060a766d7001c9b0a50c6ba6ee00.MultiValueExtendedPropertiesRequestBuilder) {
    return i2d1fe0bd4017d6a5f945a9ceefc6fde8147a060a766d7001c9b0a50c6ba6ee00.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendarView.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib73431f50c914fa465532d76bca0db38ab141314b5872ddb060382316b1988dd.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib73431f50c914fa465532d76bca0db38ab141314b5872ddb060382316b1988dd.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendar view for the calendar. Read-only.
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) Patch(options *EventRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i24e24e2af499471f8e4f73b6d94d99165b1e84d2131623dcc1a383199e8f050d.SingleValueExtendedPropertiesRequestBuilder) {
    return i24e24e2af499471f8e4f73b6d94d99165b1e84d2131623dcc1a383199e8f050d.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendarView.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i64313c875409e11e8ad455ef0785e8f2d79c5d511bf6a34bdeec84f3da8c6795.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i64313c875409e11e8ad455ef0785e8f2d79c5d511bf6a34bdeec84f3da8c6795.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i44dd09532dcb7a7d8ed2a9efa5056d3e20b725877822e9b1c509e378c8db9c90.SnoozeReminderRequestBuilder) {
    return i44dd09532dcb7a7d8ed2a9efa5056d3e20b725877822e9b1c509e378c8db9c90.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i493deeccf2b51ea98ffd5d94edc505e26b0559809a7e344b863c50a129ced7fe.TentativelyAcceptRequestBuilder) {
    return i493deeccf2b51ea98ffd5d94edc505e26b0559809a7e344b863c50a129ced7fe.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
