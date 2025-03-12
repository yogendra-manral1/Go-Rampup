package socialize

import (
	"Go-Rampup/db/models"
	"net/http"
	"slices"
	"errors"
	// "golang.org/x/crypto/bcrypt"
	// "gorm.io/gorm"
)

func (ctrl SocializeController) Follow(email string, userIds []uint) (int, []uint, error) {
	var currentUser models.User
	err := currentUser.GetUser(ctrl.DB, [][]string{{"email = ?", email}})
	if err != nil {
		return http.StatusNotFound, nil, err
	}
	var users []models.User
	ctrl.DB.Find(&users, userIds)
	var validUsersId []uint
	for _, user := range users {
		if user.Id != currentUser.Id{
			validUsersId = append(validUsersId, user.Id)
		}
	}
	var alreadyFollowing []models.Follower
	ctrl.DB.Model(&alreadyFollowing).Where(
		"follower_id = ?", currentUser.Id,
	).Where("following_id IN ?", validUsersId).Find(&alreadyFollowing)
	var alreadyFollowingUsersIds []uint
	for _, following := range alreadyFollowing {
		alreadyFollowingUsersIds = append(alreadyFollowingUsersIds, following.FollowingId)
	}
	var followersToCreate []models.Follower
	var createdFollowingUsersIds []uint
	for _, validUserId := range(validUsersId){
		if (!slices.Contains(alreadyFollowingUsersIds, validUserId)){
			followersToCreate = append(followersToCreate, models.Follower{FollowerId: currentUser.Id, FollowingId: validUserId})
			createdFollowingUsersIds = append(createdFollowingUsersIds, validUserId)
		}
	}
	if len(createdFollowingUsersIds) == 0 {
		return http.StatusBadRequest, createdFollowingUsersIds, errors.New("no valid users to follow")
	}
	result := ctrl.DB.CreateInBatches(followersToCreate, 100)
	if result.Error != nil {
		return http.StatusInternalServerError, createdFollowingUsersIds, result.Error
	}
	return http.StatusOK, createdFollowingUsersIds, nil
}

func (ctrl SocializeController) UnFollow(email string, userIds []uint) (int, []uint, error) {
	var currentUser models.User
	err := currentUser.GetUser(ctrl.DB, [][]string{{"email = ?", email}})
	if err != nil {
		return http.StatusNotFound, nil, err
	}
	var alreadyFollowing []models.Follower
	ctrl.DB.Model(&alreadyFollowing).Where(
		"follower_id = ?", currentUser.Id,
	).Where("following_id IN ?", userIds).Find(&alreadyFollowing)
	var alreadyFollowingUsersIds []uint
	for _, following := range alreadyFollowing {
		alreadyFollowingUsersIds = append(alreadyFollowingUsersIds, following.FollowingId)
	}
	ctrl.DB.Model(&alreadyFollowing).Delete(&alreadyFollowing)
	if len(alreadyFollowingUsersIds) == 0 {
		return http.StatusBadRequest, alreadyFollowingUsersIds, errors.New("no valid users to un follow")
	}
	return http.StatusOK, alreadyFollowingUsersIds, nil
}

func (ctrl SocializeController) GetFollowersList(email string) (int, []uint, error) {
	var followers []uint
	var currentUser models.User
	err := currentUser.GetUser(ctrl.DB, [][]string{{"email = ?", email}})
	if err != nil {
		return http.StatusNotFound, nil, err
	}
	ctrl.DB.Model(&models.Follower{}).Where("following_id = ?", currentUser.Id).Select("follower_id").Find(&followers)
	return http.StatusOK, followers, nil
}

func (ctrl SocializeController) GetFollowingsList(email string) (int, []uint, error) {
	var followings []uint
	var currentUser models.User
	err := currentUser.GetUser(ctrl.DB, [][]string{{"email = ?", email}})
	if err != nil {
		return http.StatusNotFound, nil, err
	}
	ctrl.DB.Model(&models.Follower{}).Where("follower_id = ?", currentUser.Id).Select("following_id").Find(&followings)
	return http.StatusOK, followings, nil
}
