package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderGetQueryParameters user experience analytics metric history
type UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsMetricHistoryId provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilder when successful
func (m *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) ByUserExperienceAnalyticsMetricHistoryId(userExperienceAnalyticsMetricHistoryId string)(*UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsMetricHistoryId != "" {
        urlTplParams["userExperienceAnalyticsMetricHistory%2Did"] = userExperienceAnalyticsMetricHistoryId
    }
    return NewUserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderInternal instantiates a new UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder and sets the default values.
func NewUserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) {
    m := &UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsMetricHistory{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder instantiates a new UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder and sets the default values.
func NewUserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsmetrichistoryCountRequestBuilder when successful
func (m *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) Count()(*UserexperienceanalyticsmetrichistoryCountRequestBuilder) {
    return NewUserexperienceanalyticsmetrichistoryCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics metric history
// returns a UserExperienceAnalyticsMetricHistoryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricHistoryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsMetricHistoryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricHistoryCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsMetricHistory for deviceManagement
// returns a UserExperienceAnalyticsMetricHistoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricHistoryable, requestConfiguration *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricHistoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricHistoryable), nil
}
// ToGetRequestInformation user experience analytics metric history
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsMetricHistory for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricHistoryable, requestConfiguration *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder when successful
func (m *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) {
    return NewUserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
