// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ieaa1d050ea8ba883c482e05cf2306cb5376cc6e2cf5966c1a6850c42c6118fa4 "github.com/microsoftgraph/msgraph-sdk-go/models/partners/billing"
)

// PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder provides operations to call the export method.
type PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilderInternal instantiates a new PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder and sets the default values.
func NewPartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder) {
    m := &PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/partners/billing/reconciliation/unbilled/microsoft.graph.partners.billing.export", pathParameters),
    }
    return m
}
// NewPartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder instantiates a new PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder and sets the default values.
func NewPartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post export the unbilled invoice reconciliation data for a specific billing period and a given currency.
// returns a Operationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/partners-billing-unbilledreconciliation-export?view=graph-rest-1.0
func (m *PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder) Post(ctx context.Context, body PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportExportPostRequestBodyable, requestConfiguration *PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration)(ieaa1d050ea8ba883c482e05cf2306cb5376cc6e2cf5966c1a6850c42c6118fa4.Operationable, error) {
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
// ToPostRequestInformation export the unbilled invoice reconciliation data for a specific billing period and a given currency.
// returns a *RequestInformation when successful
func (m *PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportExportPostRequestBodyable, requestConfiguration *PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder when successful
func (m *PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder) WithUrl(rawUrl string)(*PartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder) {
    return NewPartnersBillingReconciliationUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
