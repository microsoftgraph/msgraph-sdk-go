package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder provides operations to manage the device property of the microsoft.graph.windowsHelloForBusinessAuthenticationMethod entity.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderGetQueryParameters the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderGetQueryParameters
}
// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) CheckMemberGroups()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceCheckMemberGroupsRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) CheckMemberObjects()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceCheckMemberObjectsRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) {
    m := &ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder instantiates a new DeviceRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property device for users
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.device entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) Extensions()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceExtensionsRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.device entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) ExtensionsById(id string)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) GetMemberGroups()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberGroupsRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) GetMemberObjects()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) MemberOf()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) MemberOfById(id string)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property device in users
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable), nil
}
// RegisteredOwners provides operations to manage the registeredOwners property of the microsoft.graph.device entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) RegisteredOwners()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item.device.registeredOwners.item collection
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) RegisteredOwnersById(id string)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) RegisteredUsers()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) RegisteredUsersById(id string)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) Restore()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRestoreRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToDeleteRequestInformation delete navigation property device for users
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property device in users
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) TransitiveMemberOf()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) TransitiveMemberOfById(id string)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
