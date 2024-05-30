package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder provides operations to call the getTeamsLicensingDetails method.
type ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilderInternal instantiates a new ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder and sets the default values.
func NewItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder) {
    m := &ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/licenseDetails/getTeamsLicensingDetails()", pathParameters),
    }
    return m
}
// NewItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder instantiates a new ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder and sets the default values.
func NewItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the license status of a user in Microsoft Teams.
// returns a TeamsLicensingDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/licensedetails-getteamslicensingdetails?view=graph-rest-1.0
func (m *ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsLicensingDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamsLicensingDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsLicensingDetailsable), nil
}
// ToGetRequestInformation get the license status of a user in Microsoft Teams.
// returns a *RequestInformation when successful
func (m *ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder when successful
func (m *ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder) WithUrl(rawUrl string)(*ItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder) {
    return NewItemLicensedetailsGetteamslicensingdetailsGetTeamsLicensingDetailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
