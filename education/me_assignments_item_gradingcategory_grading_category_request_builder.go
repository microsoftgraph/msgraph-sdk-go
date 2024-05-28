package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder provides operations to manage the gradingCategory property of the microsoft.graph.educationAssignment entity.
type MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetQueryParameters when set, enables users to weight assignments differently when computing a class average grade.
type MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetQueryParameters
}
// NewMeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderInternal instantiates a new MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder and sets the default values.
func NewMeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) {
    m := &MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/gradingCategory{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder instantiates a new MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder and sets the default values.
func NewMeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get when set, enables users to weight assignments differently when computing a class average grade.
// returns a EducationGradingCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationGradingCategoryable, error) {
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
func (m *MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder when successful
func (m *MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) WithUrl(rawUrl string)(*MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) {
    return NewMeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
