package main

// ClockEntry is the basic information needed to run a channel clock.
type ClockEntry struct {
	ChannelID string `json:"channelID"`
	Timezone  string `json:"timezone"`
}

// Resources are shared throughout the application.
type Resources struct {
	Clocks   []ClockEntry `json:"clocks"`
	ModRoles []string     `json:"mod_roles"`
	Prefix   string       `json:"prefix"`
}
