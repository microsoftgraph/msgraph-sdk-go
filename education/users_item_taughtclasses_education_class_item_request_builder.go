package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemTaughtclassesEducationClassItemRequestBuilder provides operations to manage the taughtClasses property of the microsoft.graph.educationUser entity.
type UsersItemTaughtclassesEducationClassItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UsersItemTaughtclassesEducationClassItemRequestBuilderGetQueryParameters classes for which the user is a teacher.
type UsersItemTaughtclassesEducationClassItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemTaughtclassesEducationClassItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemTaughtclassesEducationClassItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemTaughtclassesEducationClassItemRequestBuilderGetQueryParameters
}
// NewUsersItemTaughtclassesEducationClassItemRequestBuilderInternal instantiates a new UsersItemTaughtclassesEducationClassItemRequestBuilder and sets the default values.
func NewUsersItemTaughtclassesEducationClassItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemTaughtclassesEducationClassItemRequestBuilder) {
    m := &UsersItemTaughtclassesEducationClassItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/users/{educationUser%2Did}/taughtClasses/{educationClass%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUsersItemTaughtclassesEducationClassItemRequestBuilder instantiates a new UsersItemTaughtclassesEducationClassItemRequestBuilder and sets the default values.
func NewUsersItemTaughtclassesEducationClassItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemTaughtclassesEducationClassItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemTaughtclassesEducationClassItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get classes for which the user is a teacher.
// returns a EducationClassable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsersItemTaughtclassesEducationClassItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemTaughtclassesEducationClassItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationClassable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationClassFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationClassable), nil
}
// ToGetRequestInformation classes for which the user is a teacher.
// returns a *RequestInformation when successful
func (m *UsersItemTaughtclassesEducationClassItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemTaughtclassesEducationClassItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UsersItemTaughtclassesEducationClassItemRequestBuilder when successful
func (m *UsersItemTaughtclassesEducationClassItemRequestBuilder) WithUrl(rawUrl string)(*UsersItemTaughtclassesEducationClassItemRequestBuilder) {
    return NewUsersItemTaughtclassesEducationClassItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
