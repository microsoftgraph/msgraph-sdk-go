package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Application()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemApplicationRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) {
    m := &ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device casts the previous resource to device.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Device()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemDeviceRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// Group casts the previous resource to group.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Group()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemGroupRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) OrgContact()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemOrgContactRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) ServicePrincipal()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemServicePrincipalRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) User()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemUserRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}