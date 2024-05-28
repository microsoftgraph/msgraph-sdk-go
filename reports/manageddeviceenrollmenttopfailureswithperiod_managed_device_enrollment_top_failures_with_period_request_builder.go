package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder provides operations to call the managedDeviceEnrollmentTopFailures method.
type ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal instantiates a new ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder and sets the default values.
func NewManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    m := &ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/managedDeviceEnrollmentTopFailures(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder instantiates a new ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder and sets the default values.
func NewManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function managedDeviceEnrollmentTopFailures
// returns a Reportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Reportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Reportable), nil
}
// ToGetRequestInformation invoke function managedDeviceEnrollmentTopFailures
// returns a *RequestInformation when successful
func (m *ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder when successful
func (m *ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) WithUrl(rawUrl string)(*ManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return NewManageddeviceenrollmenttopfailureswithperiodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
