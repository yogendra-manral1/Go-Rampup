package socialize

type FollowPayload struct{
	UserIds	[]uint	`json:"user_ids" validate:"required"`
}
