package primarychannel

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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

// PrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type PrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrimaryChannelRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrimaryChannelRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrimaryChannelRequestBuilderGetQueryParameters the general channel for the team.
type PrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrimaryChannelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrimaryChannelRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrimaryChannelRequestBuilderGetQueryParameters
}
// PrimaryChannelRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrimaryChannelRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration the completeMigration property
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*i1e4013461c5123188ab7ccd33dd297c0d5a7865ba9fea8e6c138306141eeec4a.CompleteMigrationRequestBuilder) {
    return i1e4013461c5123188ab7ccd33dd297c0d5a7865ba9fea8e6c138306141eeec4a.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}/primaryChannel{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property primaryChannel for teams
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property primaryChannel for teams
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *PrimaryChannelRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the general channel for the team.
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the general channel for the team.
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PrimaryChannelRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property primaryChannel in teams
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property primaryChannel in teams
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *PrimaryChannelRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property primaryChannel for teams
func (m *PrimaryChannelRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property primaryChannel for teams
func (m *PrimaryChannelRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *PrimaryChannelRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// FilesFolder the filesFolder property
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*id169406704fae3a188bcd329b3cbae975ed48842ae13c31a260f221158f8516a.FilesFolderRequestBuilder) {
    return id169406704fae3a188bcd329b3cbae975ed48842ae13c31a260f221158f8516a.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the general channel for the team.
func (m *PrimaryChannelRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the general channel for the team.
func (m *PrimaryChannelRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *PrimaryChannelRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChannelFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable), nil
}
// Members the members property
func (m *PrimaryChannelRequestBuilder) Members()(*i977f2dd26da5a2e1880df87bd5372aa7a19a14868a426678151d974f360c2a01.MembersRequestBuilder) {
    return i977f2dd26da5a2e1880df87bd5372aa7a19a14868a426678151d974f360c2a01.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.primaryChannel.members.item collection
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*ie734300e5bedd84acd036f8be309a4094f66ee9a58c5ec5b270079e8efc64674.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return ie734300e5bedd84acd036f8be309a4094f66ee9a58c5ec5b270079e8efc64674.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *PrimaryChannelRequestBuilder) Messages()(*i4785692a9d61d4a8638d06fe97e1daa4feb18f7b743e7ae34ff3f07f3d9425b4.MessagesRequestBuilder) {
    return i4785692a9d61d4a8638d06fe97e1daa4feb18f7b743e7ae34ff3f07f3d9425b4.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.primaryChannel.messages.item collection
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*ib00fe6f5242d43b4e0cf713930bcfb1828c83941d27088c19568b60fa6e113d8.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return ib00fe6f5242d43b4e0cf713930bcfb1828c83941d27088c19568b60fa6e113d8.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property primaryChannel in teams
func (m *PrimaryChannelRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property primaryChannel in teams
func (m *PrimaryChannelRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *PrimaryChannelRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ProvisionEmail the provisionEmail property
func (m *PrimaryChannelRequestBuilder) ProvisionEmail()(*i9cea999ed45be284a8003fe37addd7bac8685c5b60e557f3adc8dc8bcf2d27b5.ProvisionEmailRequestBuilder) {
    return i9cea999ed45be284a8003fe37addd7bac8685c5b60e557f3adc8dc8bcf2d27b5.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *PrimaryChannelRequestBuilder) RemoveEmail()(*i26a22db42badb207f51d6bdd9193eccfdc7cad972de53c3620609a905e198f7a.RemoveEmailRequestBuilder) {
    return i26a22db42badb207f51d6bdd9193eccfdc7cad972de53c3620609a905e198f7a.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tabs the tabs property
func (m *PrimaryChannelRequestBuilder) Tabs()(*ib05faad6285a5d11c891d13f0450d43dfbe0ddb57be2fa94177f219559ef0527.TabsRequestBuilder) {
    return ib05faad6285a5d11c891d13f0450d43dfbe0ddb57be2fa94177f219559ef0527.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.primaryChannel.tabs.item collection
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*i0193f7b86687af396beaaecb295b77e66ded553cf49127c0d06bc79b38a4d93f.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i0193f7b86687af396beaaecb295b77e66ded553cf49127c0d06bc79b38a4d93f.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
