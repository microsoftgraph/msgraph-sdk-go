package billing

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Billing 
type Billing struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
}
// NewBilling instantiates a new billing and sets the default values.
func NewBilling()(*Billing) {
    m := &Billing{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateBillingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBillingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBilling(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Billing) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["manifests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManifestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Manifestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Manifestable)
                }
            }
            m.SetManifests(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Operationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Operationable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["usage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsage(val.(AzureUsageable))
        }
        return nil
    }
    return res
}
// GetManifests gets the manifests property value. Represents metadata for the exported data.
func (m *Billing) GetManifests()([]Manifestable) {
    val, err := m.GetBackingStore().Get("manifests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Manifestable)
    }
    return nil
}
// GetOperations gets the operations property value. Represents an operation to export the billing data of a partner.
func (m *Billing) GetOperations()([]Operationable) {
    val, err := m.GetBackingStore().Get("operations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Operationable)
    }
    return nil
}
// GetUsage gets the usage property value. The usage property
func (m *Billing) GetUsage()(AzureUsageable) {
    val, err := m.GetBackingStore().Get("usage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AzureUsageable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Billing) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetManifests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManifests()))
        for i, v := range m.GetManifests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("manifests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("usage", m.GetUsage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManifests sets the manifests property value. Represents metadata for the exported data.
func (m *Billing) SetManifests(value []Manifestable)() {
    err := m.GetBackingStore().Set("manifests", value)
    if err != nil {
        panic(err)
    }
}
// SetOperations sets the operations property value. Represents an operation to export the billing data of a partner.
func (m *Billing) SetOperations(value []Operationable)() {
    err := m.GetBackingStore().Set("operations", value)
    if err != nil {
        panic(err)
    }
}
// SetUsage sets the usage property value. The usage property
func (m *Billing) SetUsage(value AzureUsageable)() {
    err := m.GetBackingStore().Set("usage", value)
    if err != nil {
        panic(err)
    }
}
// Billingable 
type Billingable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManifests()([]Manifestable)
    GetOperations()([]Operationable)
    GetUsage()(AzureUsageable)
    SetManifests(value []Manifestable)()
    SetOperations(value []Operationable)()
    SetUsage(value AzureUsageable)()
}
