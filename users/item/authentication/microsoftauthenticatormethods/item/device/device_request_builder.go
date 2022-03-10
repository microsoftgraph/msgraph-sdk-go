package device

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i174c93e2bf6216db7d02c40583dbf0ec0ed58e5122e2451340e395b5f387fe47 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof"
    i20a1962c9c020754c46c9e5a69dac86cf499d2d0e21eb8a6c8ceb4ea9fe34c0e "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof"
    i6830c0080fd06abbcca1525b19e421baa6b8c3c13f2dc5026bf612f8b8247f2d "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers"
    i6fbf0693744b058e513e818aa4affaf858ef325988bc7a5f5bc1adb352fa32bf "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners"
    i7ad2b3c2829892b85e796eebc9badd30d27e526b7cfc9a5ae7fd971a8b6de17f "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/extensions"
    i1bcca13b60eab3ce212dd468eb0a794407eb579d1adec903d06906dc3d20e286 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/extensions/item"
    i2796c7ec8cdee993e4a952c556cdb269b18cc6fffe387b8d7718e8081f240971 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item"
    i3420a635cda044e9f4d541d20e4a3f491085204a4694c285f75cee30a5b05324 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item"
    i76969f911a888f11441601f995b24c61cfa766e59e6efd5ee1d4facb836fc7d4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item"
    idc2cb9ffa9ded903d24e35cc8f549e83f2b748fdbfdfdd20b4f548aa1e378b39 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item"
)

// DeviceRequestBuilder provides operations to manage the device property of the microsoft.graph.microsoftAuthenticatorAuthenticationMethod entity.
type DeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceRequestBuilderDeleteOptions options for Delete
type DeviceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceRequestBuilderGetOptions options for Get
type DeviceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceRequestBuilderGetQueryParameters the registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
type DeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceRequestBuilderPatchOptions options for Patch
type DeviceRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Deviceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod_id}/device{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceRequestBuilder instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property device for users
func (m *DeviceRequestBuilder) CreateDeleteRequestInformation(options *DeviceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *DeviceRequestBuilder) CreateGetRequestInformation(options *DeviceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property device in users
func (m *DeviceRequestBuilder) CreatePatchRequestInformation(options *DeviceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property device for users
func (m *DeviceRequestBuilder) Delete(options *DeviceRequestBuilderDeleteOptions)(error) {
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
func (m *DeviceRequestBuilder) Extensions()(*i7ad2b3c2829892b85e796eebc9badd30d27e526b7cfc9a5ae7fd971a8b6de17f.ExtensionsRequestBuilder) {
    return i7ad2b3c2829892b85e796eebc9badd30d27e526b7cfc9a5ae7fd971a8b6de17f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.microsoftAuthenticatorMethods.item.device.extensions.item collection
func (m *DeviceRequestBuilder) ExtensionsById(id string)(*i1bcca13b60eab3ce212dd468eb0a794407eb579d1adec903d06906dc3d20e286.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i1bcca13b60eab3ce212dd468eb0a794407eb579d1adec903d06906dc3d20e286.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *DeviceRequestBuilder) Get(options *DeviceRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Deviceable), nil
}
func (m *DeviceRequestBuilder) MemberOf()(*i20a1962c9c020754c46c9e5a69dac86cf499d2d0e21eb8a6c8ceb4ea9fe34c0e.MemberOfRequestBuilder) {
    return i20a1962c9c020754c46c9e5a69dac86cf499d2d0e21eb8a6c8ceb4ea9fe34c0e.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.microsoftAuthenticatorMethods.item.device.memberOf.item collection
func (m *DeviceRequestBuilder) MemberOfById(id string)(*i76969f911a888f11441601f995b24c61cfa766e59e6efd5ee1d4facb836fc7d4.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i76969f911a888f11441601f995b24c61cfa766e59e6efd5ee1d4facb836fc7d4.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property device in users
func (m *DeviceRequestBuilder) Patch(options *DeviceRequestBuilderPatchOptions)(error) {
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
func (m *DeviceRequestBuilder) RegisteredOwners()(*i6fbf0693744b058e513e818aa4affaf858ef325988bc7a5f5bc1adb352fa32bf.RegisteredOwnersRequestBuilder) {
    return i6fbf0693744b058e513e818aa4affaf858ef325988bc7a5f5bc1adb352fa32bf.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.microsoftAuthenticatorMethods.item.device.registeredOwners.item collection
func (m *DeviceRequestBuilder) RegisteredOwnersById(id string)(*i3420a635cda044e9f4d541d20e4a3f491085204a4694c285f75cee30a5b05324.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i3420a635cda044e9f4d541d20e4a3f491085204a4694c285f75cee30a5b05324.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceRequestBuilder) RegisteredUsers()(*i6830c0080fd06abbcca1525b19e421baa6b8c3c13f2dc5026bf612f8b8247f2d.RegisteredUsersRequestBuilder) {
    return i6830c0080fd06abbcca1525b19e421baa6b8c3c13f2dc5026bf612f8b8247f2d.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.microsoftAuthenticatorMethods.item.device.registeredUsers.item collection
func (m *DeviceRequestBuilder) RegisteredUsersById(id string)(*i2796c7ec8cdee993e4a952c556cdb269b18cc6fffe387b8d7718e8081f240971.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i2796c7ec8cdee993e4a952c556cdb269b18cc6fffe387b8d7718e8081f240971.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceRequestBuilder) TransitiveMemberOf()(*i174c93e2bf6216db7d02c40583dbf0ec0ed58e5122e2451340e395b5f387fe47.TransitiveMemberOfRequestBuilder) {
    return i174c93e2bf6216db7d02c40583dbf0ec0ed58e5122e2451340e395b5f387fe47.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.microsoftAuthenticatorMethods.item.device.transitiveMemberOf.item collection
func (m *DeviceRequestBuilder) TransitiveMemberOfById(id string)(*idc2cb9ffa9ded903d24e35cc8f549e83f2b748fdbfdfdd20b4f548aa1e378b39.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return idc2cb9ffa9ded903d24e35cc8f549e83f2b748fdbfdfdd20b4f548aa1e378b39.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
