package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder provides operations to call the reopen method.
type CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder) {
    m := &CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/microsoft.graph.security.reopen", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder instantiates a new CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reopen an eDiscovery case that was closed. For details, see Reopen a closed case.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-reopen?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reopen an eDiscovery case that was closed. For details, see Reopen a closed case.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder when successful
func (m *CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder) {
    return NewCasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
