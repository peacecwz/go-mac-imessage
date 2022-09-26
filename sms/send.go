package sms

import (
	"fmt"
	"github.com/andybrewer/mack"
)

func Send(message, number string) error {
	args := []string{
		fmt.Sprintf("set targetBuddy to \"%s\"", number),
		"set targetService to id of 1st account whose service type = iMessage",
		"set theBuddy to participant targetBuddy of account id targetService",
		fmt.Sprintf("send \"%s\" to theBuddy", message)}
	_, err := mack.Tell("Messages", args...)
	if err != nil {
		return err
	}

	return nil
}
