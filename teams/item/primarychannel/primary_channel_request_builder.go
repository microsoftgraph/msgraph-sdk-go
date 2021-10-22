package primarychannel

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1e4013461c5123188ab7ccd33dd297c0d5a7865ba9fea8e6c138306141eeec4a "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/completemigration"
    i4785692a9d61d4a8638d06fe97e1daa4feb18f7b743e7ae34ff3f07f3d9425b4 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/messages"
    i977f2dd26da5a2e1880df87bd5372aa7a19a14868a426678151d974f360c2a01 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/members"
    ib05faad6285a5d11c891d13f0450d43dfbe0ddb57be2fa94177f219559ef0527 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/tabs"
    id169406704fae3a188bcd329b3cbae975ed48842ae13c31a260f221158f8516a "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/filesfolder"
    i0193f7b86687af396beaaecb295b77e66ded553cf49127c0d06bc79b38a4d93f "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/tabs/item"
    ib00fe6f5242d43b4e0cf713930bcfb1828c83941d27088c19568b60fa6e113d8 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/messages/item"
    ie734300e5bedd84acd036f8be309a4094f66ee9a58c5ec5b270079e8efc64674 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/primarychannel/members/item"
)

type PrimaryChannelRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PrimaryChannelRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*i1e4013461c5123188ab7ccd33dd297c0d5a7865ba9fea8e6c138306141eeec4a.CompleteMigrationRequestBuilder) {
    return i1e4013461c5123188ab7ccd33dd297c0d5a7865ba9fea8e6c138306141eeec4a.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/teams/{team_id}/primaryChannel{?select,expand}";
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
func NewPrimaryChannelRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformation(q func (value *PrimaryChannelRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PrimaryChannelRequestBuilderGetQueryParameters)
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
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrimaryChannelRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*id169406704fae3a188bcd329b3cbae975ed48842ae13c31a260f221158f8516a.FilesFolderRequestBuilder) {
    return id169406704fae3a188bcd329b3cbae975ed48842ae13c31a260f221158f8516a.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Get(q func (value *PrimaryChannelRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewChannel() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel), nil
}
func (m *PrimaryChannelRequestBuilder) Members()(*i977f2dd26da5a2e1880df87bd5372aa7a19a14868a426678151d974f360c2a01.MembersRequestBuilder) {
    return i977f2dd26da5a2e1880df87bd5372aa7a19a14868a426678151d974f360c2a01.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*ie734300e5bedd84acd036f8be309a4094f66ee9a58c5ec5b270079e8efc64674.ConversationMemberRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ie734300e5bedd84acd036f8be309a4094f66ee9a58c5ec5b270079e8efc64674.NewConversationMemberRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Messages()(*i4785692a9d61d4a8638d06fe97e1daa4feb18f7b743e7ae34ff3f07f3d9425b4.MessagesRequestBuilder) {
    return i4785692a9d61d4a8638d06fe97e1daa4feb18f7b743e7ae34ff3f07f3d9425b4.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*ib00fe6f5242d43b4e0cf713930bcfb1828c83941d27088c19568b60fa6e113d8.ChatMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return ib00fe6f5242d43b4e0cf713930bcfb1828c83941d27088c19568b60fa6e113d8.NewChatMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrimaryChannelRequestBuilder) Tabs()(*ib05faad6285a5d11c891d13f0450d43dfbe0ddb57be2fa94177f219559ef0527.TabsRequestBuilder) {
    return ib05faad6285a5d11c891d13f0450d43dfbe0ddb57be2fa94177f219559ef0527.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*i0193f7b86687af396beaaecb295b77e66ded553cf49127c0d06bc79b38a4d93f.TeamsTabRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i0193f7b86687af396beaaecb295b77e66ded553cf49127c0d06bc79b38a4d93f.NewTeamsTabRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
