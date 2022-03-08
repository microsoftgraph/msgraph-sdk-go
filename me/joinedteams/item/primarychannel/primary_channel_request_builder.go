package primarychannel

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i23c2230a9d26cc65f1377304e7e2458d1952291a8ec639fd039abd93267cf215 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/messages"
    i3772d1545a94e182a501f85827b513f4692061bccbe8b23eb73c53a831f72d14 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/removeemail"
    i42121f7ad8826a000f65df3b00059d6fc81d8554c4cbbff019865a887d20a56a "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/provisionemail"
    i57e14fb40bbf624605d5248ef685ec9b4257446262e6f8bd78b5c13c71e1e3df "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/tabs"
    i75cf55986f617df38b649b4ebcf62bc92728e8a45cbbb6b95f7040252e76e32c "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/completemigration"
    i7dcb24e562a6855f9b4081d07d2b6077073bb1827c1841433eb6644b41bca8f2 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/filesfolder"
    i8cdf95eab0d1df8385b6f1bdd3b098788d83540da047f10671bda06132b1adb3 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/members"
    i11b5e61a109dfce2102d2d81a8a621a28b579c197305053d9f61fd8e3a46a788 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/members/item"
    i7dfaad961bf25e32c3ce15ea57029e5e024ff862e31e23a16254e10f24113084 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/messages/item"
    ibeddd7f15f21eea6c1989a2c2ecf37d69b81be13bb77e9a96fd8bb8e0a63c3f7 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel/tabs/item"
)

// PrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type PrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrimaryChannelRequestBuilderDeleteOptions options for Delete
type PrimaryChannelRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrimaryChannelRequestBuilderGetOptions options for Get
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
// PrimaryChannelRequestBuilderGetQueryParameters the general channel for the team.
type PrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrimaryChannelRequestBuilderPatchOptions options for Patch
type PrimaryChannelRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channelable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*i75cf55986f617df38b649b4ebcf62bc92728e8a45cbbb6b95f7040252e76e32c.CompleteMigrationRequestBuilder) {
    return i75cf55986f617df38b649b4ebcf62bc92728e8a45cbbb6b95f7040252e76e32c.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team_id}/primaryChannel{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrimaryChannelRequestBuilder instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property primaryChannel for me
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
// CreateGetRequestInformation the general channel for the team.
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
// CreatePatchRequestInformation update the navigation property primaryChannel in me
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
// Delete delete navigation property primaryChannel for me
func (m *PrimaryChannelRequestBuilder) Delete(options *PrimaryChannelRequestBuilderDeleteOptions)(error) {
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
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*i7dcb24e562a6855f9b4081d07d2b6077073bb1827c1841433eb6644b41bca8f2.FilesFolderRequestBuilder) {
    return i7dcb24e562a6855f9b4081d07d2b6077073bb1827c1841433eb6644b41bca8f2.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the general channel for the team.
func (m *PrimaryChannelRequestBuilder) Get(options *PrimaryChannelRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateChannelFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channelable), nil
}
func (m *PrimaryChannelRequestBuilder) Members()(*i8cdf95eab0d1df8385b6f1bdd3b098788d83540da047f10671bda06132b1adb3.MembersRequestBuilder) {
    return i8cdf95eab0d1df8385b6f1bdd3b098788d83540da047f10671bda06132b1adb3.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.primaryChannel.members.item collection
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*i11b5e61a109dfce2102d2d81a8a621a28b579c197305053d9f61fd8e3a46a788.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i11b5e61a109dfce2102d2d81a8a621a28b579c197305053d9f61fd8e3a46a788.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Messages()(*i23c2230a9d26cc65f1377304e7e2458d1952291a8ec639fd039abd93267cf215.MessagesRequestBuilder) {
    return i23c2230a9d26cc65f1377304e7e2458d1952291a8ec639fd039abd93267cf215.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.primaryChannel.messages.item collection
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*i7dfaad961bf25e32c3ce15ea57029e5e024ff862e31e23a16254e10f24113084.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return i7dfaad961bf25e32c3ce15ea57029e5e024ff862e31e23a16254e10f24113084.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property primaryChannel in me
func (m *PrimaryChannelRequestBuilder) Patch(options *PrimaryChannelRequestBuilderPatchOptions)(error) {
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
func (m *PrimaryChannelRequestBuilder) ProvisionEmail()(*i42121f7ad8826a000f65df3b00059d6fc81d8554c4cbbff019865a887d20a56a.ProvisionEmailRequestBuilder) {
    return i42121f7ad8826a000f65df3b00059d6fc81d8554c4cbbff019865a887d20a56a.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) RemoveEmail()(*i3772d1545a94e182a501f85827b513f4692061bccbe8b23eb73c53a831f72d14.RemoveEmailRequestBuilder) {
    return i3772d1545a94e182a501f85827b513f4692061bccbe8b23eb73c53a831f72d14.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Tabs()(*i57e14fb40bbf624605d5248ef685ec9b4257446262e6f8bd78b5c13c71e1e3df.TabsRequestBuilder) {
    return i57e14fb40bbf624605d5248ef685ec9b4257446262e6f8bd78b5c13c71e1e3df.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.primaryChannel.tabs.item collection
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*ibeddd7f15f21eea6c1989a2c2ecf37d69b81be13bb77e9a96fd8bb8e0a63c3f7.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return ibeddd7f15f21eea6c1989a2c2ecf37d69b81be13bb77e9a96fd8bb8e0a63c3f7.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
