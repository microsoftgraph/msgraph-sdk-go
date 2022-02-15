package events

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ie15d874b5bcea33f7d9bcad18e2db2fe525af2c6393359a2b909ba2d2a99585a "github.com/microsoftgraph/msgraph-sdk-go/me/events/delta"
    ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/delta"
)

// EventsRequestBuilder builds and executes requests for operations under \me\events
type EventsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventsRequestBuilderGetOptions options for Get
type EventsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventsRequestBuilderGetQueryParameters the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
type EventsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// EventsRequestBuilderPostOptions options for Post
type EventsRequestBuilderPostOptions struct {
    // 
    Body *ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEventsRequestBuilderInternal instantiates a new EventsRequestBuilder and sets the default values.
func NewEventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventsRequestBuilder) {
    m := &EventsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events{?top,skip,filter,count,orderby,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventsRequestBuilder instantiates a new EventsRequestBuilder and sets the default values.
func NewEventsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventsRequestBuilder) CreateGetRequestInformation(options *EventsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventsRequestBuilder) CreatePostRequestInformation(options *EventsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Delta builds and executes requests for operations under \me\events\microsoft.graph.delta()
func (m *EventsRequestBuilder) Delta()(*ie15d874b5bcea33f7d9bcad18e2db2fe525af2c6393359a2b909ba2d2a99585a.DeltaRequestBuilder) {
    return ie15d874b5bcea33f7d9bcad18e2db2fe525af2c6393359a2b909ba2d2a99585a.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventsRequestBuilder) Get(options *EventsRequestBuilderGetOptions)(*EventsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEventsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*EventsResponse), nil
}
// Post the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventsRequestBuilder) Post(options *EventsRequestBuilderPostOptions)(*ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.NewEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event), nil
}
