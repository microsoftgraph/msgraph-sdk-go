package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder provides operations to call the close method.
type CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder) {
    m := &CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/microsoft.graph.security.close", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder instantiates a new CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post close an eDiscovery case. For details, see Close a case.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-close?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation close an eDiscovery case. For details, see Close a case.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder when successful
func (m *CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder) {
    return NewCasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
