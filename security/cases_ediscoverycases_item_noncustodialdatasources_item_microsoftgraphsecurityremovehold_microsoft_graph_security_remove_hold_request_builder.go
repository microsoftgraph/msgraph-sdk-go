package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder provides operations to call the removeHold method.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    m := &CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/{ediscoveryNoncustodialDataSource%2Did}/microsoft.graph.security.removeHold", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start the process of removing hold from eDiscovery non-custodial data sources. After the operation is created, you can get the status by retrieving the Location parameter from the response headers. The location provides a URL that returns an eDiscoveryHoldOperation object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverynoncustodialdatasource-removehold?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation start the process of removing hold from eDiscovery non-custodial data sources. After the operation is created, you can get the status by retrieving the Location parameter from the response headers. The location provides a URL that returns an eDiscoveryHoldOperation object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
