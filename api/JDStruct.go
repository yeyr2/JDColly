package api

type Response struct {
	StatusCode int32  `json:"status_code"`          // 0为正常,其他为异常
	StatusMsg  string `json:"status_msg,omitempty"` // 传给前端的信息
	Value      any    `json:"value"`
}

type Hot struct {
	Img       string `selector:"div.gl-i-wrap > div.p-img > a > img" attr:"data-lazy-img" json:"img,omitempty"`
	Price     string `selector:"div.gl-i-wrap > div.p-price > strong > i" json:"price,omitempty"`
	Name      string `selector:"div.gl-i-wrap > div.p-name> a > em" json:"name,omitempty"`
	ProductId string `json:"product_id,omitempty"`
	Title     string `selector:"div.gl-i-wrap > div.p-name > a" attr:"title" json:"title,omitempty"`
	Url       string `selector:"div.gl-i-wrap > div.p-img > a" attr:"href" json:"url"`
}

type AnalyzeComment struct {
	Fraction    int        `json:"fraction"`     //总评分数
	Interval    []Interval `json:"interval"`     //区间分数(5段)
	AnalyzeWord []string   `json:"analyze_word"` //评论词云
}

type Interval struct {
	Interval   int    `json:"interval"`    // 区间分数(占比)
	ScoreRange string `json:"score_range"` // 区间(例如:1 : [0.0,2.0) )
}

type JDComment struct {
	Score                 int                   `json:"score"`
	Comments              []Comments            `json:"comments"`
	MaxPage               int                   `json:"maxPage"`
	ProductCommentSummary ProductCommentSummary `json:"productCommentSummary"`
}

type Comments struct {
	EnContent       string
	Content         string `json:"content"`
	CreationTime    string `json:"creationTime"`
	Score           int    `json:"score"`
	UsefulVoteCount int    `json:"usefulVoteCount"` // 有效投票
	TextIntegral    int    `json:"textIntegral"`
	ReferenceTime   string `json:"referenceTime"`
	ReferenceImage  string `json:"referenceImage"`
	ReferenceName   string `json:"referenceName"`
}

type ProductCommentSummary struct {
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
