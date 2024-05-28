package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder provides operations to call the getMailboxUsageStorage method.
type GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilderInternal instantiates a new GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder and sets the default values.
func NewGetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder) {
    m := &GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getMailboxUsageStorage(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder instantiates a new GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder and sets the default values.
func NewGetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the amount of storage used in your organization.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getmailboxusagestorage?view=graph-rest-1.0
func (m *GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get the amount of storage used in your organization.
// returns a *RequestInformation when successful
func (m *GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder when successful
func (m *GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder) {
    return NewGetmailboxusagestoragewithperiodGetMailboxUsageStorageWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
