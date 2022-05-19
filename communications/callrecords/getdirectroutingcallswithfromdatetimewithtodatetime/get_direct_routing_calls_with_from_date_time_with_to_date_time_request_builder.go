package getdirectroutingcallswithfromdatetimewithtodatetime

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder provides operations to call the getDirectRoutingCalls method.
type GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal instantiates a new GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    m := &GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/callRecords/microsoft.graph.callRecords.getDirectRoutingCalls(fromDateTime='{fromDateTime}',toDateTime='{toDateTime}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if fromDateTime != nil {
        urlTplParams[""] = (*fromDateTime).String()
    }
    if toDateTime != nil {
        urlTplParams[""] = (*toDateTime).String()
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder instantiates a new GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function getDirectRoutingCalls
func (m *GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getDirectRoutingCalls
func (m *GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getDirectRoutingCalls
func (m *GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) Get()(GetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getDirectRoutingCalls
func (m *GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseable), nil
}
