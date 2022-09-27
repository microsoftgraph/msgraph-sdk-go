package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i020cefc69abd9d840cd84135fe23fb1bebc2c214f437d47672da3069688efb65 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/searchwithq"
    i1173f8f46182265fd51edd487d8eb3467c7304d5472add2aa3d5d0685d556c1b "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/restore"
    i277bda5a06f6ce182900160c095232c5b28b7d856a168f362c33560d2bc3cc63 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i4052aa75e096d48638bddada7f1cb8659fe3c25b1f0ad818070488e257ccbee4 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/createuploadsession"
    i413c7d3c5296661c95b513a28ba587f270281832de23c7984752b4a218f6019e "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/invite"
    i4c90a5b5af67ed60ab6adf5613edbcb1da1c9dcd2ea15d3712d431fabae295d2 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/permissions"
    i5258e8be91847b76242f45b58489af959fe0ae7cdbd06dca3e1f74c3073635e9 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/deltawithtoken"
    i6a156ef471bfddb5199d753c1547d8cfe0fed10a62269f71e40433faf09abb3c "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/checkin"
    i6d048e988ecd7a469df6e77aceb776b58b4b96a45ec000a5ade64874f4ef46af "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/subscriptions"
    i821430851f47c1fcc1bd23b2399a85e1051e8b3c866ed232120d55300cdc1128 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/validatepermission"
    i8de7c056a112237127b0dcc4ea8d6badfd1aba50e70ba16d561038ac278428e1 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/thumbnails"
    i8ed13773ce508ebebaa60ab9a3303a3ca02078bf14b6afa3519a1293144b0a57 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/versions"
    i93d215297749f08e6a96baa42da7b41e3023fd3e6eb044d22abab4d04759c6b5 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/preview"
    i9734f4cbac8b25e2d6dca093ff4a3a99bbdedd24599e80f4d0a511d780c8af84 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/unfollow"
    i977cc3182db94374759f872d55be1bd164691f41d3ef9934b2a1b2ee3a5022b9 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/listitem"
    ia4d81037bc2b75bd1c85541272c51f69d1e647b3c24dd0fe76b688dc66888981 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/checkout"
    ibad2dbd640f8438d505ec69c54f24d1b51b0428157e3df1a56b77ca2b617c307 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/content"
    ibcc5304fa8c81ec8fc2559cf01425b91f72a11b2fe9749fc949b5246a6102146 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/copy"
    ic187abdb8f33a4cb81256b996a715ff597202bed85f8f4d64b6e8319f5022059 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/children"
    ica3c1d93f94f1770bb123fd333161fa4d87f9717f224df595e97bf876ca7967e "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/createlink"
    ie46cb62de11b25a34a5de7350395445ff217862bc353808a40cd5e906b64da5a "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/getactivitiesbyinterval"
    if28afa82f9d9e53dd623400a2be39b487c3f2c06369ddb07d71922a511a981f9 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/delta"
    if6e4b55b696c1aaead2907344730bf6b12310b52afb14a9ffd1ce42161325da9 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/analytics"
    ife52184499776c3c829efc37896f9c7578be8da410d81bb1866e5c91e78e9e4c "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/follow"
    i01b95b3d57747e53543e8dadcce70427ca7796d7df1fa1a45329a65eae314015 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/versions/item"
    i0a1467d2d422249bfaa9b52d3bde594deab07162016808ecd41602cbd68904ce "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/permissions/item"
    i170fcbbf415e96234b1aa6424878a65c5d7b82ec53c144e5203f73656e625891 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/subscriptions/item"
    i33b94f1f1fae365d235ffc10e8627c15caaca7be62b15dfd7f8cb6a23342e856 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/children/item"
    icf7102c66807ffb1835310fb1cb3f4be4c0c9d50a0fe651aac0d3390144ccfad "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/thumbnails/item"
)

// DriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type DriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type DriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DriveItemItemRequestBuilderGetQueryParameters
}
// DriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics the analytics property
func (m *DriveItemItemRequestBuilder) Analytics()(*if6e4b55b696c1aaead2907344730bf6b12310b52afb14a9ffd1ce42161325da9.AnalyticsRequestBuilder) {
    return if6e4b55b696c1aaead2907344730bf6b12310b52afb14a9ffd1ce42161325da9.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *DriveItemItemRequestBuilder) Checkin()(*i6a156ef471bfddb5199d753c1547d8cfe0fed10a62269f71e40433faf09abb3c.CheckinRequestBuilder) {
    return i6a156ef471bfddb5199d753c1547d8cfe0fed10a62269f71e40433faf09abb3c.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *DriveItemItemRequestBuilder) Checkout()(*ia4d81037bc2b75bd1c85541272c51f69d1e647b3c24dd0fe76b688dc66888981.CheckoutRequestBuilder) {
    return ia4d81037bc2b75bd1c85541272c51f69d1e647b3c24dd0fe76b688dc66888981.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *DriveItemItemRequestBuilder) Children()(*ic187abdb8f33a4cb81256b996a715ff597202bed85f8f4d64b6e8319f5022059.ChildrenRequestBuilder) {
    return ic187abdb8f33a4cb81256b996a715ff597202bed85f8f4d64b6e8319f5022059.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i33b94f1f1fae365d235ffc10e8627c15caaca7be62b15dfd7f8cb6a23342e856.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return i33b94f1f1fae365d235ffc10e8627c15caaca7be62b15dfd7f8cb6a23342e856.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/items/{driveItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *DriveItemItemRequestBuilder) Content()(*ibad2dbd640f8438d505ec69c54f24d1b51b0428157e3df1a56b77ca2b617c307.ContentRequestBuilder) {
    return ibad2dbd640f8438d505ec69c54f24d1b51b0428157e3df1a56b77ca2b617c307.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *DriveItemItemRequestBuilder) Copy()(*ibcc5304fa8c81ec8fc2559cf01425b91f72a11b2fe9749fc949b5246a6102146.CopyRequestBuilder) {
    return ibcc5304fa8c81ec8fc2559cf01425b91f72a11b2fe9749fc949b5246a6102146.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for drive
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property items for drive
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DriveItemItemRequestBuilder) CreateLink()(*ica3c1d93f94f1770bb123fd333161fa4d87f9717f224df595e97bf876ca7967e.CreateLinkRequestBuilder) {
    return ica3c1d93f94f1770bb123fd333161fa4d87f9717f224df595e97bf876ca7967e.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in drive
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property items in drive
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DriveItemItemRequestBuilder) CreateUploadSession()(*i4052aa75e096d48638bddada7f1cb8659fe3c25b1f0ad818070488e257ccbee4.CreateUploadSessionRequestBuilder) {
    return i4052aa75e096d48638bddada7f1cb8659fe3c25b1f0ad818070488e257ccbee4.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property items for drive
func (m *DriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DriveItemItemRequestBuilder) Delta()(*if28afa82f9d9e53dd623400a2be39b487c3f2c06369ddb07d71922a511a981f9.DeltaRequestBuilder) {
    return if28afa82f9d9e53dd623400a2be39b487c3f2c06369ddb07d71922a511a981f9.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *DriveItemItemRequestBuilder) DeltaWithToken(token *string)(*i5258e8be91847b76242f45b58489af959fe0ae7cdbd06dca3e1f74c3073635e9.DeltaWithTokenRequestBuilder) {
    return i5258e8be91847b76242f45b58489af959fe0ae7cdbd06dca3e1f74c3073635e9.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// Follow the follow property
func (m *DriveItemItemRequestBuilder) Follow()(*ife52184499776c3c829efc37896f9c7578be8da410d81bb1866e5c91e78e9e4c.FollowRequestBuilder) {
    return ife52184499776c3c829efc37896f9c7578be8da410d81bb1866e5c91e78e9e4c.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
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
func (m *DriveItemItemRequestBuilder) GetActivitiesByInterval()(*ie46cb62de11b25a34a5de7350395445ff217862bc353808a40cd5e906b64da5a.GetActivitiesByIntervalRequestBuilder) {
    return ie46cb62de11b25a34a5de7350395445ff217862bc353808a40cd5e906b64da5a.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *DriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i277bda5a06f6ce182900160c095232c5b28b7d856a168f362c33560d2bc3cc63.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i277bda5a06f6ce182900160c095232c5b28b7d856a168f362c33560d2bc3cc63.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite the invite property
func (m *DriveItemItemRequestBuilder) Invite()(*i413c7d3c5296661c95b513a28ba587f270281832de23c7984752b4a218f6019e.InviteRequestBuilder) {
    return i413c7d3c5296661c95b513a28ba587f270281832de23c7984752b4a218f6019e.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *DriveItemItemRequestBuilder) ListItem()(*i977cc3182db94374759f872d55be1bd164691f41d3ef9934b2a1b2ee3a5022b9.ListItemRequestBuilder) {
    return i977cc3182db94374759f872d55be1bd164691f41d3ef9934b2a1b2ee3a5022b9.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in drive
func (m *DriveItemItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i4c90a5b5af67ed60ab6adf5613edbcb1da1c9dcd2ea15d3712d431fabae295d2.PermissionsRequestBuilder) {
    return i4c90a5b5af67ed60ab6adf5613edbcb1da1c9dcd2ea15d3712d431fabae295d2.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i0a1467d2d422249bfaa9b52d3bde594deab07162016808ecd41602cbd68904ce.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i0a1467d2d422249bfaa9b52d3bde594deab07162016808ecd41602cbd68904ce.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *DriveItemItemRequestBuilder) Preview()(*i93d215297749f08e6a96baa42da7b41e3023fd3e6eb044d22abab4d04759c6b5.PreviewRequestBuilder) {
    return i93d215297749f08e6a96baa42da7b41e3023fd3e6eb044d22abab4d04759c6b5.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *DriveItemItemRequestBuilder) Restore()(*i1173f8f46182265fd51edd487d8eb3467c7304d5472add2aa3d5d0685d556c1b.RestoreRequestBuilder) {
    return i1173f8f46182265fd51edd487d8eb3467c7304d5472add2aa3d5d0685d556c1b.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemItemRequestBuilder) SearchWithQ(q *string)(*i020cefc69abd9d840cd84135fe23fb1bebc2c214f437d47672da3069688efb65.SearchWithQRequestBuilder) {
    return i020cefc69abd9d840cd84135fe23fb1bebc2c214f437d47672da3069688efb65.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i6d048e988ecd7a469df6e77aceb776b58b4b96a45ec000a5ade64874f4ef46af.SubscriptionsRequestBuilder) {
    return i6d048e988ecd7a469df6e77aceb776b58b4b96a45ec000a5ade64874f4ef46af.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i170fcbbf415e96234b1aa6424878a65c5d7b82ec53c144e5203f73656e625891.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i170fcbbf415e96234b1aa6424878a65c5d7b82ec53c144e5203f73656e625891.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i8de7c056a112237127b0dcc4ea8d6badfd1aba50e70ba16d561038ac278428e1.ThumbnailsRequestBuilder) {
    return i8de7c056a112237127b0dcc4ea8d6badfd1aba50e70ba16d561038ac278428e1.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*icf7102c66807ffb1835310fb1cb3f4be4c0c9d50a0fe651aac0d3390144ccfad.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return icf7102c66807ffb1835310fb1cb3f4be4c0c9d50a0fe651aac0d3390144ccfad.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *DriveItemItemRequestBuilder) Unfollow()(*i9734f4cbac8b25e2d6dca093ff4a3a99bbdedd24599e80f4d0a511d780c8af84.UnfollowRequestBuilder) {
    return i9734f4cbac8b25e2d6dca093ff4a3a99bbdedd24599e80f4d0a511d780c8af84.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *DriveItemItemRequestBuilder) ValidatePermission()(*i821430851f47c1fcc1bd23b2399a85e1051e8b3c866ed232120d55300cdc1128.ValidatePermissionRequestBuilder) {
    return i821430851f47c1fcc1bd23b2399a85e1051e8b3c866ed232120d55300cdc1128.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *DriveItemItemRequestBuilder) Versions()(*i8ed13773ce508ebebaa60ab9a3303a3ca02078bf14b6afa3519a1293144b0a57.VersionsRequestBuilder) {
    return i8ed13773ce508ebebaa60ab9a3303a3ca02078bf14b6afa3519a1293144b0a57.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i01b95b3d57747e53543e8dadcce70427ca7796d7df1fa1a45329a65eae314015.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i01b95b3d57747e53543e8dadcce70427ca7796d7df1fa1a45329a65eae314015.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
