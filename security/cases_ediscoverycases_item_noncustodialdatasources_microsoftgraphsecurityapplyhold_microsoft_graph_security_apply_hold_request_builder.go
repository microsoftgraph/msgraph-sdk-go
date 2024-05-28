package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder provides operations to call the applyHold method.
type CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    m := &CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/microsoft.graph.security.applyHold", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyHold
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) Post(ctx context.Context, body CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdApplyHoldPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action applyHold
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdApplyHoldPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
