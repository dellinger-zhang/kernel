package util

import (
	"archive/zip"
	"bytes"
	"compress/zlib"
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

// Gob encode
func GobEncode(val interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buffer)
	if err := enc.Encode(val); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// Gob decode
func GobDecode(buf []byte, val interface{}) error {

	dec := gob.NewDecoder(bytes.NewBuffer(buf))

	if err := dec.Decode(val); err != nil {
		return err
	}

	return nil
}

// zlib compress
func ZlibCompress(buf []byte) ([]byte, error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)

	if _, err := w.Write(buf); err != nil {
		return nil, err
	}

	if err := w.Flush(); err != nil {
		fmt.Println("ZlibCompress:fail to flush,", err)
	}

	if err := w.Close(); err != nil {
		fmt.Println("ZlibCompress:fail to close,", err)
	}

	return in.Bytes(), nil
}

// zlib decompress
func ZlibUncompress(buf []byte) ([]byte, error) {

	var out bytes.Buffer
	in := bytes.NewReader(buf)
	r, err := zlib.NewReader(in)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// zlib the whole dir into target zip file
func ZipDir(dir, zipFile string) error {

	f, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	d, err := os.OpenFile(dir, os.O_RDONLY, 0655)
	//d, err := os.Create(dir)
	if err != nil {
		return err
	}

	if err = compress(d, "", writer); err != nil {
		return err
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {

	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}