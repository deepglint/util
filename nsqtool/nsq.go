package nsqtool

import (
	"time"

	"github.com/bitly/go-nsq"
	"github.com/deepglint/glog"
)

const (
	RE_INTERVAL int64 = 100
)

func NewProducer(config *nsq.Config, nsqaddr string, stopChan chan int) *nsq.Producer {
	if nsqaddr == "" {
		return nil
	}

	interval := RE_INTERVAL
	producer, err := nsq.NewProducer(nsqaddr, config)
	for {
		if err != nil {
			glog.Warningf("Error connecting to nsqd %s", nsqaddr)

			select {
			case <-stopChan:
				return nil
			case <-time.After(time.Duration(interval) * time.Millisecond):
				interval *= 2
				producer, err = nsq.NewProducer(nsqaddr, config)
				continue
			}

		}
		break
	}

	interval = RE_INTERVAL
	err = producer.Ping()
	for {
		if err != nil {
			glog.Errorf("Can not ping nsqd %s", nsqaddr)
			select {
			case <-stopChan:
				return nil
			case <-time.After(time.Duration(interval) * time.Millisecond):
				interval *= 2
				err = producer.Ping()
				continue
			}
		}
		return producer
	}
}
