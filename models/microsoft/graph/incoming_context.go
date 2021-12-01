package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IncomingContext 
type IncomingContext struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The ID of the participant that is under observation. Read-only.
    observedParticipantId *string;
    // The identity that the call is happening on behalf of.
    onBehalfOf *IdentitySet;
    // The ID of the participant that triggered the incoming call. Read-only.
    sourceParticipantId *string;
    // The identity that transferred the call.
    transferor *IdentitySet;
}
// NewIncomingContext instantiates a new incomingContext and sets the default values.
func NewIncomingContext()(*IncomingContext) {
    m := &IncomingContext{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IncomingContext) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetObservedParticipantId gets the observedParticipantId property value. The ID of the participant that is under observation. Read-only.
func (m *IncomingContext) GetObservedParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.observedParticipantId
    }
}
// GetOnBehalfOf gets the onBehalfOf property value. The identity that the call is happening on behalf of.
func (m *IncomingContext) GetOnBehalfOf()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.onBehalfOf
    }
}
// GetSourceParticipantId gets the sourceParticipantId property value. The ID of the participant that triggered the incoming call. Read-only.
func (m *IncomingContext) GetSourceParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceParticipantId
    }
}
// GetTransferor gets the transferor property value. The identity that transferred the call.
func (m *IncomingContext) GetTransferor()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.transferor
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IncomingContext) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["observedParticipantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObservedParticipantId(val)
        }
        return nil
    }
    res["onBehalfOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnBehalfOf(val.(*IdentitySet))
        }
        return nil
    }
    res["sourceParticipantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceParticipantId(val)
        }
        return nil
    }
    res["transferor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferor(val.(*IdentitySet))
        }
        return nil
    }
    return res
}
func (m *IncomingContext) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IncomingContext) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetObservedParticipantId sets the observedParticipantId property value. The ID of the participant that is under observation. Read-only.
func (m *IncomingContext) SetObservedParticipantId(value *string)() {
    if m != nil {
        m.observedParticipantId = value
    }
}
// SetOnBehalfOf sets the onBehalfOf property value. The identity that the call is happening on behalf of.
func (m *IncomingContext) SetOnBehalfOf(value *IdentitySet)() {
    if m != nil {
        m.onBehalfOf = value
    }
}
// SetSourceParticipantId sets the sourceParticipantId property value. The ID of the participant that triggered the incoming call. Read-only.
func (m *IncomingContext) SetSourceParticipantId(value *string)() {
    if m != nil {
        m.sourceParticipantId = value
    }
}
// SetTransferor sets the transferor property value. The identity that transferred the call.
func (m *IncomingContext) SetTransferor(value *IdentitySet)() {
    if m != nil {
        m.transferor = value
    }
}
