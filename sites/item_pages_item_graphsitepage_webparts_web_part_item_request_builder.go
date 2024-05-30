package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder provides operations to manage the webParts property of the microsoft.graph.sitePage entity.
type ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderGetQueryParameters collection of webparts on the SharePoint page.
type ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderGetQueryParameters
}
// ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderInternal instantiates a new ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) {
    m := &ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/webParts/{webPart%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder instantiates a new ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a webPart object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/webpart-delete?view=graph-rest-1.0
func (m *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get collection of webparts on the SharePoint page.
// returns a WebPartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWebPartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable), nil
}
// GetPositionOfWebPart provides operations to call the getPositionOfWebPart method.
// returns a *ItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) GetPositionOfWebPart()(*ItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    return NewItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property webParts in sites
// returns a WebPartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, requestConfiguration *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWebPartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable), nil
}
// ToDeleteRequestInformation delete a webPart object.
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of webparts on the SharePoint page.
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property webParts in sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, requestConfiguration *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder) {
    return NewItemPagesItemGraphsitepageWebpartsWebPartItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
