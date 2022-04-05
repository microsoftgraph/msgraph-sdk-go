package device

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i02f11ba276c0654b01548d69faebb0c77900168eb0641b68226993bf4e352627 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners"
    i2d3ed20798c35789c760a72a8a3cff923a717203a8f55c8703d063849bae61c9 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof"
    i8863c21e280ef32e5445c2b728af5fbcbb71c005cbf858df118b8ed1f88f53a5 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof"
    id777e348b88248f5da3d11b2a2524ae92ecf373d040697797a24f5fafc73ccec "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers"
    if3271ea0beb260c0136b8041947b3fabf07eb5c61529b72bf2165e70ddd5ee8a "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/extensions"
    i124bc1a2abc2ae5d3383f13c8ca5eba415dcb0a9702d7930770b07e06244958b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item"
    i2a7d0d988b8e6e11414ca2b5a2d0d798269995d56f7279e62a56a57578a2bed6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item"
    i5d95cd71373e2e78ad60d009bead769a40e126e26f3eba8633c5e4fd90ac2d9c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item"
    i66f2c54f642cba3e3d4efe9f89aec470cac5dae8b99ce568d1772298e963e7ea "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/extensions/item"
    i7e1f64d06d2f0bfeaba45f3c32002ab7002ee51d3db27487b11629e94aaed07b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item"
)

// DeviceRequestBuilder provides operations to manage the device property of the microsoft.graph.windowsHelloForBusinessAuthenticationMethod entity.
type DeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceRequestBuilderDeleteOptions options for Delete
type DeviceRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceRequestBuilderGetOptions options for Get
type DeviceRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceRequestBuilderGetQueryParameters the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
type DeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceRequestBuilderPatchOptions options for Patch
type DeviceRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod_id}/device{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceRequestBuilder instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property device for me
func (m *DeviceRequestBuilder) CreateDeleteRequestInformation(options *DeviceRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *DeviceRequestBuilder) CreateGetRequestInformation(options *DeviceRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property device in me
func (m *DeviceRequestBuilder) CreatePatchRequestInformation(options *DeviceRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property device for me
func (m *DeviceRequestBuilder) Delete(options *DeviceRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions the extensions property
func (m *DeviceRequestBuilder) Extensions()(*if3271ea0beb260c0136b8041947b3fabf07eb5c61529b72bf2165e70ddd5ee8a.ExtensionsRequestBuilder) {
    return if3271ea0beb260c0136b8041947b3fabf07eb5c61529b72bf2165e70ddd5ee8a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.windowsHelloForBusinessMethods.item.device.extensions.item collection
func (m *DeviceRequestBuilder) ExtensionsById(id string)(*i66f2c54f642cba3e3d4efe9f89aec470cac5dae8b99ce568d1772298e963e7ea.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i66f2c54f642cba3e3d4efe9f89aec470cac5dae8b99ce568d1772298e963e7ea.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *DeviceRequestBuilder) Get(options *DeviceRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable), nil
}
// MemberOf the memberOf property
func (m *DeviceRequestBuilder) MemberOf()(*i8863c21e280ef32e5445c2b728af5fbcbb71c005cbf858df118b8ed1f88f53a5.MemberOfRequestBuilder) {
    return i8863c21e280ef32e5445c2b728af5fbcbb71c005cbf858df118b8ed1f88f53a5.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.windowsHelloForBusinessMethods.item.device.memberOf.item collection
func (m *DeviceRequestBuilder) MemberOfById(id string)(*i124bc1a2abc2ae5d3383f13c8ca5eba415dcb0a9702d7930770b07e06244958b.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i124bc1a2abc2ae5d3383f13c8ca5eba415dcb0a9702d7930770b07e06244958b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property device in me
func (m *DeviceRequestBuilder) Patch(options *DeviceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RegisteredOwners the registeredOwners property
func (m *DeviceRequestBuilder) RegisteredOwners()(*i02f11ba276c0654b01548d69faebb0c77900168eb0641b68226993bf4e352627.RegisteredOwnersRequestBuilder) {
    return i02f11ba276c0654b01548d69faebb0c77900168eb0641b68226993bf4e352627.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.windowsHelloForBusinessMethods.item.device.registeredOwners.item collection
func (m *DeviceRequestBuilder) RegisteredOwnersById(id string)(*i2a7d0d988b8e6e11414ca2b5a2d0d798269995d56f7279e62a56a57578a2bed6.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i2a7d0d988b8e6e11414ca2b5a2d0d798269995d56f7279e62a56a57578a2bed6.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers the registeredUsers property
func (m *DeviceRequestBuilder) RegisteredUsers()(*id777e348b88248f5da3d11b2a2524ae92ecf373d040697797a24f5fafc73ccec.RegisteredUsersRequestBuilder) {
    return id777e348b88248f5da3d11b2a2524ae92ecf373d040697797a24f5fafc73ccec.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.windowsHelloForBusinessMethods.item.device.registeredUsers.item collection
func (m *DeviceRequestBuilder) RegisteredUsersById(id string)(*i5d95cd71373e2e78ad60d009bead769a40e126e26f3eba8633c5e4fd90ac2d9c.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i5d95cd71373e2e78ad60d009bead769a40e126e26f3eba8633c5e4fd90ac2d9c.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *DeviceRequestBuilder) TransitiveMemberOf()(*i2d3ed20798c35789c760a72a8a3cff923a717203a8f55c8703d063849bae61c9.TransitiveMemberOfRequestBuilder) {
    return i2d3ed20798c35789c760a72a8a3cff923a717203a8f55c8703d063849bae61c9.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.windowsHelloForBusinessMethods.item.device.transitiveMemberOf.item collection
func (m *DeviceRequestBuilder) TransitiveMemberOfById(id string)(*i7e1f64d06d2f0bfeaba45f3c32002ab7002ee51d3db27487b11629e94aaed07b.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i7e1f64d06d2f0bfeaba45f3c32002ab7002ee51d3db27487b11629e94aaed07b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
