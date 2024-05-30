package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderGetQueryParameters user experience analytics model scores
type UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsModelScoresId provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresItemRequestBuilder when successful
func (m *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) ByUserExperienceAnalyticsModelScoresId(userExperienceAnalyticsModelScoresId string)(*UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsModelScoresId != "" {
        urlTplParams["userExperienceAnalyticsModelScores%2Did"] = userExperienceAnalyticsModelScoresId
    }
    return NewUserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderInternal instantiates a new UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder and sets the default values.
func NewUserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) {
    m := &UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsModelScores{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder instantiates a new UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder and sets the default values.
func NewUserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsmodelscoresCountRequestBuilder when successful
func (m *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) Count()(*UserexperienceanalyticsmodelscoresCountRequestBuilder) {
    return NewUserexperienceanalyticsmodelscoresCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics model scores
// returns a UserExperienceAnalyticsModelScoresCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsModelScoresCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsModelScoresCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsModelScoresCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsModelScores for deviceManagement
// returns a UserExperienceAnalyticsModelScoresable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsModelScoresable, requestConfiguration *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsModelScoresable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsModelScoresable), nil
}
// ToGetRequestInformation user experience analytics model scores
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsModelScores for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsModelScoresable, requestConfiguration *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder when successful
func (m *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) {
    return NewUserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
