package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InformationProtection struct {
    Entity
    threatAssessmentRequests []ThreatAssessmentRequest;
}
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *NewEntity(),
    }
    return m
}
func (m *InformationProtection) GetThreatAssessmentRequests()([]ThreatAssessmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.threatAssessmentRequests
    }
}
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["threatAssessmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThreatAssessmentRequest() })
        if err != nil {
            return err
        }
        res := make([]ThreatAssessmentRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*ThreatAssessmentRequest))
        }
        m.SetThreatAssessmentRequests(res)
        return nil
    }
    return res
}
func (m *InformationProtection) IsNil()(bool) {
    return m == nil
}
func (m *InformationProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
func (m *InformationProtection) SetThreatAssessmentRequests(value []ThreatAssessmentRequest)() {
    m.threatAssessmentRequests = value
}
