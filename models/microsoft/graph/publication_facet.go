package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PublicationFacet struct {
    additionalData map[string]interface{};
    level *string;
    versionId *string;
}
func NewPublicationFacet()(*PublicationFacet) {
    m := &PublicationFacet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PublicationFacet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PublicationFacet) GetLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.level
    }
}
func (m *PublicationFacet) GetVersionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionId
    }
}
func (m *PublicationFacet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["level"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLevel(val)
        return nil
    }
    res["versionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersionId(val)
        return nil
    }
    return res
}
func (m *PublicationFacet) IsNil()(bool) {
    return m == nil
}
func (m *PublicationFacet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("level", m.GetLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("versionId", m.GetVersionId())
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
func (m *PublicationFacet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PublicationFacet) SetLevel(value *string)() {
    m.level = value
}
func (m *PublicationFacet) SetVersionId(value *string)() {
    m.versionId = value
}
