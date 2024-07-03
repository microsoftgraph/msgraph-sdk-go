package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveRequestBuilder provides operations to manage the drive property of the microsoft.graph.fileStorageContainer entity.
type FileStorageContainersItemDriveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainersItemDriveRequestBuilderGetQueryParameters the drive of the resource fileStorageContainer. Read-only.
type FileStorageContainersItemDriveRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemDriveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveRequestBuilderGetQueryParameters
}
// FileStorageContainersItemDriveRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bundles provides operations to manage the bundles property of the microsoft.graph.drive entity.
// returns a *FileStorageContainersItemDriveBundlesRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) Bundles()(*FileStorageContainersItemDriveBundlesRequestBuilder) {
    return NewFileStorageContainersItemDriveBundlesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveRequestBuilderInternal instantiates a new FileStorageContainersItemDriveRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveRequestBuilder) {
    m := &FileStorageContainersItemDriveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveRequestBuilder instantiates a new FileStorageContainersItemDriveRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageContainersItemDriveCreatedByUserRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) CreatedByUser()(*FileStorageContainersItemDriveCreatedByUserRequestBuilder) {
    return NewFileStorageContainersItemDriveCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property drive for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveRequestBuilderDeleteRequestConfiguration)(error) {
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
// Following provides operations to manage the following property of the microsoft.graph.drive entity.
// returns a *FileStorageContainersItemDriveFollowingRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) Following()(*FileStorageContainersItemDriveFollowingRequestBuilder) {
    return NewFileStorageContainersItemDriveFollowingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the drive of the resource fileStorageContainer. Read-only.
// returns a Driveable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.drive entity.
// returns a *FileStorageContainersItemDriveItemsRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) Items()(*FileStorageContainersItemDriveItemsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageContainersItemDriveLastModifiedByUserRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) LastModifiedByUser()(*FileStorageContainersItemDriveLastModifiedByUserRequestBuilder) {
    return NewFileStorageContainersItemDriveLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// List provides operations to manage the list property of the microsoft.graph.drive entity.
// returns a *FileStorageContainersItemDriveListRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) List()(*FileStorageContainersItemDriveListRequestBuilder) {
    return NewFileStorageContainersItemDriveListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property drive in storage
// returns a Driveable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, requestConfiguration *FileStorageContainersItemDriveRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable), nil
}
// Recent provides operations to call the recent method.
// returns a *FileStorageContainersItemDriveRecentRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) Recent()(*FileStorageContainersItemDriveRecentRequestBuilder) {
    return NewFileStorageContainersItemDriveRecentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Root provides operations to manage the root property of the microsoft.graph.drive entity.
// returns a *FileStorageContainersItemDriveRootRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) Root()(*FileStorageContainersItemDriveRootRequestBuilder) {
    return NewFileStorageContainersItemDriveRootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchWithQ provides operations to call the search method.
// returns a *FileStorageContainersItemDriveSearchWithQRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) SearchWithQ(q *string)(*FileStorageContainersItemDriveSearchWithQRequestBuilder) {
    return NewFileStorageContainersItemDriveSearchWithQRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, q)
}
// SharedWithMe provides operations to call the sharedWithMe method.
// returns a *FileStorageContainersItemDriveSharedWithMeRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) SharedWithMe()(*FileStorageContainersItemDriveSharedWithMeRequestBuilder) {
    return NewFileStorageContainersItemDriveSharedWithMeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Special provides operations to manage the special property of the microsoft.graph.drive entity.
// returns a *FileStorageContainersItemDriveSpecialRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) Special()(*FileStorageContainersItemDriveSpecialRequestBuilder) {
    return NewFileStorageContainersItemDriveSpecialRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property drive for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the drive of the resource fileStorageContainer. Read-only.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property drive in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, requestConfiguration *FileStorageContainersItemDriveRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveRequestBuilder when successful
func (m *FileStorageContainersItemDriveRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveRequestBuilder) {
    return NewFileStorageContainersItemDriveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
