package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i005f0821c1e57467ed9e31e0f65ebf8795241942e8afc0c6c9cc866dc3b70de2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/owneddevices"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i065d0a279baf0cb315c26855c2ba7a5c0701939d8cae543a4fb9b7bb2c7cbda5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/licensedetails"
    i0a25da984a2612b28dedf3c1e18748fd27dd92433244095b38a29e6ec842f917 "github.com/microsoftgraph/msgraph-sdk-go/users/item/checkmemberobjects"
    i0a4c7a84985f2276fcfbcf69e2fe37b80ef87c2622cc8763371022616974902d "github.com/microsoftgraph/msgraph-sdk-go/users/item/presence"
    i0a6d7999baeed47d1aa02960cd9f3f120528ecae4240bf52300542a8dc9f978d "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
    i1090e84ad9b76c4c279467ebbc90ff155a540b45b098e680706d2aae300639eb "github.com/microsoftgraph/msgraph-sdk-go/users/item/revokesigninsessions"
    i11db6285564609d0b36b9032e0d8f6d257828b887e558375da4354121f40a3d0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof"
    i152c9659f52d4d76fff6d586e3de05874eee179419588197c21a3adc1e3a6c7b "github.com/microsoftgraph/msgraph-sdk-go/users/item/ownedobjects"
    i16f13fb20dc21e9cbd93cbca432ac5cb871c47a97e5bebce67f5eab6fe8276e9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/sendmail"
    i225d659ae554743be8acfa963ab6384d1f6e82911866fa2e3c01918b6a0637b1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/settings"
    i23845c92400f42d2b469653fba5fe1555d60b50e4596b30d2d2b4c127b0ceecf "github.com/microsoftgraph/msgraph-sdk-go/users/item/getmanagedapppolicies"
    i2b4d48c4393811d79f4d4cc1f4addb3b49f58a53edcf7d1d7140e29a22ab7504 "github.com/microsoftgraph/msgraph-sdk-go/users/item/people"
    i2c0f098b49b5f2edba81f4095f0c699516263b7cd6ff316d31f469f15b8366b0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onlinemeetings"
    i2ec0496e560d7cecb96d82bcf32b4c61894a21cfd165c378253939195c546d27 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars"
    i3a4cc4e91af7383718bf555760e1dba5c1fc1de3509837d9074d3a3f0a8a0fc5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drive"
    i3fd2d5130e27e5ed481f7a0565ee17271d52e0ea15f0dec8ae436a3780d1d8d2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manager"
    i418d77dc601ca4dc2c568aef8702ce31b38463f4e1701c6c788adcbeb32ca3e8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/outlook"
    i4a474d43148dd96e52e36b3ed7c52a7d75f46d95adf4fa0ce96ace2509d33744 "github.com/microsoftgraph/msgraph-sdk-go/users/item/planner"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4af8697d288ced3f62e5f7467084c929a355a7f4c2140a3f51072b00f73663c1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/wipemanagedappregistrationsbydevicetag"
    i4cb75aef1c5547ab5604eb9b1721354472df7129995c61c066ac036f69c410c1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/createdobjects"
    i54c28e8ec50e0dd4eeee8073ed79cecee11ba12e3813ec711d8127af3d4b994c "github.com/microsoftgraph/msgraph-sdk-go/users/item/getmembergroups"
    i61d84fd6521641b94f790414a5b9abccf870ea66012c270c06f3ea7df770aad8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/todo"
    i631d94a7300d4142a2d7a4c4d0d424e98d3fb058f106e6567a8ac6476f5956b6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/registereddevices"
    i64f47462143d9c5d6613e409adca3c0fd24631de22e340968821ffdff467e657 "github.com/microsoftgraph/msgraph-sdk-go/users/item/inferenceclassification"
    i6d278e5cb359938f2362c090cc322d0c80867015903a39849ea478a6cda51689 "github.com/microsoftgraph/msgraph-sdk-go/users/item/removealldevicesfrommanagement"
    i6d5ec618068b4f9940966e3235a6b56547b161354e35247fc9843cd2fc68c280 "github.com/microsoftgraph/msgraph-sdk-go/users/item/getmailtips"
    i6fea53c850cb6db38ea2669df26d5c9986df82967fc9dcbd1415ef4c37ba3909 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives"
    i7434742f328fb34181df9069343c7bf11d068d61f4481ab2eaea71f3bc398a8c "github.com/microsoftgraph/msgraph-sdk-go/users/item/activities"
    i79cfaa5f7496a3faeffd9cee89b05478df04194a9746fc39a36637ad15ede76b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar"
    i7f31fc68d418ebfd772842294b63ed5b4dafc3cab499c7bf25b7634a901656bb "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams"
    i8acf48036fa30b3733f0a3efffd923a945db096ce1c005ede1236ec3db8f202b "github.com/microsoftgraph/msgraph-sdk-go/users/item/approleassignments"
    i8b015d6ac3b9f998d8da5aa503d326b952acdc90c4c90a03d48fdec51c4588f1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/devicemanagementtroubleshootingevents"
    i8deae6de37c6b836f34cd5c6b8fab91a73043c26ee346f7410a8dcf88ea64c99 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights"
    i8fdd3e26eeb6da18c0b9d487ac43983355a059891790fba9c732ef59d5c4f21b "github.com/microsoftgraph/msgraph-sdk-go/users/item/findmeetingtimes"
    i9164851f0a06aba10b88da1bbecb98e3600356d368c9db499b8931c4f96f1219 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contacts"
    i91af54e73e80a81d2f98c6f72fbb6c81539e67ba0480227ae8e2e5f2b8f467fa "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices"
    i91bcb6ebe27791a826a51ecff16cc35355876378f254587360afeb0054d6f984 "github.com/microsoftgraph/msgraph-sdk-go/users/item/checkmembergroups"
    i927d56fd90d9c38a0e88ef5c682feb77cc18445308d44276e0f1e095247b3634 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups"
    i94ff5ee566023b18b8b2ac8c1da92081f3bf4520205017599c00aed9faff56de "github.com/microsoftgraph/msgraph-sdk-go/users/item/teamwork"
    i9f87171771d54744395d1043beb42c14430d7cb4f1557826ebef7c5b2584f7cd "github.com/microsoftgraph/msgraph-sdk-go/users/item/restore"
    ia6a0f4f70d60f824df18f07f8918c6e9ac043f588293118e5d91eb921d5dffbf "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof"
    ia99f649e03b3020696f5c2c11f42a9ad7f094cb80add889863fc7947d3bf9ec7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/scopedrolememberof"
    ib8e8ec0d09197f1fd57167b1b67da9bcc229fa24977921ce6406f24aa89a6884 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote"
    ib946e3869f342130dddab8c76c89f264a76e899f07657349241151b4e41b4286 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
    ic31c6bdf1febcc1619240151a993f96448ebdf3c6f91fe0243fbe55f91e7f620 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders"
    icf27f39169518adb5f7aa45892636b30915e5166590a8d55b561e502c71dc5fe "github.com/microsoftgraph/msgraph-sdk-go/users/item/changepassword"
    icf56cc7fcccc9803d5655a975f93bee1c76d14d27e2e180c40ef7bb1c3c18a68 "github.com/microsoftgraph/msgraph-sdk-go/users/item/photo"
    id7592a93402c0d67508da968e58d07109c9c2561ae838fbeb1cccb27ef558ecb "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports"
    id9510e26530c8ba3b39834c060eb72230844f30169f4e2b2e6b683eb86727fa0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats"
    idc22dd282127655d723afbe2b6bcc58326fe6d9a8b1bcdee2586b77850b235b0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/agreementacceptances"
    iddc156b6ba287ee33d5592797e8da64087fd068cf179bd892944253be67ac174 "github.com/microsoftgraph/msgraph-sdk-go/users/item/reminderviewwithstartdatetimewithenddatetime"
    ide0a33943859ca09b2034a1804f60d32a108dc09bf7012d68e5d3c63ac617ddb "github.com/microsoftgraph/msgraph-sdk-go/users/item/getmemberobjects"
    ide2d4144e5fd4f180ce2f7d5791881ad155e598ebbe5bda78dbfc5f9cd8aa8d0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication"
    idea2a9c81deaaa3e33a2bcef2386387935820f4f571d71af1fc2daa24d8fa032 "github.com/microsoftgraph/msgraph-sdk-go/users/item/photos"
    idfea0d624d269e60a96eeda8949da3675c3ceb484b3fb16f118382ba0a7d3583 "github.com/microsoftgraph/msgraph-sdk-go/users/item/assignlicense"
    iea339692a4ece4755a03d26004b71e7bbfe3ad28cf8036bdc4b972512eb00359 "github.com/microsoftgraph/msgraph-sdk-go/users/item/oauth2permissiongrants"
    ieb873ea17846be13ef65aaf00090f933a16727baa2e8063b39c2d86ae07e1ad0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/getmanagedappdiagnosticstatuses"
    iedc2060488fc2f6e7d120554e9e853117ec2aab876ef464c9ed3da4696541a32 "github.com/microsoftgraph/msgraph-sdk-go/users/item/exportpersonaldata"
    if0c7df43a05145fae76c1ab18483108fb4043f7b6a0b01f47772fd66022550ee "github.com/microsoftgraph/msgraph-sdk-go/users/item/extensions"
    if12f3532feee7012e3a4dc1e3a2444ca8fab8f29e8128baafe0ef00f6ed6015c "github.com/microsoftgraph/msgraph-sdk-go/users/item/managedappregistrations"
    if19172f362f2e1d4456e041244db6eae76a05acf4f16807f05721a5ef0a0f4be "github.com/microsoftgraph/msgraph-sdk-go/users/item/followedsites"
    ifd940c65ea13f1270f41f753ee5e7761f074915bd6a6c64b2adb77f70aea8c32 "github.com/microsoftgraph/msgraph-sdk-go/users/item/reprocesslicenseassignment"
    ifd9b701a6f88aa459780777b5ac521fe3494f01b9276c6884b86827296582e54 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events"
    ifdf31129bec8038d82607a6d973830e0b090629373d5812152e5046514a4bfb0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/translateexchangeids"
    ifec6c3ff45ea43d5b46ee636228680b7e4b3bfb9eba0437938b28dad5cadcd06 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview"
    i086f912938edbc37b845789cda2e3e84c9a478d06edbee6d535f8e143d4c70ba "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item"
    i0b4dff6b82c5c98b58b9ff11aead66a7e49d0eac6b18b1490fd01bf9597cf36c "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item"
    i0be6421da86ae0eb5550fa2bf3e3bb1ee5c6cb82a7b6539fce14162126c5d994 "github.com/microsoftgraph/msgraph-sdk-go/users/item/devicemanagementtroubleshootingevents/item"
    i0c6b3d37dc64615f253a6fd9a58f59f536aa93a57f99bb5f9d4210ae83aa5112 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item"
    i129160b87cb78d35acdd1a84c7328c6872f51bf17cb02a211ba957a5431db64b "github.com/microsoftgraph/msgraph-sdk-go/users/item/extensions/item"
    i2083055461886ca3b56f8adc26419c83b6c45f08c3a0eb4fcdfa94b7e99c9bef "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item"
    i3ecd250b638acba95f41b3db07eaa507842a51f8e0cd3640cb9b75f0052ceff6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onlinemeetings/item"
    i4d1b56e56093574975d1ab3f4a52af1d6db962510708352cb01245aa891a6ef5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats/item"
    i5cf53dd103ab02b1568c2f84b946a72977e41e0de97287af6722ccafaa880d49 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item"
    i5f2b6c8998555350079828979727c205406cf7155e907008bec42440c3e9724a "github.com/microsoftgraph/msgraph-sdk-go/users/item/photos/item"
    i764f7513465c0251a3785c28eac14dab3642e586e0fad67528f13704ed1cbb64 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item"
    i8465a2d5f26666d08a0baebc755ef48f1f0dc18551f929c58c09d0110510a317 "github.com/microsoftgraph/msgraph-sdk-go/users/item/people/item"
    i88e2c30b78a734ede88b94c583c45b04ce07232920ad9d7b2366b4278bfae360 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contacts/item"
    i983bdc3dce2d947aec82e1e9dd5edd887e8d861d589287d050f96af67abb1e1f "github.com/microsoftgraph/msgraph-sdk-go/users/item/scopedrolememberof/item"
    ibbdc262c9fd17ec4015663fbf467ebc865f1a6a45daf1ffbaa73c799220a610b "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item"
    id1f5db1c2220a25bede6871151a00a067652bdb75e2db4d199ca6ffcdd70b342 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item"
    id307eb3180e14c2c4cec59b4b85564ae282ebaa15d7f7735716952f69e678279 "github.com/microsoftgraph/msgraph-sdk-go/users/item/approleassignments/item"
    if2fde6f9b7cf9a66adc51b4c84ec599a0497b6792379909cd577d454bd9d0414 "github.com/microsoftgraph/msgraph-sdk-go/users/item/activities/item"
    if33563b345c020375260161400a6b809f00973f17d5ab870445f3c67ca7323b8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item"
    ifdc39e295bcc2e19fecc18f2ae737ddb8fe9018d8753e81ef0c2de8bb37638da "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item"
)

// UserRequestBuilder builds and executes requests for operations under \users\{user-id}
type UserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserRequestBuilderDeleteOptions options for Delete
type UserRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserRequestBuilderGetOptions options for Get
type UserRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserRequestBuilderGetQueryParameters get entity from users by key
type UserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// UserRequestBuilderPatchOptions options for Patch
type UserRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UserRequestBuilder) Activities()(*i7434742f328fb34181df9069343c7bf11d068d61f4481ab2eaea71f3bc398a8c.ActivitiesRequestBuilder) {
    return i7434742f328fb34181df9069343c7bf11d068d61f4481ab2eaea71f3bc398a8c.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.activities.item collection
func (m *UserRequestBuilder) ActivitiesById(id string)(*if2fde6f9b7cf9a66adc51b4c84ec599a0497b6792379909cd577d454bd9d0414.UserActivityRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity_id"] = id
    }
    return if2fde6f9b7cf9a66adc51b4c84ec599a0497b6792379909cd577d454bd9d0414.NewUserActivityRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) AgreementAcceptances()(*idc22dd282127655d723afbe2b6bcc58326fe6d9a8b1bcdee2586b77850b235b0.AgreementAcceptancesRequestBuilder) {
    return idc22dd282127655d723afbe2b6bcc58326fe6d9a8b1bcdee2586b77850b235b0.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) AppRoleAssignments()(*i8acf48036fa30b3733f0a3efffd923a945db096ce1c005ede1236ec3db8f202b.AppRoleAssignmentsRequestBuilder) {
    return i8acf48036fa30b3733f0a3efffd923a945db096ce1c005ede1236ec3db8f202b.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.appRoleAssignments.item collection
func (m *UserRequestBuilder) AppRoleAssignmentsById(id string)(*id307eb3180e14c2c4cec59b4b85564ae282ebaa15d7f7735716952f69e678279.AppRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return id307eb3180e14c2c4cec59b4b85564ae282ebaa15d7f7735716952f69e678279.NewAppRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) AssignLicense()(*idfea0d624d269e60a96eeda8949da3675c3ceb484b3fb16f118382ba0a7d3583.AssignLicenseRequestBuilder) {
    return idfea0d624d269e60a96eeda8949da3675c3ceb484b3fb16f118382ba0a7d3583.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Authentication()(*ide2d4144e5fd4f180ce2f7d5791881ad155e598ebbe5bda78dbfc5f9cd8aa8d0.AuthenticationRequestBuilder) {
    return ide2d4144e5fd4f180ce2f7d5791881ad155e598ebbe5bda78dbfc5f9cd8aa8d0.NewAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Calendar()(*i79cfaa5f7496a3faeffd9cee89b05478df04194a9746fc39a36637ad15ede76b.CalendarRequestBuilder) {
    return i79cfaa5f7496a3faeffd9cee89b05478df04194a9746fc39a36637ad15ede76b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) CalendarGroups()(*i927d56fd90d9c38a0e88ef5c682feb77cc18445308d44276e0f1e095247b3634.CalendarGroupsRequestBuilder) {
    return i927d56fd90d9c38a0e88ef5c682feb77cc18445308d44276e0f1e095247b3634.NewCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item collection
func (m *UserRequestBuilder) CalendarGroupsById(id string)(*i2083055461886ca3b56f8adc26419c83b6c45f08c3a0eb4fcdfa94b7e99c9bef.CalendarGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup_id"] = id
    }
    return i2083055461886ca3b56f8adc26419c83b6c45f08c3a0eb4fcdfa94b7e99c9bef.NewCalendarGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Calendars()(*i2ec0496e560d7cecb96d82bcf32b4c61894a21cfd165c378253939195c546d27.CalendarsRequestBuilder) {
    return i2ec0496e560d7cecb96d82bcf32b4c61894a21cfd165c378253939195c546d27.NewCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item collection
func (m *UserRequestBuilder) CalendarsById(id string)(*i086f912938edbc37b845789cda2e3e84c9a478d06edbee6d535f8e143d4c70ba.CalendarRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar_id"] = id
    }
    return i086f912938edbc37b845789cda2e3e84c9a478d06edbee6d535f8e143d4c70ba.NewCalendarRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) CalendarView()(*ifec6c3ff45ea43d5b46ee636228680b7e4b3bfb9eba0437938b28dad5cadcd06.CalendarViewRequestBuilder) {
    return ifec6c3ff45ea43d5b46ee636228680b7e4b3bfb9eba0437938b28dad5cadcd06.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item collection
func (m *UserRequestBuilder) CalendarViewById(id string)(*i5cf53dd103ab02b1568c2f84b946a72977e41e0de97287af6722ccafaa880d49.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i5cf53dd103ab02b1568c2f84b946a72977e41e0de97287af6722ccafaa880d49.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) ChangePassword()(*icf27f39169518adb5f7aa45892636b30915e5166590a8d55b561e502c71dc5fe.ChangePasswordRequestBuilder) {
    return icf27f39169518adb5f7aa45892636b30915e5166590a8d55b561e502c71dc5fe.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Chats()(*id9510e26530c8ba3b39834c060eb72230844f30169f4e2b2e6b683eb86727fa0.ChatsRequestBuilder) {
    return id9510e26530c8ba3b39834c060eb72230844f30169f4e2b2e6b683eb86727fa0.NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.chats.item collection
func (m *UserRequestBuilder) ChatsById(id string)(*i4d1b56e56093574975d1ab3f4a52af1d6db962510708352cb01245aa891a6ef5.ChatRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat_id"] = id
    }
    return i4d1b56e56093574975d1ab3f4a52af1d6db962510708352cb01245aa891a6ef5.NewChatRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) CheckMemberGroups()(*i91bcb6ebe27791a826a51ecff16cc35355876378f254587360afeb0054d6f984.CheckMemberGroupsRequestBuilder) {
    return i91bcb6ebe27791a826a51ecff16cc35355876378f254587360afeb0054d6f984.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) CheckMemberObjects()(*i0a25da984a2612b28dedf3c1e18748fd27dd92433244095b38a29e6ec842f917.CheckMemberObjectsRequestBuilder) {
    return i0a25da984a2612b28dedf3c1e18748fd27dd92433244095b38a29e6ec842f917.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserRequestBuilder instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UserRequestBuilder) ContactFolders()(*ic31c6bdf1febcc1619240151a993f96448ebdf3c6f91fe0243fbe55f91e7f620.ContactFoldersRequestBuilder) {
    return ic31c6bdf1febcc1619240151a993f96448ebdf3c6f91fe0243fbe55f91e7f620.NewContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactFoldersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.contactFolders.item collection
func (m *UserRequestBuilder) ContactFoldersById(id string)(*ibbdc262c9fd17ec4015663fbf467ebc865f1a6a45daf1ffbaa73c799220a610b.ContactFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder_id"] = id
    }
    return ibbdc262c9fd17ec4015663fbf467ebc865f1a6a45daf1ffbaa73c799220a610b.NewContactFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Contacts()(*i9164851f0a06aba10b88da1bbecb98e3600356d368c9db499b8931c4f96f1219.ContactsRequestBuilder) {
    return i9164851f0a06aba10b88da1bbecb98e3600356d368c9db499b8931c4f96f1219.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.contacts.item collection
func (m *UserRequestBuilder) ContactsById(id string)(*i88e2c30b78a734ede88b94c583c45b04ce07232920ad9d7b2366b4278bfae360.ContactRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return i88e2c30b78a734ede88b94c583c45b04ce07232920ad9d7b2366b4278bfae360.NewContactRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete entity from users
func (m *UserRequestBuilder) CreateDeleteRequestInformation(options *UserRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *UserRequestBuilder) CreatedObjects()(*i4cb75aef1c5547ab5604eb9b1721354472df7129995c61c066ac036f69c410c1.CreatedObjectsRequestBuilder) {
    return i4cb75aef1c5547ab5604eb9b1721354472df7129995c61c066ac036f69c410c1.NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entity from users by key
func (m *UserRequestBuilder) CreateGetRequestInformation(options *UserRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in users
func (m *UserRequestBuilder) CreatePatchRequestInformation(options *UserRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from users
func (m *UserRequestBuilder) Delete(options *UserRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *UserRequestBuilder) DeviceManagementTroubleshootingEvents()(*i8b015d6ac3b9f998d8da5aa503d326b952acdc90c4c90a03d48fdec51c4588f1.DeviceManagementTroubleshootingEventsRequestBuilder) {
    return i8b015d6ac3b9f998d8da5aa503d326b952acdc90c4c90a03d48fdec51c4588f1.NewDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementTroubleshootingEventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.deviceManagementTroubleshootingEvents.item collection
func (m *UserRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*i0be6421da86ae0eb5550fa2bf3e3bb1ee5c6cb82a7b6539fce14162126c5d994.DeviceManagementTroubleshootingEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent_id"] = id
    }
    return i0be6421da86ae0eb5550fa2bf3e3bb1ee5c6cb82a7b6539fce14162126c5d994.NewDeviceManagementTroubleshootingEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) DirectReports()(*id7592a93402c0d67508da968e58d07109c9c2561ae838fbeb1cccb27ef558ecb.DirectReportsRequestBuilder) {
    return id7592a93402c0d67508da968e58d07109c9c2561ae838fbeb1cccb27ef558ecb.NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Drive()(*i3a4cc4e91af7383718bf555760e1dba5c1fc1de3509837d9074d3a3f0a8a0fc5.DriveRequestBuilder) {
    return i3a4cc4e91af7383718bf555760e1dba5c1fc1de3509837d9074d3a3f0a8a0fc5.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Drives()(*i6fea53c850cb6db38ea2669df26d5c9986df82967fc9dcbd1415ef4c37ba3909.DrivesRequestBuilder) {
    return i6fea53c850cb6db38ea2669df26d5c9986df82967fc9dcbd1415ef4c37ba3909.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item collection
func (m *UserRequestBuilder) DrivesById(id string)(*i764f7513465c0251a3785c28eac14dab3642e586e0fad67528f13704ed1cbb64.DriveRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive_id"] = id
    }
    return i764f7513465c0251a3785c28eac14dab3642e586e0fad67528f13704ed1cbb64.NewDriveRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Events()(*ifd9b701a6f88aa459780777b5ac521fe3494f01b9276c6884b86827296582e54.EventsRequestBuilder) {
    return ifd9b701a6f88aa459780777b5ac521fe3494f01b9276c6884b86827296582e54.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item collection
func (m *UserRequestBuilder) EventsById(id string)(*ifdc39e295bcc2e19fecc18f2ae737ddb8fe9018d8753e81ef0c2de8bb37638da.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return ifdc39e295bcc2e19fecc18f2ae737ddb8fe9018d8753e81ef0c2de8bb37638da.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) ExportPersonalData()(*iedc2060488fc2f6e7d120554e9e853117ec2aab876ef464c9ed3da4696541a32.ExportPersonalDataRequestBuilder) {
    return iedc2060488fc2f6e7d120554e9e853117ec2aab876ef464c9ed3da4696541a32.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Extensions()(*if0c7df43a05145fae76c1ab18483108fb4043f7b6a0b01f47772fd66022550ee.ExtensionsRequestBuilder) {
    return if0c7df43a05145fae76c1ab18483108fb4043f7b6a0b01f47772fd66022550ee.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.extensions.item collection
func (m *UserRequestBuilder) ExtensionsById(id string)(*i129160b87cb78d35acdd1a84c7328c6872f51bf17cb02a211ba957a5431db64b.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i129160b87cb78d35acdd1a84c7328c6872f51bf17cb02a211ba957a5431db64b.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) FindMeetingTimes()(*i8fdd3e26eeb6da18c0b9d487ac43983355a059891790fba9c732ef59d5c4f21b.FindMeetingTimesRequestBuilder) {
    return i8fdd3e26eeb6da18c0b9d487ac43983355a059891790fba9c732ef59d5c4f21b.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) FollowedSites()(*if19172f362f2e1d4456e041244db6eae76a05acf4f16807f05721a5ef0a0f4be.FollowedSitesRequestBuilder) {
    return if19172f362f2e1d4456e041244db6eae76a05acf4f16807f05721a5ef0a0f4be.NewFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from users by key
func (m *UserRequestBuilder) Get(options *UserRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewUser() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User), nil
}
func (m *UserRequestBuilder) GetMailTips()(*i6d5ec618068b4f9940966e3235a6b56547b161354e35247fc9843cd2fc68c280.GetMailTipsRequestBuilder) {
    return i6d5ec618068b4f9940966e3235a6b56547b161354e35247fc9843cd2fc68c280.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses builds and executes requests for operations under \users\{user-id}\microsoft.graph.getManagedAppDiagnosticStatuses()
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ieb873ea17846be13ef65aaf00090f933a16727baa2e8063b39c2d86ae07e1ad0.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ieb873ea17846be13ef65aaf00090f933a16727baa2e8063b39c2d86ae07e1ad0.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies builds and executes requests for operations under \users\{user-id}\microsoft.graph.getManagedAppPolicies()
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i23845c92400f42d2b469653fba5fe1555d60b50e4596b30d2d2b4c127b0ceecf.GetManagedAppPoliciesRequestBuilder) {
    return i23845c92400f42d2b469653fba5fe1555d60b50e4596b30d2d2b4c127b0ceecf.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) GetMemberGroups()(*i54c28e8ec50e0dd4eeee8073ed79cecee11ba12e3813ec711d8127af3d4b994c.GetMemberGroupsRequestBuilder) {
    return i54c28e8ec50e0dd4eeee8073ed79cecee11ba12e3813ec711d8127af3d4b994c.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) GetMemberObjects()(*ide0a33943859ca09b2034a1804f60d32a108dc09bf7012d68e5d3c63ac617ddb.GetMemberObjectsRequestBuilder) {
    return ide0a33943859ca09b2034a1804f60d32a108dc09bf7012d68e5d3c63ac617ddb.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) InferenceClassification()(*i64f47462143d9c5d6613e409adca3c0fd24631de22e340968821ffdff467e657.InferenceClassificationRequestBuilder) {
    return i64f47462143d9c5d6613e409adca3c0fd24631de22e340968821ffdff467e657.NewInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Insights()(*i8deae6de37c6b836f34cd5c6b8fab91a73043c26ee346f7410a8dcf88ea64c99.InsightsRequestBuilder) {
    return i8deae6de37c6b836f34cd5c6b8fab91a73043c26ee346f7410a8dcf88ea64c99.NewInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) JoinedTeams()(*i7f31fc68d418ebfd772842294b63ed5b4dafc3cab499c7bf25b7634a901656bb.JoinedTeamsRequestBuilder) {
    return i7f31fc68d418ebfd772842294b63ed5b4dafc3cab499c7bf25b7634a901656bb.NewJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeamsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item collection
func (m *UserRequestBuilder) JoinedTeamsById(id string)(*i0b4dff6b82c5c98b58b9ff11aead66a7e49d0eac6b18b1490fd01bf9597cf36c.TeamRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team_id"] = id
    }
    return i0b4dff6b82c5c98b58b9ff11aead66a7e49d0eac6b18b1490fd01bf9597cf36c.NewTeamRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) LicenseDetails()(*i065d0a279baf0cb315c26855c2ba7a5c0701939d8cae543a4fb9b7bb2c7cbda5.LicenseDetailsRequestBuilder) {
    return i065d0a279baf0cb315c26855c2ba7a5c0701939d8cae543a4fb9b7bb2c7cbda5.NewLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.licenseDetails.item collection
func (m *UserRequestBuilder) LicenseDetailsById(id string)(*i065d0a279baf0cb315c26855c2ba7a5c0701939d8cae543a4fb9b7bb2c7cbda5.LicenseDetailsRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails_id"] = id
    }
    return i065d0a279baf0cb315c26855c2ba7a5c0701939d8cae543a4fb9b7bb2c7cbda5.NewLicenseDetailsRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) MailFolders()(*ib946e3869f342130dddab8c76c89f264a76e899f07657349241151b4e41b4286.MailFoldersRequestBuilder) {
    return ib946e3869f342130dddab8c76c89f264a76e899f07657349241151b4e41b4286.NewMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MailFoldersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item collection
func (m *UserRequestBuilder) MailFoldersById(id string)(*id1f5db1c2220a25bede6871151a00a067652bdb75e2db4d199ca6ffcdd70b342.MailFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder_id"] = id
    }
    return id1f5db1c2220a25bede6871151a00a067652bdb75e2db4d199ca6ffcdd70b342.NewMailFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) ManagedAppRegistrations()(*if12f3532feee7012e3a4dc1e3a2444ca8fab8f29e8128baafe0ef00f6ed6015c.ManagedAppRegistrationsRequestBuilder) {
    return if12f3532feee7012e3a4dc1e3a2444ca8fab8f29e8128baafe0ef00f6ed6015c.NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) ManagedDevices()(*i91af54e73e80a81d2f98c6f72fbb6c81539e67ba0480227ae8e2e5f2b8f467fa.ManagedDevicesRequestBuilder) {
    return i91af54e73e80a81d2f98c6f72fbb6c81539e67ba0480227ae8e2e5f2b8f467fa.NewManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.managedDevices.item collection
func (m *UserRequestBuilder) ManagedDevicesById(id string)(*i0c6b3d37dc64615f253a6fd9a58f59f536aa93a57f99bb5f9d4210ae83aa5112.ManagedDeviceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice_id"] = id
    }
    return i0c6b3d37dc64615f253a6fd9a58f59f536aa93a57f99bb5f9d4210ae83aa5112.NewManagedDeviceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Manager()(*i3fd2d5130e27e5ed481f7a0565ee17271d52e0ea15f0dec8ae436a3780d1d8d2.ManagerRequestBuilder) {
    return i3fd2d5130e27e5ed481f7a0565ee17271d52e0ea15f0dec8ae436a3780d1d8d2.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) MemberOf()(*ia6a0f4f70d60f824df18f07f8918c6e9ac043f588293118e5d91eb921d5dffbf.MemberOfRequestBuilder) {
    return ia6a0f4f70d60f824df18f07f8918c6e9ac043f588293118e5d91eb921d5dffbf.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Messages()(*i0a6d7999baeed47d1aa02960cd9f3f120528ecae4240bf52300542a8dc9f978d.MessagesRequestBuilder) {
    return i0a6d7999baeed47d1aa02960cd9f3f120528ecae4240bf52300542a8dc9f978d.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.messages.item collection
func (m *UserRequestBuilder) MessagesById(id string)(*if33563b345c020375260161400a6b809f00973f17d5ab870445f3c67ca7323b8.MessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return if33563b345c020375260161400a6b809f00973f17d5ab870445f3c67ca7323b8.NewMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Oauth2PermissionGrants()(*iea339692a4ece4755a03d26004b71e7bbfe3ad28cf8036bdc4b972512eb00359.Oauth2PermissionGrantsRequestBuilder) {
    return iea339692a4ece4755a03d26004b71e7bbfe3ad28cf8036bdc4b972512eb00359.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Onenote()(*ib8e8ec0d09197f1fd57167b1b67da9bcc229fa24977921ce6406f24aa89a6884.OnenoteRequestBuilder) {
    return ib8e8ec0d09197f1fd57167b1b67da9bcc229fa24977921ce6406f24aa89a6884.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) OnlineMeetings()(*i2c0f098b49b5f2edba81f4095f0c699516263b7cd6ff316d31f469f15b8366b0.OnlineMeetingsRequestBuilder) {
    return i2c0f098b49b5f2edba81f4095f0c699516263b7cd6ff316d31f469f15b8366b0.NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onlineMeetings.item collection
func (m *UserRequestBuilder) OnlineMeetingsById(id string)(*i3ecd250b638acba95f41b3db07eaa507842a51f8e0cd3640cb9b75f0052ceff6.OnlineMeetingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting_id"] = id
    }
    return i3ecd250b638acba95f41b3db07eaa507842a51f8e0cd3640cb9b75f0052ceff6.NewOnlineMeetingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Outlook()(*i418d77dc601ca4dc2c568aef8702ce31b38463f4e1701c6c788adcbeb32ca3e8.OutlookRequestBuilder) {
    return i418d77dc601ca4dc2c568aef8702ce31b38463f4e1701c6c788adcbeb32ca3e8.NewOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) OwnedDevices()(*i005f0821c1e57467ed9e31e0f65ebf8795241942e8afc0c6c9cc866dc3b70de2.OwnedDevicesRequestBuilder) {
    return i005f0821c1e57467ed9e31e0f65ebf8795241942e8afc0c6c9cc866dc3b70de2.NewOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) OwnedObjects()(*i152c9659f52d4d76fff6d586e3de05874eee179419588197c21a3adc1e3a6c7b.OwnedObjectsRequestBuilder) {
    return i152c9659f52d4d76fff6d586e3de05874eee179419588197c21a3adc1e3a6c7b.NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in users
func (m *UserRequestBuilder) Patch(options *UserRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *UserRequestBuilder) People()(*i2b4d48c4393811d79f4d4cc1f4addb3b49f58a53edcf7d1d7140e29a22ab7504.PeopleRequestBuilder) {
    return i2b4d48c4393811d79f4d4cc1f4addb3b49f58a53edcf7d1d7140e29a22ab7504.NewPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PeopleById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.people.item collection
func (m *UserRequestBuilder) PeopleById(id string)(*i8465a2d5f26666d08a0baebc755ef48f1f0dc18551f929c58c09d0110510a317.PersonRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person_id"] = id
    }
    return i8465a2d5f26666d08a0baebc755ef48f1f0dc18551f929c58c09d0110510a317.NewPersonRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Photo()(*icf56cc7fcccc9803d5655a975f93bee1c76d14d27e2e180c40ef7bb1c3c18a68.PhotoRequestBuilder) {
    return icf56cc7fcccc9803d5655a975f93bee1c76d14d27e2e180c40ef7bb1c3c18a68.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Photos()(*idea2a9c81deaaa3e33a2bcef2386387935820f4f571d71af1fc2daa24d8fa032.PhotosRequestBuilder) {
    return idea2a9c81deaaa3e33a2bcef2386387935820f4f571d71af1fc2daa24d8fa032.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.photos.item collection
func (m *UserRequestBuilder) PhotosById(id string)(*i5f2b6c8998555350079828979727c205406cf7155e907008bec42440c3e9724a.ProfilePhotoRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto_id"] = id
    }
    return i5f2b6c8998555350079828979727c205406cf7155e907008bec42440c3e9724a.NewProfilePhotoRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Planner()(*i4a474d43148dd96e52e36b3ed7c52a7d75f46d95adf4fa0ce96ace2509d33744.PlannerRequestBuilder) {
    return i4a474d43148dd96e52e36b3ed7c52a7d75f46d95adf4fa0ce96ace2509d33744.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Presence()(*i0a4c7a84985f2276fcfbcf69e2fe37b80ef87c2622cc8763371022616974902d.PresenceRequestBuilder) {
    return i0a4c7a84985f2276fcfbcf69e2fe37b80ef87c2622cc8763371022616974902d.NewPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) RegisteredDevices()(*i631d94a7300d4142a2d7a4c4d0d424e98d3fb058f106e6567a8ac6476f5956b6.RegisteredDevicesRequestBuilder) {
    return i631d94a7300d4142a2d7a4c4d0d424e98d3fb058f106e6567a8ac6476f5956b6.NewRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime builds and executes requests for operations under \users\{user-id}\microsoft.graph.reminderView(StartDateTime='{StartDateTime}',EndDateTime='{EndDateTime}')
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(startDateTime *string, endDateTime *string)(*iddc156b6ba287ee33d5592797e8da64087fd068cf179bd892944253be67ac174.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return iddc156b6ba287ee33d5592797e8da64087fd068cf179bd892944253be67ac174.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime);
}
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i6d278e5cb359938f2362c090cc322d0c80867015903a39849ea478a6cda51689.RemoveAllDevicesFromManagementRequestBuilder) {
    return i6d278e5cb359938f2362c090cc322d0c80867015903a39849ea478a6cda51689.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ifd940c65ea13f1270f41f753ee5e7761f074915bd6a6c64b2adb77f70aea8c32.ReprocessLicenseAssignmentRequestBuilder) {
    return ifd940c65ea13f1270f41f753ee5e7761f074915bd6a6c64b2adb77f70aea8c32.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Restore()(*i9f87171771d54744395d1043beb42c14430d7cb4f1557826ebef7c5b2584f7cd.RestoreRequestBuilder) {
    return i9f87171771d54744395d1043beb42c14430d7cb4f1557826ebef7c5b2584f7cd.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) RevokeSignInSessions()(*i1090e84ad9b76c4c279467ebbc90ff155a540b45b098e680706d2aae300639eb.RevokeSignInSessionsRequestBuilder) {
    return i1090e84ad9b76c4c279467ebbc90ff155a540b45b098e680706d2aae300639eb.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) ScopedRoleMemberOf()(*ia99f649e03b3020696f5c2c11f42a9ad7f094cb80add889863fc7947d3bf9ec7.ScopedRoleMemberOfRequestBuilder) {
    return ia99f649e03b3020696f5c2c11f42a9ad7f094cb80add889863fc7947d3bf9ec7.NewScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.scopedRoleMemberOf.item collection
func (m *UserRequestBuilder) ScopedRoleMemberOfById(id string)(*i983bdc3dce2d947aec82e1e9dd5edd887e8d861d589287d050f96af67abb1e1f.ScopedRoleMembershipRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return i983bdc3dce2d947aec82e1e9dd5edd887e8d861d589287d050f96af67abb1e1f.NewScopedRoleMembershipRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) SendMail()(*i16f13fb20dc21e9cbd93cbca432ac5cb871c47a97e5bebce67f5eab6fe8276e9.SendMailRequestBuilder) {
    return i16f13fb20dc21e9cbd93cbca432ac5cb871c47a97e5bebce67f5eab6fe8276e9.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Settings()(*i225d659ae554743be8acfa963ab6384d1f6e82911866fa2e3c01918b6a0637b1.SettingsRequestBuilder) {
    return i225d659ae554743be8acfa963ab6384d1f6e82911866fa2e3c01918b6a0637b1.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Teamwork()(*i94ff5ee566023b18b8b2ac8c1da92081f3bf4520205017599c00aed9faff56de.TeamworkRequestBuilder) {
    return i94ff5ee566023b18b8b2ac8c1da92081f3bf4520205017599c00aed9faff56de.NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Todo()(*i61d84fd6521641b94f790414a5b9abccf870ea66012c270c06f3ea7df770aad8.TodoRequestBuilder) {
    return i61d84fd6521641b94f790414a5b9abccf870ea66012c270c06f3ea7df770aad8.NewTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) TransitiveMemberOf()(*i11db6285564609d0b36b9032e0d8f6d257828b887e558375da4354121f40a3d0.TransitiveMemberOfRequestBuilder) {
    return i11db6285564609d0b36b9032e0d8f6d257828b887e558375da4354121f40a3d0.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) TranslateExchangeIds()(*ifdf31129bec8038d82607a6d973830e0b090629373d5812152e5046514a4bfb0.TranslateExchangeIdsRequestBuilder) {
    return ifdf31129bec8038d82607a6d973830e0b090629373d5812152e5046514a4bfb0.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i4af8697d288ced3f62e5f7467084c929a355a7f4c2140a3f51072b00f73663c1.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i4af8697d288ced3f62e5f7467084c929a355a7f4c2140a3f51072b00f73663c1.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
