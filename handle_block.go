/*
Copyright 2020 The Board of Trustees of The Leland Stanford Junior University

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package lzr

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/porfirion/trie"
)

func BuildTrie() (*trie.Trie[struct{}], error) {
	filePath := getBlockedList()
	// fmt.Fprintf(os.Stderr, "FILE PATH: |%s|", filePath)

	if filePath == "" {
		return nil, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	trie := &trie.Trie[struct{}]{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := scanner.Text()
		// skip empty lines
		if ip == "" {
			fmt.Printf("EMPTY LINE FOUND! \n")
			continue
		}
		if _, ipNet, err := net.ParseCIDR(ip); err == nil {
			prefix := getCIDRPrefix(ipNet)
			// dont readd/add the prefix if it is a subnet of an existing prefix
			if _, _, found := trie.SearchPrefixInString(prefix); !found {
				trie.PutString(prefix, struct{}{})
			}
		} else if ip := net.ParseIP(ip); ip != nil {
			// dont readd/add the prefix if it is a subnet of an existing prefix
			if _, _, found := trie.SearchPrefixInString(ip.String()); !found {
				trie.PutString(ip.String(), struct{}{})
			}

		} else {
			fmt.Printf("Skipping invalid entry %s\n", ip)
		}
	}

	// for testing the trie
	// trie.Iterate(func(prefix []byte, value struct{}) {
	// 	fmt.Println(string(prefix))
	// })

	return trie, nil
}

func getCIDRPrefix(ipNet *net.IPNet) string {
	maskSize, _ := ipNet.Mask.Size()
	bytes := maskSize / 8
	newIP := strings.Split(ipNet.IP.String(), ".")
	return strings.Join(newIP[:bytes], ".")
}

func isBlocked(trie *trie.Trie[struct{}], ip string) bool {
	_, _, status := trie.SearchPrefixInString(ip)
	return status
}
