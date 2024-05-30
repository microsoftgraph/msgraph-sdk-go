package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder provides operations to call the release method.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder) {
    m := &CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/{ediscoveryNoncustodialDataSource%2Did}/microsoft.graph.security.release", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post release the non-custodial data source from the case.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverynoncustodialdatasource-release?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation release the non-custodial data source from the case.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
