package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityProtectionRoot 
type IdentityProtectionRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    riskDetections []RiskDetection;
    // 
    riskyUsers []RiskyUser;
}
// NewIdentityProtectionRoot instantiates a new IdentityProtectionRoot and sets the default values.
func NewIdentityProtectionRoot()(*IdentityProtectionRoot) {
    m := &IdentityProtectionRoot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityProtectionRoot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetRiskDetections gets the riskDetections property value. 
func (m *IdentityProtectionRoot) GetRiskDetections()([]RiskDetection) {
    if m == nil {
        return nil
    } else {
        return m.riskDetections
    }
}
// GetRiskyUsers gets the riskyUsers property value. 
func (m *IdentityProtectionRoot) GetRiskyUsers()([]RiskyUser) {
    if m == nil {
        return nil
    } else {
        return m.riskyUsers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityProtectionRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["riskDetections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRiskDetection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskDetection, len(val))
            for i, v := range val {
                res[i] = *(v.(*RiskDetection))
            }
            m.SetRiskDetections(res)
        }
        return nil
    }
    res["riskyUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRiskyUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskyUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*RiskyUser))
            }
            m.SetRiskyUsers(res)
        }
        return nil
    }
    return res
}
func (m *IdentityProtectionRoot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IdentityProtectionRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRiskDetections()))
        for i, v := range m.GetRiskDetections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("riskDetections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRiskyUsers()))
        for i, v := range m.GetRiskyUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("riskyUsers", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityProtectionRoot) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRiskDetections sets the riskDetections property value. 
func (m *IdentityProtectionRoot) SetRiskDetections(value []RiskDetection)() {
    if m != nil {
        m.riskDetections = value
    }
}
// SetRiskyUsers sets the riskyUsers property value. 
func (m *IdentityProtectionRoot) SetRiskyUsers(value []RiskyUser)() {
    if m != nil {
        m.riskyUsers = value
    }
}
