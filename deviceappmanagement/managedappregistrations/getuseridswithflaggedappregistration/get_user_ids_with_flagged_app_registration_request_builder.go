package getuseridswithflaggedappregistration

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetUserIdsWithFlaggedAppRegistrationRequestBuilder provides operations to call the getUserIdsWithFlaggedAppRegistration method.
type GetUserIdsWithFlaggedAppRegistrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal instantiates a new GetUserIdsWithFlaggedAppRegistrationRequestBuilder and sets the default values.
func NewGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    m := &GetUserIdsWithFlaggedAppRegistrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/microsoft.graph.getUserIdsWithFlaggedAppRegistration()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetUserIdsWithFlaggedAppRegistrationRequestBuilder instantiates a new GetUserIdsWithFlaggedAppRegistrationRequestBuilder and sets the default values.
func NewGetUserIdsWithFlaggedAppRegistrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getUserIdsWithFlaggedAppRegistration
func (m *GetUserIdsWithFlaggedAppRegistrationRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getUserIdsWithFlaggedAppRegistration
func (m *GetUserIdsWithFlaggedAppRegistrationRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getUserIdsWithFlaggedAppRegistration
func (m *GetUserIdsWithFlaggedAppRegistrationRequestBuilder) Get()(GetUserIdsWithFlaggedAppRegistrationResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getUserIdsWithFlaggedAppRegistration
func (m *GetUserIdsWithFlaggedAppRegistrationRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetUserIdsWithFlaggedAppRegistrationResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetUserIdsWithFlaggedAppRegistrationResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetUserIdsWithFlaggedAppRegistrationResponseable), nil
}
