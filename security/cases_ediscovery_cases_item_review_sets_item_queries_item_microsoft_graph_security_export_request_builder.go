package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilder provides operations to call the export method.
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilderInternal instantiates a new MicrosoftGraphSecurityExportRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilder) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}/microsoft.graph.security.export", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilder instantiates a new MicrosoftGraphSecurityExportRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate an export from a ediscoveryReviewSetQuery. For details, see Export documents from a review set in eDiscovery (Premium).
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoveryreviewsetquery-export?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilder) Post(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation initiate an export from a ediscoveryReviewSetQuery. For details, see Export documents from a review set in eDiscovery (Premium).
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
