package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder provides operations to call the updateIndex method.
type CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    m := &CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/microsoft.graph.security.updateIndex", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder instantiates a new CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal(urlParams, requestAdapter)
}
// Post trigger an indexOperation to make a custodian and associated sources searchable.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-updateindex?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation trigger an indexOperation to make a custodian and associated sources searchable.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
