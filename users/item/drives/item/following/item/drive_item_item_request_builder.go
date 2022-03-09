package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i1cda82cfaae0a3cc6eddb225c1216dd9520a5d594553f7feadcc4b6625c05785 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/analytics"
    i35dc5305cab82195d969ead2ac38d11b65f3adc07c524bfb9efa6d3a639f9486 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/permissions"
    i49288f1df0874f85447722b2d9ff52e8204dfd2fd069e9a783b708bc9f1331ff "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/subscriptions"
    iad37b4d9c70499d8f843c6c1616175c6b845a207b47b602bb52a729e8ec9ff91 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/versions"
    iae41ae2cba884454c122b024c26ac108220d7f6ee021067a4eeb90485c11536d "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/listitem"
    ie5f21009653891393f9ce6010e9273e1955ab814d98260959e2ba42206089568 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/thumbnails"
    iee135bd7d1ad1ef0928027ed7d3a65f9fde697679d5d73a11ce414e5a89a399c "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/children"
    ief6c96796315dc3815001870e9a2913fab943c96c7f32f8556da52088b88b394 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/content"
    i56aee357e799b7f34c517c1fab6b06640ae7c973ae614e3ea368aa4fe0a6bee0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/permissions/item"
    i91f48ed028b803b9d6e1d8ff988e6e731c6d4f527b7053b847d07d939a09f8ec "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/children/item"
    ie7cbd4a3828c12f0686ace564fbe24375bda27075b62a5538a1062d43d02adfd "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/versions/item"
    ieabe5bb8b1892d255940537e6ab7e80d413332086ab93a2e54a93b8964ceda58 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/thumbnails/item"
    iff9870cd370c80ba052d182c48935133f4f3be0ea11545f0457fa7409d831368 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/following/item/subscriptions/item"
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
func (m *DriveItemItemRequestBuilder) Analytics()(*i1cda82cfaae0a3cc6eddb225c1216dd9520a5d594553f7feadcc4b6625c05785.AnalyticsRequestBuilder) {
    return i1cda82cfaae0a3cc6eddb225c1216dd9520a5d594553f7feadcc4b6625c05785.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*iee135bd7d1ad1ef0928027ed7d3a65f9fde697679d5d73a11ce414e5a89a399c.ChildrenRequestBuilder) {
    return iee135bd7d1ad1ef0928027ed7d3a65f9fde697679d5d73a11ce414e5a89a399c.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.following.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i91f48ed028b803b9d6e1d8ff988e6e731c6d4f527b7053b847d07d939a09f8ec.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i91f48ed028b803b9d6e1d8ff988e6e731c6d4f527b7053b847d07d939a09f8ec.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/following/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*ief6c96796315dc3815001870e9a2913fab943c96c7f32f8556da52088b88b394.ContentRequestBuilder) {
    return ief6c96796315dc3815001870e9a2913fab943c96c7f32f8556da52088b88b394.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property following for users
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
// CreatePatchRequestInformation update the navigation property following in users
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
// Delete delete navigation property following for users
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
// Get the list of items the user is following. Only in OneDrive for Business.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*iae41ae2cba884454c122b024c26ac108220d7f6ee021067a4eeb90485c11536d.ListItemRequestBuilder) {
    return iae41ae2cba884454c122b024c26ac108220d7f6ee021067a4eeb90485c11536d.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property following in users
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i35dc5305cab82195d969ead2ac38d11b65f3adc07c524bfb9efa6d3a639f9486.PermissionsRequestBuilder) {
    return i35dc5305cab82195d969ead2ac38d11b65f3adc07c524bfb9efa6d3a639f9486.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.following.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i56aee357e799b7f34c517c1fab6b06640ae7c973ae614e3ea368aa4fe0a6bee0.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i56aee357e799b7f34c517c1fab6b06640ae7c973ae614e3ea368aa4fe0a6bee0.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i49288f1df0874f85447722b2d9ff52e8204dfd2fd069e9a783b708bc9f1331ff.SubscriptionsRequestBuilder) {
    return i49288f1df0874f85447722b2d9ff52e8204dfd2fd069e9a783b708bc9f1331ff.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.following.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*iff9870cd370c80ba052d182c48935133f4f3be0ea11545f0457fa7409d831368.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return iff9870cd370c80ba052d182c48935133f4f3be0ea11545f0457fa7409d831368.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*ie5f21009653891393f9ce6010e9273e1955ab814d98260959e2ba42206089568.ThumbnailsRequestBuilder) {
    return ie5f21009653891393f9ce6010e9273e1955ab814d98260959e2ba42206089568.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.following.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ieabe5bb8b1892d255940537e6ab7e80d413332086ab93a2e54a93b8964ceda58.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return ieabe5bb8b1892d255940537e6ab7e80d413332086ab93a2e54a93b8964ceda58.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*iad37b4d9c70499d8f843c6c1616175c6b845a207b47b602bb52a729e8ec9ff91.VersionsRequestBuilder) {
    return iad37b4d9c70499d8f843c6c1616175c6b845a207b47b602bb52a729e8ec9ff91.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.following.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ie7cbd4a3828c12f0686ace564fbe24375bda27075b62a5538a1062d43d02adfd.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ie7cbd4a3828c12f0686ace564fbe24375bda27075b62a5538a1062d43d02adfd.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
