package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder provides operations to manage the gradingCategory property of the microsoft.graph.educationAssignment entity.
type ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetQueryParameters when set, enables users to weight assignments differently when computing a class average grade.
type ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetQueryParameters
}
// NewClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderInternal instantiates a new ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) {
    m := &ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/gradingCategory{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder instantiates a new ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get when set, enables users to weight assignments differently when computing a class average grade.
// returns a EducationGradingCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationGradingCategoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
func (m *ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder when successful
func (m *ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) WithUrl(rawUrl string)(*ClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) {
    return NewClassesItemAssignmentsItemGradingcategoryGradingCategoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
