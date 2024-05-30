package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder provides operations to call the getSharePointActivityUserDetail method.
type GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilderInternal instantiates a new GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder and sets the default values.
func NewGetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder) {
    m := &GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSharePointActivityUserDetail(date={date})", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder instantiates a new GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder and sets the default values.
func NewGetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get details about SharePoint activity by user.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getsharepointactivityuserdetail?view=graph-rest-1.0
func (m *GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get details about SharePoint activity by user.
// returns a *RequestInformation when successful
func (m *GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder when successful
func (m *GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*GetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder) {
    return NewGetsharepointactivityuserdetailwithdateGetSharePointActivityUserDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
