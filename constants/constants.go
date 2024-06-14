package constants

type Tone struct {
	key   string
	emoji string
	value string
}

var FlirtingTones = [23]Tone{
	{key: "short", emoji: "💬", value: "Answer short & crisp"},
	{key: "long", emoji: "🗣️", value: "Answer long & detailed"},
	{key: "askq", emoji: "❓", value: "Ask a question"},
	{key: "date", emoji: "❤️", value: "Ask for date"},
	{key: "cocky", emoji: "😏", value: "Be cocky & funny"},
	{key: "tease", emoji: "😇", value: "Tease playfully"},
	{key: "intrigueq", emoji: "🤔", value: "Ask intriguing questions"},
	{key: "affection", emoji: "💞", value: "Show affection"},
	{key: "style", emoji: "🌸", value: "Compliment their style"},
	{key: "romantic", emoji: "🌹", value: "Be romantic"},
	{key: "roleplay", emoji: "🎭", value: "Use role-playing"},
	{key: "fungame", emoji: "🕹️", value: "Suggest a fun game"},
	{key: "laugh", emoji: "😁", value: "Make them laugh"},
	{key: "mysterious", emoji: "👻", value: "Be a bit mysterious"},
	{key: "future", emoji: "🕵️‍♀️", value: "Imagining future together"},
	{key: "sensual", emoji: "🍫", value: "Sensual Descriptions"},
	{key: "lovequote", emoji: "💌", value: "Send a love quote"},
	{key: "dedicatesong", emoji: "🎶", value: "Dedicate a song"},
	{key: "activities", emoji: "🌇", value: "Ask about favorite activities"},
	{key: "wishes", emoji: "🌠", value: "Ask about their wishes"},
	{key: "flirtcomment", emoji: "💫", value: "Make a flirtatious comment"},
	{key: "deeptopic", emoji: "🔍", value: "Dive deeper into a topic"},
	{key: "sweetdreams", emoji: "🌙", value: "Wish sweet dreams"},
}

var ProfessionalTones = [12]Tone{
	{key: "short", emoji: "💬", value: "Answer short & crisp"},
	{key: "long", emoji: "🗣️", value: "Answer long & detailed"},
	{key: "askq", emoji: "❓", value: "Ask a question"},
	{key: "agree", emoji: "👍", value: "Agree with their point"},
	{key: "disagree", emoji: "🚫", value: "Disagree respectfully"},
	{key: "clarify", emoji: "❓", value: "Ask for clarification"},
	{key: "changetopic", emoji: "🔄", value: "Change the topic"},
	{key: "opinion", emoji: "🗣️", value: "Express your opinion"},
	{key: "paraphrase", emoji: "💬", value: "Paraphrase their point"},
	{key: "lightenmood", emoji: "🎈", value: "Lighten the mood"},
	{key: "stayprofessional", emoji: "💼", value: "Stay professional"},
	{key: "debate", emoji: "💬", value: "Start a debate"},
}

var FriendlyTones = [22]Tone{
	{key: "short", emoji: "💬", value: "Answer short & crisp"},
	{key: "long", emoji: "🗣️", value: "Answer long & detailed"},
	{key: "askq", emoji: "❓", value: "Ask a question"},
	{key: "cocky", emoji: "😏", value: "Be cocky & funny"},
	{key: "tease", emoji: "😇", value: "Tease playfully"},
	{key: "intrigueq", emoji: "🤔", value: "Ask intriguing questions"},
	{key: "agree", emoji: "👍", value: "Agree with their point"},
	{key: "disagree", emoji: "🚫", value: "Disagree respectfully"},
	{key: "clarify", emoji: "❓", value: "Ask for clarification"},
	{key: "changetopic", emoji: "🔄", value: "Change the topic"},
	{key: "opinion", emoji: "🗣️", value: "Express your opinion"},
	{key: "paraphrase", emoji: "💬", value: "Paraphrase their point"},
	{key: "affection", emoji: "💞", value: "Show affection"},
	{key: "style", emoji: "🌸", value: "Compliment their style"},
	{key: "laugh", emoji: "😁", value: "Make them laugh"},
	{key: "mysterious", emoji: "👻", value: "Be a bit mysterious"},
	{key: "dedicatesong", emoji: "🎶", value: "Dedicate a song"},
	{key: "activities", emoji: "🌇", value: "Ask about favorite activities"},
	{key: "wishes", emoji: "🌠", value: "Ask about their wishes"},
	{key: "flirtcomment", emoji: "💫", value: "Make a flirtatious comment"},
	{key: "deeptopic", emoji: "🔍", value: "Dive deeper into a topic"},
	{key: "lightenmood", emoji: "🎈", value: "Lighten the mood"},
}
