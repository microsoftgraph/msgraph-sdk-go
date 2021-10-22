package archive

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ArchiveRequestBody struct {
    additionalData map[string]interface{};
    shouldSetSpoSiteReadOnlyForMembers *bool;
}
func NewArchiveRequestBody()(*ArchiveRequestBody) {
    m := &ArchiveRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ArchiveRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ArchiveRequestBody) GetShouldSetSpoSiteReadOnlyForMembers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shouldSetSpoSiteReadOnlyForMembers
    }
}
func (m *ArchiveRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["shouldSetSpoSiteReadOnlyForMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShouldSetSpoSiteReadOnlyForMembers(val)
        return nil
    }
    return res
}
func (m *ArchiveRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ArchiveRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("shouldSetSpoSiteReadOnlyForMembers", m.GetShouldSetSpoSiteReadOnlyForMembers())
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
func (m *ArchiveRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ArchiveRequestBody) SetShouldSetSpoSiteReadOnlyForMembers(value *bool)() {
    m.shouldSetSpoSiteReadOnlyForMembers = value
}
