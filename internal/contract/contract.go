package contract

import "github.com/LOLONEKIT2019/prof_of_work/pgk/message"

const (
	RequestTask message.Type = iota + 1
	RequestQuote
	ResponseTask
	ResponseQuote
	ResponseError
)
