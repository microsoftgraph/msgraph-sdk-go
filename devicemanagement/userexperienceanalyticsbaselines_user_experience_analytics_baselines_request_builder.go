package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderGetQueryParameters user experience analytics baselines
type UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsBaselineId provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) ByUserExperienceAnalyticsBaselineId(userExperienceAnalyticsBaselineId string)(*UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsBaselineId != "" {
        urlTplParams["userExperienceAnalyticsBaseline%2Did"] = userExperienceAnalyticsBaselineId
    }
    return NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderInternal instantiates a new UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) {
    m := &UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder instantiates a new UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsbaselinesCountRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) Count()(*UserexperienceanalyticsbaselinesCountRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics baselines
// returns a UserExperienceAnalyticsBaselineCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsBaselineCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsBaselines for deviceManagement
// returns a UserExperienceAnalyticsBaselineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable), nil
}
// ToGetRequestInformation user experience analytics baselines
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsBaselines for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsBaselineable, requestConfiguration *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
