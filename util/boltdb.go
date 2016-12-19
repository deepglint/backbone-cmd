package util

import (
	"errors"

	"github.com/boltdb/bolt"
	//"github.com/revel/revel"
	//"log"
	"encoding/gob"
	//"strconv"
	// "binary"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"time"
)

type BoltDBController struct {
	Db   *bolt.DB
	Path string
}

func NewBoltDBController(path string) *BoltDBController {
	return &BoltDBController{
		Path: path,
	}
}

func (this *BoltDBController) SaveInterface(table []byte, key []byte, v interface{}) error {

	b, err := this.ToByte(v)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return this.SetValue(table, key, b)
}
func (this *BoltDBController) GetInterface(table []byte, key []byte) (interface{}, error) {
	b, err := this.GetValue(table, key)
	if err != nil {
		return nil, err
	}
	return this.ToInterface(b)
}

func (this *BoltDBController) ToByte(s interface{}) ([]byte, error) {
	var buf = &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	// err := binary.Write(buf, binary.BigEndian, &s)
	err := enc.Encode(&s)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	return buf.Bytes(), nil

}

func (this *BoltDBController) ToInterface(b []byte) (interface{}, error) {
	var buf = &bytes.Buffer{}
	var s interface{}
	_, err := buf.Read(b)
	if err != nil {
		return s, err
	}
	var dec = gob.NewDecoder(buf)
	//err = binary.Read(buf, binary.BigEndian, &s)
	err = dec.Decode(&s)
	return s, err
}

func (this *BoltDBController) ApplyId(table string) (string, error) {
	var id string
	err := this.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			//log.Println(err)
			return err
		}
		seq, err := b.NextSequence()
		id = strconv.Itoa(int(seq))
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (this *BoltDBController) InitDB() error {
	var err error
	this.Db, err = bolt.Open(this.Path, 0777, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Println(err)
	}
	return err
}

func (this *BoltDBController) CloseDB() error {
	if this.Db != nil {
		err := this.Db.Close()
		return err
	}
	return errors.New("The DB is null")
}

func (this *BoltDBController) NewTable(name string) error {
	err := this.Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return err
}

func (this *BoltDBController) GetValue(table []byte, key []byte) ([]byte, error) {
	var result []byte
	err := this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return errors.New("The Table is Null")
		}
		result = b.Get([]byte(key))
		if len(result) <= 0 || result == nil {

			return errors.New("The Value Not Found")
		}
		return nil
	})
	return result, err
}

func (this *BoltDBController) IfTableExist(table []byte) bool {
	var res bool
	_ = this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(table)
		if b == nil {
			res = false
			return nil
		}
		res = true
		return nil
	})
	return res
}

func (this *BoltDBController) IfValueExist(table []byte, key []byte) bool {
	var result []byte
	var res bool
	_ = this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return errors.New("The Table is Null")
		}
		result = b.Get([]byte(key))
		if len(result) <= 0 || result == nil {
			res = false
			return nil
		}
		res = true
		return nil
	})
	return res

}

func (this *BoltDBController) SetValue(table []byte, key []byte, val []byte) error {
	err := this.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			//log.Println(err)
			return err
		}
		err = b.Put([]byte(key), val)
		return err
	})
	return err
}

func (this *BoltDBController) DeleteValue(table string, key string) error {
	err := this.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			//log.Println(err)
			return err
		}
		err = b.Delete([]byte(key))
		return err
	})
	return err
}

func (this *BoltDBController) GetTable(table string) (map[string][]byte, error) {
	var result = make(map[string][]byte, 0)
	err := this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return errors.New("The Table is Null")
		}
		b.ForEach(func(k, v []byte) error {
			//fmt.Printf("key=%s, value=%s\n", k, v)
			result[string(k)] = v
			return nil
		})
		return nil
	})
	return result, err
}
func (this *BoltDBController) SetValueAuto(table []byte, val []byte) (string, error) {
	var id = -1
	err := this.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			return errors.New("create bucket:" + err.Error())
		}
		seq, err := b.NextSequence()
		id = int(seq)
		if err != nil {
			return err
		}
		err = b.Put([]byte(strconv.Itoa(id)), val)
		if err != nil {
			return errors.New("put into the table:" + err.Error())
		}
		return nil
	})
	return string(strconv.Itoa(id)), err
}

///SetValueAuto
