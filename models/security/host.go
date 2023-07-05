package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Host 
type Host struct {
    Artifact
}
// NewHost instantiates a new Host and sets the default values.
func NewHost()(*Host) {
    m := &Host{
        Artifact: *NewArtifact(),
    }
    odataTypeValue := "#microsoft.graph.security.host"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateHostFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHostFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.hostname":
                        return NewHostname(), nil
                    case "#microsoft.graph.security.ipAddress":
                        return NewIpAddress(), nil
                }
            }
        }
    }
    return NewHost(), nil
}
// GetComponents gets the components property value. The hostComponents that are associated with this host.
func (m *Host) GetComponents()([]HostComponentable) {
    val, err := m.GetBackingStore().Get("components")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HostComponentable)
    }
    return nil
}
// GetCookies gets the cookies property value. The hostCookies that are associated with this host.
func (m *Host) GetCookies()([]HostCookieable) {
    val, err := m.GetBackingStore().Get("cookies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HostCookieable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Host) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Artifact.GetFieldDeserializers()
    res["components"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostComponentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostComponentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostComponentable)
                }
            }
            m.SetComponents(res)
        }
        return nil
    }
    res["cookies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostCookieFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostCookieable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostCookieable)
                }
            }
            m.SetCookies(res)
        }
        return nil
    }
    res["firstSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["passiveDns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePassiveDnsRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PassiveDnsRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PassiveDnsRecordable)
                }
            }
            m.SetPassiveDns(res)
        }
        return nil
    }
    res["passiveDnsReverse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePassiveDnsRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PassiveDnsRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PassiveDnsRecordable)
                }
            }
            m.SetPassiveDnsReverse(res)
        }
        return nil
    }
    res["reputation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHostReputationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReputation(val.(HostReputationable))
        }
        return nil
    }
    res["trackers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostTrackerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostTrackerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostTrackerable)
                }
            }
            m.SetTrackers(res)
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The first date and time this host was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *Host) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("firstSeenDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The most recent date and time when this host was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *Host) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSeenDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPassiveDns gets the passiveDns property value. Passive DNS retrieval about this host.
func (m *Host) GetPassiveDns()([]PassiveDnsRecordable) {
    val, err := m.GetBackingStore().Get("passiveDns")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PassiveDnsRecordable)
    }
    return nil
}
// GetPassiveDnsReverse gets the passiveDnsReverse property value. Reverse passive DNS retrieval about this host.
func (m *Host) GetPassiveDnsReverse()([]PassiveDnsRecordable) {
    val, err := m.GetBackingStore().Get("passiveDnsReverse")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PassiveDnsRecordable)
    }
    return nil
}
// GetReputation gets the reputation property value. Represents a calculated reputation of this host.
func (m *Host) GetReputation()(HostReputationable) {
    val, err := m.GetBackingStore().Get("reputation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(HostReputationable)
    }
    return nil
}
// GetTrackers gets the trackers property value. The hostTrackers that are associated with this host.
func (m *Host) GetTrackers()([]HostTrackerable) {
    val, err := m.GetBackingStore().Get("trackers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HostTrackerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Host) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Artifact.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetComponents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComponents()))
        for i, v := range m.GetComponents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("components", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCookies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCookies()))
        for i, v := range m.GetCookies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("cookies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPassiveDns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPassiveDns()))
        for i, v := range m.GetPassiveDns() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("passiveDns", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPassiveDnsReverse() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPassiveDnsReverse()))
        for i, v := range m.GetPassiveDnsReverse() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("passiveDnsReverse", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reputation", m.GetReputation())
        if err != nil {
            return err
        }
    }
    if m.GetTrackers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTrackers()))
        for i, v := range m.GetTrackers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("trackers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComponents sets the components property value. The hostComponents that are associated with this host.
func (m *Host) SetComponents(value []HostComponentable)() {
    err := m.GetBackingStore().Set("components", value)
    if err != nil {
        panic(err)
    }
}
// SetCookies sets the cookies property value. The hostCookies that are associated with this host.
func (m *Host) SetCookies(value []HostCookieable)() {
    err := m.GetBackingStore().Set("cookies", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The first date and time this host was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *Host) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("firstSeenDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The most recent date and time when this host was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *Host) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSeenDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPassiveDns sets the passiveDns property value. Passive DNS retrieval about this host.
func (m *Host) SetPassiveDns(value []PassiveDnsRecordable)() {
    err := m.GetBackingStore().Set("passiveDns", value)
    if err != nil {
        panic(err)
    }
}
// SetPassiveDnsReverse sets the passiveDnsReverse property value. Reverse passive DNS retrieval about this host.
func (m *Host) SetPassiveDnsReverse(value []PassiveDnsRecordable)() {
    err := m.GetBackingStore().Set("passiveDnsReverse", value)
    if err != nil {
        panic(err)
    }
}
// SetReputation sets the reputation property value. Represents a calculated reputation of this host.
func (m *Host) SetReputation(value HostReputationable)() {
    err := m.GetBackingStore().Set("reputation", value)
    if err != nil {
        panic(err)
    }
}
// SetTrackers sets the trackers property value. The hostTrackers that are associated with this host.
func (m *Host) SetTrackers(value []HostTrackerable)() {
    err := m.GetBackingStore().Set("trackers", value)
    if err != nil {
        panic(err)
    }
}
// Hostable 
type Hostable interface {
    Artifactable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComponents()([]HostComponentable)
    GetCookies()([]HostCookieable)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPassiveDns()([]PassiveDnsRecordable)
    GetPassiveDnsReverse()([]PassiveDnsRecordable)
    GetReputation()(HostReputationable)
    GetTrackers()([]HostTrackerable)
    SetComponents(value []HostComponentable)()
    SetCookies(value []HostCookieable)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPassiveDns(value []PassiveDnsRecordable)()
    SetPassiveDnsReverse(value []PassiveDnsRecordable)()
    SetReputation(value HostReputationable)()
    SetTrackers(value []HostTrackerable)()
}
