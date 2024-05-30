package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetQueryParameters retrieve a user's single temporaryAccessPassAuthenticationMethod object.
type ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetQueryParameters
}
// NewItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal instantiates a new ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    m := &ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/temporaryAccessPassMethods/{temporaryAccessPassAuthenticationMethod%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder instantiates a new ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a users's temporaryAccessPassAuthenticationMethod object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/temporaryaccesspassauthenticationmethod-delete?view=graph-rest-1.0
func (m *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve a user's single temporaryAccessPassAuthenticationMethod object.
// returns a TemporaryAccessPassAuthenticationMethodable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/temporaryaccesspassauthenticationmethod-get?view=graph-rest-1.0
func (m *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TemporaryAccessPassAuthenticationMethodable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTemporaryAccessPassAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TemporaryAccessPassAuthenticationMethodable), nil
}
// ToDeleteRequestInformation delete a users's temporaryAccessPassAuthenticationMethod object.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a user's single temporaryAccessPassAuthenticationMethod object.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder when successful
func (m *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    return NewItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
