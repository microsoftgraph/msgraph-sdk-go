package getteamsdeviceusageuserdetailwithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder provides operations to call the getTeamsDeviceUsageUserDetail method.
type GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderGetOptions options for Get
type GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal instantiates a new GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    m := &GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getTeamsDeviceUsageUserDetail(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder instantiates a new GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getTeamsDeviceUsageUserDetail
func (m *GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getTeamsDeviceUsageUserDetail
func (m *GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) Get(options *GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderGetOptions)(GetTeamsDeviceUsageUserDetailWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetTeamsDeviceUsageUserDetailWithPeriodResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetTeamsDeviceUsageUserDetailWithPeriodResponseable), nil
}
