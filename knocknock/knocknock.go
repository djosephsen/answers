package knocknock

import (
	"math/rand"
	"time"
	a "github.com/djosephsen/answers/lib"
)

var Answers = &a.Answer{
	Type : `knocknock`,
	Desc : `Knock Knock jokes1`,
	Get  : getAnswer,
	Rand : getRandAnswer,
     DB : db,
}

func getAnswer(index int) string{
	return db[index]
}

func getRandAnswer() string{
	now := time.Now()
	rand.Seed(int64(now.Unix()))
	return db[rand.Intn(len(db)-1)]
}

var db = []string{
"Canoe! \nCanoe who? \nCanoe come out and play with me today?",
"Who! \nWho who? \nThat’s what an owl says!",
"Lettuce.\n Lettuce who?\n Lettuce in, it’s cold out here.",
"Honey bee. \nHoney bee who? \nHoney bee a dear and get me some juice.",
"Wooden shoe. \nWooden shoe who? \nWooden shoe like to hear another joke?",
"A broken pencil. \nA broken pencil who. Oh never mind it’s pointless.",
"Cow says. \nCow says who?\n No silly, a cow says Mooooo!",
"Double. \nDouble who? \nW!",
"Mikey! \nMikey who? \nMikey doesn’t fit in the keyhole!",
"Atch. \nAtch who? \nBless you!",
"I am. \nI am who? \nYou don’t know who you are?",
"Ya. \nYa Who? \nWow, I’m excited to see you too.",
"Figs. \nFigs who? \nFigs the doorbell, it’s broken!",
"Boo! \nBoo who? \nDon’t cry, it’s just me.",
"The Interrupting pirate! \nThe Interrup… ARRRRRRRRRR!",
"The Interrupting cow! \nThe Interrup… MOOOOOOOOOOOOO!",
"Iva. \nIva who? \nI’ve a sore hand from knocking!",
"Avenue. \nAvenue who? \nAvenue knocked on this door before?",
"A little old lady. \nA little old lady who? I didn’t know you could yodel.",
"Banana. \nBanana who? \nKnock, knock. \nWho’s there? \nBanana. \nBanana who? \nKnock, knock. \nWho’s there? \nBanana. \nBanana who? \nKnock, knock. \nWho’s there? \nOrange. \nOrange who? \nOrange you glad I didn’t say banana?",
}
