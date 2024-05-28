package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder provides operations to manage the gradingCategories property of the microsoft.graph.educationAssignmentSettings entity.
type ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderGetQueryParameters when set, enables users to weight assignments differently when computing a class average grade.
type ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderGetQueryParameters struct {
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
// ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderGetQueryParameters
}
// ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEducationGradingCategoryId provides operations to manage the gradingCategories property of the microsoft.graph.educationAssignmentSettings entity.
// returns a *ClassesItemAssignmentsettingsGradingcategoriesEducationGradingCategoryItemRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) ByEducationGradingCategoryId(educationGradingCategoryId string)(*ClassesItemAssignmentsettingsGradingcategoriesEducationGradingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if educationGradingCategoryId != "" {
        urlTplParams["educationGradingCategory%2Did"] = educationGradingCategoryId
    }
    return NewClassesItemAssignmentsettingsGradingcategoriesEducationGradingCategoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderInternal instantiates a new ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) {
    m := &ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignmentSettings/gradingCategories{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder instantiates a new ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ClassesItemAssignmentsettingsGradingcategoriesCountRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) Count()(*ClassesItemAssignmentsettingsGradingcategoriesCountRequestBuilder) {
    return NewClassesItemAssignmentsettingsGradingcategoriesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get when set, enables users to weight assignments differently when computing a class average grade.
// returns a EducationGradingCategoryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationGradingCategoryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationGradingCategoryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationGradingCategoryCollectionResponseable), nil
}
// Post create new navigation property to gradingCategories for education
// returns a EducationGradingCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationGradingCategoryable, requestConfiguration *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationGradingCategoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationGradingCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationGradingCategoryable), nil
}
// ToGetRequestInformation when set, enables users to weight assignments differently when computing a class average grade.
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to gradingCategories for education
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationGradingCategoryable, requestConfiguration *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) WithUrl(rawUrl string)(*ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) {
    return NewClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
