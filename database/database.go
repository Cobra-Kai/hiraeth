// Copyright (c) 2019 Jon Mayo <jon@rm-f.net>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
package database

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func InitReadonly(dbfilename string) {
	db, err := bolt.Open(dbfilename, 0600, &bolt.Options{Timeout: 5 * time.Second, ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func Init(dbfilename string) {

	db, err := bolt.Open(dbfilename, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("System"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		/* write the system boot-up time */
		err = b.Put([]byte("uptime-since"), []byte(time.Now().Format(time.RFC3339)))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("System"))
		v := b.Get([]byte("uptime-since"))
		log.Println("Uptime since", string(v))
		return nil
	})
}

// TODO: output to a writer
func Dump(bucket string) {
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucket))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
}
