package leetcode355

import (
	"sort"
)

type Twitter struct {
	users    map[int]*user
	serialID int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{users: make(map[int]*user), serialID: 0}
}

type user struct {
	userId    int
	tweets    []tweet
	followees map[int]*user
}

type tweet struct {
	tweetId  int
	seriesId int
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if user, ok := this.users[userId]; ok {
		user.tweets = append(user.tweets, tweet{
			tweetId:  tweetId,
			seriesId: this.serialID,
		})
		this.serialID++
		return
	}
	this.users[userId] = &user{
		userId: userId,
		tweets: []tweet{
			{
				tweetId:  tweetId,
				seriesId: this.serialID,
			},
		},
		followees: make(map[int]*user),
	}
	this.serialID++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	user, ok := this.users[userId]
	if !ok {
		return []int{}
	}
	newTweets := make([]tweet, 0, (len(user.followees)+1)*10)
	userTweets := user.tweets
	if len(userTweets) > 10 {
		userTweets = userTweets[len(userTweets)-10:]
	}

	for i := range userTweets {
		newTweets = append(newTweets, userTweets[i])
	}

	for _, followee := range user.followees {
		tweetIds := followee.tweets
		cnt := len(tweetIds)
		if cnt > 10 {
			tweetIds = tweetIds[len(tweetIds)-10:]
		}

		for _, tweet := range tweetIds {
			newTweets = append(newTweets, tweet)
		}
	}

	sort.Slice(newTweets, func(i, j int) bool {
		if newTweets[i].seriesId > (newTweets[j].seriesId) {
			return true
		}
		return false
	})
	if len(newTweets) > 10 {
		newTweets = newTweets[:10]
	}

	var res []int
	for _, tweetId := range newTweets {
		res = append(res, tweetId.tweetId)
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	follower, followerExist := this.users[followerId]
	followee, followeeExist := this.users[followeeId]
	if !followerExist {
		follower = &user{
			userId:    followerId,
			tweets:    []tweet{},
			followees: make(map[int]*user),
		}
		this.users[followerId] = follower
	}
	if !followeeExist {
		followee = &user{
			userId:    followeeId,
			tweets:    []tweet{},
			followees: make(map[int]*user),
		}
		this.users[followeeId] = followee
	}
	follower.followees[followee.userId] = followee
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	follower, followerExist := this.users[followerId]
	if followerExist {
		delete(follower.followees, followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
