package checkmembergroups

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CheckMemberGroupsRequestBody struct {
    additionalData map[string]interface{};
    groupIds []string;
}
func NewCheckMemberGroupsRequestBody()(*CheckMemberGroupsRequestBody) {
    m := &CheckMemberGroupsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CheckMemberGroupsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CheckMemberGroupsRequestBody) GetGroupIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupIds
    }
}
func (m *CheckMemberGroupsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["groupIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetGroupIds(res)
        return nil
    }
    return res
}
func (m *CheckMemberGroupsRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *CheckMemberGroupsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("groupIds", m.GetGroupIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *CheckMemberGroupsRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CheckMemberGroupsRequestBody) SetGroupIds(value []string)() {
    m.groupIds = value
}
