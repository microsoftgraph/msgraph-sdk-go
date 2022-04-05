package primarychannel

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i071cc3ee6bc185fb886276c1f3efd8ec6e15b5714f2693722e708fe8daf4dc93 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/members"
    i2d14038ee78a67e5e81ad25ae280e9a321ce535132e47bdcd3444fcd225480b0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/tabs"
    i315462e608a5b36c83df901e69b176f16838c9fe367fb4c9689bda68ffba8689 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/provisionemail"
    i38bf35ad321205c9005bbf3779df33e66b3d0e383a60bb5e2de7cd328bf19878 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/removeemail"
    i5adc1197c87f4224fee29e19d75c57a73ac64cc8ab2046b944e0562741836da2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/messages"
    i90d115230845ab88520d70a382633d39459204d01b463530075a8d272b6bc2df "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/completemigration"
    ie108aea5d217e3f702554f6716c5fb1e0b64aafa0a2000b08a7f8a0546ff4b22 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/filesfolder"
    i29255e84bb72516c2adb78c701d232318618bffbfc6a8ad010a0fab9dcf76c2a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/tabs/item"
    icf7f667d074b27249e73afe4352cda079986d7fab80044c7430a06ab4b19ffb6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/members/item"
    id5d57e626e5ebae45655c3d42f1155aa7241ab8ccd9d616b51ea45fd735d239c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/primarychannel/messages/item"
)

// PrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type PrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrimaryChannelRequestBuilderDeleteOptions options for Delete
type PrimaryChannelRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PrimaryChannelRequestBuilderGetOptions options for Get
type PrimaryChannelRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *PrimaryChannelRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// CompleteMigration the completeMigration property
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*i90d115230845ab88520d70a382633d39459204d01b463530075a8d272b6bc2df.CompleteMigrationRequestBuilder) {
    return i90d115230845ab88520d70a382633d39459204d01b463530075a8d272b6bc2df.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/team/primaryChannel{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrimaryChannelRequestBuilder instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property primaryChannel for groups
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformation(options *PrimaryChannelRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
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
// CreateGetRequestInformation the general channel for the team.
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformation(options *PrimaryChannelRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
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
// CreatePatchRequestInformation update the navigation property primaryChannel in groups
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformation(options *PrimaryChannelRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
// Delete delete navigation property primaryChannel for groups
func (m *PrimaryChannelRequestBuilder) Delete(options *PrimaryChannelRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// FilesFolder the filesFolder property
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*ie108aea5d217e3f702554f6716c5fb1e0b64aafa0a2000b08a7f8a0546ff4b22.FilesFolderRequestBuilder) {
    return ie108aea5d217e3f702554f6716c5fb1e0b64aafa0a2000b08a7f8a0546ff4b22.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the general channel for the team.
func (m *PrimaryChannelRequestBuilder) Get(options *PrimaryChannelRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChannelFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable), nil
}
// Members the members property
func (m *PrimaryChannelRequestBuilder) Members()(*i071cc3ee6bc185fb886276c1f3efd8ec6e15b5714f2693722e708fe8daf4dc93.MembersRequestBuilder) {
    return i071cc3ee6bc185fb886276c1f3efd8ec6e15b5714f2693722e708fe8daf4dc93.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.primaryChannel.members.item collection
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*icf7f667d074b27249e73afe4352cda079986d7fab80044c7430a06ab4b19ffb6.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return icf7f667d074b27249e73afe4352cda079986d7fab80044c7430a06ab4b19ffb6.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *PrimaryChannelRequestBuilder) Messages()(*i5adc1197c87f4224fee29e19d75c57a73ac64cc8ab2046b944e0562741836da2.MessagesRequestBuilder) {
    return i5adc1197c87f4224fee29e19d75c57a73ac64cc8ab2046b944e0562741836da2.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.primaryChannel.messages.item collection
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*id5d57e626e5ebae45655c3d42f1155aa7241ab8ccd9d616b51ea45fd735d239c.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return id5d57e626e5ebae45655c3d42f1155aa7241ab8ccd9d616b51ea45fd735d239c.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property primaryChannel in groups
func (m *PrimaryChannelRequestBuilder) Patch(options *PrimaryChannelRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ProvisionEmail the provisionEmail property
func (m *PrimaryChannelRequestBuilder) ProvisionEmail()(*i315462e608a5b36c83df901e69b176f16838c9fe367fb4c9689bda68ffba8689.ProvisionEmailRequestBuilder) {
    return i315462e608a5b36c83df901e69b176f16838c9fe367fb4c9689bda68ffba8689.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *PrimaryChannelRequestBuilder) RemoveEmail()(*i38bf35ad321205c9005bbf3779df33e66b3d0e383a60bb5e2de7cd328bf19878.RemoveEmailRequestBuilder) {
    return i38bf35ad321205c9005bbf3779df33e66b3d0e383a60bb5e2de7cd328bf19878.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tabs the tabs property
func (m *PrimaryChannelRequestBuilder) Tabs()(*i2d14038ee78a67e5e81ad25ae280e9a321ce535132e47bdcd3444fcd225480b0.TabsRequestBuilder) {
    return i2d14038ee78a67e5e81ad25ae280e9a321ce535132e47bdcd3444fcd225480b0.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.primaryChannel.tabs.item collection
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*i29255e84bb72516c2adb78c701d232318618bffbfc6a8ad010a0fab9dcf76c2a.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i29255e84bb72516c2adb78c701d232318618bffbfc6a8ad010a0fab9dcf76c2a.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
