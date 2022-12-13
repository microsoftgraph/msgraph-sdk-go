package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type ItemRootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemRootRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRootRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemRootRequestBuilderGetQueryParameters retrieve the metadata for a driveItem in a drive by file system path or ID.`item-id` is the ID of a driveItem. It may also be the unique ID of a SharePoint list item.
type ItemRootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRootRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRootRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRootRequestBuilderGetQueryParameters
}
// ItemRootRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRootRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) Analytics()(*ItemRootAnalyticsRequestBuilder) {
    return NewItemRootAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin provides operations to call the checkin method.
func (m *ItemRootRequestBuilder) Checkin()(*ItemRootCheckinRequestBuilder) {
    return NewItemRootCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout provides operations to call the checkout method.
func (m *ItemRootRequestBuilder) Checkout()(*ItemRootCheckoutRequestBuilder) {
    return NewItemRootCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) Children()(*ItemRootChildrenRequestBuilder) {
    return NewItemRootChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) ChildrenById(id string)(*ItemRootChildrenDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewItemRootChildrenDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewItemRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRootRequestBuilder) {
    m := &ItemRootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewItemRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the drive entity.
func (m *ItemRootRequestBuilder) Content()(*ItemRootContentRequestBuilder) {
    return NewItemRootContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *ItemRootRequestBuilder) Copy()(*ItemRootCopyRequestBuilder) {
    return NewItemRootCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for drives
func (m *ItemRootRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemRootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemRootRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemRootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
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
func (m *ItemRootRequestBuilder) CreateLink()(*ItemRootCreateLinkRequestBuilder) {
    return NewItemRootCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property root in drives
func (m *ItemRootRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *ItemRootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateUploadSession provides operations to call the createUploadSession method.
func (m *ItemRootRequestBuilder) CreateUploadSession()(*ItemRootCreateUploadSessionRequestBuilder) {
    return NewItemRootCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for drives
func (m *ItemRootRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemRootRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemRootRequestBuilder) Delta()(*ItemRootDeltaRequestBuilder) {
    return NewItemRootDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *ItemRootRequestBuilder) DeltaWithToken(token *string)(*ItemRootDeltaWithTokenRequestBuilder) {
    return NewItemRootDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// Follow provides operations to call the follow method.
func (m *ItemRootRequestBuilder) Follow()(*ItemRootFollowRequestBuilder) {
    return NewItemRootFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the metadata for a driveItem in a drive by file system path or ID.`item-id` is the ID of a driveItem. It may also be the unique ID of a SharePoint list item.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/driveitem-get?view=graph-rest-1.0
func (m *ItemRootRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRootRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
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
func (m *ItemRootRequestBuilder) GetActivitiesByInterval()(*ItemRootGetActivitiesByIntervalRequestBuilder) {
    return NewItemRootGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemRootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemRootGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemRootGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite provides operations to call the invite method.
func (m *ItemRootRequestBuilder) Invite()(*ItemRootInviteRequestBuilder) {
    return NewItemRootInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) ListItem()(*ItemRootListItemRequestBuilder) {
    return NewItemRootListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in drives
func (m *ItemRootRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *ItemRootRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
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
func (m *ItemRootRequestBuilder) Permissions()(*ItemRootPermissionsRequestBuilder) {
    return NewItemRootPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) PermissionsById(id string)(*ItemRootPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewItemRootPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview provides operations to call the preview method.
func (m *ItemRootRequestBuilder) Preview()(*ItemRootPreviewRequestBuilder) {
    return NewItemRootPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *ItemRootRequestBuilder) Restore()(*ItemRootRestoreRequestBuilder) {
    return NewItemRootRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *ItemRootRequestBuilder) SearchWithQ(q *string)(*ItemRootSearchWithQRequestBuilder) {
    return NewItemRootSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) Subscriptions()(*ItemRootSubscriptionsRequestBuilder) {
    return NewItemRootSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) SubscriptionsById(id string)(*ItemRootSubscriptionsSubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return NewItemRootSubscriptionsSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) Thumbnails()(*ItemRootThumbnailsRequestBuilder) {
    return NewItemRootThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) ThumbnailsById(id string)(*ItemRootThumbnailsThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return NewItemRootThumbnailsThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow provides operations to call the unfollow method.
func (m *ItemRootRequestBuilder) Unfollow()(*ItemRootUnfollowRequestBuilder) {
    return NewItemRootUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *ItemRootRequestBuilder) ValidatePermission()(*ItemRootValidatePermissionRequestBuilder) {
    return NewItemRootValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) Versions()(*ItemRootVersionsRequestBuilder) {
    return NewItemRootVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemRootRequestBuilder) VersionsById(id string)(*ItemRootVersionsDriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return NewItemRootVersionsDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
