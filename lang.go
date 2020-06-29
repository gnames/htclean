package htclean

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v2"
	"github.com/dustin/go-humanize"
	"github.com/gnames/htclean/cerror"
	"github.com/gnames/htclean/keyval"
	"github.com/gnames/htclean/model"
)

func (htc *HTclean) Lang() error {
	log.Println("Getting pages with names")
	kv, err := keyval.InitKeyVal(filepath.Join(htc.WorkPath, htc.KeyValPath))
	if err != nil {
		return err
	}
	kvTxn := kv.NewTransaction(true)

	path := filepath.Join(htc.WorkPath, htc.InputFile)
	f, err := os.Open(path)
	if err != nil {
		context := fmt.Sprintf("Problem with opening of '%s'", path)
		return cerror.NewErr(context, err)
	}
	r := csv.NewReader(f)
	count := 0
	for {
		l, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			continue
		}
		count += 1
		if count%10_000_000 == 0 {
			log.Printf("Reading %sth page", humanize.Comma(int64(count)))
		}
		titleID := l[model.TitleF]
		pageID := l[model.PageF]
		key := titleID + "|" + pageID
		val := ""
		if err = kvTxn.Set([]byte(key), []byte(val)); err == badger.ErrTxnTooBig {
			err = kvTxn.Commit()
			if err != nil {
				return err
			}
			kvTxn = kv.NewTransaction(true)
			err = kvTxn.Set([]byte(key), []byte(val))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
