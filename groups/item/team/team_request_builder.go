package team

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i04e49204322d308913c8af78daf5ae0e14fb7df81490288b98bd1c408480bbb4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/group"
    i34294a7456473fba252b3ca961ed3381f55b7a770d74ee4720d752d9a45ad739 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel"
    i477873d925c8fe5f007dfb4043f82ed6374d80069f96d4931626a23aeb0f2a87 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/template"
    i495e85248bfc2c7bf36f36ede17adf8bc672f00e69df645f352a27b2b7d30e61 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule"
    i7aee2305413b3b465064d02a6aa79c0f32df27bc874d76350b572847a880a22f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/installedapps"
    i8f663f97c6eef0cee280b998b250f55e4d63fc289392b2a09af9c95210ca5d5d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels"
    i97dd0d985d4eaf1cbaeb3d771ef48c101ec691107d9a79dcd3f38213ccae863b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/members"
    i9f6a3c1c4300552f4792a53f0d0fe603636d1b865d408ff28e1cdd722683ebe3 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/operations"
    i6e500f69af937bd22bbfc2b6ee2698da77bfc652c449bb8ce6a2ee75470b9de9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/installedapps/item"
    id5fe3547558f37b7931bf00eb734fbe581be897e9aea405e91e3ce2dbeec74b3 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/operations/item"
    ie59099b1e42dc6f55628ccb467e95d4e7416a9bd33399bd6beff110ff67ea34e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/members/item"
    ie594d6624cc5e7f45c840beec42560c6f347a2da3e424188e2fbe60f77fbc850 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item"
)

// TeamRequestBuilder provides operations to manage the team property of the microsoft.graph.group entity.
type TeamRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TeamRequestBuilderDeleteOptions options for Delete
type TeamRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamRequestBuilderGetOptions options for Get
type TeamRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TeamRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamRequestBuilderGetQueryParameters get team from groups
type TeamRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TeamRequestBuilderPatchOptions options for Patch
type TeamRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Teamable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TeamRequestBuilder) Channels()(*i8f663f97c6eef0cee280b998b250f55e4d63fc289392b2a09af9c95210ca5d5d.ChannelsRequestBuilder) {
    return i8f663f97c6eef0cee280b998b250f55e4d63fc289392b2a09af9c95210ca5d5d.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.channels.item collection
func (m *TeamRequestBuilder) ChannelsById(id string)(*ie594d6624cc5e7f45c840beec42560c6f347a2da3e424188e2fbe60f77fbc850.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel_id"] = id
    }
    return ie594d6624cc5e7f45c840beec42560c6f347a2da3e424188e2fbe60f77fbc850.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamRequestBuilder) {
    m := &TeamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/team{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamRequestBuilder instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property team for groups
func (m *TeamRequestBuilder) CreateDeleteRequestInformation(options *TeamRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get team from groups
func (m *TeamRequestBuilder) CreateGetRequestInformation(options *TeamRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property team in groups
func (m *TeamRequestBuilder) CreatePatchRequestInformation(options *TeamRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property team for groups
func (m *TeamRequestBuilder) Delete(options *TeamRequestBuilderDeleteOptions)(error) {
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
// Get get team from groups
func (m *TeamRequestBuilder) Get(options *TeamRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Teamable, error) {
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
func (m *TeamRequestBuilder) Group()(*i04e49204322d308913c8af78daf5ae0e14fb7df81490288b98bd1c408480bbb4.GroupRequestBuilder) {
    return i04e49204322d308913c8af78daf5ae0e14fb7df81490288b98bd1c408480bbb4.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) InstalledApps()(*i7aee2305413b3b465064d02a6aa79c0f32df27bc874d76350b572847a880a22f.InstalledAppsRequestBuilder) {
    return i7aee2305413b3b465064d02a6aa79c0f32df27bc874d76350b572847a880a22f.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.installedApps.item collection
func (m *TeamRequestBuilder) InstalledAppsById(id string)(*i6e500f69af937bd22bbfc2b6ee2698da77bfc652c449bb8ce6a2ee75470b9de9.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return i6e500f69af937bd22bbfc2b6ee2698da77bfc652c449bb8ce6a2ee75470b9de9.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Members()(*i97dd0d985d4eaf1cbaeb3d771ef48c101ec691107d9a79dcd3f38213ccae863b.MembersRequestBuilder) {
    return i97dd0d985d4eaf1cbaeb3d771ef48c101ec691107d9a79dcd3f38213ccae863b.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.members.item collection
func (m *TeamRequestBuilder) MembersById(id string)(*ie59099b1e42dc6f55628ccb467e95d4e7416a9bd33399bd6beff110ff67ea34e.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ie59099b1e42dc6f55628ccb467e95d4e7416a9bd33399bd6beff110ff67ea34e.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Operations()(*i9f6a3c1c4300552f4792a53f0d0fe603636d1b865d408ff28e1cdd722683ebe3.OperationsRequestBuilder) {
    return i9f6a3c1c4300552f4792a53f0d0fe603636d1b865d408ff28e1cdd722683ebe3.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.operations.item collection
func (m *TeamRequestBuilder) OperationsById(id string)(*id5fe3547558f37b7931bf00eb734fbe581be897e9aea405e91e3ce2dbeec74b3.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation_id"] = id
    }
    return id5fe3547558f37b7931bf00eb734fbe581be897e9aea405e91e3ce2dbeec74b3.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property team in groups
func (m *TeamRequestBuilder) Patch(options *TeamRequestBuilderPatchOptions)(error) {
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
func (m *TeamRequestBuilder) PrimaryChannel()(*i34294a7456473fba252b3ca961ed3381f55b7a770d74ee4720d752d9a45ad739.PrimaryChannelRequestBuilder) {
    return i34294a7456473fba252b3ca961ed3381f55b7a770d74ee4720d752d9a45ad739.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) Schedule()(*i495e85248bfc2c7bf36f36ede17adf8bc672f00e69df645f352a27b2b7d30e61.ScheduleRequestBuilder) {
    return i495e85248bfc2c7bf36f36ede17adf8bc672f00e69df645f352a27b2b7d30e61.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) Template()(*i477873d925c8fe5f007dfb4043f82ed6374d80069f96d4931626a23aeb0f2a87.TemplateRequestBuilder) {
    return i477873d925c8fe5f007dfb4043f82ed6374d80069f96d4931626a23aeb0f2a87.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
