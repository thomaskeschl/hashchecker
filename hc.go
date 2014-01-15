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
)

func main() {
	inputFilePtr := flag.String("f", "", "path to the file to generate the hash for")
	inputHashPtr := flag.String("c", "", "the hash string to compare against")
	inputTypePtr := flag.String("t", "md5", "the hashing function to use (valid options are md5, sha1)")
	flag.Parse()

	// ensure a file has been entered.
	if len(*inputFilePtr) == 0 {
		fmt.Printf("Need to enter a file to hash.")
		return
	} 

	filepath := *inputFilePtr

	f,err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	hashStr := *inputTypePtr

	hash := generateSum(hashStr)
	if hash == nil {
		fmt.Printf("No such hashing function: {0}", hashStr)
	}
	hash.Write(f)

	sum := hash.Sum(nil)
	sumStr := fmt.Sprintf("%x", sum)

	if len(*inputHashPtr) != 0 {
		compareStr := *inputHashPtr
		if compareStr == sumStr {
			fmt.Printf("They match!")
		} else {
			fmt.Printf("They don't match!")
		}
		return;
	}

	fmt.Printf(sumStr)
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