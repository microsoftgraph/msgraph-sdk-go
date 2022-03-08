package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Bitlocker provides operations to manage the informationProtection singleton.
type Bitlocker struct {
    Entity
    // The recovery keys associated with the bitlocker entity.
    recoveryKeys []BitlockerRecoveryKeyable;
}
// NewBitlocker instantiates a new bitlocker and sets the default values.
func NewBitlocker()(*Bitlocker) {
    m := &Bitlocker{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBitlockerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBitlockerFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBitlocker(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Bitlocker) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["recoveryKeys"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBitlockerRecoveryKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BitlockerRecoveryKeyable, len(val))
            for i, v := range val {
                res[i] = v.(BitlockerRecoveryKeyable)
            }
            m.SetRecoveryKeys(res)
        }
        return nil
    }
    return res
}
// GetRecoveryKeys gets the recoveryKeys property value. The recovery keys associated with the bitlocker entity.
func (m *Bitlocker) GetRecoveryKeys()([]BitlockerRecoveryKeyable) {
    if m == nil {
        return nil
    } else {
        return m.recoveryKeys
    }
}
func (m *Bitlocker) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Bitlocker) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRecoveryKeys() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecoveryKeys()))
        for i, v := range m.GetRecoveryKeys() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("recoveryKeys", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRecoveryKeys sets the recoveryKeys property value. The recovery keys associated with the bitlocker entity.
func (m *Bitlocker) SetRecoveryKeys(value []BitlockerRecoveryKeyable)() {
    if m != nil {
        m.recoveryKeys = value
    }
}
