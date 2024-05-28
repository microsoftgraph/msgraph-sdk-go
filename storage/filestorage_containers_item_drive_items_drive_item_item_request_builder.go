package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
// returns a *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Analytics()(*FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
// returns a *FilestorageContainersItemDriveItemsItemAssignsensitivitylabelAssignSensitivityLabelRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) AssignSensitivityLabel()(*FilestorageContainersItemDriveItemsItemAssignsensitivitylabelAssignSensitivityLabelRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAssignsensitivitylabelAssignSensitivityLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkin provides operations to call the checkin method.
// returns a *FilestorageContainersItemDriveItemsItemCheckinRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Checkin()(*FilestorageContainersItemDriveItemsItemCheckinRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemCheckinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkout provides operations to call the checkout method.
// returns a *FilestorageContainersItemDriveItemsItemCheckoutRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Checkout()(*FilestorageContainersItemDriveItemsItemCheckoutRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemCheckoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
// returns a *FilestorageContainersItemDriveItemsItemChildrenRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Children()(*FilestorageContainersItemDriveItemsItemChildrenRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveItemsDriveItemItemRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsDriveItemItemRequestBuilder instantiates a new FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the storage entity.
// returns a *FilestorageContainersItemDriveItemsItemContentRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Content()(*FilestorageContainersItemDriveItemsItemContentRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copy provides operations to call the copy method.
// returns a *FilestorageContainersItemDriveItemsItemCopyRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Copy()(*FilestorageContainersItemDriveItemsItemCopyRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageContainersItemDriveItemsItemCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) CreatedByUser()(*FilestorageContainersItemDriveItemsItemCreatedbyuserCreatedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateLink provides operations to call the createLink method.
// returns a *FilestorageContainersItemDriveItemsItemCreatelinkCreateLinkRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) CreateLink()(*FilestorageContainersItemDriveItemsItemCreatelinkCreateLinkRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemCreatelinkCreateLinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
// returns a *FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) CreateUploadSession()(*FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property items for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *FilestorageContainersItemDriveItemsItemDeltaRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Delta()(*FilestorageContainersItemDriveItemsItemDeltaRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeltaWithToken provides operations to call the delta method.
// returns a *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) DeltaWithToken(token *string)(*FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, token)
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
// returns a *FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) ExtractSensitivityLabels()(*FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Follow provides operations to call the follow method.
// returns a *FilestorageContainersItemDriveItemsItemFollowRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Follow()(*FilestorageContainersItemDriveItemsItemFollowRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemFollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get all items contained in the drive. Read-only. Nullable.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) GetActivitiesByInterval()(*FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
// returns a *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, interval, startDateTime)
}
// Invite provides operations to call the invite method.
// returns a *FilestorageContainersItemDriveItemsItemInviteRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Invite()(*FilestorageContainersItemDriveItemsItemInviteRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemInviteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageContainersItemDriveItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) LastModifiedByUser()(*FilestorageContainersItemDriveItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
// returns a *FilestorageContainersItemDriveItemsItemListitemListItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) ListItem()(*FilestorageContainersItemDriveItemsItemListitemListItemRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemListitemListItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property items in storage
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemPermanentdeletePermanentDeleteRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) PermanentDelete()(*FilestorageContainersItemDriveItemsItemPermanentdeletePermanentDeleteRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemPermanentdeletePermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
// returns a *FilestorageContainersItemDriveItemsItemPermissionsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Permissions()(*FilestorageContainersItemDriveItemsItemPermissionsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Preview provides operations to call the preview method.
// returns a *FilestorageContainersItemDriveItemsItemPreviewRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Preview()(*FilestorageContainersItemDriveItemsItemPreviewRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *FilestorageContainersItemDriveItemsItemRestoreRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Restore()(*FilestorageContainersItemDriveItemsItemRestoreRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetentionLabel provides operations to manage the retentionLabel property of the microsoft.graph.driveItem entity.
// returns a *FilestorageContainersItemDriveItemsItemRetentionlabelRetentionLabelRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) RetentionLabel()(*FilestorageContainersItemDriveItemsItemRetentionlabelRetentionLabelRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemRetentionlabelRetentionLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchWithQ provides operations to call the search method.
// returns a *FilestorageContainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) SearchWithQ(q *string)(*FilestorageContainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, q)
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
// returns a *FilestorageContainersItemDriveItemsItemSubscriptionsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Subscriptions()(*FilestorageContainersItemDriveItemsItemSubscriptionsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
// returns a *FilestorageContainersItemDriveItemsItemThumbnailsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Thumbnails()(*FilestorageContainersItemDriveItemsItemThumbnailsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemThumbnailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property items for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Unfollow()(*FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemUnfollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ValidatePermission provides operations to call the validatePermission method.
// returns a *FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) ValidatePermission()(*FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
// returns a *FilestorageContainersItemDriveItemsItemVersionsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Versions()(*FilestorageContainersItemDriveItemsItemVersionsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsDriveItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Workbook provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsDriveItemItemRequestBuilder) Workbook()(*FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
