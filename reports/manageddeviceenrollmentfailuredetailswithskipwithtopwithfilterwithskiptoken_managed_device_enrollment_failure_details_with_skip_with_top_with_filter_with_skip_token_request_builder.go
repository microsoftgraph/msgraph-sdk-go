package reports

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder provides operations to call the managedDeviceEnrollmentFailureDetails method.
type ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal instantiates a new ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder and sets the default values.
func NewManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, filter *string, skip *int32, skipToken *string, top *int32)(*ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    m := &ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/managedDeviceEnrollmentFailureDetails(skip={skip},top={top},filter='{filter}',skipToken='{skipToken}')", pathParameters),
    }
    if filter != nil {
        m.BaseRequestBuilder.PathParameters["filter"] = *filter
    }
    if skip != nil {
        m.BaseRequestBuilder.PathParameters["skip"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*skip), 10)
    }
    if skipToken != nil {
        m.BaseRequestBuilder.PathParameters["skipToken"] = *skipToken
    }
    if top != nil {
        m.BaseRequestBuilder.PathParameters["top"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*top), 10)
    }
    return m
}
// NewManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder instantiates a new ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder and sets the default values.
func NewManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil, nil)
}
// Get invoke function managedDeviceEnrollmentFailureDetails
// returns a Reportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Reportable, error) {
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
// ToGetRequestInformation invoke function managedDeviceEnrollmentFailureDetails
// returns a *RequestInformation when successful
func (m *ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder when successful
func (m *ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) WithUrl(rawUrl string)(*ManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewManageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
