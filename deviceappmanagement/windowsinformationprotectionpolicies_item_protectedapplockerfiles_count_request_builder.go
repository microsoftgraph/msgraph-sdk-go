package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder provides operations to count the resources in the collection.
type WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderGetQueryParameters get the number of the resource
type WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderGetQueryParameters
}
// NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderInternal instantiates a new WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder) {
    m := &WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy%2Did}/protectedAppLockerFiles/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder instantiates a new WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
