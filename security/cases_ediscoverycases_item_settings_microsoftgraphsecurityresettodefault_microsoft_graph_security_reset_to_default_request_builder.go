package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder provides operations to call the resetToDefault method.
type CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder) {
    m := &CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/settings/microsoft.graph.security.resetToDefault", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder instantiates a new CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset a caseSettings object to the default values.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycasesettings-resettodefault?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reset a caseSettings object to the default values.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder) {
    return NewCasesEdiscoverycasesItemSettingsMicrosoftgraphsecurityresettodefaultMicrosoftGraphSecurityResetToDefaultRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
