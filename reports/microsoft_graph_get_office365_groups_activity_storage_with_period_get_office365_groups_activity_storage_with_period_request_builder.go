package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityStorage method.
type MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal instantiates a new GetOffice365GroupsActivityStorageWithPeriodRequestBuilder and sets the default values.
func NewMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    m := &MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getOffice365GroupsActivityStorage(period='{period}')";
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
// NewMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder instantiates a new GetOffice365GroupsActivityStorageWithPeriodRequestBuilder and sets the default values.
func NewMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365GroupsActivityStorage
func (m *MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getOffice365GroupsActivityStorage
func (m *MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
