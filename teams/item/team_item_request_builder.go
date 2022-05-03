package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0d53f661bb96ea28364f561f24c6a15b30ee8383c8b828dac50eca27ee85d3fa "github.com/microsoftgraph/msgraph-sdk-go/teams/item/operations"
    i112483150813fdda19082d152401ea7792fc69e950a26fd2de0f70593b4e9aa3 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule"
    i15af31366dc887244693eaa3d3a5b29b322f057d84108b094ab85c131706bc4d "github.com/microsoftgraph/msgraph-sdk-go/teams/item/sendactivitynotification"
    i3ab10cbdf172c82da95d8c321e7ab0c74610f25b0a5df76707ecc483fe73c949 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/unarchive"
    i6b2db9d1ed1ca06816d98e7327f880f451780be7beb8380462b23278f4c949d8 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/archive"
    i86a4744d5e45722adbed1e0f22b69f1a4b6148d0445ff82385abf3cedd90a29d "github.com/microsoftgraph/msgraph-sdk-go/teams/item/clone"
    i94154a059496bbfe5f4225f44a08db234c17a10c3a7d58124a8fb182751e3bc6 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/installedapps"
    i941dc9f0521a47c6bd3be5dd3a072e94f5381994ffe128d7e54810e76ed8e473 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/template"
    i96200494e669e195393f509bc7a1b076f717becc8f2d0e85d3bdc73e820e2615 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel"
    i9c16b9ccd15275976651d11bc1695c103bd12afe48b01714765458594b9f7975 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels"
    ic0623992ab864953cddbe4e195f8ecf82c663de82e6f4af4bd9402875dbb3257 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/group"
    id19ffdf54b88ab907e8aa643b2969b583758c77cb61e7e304e1e7fbc61c9d1fa "github.com/microsoftgraph/msgraph-sdk-go/teams/item/completemigration"
    ie3ea4358486c3530306ce0ee6275a0285de71b92eff7f63adb5e51ed2ef63235 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/members"
    i167530cb07fd4387f5ffe0c5a728e41dd5608a4963302dd28d98720ea3c093dc "github.com/microsoftgraph/msgraph-sdk-go/teams/item/members/item"
    i3efb7759b468a039683d7e9a7233c6ef567ad6e2bfdb31d83c6c7c4bfb121545 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/operations/item"
    i6b8094a2be2251280648d1c8aa1f2cb616745b9111719fe2f3f1b754098a2ea7 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/installedapps/item"
    i95bd316e1af2146251d5c5f8f94d9865c17745b134282c19c4fae4c81ba47474 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item"
)

// TeamItemRequestBuilder provides operations to manage the collection of team entities.
type TeamItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamItemRequestBuilderGetQueryParameters get entity from teams by key
type TeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamItemRequestBuilderGetQueryParameters
}
// TeamItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Archive the archive property
func (m *TeamItemRequestBuilder) Archive()(*i6b2db9d1ed1ca06816d98e7327f880f451780be7beb8380462b23278f4c949d8.ArchiveRequestBuilder) {
    return i6b2db9d1ed1ca06816d98e7327f880f451780be7beb8380462b23278f4c949d8.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels the channels property
func (m *TeamItemRequestBuilder) Channels()(*i9c16b9ccd15275976651d11bc1695c103bd12afe48b01714765458594b9f7975.ChannelsRequestBuilder) {
    return i9c16b9ccd15275976651d11bc1695c103bd12afe48b01714765458594b9f7975.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.channels.item collection
func (m *TeamItemRequestBuilder) ChannelsById(id string)(*i95bd316e1af2146251d5c5f8f94d9865c17745b134282c19c4fae4c81ba47474.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i95bd316e1af2146251d5c5f8f94d9865c17745b134282c19c4fae4c81ba47474.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone the clone property
func (m *TeamItemRequestBuilder) Clone()(*i86a4744d5e45722adbed1e0f22b69f1a4b6148d0445ff82385abf3cedd90a29d.CloneRequestBuilder) {
    return i86a4744d5e45722adbed1e0f22b69f1a4b6148d0445ff82385abf3cedd90a29d.NewCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration the completeMigration property
func (m *TeamItemRequestBuilder) CompleteMigration()(*id19ffdf54b88ab907e8aa643b2969b583758c77cb61e7e304e1e7fbc61c9d1fa.CompleteMigrationRequestBuilder) {
    return id19ffdf54b88ab907e8aa643b2969b583758c77cb61e7e304e1e7fbc61c9d1fa.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamItemRequestBuilder) {
    m := &TeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamItemRequestBuilder instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from teams
func (m *TeamItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete entity from teams
func (m *TeamItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *TeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from teams by key
func (m *TeamItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get entity from teams by key
func (m *TeamItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TeamItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in teams
func (m *TeamItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update entity in teams
func (m *TeamItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *TeamItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from teams
func (m *TeamItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete entity from teams
func (m *TeamItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *TeamItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get get entity from teams by key
func (m *TeamItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get entity from teams by key
func (m *TeamItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *TeamItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// Group the group property
func (m *TeamItemRequestBuilder) Group()(*ic0623992ab864953cddbe4e195f8ecf82c663de82e6f4af4bd9402875dbb3257.GroupRequestBuilder) {
    return ic0623992ab864953cddbe4e195f8ecf82c663de82e6f4af4bd9402875dbb3257.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledApps the installedApps property
func (m *TeamItemRequestBuilder) InstalledApps()(*i94154a059496bbfe5f4225f44a08db234c17a10c3a7d58124a8fb182751e3bc6.InstalledAppsRequestBuilder) {
    return i94154a059496bbfe5f4225f44a08db234c17a10c3a7d58124a8fb182751e3bc6.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.installedApps.item collection
func (m *TeamItemRequestBuilder) InstalledAppsById(id string)(*i6b8094a2be2251280648d1c8aa1f2cb616745b9111719fe2f3f1b754098a2ea7.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return i6b8094a2be2251280648d1c8aa1f2cb616745b9111719fe2f3f1b754098a2ea7.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *TeamItemRequestBuilder) Members()(*ie3ea4358486c3530306ce0ee6275a0285de71b92eff7f63adb5e51ed2ef63235.MembersRequestBuilder) {
    return ie3ea4358486c3530306ce0ee6275a0285de71b92eff7f63adb5e51ed2ef63235.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.members.item collection
func (m *TeamItemRequestBuilder) MembersById(id string)(*i167530cb07fd4387f5ffe0c5a728e41dd5608a4963302dd28d98720ea3c093dc.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i167530cb07fd4387f5ffe0c5a728e41dd5608a4963302dd28d98720ea3c093dc.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *TeamItemRequestBuilder) Operations()(*i0d53f661bb96ea28364f561f24c6a15b30ee8383c8b828dac50eca27ee85d3fa.OperationsRequestBuilder) {
    return i0d53f661bb96ea28364f561f24c6a15b30ee8383c8b828dac50eca27ee85d3fa.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.operations.item collection
func (m *TeamItemRequestBuilder) OperationsById(id string)(*i3efb7759b468a039683d7e9a7233c6ef567ad6e2bfdb31d83c6c7c4bfb121545.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return i3efb7759b468a039683d7e9a7233c6ef567ad6e2bfdb31d83c6c7c4bfb121545.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in teams
func (m *TeamItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update entity in teams
func (m *TeamItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *TeamItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// PrimaryChannel the primaryChannel property
func (m *TeamItemRequestBuilder) PrimaryChannel()(*i96200494e669e195393f509bc7a1b076f717becc8f2d0e85d3bdc73e820e2615.PrimaryChannelRequestBuilder) {
    return i96200494e669e195393f509bc7a1b076f717becc8f2d0e85d3bdc73e820e2615.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule the schedule property
func (m *TeamItemRequestBuilder) Schedule()(*i112483150813fdda19082d152401ea7792fc69e950a26fd2de0f70593b4e9aa3.ScheduleRequestBuilder) {
    return i112483150813fdda19082d152401ea7792fc69e950a26fd2de0f70593b4e9aa3.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification the sendActivityNotification property
func (m *TeamItemRequestBuilder) SendActivityNotification()(*i15af31366dc887244693eaa3d3a5b29b322f057d84108b094ab85c131706bc4d.SendActivityNotificationRequestBuilder) {
    return i15af31366dc887244693eaa3d3a5b29b322f057d84108b094ab85c131706bc4d.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Template the template property
func (m *TeamItemRequestBuilder) Template()(*i941dc9f0521a47c6bd3be5dd3a072e94f5381994ffe128d7e54810e76ed8e473.TemplateRequestBuilder) {
    return i941dc9f0521a47c6bd3be5dd3a072e94f5381994ffe128d7e54810e76ed8e473.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive the unarchive property
func (m *TeamItemRequestBuilder) Unarchive()(*i3ab10cbdf172c82da95d8c321e7ab0c74610f25b0a5df76707ecc483fe73c949.UnarchiveRequestBuilder) {
    return i3ab10cbdf172c82da95d8c321e7ab0c74610f25b0a5df76707ecc483fe73c949.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
