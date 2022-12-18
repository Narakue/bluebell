package redis

const (
	KeyPre           = "bluebell."
	KeyPostVoteTime  = "post.vote.time"
	KeyPostVoteScore = "post.vote.score"
	KeyPostVoteUser  = "post.vote.user."
)

func GetKey(key string) string {
	return KeyPre + key
}
