package crontab

import (
	cr "gopkg.in/robfig/cron.v2"
)

// Cron interface
type Cron interface {
	AddFunc(string, func())
	Stop()
	Start()
}

type cron struct {
	cr *cr.Cron
}

// New returns new Cron
func New() Cron {
	return &cron{cr.New()}
}

// AddFunc adds a func to the Cron to be run on the given schedule.
func (s *cron) AddFunc(expr string, f func()) {
	s.cr.AddFunc(expr, f)
}

// Stop the cron scheduler.
func (s *cron) Stop() {
	s.cr.Stop()
}

func (s *cron) Start() {
	s.cr.Start()
}
