package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0853bc5e78a4ac665d8aa887531e364debb1ba91c3a1bc736be9331cc8db0b52 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/singlevalueextendedproperties"
    i25a4d487ad09f2eb45ba28744d3fe919b7b5dbbf1c645ddd9776a9b357d0ff96 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview"
    i3abccbd6b1092e6246d7bc4dd7dcac1b6938ee6bf024fb4eccaaeecfb29ff1b6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events"
    i3e09978bd77ced5481016e4684368f72edfa92ef709a19504e9258cf8865bfc6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/allowedcalendarsharingroleswithuser"
    i847a76068ecb832435c51c746fba99e17f01da45f86bc0ea0f29b5a1d42ae5b3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/multivalueextendedproperties"
    ia0a09be4f5c1f5b3b99aa3064830cdc6c730ceea9bdb4dbc6934a7d5f0ffec6f "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/getschedule"
    iaac11e3c8bf2dd29e59be19f34577e748fb426a2bb9b94aae4be18a9b2f6474a "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarpermissions"
    i67198c8bd92c3d2fcf155bff41150fc55dc916147fb9b1eb9ca2ccf6b9d1ed6e "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item"
    i857bdaf62ca4819c99b59159f0888909252fbffeba2fa7050dfc75ce2c6b598d "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/singlevalueextendedproperties/item"
    i8783c6ceaf93b1654906d93bb1e4267cb1060f0bcdbc2d302e4f64bfac674b76 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item"
    i8b359363f8c83629e3e25072f613c3d3d807cc671fe5472a27f65565ce218c26 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarpermissions/item"
    idefd7aa8b66a92e862195e44b4d12afbb8cb1f08136af3a806efad5c689470d7 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/multivalueextendedproperties/item"
)

// Builds and executes requests for operations under \me\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}
type CalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type CalendarRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type CalendarRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CalendarRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The calendars in the calendar group. Navigation property. Read-only. Nullable.
type CalendarRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type CalendarRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Builds and executes requests for operations under \me\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
// Parameters:
//  - User : Usage: User={User}
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i3e09978bd77ced5481016e4684368f72edfa92ef709a19504e9258cf8865bfc6.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i3e09978bd77ced5481016e4684368f72edfa92ef709a19504e9258cf8865bfc6.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*iaac11e3c8bf2dd29e59be19f34577e748fb426a2bb9b94aae4be18a9b2f6474a.CalendarPermissionsRequestBuilder) {
    return iaac11e3c8bf2dd29e59be19f34577e748fb426a2bb9b94aae4be18a9b2f6474a.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendarGroups.item.calendars.item.calendarPermissions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i8b359363f8c83629e3e25072f613c3d3d807cc671fe5472a27f65565ce218c26.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i8b359363f8c83629e3e25072f613c3d3d807cc671fe5472a27f65565ce218c26.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i25a4d487ad09f2eb45ba28744d3fe919b7b5dbbf1c645ddd9776a9b357d0ff96.CalendarViewRequestBuilder) {
    return i25a4d487ad09f2eb45ba28744d3fe919b7b5dbbf1c645ddd9776a9b357d0ff96.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendarGroups.item.calendars.item.calendarView.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i8783c6ceaf93b1654906d93bb1e4267cb1060f0bcdbc2d302e4f64bfac674b76.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i8783c6ceaf93b1654906d93bb1e4267cb1060f0bcdbc2d302e4f64bfac674b76.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new CalendarRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new CalendarRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// The calendars in the calendar group. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *CalendarRequestBuilder) CreateDeleteRequestInformation(options *CalendarRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The calendars in the calendar group. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *CalendarRequestBuilder) CreateGetRequestInformation(options *CalendarRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The calendars in the calendar group. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(options *CalendarRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The calendars in the calendar group. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *CalendarRequestBuilder) Delete(options *CalendarRequestBuilderDeleteOptions)(error) {
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
func (m *CalendarRequestBuilder) Events()(*i3abccbd6b1092e6246d7bc4dd7dcac1b6938ee6bf024fb4eccaaeecfb29ff1b6.EventsRequestBuilder) {
    return i3abccbd6b1092e6246d7bc4dd7dcac1b6938ee6bf024fb4eccaaeecfb29ff1b6.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendarGroups.item.calendars.item.events.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) EventsById(id string)(*i67198c8bd92c3d2fcf155bff41150fc55dc916147fb9b1eb9ca2ccf6b9d1ed6e.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i67198c8bd92c3d2fcf155bff41150fc55dc916147fb9b1eb9ca2ccf6b9d1ed6e.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendars in the calendar group. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *CalendarRequestBuilder) Get(options *CalendarRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewCalendar() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar), nil
}
func (m *CalendarRequestBuilder) GetSchedule()(*ia0a09be4f5c1f5b3b99aa3064830cdc6c730ceea9bdb4dbc6934a7d5f0ffec6f.GetScheduleRequestBuilder) {
    return ia0a09be4f5c1f5b3b99aa3064830cdc6c730ceea9bdb4dbc6934a7d5f0ffec6f.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i847a76068ecb832435c51c746fba99e17f01da45f86bc0ea0f29b5a1d42ae5b3.MultiValueExtendedPropertiesRequestBuilder) {
    return i847a76068ecb832435c51c746fba99e17f01da45f86bc0ea0f29b5a1d42ae5b3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendarGroups.item.calendars.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*idefd7aa8b66a92e862195e44b4d12afbb8cb1f08136af3a806efad5c689470d7.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return idefd7aa8b66a92e862195e44b4d12afbb8cb1f08136af3a806efad5c689470d7.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendars in the calendar group. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *CalendarRequestBuilder) Patch(options *CalendarRequestBuilderPatchOptions)(error) {
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i0853bc5e78a4ac665d8aa887531e364debb1ba91c3a1bc736be9331cc8db0b52.SingleValueExtendedPropertiesRequestBuilder) {
    return i0853bc5e78a4ac665d8aa887531e364debb1ba91c3a1bc736be9331cc8db0b52.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendarGroups.item.calendars.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i857bdaf62ca4819c99b59159f0888909252fbffeba2fa7050dfc75ce2c6b598d.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i857bdaf62ca4819c99b59159f0888909252fbffeba2fa7050dfc75ce2c6b598d.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
