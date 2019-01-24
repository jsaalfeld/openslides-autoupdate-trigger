package motion

// MotionPoll base model
// https://github.com/OpenSlides/OpenSlides/blob/master/client/src/app/shared/models/motions/motion-poll.ts
type MotionPoll struct {
	ID           int  `json:"id"`
	Yes          int  `json:"yes"`
	No           int  `json:"no"`
	Abstain      int  `json:"abstain"`
	VotesValid   int  `json:"votesvalid"`
	VotesInvalid int  `json:"votesinvalid"`
	VotesCast    int  `json:"votescast"`
	HasVotes     bool `json:"has_votes"`
	MotionID     int  `json:"motion_id"`
}
