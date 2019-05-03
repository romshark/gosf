package gosf

// Plugin is the framework interface defining a plugin
type Plugin interface {
	Activate(app *AppSettings)
	Deactivate(app *AppSettings)
}

var plugins []*Plugin

func init() {
	App.Plugins = make(map[string]interface{})
}

// RegisterPlugin is used by the plugin to register itself for later activation
func RegisterPlugin(plugin *Plugin) {
	registration := new(Plugin)
	registration = plugin

	plugins = append(plugins, registration)
}
