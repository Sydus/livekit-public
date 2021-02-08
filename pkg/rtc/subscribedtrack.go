package rtc

import (
	"github.com/pion/ion-sfu/pkg/sfu"

	"github.com/livekit/livekit-server/pkg/utils"
)

type SubscribedTrack struct {
	dt       *sfu.DownTrack
	subMuted utils.AtomicFlag
	pubMuted utils.AtomicFlag
}

func NewSubscribedTrack(dt *sfu.DownTrack) *SubscribedTrack {
	return &SubscribedTrack{
		dt: dt,
	}
}

func (t *SubscribedTrack) DownTrack() *sfu.DownTrack {
	return t.dt
}

// has subscriber indicated it wants to mute this track
func (t *SubscribedTrack) IsMuted() bool {
	return t.subMuted.Get()
}

// set subscriber mute preference
func (t *SubscribedTrack) SetMuted(muted bool) {
	t.subMuted.TrySet(muted)
	t.updateDownTrackMute()
}

func (t *SubscribedTrack) SetPublisherMuted(muted bool) {
	t.pubMuted.TrySet(muted)
	t.updateDownTrackMute()
}

func (t *SubscribedTrack) updateDownTrackMute() {
	muted := t.subMuted.Get() || t.pubMuted.Get()
	t.dt.Mute(muted)
}
