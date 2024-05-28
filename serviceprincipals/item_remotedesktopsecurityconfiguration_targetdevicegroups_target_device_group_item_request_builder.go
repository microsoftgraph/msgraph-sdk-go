package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder provides operations to manage the targetDeviceGroups property of the microsoft.graph.remoteDesktopSecurityConfiguration entity.
type ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderGetQueryParameters read the properties and relationships of a targetDeviceGroup object for the remoteDesktopSecurityConfiguration object on the servicePrincipal.
type ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderGetQueryParameters
}
// ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderInternal instantiates a new ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder and sets the default values.
func NewItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) {
    m := &ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/remoteDesktopSecurityConfiguration/targetDeviceGroups/{targetDeviceGroup%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder instantiates a new ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder and sets the default values.
func NewItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a targetDeviceGroup object for the remoteDesktopSecurityConfiguration object on the servicePrincipal. Any user authenticating using the Microsoft Entra ID Remote Desktop Services (RDS) authentication protocol to a Microsoft Entra joined or Microsoft Entra hybrid joined device that's in the removed targetDeviceGroup doesn't get SSO prompts.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/remotedesktopsecurityconfiguration-delete-targetdevicegroups?view=graph-rest-1.0
func (m *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of a targetDeviceGroup object for the remoteDesktopSecurityConfiguration object on the servicePrincipal.
// returns a TargetDeviceGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/targetdevicegroup-get?view=graph-rest-1.0
func (m *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetDeviceGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTargetDeviceGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetDeviceGroupable), nil
}
// Patch update the properties of a targetDeviceGroup object for remoteDesktopSecurityConfiguration object on the servicePrincipal. You can configure a maximum of 10 target device groups for the remoteDesktopSecurityConfiguraiton object on the servicePrincipal.
// returns a TargetDeviceGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/targetdevicegroup-update?view=graph-rest-1.0
func (m *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetDeviceGroupable, requestConfiguration *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetDeviceGroupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTargetDeviceGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetDeviceGroupable), nil
}
// ToDeleteRequestInformation delete a targetDeviceGroup object for the remoteDesktopSecurityConfiguration object on the servicePrincipal. Any user authenticating using the Microsoft Entra ID Remote Desktop Services (RDS) authentication protocol to a Microsoft Entra joined or Microsoft Entra hybrid joined device that's in the removed targetDeviceGroup doesn't get SSO prompts.
// returns a *RequestInformation when successful
func (m *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a targetDeviceGroup object for the remoteDesktopSecurityConfiguration object on the servicePrincipal.
// returns a *RequestInformation when successful
func (m *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a targetDeviceGroup object for remoteDesktopSecurityConfiguration object on the servicePrincipal. You can configure a maximum of 10 target device groups for the remoteDesktopSecurityConfiguraiton object on the servicePrincipal.
// returns a *RequestInformation when successful
func (m *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetDeviceGroupable, requestConfiguration *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder when successful
func (m *ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) WithUrl(rawUrl string)(*ItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder) {
    return NewItemRemotedesktopsecurityconfigurationTargetdevicegroupsTargetDeviceGroupItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
