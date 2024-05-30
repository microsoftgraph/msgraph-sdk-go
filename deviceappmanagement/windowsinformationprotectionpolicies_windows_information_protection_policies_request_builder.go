package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder provides operations to manage the windowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
type WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderGetQueryParameters list properties and relationships of the windowsInformationProtectionPolicy objects.
type WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderGetQueryParameters struct {
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
// WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderGetQueryParameters
}
// WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsInformationProtectionPolicyId provides operations to manage the windowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) ByWindowsInformationProtectionPolicyId(windowsInformationProtectionPolicyId string)(*WindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsInformationProtectionPolicyId != "" {
        urlTplParams["windowsInformationProtectionPolicy%2Did"] = windowsInformationProtectionPolicyId
    }
    return NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderInternal instantiates a new WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) {
    m := &WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder instantiates a new WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsinformationprotectionpoliciesCountRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) Count()(*WindowsinformationprotectionpoliciesCountRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the windowsInformationProtectionPolicy objects.
// returns a WindowsInformationProtectionPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-windowsinformationprotectionpolicy-list?view=graph-rest-1.0
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyCollectionResponseable), nil
}
// Post create a new windowsInformationProtectionPolicy object.
// returns a WindowsInformationProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-windowsinformationprotectionpolicy-create?view=graph-rest-1.0
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable), nil
}
// ToGetRequestInformation list properties and relationships of the windowsInformationProtectionPolicy objects.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a new windowsInformationProtectionPolicy object.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionPolicyable, requestConfiguration *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
