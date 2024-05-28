package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder provides operations to call the getM365AppUserDetail method.
type Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilderInternal instantiates a new Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder and sets the default values.
func NewGetm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder) {
    m := &Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getM365AppUserDetail(date={date})", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder instantiates a new Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder and sets the default values.
func NewGetm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get a report that provides the details about which apps and platforms users have used.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get a report that provides the details about which apps and platforms users have used.
// returns a *RequestInformation when successful
func (m *Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder when successful
func (m *Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*Getm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder) {
    return NewGetm365appuserdetailwithdateGetM365AppUserDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
