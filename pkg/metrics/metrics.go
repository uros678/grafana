package metrics

var MetricStats Registry
var UseNilMetrics bool

func init() {
	// init with nil metrics
	initMetricVars(&MetricSettings{})
}

var (
	M_Instance_Start                     Counter
	M_Page_Status_200                    Counter
	M_Page_Status_500                    Counter
	M_Page_Status_404                    Counter
	M_Api_Status_500                     Counter
	M_Api_Status_404                     Counter
	M_Api_User_SignUpStarted             Counter
	M_Api_User_SignUpCompleted           Counter
	M_Api_User_SignUpInvite              Counter
	M_Api_Dashboard_Save                 Timer
	M_Api_Dashboard_Get                  Timer
	M_Api_Dashboard_Search               Timer
	M_Api_Admin_User_Create              Counter
	M_Api_Login_Post                     Counter
	M_Api_Login_OAuth                    Counter
	M_Api_Org_Create                     Counter
	M_Api_Dashboard_Snapshot_Create      Counter
	M_Api_Dashboard_Snapshot_External    Counter
	M_Api_Dashboard_Snapshot_Get         Counter
	M_Models_Dashboard_Insert            Counter
	M_Alerting_Result_Critical           Counter
	M_Alerting_Result_Warning            Counter
	M_Alerting_Result_Info               Counter
	M_Alerting_Result_Ok                 Counter
	M_Alerting_Active_Alerts             Counter
	M_Alerting_Notification_Sent_Slack   Counter
	M_Alerting_Notification_Sent_Email   Counter
	M_Alerting_Notification_Sent_Webhook Counter

	// Timers
	M_DataSource_ProxyReq_Timer Timer
	M_Alerting_Exeuction_Time   Timer
)

func initMetricVars(settings *MetricSettings) {
	UseNilMetrics = settings.Enabled == false
	MetricStats = NewRegistry()

	M_Instance_Start = RegCounter("instance_start")

	M_Page_Status_200 = RegCounter("page.resp_status", "code", "200")
	M_Page_Status_500 = RegCounter("page.resp_status", "code", "500")
	M_Page_Status_404 = RegCounter("page.resp_status", "code", "404")

	M_Api_Status_500 = RegCounter("api.resp_status", "code", "500")
	M_Api_Status_404 = RegCounter("api.resp_status", "code", "404")

	M_Api_User_SignUpStarted = RegCounter("api.user.signup_started")
	M_Api_User_SignUpCompleted = RegCounter("api.user.signup_completed")
	M_Api_User_SignUpInvite = RegCounter("api.user.signup_invite")

	M_Api_Dashboard_Save = RegTimer("api.dashboard.save")
	M_Api_Dashboard_Get = RegTimer("api.dashboard.get")
	M_Api_Dashboard_Search = RegTimer("api.dashboard.search")

	M_Api_Admin_User_Create = RegCounter("api.admin.user_create")
	M_Api_Login_Post = RegCounter("api.login.post")
	M_Api_Login_OAuth = RegCounter("api.login.oauth")
	M_Api_Org_Create = RegCounter("api.org.create")

	M_Api_Dashboard_Snapshot_Create = RegCounter("api.dashboard_snapshot.create")
	M_Api_Dashboard_Snapshot_External = RegCounter("api.dashboard_snapshot.external")
	M_Api_Dashboard_Snapshot_Get = RegCounter("api.dashboard_snapshot.get")

	M_Models_Dashboard_Insert = RegCounter("models.dashboard.insert")

	M_Alerting_Result_Critical = RegCounter("alerting.result", "severity", "critical")
	M_Alerting_Result_Warning = RegCounter("alerting.result", "severity", "warning")
	M_Alerting_Result_Info = RegCounter("alerting.result", "severity", "info")
	M_Alerting_Result_Ok = RegCounter("alerting.result", "severity", "ok")
	M_Alerting_Active_Alerts = RegCounter("alerting.active_alerts")
	M_Alerting_Notification_Sent_Slack = RegCounter("alerting.notifications_sent", "type", "slack")
	M_Alerting_Notification_Sent_Email = RegCounter("alerting.notifications_sent", "type", "email")
	M_Alerting_Notification_Sent_Webhook = RegCounter("alerting.notifications_sent", "type", "webhook")

	// Timers
	M_DataSource_ProxyReq_Timer = RegTimer("api.dataproxy.request.all")
	M_Alerting_Exeuction_Time = RegTimer("alerting.execution_time")
}
