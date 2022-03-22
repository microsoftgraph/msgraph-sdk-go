package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationProtection provides operations to manage the informationProtection singleton.
type InformationProtection struct {
    Entity
    // 
    bitlocker Bitlockerable;
    // 
    threatAssessmentRequests []ThreatAssessmentRequestable;
}
// NewInformationProtection instantiates a new informationProtection and sets the default values.
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInformationProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInformationProtection(), nil
}
// GetBitlocker gets the bitlocker property value. 
func (m *InformationProtection) GetBitlocker()(Bitlockerable) {
    if m == nil {
        return nil
    } else {
        return m.bitlocker
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bitlocker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitlockerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitlocker(val.(Bitlockerable))
        }
        return nil
    }
    res["threatAssessmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateThreatAssessmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThreatAssessmentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(ThreatAssessmentRequestable)
            }
            m.SetThreatAssessmentRequests(res)
        }
        return nil
    }
    return res
}
// GetThreatAssessmentRequests gets the threatAssessmentRequests property value. 
func (m *InformationProtection) GetThreatAssessmentRequests()([]ThreatAssessmentRequestable) {
    if m == nil {
        return nil
    } else {
        return m.threatAssessmentRequests
    }
}
// Serialize serializes information the current object
func (m *InformationProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("bitlocker", m.GetBitlocker())
        if err != nil {
            return err
        }
    }
    if m.GetThreatAssessmentRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetThreatAssessmentRequests()))
        for i, v := range m.GetThreatAssessmentRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("threatAssessmentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBitlocker sets the bitlocker property value. 
func (m *InformationProtection) SetBitlocker(value Bitlockerable)() {
    if m != nil {
        m.bitlocker = value
    }
}
// SetThreatAssessmentRequests sets the threatAssessmentRequests property value. 
func (m *InformationProtection) SetThreatAssessmentRequests(value []ThreatAssessmentRequestable)() {
    if m != nil {
        m.threatAssessmentRequests = value
    }
}
