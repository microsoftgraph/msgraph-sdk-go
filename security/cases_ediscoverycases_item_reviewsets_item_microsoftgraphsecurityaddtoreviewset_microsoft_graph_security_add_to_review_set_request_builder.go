package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder provides operations to call the addToReviewSet method.
type CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder) {
    m := &CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/microsoft.graph.security.addToReviewSet", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder instantiates a new CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start the process of adding a collection from Microsoft 365 services to a review set. After the operation is created, you can get the status of the operation by retrieving the Location parameter from the response headers. The location provides a URL that will return a Add to review set operation.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewset-addtoreviewset?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder) Post(ctx context.Context, body CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation start the process of adding a collection from Microsoft 365 services to a review set. After the operation is created, you can get the status of the operation by retrieving the Location parameter from the response headers. The location provides a URL that will return a Add to review set operation.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetMicrosoftGraphSecurityAddToReviewSetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
