package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0608c7e63daed1d8c25dae69476d30ed0865692d2083143f0218ba545bd1e6fe "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/forward"
    i07d0e4dc59220db4d7a12d39275288711f40897a0177e9303992283ab77defff "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/singlevalueextendedproperties"
    i600ccf8ecb63c2306bea2313987324f661d6a99b86ce8c4d346b598ccb947723 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/accept"
    i6ea1b94176d49fe77dc1401acdc47611ad8a52c9c6679fb9a0f4042e4a8b1e2b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/attachments"
    i6f6a379fccd6d7ecab5d208cda6abdeea04b51f72052df61fbd909540c1d7785 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/calendar"
    i80f3a65475ab0b52025d7da02f157b571a71ec1653a35767a6445d3396413731 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/cancel"
    i88a6407fe07fdde4b4f64edd95fdd05984b8f248ff13190c22eecee6198b3935 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/extensions"
    i90956c2f657489b41c77cca666d102dea0bdf9d73b17f4a49c73f6e2ed3271e9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/decline"
    i98290a55afd1d1f0fd9f466411a505e295f1873fb97266f9cc506b0d4d6917ac "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/multivalueextendedproperties"
    ia39929ce6cbf360020596a4720c2bdc855842e132ab54471ed8c11ddcf661102 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/dismissreminder"
    iac5a3fc9fdba162b0d168cefc298dc8c085d69c261727fb186025595df5972f8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/tentativelyaccept"
    if29ba39feb86bf644a6ce867550b8dc648df46f2908d6b2ac6dcd043bc00fb6e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/snoozereminder"
    i0ac354fe96550f13333780b4988ce2c8fbb97ccb242f8a0b579a9d16173fcce3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/extensions/item"
    i2ca41bd2f6db760c9dda4d70fef3da36f4fff8e684fcc1a720ef69b56eb396cf "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    i6aa496b042042eb3121230fe915280bf61f9463e3012cdaddc86bdfc32da2186 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/multivalueextendedproperties/item"
    ib7e466ee049923144d5c5bd395e80154c60e05395739ccdb4e370be08c2b4127 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/instances/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i600ccf8ecb63c2306bea2313987324f661d6a99b86ce8c4d346b598ccb947723.AcceptRequestBuilder) {
    return i600ccf8ecb63c2306bea2313987324f661d6a99b86ce8c4d346b598ccb947723.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i6ea1b94176d49fe77dc1401acdc47611ad8a52c9c6679fb9a0f4042e4a8b1e2b.AttachmentsRequestBuilder) {
    return i6ea1b94176d49fe77dc1401acdc47611ad8a52c9c6679fb9a0f4042e4a8b1e2b.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib7e466ee049923144d5c5bd395e80154c60e05395739ccdb4e370be08c2b4127.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ib7e466ee049923144d5c5bd395e80154c60e05395739ccdb4e370be08c2b4127.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i6f6a379fccd6d7ecab5d208cda6abdeea04b51f72052df61fbd909540c1d7785.CalendarRequestBuilder) {
    return i6f6a379fccd6d7ecab5d208cda6abdeea04b51f72052df61fbd909540c1d7785.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i80f3a65475ab0b52025d7da02f157b571a71ec1653a35767a6445d3396413731.CancelRequestBuilder) {
    return i80f3a65475ab0b52025d7da02f157b571a71ec1653a35767a6445d3396413731.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarView/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i90956c2f657489b41c77cca666d102dea0bdf9d73b17f4a49c73f6e2ed3271e9.DeclineRequestBuilder) {
    return i90956c2f657489b41c77cca666d102dea0bdf9d73b17f4a49c73f6e2ed3271e9.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*ia39929ce6cbf360020596a4720c2bdc855842e132ab54471ed8c11ddcf661102.DismissReminderRequestBuilder) {
    return ia39929ce6cbf360020596a4720c2bdc855842e132ab54471ed8c11ddcf661102.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i88a6407fe07fdde4b4f64edd95fdd05984b8f248ff13190c22eecee6198b3935.ExtensionsRequestBuilder) {
    return i88a6407fe07fdde4b4f64edd95fdd05984b8f248ff13190c22eecee6198b3935.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i0ac354fe96550f13333780b4988ce2c8fbb97ccb242f8a0b579a9d16173fcce3.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i0ac354fe96550f13333780b4988ce2c8fbb97ccb242f8a0b579a9d16173fcce3.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i0608c7e63daed1d8c25dae69476d30ed0865692d2083143f0218ba545bd1e6fe.ForwardRequestBuilder) {
    return i0608c7e63daed1d8c25dae69476d30ed0865692d2083143f0218ba545bd1e6fe.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i98290a55afd1d1f0fd9f466411a505e295f1873fb97266f9cc506b0d4d6917ac.MultiValueExtendedPropertiesRequestBuilder) {
    return i98290a55afd1d1f0fd9f466411a505e295f1873fb97266f9cc506b0d4d6917ac.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i6aa496b042042eb3121230fe915280bf61f9463e3012cdaddc86bdfc32da2186.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i6aa496b042042eb3121230fe915280bf61f9463e3012cdaddc86bdfc32da2186.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i07d0e4dc59220db4d7a12d39275288711f40897a0177e9303992283ab77defff.SingleValueExtendedPropertiesRequestBuilder) {
    return i07d0e4dc59220db4d7a12d39275288711f40897a0177e9303992283ab77defff.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i2ca41bd2f6db760c9dda4d70fef3da36f4fff8e684fcc1a720ef69b56eb396cf.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i2ca41bd2f6db760c9dda4d70fef3da36f4fff8e684fcc1a720ef69b56eb396cf.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*if29ba39feb86bf644a6ce867550b8dc648df46f2908d6b2ac6dcd043bc00fb6e.SnoozeReminderRequestBuilder) {
    return if29ba39feb86bf644a6ce867550b8dc648df46f2908d6b2ac6dcd043bc00fb6e.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*iac5a3fc9fdba162b0d168cefc298dc8c085d69c261727fb186025595df5972f8.TentativelyAcceptRequestBuilder) {
    return iac5a3fc9fdba162b0d168cefc298dc8c085d69c261727fb186025595df5972f8.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
