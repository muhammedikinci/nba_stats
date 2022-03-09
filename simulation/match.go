package simulation

import (
	"math/rand"
)

type match struct {
	Home        team       `json:"home"`
	Away        team       `json:"away"`
	AttackCount int        `json:"attackCount"`
	AttackTimer int        `json:"attackTimer"`
	CurrentSide int        `json:"-"`
	Random      *rand.Rand `json:"-"`
}

func (m match) getPoint() int {
	threePointProbability := m.Random.NormFloat64()
	generalProbability := m.Random.NormFloat64()

	if generalProbability > 0 && threePointProbability > 0 {
		return 3
	} else if generalProbability < 0 {
		return 0
	}

	return 2
}

func (m *match) reverseSides() {
	if m.CurrentSide == 0 {
		m.CurrentSide = 1
		return
	}

	m.CurrentSide = 0
	m.AttackCount++
	m.AttackTimer = 0
}

func (m match) check(awayPlayerProbability, homePlayerProbability float64, factors ...float64) int {
	generalProbability := m.Random.NormFloat64()

	if len(factors) == 2 {
		mean := factors[0]
		deviation := factors[1]
		generalProbability = generalProbability*mean + deviation
	}

	if generalProbability < 0 {
		return -1
	}

	if m.CurrentSide == 0 && awayPlayerProbability > homePlayerProbability {
		return 1
	}

	if m.CurrentSide == 1 && homePlayerProbability > awayPlayerProbability {
		return 0
	}

	return -1
}
