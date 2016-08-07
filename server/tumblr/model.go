package tumblr

// Post is tumblr post struct
type Post struct {
	BlogName  string   `json:"blog_name"`
	ID        int64    `json:"id"`
	PostURL   string   `json:"post_url"`
	Slug      string   `json:"slug"`
	Type      string   `json:"type"`
	Timestamp int64    `json:"timestamp"`
	Date      string   `json:"date"`
	Format    string   `json:"format"`
	ReblogKey string   `json:"reblog_key"`
	Tags      []string `json:"tags"`
	//Bookmarklet    bool     `json:"bookmarklet"`
	//Mobile         bool     `json:""`
	SourceURL      string `json:"source_url"`
	SourceTitle    string `json:"source_title"`
	Liked          bool   `json:"liked"`
	Followed       bool   `json:"followed"`
	NoteCount      int    `json:"note_count"`
	Caption        string `json:"caption"`
	ImagePermalink string `json:"image_permalink"`
	LinkUrl        string `json:"link_url"`
	// Post
	Title string `json:"title"`
	Body  string `json:"body"`
	// Photo
	Photos []Photo `json:"photos"`
	// Quote
	Text   string `json:"text"`
	Source string `json:"source"`
}

// Photo is post's photo struct
type Photo struct {
	Caption  string    `json:"caption"`
	AltSizes []AltSize `json:"alt_sizes"`
}

// AltSize is photo's alt_sizez struct
type AltSize struct {
	Width  int    `json:""`
	Height int    `json:""`
	URL    string `json:""`
}
