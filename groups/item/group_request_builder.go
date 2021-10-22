package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i040f92e34d878c9b1e2e7cdc0c6762e73b19a95fd3cad9cb82158adea802fc6e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/permissiongrants"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i126eb8e350bd8db3c0241a5ead64e6099dcc4e5b5e482c99faef02474ed9032d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites"
    i1e3ffd7e01fa4360b91692ef154671f0e71e5d0cc35b0f7ab6cd93b21fe9191f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/checkmemberobjects"
    i25afeeded75ccad738da224aa35ea4c5075fa032bec86cab3c058ab924453175 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads"
    i27bbc458e055bd736262322f8b4a908d851ea90e6431282c71068bcf65cbefcb "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof"
    i31054151b0bb87829bfe3963474fbe4ab17c82c71e6f7470f98dc872d3010e54 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview"
    i324ff89ab47fde0163709057863e71e77c5c94057ba407c20ec187dff1506fa0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members"
    i3c92c4acc0e222acc40b055f3d6ea5d09e6aa97bc4c5d7146f969c4a0c9a52e5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/unsubscribebymail"
    i3cfdfc8f36ef3c22284117091b5b0e21c71d2208e093946aa4f4e531a31d6aed "github.com/microsoftgraph/msgraph-sdk-go/groups/item/assignlicense"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4b72743ccd3ea3919530a55bb3c78d6aede553a9442ca26e3b145c91ba7b6d90 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations"
    i4c27f6ecf21cb0236a018e5fdf6e4c5a14a02dfb90e5d46f78219b25b861f9d4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/createdonbehalfof"
    i4f15e812f46be3720cebce4ecae6d3639575a9930ab35c87584933a4effb9666 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/removefavorite"
    i5336b126a77740a88a03897e9414f4cf69114c0dcff9a4bf0edfd7bc60607ed2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives"
    i543d4232a6f65a85c70d6cd7e6fe9118de280565c7f04d0b7b8ae168b8498aff "github.com/microsoftgraph/msgraph-sdk-go/groups/item/renew"
    i5b3cd370da540ada436e22f24ba1fd3b79f4345a6a464bafb2d3ac2e3f6c5f45 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/getmemberobjects"
    i5bdf25c7d741b0f149a0a9a37b84895b651d314b176e9ad86423f2fd3f6f1344 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/subscribebymail"
    i678de7f91184830396b1d4b7a46f78a5b66c413205221cc3cbe0d72da6209e85 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/photos"
    i68841ec9642599fe7197e402214c1f2ff1a9b84a5f444d4a272cbb8cf86ab097 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/checkmembergroups"
    i68aecc9a3d9efbc7cd11ecd666cdf3c9ebeda118b2842b144598d9defce5df18 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar"
    i6d7d47498bacfe34bb5923be815a017a51887291335f05da4e892f4e6faa0839 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/checkgrantedpermissionsforapp"
    i7235fca48420218e9a89986cfbc32ca5249ee8e87a7ab7e1b2e2cf0c6a3ef379 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/rejectedsenders"
    i7fccd84783e04ba49b48ffd1b4c8bf347fd854264a2c69005507b8427ac3b5ff "github.com/microsoftgraph/msgraph-sdk-go/groups/item/photo"
    i8142ad66b1f81d850ea82ba7c2d1b04b07dcd7138cfeb3e4cbbe39613e799ab0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team"
    i8e118fccf5aaa007f65dd345e24db21251fa3341071cca2ac34246fe0ca92cce "github.com/microsoftgraph/msgraph-sdk-go/groups/item/getmembergroups"
    i90733061bf48a2adb8efb9e149c1ac470817d19d24a29a4b84d124e580d38d99 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/settings"
    ia1bbdd3513a9d0a19bfbb53eba37a0c87dccf5a8784ecf2f1c48a9561f7860c0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/validateproperties"
    ia37dbe7af78a635529ec6aea76874b613f13529b238aac201827e41079c76528 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote"
    ia4d9b5e6814f0473a8235d1fcd1b0fba509840c5da363927bf12356141119b6b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/acceptedsenders"
    ia5f63161806dfdb2925614c84d36c55c4f0585bd40f214192b5b8560e1f241fb "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers"
    ib032192d61dc95db87cbde93dea6eb88484ee1160f23230aefa8adf11ac2716a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/grouplifecyclepolicies"
    ib125094fb6ca7398f85d34937572217d9648c0edcd9d750686948808cd9d7d45 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/resetunseencount"
    ib691e0fc9b9df0096773889a7b02343b3cb3fd38a786840aa53556c739e17e5b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof"
    ibe1b4bf6d56b91362a6c58dbff13888d36435ca678122388e0d2dcac9a32f565 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/extensions"
    ic622ab4965d64f3a7c740b58a10c44b0c8bc9a9837ee71eec701baad232718c8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors"
    ic8d68eb69bcf15c0a5bb6ae908d3c3e6bdb9fd8ab08e2176d9426e308fc9988a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/addfavorite"
    id75d43ba05b6bba0a51c61bc3b7ed977a5478928d413c7e05c88b66deabd1d09 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners"
    idf81ca2a070233874101a1a641a5cc974955b4c7ec50cc33774f0a4ee559bb38 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drive"
    ie34da707ff79e10048ef9f7ea711aa50a3a792dcd3d5fe2f462b4b8e01c5fcd4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner"
    ie8a0001ab578d964916038792bdc705c9f8184a4fafa5ffd68261fa822964c4d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/approleassignments"
    ief4262afcc597cdf8d4a38908ccbf9e45055870a45322db3fc0fc1420c3f0ec6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events"
    iff10248aac2fee80d2d103959cd901c3513c9862626b600ab67f077b7cc69548 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/restore"
    i0364eb181a6c8e59a0dd7c8a13994a8d5b1df88f847ca1217f353d68aec7fb0c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/permissiongrants/item"
    i08f2a0edeb243e9cfaba694782db5c61d4ff36791058cfde168e2c9472490a8c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item"
    i1e824fbe5d664ceca8d12a43b50d7d7f49e8713e55656dd7def6bfe0df703627 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/approleassignments/item"
    i441b3030f4e4a48458260f546a7160d5d7f7a51cc67d40ce8267b45eb54c0876 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/grouplifecyclepolicies/item"
    i5fa4399b660c3896f8bf2358c9bf08baa829b8c8768c9f069b443692c2c9b1f3 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/settings/item"
    ia4ceaf879d63c036e207bfe238095ee32d3e41faae41517d448ab206c335f0cb "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item"
    iaa0f64945b540b5b0648f03319b424981c2ea47cf4a0113ee649e5f593a07253 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item"
    ic5de56bd9caa7009872a581bda3b0f3ff9e503e4dc9abc07f06521e46c0472ca "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item"
    ice1094b94e12d6484840204756982a958170f3c66b8047a166bf39c98320acc7 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item"
    id632a3f6a3ad5fcc8b0136d664e83c04052ada39b1f8b92e7bccf7c61fb49848 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/photos/item"
    idb19301160b95d9f7c5034e26232a3aa92b814911db3a14c43386bd9461c541c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item"
    iddc33c1a3316b503c055f30aa899183e7ab6a6cd014c230bef3258364feeaf7c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/extensions/item"
)

type GroupRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type GroupRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *GroupRequestBuilder) AcceptedSenders()(*ia4d9b5e6814f0473a8235d1fcd1b0fba509840c5da363927bf12356141119b6b.AcceptedSendersRequestBuilder) {
    return ia4d9b5e6814f0473a8235d1fcd1b0fba509840c5da363927bf12356141119b6b.NewAcceptedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) AddFavorite()(*ic8d68eb69bcf15c0a5bb6ae908d3c3e6bdb9fd8ab08e2176d9426e308fc9988a.AddFavoriteRequestBuilder) {
    return ic8d68eb69bcf15c0a5bb6ae908d3c3e6bdb9fd8ab08e2176d9426e308fc9988a.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) AppRoleAssignments()(*ie8a0001ab578d964916038792bdc705c9f8184a4fafa5ffd68261fa822964c4d.AppRoleAssignmentsRequestBuilder) {
    return ie8a0001ab578d964916038792bdc705c9f8184a4fafa5ffd68261fa822964c4d.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) AppRoleAssignmentsById(id string)(*i1e824fbe5d664ceca8d12a43b50d7d7f49e8713e55656dd7def6bfe0df703627.AppRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return i1e824fbe5d664ceca8d12a43b50d7d7f49e8713e55656dd7def6bfe0df703627.NewAppRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) AssignLicense()(*i3cfdfc8f36ef3c22284117091b5b0e21c71d2208e093946aa4f4e531a31d6aed.AssignLicenseRequestBuilder) {
    return i3cfdfc8f36ef3c22284117091b5b0e21c71d2208e093946aa4f4e531a31d6aed.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Calendar()(*i68aecc9a3d9efbc7cd11ecd666cdf3c9ebeda118b2842b144598d9defce5df18.CalendarRequestBuilder) {
    return i68aecc9a3d9efbc7cd11ecd666cdf3c9ebeda118b2842b144598d9defce5df18.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) CalendarView()(*i31054151b0bb87829bfe3963474fbe4ab17c82c71e6f7470f98dc872d3010e54.CalendarViewRequestBuilder) {
    return i31054151b0bb87829bfe3963474fbe4ab17c82c71e6f7470f98dc872d3010e54.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) CalendarViewById(id string)(*ic5de56bd9caa7009872a581bda3b0f3ff9e503e4dc9abc07f06521e46c0472ca.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return ic5de56bd9caa7009872a581bda3b0f3ff9e503e4dc9abc07f06521e46c0472ca.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i6d7d47498bacfe34bb5923be815a017a51887291335f05da4e892f4e6faa0839.CheckGrantedPermissionsForAppRequestBuilder) {
    return i6d7d47498bacfe34bb5923be815a017a51887291335f05da4e892f4e6faa0839.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) CheckMemberGroups()(*i68841ec9642599fe7197e402214c1f2ff1a9b84a5f444d4a272cbb8cf86ab097.CheckMemberGroupsRequestBuilder) {
    return i68841ec9642599fe7197e402214c1f2ff1a9b84a5f444d4a272cbb8cf86ab097.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) CheckMemberObjects()(*i1e3ffd7e01fa4360b91692ef154671f0e71e5d0cc35b0f7ab6cd93b21fe9191f.CheckMemberObjectsRequestBuilder) {
    return i1e3ffd7e01fa4360b91692ef154671f0e71e5d0cc35b0f7ab6cd93b21fe9191f.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GroupRequestBuilder) Conversations()(*i4b72743ccd3ea3919530a55bb3c78d6aede553a9442ca26e3b145c91ba7b6d90.ConversationsRequestBuilder) {
    return i4b72743ccd3ea3919530a55bb3c78d6aede553a9442ca26e3b145c91ba7b6d90.NewConversationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) ConversationsById(id string)(*ice1094b94e12d6484840204756982a958170f3c66b8047a166bf39c98320acc7.ConversationRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["conversation_id"] = id
    }
    return ice1094b94e12d6484840204756982a958170f3c66b8047a166bf39c98320acc7.NewConversationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupRequestBuilder) CreatedOnBehalfOf()(*i4c27f6ecf21cb0236a018e5fdf6e4c5a14a02dfb90e5d46f78219b25b861f9d4.CreatedOnBehalfOfRequestBuilder) {
    return i4c27f6ecf21cb0236a018e5fdf6e4c5a14a02dfb90e5d46f78219b25b861f9d4.NewCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) CreateGetRequestInformation(q func (value *GroupRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(GroupRequestBuilderGetQueryParameters)
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
func (m *GroupRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Group, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *GroupRequestBuilder) Drive()(*idf81ca2a070233874101a1a641a5cc974955b4c7ec50cc33774f0a4ee559bb38.DriveRequestBuilder) {
    return idf81ca2a070233874101a1a641a5cc974955b4c7ec50cc33774f0a4ee559bb38.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Drives()(*i5336b126a77740a88a03897e9414f4cf69114c0dcff9a4bf0edfd7bc60607ed2.DrivesRequestBuilder) {
    return i5336b126a77740a88a03897e9414f4cf69114c0dcff9a4bf0edfd7bc60607ed2.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) DrivesById(id string)(*idb19301160b95d9f7c5034e26232a3aa92b814911db3a14c43386bd9461c541c.DriveRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["drive_id"] = id
    }
    return idb19301160b95d9f7c5034e26232a3aa92b814911db3a14c43386bd9461c541c.NewDriveRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Events()(*ief4262afcc597cdf8d4a38908ccbf9e45055870a45322db3fc0fc1420c3f0ec6.EventsRequestBuilder) {
    return ief4262afcc597cdf8d4a38908ccbf9e45055870a45322db3fc0fc1420c3f0ec6.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) EventsById(id string)(*i08f2a0edeb243e9cfaba694782db5c61d4ff36791058cfde168e2c9472490a8c.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i08f2a0edeb243e9cfaba694782db5c61d4ff36791058cfde168e2c9472490a8c.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Extensions()(*ibe1b4bf6d56b91362a6c58dbff13888d36435ca678122388e0d2dcac9a32f565.ExtensionsRequestBuilder) {
    return ibe1b4bf6d56b91362a6c58dbff13888d36435ca678122388e0d2dcac9a32f565.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) ExtensionsById(id string)(*iddc33c1a3316b503c055f30aa899183e7ab6a6cd014c230bef3258364feeaf7c.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iddc33c1a3316b503c055f30aa899183e7ab6a6cd014c230bef3258364feeaf7c.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Get(q func (value *GroupRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Group, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewGroup() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Group), nil
}
func (m *GroupRequestBuilder) GetMemberGroups()(*i8e118fccf5aaa007f65dd345e24db21251fa3341071cca2ac34246fe0ca92cce.GetMemberGroupsRequestBuilder) {
    return i8e118fccf5aaa007f65dd345e24db21251fa3341071cca2ac34246fe0ca92cce.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) GetMemberObjects()(*i5b3cd370da540ada436e22f24ba1fd3b79f4345a6a464bafb2d3ac2e3f6c5f45.GetMemberObjectsRequestBuilder) {
    return i5b3cd370da540ada436e22f24ba1fd3b79f4345a6a464bafb2d3ac2e3f6c5f45.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) GroupLifecyclePolicies()(*ib032192d61dc95db87cbde93dea6eb88484ee1160f23230aefa8adf11ac2716a.GroupLifecyclePoliciesRequestBuilder) {
    return ib032192d61dc95db87cbde93dea6eb88484ee1160f23230aefa8adf11ac2716a.NewGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) GroupLifecyclePoliciesById(id string)(*i441b3030f4e4a48458260f546a7160d5d7f7a51cc67d40ce8267b45eb54c0876.GroupLifecyclePolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy_id"] = id
    }
    return i441b3030f4e4a48458260f546a7160d5d7f7a51cc67d40ce8267b45eb54c0876.NewGroupLifecyclePolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) MemberOf()(*ib691e0fc9b9df0096773889a7b02343b3cb3fd38a786840aa53556c739e17e5b.MemberOfRequestBuilder) {
    return ib691e0fc9b9df0096773889a7b02343b3cb3fd38a786840aa53556c739e17e5b.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Members()(*i324ff89ab47fde0163709057863e71e77c5c94057ba407c20ec187dff1506fa0.MembersRequestBuilder) {
    return i324ff89ab47fde0163709057863e71e77c5c94057ba407c20ec187dff1506fa0.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) MembersWithLicenseErrors()(*ic622ab4965d64f3a7c740b58a10c44b0c8bc9a9837ee71eec701baad232718c8.MembersWithLicenseErrorsRequestBuilder) {
    return ic622ab4965d64f3a7c740b58a10c44b0c8bc9a9837ee71eec701baad232718c8.NewMembersWithLicenseErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Onenote()(*ia37dbe7af78a635529ec6aea76874b613f13529b238aac201827e41079c76528.OnenoteRequestBuilder) {
    return ia37dbe7af78a635529ec6aea76874b613f13529b238aac201827e41079c76528.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Owners()(*id75d43ba05b6bba0a51c61bc3b7ed977a5478928d413c7e05c88b66deabd1d09.OwnersRequestBuilder) {
    return id75d43ba05b6bba0a51c61bc3b7ed977a5478928d413c7e05c88b66deabd1d09.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Group, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *GroupRequestBuilder) PermissionGrants()(*i040f92e34d878c9b1e2e7cdc0c6762e73b19a95fd3cad9cb82158adea802fc6e.PermissionGrantsRequestBuilder) {
    return i040f92e34d878c9b1e2e7cdc0c6762e73b19a95fd3cad9cb82158adea802fc6e.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) PermissionGrantsById(id string)(*i0364eb181a6c8e59a0dd7c8a13994a8d5b1df88f847ca1217f353d68aec7fb0c.ResourceSpecificPermissionGrantRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant_id"] = id
    }
    return i0364eb181a6c8e59a0dd7c8a13994a8d5b1df88f847ca1217f353d68aec7fb0c.NewResourceSpecificPermissionGrantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Photo()(*i7fccd84783e04ba49b48ffd1b4c8bf347fd854264a2c69005507b8427ac3b5ff.PhotoRequestBuilder) {
    return i7fccd84783e04ba49b48ffd1b4c8bf347fd854264a2c69005507b8427ac3b5ff.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Photos()(*i678de7f91184830396b1d4b7a46f78a5b66c413205221cc3cbe0d72da6209e85.PhotosRequestBuilder) {
    return i678de7f91184830396b1d4b7a46f78a5b66c413205221cc3cbe0d72da6209e85.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) PhotosById(id string)(*id632a3f6a3ad5fcc8b0136d664e83c04052ada39b1f8b92e7bccf7c61fb49848.ProfilePhotoRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["profilePhoto_id"] = id
    }
    return id632a3f6a3ad5fcc8b0136d664e83c04052ada39b1f8b92e7bccf7c61fb49848.NewProfilePhotoRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Planner()(*ie34da707ff79e10048ef9f7ea711aa50a3a792dcd3d5fe2f462b4b8e01c5fcd4.PlannerRequestBuilder) {
    return ie34da707ff79e10048ef9f7ea711aa50a3a792dcd3d5fe2f462b4b8e01c5fcd4.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) RejectedSenders()(*i7235fca48420218e9a89986cfbc32ca5249ee8e87a7ab7e1b2e2cf0c6a3ef379.RejectedSendersRequestBuilder) {
    return i7235fca48420218e9a89986cfbc32ca5249ee8e87a7ab7e1b2e2cf0c6a3ef379.NewRejectedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) RemoveFavorite()(*i4f15e812f46be3720cebce4ecae6d3639575a9930ab35c87584933a4effb9666.RemoveFavoriteRequestBuilder) {
    return i4f15e812f46be3720cebce4ecae6d3639575a9930ab35c87584933a4effb9666.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Renew()(*i543d4232a6f65a85c70d6cd7e6fe9118de280565c7f04d0b7b8ae168b8498aff.RenewRequestBuilder) {
    return i543d4232a6f65a85c70d6cd7e6fe9118de280565c7f04d0b7b8ae168b8498aff.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) ResetUnseenCount()(*ib125094fb6ca7398f85d34937572217d9648c0edcd9d750686948808cd9d7d45.ResetUnseenCountRequestBuilder) {
    return ib125094fb6ca7398f85d34937572217d9648c0edcd9d750686948808cd9d7d45.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Restore()(*iff10248aac2fee80d2d103959cd901c3513c9862626b600ab67f077b7cc69548.RestoreRequestBuilder) {
    return iff10248aac2fee80d2d103959cd901c3513c9862626b600ab67f077b7cc69548.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Settings()(*i90733061bf48a2adb8efb9e149c1ac470817d19d24a29a4b84d124e580d38d99.SettingsRequestBuilder) {
    return i90733061bf48a2adb8efb9e149c1ac470817d19d24a29a4b84d124e580d38d99.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) SettingsById(id string)(*i5fa4399b660c3896f8bf2358c9bf08baa829b8c8768c9f069b443692c2c9b1f3.GroupSettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["groupSetting_id"] = id
    }
    return i5fa4399b660c3896f8bf2358c9bf08baa829b8c8768c9f069b443692c2c9b1f3.NewGroupSettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Sites()(*i126eb8e350bd8db3c0241a5ead64e6099dcc4e5b5e482c99faef02474ed9032d.SitesRequestBuilder) {
    return i126eb8e350bd8db3c0241a5ead64e6099dcc4e5b5e482c99faef02474ed9032d.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) SitesById(id string)(*iaa0f64945b540b5b0648f03319b424981c2ea47cf4a0113ee649e5f593a07253.SiteRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["site_id"] = id
    }
    return iaa0f64945b540b5b0648f03319b424981c2ea47cf4a0113ee649e5f593a07253.NewSiteRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) SubscribeByMail()(*i5bdf25c7d741b0f149a0a9a37b84895b651d314b176e9ad86423f2fd3f6f1344.SubscribeByMailRequestBuilder) {
    return i5bdf25c7d741b0f149a0a9a37b84895b651d314b176e9ad86423f2fd3f6f1344.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Team()(*i8142ad66b1f81d850ea82ba7c2d1b04b07dcd7138cfeb3e4cbbe39613e799ab0.TeamRequestBuilder) {
    return i8142ad66b1f81d850ea82ba7c2d1b04b07dcd7138cfeb3e4cbbe39613e799ab0.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Threads()(*i25afeeded75ccad738da224aa35ea4c5075fa032bec86cab3c058ab924453175.ThreadsRequestBuilder) {
    return i25afeeded75ccad738da224aa35ea4c5075fa032bec86cab3c058ab924453175.NewThreadsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) ThreadsById(id string)(*ia4ceaf879d63c036e207bfe238095ee32d3e41faae41517d448ab206c335f0cb.ConversationThreadRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["conversationThread_id"] = id
    }
    return ia4ceaf879d63c036e207bfe238095ee32d3e41faae41517d448ab206c335f0cb.NewConversationThreadRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) TransitiveMemberOf()(*i27bbc458e055bd736262322f8b4a908d851ea90e6431282c71068bcf65cbefcb.TransitiveMemberOfRequestBuilder) {
    return i27bbc458e055bd736262322f8b4a908d851ea90e6431282c71068bcf65cbefcb.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) TransitiveMembers()(*ia5f63161806dfdb2925614c84d36c55c4f0585bd40f214192b5b8560e1f241fb.TransitiveMembersRequestBuilder) {
    return ia5f63161806dfdb2925614c84d36c55c4f0585bd40f214192b5b8560e1f241fb.NewTransitiveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i3c92c4acc0e222acc40b055f3d6ea5d09e6aa97bc4c5d7146f969c4a0c9a52e5.UnsubscribeByMailRequestBuilder) {
    return i3c92c4acc0e222acc40b055f3d6ea5d09e6aa97bc4c5d7146f969c4a0c9a52e5.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) ValidateProperties()(*ia1bbdd3513a9d0a19bfbb53eba37a0c87dccf5a8784ecf2f1c48a9561f7860c0.ValidatePropertiesRequestBuilder) {
    return ia1bbdd3513a9d0a19bfbb53eba37a0c87dccf5a8784ecf2f1c48a9561f7860c0.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
