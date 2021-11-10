package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type InformationProtection struct {
    Entity
    // 
    bitlocker *Bitlocker;
    // 
    threatAssessmentRequests []ThreatAssessmentRequest;
}
// Instantiates a new informationProtection and sets the default values.
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the bitlocker property value. 
func (m *InformationProtection) GetBitlocker()(*Bitlocker) {
    if m == nil {
        return nil
    } else {
        return m.bitlocker
    }
}
// Gets the threatAssessmentRequests property value. 
func (m *InformationProtection) GetThreatAssessmentRequests()([]ThreatAssessmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.threatAssessmentRequests
    }
}
// The deserialization information for the current model
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bitlocker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBitlocker() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitlocker(val.(*Bitlocker))
        }
        return nil
    }
    res["threatAssessmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThreatAssessmentRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThreatAssessmentRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*ThreatAssessmentRequest))
            }
            m.SetThreatAssessmentRequests(res)
        }
        return nil
    }
    return res
}
func (m *InformationProtection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetThreatAssessmentRequests()))
        for i, v := range m.GetThreatAssessmentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("threatAssessmentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the bitlocker property value. 
// Parameters:
//  - value : Value to set for the bitlocker property.
func (m *InformationProtection) SetBitlocker(value *Bitlocker)() {
    m.bitlocker = value
}
// Sets the threatAssessmentRequests property value. 
// Parameters:
//  - value : Value to set for the threatAssessmentRequests property.
func (m *InformationProtection) SetThreatAssessmentRequests(value []ThreatAssessmentRequest)() {
    m.threatAssessmentRequests = value
}
