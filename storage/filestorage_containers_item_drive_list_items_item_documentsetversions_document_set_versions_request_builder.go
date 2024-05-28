package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
type FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetQueryParameters version information for a document set version created by a user.
type FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDocumentSetVersionId provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
// returns a *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) ByDocumentSetVersionId(documentSetVersionId string)(*FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if documentSetVersionId != "" {
        urlTplParams["documentSetVersion%2Did"] = documentSetVersionId
    }
    return NewFilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) {
    m := &FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/items/{listItem%2Did}/documentSetVersions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder instantiates a new FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *FilestorageContainersItemDriveListItemsItemDocumentsetversionsCountRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) Count()(*FilestorageContainersItemDriveListItemsItemDocumentsetversionsCountRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemDocumentsetversionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get version information for a document set version created by a user.
// returns a DocumentSetVersionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDocumentSetVersionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionCollectionResponseable), nil
}
// Post create new navigation property to documentSetVersions for storage
// returns a DocumentSetVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionable, requestConfiguration *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation version information for a document set version created by a user.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to documentSetVersions for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DocumentSetVersionable, requestConfiguration *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
