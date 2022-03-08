package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i027a46ae0e95e4e405d21a59eb6505ee2295984a990bf831237140753c099f9f "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/staffmembers"
    i4466de1bcd62b1c15cfe7be86a5394e5ed5fecb749ab217db3d3ab04a429ade6 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/services"
    i5e926d6326c2dc0b5bbb51ebc6bd50b43309aebd5643876218683cae253f9690 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/calendarview"
    i6bf58638d9684cbcac49ec92535cd01b9bdba8ece4bea2134c6d412b6ce1b72d "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/unpublish"
    i8a847b504989836f4aab54084412c3710787a77712f49c9d87a30b7f9436f181 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/customers"
    i9e67c37e4d7e393bcaa474ec71c4d548047669ab57e739c8fee48bfec1f9a6ad "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/customquestions"
    ic2d5072df0f8d04e99de862f3a1d2feb5ded47db9ddb8fabf60b325a3045df1c "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/appointments"
    ifaa3c3903221e665c7d867678d13a4a80d73f3c4698dc3b85c36793d3ae5e782 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/publish"
    i0c8d27942293fe15088aee376560cc63f01e0de684e17140e97deb80c8d67040 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/customers/item"
    i13fa9f102f5550ee0debc695a8553b7bb18a56cacbb2474b92004bac21331107 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/appointments/item"
    i1dac550ea6326e46204247bf797ccc6ac538b2fbe14da8d2f54928298a9ad4d6 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/services/item"
    i666b4f4b64d4f34838144b722929f6fa3671d9928954f8b977c2e8b643847f95 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/calendarview/item"
    i69ae791b5e1d3acb35e488b8f8f9fc402d31eddf07089c2cc865d430ceef10d9 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/customquestions/item"
    i864a68457d2653b81efe9986ef9566d379d369000a1a4808dfb1ca609c7ea560 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item/staffmembers/item"
)

// BookingBusinessItemRequestBuilder provides operations to manage the bookingBusinesses property of the microsoft.graph.solutionsRoot entity.
type BookingBusinessItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BookingBusinessItemRequestBuilderDeleteOptions options for Delete
type BookingBusinessItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BookingBusinessItemRequestBuilderGetOptions options for Get
type BookingBusinessItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BookingBusinessItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BookingBusinessItemRequestBuilderGetQueryParameters get bookingBusinesses from solutions
type BookingBusinessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BookingBusinessItemRequestBuilderPatchOptions options for Patch
type BookingBusinessItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.BookingBusinessable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *BookingBusinessItemRequestBuilder) Appointments()(*ic2d5072df0f8d04e99de862f3a1d2feb5ded47db9ddb8fabf60b325a3045df1c.AppointmentsRequestBuilder) {
    return ic2d5072df0f8d04e99de862f3a1d2feb5ded47db9ddb8fabf60b325a3045df1c.NewAppointmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppointmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.appointments.item collection
func (m *BookingBusinessItemRequestBuilder) AppointmentsById(id string)(*i13fa9f102f5550ee0debc695a8553b7bb18a56cacbb2474b92004bac21331107.BookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment_id"] = id
    }
    return i13fa9f102f5550ee0debc695a8553b7bb18a56cacbb2474b92004bac21331107.NewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessItemRequestBuilder) CalendarView()(*i5e926d6326c2dc0b5bbb51ebc6bd50b43309aebd5643876218683cae253f9690.CalendarViewRequestBuilder) {
    return i5e926d6326c2dc0b5bbb51ebc6bd50b43309aebd5643876218683cae253f9690.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.calendarView.item collection
func (m *BookingBusinessItemRequestBuilder) CalendarViewById(id string)(*i666b4f4b64d4f34838144b722929f6fa3671d9928954f8b977c2e8b643847f95.BookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment_id"] = id
    }
    return i666b4f4b64d4f34838144b722929f6fa3671d9928954f8b977c2e8b643847f95.NewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBookingBusinessItemRequestBuilderInternal instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookingBusinessItemRequestBuilder) {
    m := &BookingBusinessItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBookingBusinessItemRequestBuilder instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookingBusinessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessItemRequestBuilder) CreateDeleteRequestInformation(options *BookingBusinessItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get bookingBusinesses from solutions
func (m *BookingBusinessItemRequestBuilder) CreateGetRequestInformation(options *BookingBusinessItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property bookingBusinesses in solutions
func (m *BookingBusinessItemRequestBuilder) CreatePatchRequestInformation(options *BookingBusinessItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *BookingBusinessItemRequestBuilder) Customers()(*i8a847b504989836f4aab54084412c3710787a77712f49c9d87a30b7f9436f181.CustomersRequestBuilder) {
    return i8a847b504989836f4aab54084412c3710787a77712f49c9d87a30b7f9436f181.NewCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.customers.item collection
func (m *BookingBusinessItemRequestBuilder) CustomersById(id string)(*i0c8d27942293fe15088aee376560cc63f01e0de684e17140e97deb80c8d67040.BookingCustomerBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomerBase_id"] = id
    }
    return i0c8d27942293fe15088aee376560cc63f01e0de684e17140e97deb80c8d67040.NewBookingCustomerBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessItemRequestBuilder) CustomQuestions()(*i9e67c37e4d7e393bcaa474ec71c4d548047669ab57e739c8fee48bfec1f9a6ad.CustomQuestionsRequestBuilder) {
    return i9e67c37e4d7e393bcaa474ec71c4d548047669ab57e739c8fee48bfec1f9a6ad.NewCustomQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomQuestionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.customQuestions.item collection
func (m *BookingBusinessItemRequestBuilder) CustomQuestionsById(id string)(*i69ae791b5e1d3acb35e488b8f8f9fc402d31eddf07089c2cc865d430ceef10d9.BookingCustomQuestionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomQuestion_id"] = id
    }
    return i69ae791b5e1d3acb35e488b8f8f9fc402d31eddf07089c2cc865d430ceef10d9.NewBookingCustomQuestionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessItemRequestBuilder) Delete(options *BookingBusinessItemRequestBuilderDeleteOptions)(error) {
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
// Get get bookingBusinesses from solutions
func (m *BookingBusinessItemRequestBuilder) Get(options *BookingBusinessItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.BookingBusinessable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateBookingBusinessFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.BookingBusinessable), nil
}
// Patch update the navigation property bookingBusinesses in solutions
func (m *BookingBusinessItemRequestBuilder) Patch(options *BookingBusinessItemRequestBuilderPatchOptions)(error) {
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
func (m *BookingBusinessItemRequestBuilder) Publish()(*ifaa3c3903221e665c7d867678d13a4a80d73f3c4698dc3b85c36793d3ae5e782.PublishRequestBuilder) {
    return ifaa3c3903221e665c7d867678d13a4a80d73f3c4698dc3b85c36793d3ae5e782.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BookingBusinessItemRequestBuilder) Services()(*i4466de1bcd62b1c15cfe7be86a5394e5ed5fecb749ab217db3d3ab04a429ade6.ServicesRequestBuilder) {
    return i4466de1bcd62b1c15cfe7be86a5394e5ed5fecb749ab217db3d3ab04a429ade6.NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.services.item collection
func (m *BookingBusinessItemRequestBuilder) ServicesById(id string)(*i1dac550ea6326e46204247bf797ccc6ac538b2fbe14da8d2f54928298a9ad4d6.BookingServiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingService_id"] = id
    }
    return i1dac550ea6326e46204247bf797ccc6ac538b2fbe14da8d2f54928298a9ad4d6.NewBookingServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessItemRequestBuilder) StaffMembers()(*i027a46ae0e95e4e405d21a59eb6505ee2295984a990bf831237140753c099f9f.StaffMembersRequestBuilder) {
    return i027a46ae0e95e4e405d21a59eb6505ee2295984a990bf831237140753c099f9f.NewStaffMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StaffMembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.staffMembers.item collection
func (m *BookingBusinessItemRequestBuilder) StaffMembersById(id string)(*i864a68457d2653b81efe9986ef9566d379d369000a1a4808dfb1ca609c7ea560.BookingStaffMemberBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingStaffMemberBase_id"] = id
    }
    return i864a68457d2653b81efe9986ef9566d379d369000a1a4808dfb1ca609c7ea560.NewBookingStaffMemberBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessItemRequestBuilder) Unpublish()(*i6bf58638d9684cbcac49ec92535cd01b9bdba8ece4bea2134c6d412b6ce1b72d.UnpublishRequestBuilder) {
    return i6bf58638d9684cbcac49ec92535cd01b9bdba8ece4bea2134c6d412b6ce1b72d.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
