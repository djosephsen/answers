package chicken

import (
	"math/rand"
	"time"
	a "github.com/djosephsen/answers/lib"
)

var Answers = &a.Answer{
	Type : `chicken`,
	Desc : `Answers to 'Why did the chicken cross the road?'`,
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
	`Because the armadillo told him it was safe`,
	`Because it was being converted to Christianity`,
	`To avoid Colonel Saunders`,
	`because it was stapled to the monkey`,
	`because it was stuck on the pervert`,
	`To get the Chinese newspaper.  Get it? Neither do I! (I get The Washington Post!)`,
	`To escape the chicken jokes`,
	`No-one knows, but the road sure was pissed`,
	`Because MURICA! It can go anywhere it wants`,
	`Because it was a fundie chicken`,
	`Because the road crossed him`,
	`In Soviet Russia, the road crosses YOU!`,
	`To die. In the rain. Alone.
	-Hemingway`,
	`To boldly go where no chicken has gone before!
	-Kirk`,
	`Dammit Jim, I'm a doctor, not a barnyard psychologist!
	-Mcoy`,
	`The chicken's medium told her to "cross over"`,
	`The coyote had put a sign over there, saying "Bird Seed"`,
	`Ray Charles told her to hit the road`,
	`To lay it on the line`,
	`Because the light was green`,
	`Chicken crosses road
	Farside joke in the making
	Avoids falling sky`,
	`Because the grass was greener over there`,
	`Because it couldn't cross its fingers`,
	`Eszterhazy had long made a habit of keeping several enquiries in progress, in some manner or other, simultaneously, hopeful to inoculate himself against the sense of ennui and listlessness that often ensued upon the successful conclusion of an enquiry. And on this morning, while attempting to concentrate his mind on the far more pressing (indeed, to be sure a matter of national security and international tranquility) matter of the theft of the Cyprus Regalia from the Crypt of St. Sophie, he found his attention inexplicably but inexorably diverted to the mystery of the Chicken Who Crossed the Road. It would appear to the casual observer that the worms and corn were as abundant upon the Hither Side of the road as upon the Thither Side, the gravel bits as bright and appealing, the hens as plump and complaisant. Yet Eszterhazy, he more than many others, could readily empathize with the creeping restlessness that could make the near and comfortable side of the road appear stale through familiarity, the unknown far side an inviting field of discovery and possibilities. Just so, but it would be an error of sophomoric dimensions to assume, without more evidence, that a course of action appealing to Eszterhazy might be similarly appealing -- or appealing for similar reasons -- to a Gallus gallus. As he selected from his humidor, clipped, lit, drew, and meditatively puffed upon a Trichonopoly cheroot, Eszterhazy ruminated (ruminated? aviated? gallicated? brooded? nay, not brooded) upon the words of the so-called Baconian Addendum to the Malleus Maleficarum: "The mind of a chicken is not the same as the mind of a man." And that, indeed, might well be the answer. But, Eszterhazy wondered, to which question?`,
	`42!`,
	`To make you curious as to why a chicken might cross a road`,
	`For fowl purposes.`,
	`I don't know about the chicken, but I do know why the dinosaur crossed the road: Because the chicken had not evolved yet.`,
	`I could tell you, but then the Chicken Mafia would kill me.`,
	`Because shut the hell up, that's why.`,
	`Chicken Boo, what's the matter with you?`,
	`You don't act like the other chickens do.`,
	`You wear a disguise to look like human guys.`,
	`But you're not a man, you're a chicken, Boo.`,
	`Give me ten minutes with that chicken and we'll find out.
	-- Torquemada`,
	`Because the chicken was an engineer and that's what Stackoverflow said to do.`,
	`I dream of a better world...where chickens can cross the road without having their motives questioned`,
}
