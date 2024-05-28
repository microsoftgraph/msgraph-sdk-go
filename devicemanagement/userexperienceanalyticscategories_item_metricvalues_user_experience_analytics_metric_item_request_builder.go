package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder provides operations to manage the metricValues property of the microsoft.graph.userExperienceAnalyticsCategory entity.
type UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderGetQueryParameters the metric values for the user experience analytics category. Read-only.
type UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderInternal instantiates a new UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) {
    m := &UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsCategories/{userExperienceAnalyticsCategory%2Did}/metricValues/{userExperienceAnalyticsMetric%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder instantiates a new UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property metricValues for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the metric values for the user experience analytics category. Read-only.
// returns a UserExperienceAnalyticsMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricable), nil
}
// Patch update the navigation property metricValues in deviceManagement
// returns a UserExperienceAnalyticsMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricable, requestConfiguration *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricable), nil
}
// ToDeleteRequestInformation delete navigation property metricValues for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the metric values for the user experience analytics category. Read-only.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property metricValues in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsMetricable, requestConfiguration *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder when successful
func (m *UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder) {
    return NewUserexperienceanalyticscategoriesItemMetricvaluesUserExperienceAnalyticsMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
