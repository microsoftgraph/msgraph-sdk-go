package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GroupsGroupItemRequestBuilder provides operations to manage the collection of group entities.
type GroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsGroupItemRequestBuilderGetQueryParameters get the properties and relationships of a group object. This operation returns by default only a subset of all the available properties, as noted in the Properties section. To get properties that are _not_ returned by default, specify them in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query.
type GroupsGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsGroupItemRequestBuilderGetQueryParameters
}
// GroupsGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptedSenders provides operations to manage the acceptedSenders property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) AcceptedSenders()(*GroupsItemAcceptedSendersRequestBuilder) {
    return NewGroupsItemAcceptedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptedSendersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go.groups.item.acceptedSenders.item collection
func (m *GroupsGroupItemRequestBuilder) AcceptedSendersById(id string)(*GroupsItemAcceptedSendersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewGroupsItemAcceptedSendersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AddFavorite provides operations to call the addFavorite method.
func (m *GroupsGroupItemRequestBuilder) AddFavorite()(*GroupsItemAddFavoriteRequestBuilder) {
    return NewGroupsItemAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) AppRoleAssignments()(*GroupsItemAppRoleAssignmentsRequestBuilder) {
    return NewGroupsItemAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) AppRoleAssignmentsById(id string)(*GroupsItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewGroupsItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignLicense provides operations to call the assignLicense method.
func (m *GroupsGroupItemRequestBuilder) AssignLicense()(*GroupsItemAssignLicenseRequestBuilder) {
    return NewGroupsItemAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Calendar()(*GroupsItemCalendarRequestBuilder) {
    return NewGroupsItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) CalendarView()(*GroupsItemCalendarViewRequestBuilder) {
    return NewGroupsItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) CalendarViewById(id string)(*GroupsItemCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewGroupsItemCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckGrantedPermissionsForApp provides operations to call the checkGrantedPermissionsForApp method.
func (m *GroupsGroupItemRequestBuilder) CheckGrantedPermissionsForApp()(*GroupsItemCheckGrantedPermissionsForAppRequestBuilder) {
    return NewGroupsItemCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *GroupsGroupItemRequestBuilder) CheckMemberGroups()(*GroupsItemCheckMemberGroupsRequestBuilder) {
    return NewGroupsItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *GroupsGroupItemRequestBuilder) CheckMemberObjects()(*GroupsItemCheckMemberObjectsRequestBuilder) {
    return NewGroupsItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsGroupItemRequestBuilder) {
    m := &GroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Conversations provides operations to manage the conversations property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Conversations()(*GroupsItemConversationsRequestBuilder) {
    return NewGroupsItemConversationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConversationsById provides operations to manage the conversations property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) ConversationsById(id string)(*GroupsItemConversationsConversationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversation%2Did"] = id
    }
    return NewGroupsItemConversationsConversationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete group. When deleted, Microsoft 365 groups are moved to a temporary container and can be restored within 30 days. After that time, they're permanently deleted. This isn't applicable to Security groups and Distribution groups which are permanently deleted immediately. To learn more, see deletedItems.
func (m *GroupsGroupItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatedOnBehalfOf provides operations to manage the createdOnBehalfOf property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) CreatedOnBehalfOf()(*GroupsItemCreatedOnBehalfOfRequestBuilder) {
    return NewGroupsItemCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the properties and relationships of a group object. This operation returns by default only a subset of all the available properties, as noted in the Properties section. To get properties that are _not_ returned by default, specify them in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query.
func (m *GroupsGroupItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the properties of a group object.
func (m *GroupsGroupItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, requestConfiguration *GroupsGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete group. When deleted, Microsoft 365 groups are moved to a temporary container and can be restored within 30 days. After that time, they're permanently deleted. This isn't applicable to Security groups and Distribution groups which are permanently deleted immediately. To learn more, see deletedItems.
func (m *GroupsGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Drive provides operations to manage the drive property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Drive()(*GroupsItemDriveRequestBuilder) {
    return NewGroupsItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the drives property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Drives()(*GroupsItemDrivesRequestBuilder) {
    return NewGroupsItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) DrivesById(id string)(*GroupsItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewGroupsItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Events provides operations to manage the events property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Events()(*GroupsItemEventsRequestBuilder) {
    return NewGroupsItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById provides operations to manage the events property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) EventsById(id string)(*GroupsItemEventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewGroupsItemEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Extensions()(*GroupsItemExtensionsRequestBuilder) {
    return NewGroupsItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) ExtensionsById(id string)(*GroupsItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewGroupsItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get the properties and relationships of a group object. This operation returns by default only a subset of all the available properties, as noted in the Properties section. To get properties that are _not_ returned by default, specify them in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query.
func (m *GroupsGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsGroupItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *GroupsGroupItemRequestBuilder) GetMemberGroups()(*GroupsItemGetMemberGroupsRequestBuilder) {
    return NewGroupsItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *GroupsGroupItemRequestBuilder) GetMemberObjects()(*GroupsItemGetMemberObjectsRequestBuilder) {
    return NewGroupsItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePolicies provides operations to manage the groupLifecyclePolicies property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) GroupLifecyclePolicies()(*GroupsItemGroupLifecyclePoliciesRequestBuilder) {
    return NewGroupsItemGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePoliciesById provides operations to manage the groupLifecyclePolicies property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) GroupLifecyclePoliciesById(id string)(*GroupsItemGroupLifecyclePoliciesGroupLifecyclePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy%2Did"] = id
    }
    return NewGroupsItemGroupLifecyclePoliciesGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) MemberOf()(*GroupsItemMemberOfRequestBuilder) {
    return NewGroupsItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) MemberOfById(id string)(*GroupsItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewGroupsItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Members()(*GroupsItemMembersRequestBuilder) {
    return NewGroupsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go.groups.item.members.item collection
func (m *GroupsGroupItemRequestBuilder) MembersById(id string)(*GroupsItemMembersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewGroupsItemMembersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MembersWithLicenseErrors provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) MembersWithLicenseErrors()(*GroupsItemMembersWithLicenseErrorsRequestBuilder) {
    return NewGroupsItemMembersWithLicenseErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersWithLicenseErrorsById provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) MembersWithLicenseErrorsById(id string)(*GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewGroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Onenote()(*GroupsItemOnenoteRequestBuilder) {
    return NewGroupsItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Owners provides operations to manage the owners property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Owners()(*GroupsItemOwnersRequestBuilder) {
    return NewGroupsItemOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go.groups.item.owners.item collection
func (m *GroupsGroupItemRequestBuilder) OwnersById(id string)(*GroupsItemOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewGroupsItemOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a group object.
func (m *GroupsGroupItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, requestConfiguration *GroupsGroupItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable), nil
}
// PermissionGrants provides operations to manage the permissionGrants property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) PermissionGrants()(*GroupsItemPermissionGrantsRequestBuilder) {
    return NewGroupsItemPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById provides operations to manage the permissionGrants property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) PermissionGrantsById(id string)(*GroupsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return NewGroupsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo provides operations to manage the photo property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Photo()(*GroupsItemPhotoRequestBuilder) {
    return NewGroupsItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos provides operations to manage the photos property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Photos()(*GroupsItemPhotosRequestBuilder) {
    return NewGroupsItemPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById provides operations to manage the photos property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) PhotosById(id string)(*GroupsItemPhotosProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return NewGroupsItemPhotosProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner provides operations to manage the planner property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Planner()(*GroupsItemPlannerRequestBuilder) {
    return NewGroupsItemPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSenders provides operations to manage the rejectedSenders property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) RejectedSenders()(*GroupsItemRejectedSendersRequestBuilder) {
    return NewGroupsItemRejectedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSendersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go.groups.item.rejectedSenders.item collection
func (m *GroupsGroupItemRequestBuilder) RejectedSendersById(id string)(*GroupsItemRejectedSendersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewGroupsItemRejectedSendersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RemoveFavorite provides operations to call the removeFavorite method.
func (m *GroupsGroupItemRequestBuilder) RemoveFavorite()(*GroupsItemRemoveFavoriteRequestBuilder) {
    return NewGroupsItemRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew provides operations to call the renew method.
func (m *GroupsGroupItemRequestBuilder) Renew()(*GroupsItemRenewRequestBuilder) {
    return NewGroupsItemRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount provides operations to call the resetUnseenCount method.
func (m *GroupsGroupItemRequestBuilder) ResetUnseenCount()(*GroupsItemResetUnseenCountRequestBuilder) {
    return NewGroupsItemResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *GroupsGroupItemRequestBuilder) Restore()(*GroupsItemRestoreRequestBuilder) {
    return NewGroupsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Settings()(*GroupsItemSettingsRequestBuilder) {
    return NewGroupsItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) SettingsById(id string)(*GroupsItemSettingsGroupSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupSetting%2Did"] = id
    }
    return NewGroupsItemSettingsGroupSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the sites property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Sites()(*GroupsItemSitesRequestBuilder) {
    return NewGroupsItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the sites property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) SitesById(id string)(*GroupsItemSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return NewGroupsItemSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SubscribeByMail provides operations to call the subscribeByMail method.
func (m *GroupsGroupItemRequestBuilder) SubscribeByMail()(*GroupsItemSubscribeByMailRequestBuilder) {
    return NewGroupsItemSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Team provides operations to manage the team property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Team()(*GroupsItemTeamRequestBuilder) {
    return NewGroupsItemTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Threads provides operations to manage the threads property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) Threads()(*GroupsItemThreadsRequestBuilder) {
    return NewGroupsItemThreadsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreadsById provides operations to manage the threads property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) ThreadsById(id string)(*GroupsItemThreadsConversationThreadItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationThread%2Did"] = id
    }
    return NewGroupsItemThreadsConversationThreadItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) TransitiveMemberOf()(*GroupsItemTransitiveMemberOfRequestBuilder) {
    return NewGroupsItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) TransitiveMemberOfById(id string)(*GroupsItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewGroupsItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMembers provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) TransitiveMembers()(*GroupsItemTransitiveMembersRequestBuilder) {
    return NewGroupsItemTransitiveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMembersById provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
func (m *GroupsGroupItemRequestBuilder) TransitiveMembersById(id string)(*GroupsItemTransitiveMembersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewGroupsItemTransitiveMembersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnsubscribeByMail provides operations to call the unsubscribeByMail method.
func (m *GroupsGroupItemRequestBuilder) UnsubscribeByMail()(*GroupsItemUnsubscribeByMailRequestBuilder) {
    return NewGroupsItemUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties provides operations to call the validateProperties method.
func (m *GroupsGroupItemRequestBuilder) ValidateProperties()(*GroupsItemValidatePropertiesRequestBuilder) {
    return NewGroupsItemValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
