package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListRequestBuilder provides operations to manage the list property of the microsoft.graph.drive entity.
type FilestorageContainersItemDriveListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveListRequestBuilderGetQueryParameters for drives in SharePoint, the underlying document library list. Read-only. Nullable.
type FilestorageContainersItemDriveListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveListRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveListRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Columns provides operations to manage the columns property of the microsoft.graph.list entity.
// returns a *FilestorageContainersItemDriveListColumnsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) Columns()(*FilestorageContainersItemDriveListColumnsRequestBuilder) {
    return NewFilestorageContainersItemDriveListColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveListRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListRequestBuilder) {
    m := &FilestorageContainersItemDriveListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListRequestBuilder instantiates a new FilestorageContainersItemDriveListRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.list entity.
// returns a *FilestorageContainersItemDriveListContenttypesContentTypesRequestBuilder when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) ContentTypes()(*FilestorageContainersItemDriveListContenttypesContentTypesRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesContentTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) CreatedByUser()(*FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property list for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListRequestBuilderDeleteRequestConfiguration)(error) {
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
// Drive provides operations to manage the drive property of the microsoft.graph.list entity.
// returns a *FilestorageContainersItemDriveListDriveRequestBuilder when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) Drive()(*FilestorageContainersItemDriveListDriveRequestBuilder) {
    return NewFilestorageContainersItemDriveListDriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get for drives in SharePoint, the underlying document library list. Read-only. Nullable.
// returns a Listable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.list entity.
// returns a *FilestorageContainersItemDriveListItemsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) Items()(*FilestorageContainersItemDriveListItemsRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageContainersItemDriveListLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) LastModifiedByUser()(*FilestorageContainersItemDriveListLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveListLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.list entity.
// returns a *FilestorageContainersItemDriveListOperationsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) Operations()(*FilestorageContainersItemDriveListOperationsRequestBuilder) {
    return NewFilestorageContainersItemDriveListOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property list in storage
// returns a Listable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable, requestConfiguration *FilestorageContainersItemDriveListRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable), nil
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.list entity.
// returns a *FilestorageContainersItemDriveListSubscriptionsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) Subscriptions()(*FilestorageContainersItemDriveListSubscriptionsRequestBuilder) {
    return NewFilestorageContainersItemDriveListSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property list for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation for drives in SharePoint, the underlying document library list. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property list in storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable, requestConfiguration *FilestorageContainersItemDriveListRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveListRequestBuilder when successful
func (m *FilestorageContainersItemDriveListRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListRequestBuilder) {
    return NewFilestorageContainersItemDriveListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
