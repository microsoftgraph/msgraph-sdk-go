package getomasettingplaintextvaluewithsecretreferencevalueid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder provides operations to call the getOmaSettingPlainTextValue method.
type GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetOptions options for Get
type GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal instantiates a new GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, secretReferenceValueId *string)(*GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    m := &GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration_id}/microsoft.graph.getOmaSettingPlainTextValue(secretReferenceValueId='{secretReferenceValueId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if secretReferenceValueId != nil {
        urlTplParams["secretReferenceValueId"] = *secretReferenceValueId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder instantiates a new GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getOmaSettingPlainTextValue
func (m *GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) CreateGetRequestInformation(options *GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function getOmaSettingPlainTextValue
func (m *GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) Get(options *GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetOptions)(GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable), nil
}
