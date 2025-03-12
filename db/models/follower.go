package models

type Follower struct {
	Id         		uint `gorm:"primaryKey"`
	FollowerId 		uint `gorm:"index:follower_unique"`
	Follower   		User `gorm:"constraint:OnDelete:CASCADE"`
	FollowingId		uint `gorm:"index:follower_unique"`
	Following		User `gorm:"constraint:OnDelete:CASCADE"`
}
