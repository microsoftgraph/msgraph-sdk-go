package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder provides operations to call the applyTags method.
type CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder) {
    m := &CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}/microsoft.graph.security.applyTags", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder instantiates a new CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post apply tags to files in an eDiscovery review set. For details, see Tag documents in a review set in eDiscovery.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewsetquery-applytags?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder) Post(ctx context.Context, body CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsApplyTagsPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation apply tags to files in an eDiscovery review set. For details, see Tag documents in a review set in eDiscovery.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsApplyTagsPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
