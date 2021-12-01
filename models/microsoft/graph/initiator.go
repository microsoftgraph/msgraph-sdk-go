package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Initiator 
type Initiator struct {
    Identity
    // Type of initiator. Possible values are: user, application, system, unknownFutureValue.
    initiatorType *InitiatorType;
}
// NewInitiator instantiates a new initiator and sets the default values.
func NewInitiator()(*Initiator) {
    m := &Initiator{
        Identity: *NewIdentity(),
    }
    return m
}
// GetInitiatorType gets the initiatorType property value. Type of initiator. Possible values are: user, application, system, unknownFutureValue.
func (m *Initiator) GetInitiatorType()(*InitiatorType) {
    if m == nil {
        return nil
    } else {
        return m.initiatorType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Initiator) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["initiatorType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseInitiatorType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(InitiatorType)
            m.SetInitiatorType(&cast)
        }
        return nil
    }
    return res
}
func (m *Initiator) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Initiator) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInitiatorType() != nil {
        cast := m.GetInitiatorType().String()
        err = writer.WriteStringValue("initiatorType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiatorType sets the initiatorType property value. Type of initiator. Possible values are: user, application, system, unknownFutureValue.
func (m *Initiator) SetInitiatorType(value *InitiatorType)() {
    if m != nil {
        m.initiatorType = value
    }
}
