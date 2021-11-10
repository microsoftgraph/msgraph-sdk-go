package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new incomingContext and sets the default values.
func NewIncomingContext()(*IncomingContext) {
    m := &IncomingContext{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IncomingContext) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the observedParticipantId property value. The ID of the participant that is under observation. Read-only.
func (m *IncomingContext) GetObservedParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.observedParticipantId
    }
}
// Gets the onBehalfOf property value. The identity that the call is happening on behalf of.
func (m *IncomingContext) GetOnBehalfOf()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.onBehalfOf
    }
}
// Gets the sourceParticipantId property value. The ID of the participant that triggered the incoming call. Read-only.
func (m *IncomingContext) GetSourceParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceParticipantId
    }
}
// Gets the transferor property value. The identity that transferred the call.
func (m *IncomingContext) GetTransferor()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.transferor
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *IncomingContext) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the observedParticipantId property value. The ID of the participant that is under observation. Read-only.
// Parameters:
//  - value : Value to set for the observedParticipantId property.
func (m *IncomingContext) SetObservedParticipantId(value *string)() {
    m.observedParticipantId = value
}
// Sets the onBehalfOf property value. The identity that the call is happening on behalf of.
// Parameters:
//  - value : Value to set for the onBehalfOf property.
func (m *IncomingContext) SetOnBehalfOf(value *IdentitySet)() {
    m.onBehalfOf = value
}
// Sets the sourceParticipantId property value. The ID of the participant that triggered the incoming call. Read-only.
// Parameters:
//  - value : Value to set for the sourceParticipantId property.
func (m *IncomingContext) SetSourceParticipantId(value *string)() {
    m.sourceParticipantId = value
}
// Sets the transferor property value. The identity that transferred the call.
// Parameters:
//  - value : Value to set for the transferor property.
func (m *IncomingContext) SetTransferor(value *IdentitySet)() {
    m.transferor = value
}
