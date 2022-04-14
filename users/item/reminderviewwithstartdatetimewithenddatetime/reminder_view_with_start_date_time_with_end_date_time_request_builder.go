package reminderviewwithstartdatetimewithenddatetime

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the reminderView method.
type ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetOptions options for Get
type ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *string, startDateTime *string)(*ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.reminderView(StartDateTime='{StartDateTime}',EndDateTime='{EndDateTime}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if endDateTime != nil {
        urlTplParams[""] = *endDateTime
    }
    if startDateTime != nil {
        urlTplParams[""] = *startDateTime
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function reminderView
func (m *ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) CreateGetRequestInformation(options *ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function reminderView
func (m *ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) Get(options *ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetOptions)(ReminderViewWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateReminderViewWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(ReminderViewWithStartDateTimeWithEndDateTimeResponseable), nil
}
