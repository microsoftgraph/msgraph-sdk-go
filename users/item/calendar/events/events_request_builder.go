package events

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i3f4801a2f466b489931174c8be1727272b0ee8dd269bfb2faf583c964fba15e4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/delta"
    iacce083202b0934f81b6fb2b51edab2a4bc17d513ab0f9f29de48fc16aa138e0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/count"
)

// EventsRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
type EventsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventsRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type EventsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EventsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventsRequestBuilderGetQueryParameters
}
// EventsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEventsRequestBuilderInternal instantiates a new EventsRequestBuilder and sets the default values.
func NewEventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventsRequestBuilder) {
    m := &EventsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/events{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventsRequestBuilder instantiates a new EventsRequestBuilder and sets the default values.
func NewEventsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *EventsRequestBuilder) Count()(*iacce083202b0934f81b6fb2b51edab2a4bc17d513ab0f9f29de48fc16aa138e0.CountRequestBuilder) {
    return iacce083202b0934f81b6fb2b51edab2a4bc17d513ab0f9f29de48fc16aa138e0.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *EventsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the events in the calendar. Navigation property. Read-only.
func (m *EventsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EventsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to events for users
func (m *EventsRequestBuilder) CreatePostRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to events for users
func (m *EventsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, requestConfiguration *EventsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *EventsRequestBuilder) Delta()(*i3f4801a2f466b489931174c8be1727272b0ee8dd269bfb2faf583c964fba15e4.DeltaRequestBuilder) {
    return i3f4801a2f466b489931174c8be1727272b0ee8dd269bfb2faf583c964fba15e4.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *EventsRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the events in the calendar. Navigation property. Read-only.
func (m *EventsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *EventsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable), nil
}
// Post create new navigation property to events for users
func (m *EventsRequestBuilder) Post(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to events for users
func (m *EventsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, requestConfiguration *EventsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable), nil
}
