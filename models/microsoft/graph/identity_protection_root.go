package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityProtectionRoot 
type IdentityProtectionRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Risk detection in Azure AD Identity Protection and the associated information about the detection.
    riskDetections []RiskDetectionable;
    // Users that are flagged as at-risk by Azure AD Identity Protection.
    riskyUsers []RiskyUserable;
}
// NewIdentityProtectionRoot instantiates a new IdentityProtectionRoot and sets the default values.
func NewIdentityProtectionRoot()(*IdentityProtectionRoot) {
    m := &IdentityProtectionRoot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityProtectionRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityProtectionRootFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewIdentityProtectionRoot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityProtectionRoot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityProtectionRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["riskDetections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRiskDetectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskDetectionable, len(val))
            for i, v := range val {
                res[i] = v.(RiskDetectionable)
            }
            m.SetRiskDetections(res)
        }
        return nil
    }
    res["riskyUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRiskyUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskyUserable, len(val))
            for i, v := range val {
                res[i] = v.(RiskyUserable)
            }
            m.SetRiskyUsers(res)
        }
        return nil
    }
    return res
}
// GetRiskDetections gets the riskDetections property value. Risk detection in Azure AD Identity Protection and the associated information about the detection.
func (m *IdentityProtectionRoot) GetRiskDetections()([]RiskDetectionable) {
    if m == nil {
        return nil
    } else {
        return m.riskDetections
    }
}
// GetRiskyUsers gets the riskyUsers property value. Users that are flagged as at-risk by Azure AD Identity Protection.
func (m *IdentityProtectionRoot) GetRiskyUsers()([]RiskyUserable) {
    if m == nil {
        return nil
    } else {
        return m.riskyUsers
    }
}
// Serialize serializes information the current object
func (m *IdentityProtectionRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetRiskDetections() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRiskDetections()))
        for i, v := range m.GetRiskDetections() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("riskDetections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskyUsers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRiskyUsers()))
        for i, v := range m.GetRiskyUsers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetRiskDetections sets the riskDetections property value. Risk detection in Azure AD Identity Protection and the associated information about the detection.
func (m *IdentityProtectionRoot) SetRiskDetections(value []RiskDetectionable)() {
    if m != nil {
        m.riskDetections = value
    }
}
// SetRiskyUsers sets the riskyUsers property value. Users that are flagged as at-risk by Azure AD Identity Protection.
func (m *IdentityProtectionRoot) SetRiskyUsers(value []RiskyUserable)() {
    if m != nil {
        m.riskyUsers = value
    }
}
