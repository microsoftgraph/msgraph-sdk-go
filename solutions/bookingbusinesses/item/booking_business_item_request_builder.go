package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BookingBusinessItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BookingBusinessItemRequestBuilderGetQueryParameters get bookingBusinesses from solutions
type BookingBusinessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BookingBusinessItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookingBusinessItemRequestBuilderGetQueryParameters
}
// BookingBusinessItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Appointments the appointments property
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
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return i13fa9f102f5550ee0debc695a8553b7bb18a56cacbb2474b92004bac21331107.NewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
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
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return i666b4f4b64d4f34838144b722929f6fa3671d9928954f8b977c2e8b643847f95.NewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBookingBusinessItemRequestBuilderInternal instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessItemRequestBuilder) {
    m := &BookingBusinessItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBookingBusinessItemRequestBuilder instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *BookingBusinessItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformationWithRequestConfiguration get bookingBusinesses from solutions
func (m *BookingBusinessItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get bookingBusinesses from solutions
func (m *BookingBusinessItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *BookingBusinessItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property bookingBusinesses in solutions
func (m *BookingBusinessItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property bookingBusinesses in solutions
func (m *BookingBusinessItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, requestConfiguration *BookingBusinessItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Customers the customers property
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
        urlTplParams["bookingCustomerBase%2Did"] = id
    }
    return i0c8d27942293fe15088aee376560cc63f01e0de684e17140e97deb80c8d67040.NewBookingCustomerBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CustomQuestions the customQuestions property
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
        urlTplParams["bookingCustomQuestion%2Did"] = id
    }
    return i69ae791b5e1d3acb35e488b8f8f9fc402d31eddf07089c2cc865d430ceef10d9.NewBookingCustomQuestionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeleteWithResponseHandler delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *BookingBusinessItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *BookingBusinessItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler get bookingBusinesses from solutions
func (m *BookingBusinessItemRequestBuilder) GetWithResponseHandler(requestConfiguration *BookingBusinessItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler get bookingBusinesses from solutions
func (m *BookingBusinessItemRequestBuilder) GetWithResponseHandler(requestConfiguration *BookingBusinessItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingBusinessFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable), nil
}
// PatchWithResponseHandler update the navigation property bookingBusinesses in solutions
func (m *BookingBusinessItemRequestBuilder) PatchWithResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, requestConfiguration *BookingBusinessItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property bookingBusinesses in solutions
func (m *BookingBusinessItemRequestBuilder) PatchWithResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, requestConfiguration *BookingBusinessItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Publish the publish property
func (m *BookingBusinessItemRequestBuilder) Publish()(*ifaa3c3903221e665c7d867678d13a4a80d73f3c4698dc3b85c36793d3ae5e782.PublishRequestBuilder) {
    return ifaa3c3903221e665c7d867678d13a4a80d73f3c4698dc3b85c36793d3ae5e782.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Services the services property
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
        urlTplParams["bookingService%2Did"] = id
    }
    return i1dac550ea6326e46204247bf797ccc6ac538b2fbe14da8d2f54928298a9ad4d6.NewBookingServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// StaffMembers the staffMembers property
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
        urlTplParams["bookingStaffMemberBase%2Did"] = id
    }
    return i864a68457d2653b81efe9986ef9566d379d369000a1a4808dfb1ca609c7ea560.NewBookingStaffMemberBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *BookingBusinessItemRequestBuilder) Unpublish()(*i6bf58638d9684cbcac49ec92535cd01b9bdba8ece4bea2134c6d412b6ce1b72d.UnpublishRequestBuilder) {
    return i6bf58638d9684cbcac49ec92535cd01b9bdba8ece4bea2134c6d412b6ce1b72d.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
