package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder provides operations to call the getSkypeForBusinessActivityUserDetail method.
type GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal instantiates a new GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder and sets the default values.
func NewGetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    m := &GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSkypeForBusinessActivityUserDetail(date={date})", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder instantiates a new GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder and sets the default values.
func NewGetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get details about Skype for Business activity by user.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getskypeforbusinessactivityuserdetail?view=graph-rest-1.0
func (m *GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get details about Skype for Business activity by user.
// returns a *RequestInformation when successful
func (m *GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder when successful
func (m *GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*GetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return NewGetskypeforbusinessactivityuserdetailwithdateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
