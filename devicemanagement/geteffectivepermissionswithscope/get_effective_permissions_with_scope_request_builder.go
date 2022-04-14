package geteffectivepermissionswithscope

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetEffectivePermissionsWithScopeRequestBuilder provides operations to call the getEffectivePermissions method.
type GetEffectivePermissionsWithScopeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetEffectivePermissionsWithScopeRequestBuilderGetOptions options for Get
type GetEffectivePermissionsWithScopeRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetEffectivePermissionsWithScopeRequestBuilderInternal instantiates a new GetEffectivePermissionsWithScopeRequestBuilder and sets the default values.
func NewGetEffectivePermissionsWithScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, scope *string)(*GetEffectivePermissionsWithScopeRequestBuilder) {
    m := &GetEffectivePermissionsWithScopeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getEffectivePermissions(scope='{scope}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if scope != nil {
        urlTplParams[""] = *scope
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetEffectivePermissionsWithScopeRequestBuilder instantiates a new GetEffectivePermissionsWithScopeRequestBuilder and sets the default values.
func NewGetEffectivePermissionsWithScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetEffectivePermissionsWithScopeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetEffectivePermissionsWithScopeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation retrieves the effective permissions of the currently authenticated user
func (m *GetEffectivePermissionsWithScopeRequestBuilder) CreateGetRequestInformation(options *GetEffectivePermissionsWithScopeRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
// Get retrieves the effective permissions of the currently authenticated user
func (m *GetEffectivePermissionsWithScopeRequestBuilder) Get(options *GetEffectivePermissionsWithScopeRequestBuilderGetOptions)(GetEffectivePermissionsWithScopeResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetEffectivePermissionsWithScopeResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetEffectivePermissionsWithScopeResponseable), nil
}
