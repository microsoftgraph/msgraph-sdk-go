package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemDrivesItemItemsDriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type ItemDrivesItemItemsDriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemDrivesItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDrivesItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemDrivesItemItemsDriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type ItemDrivesItemItemsDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemDrivesItemItemsDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDrivesItemItemsDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDrivesItemItemsDriveItemItemRequestBuilderGetQueryParameters
}
// ItemDrivesItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDrivesItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Analytics()(*ItemDrivesItemItemsItemAnalyticsRequestBuilder) {
    return NewItemDrivesItemItemsItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin provides operations to call the checkin method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Checkin()(*ItemDrivesItemItemsItemCheckinRequestBuilder) {
    return NewItemDrivesItemItemsItemCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout provides operations to call the checkout method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Checkout()(*ItemDrivesItemItemsItemCheckoutRequestBuilder) {
    return NewItemDrivesItemItemsItemCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Children()(*ItemDrivesItemItemsItemChildrenRequestBuilder) {
    return NewItemDrivesItemItemsItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) ChildrenById(id string)(*ItemDrivesItemItemsItemChildrenDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return NewItemDrivesItemItemsItemChildrenDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemDrivesItemItemsDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewItemDrivesItemItemsDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDrivesItemItemsDriveItemItemRequestBuilder) {
    m := &ItemDrivesItemItemsDriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/drives/{drive%2Did}/items/{driveItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemDrivesItemItemsDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewItemDrivesItemItemsDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDrivesItemItemsDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDrivesItemItemsDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the group entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Content()(*ItemDrivesItemItemsItemContentRequestBuilder) {
    return NewItemDrivesItemItemsItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Copy()(*ItemDrivesItemItemsItemCopyRequestBuilder) {
    return NewItemDrivesItemItemsItemCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateLink provides operations to call the createLink method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) CreateLink()(*ItemDrivesItemItemsItemCreateLinkRequestBuilder) {
    return NewItemDrivesItemItemsItemCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateUploadSession provides operations to call the createUploadSession method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) CreateUploadSession()(*ItemDrivesItemItemsItemCreateUploadSessionRequestBuilder) {
    return NewItemDrivesItemItemsItemCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property items for groups
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemDrivesItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Delta provides operations to call the delta method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Delta()(*ItemDrivesItemItemsItemDeltaRequestBuilder) {
    return NewItemDrivesItemItemsItemDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) DeltaWithToken(token *string)(*ItemDrivesItemItemsItemDeltaWithTokenRequestBuilder) {
    return NewItemDrivesItemItemsItemDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// Follow provides operations to call the follow method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Follow()(*ItemDrivesItemItemsItemFollowRequestBuilder) {
    return NewItemDrivesItemItemsItemFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the drive. Read-only. Nullable.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDrivesItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// GetActivitiesByInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) GetActivitiesByInterval()(*ItemDrivesItemItemsItemGetActivitiesByIntervalRequestBuilder) {
    return NewItemDrivesItemItemsItemGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemDrivesItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemDrivesItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite provides operations to call the invite method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Invite()(*ItemDrivesItemItemsItemInviteRequestBuilder) {
    return NewItemDrivesItemItemsItemInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) ListItem()(*ItemDrivesItemItemsItemListItemRequestBuilder) {
    return NewItemDrivesItemItemsItemListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in groups
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *ItemDrivesItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Permissions()(*ItemDrivesItemItemsItemPermissionsRequestBuilder) {
    return NewItemDrivesItemItemsItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) PermissionsById(id string)(*ItemDrivesItemItemsItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewItemDrivesItemItemsItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview provides operations to call the preview method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Preview()(*ItemDrivesItemItemsItemPreviewRequestBuilder) {
    return NewItemDrivesItemItemsItemPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Restore()(*ItemDrivesItemItemsItemRestoreRequestBuilder) {
    return NewItemDrivesItemItemsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) SearchWithQ(q *string)(*ItemDrivesItemItemsItemSearchWithQRequestBuilder) {
    return NewItemDrivesItemItemsItemSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Subscriptions()(*ItemDrivesItemItemsItemSubscriptionsRequestBuilder) {
    return NewItemDrivesItemItemsItemSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) SubscriptionsById(id string)(*ItemDrivesItemItemsItemSubscriptionsSubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return NewItemDrivesItemItemsItemSubscriptionsSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Thumbnails()(*ItemDrivesItemItemsItemThumbnailsRequestBuilder) {
    return NewItemDrivesItemItemsItemThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) ThumbnailsById(id string)(*ItemDrivesItemItemsItemThumbnailsThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return NewItemDrivesItemItemsItemThumbnailsThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToDeleteRequestInformation delete navigation property items for groups
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemDrivesItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation all items contained in the drive. Read-only. Nullable.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDrivesItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property items in groups
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *ItemDrivesItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Unfollow provides operations to call the unfollow method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Unfollow()(*ItemDrivesItemItemsItemUnfollowRequestBuilder) {
    return NewItemDrivesItemItemsItemUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) ValidatePermission()(*ItemDrivesItemItemsItemValidatePermissionRequestBuilder) {
    return NewItemDrivesItemItemsItemValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) Versions()(*ItemDrivesItemItemsItemVersionsRequestBuilder) {
    return NewItemDrivesItemItemsItemVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemItemsDriveItemItemRequestBuilder) VersionsById(id string)(*ItemDrivesItemItemsItemVersionsDriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return NewItemDrivesItemItemsItemVersionsDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
