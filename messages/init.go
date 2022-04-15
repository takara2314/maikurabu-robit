package messages

import "errors"

const (
	///////////////////////////////////
	//  Robit
	///////////////////////////////////

	Connectable          = "今はサーバーが開いているみたいだね。参加できるよ！"
	ForceRebootByProblem = "サーバーは今不調みたいだね。今から強制再起動するよ！"
	VotingNow            = "今サーバーを起動するかの投票をとっているよ！上の投稿を見てね！"
	StartVote            = "サーバーを起動してほしい人は、次の投稿に「いいね」をつけよう！"
	LaunchingNow         = "今サーバーを開ける処理をしているよ！ちょっと待ってね！"
	ForceRebootingNow    = "今強制再起動中だよ！ちょっと待ってね！"
	NotMeetMinVoter      = "十分な票数が集まらなかったね… みんながいる時間にもう一度やってみよう！"
	MeetMinVoter         = "%d人以上グッドを付けてくれたね！サーバーを開けるから少し待ってね！"
	FailedLaunchServer   = "サーバーを開くことが出来なかったよ… <@%s> に言ってね！"
	LaunchServer         = "サーバーの起動スイッチを押したよ！1分ほど待ってね！"
	FailedForceReboot    = "強制再起動が失敗しちゃった… <@%s> に言ってね！"
	ForceRebooted        = "強制再起動したよ！待たせてごめんね！"

	CheckStatusWait = "現在サーバー情報を取得しているよ！ちょっと待ってね！"

	///////////////////////////////////
	//  Robit Little Brother
	///////////////////////////////////

	CannotUseOutside      = "このコマンドは、マイクラ部以外のサーバーでは使用できません。"
	AlreadyHavePermission = "既にあなたは見る権限を持っています。"
	HadNotHavePermission  = "初めから見る権限を持っていません。"
	ShowWatchChannel      = "<#%s> でマイクラ鯖のチャットを見れるようになりました！チャンネルをミュートにしておくことを推奨します。"
	HideWatchChannel      = "マイクラ鯖のチャットを見れないようにしました。"
)

var (
	ErrTimeout = errors.New("timeout")
)
