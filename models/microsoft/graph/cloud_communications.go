package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudCommunications provides operations to manage the cloudCommunications singleton.
type CloudCommunications struct {
    Entity
    // 
    callRecords []CallRecordable;
    // 
    calls []Callable;
    // 
    onlineMeetings []OnlineMeetingable;
    // 
    presences []Presenceable;
}
// NewCloudCommunications instantiates a new cloudCommunications and sets the default values.
func NewCloudCommunications()(*CloudCommunications) {
    m := &CloudCommunications{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudCommunicationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudCommunicationsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudCommunications(), nil
}
// GetCallRecords gets the callRecords property value. 
func (m *CloudCommunications) GetCallRecords()([]CallRecordable) {
    if m == nil {
        return nil
    } else {
        return m.callRecords
    }
}
// GetCalls gets the calls property value. 
func (m *CloudCommunications) GetCalls()([]Callable) {
    if m == nil {
        return nil
    } else {
        return m.calls
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudCommunications) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["callRecords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCallRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CallRecordable, len(val))
            for i, v := range val {
                res[i] = v.(CallRecordable)
            }
            m.SetCallRecords(res)
        }
        return nil
    }
    res["calls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCallFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Callable, len(val))
            for i, v := range val {
                res[i] = v.(Callable)
            }
            m.SetCalls(res)
        }
        return nil
    }
    res["onlineMeetings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnlineMeetingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnlineMeetingable, len(val))
            for i, v := range val {
                res[i] = v.(OnlineMeetingable)
            }
            m.SetOnlineMeetings(res)
        }
        return nil
    }
    res["presences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePresenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Presenceable, len(val))
            for i, v := range val {
                res[i] = v.(Presenceable)
            }
            m.SetPresences(res)
        }
        return nil
    }
    return res
}
// GetOnlineMeetings gets the onlineMeetings property value. 
func (m *CloudCommunications) GetOnlineMeetings()([]OnlineMeetingable) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetings
    }
}
// GetPresences gets the presences property value. 
func (m *CloudCommunications) GetPresences()([]Presenceable) {
    if m == nil {
        return nil
    } else {
        return m.presences
    }
}
func (m *CloudCommunications) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudCommunications) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCallRecords() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCallRecords()))
        for i, v := range m.GetCallRecords() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("callRecords", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCalls() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalls()))
        for i, v := range m.GetCalls() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("calls", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOnlineMeetings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOnlineMeetings()))
        for i, v := range m.GetOnlineMeetings() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("onlineMeetings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPresences() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPresences()))
        for i, v := range m.GetPresences() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("presences", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallRecords sets the callRecords property value. 
func (m *CloudCommunications) SetCallRecords(value []CallRecordable)() {
    if m != nil {
        m.callRecords = value
    }
}
// SetCalls sets the calls property value. 
func (m *CloudCommunications) SetCalls(value []Callable)() {
    if m != nil {
        m.calls = value
    }
}
// SetOnlineMeetings sets the onlineMeetings property value. 
func (m *CloudCommunications) SetOnlineMeetings(value []OnlineMeetingable)() {
    if m != nil {
        m.onlineMeetings = value
    }
}
// SetPresences sets the presences property value. 
func (m *CloudCommunications) SetPresences(value []Presenceable)() {
    if m != nil {
        m.presences = value
    }
}
