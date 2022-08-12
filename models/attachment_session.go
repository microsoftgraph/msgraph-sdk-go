package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttachmentSession provides operations to manage the admin singleton.
type AttachmentSession struct {
    Entity
    // The content streams that are uploaded.
    content []byte
    // The date and time in UTC when the upload session will expire. The complete file must be uploaded before this expiration time is reached.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates a single value {start} that represents the location in the file where the next upload should begin.
    nextExpectedRanges []string
}
// NewAttachmentSession instantiates a new attachmentSession and sets the default values.
func NewAttachmentSession()(*AttachmentSession) {
    m := &AttachmentSession{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.attachmentSession";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAttachmentSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttachmentSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttachmentSession(), nil
}
// GetContent gets the content property value. The content streams that are uploaded.
func (m *AttachmentSession) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The date and time in UTC when the upload session will expire. The complete file must be uploaded before this expiration time is reached.
func (m *AttachmentSession) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttachmentSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["nextExpectedRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNextExpectedRanges(res)
        }
        return nil
    }
    return res
}
// GetNextExpectedRanges gets the nextExpectedRanges property value. Indicates a single value {start} that represents the location in the file where the next upload should begin.
func (m *AttachmentSession) GetNextExpectedRanges()([]string) {
    if m == nil {
        return nil
    } else {
        return m.nextExpectedRanges
    }
}
// Serialize serializes information the current object
func (m *AttachmentSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetNextExpectedRanges() != nil {
        err = writer.WriteCollectionOfStringValues("nextExpectedRanges", m.GetNextExpectedRanges())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content streams that are uploaded.
func (m *AttachmentSession) SetContent(value []byte)() {
    if m != nil {
        m.content = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The date and time in UTC when the upload session will expire. The complete file must be uploaded before this expiration time is reached.
func (m *AttachmentSession) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetNextExpectedRanges sets the nextExpectedRanges property value. Indicates a single value {start} that represents the location in the file where the next upload should begin.
func (m *AttachmentSession) SetNextExpectedRanges(value []string)() {
    if m != nil {
        m.nextExpectedRanges = value
    }
}
