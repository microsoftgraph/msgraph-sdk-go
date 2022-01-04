package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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

// BookingBusinessRequestBuilder builds and executes requests for operations under \solutions\bookingBusinesses\{bookingBusiness-id}
type BookingBusinessRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BookingBusinessRequestBuilderDeleteOptions options for Delete
type BookingBusinessRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BookingBusinessRequestBuilderGetOptions options for Get
type BookingBusinessRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BookingBusinessRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BookingBusinessRequestBuilderGetQueryParameters get bookingBusinesses from solutions
type BookingBusinessRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BookingBusinessRequestBuilderPatchOptions options for Patch
type BookingBusinessRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.BookingBusiness;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *BookingBusinessRequestBuilder) Appointments()(*ic2d5072df0f8d04e99de862f3a1d2feb5ded47db9ddb8fabf60b325a3045df1c.AppointmentsRequestBuilder) {
    return ic2d5072df0f8d04e99de862f3a1d2feb5ded47db9ddb8fabf60b325a3045df1c.NewAppointmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppointmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.appointments.item collection
func (m *BookingBusinessRequestBuilder) AppointmentsById(id string)(*i13fa9f102f5550ee0debc695a8553b7bb18a56cacbb2474b92004bac21331107.BookingAppointmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment_id"] = id
    }
    return i13fa9f102f5550ee0debc695a8553b7bb18a56cacbb2474b92004bac21331107.NewBookingAppointmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) CalendarView()(*i5e926d6326c2dc0b5bbb51ebc6bd50b43309aebd5643876218683cae253f9690.CalendarViewRequestBuilder) {
    return i5e926d6326c2dc0b5bbb51ebc6bd50b43309aebd5643876218683cae253f9690.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.calendarView.item collection
func (m *BookingBusinessRequestBuilder) CalendarViewById(id string)(*i666b4f4b64d4f34838144b722929f6fa3671d9928954f8b977c2e8b643847f95.BookingAppointmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment_id"] = id
    }
    return i666b4f4b64d4f34838144b722929f6fa3671d9928954f8b977c2e8b643847f95.NewBookingAppointmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBookingBusinessRequestBuilderInternal instantiates a new BookingBusinessRequestBuilder and sets the default values.
func NewBookingBusinessRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookingBusinessRequestBuilder) {
    m := &BookingBusinessRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBookingBusinessRequestBuilder instantiates a new BookingBusinessRequestBuilder and sets the default values.
func NewBookingBusinessRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookingBusinessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessRequestBuilder) CreateDeleteRequestInformation(options *BookingBusinessRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *BookingBusinessRequestBuilder) CreateGetRequestInformation(options *BookingBusinessRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *BookingBusinessRequestBuilder) CreatePatchRequestInformation(options *BookingBusinessRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *BookingBusinessRequestBuilder) Customers()(*i8a847b504989836f4aab54084412c3710787a77712f49c9d87a30b7f9436f181.CustomersRequestBuilder) {
    return i8a847b504989836f4aab54084412c3710787a77712f49c9d87a30b7f9436f181.NewCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.customers.item collection
func (m *BookingBusinessRequestBuilder) CustomersById(id string)(*i0c8d27942293fe15088aee376560cc63f01e0de684e17140e97deb80c8d67040.BookingCustomerBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomerBase_id"] = id
    }
    return i0c8d27942293fe15088aee376560cc63f01e0de684e17140e97deb80c8d67040.NewBookingCustomerBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) CustomQuestions()(*i9e67c37e4d7e393bcaa474ec71c4d548047669ab57e739c8fee48bfec1f9a6ad.CustomQuestionsRequestBuilder) {
    return i9e67c37e4d7e393bcaa474ec71c4d548047669ab57e739c8fee48bfec1f9a6ad.NewCustomQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomQuestionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.customQuestions.item collection
func (m *BookingBusinessRequestBuilder) CustomQuestionsById(id string)(*i69ae791b5e1d3acb35e488b8f8f9fc402d31eddf07089c2cc865d430ceef10d9.BookingCustomQuestionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomQuestion_id"] = id
    }
    return i69ae791b5e1d3acb35e488b8f8f9fc402d31eddf07089c2cc865d430ceef10d9.NewBookingCustomQuestionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessRequestBuilder) Delete(options *BookingBusinessRequestBuilderDeleteOptions)(error) {
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
// Get get bookingBusinesses from solutions
func (m *BookingBusinessRequestBuilder) Get(options *BookingBusinessRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.BookingBusiness, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewBookingBusiness() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.BookingBusiness), nil
}
// Patch update the navigation property bookingBusinesses in solutions
func (m *BookingBusinessRequestBuilder) Patch(options *BookingBusinessRequestBuilderPatchOptions)(error) {
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
func (m *BookingBusinessRequestBuilder) Publish()(*ifaa3c3903221e665c7d867678d13a4a80d73f3c4698dc3b85c36793d3ae5e782.PublishRequestBuilder) {
    return ifaa3c3903221e665c7d867678d13a4a80d73f3c4698dc3b85c36793d3ae5e782.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) Services()(*i4466de1bcd62b1c15cfe7be86a5394e5ed5fecb749ab217db3d3ab04a429ade6.ServicesRequestBuilder) {
    return i4466de1bcd62b1c15cfe7be86a5394e5ed5fecb749ab217db3d3ab04a429ade6.NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.services.item collection
func (m *BookingBusinessRequestBuilder) ServicesById(id string)(*i1dac550ea6326e46204247bf797ccc6ac538b2fbe14da8d2f54928298a9ad4d6.BookingServiceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingService_id"] = id
    }
    return i1dac550ea6326e46204247bf797ccc6ac538b2fbe14da8d2f54928298a9ad4d6.NewBookingServiceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) StaffMembers()(*i027a46ae0e95e4e405d21a59eb6505ee2295984a990bf831237140753c099f9f.StaffMembersRequestBuilder) {
    return i027a46ae0e95e4e405d21a59eb6505ee2295984a990bf831237140753c099f9f.NewStaffMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StaffMembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item.staffMembers.item collection
func (m *BookingBusinessRequestBuilder) StaffMembersById(id string)(*i864a68457d2653b81efe9986ef9566d379d369000a1a4808dfb1ca609c7ea560.BookingStaffMemberBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingStaffMemberBase_id"] = id
    }
    return i864a68457d2653b81efe9986ef9566d379d369000a1a4808dfb1ca609c7ea560.NewBookingStaffMemberBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) Unpublish()(*i6bf58638d9684cbcac49ec92535cd01b9bdba8ece4bea2134c6d412b6ce1b72d.UnpublishRequestBuilder) {
    return i6bf58638d9684cbcac49ec92535cd01b9bdba8ece4bea2134c6d412b6ce1b72d.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
