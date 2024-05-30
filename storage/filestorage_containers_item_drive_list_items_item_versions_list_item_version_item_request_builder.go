package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder provides operations to manage the versions property of the microsoft.graph.listItem entity.
type FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderGetQueryParameters the list of previous versions of the list item.
type FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) {
    m := &FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/items/{listItem%2Did}/versions/{listItemVersion%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder instantiates a new FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property versions for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *FilestorageContainersItemDriveListItemsItemVersionsItemFieldsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) Fields()(*FilestorageContainersItemDriveListItemsItemVersionsItemFieldsRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemVersionsItemFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of previous versions of the list item.
// returns a ListItemVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateListItemVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemVersionable), nil
}
// Patch update the navigation property versions in storage
// returns a ListItemVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemVersionable, requestConfiguration *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemVersionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateListItemVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemVersionable), nil
}
// RestoreVersion provides operations to call the restoreVersion method.
// returns a *FilestorageContainersItemDriveListItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) RestoreVersion()(*FilestorageContainersItemDriveListItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property versions for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of previous versions of the list item.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property versions in storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemVersionable, requestConfiguration *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemVersionsListItemVersionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
