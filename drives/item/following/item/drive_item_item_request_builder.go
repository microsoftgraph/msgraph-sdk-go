package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0aaad28c6056577d44c56e4f6ca41e16a4932622b2e90f2d5255fde5ee897e03 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/permissions"
    i24485bce4a2e5137eda5e864f34756aff67a918e1104a222998752c08e353f24 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/subscriptions"
    i3811e7083d98f3346aeb13b52f6b1cd47f750c3344259843eed30986fe674956 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/content"
    i3c1990cfc2ae5fa1c2601063c8d23b72d1b81cb55ce739479bfc3cb0fa1a15a7 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/thumbnails"
    i42616730ee47222962baf33952b334f8366dcb59aa7b3eab4948c1322f3ac708 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/analytics"
    ia4e343e4d0a60031f2bed69eaf429356014f5888dddc10bfd35fcb29c22b2a6d "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/listitem"
    iaa9832423f74f4ecb27f740704ae3c1fdc5ebe2b04b3d7e83840d6d78a79aeda "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/versions"
    if3264640038e5bb1853138efb7d84098c326d162bb5f1250af2247b0ce9a52ae "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/children"
    i0781995521e0f5d2b14b11b3e420cd8c2b5ba92387c34ef876e7d4b29a7e417c "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/thumbnails/item"
    i24151692ba6e304c0f09577c30c07c701ab3fee1367ffd6fd426246494b8d037 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/permissions/item"
    i28fd3cc435d921b531b6ed8322ba8b61635b910a1eb29a1224dcbcc95ae8e82f "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/children/item"
    ibea64976d0e4bc1184492fcf4fc87c07843d128bf7f55188228786f9e087fec0 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/subscriptions/item"
    ic56afe40b31b7037cd6e81717bba9b72e25c8a1761bc6ac871f0c1f1acf7815e "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item/versions/item"
)

// DriveItemItemRequestBuilder provides operations to manage the following property of the microsoft.graph.drive entity.
type DriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemItemRequestBuilderDeleteOptions options for Delete
type DriveItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemItemRequestBuilderGetOptions options for Get
type DriveItemItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemItemRequestBuilderGetQueryParameters the list of items the user is following. Only in OneDrive for Business.
type DriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemItemRequestBuilderPatchOptions options for Patch
type DriveItemItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i42616730ee47222962baf33952b334f8366dcb59aa7b3eab4948c1322f3ac708.AnalyticsRequestBuilder) {
    return i42616730ee47222962baf33952b334f8366dcb59aa7b3eab4948c1322f3ac708.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*if3264640038e5bb1853138efb7d84098c326d162bb5f1250af2247b0ce9a52ae.ChildrenRequestBuilder) {
    return if3264640038e5bb1853138efb7d84098c326d162bb5f1250af2247b0ce9a52ae.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.following.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i28fd3cc435d921b531b6ed8322ba8b61635b910a1eb29a1224dcbcc95ae8e82f.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i28fd3cc435d921b531b6ed8322ba8b61635b910a1eb29a1224dcbcc95ae8e82f.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/following/{driveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DriveItemItemRequestBuilder) Content()(*i3811e7083d98f3346aeb13b52f6b1cd47f750c3344259843eed30986fe674956.ContentRequestBuilder) {
    return i3811e7083d98f3346aeb13b52f6b1cd47f750c3344259843eed30986fe674956.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property following for drives
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of items the user is following. Only in OneDrive for Business.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation(options *DriveItemItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property following in drives
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property following for drives
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of items the user is following. Only in OneDrive for Business.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable), nil
}
func (m *DriveItemItemRequestBuilder) ListItem()(*ia4e343e4d0a60031f2bed69eaf429356014f5888dddc10bfd35fcb29c22b2a6d.ListItemRequestBuilder) {
    return ia4e343e4d0a60031f2bed69eaf429356014f5888dddc10bfd35fcb29c22b2a6d.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property following in drives
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveItemItemRequestBuilder) Permissions()(*i0aaad28c6056577d44c56e4f6ca41e16a4932622b2e90f2d5255fde5ee897e03.PermissionsRequestBuilder) {
    return i0aaad28c6056577d44c56e4f6ca41e16a4932622b2e90f2d5255fde5ee897e03.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.following.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i24151692ba6e304c0f09577c30c07c701ab3fee1367ffd6fd426246494b8d037.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i24151692ba6e304c0f09577c30c07c701ab3fee1367ffd6fd426246494b8d037.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i24485bce4a2e5137eda5e864f34756aff67a918e1104a222998752c08e353f24.SubscriptionsRequestBuilder) {
    return i24485bce4a2e5137eda5e864f34756aff67a918e1104a222998752c08e353f24.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.following.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*ibea64976d0e4bc1184492fcf4fc87c07843d128bf7f55188228786f9e087fec0.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ibea64976d0e4bc1184492fcf4fc87c07843d128bf7f55188228786f9e087fec0.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i3c1990cfc2ae5fa1c2601063c8d23b72d1b81cb55ce739479bfc3cb0fa1a15a7.ThumbnailsRequestBuilder) {
    return i3c1990cfc2ae5fa1c2601063c8d23b72d1b81cb55ce739479bfc3cb0fa1a15a7.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.following.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i0781995521e0f5d2b14b11b3e420cd8c2b5ba92387c34ef876e7d4b29a7e417c.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i0781995521e0f5d2b14b11b3e420cd8c2b5ba92387c34ef876e7d4b29a7e417c.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*iaa9832423f74f4ecb27f740704ae3c1fdc5ebe2b04b3d7e83840d6d78a79aeda.VersionsRequestBuilder) {
    return iaa9832423f74f4ecb27f740704ae3c1fdc5ebe2b04b3d7e83840d6d78a79aeda.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.following.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ic56afe40b31b7037cd6e81717bba9b72e25c8a1761bc6ac871f0c1f1acf7815e.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ic56afe40b31b7037cd6e81717bba9b72e25c8a1761bc6ac871f0c1f1acf7815e.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
