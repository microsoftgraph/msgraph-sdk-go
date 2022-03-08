package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i347de3ec81ee72cada957f9274a2049b35b0a5adeeea551ab42adea127d3020a "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/versions"
    i5304ce2548f901afe419817e19c86188c2a91f438d46d99102bb441eeff4f0b3 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/content"
    i5911408c9fb6d0b7b5506ed43692ffa0e9ef7ea4aa6bb400252473412d679648 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/subscriptions"
    i697fe94c20cf8aa9d9ca8670aa74113ca512fc7e962733fbc87546eebac86d99 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/thumbnails"
    i746d7e15a556f6f04633a940990ca8cd5733284011c498aa2f3efda3d02ec850 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/listitem"
    ibaffe2d11367d67304962b88ba02d998906cb5eb0f54d8ddbb61bee317476e87 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/analytics"
    ibe02574c8f615741fc45ce6218a6cadf885d6fe1213b9a2728f6c2560f64bc8b "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/children"
    ifab0ee1ab8b0ac70ace5f3a9a8a6be20f9f7bcaa0a4a597c12b0193c63a01bfc "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/permissions"
    i41715f4164032be0448074d9b13e0a1f3deea014767e99c2401d7c0c2be460fa "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/permissions/item"
    i43d5a4f8297869aa1d898f64c47af83c1f2f56e21d6cbe49173818f05cdb50e6 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/versions/item"
    i6074b6299f116dcb8c8e4a9411037fe415721114dd52ebb0cc1f250577d97877 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/subscriptions/item"
    ibbd03e76ea11da931ba857caee75fe5915f32400c809adf19ab65c9715024303 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/thumbnails/item"
    ic1d7a0ae14cc9523d32ba0aac77590208358c9eb1f33aa8aae89b53a722a2f34 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item/children/item"
)

// DriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
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
// DriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) Analytics()(*ibaffe2d11367d67304962b88ba02d998906cb5eb0f54d8ddbb61bee317476e87.AnalyticsRequestBuilder) {
    return ibaffe2d11367d67304962b88ba02d998906cb5eb0f54d8ddbb61bee317476e87.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*ibe02574c8f615741fc45ce6218a6cadf885d6fe1213b9a2728f6c2560f64bc8b.ChildrenRequestBuilder) {
    return ibe02574c8f615741fc45ce6218a6cadf885d6fe1213b9a2728f6c2560f64bc8b.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*ic1d7a0ae14cc9523d32ba0aac77590208358c9eb1f33aa8aae89b53a722a2f34.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return ic1d7a0ae14cc9523d32ba0aac77590208358c9eb1f33aa8aae89b53a722a2f34.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/items/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i5304ce2548f901afe419817e19c86188c2a91f438d46d99102bb441eeff4f0b3.ContentRequestBuilder) {
    return i5304ce2548f901afe419817e19c86188c2a91f438d46d99102bb441eeff4f0b3.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for me
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
// CreateGetRequestInformation all items contained in the drive. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property items in me
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
// Delete delete navigation property items for me
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
// Get all items contained in the drive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i746d7e15a556f6f04633a940990ca8cd5733284011c498aa2f3efda3d02ec850.ListItemRequestBuilder) {
    return i746d7e15a556f6f04633a940990ca8cd5733284011c498aa2f3efda3d02ec850.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in me
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
func (m *DriveItemItemRequestBuilder) Permissions()(*ifab0ee1ab8b0ac70ace5f3a9a8a6be20f9f7bcaa0a4a597c12b0193c63a01bfc.PermissionsRequestBuilder) {
    return ifab0ee1ab8b0ac70ace5f3a9a8a6be20f9f7bcaa0a4a597c12b0193c63a01bfc.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i41715f4164032be0448074d9b13e0a1f3deea014767e99c2401d7c0c2be460fa.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i41715f4164032be0448074d9b13e0a1f3deea014767e99c2401d7c0c2be460fa.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i5911408c9fb6d0b7b5506ed43692ffa0e9ef7ea4aa6bb400252473412d679648.SubscriptionsRequestBuilder) {
    return i5911408c9fb6d0b7b5506ed43692ffa0e9ef7ea4aa6bb400252473412d679648.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i6074b6299f116dcb8c8e4a9411037fe415721114dd52ebb0cc1f250577d97877.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i6074b6299f116dcb8c8e4a9411037fe415721114dd52ebb0cc1f250577d97877.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i697fe94c20cf8aa9d9ca8670aa74113ca512fc7e962733fbc87546eebac86d99.ThumbnailsRequestBuilder) {
    return i697fe94c20cf8aa9d9ca8670aa74113ca512fc7e962733fbc87546eebac86d99.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ibbd03e76ea11da931ba857caee75fe5915f32400c809adf19ab65c9715024303.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return ibbd03e76ea11da931ba857caee75fe5915f32400c809adf19ab65c9715024303.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i347de3ec81ee72cada957f9274a2049b35b0a5adeeea551ab42adea127d3020a.VersionsRequestBuilder) {
    return i347de3ec81ee72cada957f9274a2049b35b0a5adeeea551ab42adea127d3020a.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i43d5a4f8297869aa1d898f64c47af83c1f2f56e21d6cbe49173818f05cdb50e6.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i43d5a4f8297869aa1d898f64c47af83c1f2f56e21d6cbe49173818f05cdb50e6.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
