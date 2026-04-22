package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type YoutubeStats struct {
	Subscribers    int    `json:"subscribers"`
	ChannelName    string `json:"channelName"`
	MinutesWatched int    `json:"minutesWatched"`
	Views          int    `json:"views"`
}

func getChannelStats() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		yt := YoutubeStats{
			Subscribers:    5,
			ChannelName:    "yosenup",
			MinutesWatched: 50,
			Views:          111,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(yt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
