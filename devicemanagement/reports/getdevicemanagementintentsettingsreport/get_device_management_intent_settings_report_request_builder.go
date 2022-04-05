package getdevicemanagementintentsettingsreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetDeviceManagementIntentSettingsReportRequestBuilder provides operations to call the getDeviceManagementIntentSettingsReport method.
type GetDeviceManagementIntentSettingsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetDeviceManagementIntentSettingsReportRequestBuilderPostOptions options for Post
type GetDeviceManagementIntentSettingsReportRequestBuilderPostOptions struct {
    // 
    Body GetDeviceManagementIntentSettingsReportRequestBodyable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewGetDeviceManagementIntentSettingsReportRequestBuilderInternal instantiates a new GetDeviceManagementIntentSettingsReportRequestBuilder and sets the default values.
func NewGetDeviceManagementIntentSettingsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDeviceManagementIntentSettingsReportRequestBuilder) {
    m := &GetDeviceManagementIntentSettingsReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getDeviceManagementIntentSettingsReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetDeviceManagementIntentSettingsReportRequestBuilder instantiates a new GetDeviceManagementIntentSettingsReportRequestBuilder and sets the default values.
func NewGetDeviceManagementIntentSettingsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDeviceManagementIntentSettingsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDeviceManagementIntentSettingsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getDeviceManagementIntentSettingsReport
func (m *GetDeviceManagementIntentSettingsReportRequestBuilder) CreatePostRequestInformation(options *GetDeviceManagementIntentSettingsReportRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Post invoke action getDeviceManagementIntentSettingsReport
func (m *GetDeviceManagementIntentSettingsReportRequestBuilder) Post(options *GetDeviceManagementIntentSettingsReportRequestBuilderPostOptions)(GetDeviceManagementIntentSettingsReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetDeviceManagementIntentSettingsReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetDeviceManagementIntentSettingsReportResponseable), nil
}
