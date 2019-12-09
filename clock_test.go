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

func (s *ClockSuite) TestSystemClockIsInUTC() {
	clock := &clk.SystemClock{}
	now := clock.Now()
	name, _ := now.Zone()
	s.Equal("UTC", name)
}

func (s *ClockSuite) TestSystemClockIsInTruncatedAtMilliseconds() {
	clock := &clk.SystemClock{}
	now := clock.Now()
	s.Equal(int64(0), now.UnixNano()%100000)
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

func (s *ClockSuite) TestUnix() {
	t := clk.Unix(1575876119)
	name, _ := t.Zone()
	s.Equal("UTC", name)
}

func (s *ClockSuite) TestDate() {
	t := clk.Date(2019, time.February, 25, 10, 32, 40)
	s.Equal(2019, t.Year())
	s.Equal(time.February, t.Month())
	s.Equal(25, t.Day())
	s.Equal(10, t.Hour())
	s.Equal(32, t.Minute())
	s.Equal(40, t.Second())
	s.Equal(0, t.Nanosecond())
	name, _ := t.Zone()
	s.Equal("UTC", name)
}
