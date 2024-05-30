package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder provides operations to call the export method.
type CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) {
    m := &CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/microsoft.graph.security.export", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder instantiates a new CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate an export from a ediscoveryReviewSet. For details, see Export documents from a review set in eDiscovery (Premium).
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewset-export?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) Post(ctx context.Context, body CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation initiate an export from a ediscoveryReviewSet. For details, see Export documents from a review set in eDiscovery (Premium).
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
