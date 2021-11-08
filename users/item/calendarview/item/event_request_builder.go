package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0f91cc3524d407ce127cf52c636ea3d70741ec93c24f733e934c355ef01e9673 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/forward"
    i208443f24ffa0e9d495211e2061b0c16c1cc4e5aebf7a138ffbe44f915a17e0b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar"
    i2122e565a8aa4c89017cd0e7f55859c05eae01b9b38ffdc65ae46e71e9707ee4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/accept"
    i341a1e7e328f323fd2b5b915fd2dbad7995b2bf8f3418f1e5b5e4ac55bb815c4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/decline"
    i3bb2079b15e2279e42f683d8f17113622024239edddadace92a48a4f793d9522 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/tentativelyaccept"
    i85b3ecb1df5dabd5d3ad4c4e6a54e2d1660771b26010ff83b38fa7db3649b996 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/dismissreminder"
    i9407c9bcb1ee71f799ed9d8c791c167b82daf93c1270274e376f805cc9f76d5f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/extensions"
    i994428626bd985a642668dee4c95629ec8caffe540062bae0249e43b54dfc697 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances"
    i9eddc52577a5555a33aafc56b6ec4c751b6c5272a204d1db9d90ab5f60ff902e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/singlevalueextendedproperties"
    ia5d34a53a5d690fc00e53f31e3ea500bbc666ac906da17540530fa20f6d74e70 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/attachments"
    iad6745b576f0cdfa153c080e6ed44bd19ba5fa213eacc4c29dc1e699984f0e3c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/multivalueextendedproperties"
    ibcb96cadb1d3953182dae076a8d0d2f4d74b8522757c415c12ef16502beb9e53 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/cancel"
    ie4ee11a97dc68135822d054e4649c9f9aa8b85b739ea56737a94898a6de75e4f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/snoozereminder"
    i36231df696b83109272fd3bb1ff522ae5b1c7890f3c08c8d848fe767a18cd47a "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/extensions/item"
    i4b0486fcdf6c612baf1d1ebbb06ea4f753f858e537fec0fc098234b5a3bf66f9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/attachments/item"
    i5df7d90865ed74cfa89b9670f8906bc275af7e716c852c939338962b443c4fdb "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item"
    i7df1508e513db666e7370befb3f751526bf7cdf8a085672e191438aec048b505 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/multivalueextendedproperties/item"
    i8bde4d3140c0c2f25ed25759ec43c0e794a2a480b2fc8d25b5f99b0a48d2c100 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/singlevalueextendedproperties/item"
)

// Builds and executes requests for operations under \users\{user-id}\calendarView\{event-id}
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
// The calendar view for the calendar. Read-only. Nullable.
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
func (m *EventRequestBuilder) Accept()(*i2122e565a8aa4c89017cd0e7f55859c05eae01b9b38ffdc65ae46e71e9707ee4.AcceptRequestBuilder) {
    return i2122e565a8aa4c89017cd0e7f55859c05eae01b9b38ffdc65ae46e71e9707ee4.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ia5d34a53a5d690fc00e53f31e3ea500bbc666ac906da17540530fa20f6d74e70.AttachmentsRequestBuilder) {
    return ia5d34a53a5d690fc00e53f31e3ea500bbc666ac906da17540530fa20f6d74e70.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.calendarView.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) AttachmentsById(id string)(*i4b0486fcdf6c612baf1d1ebbb06ea4f753f858e537fec0fc098234b5a3bf66f9.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i4b0486fcdf6c612baf1d1ebbb06ea4f753f858e537fec0fc098234b5a3bf66f9.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i208443f24ffa0e9d495211e2061b0c16c1cc4e5aebf7a138ffbe44f915a17e0b.CalendarRequestBuilder) {
    return i208443f24ffa0e9d495211e2061b0c16c1cc4e5aebf7a138ffbe44f915a17e0b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*ibcb96cadb1d3953182dae076a8d0d2f4d74b8522757c415c12ef16502beb9e53.CancelRequestBuilder) {
    return ibcb96cadb1d3953182dae076a8d0d2f4d74b8522757c415c12ef16502beb9e53.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarView/{event_id}{?startDateTime,endDateTime,select}";
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
// The calendar view for the calendar. Read-only. Nullable.
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
// The calendar view for the calendar. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) CreateGetRequestInformation(options *EventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The calendar view for the calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) Decline()(*i341a1e7e328f323fd2b5b915fd2dbad7995b2bf8f3418f1e5b5e4ac55bb815c4.DeclineRequestBuilder) {
    return i341a1e7e328f323fd2b5b915fd2dbad7995b2bf8f3418f1e5b5e4ac55bb815c4.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) DismissReminder()(*i85b3ecb1df5dabd5d3ad4c4e6a54e2d1660771b26010ff83b38fa7db3649b996.DismissReminderRequestBuilder) {
    return i85b3ecb1df5dabd5d3ad4c4e6a54e2d1660771b26010ff83b38fa7db3649b996.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i9407c9bcb1ee71f799ed9d8c791c167b82daf93c1270274e376f805cc9f76d5f.ExtensionsRequestBuilder) {
    return i9407c9bcb1ee71f799ed9d8c791c167b82daf93c1270274e376f805cc9f76d5f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.calendarView.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) ExtensionsById(id string)(*i36231df696b83109272fd3bb1ff522ae5b1c7890f3c08c8d848fe767a18cd47a.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i36231df696b83109272fd3bb1ff522ae5b1c7890f3c08c8d848fe767a18cd47a.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i0f91cc3524d407ce127cf52c636ea3d70741ec93c24f733e934c355ef01e9673.ForwardRequestBuilder) {
    return i0f91cc3524d407ce127cf52c636ea3d70741ec93c24f733e934c355ef01e9673.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) Instances()(*i994428626bd985a642668dee4c95629ec8caffe540062bae0249e43b54dfc697.InstancesRequestBuilder) {
    return i994428626bd985a642668dee4c95629ec8caffe540062bae0249e43b54dfc697.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.calendarView.item.instances.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) InstancesById(id string)(*i5df7d90865ed74cfa89b9670f8906bc275af7e716c852c939338962b443c4fdb.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i5df7d90865ed74cfa89b9670f8906bc275af7e716c852c939338962b443c4fdb.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*iad6745b576f0cdfa153c080e6ed44bd19ba5fa213eacc4c29dc1e699984f0e3c.MultiValueExtendedPropertiesRequestBuilder) {
    return iad6745b576f0cdfa153c080e6ed44bd19ba5fa213eacc4c29dc1e699984f0e3c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.calendarView.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7df1508e513db666e7370befb3f751526bf7cdf8a085672e191438aec048b505.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i7df1508e513db666e7370befb3f751526bf7cdf8a085672e191438aec048b505.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendar view for the calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i9eddc52577a5555a33aafc56b6ec4c751b6c5272a204d1db9d90ab5f60ff902e.SingleValueExtendedPropertiesRequestBuilder) {
    return i9eddc52577a5555a33aafc56b6ec4c751b6c5272a204d1db9d90ab5f60ff902e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.calendarView.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8bde4d3140c0c2f25ed25759ec43c0e794a2a480b2fc8d25b5f99b0a48d2c100.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i8bde4d3140c0c2f25ed25759ec43c0e794a2a480b2fc8d25b5f99b0a48d2c100.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*ie4ee11a97dc68135822d054e4649c9f9aa8b85b739ea56737a94898a6de75e4f.SnoozeReminderRequestBuilder) {
    return ie4ee11a97dc68135822d054e4649c9f9aa8b85b739ea56737a94898a6de75e4f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i3bb2079b15e2279e42f683d8f17113622024239edddadace92a48a4f793d9522.TentativelyAcceptRequestBuilder) {
    return i3bb2079b15e2279e42f683d8f17113622024239edddadace92a48a4f793d9522.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
