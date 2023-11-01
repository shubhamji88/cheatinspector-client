package watchman

import (
	"fmt"
	"strconv"
	"time"

	"github.com/YashKumarVerma/cheatinspector-client/internal/ably"
	"github.com/YashKumarVerma/cheatinspector-client/internal/config"
	"github.com/YashKumarVerma/cheatinspector-client/internal/sensor"
)

func NotifyBackend(entropy uint64) bool {
	ably.UserWritingCodePresence(sensor.Load.MachineID)
	channel := ably.ClientChannel
	channel.PublishAsync("entropy", string(strconv.FormatUint(entropy, 10)), func(err error) {
		if err != nil {
			fmt.Println("failed to publish", err)
		} else {
			fmt.Println("publish ok" + sensor.Load.MachineID)
		}
	})

	if config.Load.Debug == true {
		fmt.Println("Publisher Entropy Call : Params : ", strconv.FormatUint(entropy, 10))
		fmt.Println("Publisher Entropy Call : Stamp : ", time.Unix(time.Now().Unix(), 0).Format(time.RFC1123Z))
	} else {
		fmt.Println("Entropy Transmitted.")
	}

	return true
}

func UpdateSnapshot(data uint64) bool {

	channel := ably.ClientChannel
	channel.PublishAsync("snapshot", string(strconv.FormatUint(data, 10)), func(err error) {
		if err != nil {
			fmt.Println("failed to publish", err)
		} else {
			fmt.Println("publish ok" + sensor.Load.MachineID)
		}
	})

	if config.Load.Debug == true {
		fmt.Println("Publisher Snapshot Call : Params : ", strconv.FormatUint(data, 10))
		fmt.Println("Publisher Snapshot Call : Stamp : ", time.Unix(time.Now().Unix(), 0).Format(time.RFC1123Z))
	} else {
		fmt.Println("Snapshot transmitted.")
	}

	return true
}
