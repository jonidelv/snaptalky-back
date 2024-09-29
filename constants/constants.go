package constants

type Tone struct {
	key   string
	emoji string
	value string
}

var FlirtingTones = [10]Tone{
	{key: "date", emoji: "❤️", value: "Ask for a date"},
	{key: "cocky", emoji: "😏", value: "Be cocky & funny"},
	{key: "tease", emoji: "😇", value: "Tease playfully"},
	{key: "intrigueq", emoji: "🤔", value: "Ask intriguing questions"},
	{key: "affection", emoji: "💞", value: "Show affection"},
	{key: "romantic", emoji: "🌹", value: "Be romantic"},
	{key: "pickupline", emoji: "👻", value: "Generate a pickup line"},
	{key: "flirtcomment", emoji: "💫", value: "Make a flirtatious comment"},
	{key: "nsfw", emoji: "😈", value: "NSFW"},
	{key: "sweetdreams", emoji: "🌙", value: "Wish sweet dreams"},
}

var FormalTones = [10]Tone{
	{key: "short", emoji: "💬", value: "Answer short & crisp"},
	{key: "long", emoji: "🗣️", value: "Answer long & detailed"},
	{key: "agree", emoji: "👍", value: "Agree with their point"},
	{key: "disagree", emoji: "🚫", value: "Disagree respectfully"},
	{key: "clarify", emoji: "❓", value: "Ask for clarification"},
	{key: "changetopic", emoji: "🔄", value: "Change the topic"},
	{key: "stayprofessional", emoji: "💼", value: "Stay professional"},
	{key: "deeptopic", emoji: "🔍", value: "Dive deeper into the topic"},
	{key: "providefeedback", emoji: "💡", value: "Provide constructive feedback"},
	{key: "empathize", emoji: "🤝", value: "Express empathy and understanding"},
}

var FriendlyTones = [10]Tone{
	{key: "askq", emoji: "❓", value: "Ask a question"},
	{key: "cocky", emoji: "😏", value: "Be cocky & funny"},
	{key: "tease", emoji: "😇", value: "Tease playfully"},
	{key: "intrigueq", emoji: "🤔", value: "Ask intriguing questions"},
	{key: "disagree", emoji: "🚫", value: "Disagree respectfully"},
	{key: "changetopic", emoji: "🔄", value: "Change the topic"},
	{key: "affection", emoji: "💞", value: "Show affection"},
	{key: "style", emoji: "🌸", value: "Compliment their style"},
	{key: "deeptopic", emoji: "🔍", value: "Dive deeper into a topic"},
	{key: "lightenmood", emoji: "🎈", value: "Lighten the mood"},
}

const TokenValidDays = 30
