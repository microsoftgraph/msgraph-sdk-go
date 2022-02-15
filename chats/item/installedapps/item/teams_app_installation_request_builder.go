package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i35ad4495bd8fba4273f7824b74ad010646602ceb162a7ff8bf77ddc51b7fcc0d "github.com/microsoftgraph/msgraph-sdk-go/chats/item/installedapps/item/teamsappdefinition"
    ia746748f624c7fd88e59d5553c3ef6fa70382acf2b670e75d11ef27cff15bfc4 "github.com/microsoftgraph/msgraph-sdk-go/chats/item/installedapps/item/upgrade"
    ia96b6fb366243413b3c64f01d0b11020e423792cf3dd3d1045d1fd9067c440d3 "github.com/microsoftgraph/msgraph-sdk-go/chats/item/installedapps/item/teamsapp"
)

// TeamsAppInstallationRequestBuilder builds and executes requests for operations under \chats\{chat-id}\installedApps\{teamsAppInstallation-id}
type TeamsAppInstallationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TeamsAppInstallationRequestBuilderDeleteOptions options for Delete
type TeamsAppInstallationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamsAppInstallationRequestBuilderGetOptions options for Get
type TeamsAppInstallationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TeamsAppInstallationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamsAppInstallationRequestBuilderGetQueryParameters a collection of all the apps in the chat. Nullable.
type TeamsAppInstallationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TeamsAppInstallationRequestBuilderPatchOptions options for Patch
type TeamsAppInstallationRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamsAppInstallation;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTeamsAppInstallationRequestBuilderInternal instantiates a new TeamsAppInstallationRequestBuilder and sets the default values.
func NewTeamsAppInstallationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamsAppInstallationRequestBuilder) {
    m := &TeamsAppInstallationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/chats/{chat_id}/installedApps/{teamsAppInstallation_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamsAppInstallationRequestBuilder instantiates a new TeamsAppInstallationRequestBuilder and sets the default values.
func NewTeamsAppInstallationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamsAppInstallationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsAppInstallationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) CreateDeleteRequestInformation(options *TeamsAppInstallationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) CreateGetRequestInformation(options *TeamsAppInstallationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) CreatePatchRequestInformation(options *TeamsAppInstallationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) Delete(options *TeamsAppInstallationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) Get(options *TeamsAppInstallationRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamsAppInstallation, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTeamsAppInstallation() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamsAppInstallation), nil
}
// Patch a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) Patch(options *TeamsAppInstallationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamsAppInstallationRequestBuilder) TeamsApp()(*ia96b6fb366243413b3c64f01d0b11020e423792cf3dd3d1045d1fd9067c440d3.TeamsAppRequestBuilder) {
    return ia96b6fb366243413b3c64f01d0b11020e423792cf3dd3d1045d1fd9067c440d3.NewTeamsAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamsAppInstallationRequestBuilder) TeamsAppDefinition()(*i35ad4495bd8fba4273f7824b74ad010646602ceb162a7ff8bf77ddc51b7fcc0d.TeamsAppDefinitionRequestBuilder) {
    return i35ad4495bd8fba4273f7824b74ad010646602ceb162a7ff8bf77ddc51b7fcc0d.NewTeamsAppDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamsAppInstallationRequestBuilder) Upgrade()(*ia746748f624c7fd88e59d5553c3ef6fa70382acf2b670e75d11ef27cff15bfc4.UpgradeRequestBuilder) {
    return ia746748f624c7fd88e59d5553c3ef6fa70382acf2b670e75d11ef27cff15bfc4.NewUpgradeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
