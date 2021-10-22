package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i28a830c1e815101c0b85504c7b909848fd0dc929f2fa98ecbf66ae5908cc0e7e "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/copy"
    i57ff9c1da28e32d1438213a00ffd1d855134d619642894e7fd87e4532027a864 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/multivalueextendedproperties"
    i7f0870b87ccb71db8ecb5bfea28fb9e54778c808e443d1a99c4580eea1bd3f52 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/move"
    i8f7120f978f78e56c6efaa74e7732bcb02d100ea8df97bff99e5f8e11b74afc0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages"
    ic134c9b00254f702eada75fd8ad921541915d74fcca08820a255eae64692f139 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/singlevalueextendedproperties"
    id31d137ab3a6ab25d27f953a9a06af2a8a3fec04189a4cb157e4d5cb6cdaebf4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messagerules"
    if9a47801299e0be05f05c57b8b33e6c23c8bf48add445922ff4931f33ac0a26e "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders"
    i1285f2feb19f3fd1b1b5a3d0dc5440e58dafd55c4f69daede162ea2180f7adf5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item"
    i335a0ecd907517c064bd273cae9902fe9447a3a8ec23fe4de8b215b5356fc26c "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/multivalueextendedproperties/item"
    i9c44f0ee1a0a3f651db0ddfbaf98ec44de5c06f226dce2ae35c6c384f3a2bad4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item"
    ic153e3fa3efccd7712fa90e3c5117dc286ddf11d28ccac6a7570f274e8f78032 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/singlevalueextendedproperties/item"
    if56e053996654e0ad839a540d45995ed16925363ccece02b277bf8d9fe87dd53 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messagerules/item"
)

type MailFolderRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type MailFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *MailFolderRequestBuilder) ChildFolders()(*if9a47801299e0be05f05c57b8b33e6c23c8bf48add445922ff4931f33ac0a26e.ChildFoldersRequestBuilder) {
    return if9a47801299e0be05f05c57b8b33e6c23c8bf48add445922ff4931f33ac0a26e.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) ChildFoldersById(id string)(*i1285f2feb19f3fd1b1b5a3d0dc5440e58dafd55c4f69daede162ea2180f7adf5.MailFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["mailFolder_id1"] = id
    }
    return i1285f2feb19f3fd1b1b5a3d0dc5440e58dafd55c4f69daede162ea2180f7adf5.NewMailFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewMailFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderRequestBuilder) {
    m := &MailFolderRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/mailFolders/{mailFolder_id}{?select,expand}";
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
func NewMailFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFolderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MailFolderRequestBuilder) Copy()(*i28a830c1e815101c0b85504c7b909848fd0dc929f2fa98ecbf66ae5908cc0e7e.CopyRequestBuilder) {
    return i28a830c1e815101c0b85504c7b909848fd0dc929f2fa98ecbf66ae5908cc0e7e.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MailFolderRequestBuilder) CreateGetRequestInformation(q func (value *MailFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(MailFolderRequestBuilderGetQueryParameters)
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
func (m *MailFolderRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MailFolderRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *MailFolderRequestBuilder) Get(q func (value *MailFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMailFolder() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolder), nil
}
func (m *MailFolderRequestBuilder) MessageRules()(*id31d137ab3a6ab25d27f953a9a06af2a8a3fec04189a4cb157e4d5cb6cdaebf4.MessageRulesRequestBuilder) {
    return id31d137ab3a6ab25d27f953a9a06af2a8a3fec04189a4cb157e4d5cb6cdaebf4.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MessageRulesById(id string)(*if56e053996654e0ad839a540d45995ed16925363ccece02b277bf8d9fe87dd53.MessageRuleRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["messageRule_id"] = id
    }
    return if56e053996654e0ad839a540d45995ed16925363ccece02b277bf8d9fe87dd53.NewMessageRuleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) Messages()(*i8f7120f978f78e56c6efaa74e7732bcb02d100ea8df97bff99e5f8e11b74afc0.MessagesRequestBuilder) {
    return i8f7120f978f78e56c6efaa74e7732bcb02d100ea8df97bff99e5f8e11b74afc0.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MessagesById(id string)(*i9c44f0ee1a0a3f651db0ddfbaf98ec44de5c06f226dce2ae35c6c384f3a2bad4.MessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return i9c44f0ee1a0a3f651db0ddfbaf98ec44de5c06f226dce2ae35c6c384f3a2bad4.NewMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) Move()(*i7f0870b87ccb71db8ecb5bfea28fb9e54778c808e443d1a99c4580eea1bd3f52.MoveRequestBuilder) {
    return i7f0870b87ccb71db8ecb5bfea28fb9e54778c808e443d1a99c4580eea1bd3f52.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MultiValueExtendedProperties()(*i57ff9c1da28e32d1438213a00ffd1d855134d619642894e7fd87e4532027a864.MultiValueExtendedPropertiesRequestBuilder) {
    return i57ff9c1da28e32d1438213a00ffd1d855134d619642894e7fd87e4532027a864.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i335a0ecd907517c064bd273cae9902fe9447a3a8ec23fe4de8b215b5356fc26c.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i335a0ecd907517c064bd273cae9902fe9447a3a8ec23fe4de8b215b5356fc26c.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *MailFolderRequestBuilder) SingleValueExtendedProperties()(*ic134c9b00254f702eada75fd8ad921541915d74fcca08820a255eae64692f139.SingleValueExtendedPropertiesRequestBuilder) {
    return ic134c9b00254f702eada75fd8ad921541915d74fcca08820a255eae64692f139.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic153e3fa3efccd7712fa90e3c5117dc286ddf11d28ccac6a7570f274e8f78032.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ic153e3fa3efccd7712fa90e3c5117dc286ddf11d28ccac6a7570f274e8f78032.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
