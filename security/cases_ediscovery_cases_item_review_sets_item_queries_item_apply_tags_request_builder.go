package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder provides operations to call the applyTags method.
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilderInternal instantiates a new ApplyTagsRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}/microsoft.graph.security.applyTags";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder instantiates a new ApplyTagsRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation apply tags to files in an eDiscovery review set. For details, see Tag documents in a review set in eDiscovery.
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post apply tags to files in an eDiscovery review set. For details, see Tag documents in a review set in eDiscovery.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoveryreviewsetquery-applytags?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder) Post(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
