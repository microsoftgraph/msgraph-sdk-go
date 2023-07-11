package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventMessageResponse 
type EventMessageResponse struct {
    EventMessage
}
// NewEventMessageResponse instantiates a new eventMessageResponse and sets the default values.
func NewEventMessageResponse()(*EventMessageResponse) {
    m := &EventMessageResponse{
        EventMessage: *NewEventMessage(),
    }
    odataTypeValue := "#microsoft.graph.eventMessageResponse"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEventMessageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEventMessageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEventMessageResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EventMessageResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessage.GetFieldDeserializers()
    res["proposedNewTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposedNewTime(val.(TimeSlotable))
        }
        return nil
    }
    res["responseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResponseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseType(val.(*ResponseType))
        }
        return nil
    }
    return res
}
// GetProposedNewTime gets the proposedNewTime property value. The proposedNewTime property
func (m *EventMessageResponse) GetProposedNewTime()(TimeSlotable) {
    val, err := m.GetBackingStore().Get("proposedNewTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TimeSlotable)
    }
    return nil
}
// GetResponseType gets the responseType property value. The responseType property
func (m *EventMessageResponse) GetResponseType()(*ResponseType) {
    val, err := m.GetBackingStore().Get("responseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ResponseType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EventMessageResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessage.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("proposedNewTime", m.GetProposedNewTime())
        if err != nil {
            return err
        }
    }
    if m.GetResponseType() != nil {
        cast := (*m.GetResponseType()).String()
        err = writer.WriteStringValue("responseType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetProposedNewTime sets the proposedNewTime property value. The proposedNewTime property
func (m *EventMessageResponse) SetProposedNewTime(value TimeSlotable)() {
    err := m.GetBackingStore().Set("proposedNewTime", value)
    if err != nil {
        panic(err)
    }
}
// SetResponseType sets the responseType property value. The responseType property
func (m *EventMessageResponse) SetResponseType(value *ResponseType)() {
    err := m.GetBackingStore().Set("responseType", value)
    if err != nil {
        panic(err)
    }
}
// EventMessageResponseable 
type EventMessageResponseable interface {
    EventMessageable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProposedNewTime()(TimeSlotable)
    GetResponseType()(*ResponseType)
    SetProposedNewTime(value TimeSlotable)()
    SetResponseType(value *ResponseType)()
}
