package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i04a148e32be31a86cd21b897c8b55a4508d63dbae47a02eaeb122662dbd2ff9b "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/template"
    i5cc953f63726e53531e5d00c609bd14469e1ce23944c22b6aea0986c18fb0043 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel"
    i8a1cdbeac728d5d9d3409d0d7085c53384ad37435e0292d966ed94bbc4155a05 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/group"
    i9fa0e9d329dc2b42ce0cc0330991bb8f8e864efaaef5061789d895e28321a6b2 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/operations"
    icf925d6e8373dd15bb408b246595c1d2598e1881b555a62a02a76300fc5b7cd2 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/installedapps"
    id463d65124ba412b3980ec713bebe8eb90e4a925515f5f993cd79d5b01b70907 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels"
    ifc923048eff969ef232f17cdaf6c11e18676c5c9e2918bc19001d30cbdc4b5c1 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/members"
    ifd53534f50d40567e607c2213e794582e29aa46a0c07e2d406db231a42a0140a "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule"
    i03b2a35397701f29dd92bcf0dd549e252ab75a1bd539b1b1ca11a58738ba2c51 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/members/item"
    i77d75bcb3611254a7012d361f29a795e840693affe1de5861c58d310feff299c "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/operations/item"
    id12792e5e45c250abd3ecec9bbc023cd6dc63edea08efc7907dcad0da46a8a0d "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item"
    ifaff06990f5e17731e64f75f4cb575f1ada76645d84341378dccf3aef4f66fab "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/installedapps/item"
)

// TeamItemRequestBuilder provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
type TeamItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TeamItemRequestBuilderDeleteOptions options for Delete
type TeamItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamItemRequestBuilderGetOptions options for Get
type TeamItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TeamItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamItemRequestBuilderGetQueryParameters the Microsoft Teams teams that the user is a member of. Read-only. Nullable.
type TeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TeamItemRequestBuilderPatchOptions options for Patch
type TeamItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Teamable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TeamItemRequestBuilder) Channels()(*id463d65124ba412b3980ec713bebe8eb90e4a925515f5f993cd79d5b01b70907.ChannelsRequestBuilder) {
    return id463d65124ba412b3980ec713bebe8eb90e4a925515f5f993cd79d5b01b70907.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.channels.item collection
func (m *TeamItemRequestBuilder) ChannelsById(id string)(*id12792e5e45c250abd3ecec9bbc023cd6dc63edea08efc7907dcad0da46a8a0d.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel_id"] = id
    }
    return id12792e5e45c250abd3ecec9bbc023cd6dc63edea08efc7907dcad0da46a8a0d.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamItemRequestBuilder) {
    m := &TeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamItemRequestBuilder instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property joinedTeams for me
func (m *TeamItemRequestBuilder) CreateDeleteRequestInformation(options *TeamItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *TeamItemRequestBuilder) CreateGetRequestInformation(options *TeamItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property joinedTeams in me
func (m *TeamItemRequestBuilder) CreatePatchRequestInformation(options *TeamItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property joinedTeams for me
func (m *TeamItemRequestBuilder) Delete(options *TeamItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *TeamItemRequestBuilder) Get(options *TeamItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Teamable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateTeamFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Teamable), nil
}
func (m *TeamItemRequestBuilder) Group()(*i8a1cdbeac728d5d9d3409d0d7085c53384ad37435e0292d966ed94bbc4155a05.GroupRequestBuilder) {
    return i8a1cdbeac728d5d9d3409d0d7085c53384ad37435e0292d966ed94bbc4155a05.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) InstalledApps()(*icf925d6e8373dd15bb408b246595c1d2598e1881b555a62a02a76300fc5b7cd2.InstalledAppsRequestBuilder) {
    return icf925d6e8373dd15bb408b246595c1d2598e1881b555a62a02a76300fc5b7cd2.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.installedApps.item collection
func (m *TeamItemRequestBuilder) InstalledAppsById(id string)(*ifaff06990f5e17731e64f75f4cb575f1ada76645d84341378dccf3aef4f66fab.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return ifaff06990f5e17731e64f75f4cb575f1ada76645d84341378dccf3aef4f66fab.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) Members()(*ifc923048eff969ef232f17cdaf6c11e18676c5c9e2918bc19001d30cbdc4b5c1.MembersRequestBuilder) {
    return ifc923048eff969ef232f17cdaf6c11e18676c5c9e2918bc19001d30cbdc4b5c1.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.members.item collection
func (m *TeamItemRequestBuilder) MembersById(id string)(*i03b2a35397701f29dd92bcf0dd549e252ab75a1bd539b1b1ca11a58738ba2c51.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i03b2a35397701f29dd92bcf0dd549e252ab75a1bd539b1b1ca11a58738ba2c51.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) Operations()(*i9fa0e9d329dc2b42ce0cc0330991bb8f8e864efaaef5061789d895e28321a6b2.OperationsRequestBuilder) {
    return i9fa0e9d329dc2b42ce0cc0330991bb8f8e864efaaef5061789d895e28321a6b2.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.operations.item collection
func (m *TeamItemRequestBuilder) OperationsById(id string)(*i77d75bcb3611254a7012d361f29a795e840693affe1de5861c58d310feff299c.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation_id"] = id
    }
    return i77d75bcb3611254a7012d361f29a795e840693affe1de5861c58d310feff299c.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property joinedTeams in me
func (m *TeamItemRequestBuilder) Patch(options *TeamItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamItemRequestBuilder) PrimaryChannel()(*i5cc953f63726e53531e5d00c609bd14469e1ce23944c22b6aea0986c18fb0043.PrimaryChannelRequestBuilder) {
    return i5cc953f63726e53531e5d00c609bd14469e1ce23944c22b6aea0986c18fb0043.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) Schedule()(*ifd53534f50d40567e607c2213e794582e29aa46a0c07e2d406db231a42a0140a.ScheduleRequestBuilder) {
    return ifd53534f50d40567e607c2213e794582e29aa46a0c07e2d406db231a42a0140a.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) Template()(*i04a148e32be31a86cd21b897c8b55a4508d63dbae47a02eaeb122662dbd2ff9b.TemplateRequestBuilder) {
    return i04a148e32be31a86cd21b897c8b55a4508d63dbae47a02eaeb122662dbd2ff9b.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
