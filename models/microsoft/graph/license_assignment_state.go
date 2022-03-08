package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LicenseAssignmentState provides operations to manage the drive singleton.
type LicenseAssignmentState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
    assignedByGroup *string;
    // The service plans that are disabled in this assignment. Read-Only.
    disabledPlans []string;
    // License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. The possible values are CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Other. For more information on how to identify and resolve license assignment errors see here.
    error *string;
    // The timestamp when the state of the license assignment was last updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The unique identifier for the SKU. Read-Only.
    skuId *string;
    // Indicate the current state of this assignment. Read-Only. The possible values are Active, ActiveWithError, Disabled, and Error.
    state *string;
}
// NewLicenseAssignmentState instantiates a new licenseAssignmentState and sets the default values.
func NewLicenseAssignmentState()(*LicenseAssignmentState) {
    m := &LicenseAssignmentState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLicenseAssignmentStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLicenseAssignmentStateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewLicenseAssignmentState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
// GetError gets the error property value. License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. The possible values are CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Other. For more information on how to identify and resolve license assignment errors see here.
func (m *LicenseAssignmentState) GetError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.error
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
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
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
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The timestamp when the state of the license assignment was last updated.
func (m *LicenseAssignmentState) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
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
// GetState gets the state property value. Indicate the current state of this assignment. Read-Only. The possible values are Active, ActiveWithError, Disabled, and Error.
func (m *LicenseAssignmentState) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
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
    if m.GetDisabledPlans() != nil {
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
        err := writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LicenseAssignmentState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignedByGroup sets the assignedByGroup property value. The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
func (m *LicenseAssignmentState) SetAssignedByGroup(value *string)() {
    if m != nil {
        m.assignedByGroup = value
    }
}
// SetDisabledPlans sets the disabledPlans property value. The service plans that are disabled in this assignment. Read-Only.
func (m *LicenseAssignmentState) SetDisabledPlans(value []string)() {
    if m != nil {
        m.disabledPlans = value
    }
}
// SetError sets the error property value. License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. The possible values are CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Other. For more information on how to identify and resolve license assignment errors see here.
func (m *LicenseAssignmentState) SetError(value *string)() {
    if m != nil {
        m.error = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The timestamp when the state of the license assignment was last updated.
func (m *LicenseAssignmentState) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetSkuId sets the skuId property value. The unique identifier for the SKU. Read-Only.
func (m *LicenseAssignmentState) SetSkuId(value *string)() {
    if m != nil {
        m.skuId = value
    }
}
// SetState sets the state property value. Indicate the current state of this assignment. Read-Only. The possible values are Active, ActiveWithError, Disabled, and Error.
func (m *LicenseAssignmentState) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
