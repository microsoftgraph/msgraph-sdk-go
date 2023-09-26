package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder casts the previous resource to microsoftStoreForBusinessApp.
type MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.microsoftStoreForBusinessApp
type MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) Assignments()(*MobileAppsItemGraphMicrosoftStoreForBusinessAppAssignmentsRequestBuilder) {
    return NewMobileAppsItemGraphMicrosoftStoreForBusinessAppAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) Categories()(*MobileAppsItemGraphMicrosoftStoreForBusinessAppCategoriesRequestBuilder) {
    return NewMobileAppsItemGraphMicrosoftStoreForBusinessAppCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderInternal instantiates a new GraphMicrosoftStoreForBusinessAppRequestBuilder and sets the default values.
func NewMobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    m := &MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.microsoftStoreForBusinessApp{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder instantiates a new GraphMicrosoftStoreForBusinessAppRequestBuilder and sets the default values.
func NewMobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.microsoftStoreForBusinessApp
func (m *MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MicrosoftStoreForBusinessAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMicrosoftStoreForBusinessAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MicrosoftStoreForBusinessAppable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.microsoftStoreForBusinessApp
func (m *MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    return NewMobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
