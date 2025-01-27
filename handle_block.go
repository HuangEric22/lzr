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
    "os"
    "strings"
	"net"
	"github.com/yl2chen/cidranger"
)

func BuildTrie() (cidranger.Ranger, error) {
    filePath := getBlockedList()

    if filePath == "" {
        return nil, nil
    }

    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    trie := cidranger.NewPCTrieRanger();

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        ip := scanner.Text()
        // skip empty lines
        if ip == "" {
            fmt.Printf("EMPTY LINE FOUND! \n")
            continue
        }

        if idx := strings.Index(ip, "#"); idx != -1 {
            ip = ip[:idx]
        }

        ip = strings.TrimSpace(ip)

		_, prefix, err := net.ParseCIDR(ip);

		if err != nil {
			return nil, fmt.Errorf("failed to parse CIDR '%s': %w", ip, err)
		}
		
		
		err = trie.Insert(cidranger.NewBasicRangerEntry(*prefix))
		if err != nil {
			return nil, fmt.Errorf("failed to insert CIDR '%s': %w", ip, err)
		}		
	}

	// For debugging and printing out the trie
	// allCIDRs, err := trie.CoveredNetworks(net.IPNet{
	// 	IP:   net.IPv4(0, 0, 0, 0),
	// 	Mask: net.CIDRMask(0, 32),
	// })
	// fmt.Println("CIDR Blocks in the Ranger:")
	// for _, cidr := range allCIDRs {
	// 	fmt.Println(cidr)
	// }

    fmt.Println("BLOCK LIST INITIALIZED")

	return trie, nil;
}

func isBlocked(trie cidranger.Ranger, ip string) bool {
	status, _ := trie.Contains(net.ParseIP(ip));
	return status;
}