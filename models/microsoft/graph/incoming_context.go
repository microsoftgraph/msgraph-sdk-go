package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IncomingContext struct {
    additionalData map[string]interface{};
    observedParticipantId *string;
    onBehalfOf *IdentitySet;
    sourceParticipantId *string;
    transferor *IdentitySet;
}
func NewIncomingContext()(*IncomingContext) {
    m := &IncomingContext{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *IncomingContext) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *IncomingContext) GetObservedParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.observedParticipantId
    }
}
func (m *IncomingContext) GetOnBehalfOf()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.onBehalfOf
    }
}
func (m *IncomingContext) GetSourceParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceParticipantId
    }
}
func (m *IncomingContext) GetTransferor()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.transferor
    }
}
func (m *IncomingContext) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["observedParticipantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetObservedParticipantId(val)
        return nil
    }
    res["onBehalfOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetOnBehalfOf(val.(*IdentitySet))
        return nil
    }
    res["sourceParticipantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceParticipantId(val)
        return nil
    }
    res["transferor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetTransferor(val.(*IdentitySet))
        return nil
    }
    return res
}
func (m *IncomingContext) IsNil()(bool) {
    return m == nil
}
func (m *IncomingContext) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("observedParticipantId", m.GetObservedParticipantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("onBehalfOf", m.GetOnBehalfOf())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceParticipantId", m.GetSourceParticipantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("transferor", m.GetTransferor())
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
func (m *IncomingContext) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *IncomingContext) SetObservedParticipantId(value *string)() {
    m.observedParticipantId = value
}
func (m *IncomingContext) SetOnBehalfOf(value *IdentitySet)() {
    m.onBehalfOf = value
}
func (m *IncomingContext) SetSourceParticipantId(value *string)() {
    m.sourceParticipantId = value
}
func (m *IncomingContext) SetTransferor(value *IdentitySet)() {
    m.transferor = value
}
