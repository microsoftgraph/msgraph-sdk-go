package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LicenseAssignmentState 
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
// NewLicenseAssignmentState instantiates a new licenseAssignmentState and sets the default values.
func NewLicenseAssignmentState()(*LicenseAssignmentState) {
    m := &LicenseAssignmentState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LicenseAssignmentState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignedByGroup gets the assignedByGroup property value. The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
func (m *LicenseAssignmentState) GetAssignedByGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedByGroup
    }
}
// GetDisabledPlans gets the disabledPlans property value. The service plans that are disabled in this assignment. Read-Only.
func (m *LicenseAssignmentState) GetDisabledPlans()([]string) {
    if m == nil {
        return nil
    } else {
        return m.disabledPlans
    }
}
// GetError gets the error property value. License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. Possible values: CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Others. For more information on how to identify and resolve license assignment errors see here.
func (m *LicenseAssignmentState) GetError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetSkuId gets the skuId property value. The unique identifier for the SKU. Read-Only.
func (m *LicenseAssignmentState) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// GetState gets the state property value. Indicate the current state of this assignment. Read-Only. Possible values: Active, ActiveWithError, Disabled and Error.
func (m *LicenseAssignmentState) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LicenseAssignmentState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignedByGroup sets the assignedByGroup property value. The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
func (m *LicenseAssignmentState) SetAssignedByGroup(value *string)() {
    m.assignedByGroup = value
}
// SetDisabledPlans sets the disabledPlans property value. The service plans that are disabled in this assignment. Read-Only.
func (m *LicenseAssignmentState) SetDisabledPlans(value []string)() {
    m.disabledPlans = value
}
// SetError sets the error property value. License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. Possible values: CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Others. For more information on how to identify and resolve license assignment errors see here.
func (m *LicenseAssignmentState) SetError(value *string)() {
    m.error = value
}
// SetSkuId sets the skuId property value. The unique identifier for the SKU. Read-Only.
func (m *LicenseAssignmentState) SetSkuId(value *string)() {
    m.skuId = value
}
// SetState sets the state property value. Indicate the current state of this assignment. Read-Only. Possible values: Active, ActiveWithError, Disabled and Error.
func (m *LicenseAssignmentState) SetState(value *string)() {
    m.state = value
}
