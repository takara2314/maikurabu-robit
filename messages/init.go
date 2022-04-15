package messages

import "errors"

const (
	///////////////////////////////////
	//  Robit
	///////////////////////////////////

	CheckStatusWait = "現在サーバー情報を取得しているよ！ちょっと待ってね！"

	FailedStartByLocking = "ごめんね。操作ロックがかけられているよ！もう少し待ったら解除されるかも…！"
	Connectable          = "今はサーバーが開いているみたいだね。参加できるよ！"
	ForceRebootByProblem = "サーバーは今不調みたいだね。今から強制再起動するよ！"
	VotingNow            = "今サーバーを起動するかの投票をとっているよ！上の投稿を見てね！"
	StartVote            = "サーバーを起動してほしい人は、次の投稿に「いいね」をつけよう！"
	LaunchingNow         = "今サーバーを開ける処理をしているよ！ちょっと待ってね！"
	ForceRebootingNow    = "今強制再起動中だよ！ちょっと待ってね！"
	VoteMessageNotFound  = "こらー！！！投票メッセージを消すなー！！怒るよ？"
	NotMeetMinVoter      = "十分な票数が集まらなかったね… みんながいる時間にもう一度やってみよう！"
	MeetMinVoter         = "%d人以上グッドを付けてくれたね！サーバーを開けるから少し待ってね！"
	FailedLaunchServer   = "サーバーを開くことが出来なかったよ… <@%s> に言ってね！"
	LaunchServer         = "サーバーの起動スイッチを押したよ！1分ほど待ってね！"
	FailedForceReboot    = "強制再起動が失敗しちゃった… <@%s> に言ってね！"
	ForceRebooted        = "強制再起動したよ！待たせてごめんね！"

	///////////////////////////////////
	//  Robit Little Brother
	///////////////////////////////////

	CannotUseOutside      = "このコマンドは、マイクラ部以外のサーバーでは使用できません。"
	AlreadyHavePermission = "既にあなたは見る権限を持っています。"
	HadNotHavePermission  = "初めから見る権限を持っていません。"
	ShowWatchChannel      = "<#%s> でマイクラ鯖のチャットを見れるようになりました！チャンネルをミュートにしておくことを推奨します。"
	HideWatchChannel      = "マイクラ鯖のチャットを見れないようにしました。"

	///////////////////////////////////
	//  Activity
	///////////////////////////////////

	Closing        = "サーバーは閉まっています :("
	Timeouted      = "タイムアウトしました :("
	HowManyPlayers = "%d人がログイン中だよ :)"
	NoPlayers      = "誰もログインしていないよ :|"
	LowPing        = "%dms 遅延 - 快適です :)"
	HighPing       = "%dms 遅延 - ラグいです :("

	///////////////////////////////////
	//  Auto stopping
	///////////////////////////////////

	AutoStopping = "%d分間サーバーに誰もいないので、サーバーを終了するよ。"

	///////////////////////////////////
	//  Start command locking
	///////////////////////////////////

	StartCMDLocked    = "操作ロックしたよ！これで集中して作業できるね！"
	StartCMDUnlocked  = "操作ロック解除したよ！これでみんなと接しれる！"
	StartCMDLocking   = "現在操作ロックされているよ！安全だね！"
	StartCMDUnlocking = "現在操作ロックが解除されているよ… 作業するならロックしておこう！"
)

var (
	ErrTimeout  = errors.New("timeout")
	ErrNotFound = errors.New("not found")
)
