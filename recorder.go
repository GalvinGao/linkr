package main

import "net/http"

func recordLinkVisit(r *http.Request, linkId uint) error {
	return DB.Create(&LinkRecord{
		ParentLinkID: linkId,
		Referer:      r.Referer(),
		UserAgent:    r.UserAgent(),
	}).Error
}
