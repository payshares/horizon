package horizon

import (
	"github.com/payshares/go/psrkey"
	"github.com/payshares/horizon/friendbot"
)

func initFriendbot(app *App) {
	if app.config.FriendbotSecret == "" {
		return
	}

	// ensure its a seed if its not blank
	psrkey.MustDecode(psrkey.VersionByteSeed, app.config.FriendbotSecret)

	app.friendbot = &friendbot.Bot{
		Secret:    app.config.FriendbotSecret,
		Submitter: app.submitter,
		Network:   app.networkPassphrase,
	}

}

func init() {
	appInit.Add("friendbot", initFriendbot, "txsub", "paysharesCoreInfo")
}
