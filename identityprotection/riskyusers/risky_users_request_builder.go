package riskyusers

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i5a8d5869e42d3c9c54ea08a341c5297a0b0865928dbe0569fabaee2c8e65008e "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers/dismiss"
    ic0a458e5f0f3b529bf41d03f94eacd03ae5394d2801818addefd2ef0ad4c34bc "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers/confirmcompromised"
    icf8431a0bea7943ea4c27e26967395f2983c85c27c9bf1afbd36da7444ddb779 "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers/count"
)

// RiskyUsersRequestBuilder provides operations to manage the riskyUsers property of the microsoft.graph.identityProtectionRoot entity.
type RiskyUsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RiskyUsersRequestBuilderGetQueryParameters users that are flagged as at-risk by Azure AD Identity Protection.
type RiskyUsersRequestBuilderGetQueryParameters struct {
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
// RiskyUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RiskyUsersRequestBuilderGetQueryParameters
}
// RiskyUsersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyUsersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConfirmCompromised the confirmCompromised property
func (m *RiskyUsersRequestBuilder) ConfirmCompromised()(*ic0a458e5f0f3b529bf41d03f94eacd03ae5394d2801818addefd2ef0ad4c34bc.ConfirmCompromisedRequestBuilder) {
    return ic0a458e5f0f3b529bf41d03f94eacd03ae5394d2801818addefd2ef0ad4c34bc.NewConfirmCompromisedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRiskyUsersRequestBuilderInternal instantiates a new RiskyUsersRequestBuilder and sets the default values.
func NewRiskyUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyUsersRequestBuilder) {
    m := &RiskyUsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyUsers{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRiskyUsersRequestBuilder instantiates a new RiskyUsersRequestBuilder and sets the default values.
func NewRiskyUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *RiskyUsersRequestBuilder) Count()(*icf8431a0bea7943ea4c27e26967395f2983c85c27c9bf1afbd36da7444ddb779.CountRequestBuilder) {
    return icf8431a0bea7943ea4c27e26967395f2983c85c27c9bf1afbd36da7444ddb779.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUsersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUsersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RiskyUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to riskyUsers for identityProtection
func (m *RiskyUsersRequestBuilder) CreatePostRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to riskyUsers for identityProtection
func (m *RiskyUsersRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserable, requestConfiguration *RiskyUsersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Dismiss the dismiss property
func (m *RiskyUsersRequestBuilder) Dismiss()(*i5a8d5869e42d3c9c54ea08a341c5297a0b0865928dbe0569fabaee2c8e65008e.DismissRequestBuilder) {
    return i5a8d5869e42d3c9c54ea08a341c5297a0b0865928dbe0569fabaee2c8e65008e.NewDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUsersRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUsersRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RiskyUsersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRiskyUserCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserCollectionResponseable), nil
}
// Post create new navigation property to riskyUsers for identityProtection
func (m *RiskyUsersRequestBuilder) Post(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserable)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to riskyUsers for identityProtection
func (m *RiskyUsersRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserable, requestConfiguration *RiskyUsersRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRiskyUserFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyUserable), nil
}
