package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i08263b120d4a64d516da1a9a009102f68f1ea3b72038e62049453064f3520daa "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/getschedule"
    i5b92fe7490f3adfa9b9eb8b7b27fcf837f66e0bf11a120b311cab89a7ab5ccb4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/multivalueextendedproperties"
    i636df8302013a50dd4f709806f91e209c2a39540a8c8acb8ff6d284173bc9611 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events"
    i790abfff1579d62421ba888c419e99fcdb7f3342ca621795f1d4878bdf41107d "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarpermissions"
    i8d0a3a589be44f5a571e77f994200094b296ec5e5a51363c20dd536fdd042980 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/singlevalueextendedproperties"
    ib60e8318b810b2b05ebbda7d142451b6e905523a0d513153c156b9da64697fee "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview"
    if635722f2fe253209876c19e8ff9c2f5ed8ce31cb7fb3a6ef7c95fc4fa600390 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/allowedcalendarsharingroleswithuser"
    i0159be40cebef16f606f3683bed2c527ac757b7c0595a04fd5a7980fc524e986 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/multivalueextendedproperties/item"
    i5eeae220666e7414f299626922168ed676e16d02eb4326b487e9a593c4ea62b5 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarpermissions/item"
    i72d2962cc7d6783d327cd6ac936abb7f3a115f64393795d9f16cbed9c34fbc1a "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item"
    i7441dd38f6817f39599227c0b3a93087e04dc084baa163401adef30ac3874560 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/singlevalueextendedproperties/item"
    if4d2b8a420506859db7a9f0c9c655e7bda330951aa669e4ab40546cf6bb9ce22 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item"
)

// CalendarItemRequestBuilder provides operations to manage the calendars property of the microsoft.graph.user entity.
type CalendarItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CalendarItemRequestBuilderDeleteOptions options for Delete
type CalendarItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CalendarItemRequestBuilderGetOptions options for Get
type CalendarItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CalendarItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CalendarItemRequestBuilderGetQueryParameters the user's calendars. Read-only. Nullable.
type CalendarItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// CalendarItemRequestBuilderPatchOptions options for Patch
type CalendarItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendarable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *CalendarItemRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*if635722f2fe253209876c19e8ff9c2f5ed8ce31cb7fb3a6ef7c95fc4fa600390.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return if635722f2fe253209876c19e8ff9c2f5ed8ce31cb7fb3a6ef7c95fc4fa600390.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarItemRequestBuilder) CalendarPermissions()(*i790abfff1579d62421ba888c419e99fcdb7f3342ca621795f1d4878bdf41107d.CalendarPermissionsRequestBuilder) {
    return i790abfff1579d62421ba888c419e99fcdb7f3342ca621795f1d4878bdf41107d.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.calendarPermissions.item collection
func (m *CalendarItemRequestBuilder) CalendarPermissionsById(id string)(*i5eeae220666e7414f299626922168ed676e16d02eb4326b487e9a593c4ea62b5.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i5eeae220666e7414f299626922168ed676e16d02eb4326b487e9a593c4ea62b5.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarItemRequestBuilder) CalendarView()(*ib60e8318b810b2b05ebbda7d142451b6e905523a0d513153c156b9da64697fee.CalendarViewRequestBuilder) {
    return ib60e8318b810b2b05ebbda7d142451b6e905523a0d513153c156b9da64697fee.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.calendarView.item collection
func (m *CalendarItemRequestBuilder) CalendarViewById(id string)(*if4d2b8a420506859db7a9f0c9c655e7bda330951aa669e4ab40546cf6bb9ce22.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return if4d2b8a420506859db7a9f0c9c655e7bda330951aa669e4ab40546cf6bb9ce22.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarItemRequestBuilderInternal instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarItemRequestBuilder) {
    m := &CalendarItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarItemRequestBuilder instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) CreateDeleteRequestInformation(options *CalendarItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the user's calendars. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) CreateGetRequestInformation(options *CalendarItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) CreatePatchRequestInformation(options *CalendarItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) Delete(options *CalendarItemRequestBuilderDeleteOptions)(error) {
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
func (m *CalendarItemRequestBuilder) Events()(*i636df8302013a50dd4f709806f91e209c2a39540a8c8acb8ff6d284173bc9611.EventsRequestBuilder) {
    return i636df8302013a50dd4f709806f91e209c2a39540a8c8acb8ff6d284173bc9611.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item collection
func (m *CalendarItemRequestBuilder) EventsById(id string)(*i72d2962cc7d6783d327cd6ac936abb7f3a115f64393795d9f16cbed9c34fbc1a.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i72d2962cc7d6783d327cd6ac936abb7f3a115f64393795d9f16cbed9c34fbc1a.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the user's calendars. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) Get(options *CalendarItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendarable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateCalendarFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendarable), nil
}
func (m *CalendarItemRequestBuilder) GetSchedule()(*i08263b120d4a64d516da1a9a009102f68f1ea3b72038e62049453064f3520daa.GetScheduleRequestBuilder) {
    return i08263b120d4a64d516da1a9a009102f68f1ea3b72038e62049453064f3520daa.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarItemRequestBuilder) MultiValueExtendedProperties()(*i5b92fe7490f3adfa9b9eb8b7b27fcf837f66e0bf11a120b311cab89a7ab5ccb4.MultiValueExtendedPropertiesRequestBuilder) {
    return i5b92fe7490f3adfa9b9eb8b7b27fcf837f66e0bf11a120b311cab89a7ab5ccb4.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0159be40cebef16f606f3683bed2c527ac757b7c0595a04fd5a7980fc524e986.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0159be40cebef16f606f3683bed2c527ac757b7c0595a04fd5a7980fc524e986.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) Patch(options *CalendarItemRequestBuilderPatchOptions)(error) {
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
func (m *CalendarItemRequestBuilder) SingleValueExtendedProperties()(*i8d0a3a589be44f5a571e77f994200094b296ec5e5a51363c20dd536fdd042980.SingleValueExtendedPropertiesRequestBuilder) {
    return i8d0a3a589be44f5a571e77f994200094b296ec5e5a51363c20dd536fdd042980.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7441dd38f6817f39599227c0b3a93087e04dc084baa163401adef30ac3874560.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i7441dd38f6817f39599227c0b3a93087e04dc084baa163401adef30ac3874560.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
