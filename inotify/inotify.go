package inotify

import (
	"errors"
	"fmt"

	"github.com/deepglint/glog"
	"github.com/deepglint/util/filetool"
	"github.com/howeyc/fsnotify"
)

type PathWatcher struct {
	Path         string
	Watcher      *fsnotify.Watcher
	Stop         chan bool
	FileHandlers map[string]func()
}

func NewPathWatcher(path string, flags uint32, stopChan chan bool) (*PathWatcher, error) {
	pathWatcher := new(PathWatcher)
	pathWatcher.Stop = stopChan
	pathWatcher.FileHandlers = make(map[string]func())

	if !filetool.IsExist(path) {
		return nil, errors.New(fmt.Sprintf("'%s' not found", path))
	}
	pathWatcher.Path = path

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	err = watcher.WatchFlags(path, flags)
	if err != nil {
		return nil, err
	}
	pathWatcher.Watcher = watcher

	return pathWatcher, err
}

func (this *PathWatcher) AddFileEvent(file string, handler func()) {
	if !filetool.IsFile(file) {
		return
	}

	this.FileHandlers[file] = handler
}

func (this *PathWatcher) Run() {
	go func() {
		for {
			select {
			case <-this.Stop:
				glog.Infoln("close")
				break
			case event := <-this.Watcher.Event:
				glog.Infoln(event)

				for file, handler := range this.FileHandlers {
					if event.Name == file && event.IsCreate() { // || event.IsModify()
						go handler()
					}
				}
			case err := <-this.Watcher.Error:
				glog.Errorln(err.Error())
			}
		}
		glog.Infoln("yeah")
	}()
}

func (this *PathWatcher) Close() {
	this.Stop <- true

	this.Watcher.Close()
}
