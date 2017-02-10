package controller

import (
	"log"
	"os"

	util "git.coding.net/deepglint/devo/app/db"
)

func BoltLs(boltsrc, table string) {
	bc := util.NewBoltDBController(boltsrc)
	e := bc.InitDB()
	defer bc.CloseDB()
	if e != nil {
		log.Println(e.Error())
		return
	}
	t, e := bc.GetTable(table)
	if e != nil {
		log.Println(e.Error())
		return
	}
	print("key")
	print("\t" + "value")
	print("\n")
	for k, v := range t {
		print(string(k))
		print("\t" + string(v))
		print("\n")
	}
	return
}

func BoltLsTable(boltsrc string) {
	bc := util.NewBoltDBController(boltsrc)
	e := bc.InitDB()
	if e != nil {
		log.Println(e.Error())
		return
	}
	defer bc.CloseDB()
	t, e := bc.GetTables()
	if e != nil {
		log.Println(e.Error())
		return
	}
	for _, v := range t {
		println(v)
	}
}

func BoltLs2File(boltsrc, table string, dest string) {
	bc := util.NewBoltDBController(boltsrc)
	e := bc.InitDB()
	defer bc.CloseDB()
	if e != nil {
		log.Println(e.Error())
		return
	}
	t, e := bc.GetTable(table)
	if e != nil {
		log.Println(e.Error())
		return
	}
	// print("key")
	// print("\t" + "value")
	// print("\n")

	if dest == "" {
		println("Error for file name")
		return
	}
	f, e := os.Create(dest)
	if e != nil {
		log.Println(e.Error())
		return
	}
	for k, v := range t {
		f.WriteString(string(k))
		f.WriteString("\t" + string(v))
		f.WriteString("\n")
	}
	f.Close()
	return
}
