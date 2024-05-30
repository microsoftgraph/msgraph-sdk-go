package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder provides operations to call the applyHold method.
type CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    m := &CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/microsoft.graph.security.applyHold", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder instantiates a new CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyHold
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) Post(ctx context.Context, body CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdApplyHoldPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(error) {
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
func (m *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdApplyHoldPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
