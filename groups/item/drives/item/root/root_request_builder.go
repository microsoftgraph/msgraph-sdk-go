package root

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0082410670daef638919e52ce33d639b7b6ea63fd8e9d8e08f3d06b1477a4b3f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/delta"
    i0c944d7cde2fcf588ca052b523e87a74c49ddc3e86ebd26b0a5119bec6f41ca0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/createuploadsession"
    i1404d06852ae31f933f39ab7cd1c6423bda229061d3ef6e592f61708a2d2554b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/analytics"
    i2683147c078f6273eb1a60f1531618bcf66ada7273b5ad2ed06a37aa94d1e0dc "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/copy"
    i26d61efdfa482aac4ed50cfc52e1f5660fc7b4dc4d1545d9fe3f267351496386 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/invite"
    i3c3d2dc90a4d45ba89a1fe68c7a438145d66dfa0f836ab7266ef55d646087126 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/children"
    i61d3c688f098a379da87e2a24785d28417deba346b80d618a9b689330a9ea804 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/unfollow"
    i634bb2571b0cacd8fd3fbe567e64022f43ec60b68442b9c695c06cd1712f12fc "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/thumbnails"
    i653de1968cf538e4cbd997e1c2243d91a2cd183dd3d96a161448e115529705fa "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/content"
    i74527195d199bdffd41f1d994acb52433f15c00f369e0f17bab6842f9aa60737 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/validatepermission"
    i77ec29a690464da49931d35abc08f8231d99a5a205a4959708a764b545b39caa "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/restore"
    i7a72600b407dc115ad5510c6b902866ed3c770e1d99da81feba8c78c29f643db "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/versions"
    i841ff79e037b2788b1924595c099ff799ce213e48c1dcce85c506fa052931d79 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/permissions"
    i86dfe3c82f7f2a36e71e9ebf2cf41415bbcbee9a6747e54b671ce0bcfb1e630f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/listitem"
    i8cc34195b9995e5434b4ce17672018549295334561bba17b731c8fa49654a747 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/checkout"
    i8f62022515317d7e72ed20f4b57475c6bb4b57a99dc994e8cd67feaf64817c5d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/createlink"
    i985d0e8fef0e465708a6633b299f61ed0881a5fcce55307950e3b25d1bd6728b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/follow"
    ia6c6ad50a86683840e630a2c7059e913084401af16fdbf4b17b77348066f6a48 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/getactivitiesbyinterval"
    ib21af4cf6809a4b31cf62e0b926a8ad87dfa897d6069113bd49bbaa7ae87a7ef "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ibfb26e769e14288ce7267db8e8d0a7ea7ba9d2636ded1a5fcf7b20d2ae179424 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/subscriptions"
    ic09ea497c0e0db9cb546554a03887378495c0ee93c4ea92efee63551c17cb915 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/deltawithtoken"
    idbfb688b4f084fd3caf043ade047c046dba7eca87c1cfdb0f95b9166659732b8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/preview"
    idf33142f46f6bf05a63edf7b3efb0a264a47ca945c86d8557dabf550ca8835f2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/searchwithq"
    if71ce21fc5dd5fe82476a1263b022424ec8aed34d59a10713e62b1941195acdf "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/checkin"
    i07b9939ec119ea6111987e32cf3bba70d45517ee87ccac971a870a9f988b18c9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/children/item"
    i45a80764a6f657f684e135272e30ca2f11c3a27841c105a02a27c2b6cf59ae9b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/thumbnails/item"
    i78e358a55c5906dadd43e4729cbb82212ead2f8293a7560b71fb9e1a6d3debcd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/permissions/item"
    ib612109492ba4022df9b33ed40472711c65e7875cfb012ca9e987006140115ab "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/subscriptions/item"
    ib6aec01254a524eb447e9f58964dac48e58f362931827f812b0aed386bfb42f8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/versions/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RootRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RootRequestBuilderGetQueryParameters retrieve the metadata for a driveItem in a drive by file system path or ID.`item-id` is the ID of a driveItem. It may also be the unique ID of a SharePoint list item.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RootRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RootRequestBuilderGetQueryParameters
}
// RootRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*i1404d06852ae31f933f39ab7cd1c6423bda229061d3ef6e592f61708a2d2554b.AnalyticsRequestBuilder) {
    return i1404d06852ae31f933f39ab7cd1c6423bda229061d3ef6e592f61708a2d2554b.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *RootRequestBuilder) Checkin()(*if71ce21fc5dd5fe82476a1263b022424ec8aed34d59a10713e62b1941195acdf.CheckinRequestBuilder) {
    return if71ce21fc5dd5fe82476a1263b022424ec8aed34d59a10713e62b1941195acdf.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *RootRequestBuilder) Checkout()(*i8cc34195b9995e5434b4ce17672018549295334561bba17b731c8fa49654a747.CheckoutRequestBuilder) {
    return i8cc34195b9995e5434b4ce17672018549295334561bba17b731c8fa49654a747.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i3c3d2dc90a4d45ba89a1fe68c7a438145d66dfa0f836ab7266ef55d646087126.ChildrenRequestBuilder) {
    return i3c3d2dc90a4d45ba89a1fe68c7a438145d66dfa0f836ab7266ef55d646087126.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i07b9939ec119ea6111987e32cf3bba70d45517ee87ccac971a870a9f988b18c9.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i07b9939ec119ea6111987e32cf3bba70d45517ee87ccac971a870a9f988b18c9.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *RootRequestBuilder) Content()(*i653de1968cf538e4cbd997e1c2243d91a2cd183dd3d96a161448e115529705fa.ContentRequestBuilder) {
    return i653de1968cf538e4cbd997e1c2243d91a2cd183dd3d96a161448e115529705fa.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *RootRequestBuilder) Copy()(*i2683147c078f6273eb1a60f1531618bcf66ada7273b5ad2ed06a37aa94d1e0dc.CopyRequestBuilder) {
    return i2683147c078f6273eb1a60f1531618bcf66ada7273b5ad2ed06a37aa94d1e0dc.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for groups
func (m *RootRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for groups
func (m *RootRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation retrieve the metadata for a driveItem in a drive by file system path or ID.`item-id` is the ID of a driveItem. It may also be the unique ID of a SharePoint list item.
func (m *RootRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieve the metadata for a driveItem in a drive by file system path or ID.`item-id` is the ID of a driveItem. It may also be the unique ID of a SharePoint list item.
func (m *RootRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateLink the createLink property
func (m *RootRequestBuilder) CreateLink()(*i8f62022515317d7e72ed20f4b57475c6bb4b57a99dc994e8cd67feaf64817c5d.CreateLinkRequestBuilder) {
    return i8f62022515317d7e72ed20f4b57475c6bb4b57a99dc994e8cd67feaf64817c5d.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property root in groups
func (m *RootRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in groups
func (m *RootRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateUploadSession the createUploadSession property
func (m *RootRequestBuilder) CreateUploadSession()(*i0c944d7cde2fcf588ca052b523e87a74c49ddc3e86ebd26b0a5119bec6f41ca0.CreateUploadSessionRequestBuilder) {
    return i0c944d7cde2fcf588ca052b523e87a74c49ddc3e86ebd26b0a5119bec6f41ca0.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for groups
func (m *RootRequestBuilder) Delete(ctx context.Context, requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *RootRequestBuilder) Delta()(*i0082410670daef638919e52ce33d639b7b6ea63fd8e9d8e08f3d06b1477a4b3f.DeltaRequestBuilder) {
    return i0082410670daef638919e52ce33d639b7b6ea63fd8e9d8e08f3d06b1477a4b3f.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *RootRequestBuilder) DeltaWithToken(token *string)(*ic09ea497c0e0db9cb546554a03887378495c0ee93c4ea92efee63551c17cb915.DeltaWithTokenRequestBuilder) {
    return ic09ea497c0e0db9cb546554a03887378495c0ee93c4ea92efee63551c17cb915.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// Follow the follow property
func (m *RootRequestBuilder) Follow()(*i985d0e8fef0e465708a6633b299f61ed0881a5fcce55307950e3b25d1bd6728b.FollowRequestBuilder) {
    return i985d0e8fef0e465708a6633b299f61ed0881a5fcce55307950e3b25d1bd6728b.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the metadata for a driveItem in a drive by file system path or ID.`item-id` is the ID of a driveItem. It may also be the unique ID of a SharePoint list item.
func (m *RootRequestBuilder) Get(ctx context.Context, requestConfiguration *RootRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *RootRequestBuilder) GetActivitiesByInterval()(*ia6c6ad50a86683840e630a2c7059e913084401af16fdbf4b17b77348066f6a48.GetActivitiesByIntervalRequestBuilder) {
    return ia6c6ad50a86683840e630a2c7059e913084401af16fdbf4b17b77348066f6a48.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *RootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ib21af4cf6809a4b31cf62e0b926a8ad87dfa897d6069113bd49bbaa7ae87a7ef.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ib21af4cf6809a4b31cf62e0b926a8ad87dfa897d6069113bd49bbaa7ae87a7ef.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite the invite property
func (m *RootRequestBuilder) Invite()(*i26d61efdfa482aac4ed50cfc52e1f5660fc7b4dc4d1545d9fe3f267351496386.InviteRequestBuilder) {
    return i26d61efdfa482aac4ed50cfc52e1f5660fc7b4dc4d1545d9fe3f267351496386.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *RootRequestBuilder) ListItem()(*i86dfe3c82f7f2a36e71e9ebf2cf41415bbcbee9a6747e54b671ce0bcfb1e630f.ListItemRequestBuilder) {
    return i86dfe3c82f7f2a36e71e9ebf2cf41415bbcbee9a6747e54b671ce0bcfb1e630f.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in groups
func (m *RootRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Permissions the permissions property
func (m *RootRequestBuilder) Permissions()(*i841ff79e037b2788b1924595c099ff799ce213e48c1dcce85c506fa052931d79.PermissionsRequestBuilder) {
    return i841ff79e037b2788b1924595c099ff799ce213e48c1dcce85c506fa052931d79.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i78e358a55c5906dadd43e4729cbb82212ead2f8293a7560b71fb9e1a6d3debcd.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i78e358a55c5906dadd43e4729cbb82212ead2f8293a7560b71fb9e1a6d3debcd.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *RootRequestBuilder) Preview()(*idbfb688b4f084fd3caf043ade047c046dba7eca87c1cfdb0f95b9166659732b8.PreviewRequestBuilder) {
    return idbfb688b4f084fd3caf043ade047c046dba7eca87c1cfdb0f95b9166659732b8.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *RootRequestBuilder) Restore()(*i77ec29a690464da49931d35abc08f8231d99a5a205a4959708a764b545b39caa.RestoreRequestBuilder) {
    return i77ec29a690464da49931d35abc08f8231d99a5a205a4959708a764b545b39caa.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *RootRequestBuilder) SearchWithQ(q *string)(*idf33142f46f6bf05a63edf7b3efb0a264a47ca945c86d8557dabf550ca8835f2.SearchWithQRequestBuilder) {
    return idf33142f46f6bf05a63edf7b3efb0a264a47ca945c86d8557dabf550ca8835f2.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*ibfb26e769e14288ce7267db8e8d0a7ea7ba9d2636ded1a5fcf7b20d2ae179424.SubscriptionsRequestBuilder) {
    return ibfb26e769e14288ce7267db8e8d0a7ea7ba9d2636ded1a5fcf7b20d2ae179424.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*ib612109492ba4022df9b33ed40472711c65e7875cfb012ca9e987006140115ab.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return ib612109492ba4022df9b33ed40472711c65e7875cfb012ca9e987006140115ab.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i634bb2571b0cacd8fd3fbe567e64022f43ec60b68442b9c695c06cd1712f12fc.ThumbnailsRequestBuilder) {
    return i634bb2571b0cacd8fd3fbe567e64022f43ec60b68442b9c695c06cd1712f12fc.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i45a80764a6f657f684e135272e30ca2f11c3a27841c105a02a27c2b6cf59ae9b.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i45a80764a6f657f684e135272e30ca2f11c3a27841c105a02a27c2b6cf59ae9b.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *RootRequestBuilder) Unfollow()(*i61d3c688f098a379da87e2a24785d28417deba346b80d618a9b689330a9ea804.UnfollowRequestBuilder) {
    return i61d3c688f098a379da87e2a24785d28417deba346b80d618a9b689330a9ea804.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *RootRequestBuilder) ValidatePermission()(*i74527195d199bdffd41f1d994acb52433f15c00f369e0f17bab6842f9aa60737.ValidatePermissionRequestBuilder) {
    return i74527195d199bdffd41f1d994acb52433f15c00f369e0f17bab6842f9aa60737.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*i7a72600b407dc115ad5510c6b902866ed3c770e1d99da81feba8c78c29f643db.VersionsRequestBuilder) {
    return i7a72600b407dc115ad5510c6b902866ed3c770e1d99da81feba8c78c29f643db.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*ib6aec01254a524eb447e9f58964dac48e58f362931827f812b0aed386bfb42f8.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return ib6aec01254a524eb447e9f58964dac48e58f362931827f812b0aed386bfb42f8.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
