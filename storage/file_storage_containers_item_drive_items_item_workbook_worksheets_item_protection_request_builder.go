package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder provides operations to manage the protection property of the microsoft.graph.workbookWorksheet entity.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderGetQueryParameters returns sheet protection object for a worksheet. Read-only.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderGetQueryParameters
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/protection{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property protection for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns sheet protection object for a worksheet. Read-only.
// returns a WorkbookWorksheetProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookWorksheetProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetProtectionable), nil
}
// Patch update the navigation property protection in storage
// returns a WorkbookWorksheetProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetProtectionable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetProtectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookWorksheetProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetProtectionable), nil
}
// Protect provides operations to call the protect method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionProtectRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) Protect()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionProtectRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionProtectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property protection for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns sheet protection object for a worksheet. Read-only.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property protection in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetProtectionable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unprotect provides operations to call the unprotect method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionUnprotectRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) Unprotect()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionUnprotectRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionUnprotectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
