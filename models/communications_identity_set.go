package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsIdentitySet 
type CommunicationsIdentitySet struct {
    IdentitySet
}
// NewCommunicationsIdentitySet instantiates a new communicationsIdentitySet and sets the default values.
func NewCommunicationsIdentitySet()(*CommunicationsIdentitySet) {
    m := &CommunicationsIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    odataTypeValue := "#microsoft.graph.communicationsIdentitySet"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCommunicationsIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsIdentitySet(), nil
}
// GetApplicationInstance gets the applicationInstance property value. The applicationInstance property
func (m *CommunicationsIdentitySet) GetApplicationInstance()(Identityable) {
    val, err := m.GetBackingStore().Get("applicationInstance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetAssertedIdentity gets the assertedIdentity property value. The assertedIdentity property
func (m *CommunicationsIdentitySet) GetAssertedIdentity()(Identityable) {
    val, err := m.GetBackingStore().Get("assertedIdentity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetAzureCommunicationServicesUser gets the azureCommunicationServicesUser property value. The azureCommunicationServicesUser property
func (m *CommunicationsIdentitySet) GetAzureCommunicationServicesUser()(Identityable) {
    val, err := m.GetBackingStore().Get("azureCommunicationServicesUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetEncrypted gets the encrypted property value. The encrypted property
func (m *CommunicationsIdentitySet) GetEncrypted()(Identityable) {
    val, err := m.GetBackingStore().Get("encrypted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetEndpointType gets the endpointType property value. The endpointType property
func (m *CommunicationsIdentitySet) GetEndpointType()(*EndpointType) {
    val, err := m.GetBackingStore().Get("endpointType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EndpointType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["applicationInstance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationInstance(val.(Identityable))
        }
        return nil
    }
    res["assertedIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssertedIdentity(val.(Identityable))
        }
        return nil
    }
    res["azureCommunicationServicesUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureCommunicationServicesUser(val.(Identityable))
        }
        return nil
    }
    res["encrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncrypted(val.(Identityable))
        }
        return nil
    }
    res["endpointType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndpointType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointType(val.(*EndpointType))
        }
        return nil
    }
    res["guest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuest(val.(Identityable))
        }
        return nil
    }
    res["onPremises"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremises(val.(Identityable))
        }
        return nil
    }
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val.(Identityable))
        }
        return nil
    }
    return res
}
// GetGuest gets the guest property value. The guest property
func (m *CommunicationsIdentitySet) GetGuest()(Identityable) {
    val, err := m.GetBackingStore().Get("guest")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetOnPremises gets the onPremises property value. The onPremises property
func (m *CommunicationsIdentitySet) GetOnPremises()(Identityable) {
    val, err := m.GetBackingStore().Get("onPremises")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetPhone gets the phone property value. The phone property
func (m *CommunicationsIdentitySet) GetPhone()(Identityable) {
    val, err := m.GetBackingStore().Get("phone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CommunicationsIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applicationInstance", m.GetApplicationInstance())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assertedIdentity", m.GetAssertedIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("azureCommunicationServicesUser", m.GetAzureCommunicationServicesUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("encrypted", m.GetEncrypted())
        if err != nil {
            return err
        }
    }
    if m.GetEndpointType() != nil {
        cast := (*m.GetEndpointType()).String()
        err = writer.WriteStringValue("endpointType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("guest", m.GetGuest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onPremises", m.GetOnPremises())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationInstance sets the applicationInstance property value. The applicationInstance property
func (m *CommunicationsIdentitySet) SetApplicationInstance(value Identityable)() {
    err := m.GetBackingStore().Set("applicationInstance", value)
    if err != nil {
        panic(err)
    }
}
// SetAssertedIdentity sets the assertedIdentity property value. The assertedIdentity property
func (m *CommunicationsIdentitySet) SetAssertedIdentity(value Identityable)() {
    err := m.GetBackingStore().Set("assertedIdentity", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureCommunicationServicesUser sets the azureCommunicationServicesUser property value. The azureCommunicationServicesUser property
func (m *CommunicationsIdentitySet) SetAzureCommunicationServicesUser(value Identityable)() {
    err := m.GetBackingStore().Set("azureCommunicationServicesUser", value)
    if err != nil {
        panic(err)
    }
}
// SetEncrypted sets the encrypted property value. The encrypted property
func (m *CommunicationsIdentitySet) SetEncrypted(value Identityable)() {
    err := m.GetBackingStore().Set("encrypted", value)
    if err != nil {
        panic(err)
    }
}
// SetEndpointType sets the endpointType property value. The endpointType property
func (m *CommunicationsIdentitySet) SetEndpointType(value *EndpointType)() {
    err := m.GetBackingStore().Set("endpointType", value)
    if err != nil {
        panic(err)
    }
}
// SetGuest sets the guest property value. The guest property
func (m *CommunicationsIdentitySet) SetGuest(value Identityable)() {
    err := m.GetBackingStore().Set("guest", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremises sets the onPremises property value. The onPremises property
func (m *CommunicationsIdentitySet) SetOnPremises(value Identityable)() {
    err := m.GetBackingStore().Set("onPremises", value)
    if err != nil {
        panic(err)
    }
}
// SetPhone sets the phone property value. The phone property
func (m *CommunicationsIdentitySet) SetPhone(value Identityable)() {
    err := m.GetBackingStore().Set("phone", value)
    if err != nil {
        panic(err)
    }
}
// CommunicationsIdentitySetable 
type CommunicationsIdentitySetable interface {
    IdentitySetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationInstance()(Identityable)
    GetAssertedIdentity()(Identityable)
    GetAzureCommunicationServicesUser()(Identityable)
    GetEncrypted()(Identityable)
    GetEndpointType()(*EndpointType)
    GetGuest()(Identityable)
    GetOnPremises()(Identityable)
    GetPhone()(Identityable)
    SetApplicationInstance(value Identityable)()
    SetAssertedIdentity(value Identityable)()
    SetAzureCommunicationServicesUser(value Identityable)()
    SetEncrypted(value Identityable)()
    SetEndpointType(value *EndpointType)()
    SetGuest(value Identityable)()
    SetOnPremises(value Identityable)()
    SetPhone(value Identityable)()
}
