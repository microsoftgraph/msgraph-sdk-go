package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CloudCommunications struct {
    Entity
    callRecords []CallRecord;
    calls []Call;
    onlineMeetings []OnlineMeeting;
    presences []Presence;
}
func NewCloudCommunications()(*CloudCommunications) {
    m := &CloudCommunications{
        Entity: *NewEntity(),
    }
    return m
}
func (m *CloudCommunications) GetCallRecords()([]CallRecord) {
    if m == nil {
        return nil
    } else {
        return m.callRecords
    }
}
func (m *CloudCommunications) GetCalls()([]Call) {
    if m == nil {
        return nil
    } else {
        return m.calls
    }
}
func (m *CloudCommunications) GetOnlineMeetings()([]OnlineMeeting) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetings
    }
}
func (m *CloudCommunications) GetPresences()([]Presence) {
    if m == nil {
        return nil
    } else {
        return m.presences
    }
}
func (m *CloudCommunications) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["callRecords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCallRecord() })
        if err != nil {
            return err
        }
        res := make([]CallRecord, len(val))
        for i, v := range val {
            res[i] = *(v.(*CallRecord))
        }
        m.SetCallRecords(res)
        return nil
    }
    res["calls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCall() })
        if err != nil {
            return err
        }
        res := make([]Call, len(val))
        for i, v := range val {
            res[i] = *(v.(*Call))
        }
        m.SetCalls(res)
        return nil
    }
    res["onlineMeetings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnlineMeeting() })
        if err != nil {
            return err
        }
        res := make([]OnlineMeeting, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnlineMeeting))
        }
        m.SetOnlineMeetings(res)
        return nil
    }
    res["presences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPresence() })
        if err != nil {
            return err
        }
        res := make([]Presence, len(val))
        for i, v := range val {
            res[i] = *(v.(*Presence))
        }
        m.SetPresences(res)
        return nil
    }
    return res
}
func (m *CloudCommunications) IsNil()(bool) {
    return m == nil
}
func (m *CloudCommunications) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCallRecords()))
        for i, v := range m.GetCallRecords() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("callRecords", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalls()))
        for i, v := range m.GetCalls() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calls", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOnlineMeetings()))
        for i, v := range m.GetOnlineMeetings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("onlineMeetings", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPresences()))
        for i, v := range m.GetPresences() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("presences", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *CloudCommunications) SetCallRecords(value []CallRecord)() {
    m.callRecords = value
}
func (m *CloudCommunications) SetCalls(value []Call)() {
    m.calls = value
}
func (m *CloudCommunications) SetOnlineMeetings(value []OnlineMeeting)() {
    m.onlineMeetings = value
}
func (m *CloudCommunications) SetPresences(value []Presence)() {
    m.presences = value
}
