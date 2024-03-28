package storage 

import (
	"fmt"
	"strings"
)

type ConnectionString struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Opts     []string
}

func NewConnectionString(opts ...func(*ConnectionString)) ConnectionString {
	cs := ConnectionString{}

	for _, f := range opts {
		f(&cs)
	}

	return cs
}

func WithDriver(driver string) func(*ConnectionString) {
	return func(cs *ConnectionString) {
		cs.Driver = driver
	}
}

func WithUser(user string) func(*ConnectionString) {
	return func(cs *ConnectionString) {
		cs.User = user
	}
}

func WithPassword(password string) func(*ConnectionString) {
	return func(cs *ConnectionString) {
		cs.Password = password
	}
}

func WithHost(host string) func(*ConnectionString) {
	return func(cs *ConnectionString) {
		cs.Host = host
	}
}
func WithPort(port int) func(*ConnectionString) {
	return func(cs *ConnectionString) {
		cs.Port = port
	}
}

func WithDatabase(database string) func(*ConnectionString) {
	return func(cs *ConnectionString) {
		cs.Database = database
	}
}

func WithOpts(opts map[string]string) func(*ConnectionString) {
	return func(cs *ConnectionString) {
		for k, v := range opts {
			cs.Opts = append(cs.Opts, fmt.Sprintf("%s=%s", k, v))
		}
	}
}

func (cs ConnectionString) String() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?%s", cs.Driver, cs.User, cs.Password, cs.Host, cs.Port, cs.Database, strings.Join(cs.Opts, "&"))
}
