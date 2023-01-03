package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemDrivesItemRootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type ItemDrivesItemRootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemDrivesItemRootRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDrivesItemRootRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemDrivesItemRootRequestBuilderGetQueryParameters retrieve the metadata for a driveItem in a drive by file system path or ID.`item-id` is the ID of a driveItem. It may also be the unique ID of a SharePoint list item.
type ItemDrivesItemRootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemDrivesItemRootRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDrivesItemRootRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDrivesItemRootRequestBuilderGetQueryParameters
}
// ItemDrivesItemRootRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDrivesItemRootRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) Analytics()(*ItemDrivesItemRootAnalyticsRequestBuilder) {
    return NewItemDrivesItemRootAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin provides operations to call the checkin method.
func (m *ItemDrivesItemRootRequestBuilder) Checkin()(*ItemDrivesItemRootCheckinRequestBuilder) {
    return NewItemDrivesItemRootCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout provides operations to call the checkout method.
func (m *ItemDrivesItemRootRequestBuilder) Checkout()(*ItemDrivesItemRootCheckoutRequestBuilder) {
    return NewItemDrivesItemRootCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) Children()(*ItemDrivesItemRootChildrenRequestBuilder) {
    return NewItemDrivesItemRootChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) ChildrenById(id string)(*DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemDrivesItemRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewItemDrivesItemRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDrivesItemRootRequestBuilder) {
    m := &ItemDrivesItemRootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemDrivesItemRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewItemDrivesItemRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDrivesItemRootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDrivesItemRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *ItemDrivesItemRootRequestBuilder) Content()(*ItemDrivesItemRootContentRequestBuilder) {
    return NewItemDrivesItemRootContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *ItemDrivesItemRootRequestBuilder) Copy()(*ItemDrivesItemRootCopyRequestBuilder) {
    return NewItemDrivesItemRootCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for users
func (m *ItemDrivesItemRootRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemDrivesItemRootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the metadata for a driveItem in a drive by file system path or ID.`item-id` is the ID of a driveItem. It may also be the unique ID of a SharePoint list item.
func (m *ItemDrivesItemRootRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemDrivesItemRootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateLink provides operations to call the createLink method.
func (m *ItemDrivesItemRootRequestBuilder) CreateLink()(*ItemDrivesItemRootCreateLinkRequestBuilder) {
    return NewItemDrivesItemRootCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property root in users
func (m *ItemDrivesItemRootRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *ItemDrivesItemRootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateUploadSession provides operations to call the createUploadSession method.
func (m *ItemDrivesItemRootRequestBuilder) CreateUploadSession()(*ItemDrivesItemRootCreateUploadSessionRequestBuilder) {
    return NewItemDrivesItemRootCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for users
func (m *ItemDrivesItemRootRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemDrivesItemRootRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
func (m *ItemDrivesItemRootRequestBuilder) Delta()(*ItemDrivesItemRootDeltaRequestBuilder) {
    return NewItemDrivesItemRootDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *ItemDrivesItemRootRequestBuilder) DeltaWithToken(token *string)(*ItemDrivesItemRootDeltaWithTokenRequestBuilder) {
    return NewItemDrivesItemRootDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// Follow provides operations to call the follow method.
func (m *ItemDrivesItemRootRequestBuilder) Follow()(*ItemDrivesItemRootFollowRequestBuilder) {
    return NewItemDrivesItemRootFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the metadata for a driveItem in a drive by file system path or ID.`item-id` is the ID of a driveItem. It may also be the unique ID of a SharePoint list item.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/driveitem-get?view=graph-rest-1.0
func (m *ItemDrivesItemRootRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDrivesItemRootRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
func (m *ItemDrivesItemRootRequestBuilder) GetActivitiesByInterval()(*ItemDrivesItemRootGetActivitiesByIntervalRequestBuilder) {
    return NewItemDrivesItemRootGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemDrivesItemRootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemDrivesItemRootGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemDrivesItemRootGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite provides operations to call the invite method.
func (m *ItemDrivesItemRootRequestBuilder) Invite()(*ItemDrivesItemRootInviteRequestBuilder) {
    return NewItemDrivesItemRootInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) ListItem()(*ItemDrivesItemRootListItemRequestBuilder) {
    return NewItemDrivesItemRootListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in users
func (m *ItemDrivesItemRootRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *ItemDrivesItemRootRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
func (m *ItemDrivesItemRootRequestBuilder) Permissions()(*ItemDrivesItemRootPermissionsRequestBuilder) {
    return NewItemDrivesItemRootPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) PermissionsById(id string)(*PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview provides operations to call the preview method.
func (m *ItemDrivesItemRootRequestBuilder) Preview()(*ItemDrivesItemRootPreviewRequestBuilder) {
    return NewItemDrivesItemRootPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *ItemDrivesItemRootRequestBuilder) Restore()(*ItemDrivesItemRootRestoreRequestBuilder) {
    return NewItemDrivesItemRootRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *ItemDrivesItemRootRequestBuilder) SearchWithQ(q *string)(*ItemDrivesItemRootSearchWithQRequestBuilder) {
    return NewItemDrivesItemRootSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) Subscriptions()(*ItemDrivesItemRootSubscriptionsRequestBuilder) {
    return NewItemDrivesItemRootSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) SubscriptionsById(id string)(*SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) Thumbnails()(*ItemDrivesItemRootThumbnailsRequestBuilder) {
    return NewItemDrivesItemRootThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) ThumbnailsById(id string)(*ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow provides operations to call the unfollow method.
func (m *ItemDrivesItemRootRequestBuilder) Unfollow()(*ItemDrivesItemRootUnfollowRequestBuilder) {
    return NewItemDrivesItemRootUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *ItemDrivesItemRootRequestBuilder) ValidatePermission()(*ItemDrivesItemRootValidatePermissionRequestBuilder) {
    return NewItemDrivesItemRootValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) Versions()(*ItemDrivesItemRootVersionsRequestBuilder) {
    return NewItemDrivesItemRootVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemDrivesItemRootRequestBuilder) VersionsById(id string)(*DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
