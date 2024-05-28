package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder provides operations to manage the parentNotebook property of the microsoft.graph.onenoteSection entity.
type ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderGetQueryParameters the notebook that contains the section.  Read-only.
type ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderGetQueryParameters
}
// NewItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderInternal instantiates a new ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder) {
    m := &ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/sections/{onenoteSection%2Did}/parentNotebook{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder instantiates a new ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the notebook that contains the section.  Read-only.
// returns a Notebookable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Notebookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateNotebookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Notebookable), nil
}
// ToGetRequestInformation the notebook that contains the section.  Read-only.
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder) {
    return NewItemSitesItemOnenoteSectionsItemParentnotebookParentNotebookRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
