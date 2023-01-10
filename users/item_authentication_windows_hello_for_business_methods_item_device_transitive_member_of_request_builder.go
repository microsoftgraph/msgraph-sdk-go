package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderGetQueryParameters groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) Application()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfApplicationRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderInternal instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) {
    m := &ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/transitiveMemberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) Count()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfCountRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Device casts the previous resource to device.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) Device()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDeviceRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/device-list-transitivememberof?view=graph-rest-1.0
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// Group casts the previous resource to group.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) Group()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfGroupRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) OrgContact()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) ServicePrincipal()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfServicePrincipalRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToGetRequestInformation groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// User casts the previous resource to user.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfRequestBuilder) User()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfUserRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
