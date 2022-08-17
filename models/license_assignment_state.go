package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LicenseAssignmentState 
type LicenseAssignmentState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
    assignedByGroup *string
    // The service plans that are disabled in this assignment. Read-Only.
    disabledPlans []string
    // License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. The possible values are CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Other. For more information on how to identify and resolve license assignment errors see here.
    error *string
    // The timestamp when the state of the license assignment was last updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The unique identifier for the SKU. Read-Only.
    skuId *string
    // Indicate the current state of this assignment. Read-Only. The possible values are Active, ActiveWithError, Disabled, and Error.
    state *string
}
// NewLicenseAssignmentState instantiates a new licenseAssignmentState and sets the default values.
func NewLicenseAssignmentState()(*LicenseAssignmentState) {
    m := &LicenseAssignmentState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.licenseAssignmentState";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateLicenseAssignmentStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLicenseAssignmentStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLicenseAssignmentState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LicenseAssignmentState) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignedByGroup gets the assignedByGroup property value. The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
func (m *LicenseAssignmentState) GetAssignedByGroup()(*string) {
    return m.assignedByGroup
}
// GetDisabledPlans gets the disabledPlans property value. The service plans that are disabled in this assignment. Read-Only.
func (m *LicenseAssignmentState) GetDisabledPlans()([]string) {
    return m.disabledPlans
}
// GetError gets the error property value. License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. The possible values are CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Other. For more information on how to identify and resolve license assignment errors see here.
func (m *LicenseAssignmentState) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LicenseAssignmentState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedByGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedByGroup(val)
        }
        return nil
    }
    res["disabledPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["skuId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    return m.lastUpdatedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LicenseAssignmentState) GetOdataType()(*string) {
    return m.odataType
}
// GetSkuId gets the skuId property value. The unique identifier for the SKU. Read-Only.
func (m *LicenseAssignmentState) GetSkuId()(*string) {
    return m.skuId
}
// GetState gets the state property value. Indicate the current state of this assignment. Read-Only. The possible values are Active, ActiveWithError, Disabled, and Error.
func (m *LicenseAssignmentState) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *LicenseAssignmentState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetError sets the error property value. License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. The possible values are CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Other. For more information on how to identify and resolve license assignment errors see here.
func (m *LicenseAssignmentState) SetError(value *string)() {
    m.error = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The timestamp when the state of the license assignment was last updated.
func (m *LicenseAssignmentState) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LicenseAssignmentState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSkuId sets the skuId property value. The unique identifier for the SKU. Read-Only.
func (m *LicenseAssignmentState) SetSkuId(value *string)() {
    m.skuId = value
}
// SetState sets the state property value. Indicate the current state of this assignment. Read-Only. The possible values are Active, ActiveWithError, Disabled, and Error.
func (m *LicenseAssignmentState) SetState(value *string)() {
    m.state = value
}
