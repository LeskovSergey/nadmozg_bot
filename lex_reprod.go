package main

import (
	"strings"
	//"log"
	"errors"
	"math/rand"
	"time"
	"log"
	"strconv"
)

var (
	mat    = []string{"ху", "пизд"}
	gopnik = []string{
		"Пацанчик, ты в натуре попутал, ты чо тут лечишь мне? Ты чо переобуваешься?",
		"Слышь, есть по сигарете?",
		"Опаньки. Да ты попутал… Собираешься по беспределу наехать?",
		"Нука быстро коня нарисовала!",
		"Сышь, благодарочка, ну ланенько я потяпал.",
		"Есть чё по мелочи?",
		"Ты вообще кто по жизни?",
		"Дай проездной почитать. Васа, о чем ты шепчешь?!А ты не спотыкаешься?",
		"Спрячь штакетину. Попадалово. ",
		"Чо ты мне исполняешь?! ",
		"А чё у тя штаны такие узкие, у сестрёнки забрал? ",
		"А может на кортах поговорим? ",
		"Братишка, есть чо помелочи то? Найди что ли, попрыгай там… ",
		"Да чё ты мажешься, чё не пацан…}",
		"Телефон есть? А если найду?",
		"Слышь ты, ты ваще не прав, ты осознай это.",
		"Слышь, брателло, ты какой масти?",
		"Слышь ты, ты обоснуй свои слова, чтобы потом о здоровье не пожалеть.",
		"Жиги не найдётся?",
		"А чо ты мне лечишь? Чо ты беса гонишь?",
		"Дай куртку, пацану холодно. Слышь?! а чё ты такой умный? Школу закончил? Эй ты, чорт! А чо ты такой волосатый как девка?",
		"Опа, а ты чё, предъявить мне хочешь? Ну так ты обзовись сначала.",
		"Я че на фраера похож?",
		"Чем меньше у человека зубов, тем он лучше фильтрует базар.",
		"Слышь пацанчик, а есть чёнить? а? А найду? ",
		"Я чёткий пацан, я по жизни в теме!!! Понял!",
		"А чё проблемы? Нет? Ща будут… ",
		"Пацан сказал! — Пацан забыл!",
		"Опа. Ты чо, предъявить мне хочешь, брателло?",
		"Дай куртку, пацану холодно.",
		"Опаньки. Да ты попутал, чепушила.",
		"Васа, о чем ты шепчешь?!",
		"Стрела в восемь? Я приду по—любому.",
		"А дай деньги подержать!",
		"Ты че, рамсы попутал?",
		"Ты чё из Продижи ? Я думал вы распались!",
		"Опа! Опаньки! Опа, нихуя себе вштырило!",
		"Ты, да ты сам себя под лоха подписал.",
		"Ты, да ты чё такой охуевший?",
		"Есть 25 копочек (копеек), до дома доехать?",
		"А че есть есть пару рублей для пацанов?",
		"А чё у тя штаны такие узкие, у сестрёнки спиздил?",
		"Эээй...браток, подогрей сигаретой.",
		"Слишком часто улыбаться нельзя — зубы могут окислиться.",
		"Слышь бля!",
		"Мужик не тот, кто яйца брил, а тот кто в армии служил.",
		"А чо ты мне лечишь? Чо ты беса гонишь?",
		"Че ты палишь? Спрячь штакетину. Попадалово.",
		"Слышь, ну ты в натуре определись: пацан ты или лох.",
		"Ты за базар отвечаешь? Ты кого знаешь?",
		"Слышь ты, ты ваще не прав, ты осознай это.",
		"Есть чо?",
		"Пацанчик, ты в натуре попутал за всю хуйню! Ты чо тут лечишь мне? Ты чо переобуваешься?",
		"Э, братуха. Дай чисто сотку позвонить не в падлу, у меня симка своя.",
		"Ты, да ты чертила мутный, иди сюда, сука!",
		"Мы сегодня раскумаримся по—любому.",
		"Слышь, братишка. Тормозни—ка, базар есть. А чо ты рамсишь—то? Ты чо такой голимый?",
		"Бля, прикинь — только штакет взорвали, и мусора. Я даже на измену подсел.",
		"Ты чё, страх потерял, шерсть?",
		"Дела ништяк. Все пучком.",
		"Собираешься по беспределу наехать?",
		"Что, западло с нормальными пацанами побазарить?",
		"Ты че такой борзый?",
		"А может на кортах поп*здим?",
		"Ты че, не уважаешь реальных пацанов?",
		"Давай с тобой ЧЕТКО познакомимся!",
		"Пацик дай берцы потаскать!",
		"Братишка, есть чо помелочи то? Найди что ли, попрыгай там...",
		"Сыш братуха, семки есть?",
		"Монету дай посмотреть!",
		"А че, есть че?",
		"Ты чё, не пацан?",
		"Слышь, есть по сигарете?",
		"Будь пацаном, вывози за базар!",
		"Ты чо распустил тут свои выебоны?",
		"Ты чо, не вкурил куда ты зашел?",
		"Эй пацан! Ты с какого района?!",
		"Парень, ты не прав, не по понятиям.",
		"Чем меньше у человека зубов, тем он лучше фильтрует базар.",
		"Сюда иди бля!",
		"Хуля пиздиш?",
		"Слышь,ты ахуел?",
		"Ты чо,конченный??",
		"Присядь, поговорим.",
		"Я тебе ноги поломаю!",
		"Э-э слышь..",
		"Ты чёта попутал вроде...",
		"Чё ты меня лечишь..",
		"ты че, не уважаешь нас?",
		"ты не прав по понятиям",
		"я на зоне сидел",
		"ты знаешь кто я такой?",
		"По хую мороз",
		"У меня нет перстня-поэтому целуйте меня в жопу.",
		"А кто это к нам колеса катит?",
		"Отскочем побормочем",
		"Ты чё берега попутал",
		"Ебалом хлопни, или я тебя хлопну.",
		"Слышь ты местный?",
		"Че-то ты не по сезону шелестишь",
		"Присядь на уровень, поговорим.",
		"иди катай матрас",
		"Эээ а чё у тебя руки в карманах?",
		"Эй, сука, иди сюда!",
		"Ты чо бля олень охуел?",
		"Дай рубалек.",
		"Дахулиты ведешся?",
		"Чоты ебать?в ебло дать?",
		"Хули ты тупишь, шмель ебаный",
		"У тя походу полтаха храпит?",
		"Ты чоооооо бля барбос?",
		"Тупишь быдла бля? Да ты походу тёплый?",
		"Слышь, а ты по каким понятиям ваще живёшь?",
		"Жопу на весь район ставишь?",
		"Хули ты вякаешь?!",
		"Ебало завали, рысявка!",
	}
)

func huifma(word string) (string, error) {

	v_index := vowel_index(&word)
	if v_index != -1 {
		utf8word := []rune(strings.ToLower(word))
		part2 := utf8word[vowel_index(&word) : len(word)-1]
		switch part2[0] {
		case 'а':
			part2[0] = 'я'
		case 'о':
			part2[0] = 'ё'
		case 'э':
			part2[0] = 'е'
		case 'ы':
			part2[0] = 'и'
		case 'у':
			part2[0] = 'ю'
		}
		sr := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(sr)
		return mat[r1.Intn(len(mat))] + string(part2), nil
	}
	return strings.ToLower(word), errors.New("vowels not found")
}

func gopnic_answer() (string){
	sr := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(sr)
	return gopnik[r1.Intn(len(gopnik))]
}

func vowel_index(word *string) int {
	var i int
	for _, value := range *word {
		switch value {
		case 'а', 'о', 'и', 'е', 'ё', 'э', 'ы', 'у', 'ю', 'я', 'А', 'О', 'И', 'Е', 'Ё', 'Э', 'Ы', 'У', 'Ю', 'Я':
			return i
		}
		i++
	}
	return -1
}

func probably(per int) (bool) {
	if (per < 100) && (per > 0) {
		sr := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(sr)
		cur_n := r1.Intn(100)
		log.Println("rand "+ strconv.Itoa(cur_n) + " of " + strconv.Itoa(per))
		if cur_n < per {
			return true
		}
	}
	return false
}