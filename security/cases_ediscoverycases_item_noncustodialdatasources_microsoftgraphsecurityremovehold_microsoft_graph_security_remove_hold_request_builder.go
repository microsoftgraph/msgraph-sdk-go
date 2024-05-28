package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder provides operations to call the removeHold method.
type CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    m := &CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/microsoft.graph.security.removeHold", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action removeHold
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) Post(ctx context.Context, body CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdRemoveHoldPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action removeHold
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdRemoveHoldPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
