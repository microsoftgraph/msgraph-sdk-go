package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder provides operations to call the removeHold method.
type CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilderInternal instantiates a new RemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder) {
    m := &CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/microsoft.graph.security.removeHold";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder instantiates a new RemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation start the process of removing hold from eDiscovery custodians. After the operation is created, you can get the status by retrieving the `Location` parameter from the response headers. The location provides a URL that will return an eDiscoveryHoldOperation object.
func (m *CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post start the process of removing hold from eDiscovery custodians. After the operation is created, you can get the status by retrieving the `Location` parameter from the response headers. The location provides a URL that will return an eDiscoveryHoldOperation object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoverycustodian-removehold?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
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
