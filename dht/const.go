package dht

import (
	"time"
)

const K int = 8

const ALPHA int = 3

const (
	NEW    = iota
	SENIOR = iota
)

var BOOTSTRAP []string = []string{
	"dht.transmissionbt.com:6881",
	"service.ygrek.org.ua:6881",
	"router.utorrent.com:6881",
	"router.transmission.com:6881"}

const MAXSIZE = 500

const (
	GOOD           = iota
	QUESTIONABLE_1 = iota //MISS once
	QUESTIONABLE_2 = iota //MISS twice
	BAD            = iota
)

//username:password@protocol(address)/dbname?param=value
const DSN = "root:123456@unix(/opt/local/var/run/mysql55/mysqld.sock)/dhtrobot?charset=utf8"

const (
	NODENUM = 100
)

const (
	EXPIRE_DURATION = time.Minute * 25
)
