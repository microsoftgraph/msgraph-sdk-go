package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder provides operations to call the setMobileDeviceManagementAuthority method.
type ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilderInternal instantiates a new ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder) {
    m := &ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/setMobileDeviceManagementAuthority", pathParameters),
    }
    return m
}
// NewItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder instantiates a new ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set mobile device management authority
// Deprecated: This method is obsolete. Use PostAsSetMobileDeviceManagementAuthorityPostResponse instead.
// returns a ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-organization-setmobiledevicemanagementauthority?view=graph-rest-1.0
func (m *ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration)(ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityResponseable), nil
}
// PostAsSetMobileDeviceManagementAuthorityPostResponse set mobile device management authority
// returns a ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-organization-setmobiledevicemanagementauthority?view=graph-rest-1.0
func (m *ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder) PostAsSetMobileDeviceManagementAuthorityPostResponse(ctx context.Context, requestConfiguration *ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration)(ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityPostResponseable), nil
}
// ToPostRequestInformation set mobile device management authority
// returns a *RequestInformation when successful
func (m *ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder when successful
func (m *ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder) WithUrl(rawUrl string)(*ItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder) {
    return NewItemSetmobiledevicemanagementauthoritySetMobileDeviceManagementAuthorityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
