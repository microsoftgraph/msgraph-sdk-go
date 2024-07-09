package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder provides operations to manage the special property of the microsoft.graph.drive entity.
type FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderGetQueryParameters collection of common folders available in OneDrive. Read-only. Nullable.
type FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderGetQueryParameters
}
// NewFileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderInternal instantiates a new FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder) {
    m := &FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/special/{driveItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder instantiates a new FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the storage entity.
// returns a *FileStorageContainersItemDriveSpecialItemContentRequestBuilder when successful
func (m *FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder) Content()(*FileStorageContainersItemDriveSpecialItemContentRequestBuilder) {
    return NewFileStorageContainersItemDriveSpecialItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of common folders available in OneDrive. Read-only. Nullable.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// ToGetRequestInformation collection of common folders available in OneDrive. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder when successful
func (m *FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder) {
    return NewFileStorageContainersItemDriveSpecialDriveItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
