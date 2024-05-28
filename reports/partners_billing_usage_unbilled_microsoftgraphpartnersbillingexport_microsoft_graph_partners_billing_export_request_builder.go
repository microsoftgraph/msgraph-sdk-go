package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ieaa1d050ea8ba883c482e05cf2306cb5376cc6e2cf5966c1a6850c42c6118fa4 "github.com/microsoftgraph/msgraph-sdk-go/models/partners/billing"
)

// PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder provides operations to call the export method.
type PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderInternal instantiates a new PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder and sets the default values.
func NewPartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) {
    m := &PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/partners/billing/usage/unbilled/microsoft.graph.partners.billing.export", pathParameters),
    }
    return m
}
// NewPartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder instantiates a new PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder and sets the default values.
func NewPartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post export unbilled Azure usage data for a specific billing period and currency.
// returns a Operationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/partners-billing-unbilledusage-export?view=graph-rest-1.0
func (m *PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) Post(ctx context.Context, body PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportExportPostRequestBodyable, requestConfiguration *PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration)(ieaa1d050ea8ba883c482e05cf2306cb5376cc6e2cf5966c1a6850c42c6118fa4.Operationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieaa1d050ea8ba883c482e05cf2306cb5376cc6e2cf5966c1a6850c42c6118fa4.CreateOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieaa1d050ea8ba883c482e05cf2306cb5376cc6e2cf5966c1a6850c42c6118fa4.Operationable), nil
}
// ToPostRequestInformation export unbilled Azure usage data for a specific billing period and currency.
// returns a *RequestInformation when successful
func (m *PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportExportPostRequestBodyable, requestConfiguration *PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder when successful
func (m *PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) WithUrl(rawUrl string)(*PartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) {
    return NewPartnersBillingUsageUnbilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
