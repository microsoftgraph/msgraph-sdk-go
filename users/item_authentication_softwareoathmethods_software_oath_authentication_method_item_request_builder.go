package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderGetQueryParameters the software OATH time-based one-time password (TOTP) applications registered to a user for authentication.
type ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderGetQueryParameters
}
// NewItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderInternal instantiates a new ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder) {
    m := &ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/softwareOathMethods/{softwareOathAuthenticationMethod%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder instantiates a new ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a user's Software OATH token authentication method object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/softwareoathauthenticationmethod-delete?view=graph-rest-1.0
func (m *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the software OATH time-based one-time password (TOTP) applications registered to a user for authentication.
// returns a SoftwareOathAuthenticationMethodable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SoftwareOathAuthenticationMethodable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSoftwareOathAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SoftwareOathAuthenticationMethodable), nil
}
// ToDeleteRequestInformation delete a user's Software OATH token authentication method object.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the software OATH time-based one-time password (TOTP) applications registered to a user for authentication.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder when successful
func (m *ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder) {
    return NewItemAuthenticationSoftwareoathmethodsSoftwareOathAuthenticationMethodItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
