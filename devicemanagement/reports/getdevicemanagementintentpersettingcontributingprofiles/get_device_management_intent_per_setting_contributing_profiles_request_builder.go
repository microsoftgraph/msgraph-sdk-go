package getdevicemanagementintentpersettingcontributingprofiles

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
type GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostOptions options for Post
type GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostOptions struct {
    // 
    Body GetDeviceManagementIntentPerSettingContributingProfilesRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal instantiates a new GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder and sets the default values.
func NewGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    m := &GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getDeviceManagementIntentPerSettingContributingProfiles";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder instantiates a new GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder and sets the default values.
func NewGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getDeviceManagementIntentPerSettingContributingProfiles
func (m *GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) CreatePostRequestInformation(options *GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action getDeviceManagementIntentPerSettingContributingProfiles
func (m *GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) Post(options *GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostOptions)(GetDeviceManagementIntentPerSettingContributingProfilesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetDeviceManagementIntentPerSettingContributingProfilesResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetDeviceManagementIntentPerSettingContributingProfilesResponseable), nil
}
