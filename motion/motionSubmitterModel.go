package motion

// Motion Submitter Base Class
// https://github.com/OpenSlides/OpenSlides/blob/master/client/src/app/shared/models/motions/motion-submitter.ts
type MotionSubmitter struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	MotionID int `json:"motion_id"`
	Weight   int `json:"weight"`
}
