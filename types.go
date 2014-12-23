package myanimelist

import "encoding/xml"

type User struct {
	Id                string `xml:"user_id"`
	Name              string `xml:"user_name"`
	Watching          string `xml:"user_watching"`
	Completed         string `xml:"user_completed"`
	OnHold            string `xml:"user_onhold"`
	Dropped           string `xml:"user_dropped"`
	PlanToWatch       string `xml:"user_plantowatch"`
	DaysSpentWatching string `xml:"user_days_spent_watching"`
}

const (
	STATUS_WATCHING      = 1
	STATUS_WATCHED       = 2
	STATUS_PLAN_TO_WATCH = 6
)

type Anime struct {
	AnimedbId         string `xml:"series_animedb_id"`
	Title             string `xml:"series_title"`
	Synonyms          string `xml:"series_synonyms"`
	Type              string `xml:"series_type"`
	Episodes          string `xml:"series_episodes"`
	Status            string `xml:"series_status"`
	Start             string `xml:"series_start"`
	End               string `xml:"series_end"`
	Image             string `xml:"series_image"`
	MyId              string `xml:"my_id"`
	MyWatchedEpisodes string `xml:"my_watched_episodes"`
	MyStartDate       string `xml:"my_start_date"`
	MyFinishDate      string `xml:"my_finish_date"`
	MyScore           string `xml:"my_score"`
	MyStatus          string `xml:"my_status"`
	MyRewatching      string `xml:"my_rewatching"`
	MyRewatchingEp    string `xml:"my_rewatching_ep"`
	MyLastUpdated     string `xml:"my_last_updated"`
}

type Result struct {
	XMLName xml.Name `xml:"myanimelist"`
	MyInfo  User     `xml:"myinfo"`
	Animes  []Anime  `xml:"anime"`
}
