package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder provides operations to call the getOffice365ActivationsUserDetail method.
type Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilderInternal instantiates a new Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder and sets the default values.
func NewGetoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder) {
    m := &Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365ActivationsUserDetail()", pathParameters),
    }
    return m
}
// NewGetoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder instantiates a new Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder and sets the default values.
func NewGetoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get details about users who have activated Microsoft 365.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getoffice365activationsuserdetail?view=graph-rest-1.0
func (m *Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get details about users who have activated Microsoft 365.
// returns a *RequestInformation when successful
func (m *Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder when successful
func (m *Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder) WithUrl(rawUrl string)(*Getoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder) {
    return NewGetoffice365activationsuserdetailGetOffice365ActivationsUserDetailRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
