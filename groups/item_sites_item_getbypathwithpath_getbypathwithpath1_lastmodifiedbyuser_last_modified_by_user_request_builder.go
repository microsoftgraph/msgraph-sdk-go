package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
type ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderGetQueryParameters identity of the user who last modified the item. Read-only.
type ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderGetQueryParameters
}
// NewItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderInternal instantiates a new ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    m := &ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/lastModifiedByUser{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder instantiates a new ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get identity of the user who last modified the item. Read-only.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// ToGetRequestInformation identity of the user who last modified the item. Read-only.
// returns a *RequestInformation when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
