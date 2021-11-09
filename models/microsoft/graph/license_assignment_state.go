package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type LicenseAssignmentState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
    assignedByGroup *string;
    // The service plans that are disabled in this assignment. Read-Only.
    disabledPlans []string;
    // License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. Possible values: CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Others. For more information on how to identify and resolve license assignment errors see here.
    error *string;
    // The unique identifier for the SKU. Read-Only.
    skuId *string;
    // Indicate the current state of this assignment. Read-Only. Possible values: Active, ActiveWithError, Disabled and Error.
    state *string;
}
// Instantiates a new licenseAssignmentState and sets the default values.
func NewLicenseAssignmentState()(*LicenseAssignmentState) {
    m := &LicenseAssignmentState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LicenseAssignmentState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the assignedByGroup property value. The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
func (m *LicenseAssignmentState) GetAssignedByGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedByGroup
    }
}
// Gets the disabledPlans property value. The service plans that are disabled in this assignment. Read-Only.
func (m *LicenseAssignmentState) GetDisabledPlans()([]string) {
    if m == nil {
        return nil
    } else {
        return m.disabledPlans
    }
}
// Gets the error property value. License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. Possible values: CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Others. For more information on how to identify and resolve license assignment errors see here.
func (m *LicenseAssignmentState) GetError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// Gets the skuId property value. The unique identifier for the SKU. Read-Only.
func (m *LicenseAssignmentState) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// Gets the state property value. Indicate the current state of this assignment. Read-Only. Possible values: Active, ActiveWithError, Disabled and Error.
func (m *LicenseAssignmentState) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *LicenseAssignmentState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignedByGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedByGroup(val)
        }
        return nil
    }
    res["disabledPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDisabledPlans(res)
        }
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["skuId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
func (m *LicenseAssignmentState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LicenseAssignmentState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignedByGroup", m.GetAssignedByGroup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("disabledPlans", m.GetDisabledPlans())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("skuId", m.GetSkuId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *LicenseAssignmentState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the assignedByGroup property value. The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
// Parameters:
//  - value : Value to set for the assignedByGroup property.
func (m *LicenseAssignmentState) SetAssignedByGroup(value *string)() {
    m.assignedByGroup = value
}
// Sets the disabledPlans property value. The service plans that are disabled in this assignment. Read-Only.
// Parameters:
//  - value : Value to set for the disabledPlans property.
func (m *LicenseAssignmentState) SetDisabledPlans(value []string)() {
    m.disabledPlans = value
}
// Sets the error property value. License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. Possible values: CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Others. For more information on how to identify and resolve license assignment errors see here.
// Parameters:
//  - value : Value to set for the error property.
func (m *LicenseAssignmentState) SetError(value *string)() {
    m.error = value
}
// Sets the skuId property value. The unique identifier for the SKU. Read-Only.
// Parameters:
//  - value : Value to set for the skuId property.
func (m *LicenseAssignmentState) SetSkuId(value *string)() {
    m.skuId = value
}
// Sets the state property value. Indicate the current state of this assignment. Read-Only. Possible values: Active, ActiveWithError, Disabled and Error.
// Parameters:
//  - value : Value to set for the state property.
func (m *LicenseAssignmentState) SetState(value *string)() {
    m.state = value
}
