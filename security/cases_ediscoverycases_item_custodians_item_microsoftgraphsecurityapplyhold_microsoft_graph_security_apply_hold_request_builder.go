package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder provides operations to call the applyHold method.
type CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    m := &CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/microsoft.graph.security.applyHold", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder instantiates a new CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start the process of applying hold on eDiscovery custodians. After the operation is created, you can get the status by retrieving the Location parameter from the response headers. The location provides a URL that will return an eDiscoveryHoldOperation object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-applyhold?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation start the process of applying hold on eDiscovery custodians. After the operation is created, you can get the status by retrieving the Location parameter from the response headers. The location provides a URL that will return an eDiscoveryHoldOperation object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
