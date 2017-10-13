package sysvipc

// import (
// 	"fmt"
// 	"yadda/shm/sysvipc"
// )

// func main() {
// 	//sysvipcsysv.GetSharedMem(128, 128, flags*SHMFlags)(*SharedMem, error)
// 	m, err := sysvipc.GetSharedMem(128, 128, &sysvipc.SHMFlags{Create: true, Perms: 0666}) //(*SharedMem, error)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	mount, err := m.Attach(&sysvipc.SHMAttachFlags{ReadOnly: true})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	p := make([]byte, 128)
// 	l, err := mount.Read(p)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	sem, err := sysvipc.GetSemSet(111, 1, &sysvipc.SemSetFlags{Create: true, Exclusive: false, Perms: 0666})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	err = sem.P()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(l, string(p))
// 	err = sem.V()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
