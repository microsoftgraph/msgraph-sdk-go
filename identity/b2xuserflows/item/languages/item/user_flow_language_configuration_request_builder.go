package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i6efee475ffa3033985a2f1ade73cee7661b1f6f7afcf650570d574bea030ce03 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages/item/overridespages"
    i97138cb6b4f142a12dfbe5eeba4da3da641e9238e6170e8d10e44cdf6c46efb2 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages/item/defaultpages"
    i3ea30c4cbd180dc3a448c556906b82bed378364ac709419c5058d94bee5e216a "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages/item/overridespages/item"
    iaa3b91f02a3db748d21d8b30fa842ec6a841f6d30579d8378072348ce9f6a367 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages/item/defaultpages/item"
)

// UserFlowLanguageConfigurationRequestBuilder builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\languages\{userFlowLanguageConfiguration-id}
type UserFlowLanguageConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserFlowLanguageConfigurationRequestBuilderDeleteOptions options for Delete
type UserFlowLanguageConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserFlowLanguageConfigurationRequestBuilderGetOptions options for Get
type UserFlowLanguageConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserFlowLanguageConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserFlowLanguageConfigurationRequestBuilderGetQueryParameters the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
type UserFlowLanguageConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserFlowLanguageConfigurationRequestBuilderPatchOptions options for Patch
type UserFlowLanguageConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UserFlowLanguageConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserFlowLanguageConfigurationRequestBuilderInternal instantiates a new UserFlowLanguageConfigurationRequestBuilder and sets the default values.
func NewUserFlowLanguageConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowLanguageConfigurationRequestBuilder) {
    m := &UserFlowLanguageConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow_id}/languages/{userFlowLanguageConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserFlowLanguageConfigurationRequestBuilder instantiates a new UserFlowLanguageConfigurationRequestBuilder and sets the default values.
func NewUserFlowLanguageConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowLanguageConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserFlowLanguageConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *UserFlowLanguageConfigurationRequestBuilder) CreateDeleteRequestInformation(options *UserFlowLanguageConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *UserFlowLanguageConfigurationRequestBuilder) CreateGetRequestInformation(options *UserFlowLanguageConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// CreatePatchRequestInformation the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *UserFlowLanguageConfigurationRequestBuilder) CreatePatchRequestInformation(options *UserFlowLanguageConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
func (m *UserFlowLanguageConfigurationRequestBuilder) DefaultPages()(*i97138cb6b4f142a12dfbe5eeba4da3da641e9238e6170e8d10e44cdf6c46efb2.DefaultPagesRequestBuilder) {
    return i97138cb6b4f142a12dfbe5eeba4da3da641e9238e6170e8d10e44cdf6c46efb2.NewDefaultPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultPagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.languages.item.defaultPages.item collection
func (m *UserFlowLanguageConfigurationRequestBuilder) DefaultPagesById(id string)(*iaa3b91f02a3db748d21d8b30fa842ec6a841f6d30579d8378072348ce9f6a367.UserFlowLanguagePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage_id"] = id
    }
    return iaa3b91f02a3db748d21d8b30fa842ec6a841f6d30579d8378072348ce9f6a367.NewUserFlowLanguagePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *UserFlowLanguageConfigurationRequestBuilder) Delete(options *UserFlowLanguageConfigurationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *UserFlowLanguageConfigurationRequestBuilder) Get(options *UserFlowLanguageConfigurationRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UserFlowLanguageConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewUserFlowLanguageConfiguration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UserFlowLanguageConfiguration), nil
}
func (m *UserFlowLanguageConfigurationRequestBuilder) OverridesPages()(*i6efee475ffa3033985a2f1ade73cee7661b1f6f7afcf650570d574bea030ce03.OverridesPagesRequestBuilder) {
    return i6efee475ffa3033985a2f1ade73cee7661b1f6f7afcf650570d574bea030ce03.NewOverridesPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OverridesPagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.languages.item.overridesPages.item collection
func (m *UserFlowLanguageConfigurationRequestBuilder) OverridesPagesById(id string)(*i3ea30c4cbd180dc3a448c556906b82bed378364ac709419c5058d94bee5e216a.UserFlowLanguagePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage_id"] = id
    }
    return i3ea30c4cbd180dc3a448c556906b82bed378364ac709419c5058d94bee5e216a.NewUserFlowLanguagePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *UserFlowLanguageConfigurationRequestBuilder) Patch(options *UserFlowLanguageConfigurationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
