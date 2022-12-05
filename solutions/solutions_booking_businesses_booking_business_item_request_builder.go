package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SolutionsBookingBusinessesBookingBusinessItemRequestBuilder provides operations to manage the bookingBusinesses property of the microsoft.graph.solutionsRoot entity.
type SolutionsBookingBusinessesBookingBusinessItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SolutionsBookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsBookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SolutionsBookingBusinessesBookingBusinessItemRequestBuilderGetQueryParameters get bookingBusinesses from solutions
type SolutionsBookingBusinessesBookingBusinessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SolutionsBookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsBookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SolutionsBookingBusinessesBookingBusinessItemRequestBuilderGetQueryParameters
}
// SolutionsBookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsBookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Appointments provides operations to manage the appointments property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) Appointments()(*SolutionsBookingBusinessesItemAppointmentsRequestBuilder) {
    return NewSolutionsBookingBusinessesItemAppointmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppointmentsById provides operations to manage the appointments property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) AppointmentsById(id string)(*SolutionsBookingBusinessesItemAppointmentsBookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return NewSolutionsBookingBusinessesItemAppointmentsBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) CalendarView()(*SolutionsBookingBusinessesItemCalendarViewRequestBuilder) {
    return NewSolutionsBookingBusinessesItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) CalendarViewById(id string)(*SolutionsBookingBusinessesItemCalendarViewBookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return NewSolutionsBookingBusinessesItemCalendarViewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSolutionsBookingBusinessesBookingBusinessItemRequestBuilderInternal instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewSolutionsBookingBusinessesBookingBusinessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) {
    m := &SolutionsBookingBusinessesBookingBusinessItemRequestBuilder{
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
// NewSolutionsBookingBusinessesBookingBusinessItemRequestBuilder instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewSolutionsBookingBusinessesBookingBusinessItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSolutionsBookingBusinessesBookingBusinessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property bookingBusinesses for solutions
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SolutionsBookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get bookingBusinesses from solutions
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SolutionsBookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property bookingBusinesses in solutions
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, requestConfiguration *SolutionsBookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Customers provides operations to manage the customers property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) Customers()(*SolutionsBookingBusinessesItemCustomersRequestBuilder) {
    return NewSolutionsBookingBusinessesItemCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomersById provides operations to manage the customers property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) CustomersById(id string)(*SolutionsBookingBusinessesItemCustomersBookingCustomerBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomerBase%2Did"] = id
    }
    return NewSolutionsBookingBusinessesItemCustomersBookingCustomerBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CustomQuestions provides operations to manage the customQuestions property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) CustomQuestions()(*SolutionsBookingBusinessesItemCustomQuestionsRequestBuilder) {
    return NewSolutionsBookingBusinessesItemCustomQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomQuestionsById provides operations to manage the customQuestions property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) CustomQuestionsById(id string)(*SolutionsBookingBusinessesItemCustomQuestionsBookingCustomQuestionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomQuestion%2Did"] = id
    }
    return NewSolutionsBookingBusinessesItemCustomQuestionsBookingCustomQuestionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property bookingBusinesses for solutions
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SolutionsBookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get bookingBusinesses from solutions
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SolutionsBookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingBusinessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable), nil
}
// GetStaffAvailability provides operations to call the getStaffAvailability method.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) GetStaffAvailability()(*SolutionsBookingBusinessesItemGetStaffAvailabilityRequestBuilder) {
    return NewSolutionsBookingBusinessesItemGetStaffAvailabilityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bookingBusinesses in solutions
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, requestConfiguration *SolutionsBookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingBusinessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable), nil
}
// Publish provides operations to call the publish method.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) Publish()(*SolutionsBookingBusinessesItemPublishRequestBuilder) {
    return NewSolutionsBookingBusinessesItemPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Services provides operations to manage the services property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) Services()(*SolutionsBookingBusinessesItemServicesRequestBuilder) {
    return NewSolutionsBookingBusinessesItemServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicesById provides operations to manage the services property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) ServicesById(id string)(*SolutionsBookingBusinessesItemServicesBookingServiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingService%2Did"] = id
    }
    return NewSolutionsBookingBusinessesItemServicesBookingServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// StaffMembers provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) StaffMembers()(*SolutionsBookingBusinessesItemStaffMembersRequestBuilder) {
    return NewSolutionsBookingBusinessesItemStaffMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StaffMembersById provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) StaffMembersById(id string)(*SolutionsBookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingStaffMemberBase%2Did"] = id
    }
    return NewSolutionsBookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unpublish provides operations to call the unpublish method.
func (m *SolutionsBookingBusinessesBookingBusinessItemRequestBuilder) Unpublish()(*SolutionsBookingBusinessesItemUnpublishRequestBuilder) {
    return NewSolutionsBookingBusinessesItemUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
