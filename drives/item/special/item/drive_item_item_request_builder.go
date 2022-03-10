package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i2332bd8c8dc48236f7d5e920135a8191956f518d77854c3114bd2c1ac21c22bc "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/analytics"
    i317a601c3a6ac36bf484b65d5efe6abd1162bb3713f8bbd2f80b5d3e9d630e68 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/listitem"
    i5a02dc3a6d7bdd8984fe2d5f31bfc6f23b910db189e5a5a7d16d821e92b0cb90 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/children"
    i76209ff0c72fddbb47a0939909eb0286899bd7a679fa22cbcda664f1cdb08ca2 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/subscriptions"
    i9bc044cb5d2b45138a4dbcf051071d66fb1e769e2fb86a60c1fbdb0892370225 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/versions"
    i9f7890c2f1e68ca275270f4cc6918695242510c1392ead7bd6003b900e8479b2 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/content"
    ibcf2edd0a9b026b521679e290144aa990d15c2e558eb8b81a8bb64f95edf06c6 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/permissions"
    ief8ed661b353d52760c4e10394cf7c63798d8d5946189b4475949528bd973d5b "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/thumbnails"
    i1506abd30263be2ce02d22772e7d1cc90c84c7de3b5236ac842dafaebc016170 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/permissions/item"
    i1e61da277605ae5b539b1381173379b94b68bf2e129f2c9217994535e8641edd "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/thumbnails/item"
    ic6573023f09fd581736bc8a981a41db25be82d435f900767cf5e6889316cdc44 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/children/item"
    idbb9b67d8a8b376db8cdd8097fbcbe8a577dd240eac155ab17646f33e4bb33c5 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/subscriptions/item"
    iff2ee26348e4509ddbe3900e67b0bf1fc186801ab1d002ee28cf5f8334fbf51a "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item/versions/item"
)

// DriveItemItemRequestBuilder provides operations to manage the special property of the microsoft.graph.drive entity.
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
// DriveItemItemRequestBuilderGetQueryParameters collection of common folders available in OneDrive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) Analytics()(*i2332bd8c8dc48236f7d5e920135a8191956f518d77854c3114bd2c1ac21c22bc.AnalyticsRequestBuilder) {
    return i2332bd8c8dc48236f7d5e920135a8191956f518d77854c3114bd2c1ac21c22bc.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i5a02dc3a6d7bdd8984fe2d5f31bfc6f23b910db189e5a5a7d16d821e92b0cb90.ChildrenRequestBuilder) {
    return i5a02dc3a6d7bdd8984fe2d5f31bfc6f23b910db189e5a5a7d16d821e92b0cb90.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.special.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*ic6573023f09fd581736bc8a981a41db25be82d435f900767cf5e6889316cdc44.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return ic6573023f09fd581736bc8a981a41db25be82d435f900767cf5e6889316cdc44.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/special/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i9f7890c2f1e68ca275270f4cc6918695242510c1392ead7bd6003b900e8479b2.ContentRequestBuilder) {
    return i9f7890c2f1e68ca275270f4cc6918695242510c1392ead7bd6003b900e8479b2.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property special for drives
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
// CreateGetRequestInformation collection of common folders available in OneDrive. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property special in drives
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
// Delete delete navigation property special for drives
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of common folders available in OneDrive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable), nil
}
func (m *DriveItemItemRequestBuilder) ListItem()(*i317a601c3a6ac36bf484b65d5efe6abd1162bb3713f8bbd2f80b5d3e9d630e68.ListItemRequestBuilder) {
    return i317a601c3a6ac36bf484b65d5efe6abd1162bb3713f8bbd2f80b5d3e9d630e68.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property special in drives
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveItemItemRequestBuilder) Permissions()(*ibcf2edd0a9b026b521679e290144aa990d15c2e558eb8b81a8bb64f95edf06c6.PermissionsRequestBuilder) {
    return ibcf2edd0a9b026b521679e290144aa990d15c2e558eb8b81a8bb64f95edf06c6.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.special.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i1506abd30263be2ce02d22772e7d1cc90c84c7de3b5236ac842dafaebc016170.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i1506abd30263be2ce02d22772e7d1cc90c84c7de3b5236ac842dafaebc016170.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i76209ff0c72fddbb47a0939909eb0286899bd7a679fa22cbcda664f1cdb08ca2.SubscriptionsRequestBuilder) {
    return i76209ff0c72fddbb47a0939909eb0286899bd7a679fa22cbcda664f1cdb08ca2.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.special.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*idbb9b67d8a8b376db8cdd8097fbcbe8a577dd240eac155ab17646f33e4bb33c5.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return idbb9b67d8a8b376db8cdd8097fbcbe8a577dd240eac155ab17646f33e4bb33c5.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*ief8ed661b353d52760c4e10394cf7c63798d8d5946189b4475949528bd973d5b.ThumbnailsRequestBuilder) {
    return ief8ed661b353d52760c4e10394cf7c63798d8d5946189b4475949528bd973d5b.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.special.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i1e61da277605ae5b539b1381173379b94b68bf2e129f2c9217994535e8641edd.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i1e61da277605ae5b539b1381173379b94b68bf2e129f2c9217994535e8641edd.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i9bc044cb5d2b45138a4dbcf051071d66fb1e769e2fb86a60c1fbdb0892370225.VersionsRequestBuilder) {
    return i9bc044cb5d2b45138a4dbcf051071d66fb1e769e2fb86a60c1fbdb0892370225.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.special.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*iff2ee26348e4509ddbe3900e67b0bf1fc186801ab1d002ee28cf5f8334fbf51a.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return iff2ee26348e4509ddbe3900e67b0bf1fc186801ab1d002ee28cf5f8334fbf51a.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
