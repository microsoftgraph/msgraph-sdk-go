package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Alert struct {
    Entity
    activityGroupName *string;
    alertDetections []AlertDetection;
    assignedTo *string;
    azureSubscriptionId *string;
    azureTenantId *string;
    category *string;
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    cloudAppStates []CloudAppSecurityState;
    comments []string;
    confidence *int32;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    description *string;
    detectionIds []string;
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    feedback *AlertFeedback;
    fileStates []FileSecurityState;
    historyStates []AlertHistoryState;
    hostStates []HostSecurityState;
    incidentIds []string;
    investigationSecurityStates []InvestigationSecurityState;
    lastEventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    malwareStates []MalwareState;
    messageSecurityStates []MessageSecurityState;
    networkConnections []NetworkConnection;
    processes []Process;
    recommendedActions []string;
    registryKeyStates []RegistryKeyState;
    securityResources []SecurityResource;
    severity *AlertSeverity;
    sourceMaterials []string;
    status *AlertStatus;
    tags []string;
    title *string;
    triggers []AlertTrigger;
    uriClickSecurityStates []UriClickSecurityState;
    userStates []UserSecurityState;
    vendorInformation *SecurityVendorInformation;
    vulnerabilityStates []VulnerabilityState;
}
func NewAlert()(*Alert) {
    m := &Alert{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Alert) GetActivityGroupName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityGroupName
    }
}
func (m *Alert) GetAlertDetections()([]AlertDetection) {
    if m == nil {
        return nil
    } else {
        return m.alertDetections
    }
}
func (m *Alert) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
func (m *Alert) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
func (m *Alert) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
func (m *Alert) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
func (m *Alert) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.closedDateTime
    }
}
func (m *Alert) GetCloudAppStates()([]CloudAppSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.cloudAppStates
    }
}
func (m *Alert) GetComments()([]string) {
    if m == nil {
        return nil
    } else {
        return m.comments
    }
}
func (m *Alert) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
func (m *Alert) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *Alert) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *Alert) GetDetectionIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.detectionIds
    }
}
func (m *Alert) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
func (m *Alert) GetFeedback()(*AlertFeedback) {
    if m == nil {
        return nil
    } else {
        return m.feedback
    }
}
func (m *Alert) GetFileStates()([]FileSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.fileStates
    }
}
func (m *Alert) GetHistoryStates()([]AlertHistoryState) {
    if m == nil {
        return nil
    } else {
        return m.historyStates
    }
}
func (m *Alert) GetHostStates()([]HostSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.hostStates
    }
}
func (m *Alert) GetIncidentIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.incidentIds
    }
}
func (m *Alert) GetInvestigationSecurityStates()([]InvestigationSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.investigationSecurityStates
    }
}
func (m *Alert) GetLastEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastEventDateTime
    }
}
func (m *Alert) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *Alert) GetMalwareStates()([]MalwareState) {
    if m == nil {
        return nil
    } else {
        return m.malwareStates
    }
}
func (m *Alert) GetMessageSecurityStates()([]MessageSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.messageSecurityStates
    }
}
func (m *Alert) GetNetworkConnections()([]NetworkConnection) {
    if m == nil {
        return nil
    } else {
        return m.networkConnections
    }
}
func (m *Alert) GetProcesses()([]Process) {
    if m == nil {
        return nil
    } else {
        return m.processes
    }
}
func (m *Alert) GetRecommendedActions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedActions
    }
}
func (m *Alert) GetRegistryKeyStates()([]RegistryKeyState) {
    if m == nil {
        return nil
    } else {
        return m.registryKeyStates
    }
}
func (m *Alert) GetSecurityResources()([]SecurityResource) {
    if m == nil {
        return nil
    } else {
        return m.securityResources
    }
}
func (m *Alert) GetSeverity()(*AlertSeverity) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
func (m *Alert) GetSourceMaterials()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sourceMaterials
    }
}
func (m *Alert) GetStatus()(*AlertStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *Alert) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
func (m *Alert) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
func (m *Alert) GetTriggers()([]AlertTrigger) {
    if m == nil {
        return nil
    } else {
        return m.triggers
    }
}
func (m *Alert) GetUriClickSecurityStates()([]UriClickSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.uriClickSecurityStates
    }
}
func (m *Alert) GetUserStates()([]UserSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.userStates
    }
}
func (m *Alert) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
func (m *Alert) GetVulnerabilityStates()([]VulnerabilityState) {
    if m == nil {
        return nil
    } else {
        return m.vulnerabilityStates
    }
}
func (m *Alert) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityGroupName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivityGroupName(val)
        return nil
    }
    res["alertDetections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlertDetection() })
        if err != nil {
            return err
        }
        res := make([]AlertDetection, len(val))
        for i, v := range val {
            res[i] = *(v.(*AlertDetection))
        }
        m.SetAlertDetections(res)
        return nil
    }
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignedTo(val)
        return nil
    }
    res["azureSubscriptionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureSubscriptionId(val)
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureTenantId(val)
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory(val)
        return nil
    }
    res["closedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetClosedDateTime(val)
        return nil
    }
    res["cloudAppStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudAppSecurityState() })
        if err != nil {
            return err
        }
        res := make([]CloudAppSecurityState, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudAppSecurityState))
        }
        m.SetCloudAppStates(res)
        return nil
    }
    res["comments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetComments(res)
        return nil
    }
    res["confidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfidence(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["detectionIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDetectionIds(res)
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEventDateTime(val)
        return nil
    }
    res["feedback"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertFeedback)
        if err != nil {
            return err
        }
        cast := val.(AlertFeedback)
        m.SetFeedback(&cast)
        return nil
    }
    res["fileStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileSecurityState() })
        if err != nil {
            return err
        }
        res := make([]FileSecurityState, len(val))
        for i, v := range val {
            res[i] = *(v.(*FileSecurityState))
        }
        m.SetFileStates(res)
        return nil
    }
    res["historyStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlertHistoryState() })
        if err != nil {
            return err
        }
        res := make([]AlertHistoryState, len(val))
        for i, v := range val {
            res[i] = *(v.(*AlertHistoryState))
        }
        m.SetHistoryStates(res)
        return nil
    }
    res["hostStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHostSecurityState() })
        if err != nil {
            return err
        }
        res := make([]HostSecurityState, len(val))
        for i, v := range val {
            res[i] = *(v.(*HostSecurityState))
        }
        m.SetHostStates(res)
        return nil
    }
    res["incidentIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncidentIds(res)
        return nil
    }
    res["investigationSecurityStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInvestigationSecurityState() })
        if err != nil {
            return err
        }
        res := make([]InvestigationSecurityState, len(val))
        for i, v := range val {
            res[i] = *(v.(*InvestigationSecurityState))
        }
        m.SetInvestigationSecurityStates(res)
        return nil
    }
    res["lastEventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastEventDateTime(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["malwareStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMalwareState() })
        if err != nil {
            return err
        }
        res := make([]MalwareState, len(val))
        for i, v := range val {
            res[i] = *(v.(*MalwareState))
        }
        m.SetMalwareStates(res)
        return nil
    }
    res["messageSecurityStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessageSecurityState() })
        if err != nil {
            return err
        }
        res := make([]MessageSecurityState, len(val))
        for i, v := range val {
            res[i] = *(v.(*MessageSecurityState))
        }
        m.SetMessageSecurityStates(res)
        return nil
    }
    res["networkConnections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNetworkConnection() })
        if err != nil {
            return err
        }
        res := make([]NetworkConnection, len(val))
        for i, v := range val {
            res[i] = *(v.(*NetworkConnection))
        }
        m.SetNetworkConnections(res)
        return nil
    }
    res["processes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProcess() })
        if err != nil {
            return err
        }
        res := make([]Process, len(val))
        for i, v := range val {
            res[i] = *(v.(*Process))
        }
        m.SetProcesses(res)
        return nil
    }
    res["recommendedActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRecommendedActions(res)
        return nil
    }
    res["registryKeyStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRegistryKeyState() })
        if err != nil {
            return err
        }
        res := make([]RegistryKeyState, len(val))
        for i, v := range val {
            res[i] = *(v.(*RegistryKeyState))
        }
        m.SetRegistryKeyStates(res)
        return nil
    }
    res["securityResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityResource() })
        if err != nil {
            return err
        }
        res := make([]SecurityResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecurityResource))
        }
        m.SetSecurityResources(res)
        return nil
    }
    res["severity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertSeverity)
        if err != nil {
            return err
        }
        cast := val.(AlertSeverity)
        m.SetSeverity(&cast)
        return nil
    }
    res["sourceMaterials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSourceMaterials(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertStatus)
        if err != nil {
            return err
        }
        cast := val.(AlertStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTags(res)
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
        return nil
    }
    res["triggers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlertTrigger() })
        if err != nil {
            return err
        }
        res := make([]AlertTrigger, len(val))
        for i, v := range val {
            res[i] = *(v.(*AlertTrigger))
        }
        m.SetTriggers(res)
        return nil
    }
    res["uriClickSecurityStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUriClickSecurityState() })
        if err != nil {
            return err
        }
        res := make([]UriClickSecurityState, len(val))
        for i, v := range val {
            res[i] = *(v.(*UriClickSecurityState))
        }
        m.SetUriClickSecurityStates(res)
        return nil
    }
    res["userStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSecurityState() })
        if err != nil {
            return err
        }
        res := make([]UserSecurityState, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSecurityState))
        }
        m.SetUserStates(res)
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        m.SetVendorInformation(val.(*SecurityVendorInformation))
        return nil
    }
    res["vulnerabilityStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVulnerabilityState() })
        if err != nil {
            return err
        }
        res := make([]VulnerabilityState, len(val))
        for i, v := range val {
            res[i] = *(v.(*VulnerabilityState))
        }
        m.SetVulnerabilityStates(res)
        return nil
    }
    return res
}
func (m *Alert) IsNil()(bool) {
    return m == nil
}
func (m *Alert) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activityGroupName", m.GetActivityGroupName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlertDetections()))
        for i, v := range m.GetAlertDetections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("alertDetections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureSubscriptionId", m.GetAzureSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("closedDateTime", m.GetClosedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCloudAppStates()))
        for i, v := range m.GetCloudAppStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cloudAppStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("detectionIds", m.GetDetectionIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetFeedback() != nil {
        cast := m.GetFeedback().String()
        err = writer.WriteStringValue("feedback", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFileStates()))
        for i, v := range m.GetFileStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("fileStates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistoryStates()))
        for i, v := range m.GetHistoryStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("historyStates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHostStates()))
        for i, v := range m.GetHostStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("hostStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("incidentIds", m.GetIncidentIds())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInvestigationSecurityStates()))
        for i, v := range m.GetInvestigationSecurityStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("investigationSecurityStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastEventDateTime", m.GetLastEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMalwareStates()))
        for i, v := range m.GetMalwareStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("malwareStates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessageSecurityStates()))
        for i, v := range m.GetMessageSecurityStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("messageSecurityStates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNetworkConnections()))
        for i, v := range m.GetNetworkConnections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("networkConnections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProcesses()))
        for i, v := range m.GetProcesses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("processes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("recommendedActions", m.GetRecommendedActions())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRegistryKeyStates()))
        for i, v := range m.GetRegistryKeyStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("registryKeyStates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecurityResources()))
        for i, v := range m.GetSecurityResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("securityResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := m.GetSeverity().String()
        err = writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("sourceMaterials", m.GetSourceMaterials())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTriggers()))
        for i, v := range m.GetTriggers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("triggers", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUriClickSecurityStates()))
        for i, v := range m.GetUriClickSecurityStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("uriClickSecurityStates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStates()))
        for i, v := range m.GetUserStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVulnerabilityStates()))
        for i, v := range m.GetVulnerabilityStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("vulnerabilityStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Alert) SetActivityGroupName(value *string)() {
    m.activityGroupName = value
}
func (m *Alert) SetAlertDetections(value []AlertDetection)() {
    m.alertDetections = value
}
func (m *Alert) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
func (m *Alert) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
func (m *Alert) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
func (m *Alert) SetCategory(value *string)() {
    m.category = value
}
func (m *Alert) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closedDateTime = value
}
func (m *Alert) SetCloudAppStates(value []CloudAppSecurityState)() {
    m.cloudAppStates = value
}
func (m *Alert) SetComments(value []string)() {
    m.comments = value
}
func (m *Alert) SetConfidence(value *int32)() {
    m.confidence = value
}
func (m *Alert) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *Alert) SetDescription(value *string)() {
    m.description = value
}
func (m *Alert) SetDetectionIds(value []string)() {
    m.detectionIds = value
}
func (m *Alert) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
func (m *Alert) SetFeedback(value *AlertFeedback)() {
    m.feedback = value
}
func (m *Alert) SetFileStates(value []FileSecurityState)() {
    m.fileStates = value
}
func (m *Alert) SetHistoryStates(value []AlertHistoryState)() {
    m.historyStates = value
}
func (m *Alert) SetHostStates(value []HostSecurityState)() {
    m.hostStates = value
}
func (m *Alert) SetIncidentIds(value []string)() {
    m.incidentIds = value
}
func (m *Alert) SetInvestigationSecurityStates(value []InvestigationSecurityState)() {
    m.investigationSecurityStates = value
}
func (m *Alert) SetLastEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastEventDateTime = value
}
func (m *Alert) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *Alert) SetMalwareStates(value []MalwareState)() {
    m.malwareStates = value
}
func (m *Alert) SetMessageSecurityStates(value []MessageSecurityState)() {
    m.messageSecurityStates = value
}
func (m *Alert) SetNetworkConnections(value []NetworkConnection)() {
    m.networkConnections = value
}
func (m *Alert) SetProcesses(value []Process)() {
    m.processes = value
}
func (m *Alert) SetRecommendedActions(value []string)() {
    m.recommendedActions = value
}
func (m *Alert) SetRegistryKeyStates(value []RegistryKeyState)() {
    m.registryKeyStates = value
}
func (m *Alert) SetSecurityResources(value []SecurityResource)() {
    m.securityResources = value
}
func (m *Alert) SetSeverity(value *AlertSeverity)() {
    m.severity = value
}
func (m *Alert) SetSourceMaterials(value []string)() {
    m.sourceMaterials = value
}
func (m *Alert) SetStatus(value *AlertStatus)() {
    m.status = value
}
func (m *Alert) SetTags(value []string)() {
    m.tags = value
}
func (m *Alert) SetTitle(value *string)() {
    m.title = value
}
func (m *Alert) SetTriggers(value []AlertTrigger)() {
    m.triggers = value
}
func (m *Alert) SetUriClickSecurityStates(value []UriClickSecurityState)() {
    m.uriClickSecurityStates = value
}
func (m *Alert) SetUserStates(value []UserSecurityState)() {
    m.userStates = value
}
func (m *Alert) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}
func (m *Alert) SetVulnerabilityStates(value []VulnerabilityState)() {
    m.vulnerabilityStates = value
}
