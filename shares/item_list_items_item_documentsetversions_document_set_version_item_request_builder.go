package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
type ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderGetQueryParameters version information for a document set version created by a user.
type ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderGetQueryParameters
}
// ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderInternal instantiates a new ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder and sets the default values.
func NewItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) {
    m := &ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shares/{sharedDriveItem%2Did}/list/items/{listItem%2Did}/documentSetVersions/{documentSetVersion%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder instantiates a new ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder and sets the default values.
func NewItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property documentSetVersions for shares
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Fields provides operations to manage the fields property of the microsoft.graph.listItemVersion entity.
// returns a *ItemListItemsItemDocumentsetversionsItemFieldsRequestBuilder when successful
func (m *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) Fields()(*ItemListItemsItemDocumentsetversionsItemFieldsRequestBuilder) {
    return NewItemListItemsItemDocumentsetversionsItemFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get version information for a document set version created by a user.
// returns a DocumentSetVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDocumentSetVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionable), nil
}
// Patch update the navigation property documentSetVersions in shares
// returns a DocumentSetVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionable, requestConfiguration *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDocumentSetVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionable), nil
}
// Restore provides operations to call the restore method.
// returns a *ItemListItemsItemDocumentsetversionsItemRestoreRequestBuilder when successful
func (m *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) Restore()(*ItemListItemsItemDocumentsetversionsItemRestoreRequestBuilder) {
    return NewItemListItemsItemDocumentsetversionsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property documentSetVersions for shares
// returns a *RequestInformation when successful
func (m *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation version information for a document set version created by a user.
// returns a *RequestInformation when successful
func (m *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property documentSetVersions in shares
// returns a *RequestInformation when successful
func (m *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionable, requestConfiguration *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder when successful
func (m *ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) WithUrl(rawUrl string)(*ItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) {
    return NewItemListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
