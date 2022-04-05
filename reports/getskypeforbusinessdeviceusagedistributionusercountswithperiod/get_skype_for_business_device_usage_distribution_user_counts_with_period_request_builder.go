package getskypeforbusinessdeviceusagedistributionusercountswithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
type GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetOptions options for Get
type GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal instantiates a new GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    m := &GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getSkypeForBusinessDeviceUsageDistributionUserCounts(period='{period}')";
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
// NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder instantiates a new GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getSkypeForBusinessDeviceUsageDistributionUserCounts
func (m *GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getSkypeForBusinessDeviceUsageDistributionUserCounts
func (m *GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) Get(options *GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Reportable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateReportFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Reportable), nil
}
