package sysvipc

/*
#cgo LDFLAGS: -lstdc++ -L.
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/sem.h>
int semget(key_t key, int nsems, int semflg);


union arg4 {
	int             val;
	struct semid_ds *buf;
	unsigned short  *array;
};

static struct sembuf p_buf_ = { 0, -1, SEM_UNDO };
static struct sembuf v_buf_ = { 0, 1, SEM_UNDO };


int sem_p(int semid) {//-1 means failed
	int ret = semop(semid, &p_buf_, 1);
	return ret;
};

int sem_v(int semid) {//-1 means failed
	int ret = semop(semid, &v_buf_, 1);
	return ret;
};

int getVal(int semid) {//-1 means failed
	int val = semctl(semid, 0, GETVAL, 0);
	return val;
};

int setVal(int semid,int val) {//-1 means failed
	union arg4 arg;
	arg.val = val;
	return semctl(semid, 0, SETVAL, arg);
};

*/
import "C"
import (
	"errors"
	//"fmt"
	"time"
)

// SemaphoreSet is a kernel-maintained collection of semaphores.
type SemaphoreSet struct {
	id    int64
	count uint
}

// GetSemSet creates or retrieves the semaphore set for a given IPC key.
func GetSemSet(key, count int64, flags *SemSetFlags) (*SemaphoreSet, error) {
	rc, err := C.semget(C.key_t(key), C.int(count), C.int(flags.flags()))
	if rc == -1 {
		return nil, err
	}
	return &SemaphoreSet{int64(rc), uint(count)}, nil
}

func (ss *SemaphoreSet) P_Timeout(s int) error {
	rs := make(chan int)
	time.AfterFunc(time.Second*time.Duration(s), func() {
		select {
		case rs <- -2:

		default:
			return
		}

	})
	go func() {
		rc := C.sem_p(C.int(ss.id))
		rs <- int(rc)
		// if rc == -1 {
		// 	return errors.New("p faild")
		// }
	}()
	select {
	case r := <-rs:
		if r == -1 {
			return errors.New("p faild")
		}
		if r == -2 {
			//go ss.V()
			return errors.New("time out")
		}
	}
	return nil
}

func (ss *SemaphoreSet) P() error {

	rc := C.sem_p(C.int(ss.id))
	if rc == -1 {
		return errors.New("p faild")
	}
	return nil

}

func (ss *SemaphoreSet) V() error {
	rc := C.sem_v(C.int(ss.id))
	if rc == -1 {
		return errors.New("v faild")
	}
	return nil
}

func (ss *SemaphoreSet) Get() int {
	rc := C.getVal(C.int(ss.id))
	return int(rc)
}

func (ss *SemaphoreSet) Set(v int) int {
	rc := C.setVal(C.int(ss.id), C.int(v))
	return int(rc)
}

// // Remove deletes the semaphore set.
// // This will also awake anyone blocked on the set with EIDRM.
// func (ss *SemaphoreSet) Remove() error {
// 	rc, err := C.semctl_noarg(C.int(ss.id), 0, C.IPC_RMID)
// 	if rc == -1 {
// 		return err
// 	}
// 	return nil
// }

// SemOps is a collection of operations submitted to SemaphoreSet.Run.

// SemSetFlags holds the options for a GetSemSet() call
type SemSetFlags struct {
	// Create controls whether to create the set if it doens't already exist.
	Create bool

	// Exclusive causes GetSemSet to fail if the semaphore set already exists
	// (only useful with Create).
	Exclusive bool

	// Perms is the file-style (rwxrwxrwx) permissions with which to create the
	// semaphore set (also only useful with Create).
	Perms int
}

func (sf *SemSetFlags) flags() int64 {
	if sf == nil {
		return 0
	}

	var f int64 = int64(sf.Perms) & 0777
	if sf.Create {
		f |= int64(C.IPC_CREAT)
	}
	if sf.Exclusive {
		f |= int64(C.IPC_EXCL)
	}

	return f
}
