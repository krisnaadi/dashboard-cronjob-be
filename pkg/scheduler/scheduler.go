package scheduler

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Scheduler struct {
	ticker *time.Ticker
	jobs   []*job
	sync.RWMutex
}

type job struct {
	min       map[int]struct{}
	hour      map[int]struct{}
	day       map[int]struct{}
	month     map[int]struct{}
	dayOfWeek map[int]struct{}

	fn   interface{}
	args []interface{}
	sync.RWMutex
}

type tick struct {
	min       int
	hour      int
	day       int
	month     int
	dayOfWeek int
}

func New() *Scheduler {
	return new(time.Minute)
}

func new(t time.Duration) *Scheduler {
	c := &Scheduler{
		ticker: time.NewTicker(t),
		jobs:   []*job{},
	}

	go func() {
		for t := range c.ticker.C {
			c.runScheduled(t)
		}
	}()

	return c
}

func (c *Scheduler) AddJob(schedule string, fn interface{}, args ...interface{}) error {
	j, err := parseSchedule(schedule)
	c.Lock()
	defer c.Unlock()
	if err != nil {
		return err
	}

	if fn == nil || reflect.ValueOf(fn).Kind() != reflect.Func {
		return fmt.Errorf("cron job must be func()")
	}

	fnType := reflect.TypeOf(fn)
	if len(args) != fnType.NumIn() {
		return fmt.Errorf("number of func() params and number of provided params doesn't match")
	}

	for i := 0; i < fnType.NumIn(); i++ {
		a := args[i]
		t1 := fnType.In(i)
		t2 := reflect.TypeOf(a)

		if t1 != t2 {
			if t1.Kind() != reflect.Interface {
				return fmt.Errorf("param with index %d shold be `%s` not `%s`", i, t1, t2)
			}
			if !t2.Implements(t1) {
				return fmt.Errorf("param with index %d of type `%s` doesn't implement interface `%s`", i, t2, t1)
			}
		}
	}

	// all checked, add job to scheduler
	j.fn = fn
	j.args = args
	c.jobs = append(c.jobs, j)
	return nil
}

func (c *Scheduler) Shutdown() {
	c.ticker.Stop()
}

func (c *Scheduler) Clear() {
	c.Lock()
	c.jobs = []*job{}
	c.Unlock()
}

func (c *Scheduler) RunAll() {
	c.RLock()
	defer c.RUnlock()
	for _, j := range c.jobs {
		go j.run()
	}
}

func (c *Scheduler) runScheduled(t time.Time) {
	tick := getTick(t)
	c.RLock()
	defer c.RUnlock()

	for _, j := range c.jobs {
		if j.tick(tick) {
			go j.run()
		}
	}
}

func (j *job) run() {
	j.RLock()
	defer func() {
		if r := recover(); r != nil {
			log.Println("Scheduler error", r)
		}
	}()
	v := reflect.ValueOf(j.fn)
	rargs := make([]reflect.Value, len(j.args))
	for i, a := range j.args {
		rargs[i] = reflect.ValueOf(a)
	}
	j.RUnlock()
	v.Call(rargs)
}

func (j *job) tick(t tick) bool {
	j.RLock()
	defer j.RUnlock()
	if _, ok := j.min[t.min]; !ok {
		return false
	}

	if _, ok := j.hour[t.hour]; !ok {
		return false
	}

	_, day := j.day[t.day]
	_, dayOfWeek := j.dayOfWeek[t.dayOfWeek]
	if !day && !dayOfWeek {
		return false
	}

	if _, ok := j.month[t.month]; !ok {
		return false
	}

	return true
}

var (
	matchSpaces = regexp.MustCompile(`\s+`)
	matchN      = regexp.MustCompile(`(.*)/(\d+)`)
	matchRange  = regexp.MustCompile(`^(\d+)-(\d+)$`)
)

func parseSchedule(s string) (*job, error) {
	var err error
	j := &job{}
	j.Lock()
	defer j.Unlock()
	s = matchSpaces.ReplaceAllLiteralString(s, " ")
	parts := strings.Split(s, " ")
	if len(parts) != 5 {
		return j, errors.New("schedule string must have five components like * * * * *")
	}

	j.min, err = parsePart(parts[0], 0, 59)
	if err != nil {
		return j, err
	}

	j.hour, err = parsePart(parts[1], 0, 23)
	if err != nil {
		return j, err
	}

	j.day, err = parsePart(parts[2], 1, 31)
	if err != nil {
		return j, err
	}

	j.month, err = parsePart(parts[3], 1, 12)
	if err != nil {
		return j, err
	}

	j.dayOfWeek, err = parsePart(parts[4], 0, 6)
	if err != nil {
		return j, err
	}

	switch {
	case len(j.day) < 31 && len(j.dayOfWeek) == 7: // day set, but not dayOfWeek, clear dayOfWeek
		j.dayOfWeek = make(map[int]struct{})
	case len(j.dayOfWeek) < 7 && len(j.day) == 31: // dayOfWeek set, but not day, clear day
		j.day = make(map[int]struct{})
	default:
	}

	return j, nil
}

func parsePart(s string, min, max int) (map[int]struct{}, error) {

	r := make(map[int]struct{})

	// wildcard pattern
	if s == "*" {
		for i := min; i <= max; i++ {
			r[i] = struct{}{}
		}
		return r, nil
	}

	// */2 1-59/5 pattern
	if matches := matchN.FindStringSubmatch(s); matches != nil {
		localMin := min
		localMax := max
		if matches[1] != "" && matches[1] != "*" {
			if rng := matchRange.FindStringSubmatch(matches[1]); rng != nil {
				localMin, _ = strconv.Atoi(rng[1])
				localMax, _ = strconv.Atoi(rng[2])
				if localMin < min || localMax > max {
					return nil, fmt.Errorf("out of range for %s in %s. %s must be in range %d-%d", rng[1], s, rng[1], min, max)
				}
			} else {
				return nil, fmt.Errorf("unable to parse %s part in %s", matches[1], s)
			}
		}
		n, _ := strconv.Atoi(matches[2])
		for i := localMin; i <= localMax; i += n {
			r[i] = struct{}{}
		}
		return r, nil
	}

	// 1,2,4  or 1,2,10-15,20,30-45 pattern
	parts := strings.Split(s, ",")
	for _, x := range parts {
		if rng := matchRange.FindStringSubmatch(x); rng != nil {
			localMin, _ := strconv.Atoi(rng[1])
			localMax, _ := strconv.Atoi(rng[2])
			if localMin < min || localMax > max {
				return nil, fmt.Errorf("out of range for %s in %s. %s must be in range %d-%d", x, s, x, min, max)
			}
			for i := localMin; i <= localMax; i++ {
				r[i] = struct{}{}
			}
		} else if i, err := strconv.Atoi(x); err == nil {
			if i < min || i > max {
				return nil, fmt.Errorf("out of range for %d in %s. %d must be in range %d-%d", i, s, i, min, max)
			}
			r[i] = struct{}{}
		} else {
			return nil, fmt.Errorf("unable to parse %s part in %s", x, s)
		}
	}

	if len(r) == 0 {
		return nil, fmt.Errorf("unable to parse %s", s)
	}

	return r, nil
}

// getTick returns the tick struct from time
func getTick(t time.Time) tick {
	return tick{
		min:       t.Minute(),
		hour:      t.Hour(),
		day:       t.Day(),
		month:     int(t.Month()),
		dayOfWeek: int(t.Weekday()),
	}
}
