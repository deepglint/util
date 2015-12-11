package targz

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/deepglint/glog"
	"github.com/deepglint/muses/autobot/utils/filetool"
)

/*
   compress 'path' dir to 'file'.tar.gz
   using tar & gzip
*/
func CompressTargz(file, path string) (err error) {
	// create a writer for the target file
	fw, err := os.Create(file)
	if err != nil {
		return
	}
	defer fw.Close()

	// create a gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// create a tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// open the source dir
	dir, err := os.Open(path)
	if err != nil {
		return
	}
	defer dir.Close()

	// read the file list
	filetool.ResetFileStrings()
	filetool.FilesOfPath(path, false)

	// fis, err = dir.Readdir(0)
	// if err != nil {
	// 	return
	// }

	// iterate the file list
	for _, fileString := range filetool.FileStrings {
		// print filename
		glog.Infoln(fileString)

		// open the file
		fc, err := os.Open(fileString)
		if err != nil {
			glog.Infoln(err)
			continue
		}
		defer fc.Close()

		fi, err := fc.Stat()
		if err != nil {
			glog.Infoln(err)
			continue
		}

		// head
		h := new(tar.Header)
		// h.Name = fi.Name()
		// h.Name = fileString
		index := strings.LastIndex(fileString, "./")
		// glog.Infoln (fileString[index+1:])
		h.Name = fileString[index+1:]
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()

		// write head
		err = tw.WriteHeader(h)
		if err != nil {
			glog.Infoln(err)
			continue
		}

		// write content
		_, err = io.Copy(tw, fc)
		if err != nil {
			glog.Infoln(err)
			continue
		}
	}

	return
}

/*
	uncompress 'file' tar.gz to 'path' dir
	using tar & gzip
*/
func UncompressTargz(file, path string) (err error) {
	// open tar.gz file
	f, err := os.Open(file)
	if err != nil {
		glog.Infoln(err)
		return
	}
	defer f.Close()

	// read from gzip
	gr, err := gzip.NewReader(f)
	if err != nil {
		glog.Infoln(err)
		return
	}
	defer gr.Close()

	// read from tar
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			glog.Infoln(err)
			if err == io.EOF {
				break
			}
			return err
		}

		hdr_path := filepath.Dir(hdr.Name)
		// glog.Infoln (hdr_path)

		err = filetool.CreateDirRecursively(filepath.Join(path, hdr_path), os.FileMode(0777))
		if err != nil {
			glog.Infoln(err)
			return err
		}

		curfile := filepath.Join(path + string(os.PathSeparator) + hdr.Name)

		// if filetool.IsFile(curfile) {
		fw, err := os.OpenFile(curfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, hdr.FileInfo().Mode())
		if err != nil {
			glog.Infoln(err)
			return err
		}
		defer fw.Close()

		_, err = io.Copy(fw, tr)
		if err != nil {
			glog.Infoln(err)
			return err
		}
		// }
	}
	return nil
}
