package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder provides operations to call the getByPath method.
type ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilderInternal instantiates a new GetByPathWithPathRequestBuilder and sets the default values.
func NewItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, path *string)(*ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder) {
    m := &ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/microsoft.graph.getByPath(path='{path}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if path != nil {
        urlTplParams["path"] = *path
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder instantiates a new GetByPathWithPathRequestBuilder and sets the default values.
func NewItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getByPath
func (m *ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable), nil
}
// ToGetRequestInformation invoke function getByPath
func (m *ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
