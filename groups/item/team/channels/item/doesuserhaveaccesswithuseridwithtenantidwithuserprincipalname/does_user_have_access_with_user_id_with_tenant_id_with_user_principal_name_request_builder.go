package doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder provides operations to call the doesUserHaveAccess method.
type DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetQueryParameters invoke function doesUserHaveAccess
type DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetQueryParameters struct {
    // Usage: tenantId='{tenantId}'
    TenantId *string
    // Usage: userId='{userId}'
    UserId *string
    // Usage: userPrincipalName='{userPrincipalName}'
    UserPrincipalName *string
}
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetQueryParameters
}
// NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal instantiates a new DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder and sets the default values.
func NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    m := &DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/channels/{channel%2Did}/microsoft.graph.doesUserHaveAccess(userId='{userId}',tenantId='{tenantId}',userPrincipalName='{userPrincipalName}'){?userId,tenantId,userPrincipalName}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder instantiates a new DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder and sets the default values.
func NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) Get()(DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseable), nil
}
