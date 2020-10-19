package application

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DiscordClocks/models"
)

var clockFaces = []string{"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"}
var halfHourClockFaces = []string{"🕜", "🕝", "🕞", "🕟", "🕠", "🕡", "🕢", "🕣", "🕤", "🕥", "🕦", "🕧"}

// RunClocks runs the clocks.
func RunClocks(client *discordgo.Session, resources *models.Resources) {
	startNow := true

	for {
		for _, clock := range resources.Clocks {
			now := time.Now()

			minute := now.Minute()
			hour := now.Hour()
			var clockEmoji string
			if minute < 30 {
				clockEmoji = clockFaces[hour-1]
			} else {
				clockEmoji = halfHourClockFaces[hour-1]
			}

			timeString := clockEmoji + " " + now.Format("h:mm A")
			if clock.TzCode != "" {
				timeString += " " + clock.TzCode
			}

			channel, err := client.Channel(clock.ChannelID)
			if err != nil {
				log.Println(err)
				continue
			}

			if minute%5 == 0 && channel.Name != timeString {
				if _, err := client.ChannelEdit(channel.ID, timeString); err != nil {
					log.Println(err)
				}
			}
		}

		if startNow {
			startNow = false
		} else {
			time.Sleep(10000)
		}
	}
}
