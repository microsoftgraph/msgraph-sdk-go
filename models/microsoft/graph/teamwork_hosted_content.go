package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamworkHostedContent struct {
    Entity
    contentBytes []byte;
    contentType *string;
}
func NewTeamworkHostedContent()(*TeamworkHostedContent) {
    m := &TeamworkHostedContent{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TeamworkHostedContent) GetContentBytes()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.contentBytes
    }
}
func (m *TeamworkHostedContent) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
func (m *TeamworkHostedContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetContentBytes(val)
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentType(val)
        return nil
    }
    return res
}
func (m *TeamworkHostedContent) IsNil()(bool) {
    return m == nil
}
func (m *TeamworkHostedContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("contentBytes", m.GetContentBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TeamworkHostedContent) SetContentBytes(value []byte)() {
    m.contentBytes = value
}
func (m *TeamworkHostedContent) SetContentType(value *string)() {
    m.contentType = value
}
