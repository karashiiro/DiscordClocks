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
			loc, err := time.LoadLocation(clock.Timezone)
			if err != nil {
				log.Println(err)
				continue
			}

			now := time.Now().In(loc)

			minute := now.Minute()
			hour := now.Hour() + 1
			if hour > 12 {
				hour -= 12
			}

			var clockEmoji string
			if minute < 30 {
				clockEmoji = clockFaces[hour-1]
			} else {
				clockEmoji = halfHourClockFaces[hour-1]
			}

			timeFormat := "3:04 PM"
			if clock.TzCode == "" {
				timeFormat += " MST"
			}

			timeString := clockEmoji + " " + now.Format(timeFormat)
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
