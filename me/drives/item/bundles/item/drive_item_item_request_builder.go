package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i04c509e64f95c836e16edb18bcdb3cf4fef921946a4242b1d0545eaf4d7e7184 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/analytics"
    i260fdbaa1575bee3c708c92f7bd31fa90f0de284de33c14ecd9e7501a1058e44 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/thumbnails"
    i2b5cd74eb16c23b495dfdac8f550135097311a9d4e7591c0b4450729d6bf24c6 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/listitem"
    i84ebba09883f9cd90ed039449dd4cb52a7b8ec6cab8e56d11d966f5b165ecf4b "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/versions"
    i8eb04ff7cf094a0d7467b9c99cb8f421cb4ebe51f72abe2ffb3a9b340a4afa1b "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/subscriptions"
    i928940598df1d4614098d1a0bb66c835aeee94494871247c9b03479d166738b5 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/content"
    i999e1b9227e3ddec34373d34e5f47774a42ef890e53c7c6cc95c0905805dd35b "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/permissions"
    id3c5f4de366e3560569e111388510acc7970726c8eeb84425bad9e05f365af3a "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/children"
    i10f5e1395edfa5a8ebb6c6eeb6934bc1ee86a79a1e4ed926e7a39cb7458bbe3f "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/children/item"
    i392f5ae3c8f654fb7df4ab08d25515615ac32e5112c5f4757e624e0573628431 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/subscriptions/item"
    i935d18124e79fd9550a900d8c341aad3989511bea29ea2e680d9c01a36314417 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/thumbnails/item"
    id9e8d3832206c94e302524b98708c43e6b2754dd2cd9d6b3c668258c63ba3502 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/permissions/item"
    if3f714e8196f2cab0155d1951fb8fd9f8988378ae14bd44c6a7fc8743f711dcf "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item/versions/item"
)

// DriveItemItemRequestBuilder provides operations to manage the bundles property of the microsoft.graph.drive entity.
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
// DriveItemItemRequestBuilderGetQueryParameters collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
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
func (m *DriveItemItemRequestBuilder) Analytics()(*i04c509e64f95c836e16edb18bcdb3cf4fef921946a4242b1d0545eaf4d7e7184.AnalyticsRequestBuilder) {
    return i04c509e64f95c836e16edb18bcdb3cf4fef921946a4242b1d0545eaf4d7e7184.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*id3c5f4de366e3560569e111388510acc7970726c8eeb84425bad9e05f365af3a.ChildrenRequestBuilder) {
    return id3c5f4de366e3560569e111388510acc7970726c8eeb84425bad9e05f365af3a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.bundles.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i10f5e1395edfa5a8ebb6c6eeb6934bc1ee86a79a1e4ed926e7a39cb7458bbe3f.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i10f5e1395edfa5a8ebb6c6eeb6934bc1ee86a79a1e4ed926e7a39cb7458bbe3f.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/bundles/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i928940598df1d4614098d1a0bb66c835aeee94494871247c9b03479d166738b5.ContentRequestBuilder) {
    return i928940598df1d4614098d1a0bb66c835aeee94494871247c9b03479d166738b5.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property bundles for me
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
// CreateGetRequestInformation collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
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
// CreatePatchRequestInformation update the navigation property bundles in me
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
// Delete delete navigation property bundles for me
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
// Get collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i2b5cd74eb16c23b495dfdac8f550135097311a9d4e7591c0b4450729d6bf24c6.ListItemRequestBuilder) {
    return i2b5cd74eb16c23b495dfdac8f550135097311a9d4e7591c0b4450729d6bf24c6.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bundles in me
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i999e1b9227e3ddec34373d34e5f47774a42ef890e53c7c6cc95c0905805dd35b.PermissionsRequestBuilder) {
    return i999e1b9227e3ddec34373d34e5f47774a42ef890e53c7c6cc95c0905805dd35b.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.bundles.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*id9e8d3832206c94e302524b98708c43e6b2754dd2cd9d6b3c668258c63ba3502.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return id9e8d3832206c94e302524b98708c43e6b2754dd2cd9d6b3c668258c63ba3502.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i8eb04ff7cf094a0d7467b9c99cb8f421cb4ebe51f72abe2ffb3a9b340a4afa1b.SubscriptionsRequestBuilder) {
    return i8eb04ff7cf094a0d7467b9c99cb8f421cb4ebe51f72abe2ffb3a9b340a4afa1b.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.bundles.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i392f5ae3c8f654fb7df4ab08d25515615ac32e5112c5f4757e624e0573628431.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i392f5ae3c8f654fb7df4ab08d25515615ac32e5112c5f4757e624e0573628431.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i260fdbaa1575bee3c708c92f7bd31fa90f0de284de33c14ecd9e7501a1058e44.ThumbnailsRequestBuilder) {
    return i260fdbaa1575bee3c708c92f7bd31fa90f0de284de33c14ecd9e7501a1058e44.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.bundles.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i935d18124e79fd9550a900d8c341aad3989511bea29ea2e680d9c01a36314417.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i935d18124e79fd9550a900d8c341aad3989511bea29ea2e680d9c01a36314417.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i84ebba09883f9cd90ed039449dd4cb52a7b8ec6cab8e56d11d966f5b165ecf4b.VersionsRequestBuilder) {
    return i84ebba09883f9cd90ed039449dd4cb52a7b8ec6cab8e56d11d966f5b165ecf4b.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.bundles.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*if3f714e8196f2cab0155d1951fb8fd9f8988378ae14bd44c6a7fc8743f711dcf.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return if3f714e8196f2cab0155d1951fb8fd9f8988378ae14bd44c6a7fc8743f711dcf.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
