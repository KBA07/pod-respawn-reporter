package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodNotifRestartSpec defines the desired state of PodNotifRestart
type PodNotifRestartSpec struct {
	// NamespacesToMonitor specifies which namespaces to monitor.
	// If empty, all namespaces will be monitored.
	NamespacesToMonitor []string `json:"namespacesToMonitor,omitempty"`

	// MinRestarts specifies the minimum number of restarts before sending a notification.
	MinRestarts int32 `json:"minRestarts,omitempty"`

	// TeamsWebhookURL is the webhook URL for Microsoft Teams notifications.
	TeamsWebhookURL string `json:"teamsWebhookURL,omitempty"`

	// SlackWebhookURL is the webhook URL for Slack notifications.
	SlackWebhookURL string `json:"slackWebhookURL,omitempty"`

	WebhookURL string `json:"webhookURL"`
}

// PodNotifRestartStatus defines the observed state of PodNotifRestart
type PodNotifRestartStatus struct {
	// LastChecked is the timestamp when the pods were last checked.
	LastChecked metav1.Time `json:"lastChecked"`

	// NotificationsSent is the number of notifications sent so far.
	NotificationsSent int32 `json:"notificationsSent"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodNotifRestart is the Schema for the podnotifrestarts API
type PodNotifRestart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodNotifRestartSpec   `json:"spec,omitempty"`
	Status PodNotifRestartStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodNotifRestartList contains a list of PodNotifRestart
type PodNotifRestartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodNotifRestart `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodNotifRestart{}, &PodNotifRestartList{})
}

// GetNamespacesToMonitor returns the namespaces to monitor.
func (p *PodNotifRestart) GetNamespacesToMonitor() []string {
	if p.Spec.NamespacesToMonitor == nil {
		return []string{}
	}
	return p.Spec.NamespacesToMonitor
}

// GetMinRestarts returns the minimum number of restarts before sending a notification.
func (p *PodNotifRestart) GetMinRestarts() int32 {
	if p.Spec.MinRestarts == 0 {
		return 1
	}
	return p.Spec.MinRestarts
}

// GetTeamsWebhookURL returns the Microsoft Teams webhook URL.
func (p *PodNotifRestart) GetTeamsWebhookURL() string {
	return p.Spec.TeamsWebhookURL
}

// GetSlackWebhookURL returns the Slack webhook URL.
func (p *PodNotifRestart) GetSlackWebhookURL() string {
	return p.Spec.SlackWebhookURL
}

// GetLastChecked returns the timestamp when the pods were last checked.
func (p *PodNotifRestart) GetLastChecked() metav1.Time {
	return p.Status.LastChecked
}

// GetNotificationsSent returns the number of notifications sent so far.
func (p *PodNotifRestart) GetNotificationsSent() int32 {
	return p.Status.NotificationsSent
}

// SetLastChecked sets the timestamp when the pods were last checked.
func (p *PodNotifRestart) SetLastChecked(t metav1.Time) {
	p.Status.LastChecked = t
}

// SetNotificationsSent sets the number of notifications sent so far.
func (p *PodNotifRestart) SetNotificationsSent(n int32) {
	p.Status.NotificationsSent = n
}

// Validate webhook URL for example if it is try to put null
func (p *PodNotifRestart) Validate() bool {
	if p.Spec.TeamsWebhookURL == "" && p.Spec.SlackWebhookURL == "" {
		return false
	}
	return true
}

// GetWebhookURL returns the webhook URL.
func (p *PodNotifRestart) GetWebhookURL() string {
	if p.Spec.TeamsWebhookURL != "" {
		return p.Spec.TeamsWebhookURL
	}
	if p.Spec.SlackWebhookURL != "" {
		return p.Spec.SlackWebhookURL
	}
	return ""
}

// AddNotificationSent adds 1 to the number of notifications sent so far.
func (p *PodNotifRestart) AddNotificationSent() {
	p.Status.NotificationsSent++
}
