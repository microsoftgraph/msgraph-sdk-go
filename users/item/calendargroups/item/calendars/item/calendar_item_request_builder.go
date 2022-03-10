package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i0f5bcb7a24d4beb4ae67784520a3579932363bb2bfdb339ab0cf8221245ecc58 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/getschedule"
    i13913393b81b24c472b0854f70ed58ba370e3d688d3fc286f015e296e37bb3ae "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/allowedcalendarsharingroleswithuser"
    i2344dbbca462c86fbf066a3ffe12bd31610653945d64466880e923eb7f53aec6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/multivalueextendedproperties"
    i54e0033778f8c6c67b85a50409c9a29aebd1e7144d1e2e3c5bf118ec80a65ad7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/singlevalueextendedproperties"
    i75bc86155f6d06030ba24c9798e67ff834e381553541aacb8e0b1a22654f7c4d "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events"
    ic96778109cd4e1febe6211ff2686fa6e61d607c5070ff396a56168eac4d71eb0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarpermissions"
    ica635d049e7e80c35a6a124c9367cbf28debacf14e8a5b657ad008d352660570 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview"
    i82da6469abf82a1ac4ab553a277e83d9a202c2dc9468f3e513118568edcf7fb5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/singlevalueextendedproperties/item"
    i8aff0cbfd54a553c54c6aa7d6e012863c3635a1ec2ebb8862be1543ca198fac8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/multivalueextendedproperties/item"
    ia267b2781f42b15f11e0527b74cbe59b84de1ec5a54c86b5ae396a228cebf4ec "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item"
    ia388db07eba6a77790441359b2e24a78a8eee9665bc9b6af063c1c5883221794 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarpermissions/item"
    ief0d47706e51252f7f79b390fa08f36344fdcb0ee6a1e77eac4a8f315e66b285 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item"
)

// CalendarItemRequestBuilder provides operations to manage the calendars property of the microsoft.graph.calendarGroup entity.
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
// CalendarItemRequestBuilderGetQueryParameters the calendars in the calendar group. Navigation property. Read-only. Nullable.
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
func (m *CalendarItemRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i13913393b81b24c472b0854f70ed58ba370e3d688d3fc286f015e296e37bb3ae.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i13913393b81b24c472b0854f70ed58ba370e3d688d3fc286f015e296e37bb3ae.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarItemRequestBuilder) CalendarPermissions()(*ic96778109cd4e1febe6211ff2686fa6e61d607c5070ff396a56168eac4d71eb0.CalendarPermissionsRequestBuilder) {
    return ic96778109cd4e1febe6211ff2686fa6e61d607c5070ff396a56168eac4d71eb0.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarPermissions.item collection
func (m *CalendarItemRequestBuilder) CalendarPermissionsById(id string)(*ia388db07eba6a77790441359b2e24a78a8eee9665bc9b6af063c1c5883221794.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return ia388db07eba6a77790441359b2e24a78a8eee9665bc9b6af063c1c5883221794.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarItemRequestBuilder) CalendarView()(*ica635d049e7e80c35a6a124c9367cbf28debacf14e8a5b657ad008d352660570.CalendarViewRequestBuilder) {
    return ica635d049e7e80c35a6a124c9367cbf28debacf14e8a5b657ad008d352660570.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item collection
func (m *CalendarItemRequestBuilder) CalendarViewById(id string)(*ief0d47706e51252f7f79b390fa08f36344fdcb0ee6a1e77eac4a8f315e66b285.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return ief0d47706e51252f7f79b390fa08f36344fdcb0ee6a1e77eac4a8f315e66b285.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarItemRequestBuilderInternal instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarItemRequestBuilder) {
    m := &CalendarItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}{?select}";
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
// CreateDeleteRequestInformation delete navigation property calendars for users
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
// CreateGetRequestInformation the calendars in the calendar group. Navigation property. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property calendars in users
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
// Delete delete navigation property calendars for users
func (m *CalendarItemRequestBuilder) Delete(options *CalendarItemRequestBuilderDeleteOptions)(error) {
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
func (m *CalendarItemRequestBuilder) Events()(*i75bc86155f6d06030ba24c9798e67ff834e381553541aacb8e0b1a22654f7c4d.EventsRequestBuilder) {
    return i75bc86155f6d06030ba24c9798e67ff834e381553541aacb8e0b1a22654f7c4d.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item collection
func (m *CalendarItemRequestBuilder) EventsById(id string)(*ia267b2781f42b15f11e0527b74cbe59b84de1ec5a54c86b5ae396a228cebf4ec.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return ia267b2781f42b15f11e0527b74cbe59b84de1ec5a54c86b5ae396a228cebf4ec.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) Get(options *CalendarItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendarable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateCalendarFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendarable), nil
}
func (m *CalendarItemRequestBuilder) GetSchedule()(*i0f5bcb7a24d4beb4ae67784520a3579932363bb2bfdb339ab0cf8221245ecc58.GetScheduleRequestBuilder) {
    return i0f5bcb7a24d4beb4ae67784520a3579932363bb2bfdb339ab0cf8221245ecc58.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarItemRequestBuilder) MultiValueExtendedProperties()(*i2344dbbca462c86fbf066a3ffe12bd31610653945d64466880e923eb7f53aec6.MultiValueExtendedPropertiesRequestBuilder) {
    return i2344dbbca462c86fbf066a3ffe12bd31610653945d64466880e923eb7f53aec6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i8aff0cbfd54a553c54c6aa7d6e012863c3635a1ec2ebb8862be1543ca198fac8.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i8aff0cbfd54a553c54c6aa7d6e012863c3635a1ec2ebb8862be1543ca198fac8.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendars in users
func (m *CalendarItemRequestBuilder) Patch(options *CalendarItemRequestBuilderPatchOptions)(error) {
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
func (m *CalendarItemRequestBuilder) SingleValueExtendedProperties()(*i54e0033778f8c6c67b85a50409c9a29aebd1e7144d1e2e3c5bf118ec80a65ad7.SingleValueExtendedPropertiesRequestBuilder) {
    return i54e0033778f8c6c67b85a50409c9a29aebd1e7144d1e2e3c5bf118ec80a65ad7.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i82da6469abf82a1ac4ab553a277e83d9a202c2dc9468f3e513118568edcf7fb5.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i82da6469abf82a1ac4ab553a277e83d9a202c2dc9468f3e513118568edcf7fb5.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
