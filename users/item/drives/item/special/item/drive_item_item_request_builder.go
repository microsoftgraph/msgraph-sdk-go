package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i01519f6527cd27de0d0206601e6a99e65e1d02b2cc582b523d34fa09d861e0e2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/analytics"
    i16f286c80a97a30511cfd598f3bc4eb419cd0a68fb00ce06fbfbc67c97e1114f "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/versions"
    i19bd11d1cd6787496925edd7b818a82834d022e827becd10062880fd77aa0a1c "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/children"
    i7b005ff8448f500405ea534a1724e1d349c087dc25092f66a4296dec55221fcc "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/subscriptions"
    i82c51906a8f8e5995521d546d21b7f258d44003c77ada3e6f8a7d0a71494d7be "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/thumbnails"
    ia2e39a4c866ff367b398d30e099d94047a5cc6fe607bf52e8ab0b48d4e65c2bb "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/listitem"
    ib79c0230ea7b5afdffec69d21e9a9441ef851b72d9a0f5dd19923ae63e662025 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/content"
    idd80c801d7860c1d12a236550e4ca9d092abaffbfe2250e9b953e3bd93aae3f2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/permissions"
    i20d350a7d9efb08a78cee6c2c9f0251dacea33b325d63d77796a12151b05934b "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/subscriptions/item"
    i39c0c9e01bc01eff72ee5b18410e29de41c051d189f03d16e7732bcc6aab5018 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/permissions/item"
    i43cd1e815a6d8136f5d9dc6120d6a900872413877d6ef947b189fac6a7359b35 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/children/item"
    i51f749cb530aa4a2c132d8a1ddaa5c29f022af8b45617699c428291ed0f5018e "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/versions/item"
    ib57162d336ed51ae0e881d55b1433a3969e660a3fbd45b028160ef5bc285d72d "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/special/item/thumbnails/item"
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
func (m *DriveItemItemRequestBuilder) Analytics()(*i01519f6527cd27de0d0206601e6a99e65e1d02b2cc582b523d34fa09d861e0e2.AnalyticsRequestBuilder) {
    return i01519f6527cd27de0d0206601e6a99e65e1d02b2cc582b523d34fa09d861e0e2.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i19bd11d1cd6787496925edd7b818a82834d022e827becd10062880fd77aa0a1c.ChildrenRequestBuilder) {
    return i19bd11d1cd6787496925edd7b818a82834d022e827becd10062880fd77aa0a1c.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.special.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i43cd1e815a6d8136f5d9dc6120d6a900872413877d6ef947b189fac6a7359b35.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i43cd1e815a6d8136f5d9dc6120d6a900872413877d6ef947b189fac6a7359b35.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/special/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*ib79c0230ea7b5afdffec69d21e9a9441ef851b72d9a0f5dd19923ae63e662025.ContentRequestBuilder) {
    return ib79c0230ea7b5afdffec69d21e9a9441ef851b72d9a0f5dd19923ae63e662025.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property special for users
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
// CreatePatchRequestInformation update the navigation property special in users
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
// Delete delete navigation property special for users
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
// Get collection of common folders available in OneDrive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*ia2e39a4c866ff367b398d30e099d94047a5cc6fe607bf52e8ab0b48d4e65c2bb.ListItemRequestBuilder) {
    return ia2e39a4c866ff367b398d30e099d94047a5cc6fe607bf52e8ab0b48d4e65c2bb.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property special in users
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
func (m *DriveItemItemRequestBuilder) Permissions()(*idd80c801d7860c1d12a236550e4ca9d092abaffbfe2250e9b953e3bd93aae3f2.PermissionsRequestBuilder) {
    return idd80c801d7860c1d12a236550e4ca9d092abaffbfe2250e9b953e3bd93aae3f2.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.special.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i39c0c9e01bc01eff72ee5b18410e29de41c051d189f03d16e7732bcc6aab5018.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i39c0c9e01bc01eff72ee5b18410e29de41c051d189f03d16e7732bcc6aab5018.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i7b005ff8448f500405ea534a1724e1d349c087dc25092f66a4296dec55221fcc.SubscriptionsRequestBuilder) {
    return i7b005ff8448f500405ea534a1724e1d349c087dc25092f66a4296dec55221fcc.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.special.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i20d350a7d9efb08a78cee6c2c9f0251dacea33b325d63d77796a12151b05934b.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i20d350a7d9efb08a78cee6c2c9f0251dacea33b325d63d77796a12151b05934b.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i82c51906a8f8e5995521d546d21b7f258d44003c77ada3e6f8a7d0a71494d7be.ThumbnailsRequestBuilder) {
    return i82c51906a8f8e5995521d546d21b7f258d44003c77ada3e6f8a7d0a71494d7be.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.special.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ib57162d336ed51ae0e881d55b1433a3969e660a3fbd45b028160ef5bc285d72d.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return ib57162d336ed51ae0e881d55b1433a3969e660a3fbd45b028160ef5bc285d72d.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i16f286c80a97a30511cfd598f3bc4eb419cd0a68fb00ce06fbfbc67c97e1114f.VersionsRequestBuilder) {
    return i16f286c80a97a30511cfd598f3bc4eb419cd0a68fb00ce06fbfbc67c97e1114f.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.special.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i51f749cb530aa4a2c132d8a1ddaa5c29f022af8b45617699c428291ed0f5018e.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i51f749cb530aa4a2c132d8a1ddaa5c29f022af8b45617699c428291ed0f5018e.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
