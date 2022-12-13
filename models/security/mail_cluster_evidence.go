package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailClusterEvidence 
type MailClusterEvidence struct {
    AlertEvidence
    // The clusterBy property
    clusterBy *string
    // The clusterByValue property
    clusterByValue *string
    // The emailCount property
    emailCount *int64
    // The networkMessageIds property
    networkMessageIds []string
    // The query property
    query *string
    // The urn property
    urn *string
}
// NewMailClusterEvidence instantiates a new MailClusterEvidence and sets the default values.
func NewMailClusterEvidence()(*MailClusterEvidence) {
    m := &MailClusterEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateMailClusterEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailClusterEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailClusterEvidence(), nil
}
// GetClusterBy gets the clusterBy property value. The clusterBy property
func (m *MailClusterEvidence) GetClusterBy()(*string) {
    return m.clusterBy
}
// GetClusterByValue gets the clusterByValue property value. The clusterByValue property
func (m *MailClusterEvidence) GetClusterByValue()(*string) {
    return m.clusterByValue
}
// GetEmailCount gets the emailCount property value. The emailCount property
func (m *MailClusterEvidence) GetEmailCount()(*int64) {
    return m.emailCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailClusterEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["clusterBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterBy(val)
        }
        return nil
    }
    res["clusterByValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterByValue(val)
        }
        return nil
    }
    res["emailCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailCount(val)
        }
        return nil
    }
    res["networkMessageIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNetworkMessageIds(res)
        }
        return nil
    }
    res["query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    res["urn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrn(val)
        }
        return nil
    }
    return res
}
// GetNetworkMessageIds gets the networkMessageIds property value. The networkMessageIds property
func (m *MailClusterEvidence) GetNetworkMessageIds()([]string) {
    return m.networkMessageIds
}
// GetQuery gets the query property value. The query property
func (m *MailClusterEvidence) GetQuery()(*string) {
    return m.query
}
// GetUrn gets the urn property value. The urn property
func (m *MailClusterEvidence) GetUrn()(*string) {
    return m.urn
}
// Serialize serializes information the current object
func (m *MailClusterEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clusterBy", m.GetClusterBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clusterByValue", m.GetClusterByValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("emailCount", m.GetEmailCount())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkMessageIds() != nil {
        err = writer.WriteCollectionOfStringValues("networkMessageIds", m.GetNetworkMessageIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("urn", m.GetUrn())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClusterBy sets the clusterBy property value. The clusterBy property
func (m *MailClusterEvidence) SetClusterBy(value *string)() {
    m.clusterBy = value
}
// SetClusterByValue sets the clusterByValue property value. The clusterByValue property
func (m *MailClusterEvidence) SetClusterByValue(value *string)() {
    m.clusterByValue = value
}
// SetEmailCount sets the emailCount property value. The emailCount property
func (m *MailClusterEvidence) SetEmailCount(value *int64)() {
    m.emailCount = value
}
// SetNetworkMessageIds sets the networkMessageIds property value. The networkMessageIds property
func (m *MailClusterEvidence) SetNetworkMessageIds(value []string)() {
    m.networkMessageIds = value
}
// SetQuery sets the query property value. The query property
func (m *MailClusterEvidence) SetQuery(value *string)() {
    m.query = value
}
// SetUrn sets the urn property value. The urn property
func (m *MailClusterEvidence) SetUrn(value *string)() {
    m.urn = value
}
