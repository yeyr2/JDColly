package api

type Hot struct {
	Img       string `selector:"div.gl-i-wrap > div.p-img > a > img" attr:"data-lazy-img" json:"img,omitempty"`
	Price     string `selector:"div.gl-i-wrap > div.p-price > strong > i" json:"price,omitempty"`
	Name      string `selector:"div.gl-i-wrap > div.p-name> a > em" json:"name,omitempty"`
	ProductId string `json:"productId,omitempty"`
	Title     string `selector:"div.gl-i-wrap > div.p-name > a" attr:"title" json:"title,omitempty"`
	Url       string `selector:"div.gl-i-wrap > div.p-img > a" attr:"href"`
}

type AnalyzeComment struct {
	fraction string     //总评分数
	interval []Interval //区间分数(5段)
	analyze  []string   //评论词云
}

type Interval struct {
	interval   string // 区间分数(占比)
	ScoreRange string // 区间(例如:[0.0,2.0) )
}

type JDComment struct {
	JwotestProduct          interface{}               `json:"jwotestProduct"`
	Score                   int                       `json:"score"`
	Comments                []Comments                `json:"comments"`
	SoType                  int                       `json:"soType"`
	Csv                     string                    `json:"csv"`
	ImageListCount          int                       `json:"imageListCount"`
	HotCommentTagStatistics []HotCommentTagStatistics `json:"hotCommentTagStatistics"`
	TestID                  string                    `json:"testId"`
	VTagStatistics          interface{}               `json:"vTagStatistics"`
	MaxPage                 int                       `json:"maxPage"`
	ProductCommentSummary   ProductCommentSummary     `json:"productCommentSummary"`
	ProductAttr             interface{}               `json:"productAttr"`
}

type Images struct {
	ID       int    `json:"id"`
	ImgURL   string `json:"imgUrl"`
	ImgTitle string `json:"imgTitle"`
	Status   int    `json:"status"`
}

type Videos struct {
	ID          int    `json:"id"`
	MainURL     string `json:"mainUrl"`
	VideoHeight int    `json:"videoHeight"`
	VideoWidth  int    `json:"videoWidth"`
	VideoLength int    `json:"videoLength"`
	VideoTitle  string `json:"videoTitle"`
	VideoURL    int    `json:"videoUrl"`
	VideoID     int    `json:"videoId"`
	Status      int    `json:"status"`
	Remark      string `json:"remark"`
}

type ExtMap struct {
	BuyCount int `json:"buyCount"`
}

type Comments struct {
	ID               int64         `json:"id"`
	GUID             string        `json:"guid"`
	Content          string        `json:"content"`
	CreationTime     string        `json:"creationTime"`
	IsDelete         bool          `json:"isDelete"`
	IsTop            bool          `json:"isTop"`
	UserImageURL     string        `json:"userImageUrl"`
	Topped           int           `json:"topped"`
	Replies          []interface{} `json:"replies"`
	ReplyCount       int           `json:"replyCount"`
	Score            int           `json:"score"`
	ImageStatus      int           `json:"imageStatus"`
	UsefulVoteCount  int           `json:"usefulVoteCount"`
	UserClient       int           `json:"userClient"`
	DiscussionID     int           `json:"discussionId"`
	ImageCount       int           `json:"imageCount"`
	AnonymousFlag    int           `json:"anonymousFlag"`
	PlusAvailable    int           `json:"plusAvailable"`
	MobileVersion    string        `json:"mobileVersion"`
	Images           []Images      `json:"images"`
	Videos           []Videos      `json:"videos"`
	MergeOrderStatus int           `json:"mergeOrderStatus"`
	ProductColor     string        `json:"productColor"`
	ProductSize      string        `json:"productSize"`
	TextIntegral     int           `json:"textIntegral"`
	ImageIntegral    int           `json:"imageIntegral"`
	ExtMap           ExtMap        `json:"extMap"`
	Status           int           `json:"status"`
	ReferenceID      string        `json:"referenceId"`
	ReferenceTime    string        `json:"referenceTime"`
	Nickname         string        `json:"nickname"`
	ReplyCount2      int           `json:"replyCount2"`
	UserImage        string        `json:"userImage"`
	OrderID          int           `json:"orderId"`
	Integral         int           `json:"integral"`
	ProductSales     string        `json:"productSales"`
	ReferenceImage   string        `json:"referenceImage"`
	ReferenceName    string        `json:"referenceName"`
	FirstCategory    int           `json:"firstCategory"`
	SecondCategory   int           `json:"secondCategory"`
	ThirdCategory    int           `json:"thirdCategory"`
	AesPin           interface{}   `json:"aesPin"`
	Days             int           `json:"days"`
	AfterDays        int           `json:"afterDays"`
}
type HotCommentTagStatistics struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Count          int    `json:"count"`
	Type           int    `json:"type"`
	CanBeFiltered  bool   `json:"canBeFiltered"`
	Stand          int    `json:"stand"`
	Rid            string `json:"rid"`
	CkeKeyWordBury string `json:"ckeKeyWordBury"`
}
type ProductCommentSummary struct {
	SkuID               int64   `json:"skuId"`
	AverageScore        int     `json:"averageScore"`
	DefaultGoodCount    int     `json:"defaultGoodCount"`
	DefaultGoodCountStr string  `json:"defaultGoodCountStr"`
	CommentCount        int     `json:"commentCount"`
	CommentCountStr     string  `json:"commentCountStr"`
	GoodCount           int     `json:"goodCount"`
	GoodCountStr        string  `json:"goodCountStr"`
	GoodRate            float64 `json:"goodRate"`
	GoodRateShow        int     `json:"goodRateShow"`
	GeneralCount        int     `json:"generalCount"`
	GeneralCountStr     string  `json:"generalCountStr"`
	GeneralRate         float64 `json:"generalRate"`
	GeneralRateShow     int     `json:"generalRateShow"`
	PoorCount           int     `json:"poorCount"`
	PoorCountStr        string  `json:"poorCountStr"`
	PoorRate            float64 `json:"poorRate"`
	PoorRateShow        int     `json:"poorRateShow"`
	VideoCount          int     `json:"videoCount"`
	VideoCountStr       string  `json:"videoCountStr"`
	AfterCount          int     `json:"afterCount"`
	AfterCountStr       string  `json:"afterCountStr"`
	ShowCount           int     `json:"showCount"`
	ShowCountStr        string  `json:"showCountStr"`
	OneYear             int     `json:"oneYear"`
	SensitiveBook       int     `json:"sensitiveBook"`
	FixCount            int     `json:"fixCount"`
	PlusCount           int     `json:"plusCount"`
	PlusCountStr        string  `json:"plusCountStr"`
	BuyerShow           int     `json:"buyerShow"`
	PoorRateStyle       int     `json:"poorRateStyle"`
	GeneralRateStyle    int     `json:"generalRateStyle"`
	GoodRateStyle       int     `json:"goodRateStyle"`
	InstallRate         int     `json:"installRate"`
	ProductID           int64   `json:"productId"`
	Score1Count         int     `json:"score1Count"`
	Score2Count         int     `json:"score2Count"`
	Score3Count         int     `json:"score3Count"`
	Score4Count         int     `json:"score4Count"`
	Score5Count         int     `json:"score5Count"`
}
