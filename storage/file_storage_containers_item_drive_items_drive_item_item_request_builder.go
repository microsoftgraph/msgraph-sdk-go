package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters
}
// FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
// returns a *FileStorageContainersItemDriveItemsItemAnalyticsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Analytics()(*FileStorageContainersItemDriveItemsItemAnalyticsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
// returns a *FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) AssignSensitivityLabel()(*FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkin provides operations to call the checkin method.
// returns a *FileStorageContainersItemDriveItemsItemCheckinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Checkin()(*FileStorageContainersItemDriveItemsItemCheckinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemCheckinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkout provides operations to call the checkout method.
// returns a *FileStorageContainersItemDriveItemsItemCheckoutRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Checkout()(*FileStorageContainersItemDriveItemsItemCheckoutRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemCheckoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
// returns a *FileStorageContainersItemDriveItemsItemChildrenRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Children()(*FileStorageContainersItemDriveItemsItemChildrenRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveItemsDriveItemItemRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsDriveItemItemRequestBuilder instantiates a new FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the storage entity.
// returns a *FileStorageContainersItemDriveItemsItemContentRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Content()(*FileStorageContainersItemDriveItemsItemContentRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copy provides operations to call the copy method.
// returns a *FileStorageContainersItemDriveItemsItemCopyRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Copy()(*FileStorageContainersItemDriveItemsItemCopyRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageContainersItemDriveItemsItemCreatedByUserRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) CreatedByUser()(*FileStorageContainersItemDriveItemsItemCreatedByUserRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateLink provides operations to call the createLink method.
// returns a *FileStorageContainersItemDriveItemsItemCreateLinkRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) CreateLink()(*FileStorageContainersItemDriveItemsItemCreateLinkRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemCreateLinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
// returns a *FileStorageContainersItemDriveItemsItemCreateUploadSessionRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) CreateUploadSession()(*FileStorageContainersItemDriveItemsItemCreateUploadSessionRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemCreateUploadSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property items for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Delta provides operations to call the delta method.
// returns a *FileStorageContainersItemDriveItemsItemDeltaRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Delta()(*FileStorageContainersItemDriveItemsItemDeltaRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeltaWithToken provides operations to call the delta method.
// returns a *FileStorageContainersItemDriveItemsItemDeltaWithTokenRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) DeltaWithToken(token *string)(*FileStorageContainersItemDriveItemsItemDeltaWithTokenRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemDeltaWithTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, token)
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
// returns a *FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) ExtractSensitivityLabels()(*FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Follow provides operations to call the follow method.
// returns a *FileStorageContainersItemDriveItemsItemFollowRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Follow()(*FileStorageContainersItemDriveItemsItemFollowRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemFollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get all items contained in the drive. Read-only. Nullable.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
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
// GetActivitiesByInterval provides operations to call the getActivitiesByInterval method.
// returns a *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) GetActivitiesByInterval()(*FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
// returns a *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, interval, startDateTime)
}
// Invite provides operations to call the invite method.
// returns a *FileStorageContainersItemDriveItemsItemInviteRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Invite()(*FileStorageContainersItemDriveItemsItemInviteRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemInviteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageContainersItemDriveItemsItemLastModifiedByUserRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) LastModifiedByUser()(*FileStorageContainersItemDriveItemsItemLastModifiedByUserRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
// returns a *FileStorageContainersItemDriveItemsItemListItemRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) ListItem()(*FileStorageContainersItemDriveItemsItemListItemRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemListItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property items in storage
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *FileStorageContainersItemDriveItemsItemPermanentDeleteRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) PermanentDelete()(*FileStorageContainersItemDriveItemsItemPermanentDeleteRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
// returns a *FileStorageContainersItemDriveItemsItemPermissionsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Permissions()(*FileStorageContainersItemDriveItemsItemPermissionsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Preview provides operations to call the preview method.
// returns a *FileStorageContainersItemDriveItemsItemPreviewRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Preview()(*FileStorageContainersItemDriveItemsItemPreviewRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *FileStorageContainersItemDriveItemsItemRestoreRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Restore()(*FileStorageContainersItemDriveItemsItemRestoreRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetentionLabel provides operations to manage the retentionLabel property of the microsoft.graph.driveItem entity.
// returns a *FileStorageContainersItemDriveItemsItemRetentionLabelRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) RetentionLabel()(*FileStorageContainersItemDriveItemsItemRetentionLabelRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemRetentionLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchWithQ provides operations to call the search method.
// returns a *FileStorageContainersItemDriveItemsItemSearchWithQRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) SearchWithQ(q *string)(*FileStorageContainersItemDriveItemsItemSearchWithQRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemSearchWithQRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, q)
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
// returns a *FileStorageContainersItemDriveItemsItemSubscriptionsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Subscriptions()(*FileStorageContainersItemDriveItemsItemSubscriptionsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
// returns a *FileStorageContainersItemDriveItemsItemThumbnailsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Thumbnails()(*FileStorageContainersItemDriveItemsItemThumbnailsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemThumbnailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property items for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation all items contained in the drive. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property items in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unfollow provides operations to call the unfollow method.
// returns a *FileStorageContainersItemDriveItemsItemUnfollowRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Unfollow()(*FileStorageContainersItemDriveItemsItemUnfollowRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemUnfollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ValidatePermission provides operations to call the validatePermission method.
// returns a *FileStorageContainersItemDriveItemsItemValidatePermissionRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) ValidatePermission()(*FileStorageContainersItemDriveItemsItemValidatePermissionRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemValidatePermissionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
// returns a *FileStorageContainersItemDriveItemsItemVersionsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Versions()(*FileStorageContainersItemDriveItemsItemVersionsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsDriveItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Workbook provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsDriveItemItemRequestBuilder) Workbook()(*FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
