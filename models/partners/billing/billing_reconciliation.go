// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package billing

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

type BillingReconciliation struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
}
// NewBillingReconciliation instantiates a new BillingReconciliation and sets the default values.
func NewBillingReconciliation()(*BillingReconciliation) {
    m := &BillingReconciliation{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateBillingReconciliationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBillingReconciliationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBillingReconciliation(), nil
}
// GetBilled gets the billed property value. The billed property
// returns a BilledReconciliationable when successful
func (m *BillingReconciliation) GetBilled()(BilledReconciliationable) {
    val, err := m.GetBackingStore().Get("billed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BilledReconciliationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BillingReconciliation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["billed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBilledReconciliationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBilled(val.(BilledReconciliationable))
        }
        return nil
    }
    res["unbilled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnbilledReconciliationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnbilled(val.(UnbilledReconciliationable))
        }
        return nil
    }
    return res
}
// GetUnbilled gets the unbilled property value. The unbilled property
// returns a UnbilledReconciliationable when successful
func (m *BillingReconciliation) GetUnbilled()(UnbilledReconciliationable) {
    val, err := m.GetBackingStore().Get("unbilled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UnbilledReconciliationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BillingReconciliation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("billed", m.GetBilled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("unbilled", m.GetUnbilled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBilled sets the billed property value. The billed property
func (m *BillingReconciliation) SetBilled(value BilledReconciliationable)() {
    err := m.GetBackingStore().Set("billed", value)
    if err != nil {
        panic(err)
    }
}
// SetUnbilled sets the unbilled property value. The unbilled property
func (m *BillingReconciliation) SetUnbilled(value UnbilledReconciliationable)() {
    err := m.GetBackingStore().Set("unbilled", value)
    if err != nil {
        panic(err)
    }
}
type BillingReconciliationable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBilled()(BilledReconciliationable)
    GetUnbilled()(UnbilledReconciliationable)
    SetBilled(value BilledReconciliationable)()
    SetUnbilled(value UnbilledReconciliationable)()
}
