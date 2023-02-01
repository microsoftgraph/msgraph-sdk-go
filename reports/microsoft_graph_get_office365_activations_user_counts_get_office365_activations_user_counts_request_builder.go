package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder provides operations to call the getOffice365ActivationsUserCounts method.
type MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilderInternal instantiates a new GetOffice365ActivationsUserCountsRequestBuilder and sets the default values.
func NewMicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder) {
    m := &MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getOffice365ActivationsUserCounts()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder instantiates a new GetOffice365ActivationsUserCountsRequestBuilder and sets the default values.
func NewMicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getOffice365ActivationsUserCounts
func (m *MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation invoke function getOffice365ActivationsUserCounts
func (m *MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
