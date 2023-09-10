package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ThreatIntelligence 
type ThreatIntelligence struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
}
// NewThreatIntelligence instantiates a new threatIntelligence and sets the default values.
func NewThreatIntelligence()(*ThreatIntelligence) {
    m := &ThreatIntelligence{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateThreatIntelligenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateThreatIntelligenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThreatIntelligence(), nil
}
// GetArticleIndicators gets the articleIndicators property value. Refers to indicators of threat or compromise highlighted in an microsoft.graph.security.article.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) GetArticleIndicators()([]ArticleIndicatorable) {
    val, err := m.GetBackingStore().Get("articleIndicators")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ArticleIndicatorable)
    }
    return nil
}
// GetArticles gets the articles property value. A list of article objects.
func (m *ThreatIntelligence) GetArticles()([]Articleable) {
    val, err := m.GetBackingStore().Get("articles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Articleable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThreatIntelligence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["articleIndicators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateArticleIndicatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ArticleIndicatorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ArticleIndicatorable)
                }
            }
            m.SetArticleIndicators(res)
        }
        return nil
    }
    res["articles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateArticleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Articleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Articleable)
                }
            }
            m.SetArticles(res)
        }
        return nil
    }
    res["hostComponents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetHostComponents(res)
        }
        return nil
    }
    res["hostCookies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetHostCookies(res)
        }
        return nil
    }
    res["hosts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Hostable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Hostable)
                }
            }
            m.SetHosts(res)
        }
        return nil
    }
    res["hostTrackers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetHostTrackers(res)
        }
        return nil
    }
    res["intelligenceProfileIndicators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntelligenceProfileIndicatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntelligenceProfileIndicatorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IntelligenceProfileIndicatorable)
                }
            }
            m.SetIntelligenceProfileIndicators(res)
        }
        return nil
    }
    res["intelProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntelligenceProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntelligenceProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IntelligenceProfileable)
                }
            }
            m.SetIntelProfiles(res)
        }
        return nil
    }
    res["passiveDnsRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPassiveDnsRecords(res)
        }
        return nil
    }
    res["vulnerabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVulnerabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Vulnerabilityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Vulnerabilityable)
                }
            }
            m.SetVulnerabilities(res)
        }
        return nil
    }
    return res
}
// GetHostComponents gets the hostComponents property value. Retrieve details about microsoft.graph.security.hostComponent objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) GetHostComponents()([]HostComponentable) {
    val, err := m.GetBackingStore().Get("hostComponents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HostComponentable)
    }
    return nil
}
// GetHostCookies gets the hostCookies property value. Retrieve details about microsoft.graph.security.hostCookie objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) GetHostCookies()([]HostCookieable) {
    val, err := m.GetBackingStore().Get("hostCookies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HostCookieable)
    }
    return nil
}
// GetHosts gets the hosts property value. Refers to microsoft.graph.security.host objects that Microsoft Threat Intelligence has observed.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) GetHosts()([]Hostable) {
    val, err := m.GetBackingStore().Get("hosts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Hostable)
    }
    return nil
}
// GetHostTrackers gets the hostTrackers property value. Retrieve details about microsoft.graph.security.hostTracker objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) GetHostTrackers()([]HostTrackerable) {
    val, err := m.GetBackingStore().Get("hostTrackers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HostTrackerable)
    }
    return nil
}
// GetIntelligenceProfileIndicators gets the intelligenceProfileIndicators property value. The intelligenceProfileIndicators property
func (m *ThreatIntelligence) GetIntelligenceProfileIndicators()([]IntelligenceProfileIndicatorable) {
    val, err := m.GetBackingStore().Get("intelligenceProfileIndicators")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IntelligenceProfileIndicatorable)
    }
    return nil
}
// GetIntelProfiles gets the intelProfiles property value. A list of intelligenceProfile objects.
func (m *ThreatIntelligence) GetIntelProfiles()([]IntelligenceProfileable) {
    val, err := m.GetBackingStore().Get("intelProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IntelligenceProfileable)
    }
    return nil
}
// GetPassiveDnsRecords gets the passiveDnsRecords property value. Retrieve details about microsoft.graph.security.passiveDnsRecord objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) GetPassiveDnsRecords()([]PassiveDnsRecordable) {
    val, err := m.GetBackingStore().Get("passiveDnsRecords")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PassiveDnsRecordable)
    }
    return nil
}
// GetVulnerabilities gets the vulnerabilities property value. Retrieve details about microsoft.graph.security.vulnerabilities.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) GetVulnerabilities()([]Vulnerabilityable) {
    val, err := m.GetBackingStore().Get("vulnerabilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Vulnerabilityable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ThreatIntelligence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetArticleIndicators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetArticleIndicators()))
        for i, v := range m.GetArticleIndicators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("articleIndicators", cast)
        if err != nil {
            return err
        }
    }
    if m.GetArticles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetArticles()))
        for i, v := range m.GetArticles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("articles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostComponents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostComponents()))
        for i, v := range m.GetHostComponents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostComponents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostCookies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostCookies()))
        for i, v := range m.GetHostCookies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostCookies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHosts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHosts()))
        for i, v := range m.GetHosts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hosts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostTrackers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostTrackers()))
        for i, v := range m.GetHostTrackers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostTrackers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntelligenceProfileIndicators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntelligenceProfileIndicators()))
        for i, v := range m.GetIntelligenceProfileIndicators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("intelligenceProfileIndicators", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntelProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntelProfiles()))
        for i, v := range m.GetIntelProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("intelProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPassiveDnsRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPassiveDnsRecords()))
        for i, v := range m.GetPassiveDnsRecords() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("passiveDnsRecords", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVulnerabilities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVulnerabilities()))
        for i, v := range m.GetVulnerabilities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("vulnerabilities", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArticleIndicators sets the articleIndicators property value. Refers to indicators of threat or compromise highlighted in an microsoft.graph.security.article.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetArticleIndicators(value []ArticleIndicatorable)() {
    err := m.GetBackingStore().Set("articleIndicators", value)
    if err != nil {
        panic(err)
    }
}
// SetArticles sets the articles property value. A list of article objects.
func (m *ThreatIntelligence) SetArticles(value []Articleable)() {
    err := m.GetBackingStore().Set("articles", value)
    if err != nil {
        panic(err)
    }
}
// SetHostComponents sets the hostComponents property value. Retrieve details about microsoft.graph.security.hostComponent objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHostComponents(value []HostComponentable)() {
    err := m.GetBackingStore().Set("hostComponents", value)
    if err != nil {
        panic(err)
    }
}
// SetHostCookies sets the hostCookies property value. Retrieve details about microsoft.graph.security.hostCookie objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHostCookies(value []HostCookieable)() {
    err := m.GetBackingStore().Set("hostCookies", value)
    if err != nil {
        panic(err)
    }
}
// SetHosts sets the hosts property value. Refers to microsoft.graph.security.host objects that Microsoft Threat Intelligence has observed.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHosts(value []Hostable)() {
    err := m.GetBackingStore().Set("hosts", value)
    if err != nil {
        panic(err)
    }
}
// SetHostTrackers sets the hostTrackers property value. Retrieve details about microsoft.graph.security.hostTracker objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHostTrackers(value []HostTrackerable)() {
    err := m.GetBackingStore().Set("hostTrackers", value)
    if err != nil {
        panic(err)
    }
}
// SetIntelligenceProfileIndicators sets the intelligenceProfileIndicators property value. The intelligenceProfileIndicators property
func (m *ThreatIntelligence) SetIntelligenceProfileIndicators(value []IntelligenceProfileIndicatorable)() {
    err := m.GetBackingStore().Set("intelligenceProfileIndicators", value)
    if err != nil {
        panic(err)
    }
}
// SetIntelProfiles sets the intelProfiles property value. A list of intelligenceProfile objects.
func (m *ThreatIntelligence) SetIntelProfiles(value []IntelligenceProfileable)() {
    err := m.GetBackingStore().Set("intelProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetPassiveDnsRecords sets the passiveDnsRecords property value. Retrieve details about microsoft.graph.security.passiveDnsRecord objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetPassiveDnsRecords(value []PassiveDnsRecordable)() {
    err := m.GetBackingStore().Set("passiveDnsRecords", value)
    if err != nil {
        panic(err)
    }
}
// SetVulnerabilities sets the vulnerabilities property value. Retrieve details about microsoft.graph.security.vulnerabilities.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetVulnerabilities(value []Vulnerabilityable)() {
    err := m.GetBackingStore().Set("vulnerabilities", value)
    if err != nil {
        panic(err)
    }
}
// ThreatIntelligenceable 
type ThreatIntelligenceable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArticleIndicators()([]ArticleIndicatorable)
    GetArticles()([]Articleable)
    GetHostComponents()([]HostComponentable)
    GetHostCookies()([]HostCookieable)
    GetHosts()([]Hostable)
    GetHostTrackers()([]HostTrackerable)
    GetIntelligenceProfileIndicators()([]IntelligenceProfileIndicatorable)
    GetIntelProfiles()([]IntelligenceProfileable)
    GetPassiveDnsRecords()([]PassiveDnsRecordable)
    GetVulnerabilities()([]Vulnerabilityable)
    SetArticleIndicators(value []ArticleIndicatorable)()
    SetArticles(value []Articleable)()
    SetHostComponents(value []HostComponentable)()
    SetHostCookies(value []HostCookieable)()
    SetHosts(value []Hostable)()
    SetHostTrackers(value []HostTrackerable)()
    SetIntelligenceProfileIndicators(value []IntelligenceProfileIndicatorable)()
    SetIntelProfiles(value []IntelligenceProfileable)()
    SetPassiveDnsRecords(value []PassiveDnsRecordable)()
    SetVulnerabilities(value []Vulnerabilityable)()
}
