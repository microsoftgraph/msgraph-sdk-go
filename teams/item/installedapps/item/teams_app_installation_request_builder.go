package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i39c43799289ea05ffa89f70f403a251dd73339eb42892e687ec303918b2e3268 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/installedapps/item/teamsappdefinition"
    i566d96240508f5329f9999b97147cc637c3eaf2b4471817ddc74c72d8e24d796 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/installedapps/item/upgrade"
    ief8ef10571369d4c38113b3147db690970d1488c5c2f21c4462380a4e31ac98c "github.com/microsoftgraph/msgraph-sdk-go/teams/item/installedapps/item/teamsapp"
)

// Builds and executes requests for operations under \teams\{team-id}\installedApps\{teamsAppInstallation-id}
type TeamsAppInstallationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The apps installed in this team.
type TeamsAppInstallationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new TeamsAppInstallationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTeamsAppInstallationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamsAppInstallationRequestBuilder) {
    m := &TeamsAppInstallationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/teams/{team_id}/installedApps/{teamsAppInstallation_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TeamsAppInstallationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTeamsAppInstallationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamsAppInstallationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsAppInstallationRequestBuilderInternal(urlParams, requestAdapter)
}
// The apps installed in this team.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *TeamsAppInstallationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The apps installed in this team.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *TeamsAppInstallationRequestBuilder) CreateGetRequestInformation(q func (value *TeamsAppInstallationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TeamsAppInstallationRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The apps installed in this team.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *TeamsAppInstallationRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamsAppInstallation, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The apps installed in this team.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TeamsAppInstallationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// The apps installed in this team.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TeamsAppInstallationRequestBuilder) Get(q func (value *TeamsAppInstallationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamsAppInstallation, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTeamsAppInstallation() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamsAppInstallation), nil
}
// The apps installed in this team.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TeamsAppInstallationRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamsAppInstallation, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamsAppInstallationRequestBuilder) TeamsApp()(*ief8ef10571369d4c38113b3147db690970d1488c5c2f21c4462380a4e31ac98c.TeamsAppRequestBuilder) {
    return ief8ef10571369d4c38113b3147db690970d1488c5c2f21c4462380a4e31ac98c.NewTeamsAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamsAppInstallationRequestBuilder) TeamsAppDefinition()(*i39c43799289ea05ffa89f70f403a251dd73339eb42892e687ec303918b2e3268.TeamsAppDefinitionRequestBuilder) {
    return i39c43799289ea05ffa89f70f403a251dd73339eb42892e687ec303918b2e3268.NewTeamsAppDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamsAppInstallationRequestBuilder) Upgrade()(*i566d96240508f5329f9999b97147cc637c3eaf2b4471817ddc74c72d8e24d796.UpgradeRequestBuilder) {
    return i566d96240508f5329f9999b97147cc637c3eaf2b4471817ddc74c72d8e24d796.NewUpgradeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
