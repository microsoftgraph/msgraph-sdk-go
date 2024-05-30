package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder provides operations to call the export method.
type CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) {
    m := &CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}/microsoft.graph.security.export", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder instantiates a new CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate an export from a ediscoveryReviewSetQuery. For details, see Export documents from a review set in eDiscovery (Premium).
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewsetquery-export?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) Post(ctx context.Context, body CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportExportPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation initiate an export from a ediscoveryReviewSetQuery. For details, see Export documents from a review set in eDiscovery (Premium).
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportExportPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
