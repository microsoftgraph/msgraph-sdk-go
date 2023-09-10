package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestEnumeratedMailboxLocation 
type SubjectRightsRequestEnumeratedMailboxLocation struct {
    SubjectRightsRequestMailboxLocation
}
// NewSubjectRightsRequestEnumeratedMailboxLocation instantiates a new subjectRightsRequestEnumeratedMailboxLocation and sets the default values.
func NewSubjectRightsRequestEnumeratedMailboxLocation()(*SubjectRightsRequestEnumeratedMailboxLocation) {
    m := &SubjectRightsRequestEnumeratedMailboxLocation{
        SubjectRightsRequestMailboxLocation: *NewSubjectRightsRequestMailboxLocation(),
    }
    odataTypeValue := "#microsoft.graph.subjectRightsRequestEnumeratedMailboxLocation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSubjectRightsRequestEnumeratedMailboxLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestEnumeratedMailboxLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequestEnumeratedMailboxLocation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestEnumeratedMailboxLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectRightsRequestMailboxLocation.GetFieldDeserializers()
    res["userPrincipalNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetUserPrincipalNames(res)
        }
        return nil
    }
    return res
}
// GetUserPrincipalNames gets the userPrincipalNames property value. The userPrincipalNames property
func (m *SubjectRightsRequestEnumeratedMailboxLocation) GetUserPrincipalNames()([]string) {
    val, err := m.GetBackingStore().Get("userPrincipalNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SubjectRightsRequestEnumeratedMailboxLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectRightsRequestMailboxLocation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserPrincipalNames() != nil {
        err = writer.WriteCollectionOfStringValues("userPrincipalNames", m.GetUserPrincipalNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserPrincipalNames sets the userPrincipalNames property value. The userPrincipalNames property
func (m *SubjectRightsRequestEnumeratedMailboxLocation) SetUserPrincipalNames(value []string)() {
    err := m.GetBackingStore().Set("userPrincipalNames", value)
    if err != nil {
        panic(err)
    }
}
// SubjectRightsRequestEnumeratedMailboxLocationable 
type SubjectRightsRequestEnumeratedMailboxLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SubjectRightsRequestMailboxLocationable
    GetUserPrincipalNames()([]string)
    SetUserPrincipalNames(value []string)()
}
