package setmobiledevicemanagementauthority

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SetMobileDeviceManagementAuthorityRequestBuilder provides operations to call the setMobileDeviceManagementAuthority method.
type SetMobileDeviceManagementAuthorityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SetMobileDeviceManagementAuthorityRequestBuilderPostOptions options for Post
type SetMobileDeviceManagementAuthorityRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSetMobileDeviceManagementAuthorityRequestBuilderInternal instantiates a new SetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewSetMobileDeviceManagementAuthorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetMobileDeviceManagementAuthorityRequestBuilder) {
    m := &SetMobileDeviceManagementAuthorityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization_id}/microsoft.graph.setMobileDeviceManagementAuthority";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetMobileDeviceManagementAuthorityRequestBuilder instantiates a new SetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewSetMobileDeviceManagementAuthorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetMobileDeviceManagementAuthorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetMobileDeviceManagementAuthorityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation set mobile device management authority
func (m *SetMobileDeviceManagementAuthorityRequestBuilder) CreatePostRequestInformation(options *SetMobileDeviceManagementAuthorityRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Post set mobile device management authority
func (m *SetMobileDeviceManagementAuthorityRequestBuilder) Post(options *SetMobileDeviceManagementAuthorityRequestBuilderPostOptions)(SetMobileDeviceManagementAuthorityResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSetMobileDeviceManagementAuthorityResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(SetMobileDeviceManagementAuthorityResponseable), nil
}
