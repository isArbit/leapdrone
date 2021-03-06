/*
	Package leap implements all leap motion needed functions
*/
package leap

import (
	lp "github.com/hybridgroup/gobot/platforms/leap"
)

// NewHand initiate a hand struct
func NewHand(h lp.Hand, th float64) *Hand {
	return &Hand{
		Hand:      h,
		ID:        h.ID,
		Threshold: th,
	}
}

// Hand is a extension of a leap motion hand by adding move direction
// detection.
type Hand struct {
	// embeded a leap motion hand
	lp.Hand

	ID        int
	Threshold float64
}

// IsUpward determines if hand is moving forward.
func (h *Hand) IsForward() bool {
	// negative Z-axis movement
	if h.PalmVelocity[2] < -(h.Threshold) {
		return true
	}
	return false
}

// IsUpward determines if hand is moving back.
func (h *Hand) IsBackward() bool {
	// positive Z-axis movement
	if h.PalmVelocity[2] > (h.Threshold) {
		return true
	}
	return false
}

// IsUpward determines if hand is moving up.
func (h *Hand) IsUpward() bool {
	// positive Y-axis movement
	if h.PalmVelocity[1] > (h.Threshold) {
		return true
	}
	return false
}

// IsDownward determines if hand is moving down.
func (h *Hand) IsDownward() bool {
	// negative Y-axis movement
	if h.PalmVelocity[1] < -(h.Threshold) {
		return true
	}
	return false
}

// IsRight determines if hand is moving right.
func (h *Hand) IsRight() bool {
	// positive X-axis movement
	if h.PalmVelocity[0] > (h.Threshold) {
		return true
	}
	return false
}

// IsLeft determines if hand is moving left.
func (h *Hand) IsLeft() bool {
	// negative X-axis movement
	if h.PalmVelocity[0] < -(h.Threshold) {
		return true
	}
	return false
}
