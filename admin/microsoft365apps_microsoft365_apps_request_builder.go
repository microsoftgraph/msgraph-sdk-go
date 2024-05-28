package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Microsoft365appsMicrosoft365AppsRequestBuilder provides operations to manage the microsoft365Apps property of the microsoft.graph.admin entity.
type Microsoft365appsMicrosoft365AppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Microsoft365appsMicrosoft365AppsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Microsoft365appsMicrosoft365AppsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Microsoft365appsMicrosoft365AppsRequestBuilderGetQueryParameters a container for the Microsoft 365 apps admin functionality.
type Microsoft365appsMicrosoft365AppsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// Microsoft365appsMicrosoft365AppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Microsoft365appsMicrosoft365AppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Microsoft365appsMicrosoft365AppsRequestBuilderGetQueryParameters
}
// Microsoft365appsMicrosoft365AppsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Microsoft365appsMicrosoft365AppsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoft365appsMicrosoft365AppsRequestBuilderInternal instantiates a new Microsoft365appsMicrosoft365AppsRequestBuilder and sets the default values.
func NewMicrosoft365appsMicrosoft365AppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Microsoft365appsMicrosoft365AppsRequestBuilder) {
    m := &Microsoft365appsMicrosoft365AppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/microsoft365Apps{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMicrosoft365appsMicrosoft365AppsRequestBuilder instantiates a new Microsoft365appsMicrosoft365AppsRequestBuilder and sets the default values.
func NewMicrosoft365appsMicrosoft365AppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Microsoft365appsMicrosoft365AppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoft365appsMicrosoft365AppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property microsoft365Apps for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Microsoft365appsMicrosoft365AppsRequestBuilder) Delete(ctx context.Context, requestConfiguration *Microsoft365appsMicrosoft365AppsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a container for the Microsoft 365 apps admin functionality.
// returns a AdminMicrosoft365Appsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Microsoft365appsMicrosoft365AppsRequestBuilder) Get(ctx context.Context, requestConfiguration *Microsoft365appsMicrosoft365AppsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdminMicrosoft365Appsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAdminMicrosoft365AppsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdminMicrosoft365Appsable), nil
}
// InstallationOptions provides operations to manage the installationOptions property of the microsoft.graph.adminMicrosoft365Apps entity.
// returns a *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder when successful
func (m *Microsoft365appsMicrosoft365AppsRequestBuilder) InstallationOptions()(*Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) {
    return NewMicrosoft365appsInstallationoptionsInstallationOptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property microsoft365Apps in admin
// returns a AdminMicrosoft365Appsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Microsoft365appsMicrosoft365AppsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdminMicrosoft365Appsable, requestConfiguration *Microsoft365appsMicrosoft365AppsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdminMicrosoft365Appsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAdminMicrosoft365AppsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdminMicrosoft365Appsable), nil
}
// ToDeleteRequestInformation delete navigation property microsoft365Apps for admin
// returns a *RequestInformation when successful
func (m *Microsoft365appsMicrosoft365AppsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *Microsoft365appsMicrosoft365AppsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a container for the Microsoft 365 apps admin functionality.
// returns a *RequestInformation when successful
func (m *Microsoft365appsMicrosoft365AppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Microsoft365appsMicrosoft365AppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property microsoft365Apps in admin
// returns a *RequestInformation when successful
func (m *Microsoft365appsMicrosoft365AppsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdminMicrosoft365Appsable, requestConfiguration *Microsoft365appsMicrosoft365AppsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *Microsoft365appsMicrosoft365AppsRequestBuilder when successful
func (m *Microsoft365appsMicrosoft365AppsRequestBuilder) WithUrl(rawUrl string)(*Microsoft365appsMicrosoft365AppsRequestBuilder) {
    return NewMicrosoft365appsMicrosoft365AppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
