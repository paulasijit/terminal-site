package structDefination

type Author struct {
	FirstName string
	LastName  string
	Rating    float64
	TotalPost int64
}

type Contact struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	State       string
	Business    bool
	Message     string
}

type PostTags struct {
	Genre      string
	Popularity float64
	Rank       int64
	TotalPost  int64
	Like       int64
	Dislike    int64
}

type Post struct {
	Title      string
	Subheader  string
	Body       string
	Author     Author
	PublisedAt int64
	ReadTime   int64
	Tags       PostTags
}
