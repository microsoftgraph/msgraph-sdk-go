package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersFileStorageContainerItemRequestBuilder provides operations to manage the containers property of the microsoft.graph.fileStorage entity.
type FilestorageContainersFileStorageContainerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersFileStorageContainerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersFileStorageContainerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersFileStorageContainerItemRequestBuilderGetQueryParameters get containers from storage
type FilestorageContainersFileStorageContainerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersFileStorageContainerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersFileStorageContainerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersFileStorageContainerItemRequestBuilderGetQueryParameters
}
// FilestorageContainersFileStorageContainerItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersFileStorageContainerItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate provides operations to call the activate method.
// returns a *FilestorageContainersItemActivateRequestBuilder when successful
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) Activate()(*FilestorageContainersItemActivateRequestBuilder) {
    return NewFilestorageContainersItemActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersFileStorageContainerItemRequestBuilderInternal instantiates a new FilestorageContainersFileStorageContainerItemRequestBuilder and sets the default values.
func NewFilestorageContainersFileStorageContainerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersFileStorageContainerItemRequestBuilder) {
    m := &FilestorageContainersFileStorageContainerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersFileStorageContainerItemRequestBuilder instantiates a new FilestorageContainersFileStorageContainerItemRequestBuilder and sets the default values.
func NewFilestorageContainersFileStorageContainerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersFileStorageContainerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersFileStorageContainerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property containers for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersFileStorageContainerItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Drive provides operations to manage the drive property of the microsoft.graph.fileStorageContainer entity.
// returns a *FilestorageContainersItemDriveRequestBuilder when successful
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) Drive()(*FilestorageContainersItemDriveRequestBuilder) {
    return NewFilestorageContainersItemDriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get containers from storage
// returns a FileStorageContainerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersFileStorageContainerItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFileStorageContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable), nil
}
// Patch update the navigation property containers in storage
// returns a FileStorageContainerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable, requestConfiguration *FilestorageContainersFileStorageContainerItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFileStorageContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.fileStorageContainer entity.
// returns a *FilestorageContainersItemPermissionsRequestBuilder when successful
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) Permissions()(*FilestorageContainersItemPermissionsRequestBuilder) {
    return NewFilestorageContainersItemPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property containers for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersFileStorageContainerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get containers from storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersFileStorageContainerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property containers in storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable, requestConfiguration *FilestorageContainersFileStorageContainerItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersFileStorageContainerItemRequestBuilder when successful
func (m *FilestorageContainersFileStorageContainerItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersFileStorageContainerItemRequestBuilder) {
    return NewFilestorageContainersFileStorageContainerItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
