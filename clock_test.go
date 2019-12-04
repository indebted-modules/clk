package clk_test

import (
	"testing"
	"time"

	"github.com/indebted-modules/clk"
	"github.com/stretchr/testify/suite"
)

type ClockSuite struct {
	suite.Suite
}

func TestClockSuite(t *testing.T) {
	suite.Run(t, new(ClockSuite))
}

func (s *ClockSuite) TestSystemClock() {
	clock := &clk.SystemClock{}
	s.NotEmpty(clock.Now())
}

func (s *ClockSuite) TestFrozenClock() {
	clock := &clk.FrozenClock{}
	s.Equal(clk.FrozenTime, clock.Now())
	s.NotEmpty(clock.Now())
}

func (s *ClockSuite) TestFrozenClockCustomTime() {
	clock := &clk.FrozenClock{ClockTime: time.Date(2015, 10, 21, 21, 00, 0, 0, time.UTC)}
	s.Equal(time.Date(2015, 10, 21, 21, 00, 0, 0, time.UTC), clock.Now())
	s.NotEmpty(clock.Now())
}

func (s *ClockSuite) TestDaysSince() {
	oneWeekAgo := time.Now().Add(time.Hour * 24 * -7)
	s.Equal(int64(7), clk.DaysSince(oneWeekAgo))
}

func (s *ClockSuite) TestDaysSinceFloorHours() {
	_23hAgo := time.Now().Add(time.Hour * -23)
	s.Equal(int64(0), clk.DaysSince(_23hAgo))
}

func (s *ClockSuite) TestParseDate() {
	t, err := clk.ParseDate("2019-06-30")
	s.Equal(time.Date(2019, 6, 30, 0, 0, 0, 0, time.UTC), t)
	s.Nil(err)
}

func (s *ClockSuite) TestParseDateInvalidFormat() {
	t, err := clk.ParseDate("2019-30-30")
	s.Zero(t)
	s.Error(err)
}
