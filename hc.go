// Copyright 2014 Thomas Keschl. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in
// the LICENSE file.

package main 

import(
	"io/ioutil"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"flag"
	"hash"
	"os"
)

func main() {
	inputFilePtr := flag.String("f", "", "path to the file to generate the hash for")
	inputHashPtr := flag.String("c", "", "the hash string to compare against")
	inputTypePtr := flag.String("t", "md5", "the hashing function to use (valid options are md5, sha1)")
	flag.Parse()

	// ensure a file has been entered.
	if len(*inputFilePtr) == 0 {
		fmt.Println("Need to enter a file to hash.")
		os.Exit(1)
	} 

	filepath := *inputFilePtr

	f,err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("%s", err)
		os.Exit(2)
	}

	hashStr := *inputTypePtr

	hash := generateSum(hashStr)
	if hash == nil {
		fmt.Println("No such hashing function: ", hashStr)
		os.Exit(3)
	}
	hash.Write(f)

	sum := hash.Sum(nil)
	sumStr := fmt.Sprintf("%x", sum)

	if len(*inputHashPtr) != 0 {
		compareStr := *inputHashPtr
		if compareStr == sumStr {
			fmt.Println("They match!")
		} else {
			fmt.Println("They don't match!")
		}
		os.Exit(0)
	}

	fmt.Println(sumStr)
	os.Exit(0)
}

func generateSum(hashStr string) hash.Hash {
	switch {
		case hashStr == "md5":
			return md5.New()
		case hashStr == "sha1":
			return sha1.New()
	}
	return nil
}