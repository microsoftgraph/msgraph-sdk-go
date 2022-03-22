package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentApprovalSettings 
type AccessPackageAssignmentApprovalSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    isApprovalRequiredForAdd *bool;
    // 
    isApprovalRequiredForUpdate *bool;
    // 
    stages []AccessPackageApprovalStageable;
}
// NewAccessPackageAssignmentApprovalSettings instantiates a new accessPackageAssignmentApprovalSettings and sets the default values.
func NewAccessPackageAssignmentApprovalSettings()(*AccessPackageAssignmentApprovalSettings) {
    m := &AccessPackageAssignmentApprovalSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessPackageAssignmentApprovalSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentApprovalSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessPackageAssignmentApprovalSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentApprovalSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentApprovalSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isApprovalRequiredForAdd"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForAdd(val)
        }
        return nil
    }
    res["isApprovalRequiredForUpdate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForUpdate(val)
        }
        return nil
    }
    res["stages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageApprovalStageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageApprovalStageable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageApprovalStageable)
            }
            m.SetStages(res)
        }
        return nil
    }
    return res
}
// GetIsApprovalRequiredForAdd gets the isApprovalRequiredForAdd property value. 
func (m *AccessPackageAssignmentApprovalSettings) GetIsApprovalRequiredForAdd()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequiredForAdd
    }
}
// GetIsApprovalRequiredForUpdate gets the isApprovalRequiredForUpdate property value. 
func (m *AccessPackageAssignmentApprovalSettings) GetIsApprovalRequiredForUpdate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequiredForUpdate
    }
}
// GetStages gets the stages property value. 
func (m *AccessPackageAssignmentApprovalSettings) GetStages()([]AccessPackageApprovalStageable) {
    if m == nil {
        return nil
    } else {
        return m.stages
    }
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentApprovalSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isApprovalRequiredForAdd", m.GetIsApprovalRequiredForAdd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequiredForUpdate", m.GetIsApprovalRequiredForUpdate())
        if err != nil {
            return err
        }
    }
    if m.GetStages() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStages()))
        for i, v := range m.GetStages() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("stages", cast)
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
func (m *AccessPackageAssignmentApprovalSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsApprovalRequiredForAdd sets the isApprovalRequiredForAdd property value. 
func (m *AccessPackageAssignmentApprovalSettings) SetIsApprovalRequiredForAdd(value *bool)() {
    if m != nil {
        m.isApprovalRequiredForAdd = value
    }
}
// SetIsApprovalRequiredForUpdate sets the isApprovalRequiredForUpdate property value. 
func (m *AccessPackageAssignmentApprovalSettings) SetIsApprovalRequiredForUpdate(value *bool)() {
    if m != nil {
        m.isApprovalRequiredForUpdate = value
    }
}
// SetStages sets the stages property value. 
func (m *AccessPackageAssignmentApprovalSettings) SetStages(value []AccessPackageApprovalStageable)() {
    if m != nil {
        m.stages = value
    }
}
