package root

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1404d06852ae31f933f39ab7cd1c6423bda229061d3ef6e592f61708a2d2554b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/analytics"
    i3c3d2dc90a4d45ba89a1fe68c7a438145d66dfa0f836ab7266ef55d646087126 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/children"
    i634bb2571b0cacd8fd3fbe567e64022f43ec60b68442b9c695c06cd1712f12fc "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/thumbnails"
    i653de1968cf538e4cbd997e1c2243d91a2cd183dd3d96a161448e115529705fa "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/content"
    i7a72600b407dc115ad5510c6b902866ed3c770e1d99da81feba8c78c29f643db "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/versions"
    i841ff79e037b2788b1924595c099ff799ce213e48c1dcce85c506fa052931d79 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/permissions"
    i86dfe3c82f7f2a36e71e9ebf2cf41415bbcbee9a6747e54b671ce0bcfb1e630f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/listitem"
    ibfb26e769e14288ce7267db8e8d0a7ea7ba9d2636ded1a5fcf7b20d2ae179424 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/subscriptions"
    i07b9939ec119ea6111987e32cf3bba70d45517ee87ccac971a870a9f988b18c9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/children/item"
    i45a80764a6f657f684e135272e30ca2f11c3a27841c105a02a27c2b6cf59ae9b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/thumbnails/item"
    i78e358a55c5906dadd43e4729cbb82212ead2f8293a7560b71fb9e1a6d3debcd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/permissions/item"
    ib612109492ba4022df9b33ed40472711c65e7875cfb012ca9e987006140115ab "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/subscriptions/item"
    ib6aec01254a524eb447e9f58964dac48e58f362931827f812b0aed386bfb42f8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root/versions/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RootRequestBuilderDeleteOptions options for Delete
type RootRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RootRequestBuilderGetOptions options for Get
type RootRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RootRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RootRequestBuilderPatchOptions options for Patch
type RootRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RootRequestBuilder) Analytics()(*i1404d06852ae31f933f39ab7cd1c6423bda229061d3ef6e592f61708a2d2554b.AnalyticsRequestBuilder) {
    return i1404d06852ae31f933f39ab7cd1c6423bda229061d3ef6e592f61708a2d2554b.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["driveItem_id"] = id
    }
    return i07b9939ec119ea6111987e32cf3bba70d45517ee87ccac971a870a9f988b18c9.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}/root{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RootRequestBuilder) Content()(*i653de1968cf538e4cbd997e1c2243d91a2cd183dd3d96a161448e115529705fa.ContentRequestBuilder) {
    return i653de1968cf538e4cbd997e1c2243d91a2cd183dd3d96a161448e115529705fa.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for groups
func (m *RootRequestBuilder) CreateDeleteRequestInformation(options *RootRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation(options *RootRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property root in groups
func (m *RootRequestBuilder) CreatePatchRequestInformation(options *RootRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property root for groups
func (m *RootRequestBuilder) Delete(options *RootRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get(options *RootRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable), nil
}
func (m *RootRequestBuilder) ListItem()(*i86dfe3c82f7f2a36e71e9ebf2cf41415bbcbee9a6747e54b671ce0bcfb1e630f.ListItemRequestBuilder) {
    return i86dfe3c82f7f2a36e71e9ebf2cf41415bbcbee9a6747e54b671ce0bcfb1e630f.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in groups
func (m *RootRequestBuilder) Patch(options *RootRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["permission_id"] = id
    }
    return i78e358a55c5906dadd43e4729cbb82212ead2f8293a7560b71fb9e1a6d3debcd.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["subscription_id"] = id
    }
    return ib612109492ba4022df9b33ed40472711c65e7875cfb012ca9e987006140115ab.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["thumbnailSet_id"] = id
    }
    return i45a80764a6f657f684e135272e30ca2f11c3a27841c105a02a27c2b6cf59ae9b.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["driveItemVersion_id"] = id
    }
    return ib6aec01254a524eb447e9f58964dac48e58f362931827f812b0aed386bfb42f8.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
