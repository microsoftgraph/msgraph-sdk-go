package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5363721350f81d4226008b2a0abc5fa76fabff34763dc7de5066c621e0c0f6be "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/operations"
    i539671039e993ec9795209a4951310067d75618d6c4524b564660e6d7f6dec35 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/template"
    i6ea3db8c980a5b6b549dcba6aaa6b387b8c6b1cb073e6178a10627e6fb06a389 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/primarychannel"
    i81bd7dd61df3d95dbc4620ad1479451afe1469239e95f230e4f09a46b785da75 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule"
    ib1deb2599b9a847d5026bfa2b1b44e59f2a829010f536668074417c636ce36db "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels"
    ic44d4f0e5acd07f5481d8f8b84a79cab11dfbee059cf6770fd4e2b23db0c1a93 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/group"
    ic914f26c893a65c2d7a0c81250eacc4c0f9d860f64ab1b1cf1c364c6743699a8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/installedapps"
    ief730b7100e12b3385f77770a9d755bebd575023d52fdbb9414e8c86230b027d "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/members"
    i6d038dcd104603d7281b4eac363006d79b67e92eca3e8c364b49d02b84649d79 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item"
    i98b219e6f0357bddb879690365c172b8f5fa9226021c11592a8062b6d8a9415e "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/members/item"
    ie5af489748d2fc36f15958b988ad08901566118f084112092561d1d1b446b768 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/installedapps/item"
    ie78978d46a3704a561b7ecbd9e66a5ee9be50b99960e26229ec758c80ee98271 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/operations/item"
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
func (m *TeamItemRequestBuilder) Channels()(*ib1deb2599b9a847d5026bfa2b1b44e59f2a829010f536668074417c636ce36db.ChannelsRequestBuilder) {
    return ib1deb2599b9a847d5026bfa2b1b44e59f2a829010f536668074417c636ce36db.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.channels.item collection
func (m *TeamItemRequestBuilder) ChannelsById(id string)(*i6d038dcd104603d7281b4eac363006d79b67e92eca3e8c364b49d02b84649d79.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel_id"] = id
    }
    return i6d038dcd104603d7281b4eac363006d79b67e92eca3e8c364b49d02b84649d79.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamItemRequestBuilder) {
    m := &TeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/joinedTeams/{team_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property joinedTeams for users
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
// CreatePatchRequestInformation update the navigation property joinedTeams in users
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
// Delete delete navigation property joinedTeams for users
func (m *TeamItemRequestBuilder) Delete(options *TeamItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
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
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateTeamFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Teamable), nil
}
func (m *TeamItemRequestBuilder) Group()(*ic44d4f0e5acd07f5481d8f8b84a79cab11dfbee059cf6770fd4e2b23db0c1a93.GroupRequestBuilder) {
    return ic44d4f0e5acd07f5481d8f8b84a79cab11dfbee059cf6770fd4e2b23db0c1a93.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) InstalledApps()(*ic914f26c893a65c2d7a0c81250eacc4c0f9d860f64ab1b1cf1c364c6743699a8.InstalledAppsRequestBuilder) {
    return ic914f26c893a65c2d7a0c81250eacc4c0f9d860f64ab1b1cf1c364c6743699a8.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.installedApps.item collection
func (m *TeamItemRequestBuilder) InstalledAppsById(id string)(*ie5af489748d2fc36f15958b988ad08901566118f084112092561d1d1b446b768.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return ie5af489748d2fc36f15958b988ad08901566118f084112092561d1d1b446b768.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) Members()(*ief730b7100e12b3385f77770a9d755bebd575023d52fdbb9414e8c86230b027d.MembersRequestBuilder) {
    return ief730b7100e12b3385f77770a9d755bebd575023d52fdbb9414e8c86230b027d.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.members.item collection
func (m *TeamItemRequestBuilder) MembersById(id string)(*i98b219e6f0357bddb879690365c172b8f5fa9226021c11592a8062b6d8a9415e.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i98b219e6f0357bddb879690365c172b8f5fa9226021c11592a8062b6d8a9415e.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) Operations()(*i5363721350f81d4226008b2a0abc5fa76fabff34763dc7de5066c621e0c0f6be.OperationsRequestBuilder) {
    return i5363721350f81d4226008b2a0abc5fa76fabff34763dc7de5066c621e0c0f6be.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.operations.item collection
func (m *TeamItemRequestBuilder) OperationsById(id string)(*ie78978d46a3704a561b7ecbd9e66a5ee9be50b99960e26229ec758c80ee98271.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation_id"] = id
    }
    return ie78978d46a3704a561b7ecbd9e66a5ee9be50b99960e26229ec758c80ee98271.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property joinedTeams in users
func (m *TeamItemRequestBuilder) Patch(options *TeamItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamItemRequestBuilder) PrimaryChannel()(*i6ea3db8c980a5b6b549dcba6aaa6b387b8c6b1cb073e6178a10627e6fb06a389.PrimaryChannelRequestBuilder) {
    return i6ea3db8c980a5b6b549dcba6aaa6b387b8c6b1cb073e6178a10627e6fb06a389.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) Schedule()(*i81bd7dd61df3d95dbc4620ad1479451afe1469239e95f230e4f09a46b785da75.ScheduleRequestBuilder) {
    return i81bd7dd61df3d95dbc4620ad1479451afe1469239e95f230e4f09a46b785da75.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamItemRequestBuilder) Template()(*i539671039e993ec9795209a4951310067d75618d6c4524b564660e6d7f6dec35.TemplateRequestBuilder) {
    return i539671039e993ec9795209a4951310067d75618d6c4524b564660e6d7f6dec35.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
