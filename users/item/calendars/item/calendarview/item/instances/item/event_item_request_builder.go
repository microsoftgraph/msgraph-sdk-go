package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i272324389d6bf1f60ccbfa77144be57c1d70c0718d79da10d8dee55e5f62919e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/cancel"
    i2ffebe5209de4076e8a82adab1a4bfca6c05cdf6f208c455125b48c57db65904 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/multivalueextendedproperties"
    i44867788db34306f0996cf494a44a02c2fdfeba2f1d248dfbf3719ef279f5d35 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/decline"
    i48b536acd4f01ac6a9b0c428b9de43a1cbde1b708b519c38213a57d87b96504f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties"
    i7cfc21431c7f3715849334d5beda9eb12748eea231818a18c33f9cd7311e2ea7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/calendar"
    i7d3117e3afb0fbd44802c1ad6d61175bbd7b795db5acdd75c10d1c1da8aed689 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/extensions"
    i81701a116c93dee173023918621f350692ddfe7873f6de037bccfa9bc683dc45 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/accept"
    i83473d2a0bc99449ff6981736c3c17fbee0979720e884cddb635db5c58ed629c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/forward"
    i9bf36797cd9f9784c47c527ad0e9d3362b3dc78d433e24e67f96bcd93f8c8dd2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/dismissreminder"
    iba1c2f9e853ce3684320c2cf66b1bf622a32492dfdb465c4f5900f2609cac816 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/attachments"
    ic2b75b590eb1afa68b5b139fda2fa1f65d5573ce52b3716a67f1e95fa70569e8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/snoozereminder"
    idcd892355bb9b0a3694ef802b374e6be28dd95de1898988e0f5500a51feedf6f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/tentativelyaccept"
    i0a8a1ba6dbe110e4b0d2005064e6c047ae4063f6e259f00465d057bc61bf5204 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    i7d0e215fee21f7bd421540d9e6919033aaaf98c664446238d006381d73e70691 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/attachments/item"
    icc005ad8da2b3b97a3de6350a707a53f032a558960b40cd2afcc1c28e246ada6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/extensions/item"
    ifcf0b211e141030ced11d35ff3b92d777831766af0ce38e3192dd97c536b3221 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i81701a116c93dee173023918621f350692ddfe7873f6de037bccfa9bc683dc45.AcceptRequestBuilder) {
    return i81701a116c93dee173023918621f350692ddfe7873f6de037bccfa9bc683dc45.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*iba1c2f9e853ce3684320c2cf66b1bf622a32492dfdb465c4f5900f2609cac816.AttachmentsRequestBuilder) {
    return iba1c2f9e853ce3684320c2cf66b1bf622a32492dfdb465c4f5900f2609cac816.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i7d0e215fee21f7bd421540d9e6919033aaaf98c664446238d006381d73e70691.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i7d0e215fee21f7bd421540d9e6919033aaaf98c664446238d006381d73e70691.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i7cfc21431c7f3715849334d5beda9eb12748eea231818a18c33f9cd7311e2ea7.CalendarRequestBuilder) {
    return i7cfc21431c7f3715849334d5beda9eb12748eea231818a18c33f9cd7311e2ea7.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i272324389d6bf1f60ccbfa77144be57c1d70c0718d79da10d8dee55e5f62919e.CancelRequestBuilder) {
    return i272324389d6bf1f60ccbfa77144be57c1d70c0718d79da10d8dee55e5f62919e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}/calendarView/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i44867788db34306f0996cf494a44a02c2fdfeba2f1d248dfbf3719ef279f5d35.DeclineRequestBuilder) {
    return i44867788db34306f0996cf494a44a02c2fdfeba2f1d248dfbf3719ef279f5d35.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i9bf36797cd9f9784c47c527ad0e9d3362b3dc78d433e24e67f96bcd93f8c8dd2.DismissReminderRequestBuilder) {
    return i9bf36797cd9f9784c47c527ad0e9d3362b3dc78d433e24e67f96bcd93f8c8dd2.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i7d3117e3afb0fbd44802c1ad6d61175bbd7b795db5acdd75c10d1c1da8aed689.ExtensionsRequestBuilder) {
    return i7d3117e3afb0fbd44802c1ad6d61175bbd7b795db5acdd75c10d1c1da8aed689.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*icc005ad8da2b3b97a3de6350a707a53f032a558960b40cd2afcc1c28e246ada6.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return icc005ad8da2b3b97a3de6350a707a53f032a558960b40cd2afcc1c28e246ada6.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i83473d2a0bc99449ff6981736c3c17fbee0979720e884cddb635db5c58ed629c.ForwardRequestBuilder) {
    return i83473d2a0bc99449ff6981736c3c17fbee0979720e884cddb635db5c58ed629c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i2ffebe5209de4076e8a82adab1a4bfca6c05cdf6f208c455125b48c57db65904.MultiValueExtendedPropertiesRequestBuilder) {
    return i2ffebe5209de4076e8a82adab1a4bfca6c05cdf6f208c455125b48c57db65904.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ifcf0b211e141030ced11d35ff3b92d777831766af0ce38e3192dd97c536b3221.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ifcf0b211e141030ced11d35ff3b92d777831766af0ce38e3192dd97c536b3221.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i48b536acd4f01ac6a9b0c428b9de43a1cbde1b708b519c38213a57d87b96504f.SingleValueExtendedPropertiesRequestBuilder) {
    return i48b536acd4f01ac6a9b0c428b9de43a1cbde1b708b519c38213a57d87b96504f.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0a8a1ba6dbe110e4b0d2005064e6c047ae4063f6e259f00465d057bc61bf5204.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i0a8a1ba6dbe110e4b0d2005064e6c047ae4063f6e259f00465d057bc61bf5204.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ic2b75b590eb1afa68b5b139fda2fa1f65d5573ce52b3716a67f1e95fa70569e8.SnoozeReminderRequestBuilder) {
    return ic2b75b590eb1afa68b5b139fda2fa1f65d5573ce52b3716a67f1e95fa70569e8.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*idcd892355bb9b0a3694ef802b374e6be28dd95de1898988e0f5500a51feedf6f.TentativelyAcceptRequestBuilder) {
    return idcd892355bb9b0a3694ef802b374e6be28dd95de1898988e0f5500a51feedf6f.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
