package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryObject 
type DirectoryObject struct {
    Entity
    // Date and time when this object was deleted. Always null when the object hasn't been deleted.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewDirectoryObject instantiates a new directoryObject and sets the default values.
func NewDirectoryObject()(*DirectoryObject) {
    m := &DirectoryObject{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDirectoryObjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryObjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryObject(), nil
}
// GetDeletedDateTime gets the deletedDateTime property value. Date and time when this object was deleted. Always null when the object hasn't been deleted.
func (m *DirectoryObject) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deletedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryObject) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deletedDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDateTime(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DirectoryObject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("deletedDateTime", m.GetDeletedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedDateTime sets the deletedDateTime property value. Date and time when this object was deleted. Always null when the object hasn't been deleted.
func (m *DirectoryObject) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.deletedDateTime = value
    }
}
