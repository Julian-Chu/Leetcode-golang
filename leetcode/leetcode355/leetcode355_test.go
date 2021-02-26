package leetcode355

import (
	"reflect"
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	twitter := Constructor()

	// User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 5)

	// User 1's news feed should return a list with 1 tweet id -> [5].
	res := twitter.GetNewsFeed(1)
	if !reflect.DeepEqual(res, []int{5}) {
		t.Errorf("want {5}")
	}

	// User 1 follows user 2.
	twitter.Follow(1, 2)

	// User 2 posts a new tweet (id = 6).
	twitter.PostTweet(2, 6)

	// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
	// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	res = twitter.GetNewsFeed(1)
	if !reflect.DeepEqual(res, []int{6, 5}) {
		t.Errorf("want {6,5}")
	}
	// User 1 unfollows user 2.
	twitter.Unfollow(1, 2)

	// User 1's news feed should return a list with 1 tweet id -> [5],
	// since user 1 is no longer following user 2.
	res = twitter.GetNewsFeed(1)
	if !reflect.DeepEqual(res, []int{5}) {
		t.Errorf("want {5}")
	}
}

func TestCase2(t *testing.T) {
	twitter := Constructor()

	// User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 5)
	twitter.PostTweet(1, 3)

	// User 1's news feed should return a list with 1 tweet id -> [5].
	res := twitter.GetNewsFeed(1)
	if !reflect.DeepEqual(res, []int{3, 5}) {
		t.Errorf("want {3,5}")
	}
}
func TestCase3(t *testing.T) {
	twitter := Constructor()

	// User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 5)
	twitter.Follow(1, 1)

	// User 1's news feed should return a list with 1 tweet id -> [5].
	got := twitter.GetNewsFeed(1)
	want := []int{5}
	if !reflect.DeepEqual(got, want) {
		gotStr := "{"
		for i := 0; i < len(got); i++ {
			gotStr += strconv.Itoa(got[i])
			if i >= len(got)-1 {
				break
			}
			gotStr += ","
		}
		gotStr += "}"
		wantStr := "{"
		for i := 0; i < len(want); i++ {
			wantStr += strconv.Itoa(want[i])
			if i >= len(want)-1 {
				break
			}
			wantStr += ","
		}
		wantStr += "}"
		t.Errorf("want %s, but got %s", wantStr, gotStr)
	}
}
