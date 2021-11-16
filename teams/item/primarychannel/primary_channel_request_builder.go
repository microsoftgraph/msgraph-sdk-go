package primarychannel

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1e4013461c5123188ab7ccd33dd297c0d5a7865ba9fea8e6c138306141eeec4a "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/completemigration"
    i26a22db42badb207f51d6bdd9193eccfdc7cad972de53c3620609a905e198f7a "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/removeemail"
    i4785692a9d61d4a8638d06fe97e1daa4feb18f7b743e7ae34ff3f07f3d9425b4 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/messages"
    i977f2dd26da5a2e1880df87bd5372aa7a19a14868a426678151d974f360c2a01 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/members"
    i9cea999ed45be284a8003fe37addd7bac8685c5b60e557f3adc8dc8bcf2d27b5 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/provisionemail"
    ib05faad6285a5d11c891d13f0450d43dfbe0ddb57be2fa94177f219559ef0527 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/tabs"
    id169406704fae3a188bcd329b3cbae975ed48842ae13c31a260f221158f8516a "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/filesfolder"
    i0193f7b86687af396beaaecb295b77e66ded553cf49127c0d06bc79b38a4d93f "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/tabs/item"
    ib00fe6f5242d43b4e0cf713930bcfb1828c83941d27088c19568b60fa6e113d8 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/messages/item"
    ie734300e5bedd84acd036f8be309a4094f66ee9a58c5ec5b270079e8efc64674 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/members/item"
)

// Builds and executes requests for operations under \teams\{team-id}\primaryChannel
type PrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type PrimaryChannelRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type PrimaryChannelRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrimaryChannelRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The general channel for the team.
type PrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type PrimaryChannelRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*i1e4013461c5123188ab7ccd33dd297c0d5a7865ba9fea8e6c138306141eeec4a.CompleteMigrationRequestBuilder) {
    return i1e4013461c5123188ab7ccd33dd297c0d5a7865ba9fea8e6c138306141eeec4a.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new PrimaryChannelRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}/primaryChannel{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PrimaryChannelRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrimaryChannelRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// The general channel for the team.
// Parameters:
//  - options : Options for the request
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformation(options *PrimaryChannelRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The general channel for the team.
// Parameters:
//  - options : Options for the request
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformation(options *PrimaryChannelRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The general channel for the team.
// Parameters:
//  - options : Options for the request
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformation(options *PrimaryChannelRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The general channel for the team.
// Parameters:
//  - options : Options for the request
func (m *PrimaryChannelRequestBuilder) Delete(options *PrimaryChannelRequestBuilderDeleteOptions)(error) {
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
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*id169406704fae3a188bcd329b3cbae975ed48842ae13c31a260f221158f8516a.FilesFolderRequestBuilder) {
    return id169406704fae3a188bcd329b3cbae975ed48842ae13c31a260f221158f8516a.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The general channel for the team.
// Parameters:
//  - options : Options for the request
func (m *PrimaryChannelRequestBuilder) Get(options *PrimaryChannelRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewChannel() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel), nil
}
func (m *PrimaryChannelRequestBuilder) Members()(*i977f2dd26da5a2e1880df87bd5372aa7a19a14868a426678151d974f360c2a01.MembersRequestBuilder) {
    return i977f2dd26da5a2e1880df87bd5372aa7a19a14868a426678151d974f360c2a01.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.primaryChannel.members.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*ie734300e5bedd84acd036f8be309a4094f66ee9a58c5ec5b270079e8efc64674.ConversationMemberRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ie734300e5bedd84acd036f8be309a4094f66ee9a58c5ec5b270079e8efc64674.NewConversationMemberRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Messages()(*i4785692a9d61d4a8638d06fe97e1daa4feb18f7b743e7ae34ff3f07f3d9425b4.MessagesRequestBuilder) {
    return i4785692a9d61d4a8638d06fe97e1daa4feb18f7b743e7ae34ff3f07f3d9425b4.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.primaryChannel.messages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*ib00fe6f5242d43b4e0cf713930bcfb1828c83941d27088c19568b60fa6e113d8.ChatMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return ib00fe6f5242d43b4e0cf713930bcfb1828c83941d27088c19568b60fa6e113d8.NewChatMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The general channel for the team.
// Parameters:
//  - options : Options for the request
func (m *PrimaryChannelRequestBuilder) Patch(options *PrimaryChannelRequestBuilderPatchOptions)(error) {
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
func (m *PrimaryChannelRequestBuilder) ProvisionEmail()(*i9cea999ed45be284a8003fe37addd7bac8685c5b60e557f3adc8dc8bcf2d27b5.ProvisionEmailRequestBuilder) {
    return i9cea999ed45be284a8003fe37addd7bac8685c5b60e557f3adc8dc8bcf2d27b5.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) RemoveEmail()(*i26a22db42badb207f51d6bdd9193eccfdc7cad972de53c3620609a905e198f7a.RemoveEmailRequestBuilder) {
    return i26a22db42badb207f51d6bdd9193eccfdc7cad972de53c3620609a905e198f7a.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Tabs()(*ib05faad6285a5d11c891d13f0450d43dfbe0ddb57be2fa94177f219559ef0527.TabsRequestBuilder) {
    return ib05faad6285a5d11c891d13f0450d43dfbe0ddb57be2fa94177f219559ef0527.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.primaryChannel.tabs.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*i0193f7b86687af396beaaecb295b77e66ded553cf49127c0d06bc79b38a4d93f.TeamsTabRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i0193f7b86687af396beaaecb295b77e66ded553cf49127c0d06bc79b38a4d93f.NewTeamsTabRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
