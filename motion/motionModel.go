package motion

// Motion Base Type
// https://github.com/OpenSlides/OpenSlides/blob/master/client/src/app/shared/models/motions/motion.ts
type Motion struct {
	ID                           int               `json:"id"`
	Identifier                   string            `json:"identifier"`
	Title                        string            `json:"title"`
	Text                         string            `json:"text"`
	Reason                       string            `json:"reason"`
	AmendmentParagraphs          []string          `json:"amendment_paragraphs"`
	ModifiedFinalVersion         string            `json:"modified_final_version"`
	ParentID                     int               `json:"parent_id"`
	CategoryID                   int               `json:"category_id"`
	MotionBlockID                int               `json:"motion_block_id"`
	Origin                       string            `json:"origin"`
	Submitters                   []MotionSubmitter `json:"submitters"` //MotionSubmitter[]
	SupportersID                 []int             `json:"supporters_id"`
	Comments                     []string          `json:"comments"` //MotionComment[]
	WorkflowID                   int               `json:"workflow_id"`
	StateID                      int               `json:"state_id"`
	StateExtension               string            `json:"state_extension"`
	StateRequiredPermissionToSee string            `json:"state_required_permission_to_see"`
	StatuteParagraphID           int               `json:"statute_paragraph_id"`
	RecommendationID             int               `json:"recommendation_id"`
	RecommendationExtension      string            `json:"recommendation_extension"`
	TagsID                       []int             `json:"tags_id"`
	AttachmentsID                []int             `json:"attachments_id"`
	Polls                        []MotionPoll      `json:"polls"` //MotionPoll[]
	AgendaItemID                 int               `json:"agenda_item_id"`
	LogMessages                  []string          `json:"log_messages"` //MotionLog[]
	Weight                       int               `json:"weight"`
	SortParentID                 int               `json:"sort_parent_id"`
	Created                      string            `json:"created"`
	LastModified                 string            `json:"last_modified"`
}
