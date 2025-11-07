package goroutines

import (
	"fmt"
	"time"
)

var word string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed nec diam leo. Mauris sed neque quam. Phasellus sit amet venenatis arcu. Nulla malesuada lorem sollicitudin sem auctor facilisis. Integer nec eros ut turpis gravida suscipit nec non magna. Proin non est accumsan justo finibus tincidunt. Nam varius sit amet est quis ullamcorper. Integer lacinia bibendum tincidunt. Pellentesque ultrices, ligula tincidunt sagittis tincidunt, mi nisl dignissim arcu, sit amet viverra nisi erat et odio.
Etiam ac ex et lectus semper aliquam a eu dolor. Ut non eros imperdiet nisi pretium maximus. Proin malesuada, eros non scelerisque vestibulum, libero urna feugiat nulla, interdum accumsan augue augue quis arcu. Duis aliquet elementum velit a pretium. Vestibulum ac sapien in sem ornare cursus nec et neque. Ut sed massa nibh. Morbi eros neque, ultricies eu sollicitudin id, consectetur sed est. Mauris tincidunt facilisis feugiat. Nunc ultrices, diam vel sodales accumsan, arcu dui scelerisque neque, quis vehicula ante metus id tortor. Morbi sed mi id sapien bibendum convallis. Sed ultricies leo metus, id placerat justo pretium a. Cras consectetur mi nunc, quis aliquet nunc imperdiet eget. Curabitur ut tortor efficitur, fringilla neque congue, rutrum justo. Sed ut quam et orci sagittis luctus. Aliquam erat volutpat. Integer mattis enim et rhoncus mattis.
Sed mattis nisl massa, ac hendrerit urna pulvinar et. Aliquam tempor, augue vel congue vestibulum, enim nunc posuere augue, convallis pulvinar libero odio at sapien. Mauris faucibus rhoncus pellentesque. Nam fermentum odio dolor, sed scelerisque felis egestas nec. Fusce consequat dolor ligula, nec posuere ipsum accumsan quis. Curabitur egestas tellus id neque mollis, nec pharetra ipsum egestas. Sed tincidunt arcu metus, ut mattis ligula commodo malesuada. Aenean interdum erat aliquet nulla vestibulum, et mattis lorem fringilla.
In ultrices nisi vitae vulputate pretium. Mauris justo sem, gravida a vehicula ac, convallis non urna. Proin turpis urna, maximus varius lacinia eu, dignissim nec sapien. Integer pulvinar enim in commodo vehicula. Donec dolor lacus, tristique et lobortis nec, vulputate non enim. Sed sollicitudin sed massa ac mattis. Pellentesque placerat neque neque, eget dignissim augue egestas ut. Morbi at purus vitae eros vehicula viverra. Suspendisse potenti. Aliquam vel nisi blandit, accumsan tellus sit amet, faucibus lorem.
Maecenas sit amet tristique orci, ac hendrerit quam. Phasellus id arcu ac lectus elementum aliquam dignissim eget velit. Curabitur ante purus, lacinia quis fringilla nec, luctus a dolor. Aliquam ultricies arcu non tincidunt sagittis. Donec ac varius est. Nulla sed tincidunt ante. Nunc aliquet ligula id viverra viverra. Sed ut dui at massa tristique suscipit non tristique est. Ut et dolor pulvinar, condimentum neque et, finibus eros. Vivamus ut arcu eros. Donec molestie mi commodo, varius felis non, vulputate elit. Vestibulum nec lorem eu erat tempor dapibus non et felis. Pellentesque ut sem mauris.
Donec lobortis, turpis et vestibulum commodo, turpis est congue diam, ut consequat lorem libero nec metus. Praesent cursus lacinia quam, a molestie nisl ullamcorper sodales. Aenean pharetra nisl lectus, ac placerat nisl dignissim quis. Pellentesque finibus a turpis vitae mattis. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam dignissim sodales nulla quis rhoncus. Fusce venenatis a elit non fermentum. Aliquam sed erat vel mauris cursus tincidunt.
Nulla a erat ante. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris vel posuere ipsum. Nulla ullamcorper arcu nulla, in vehicula justo venenatis non. Suspendisse efficitur ut sapien ac hendrerit. Quisque vel enim libero. Nullam vel arcu suscipit, lobortis quam et, ultricies felis. Pellentesque consequat sodales condimentum. Proin vitae turpis ac diam lacinia placerat.
In hac habitasse platea dictumst. In auctor nulla at eleifend fermentum. Interdum et malesuada fames ac ante ipsum primis in faucibus. Ut dignissim scelerisque lectus vitae porttitor. Proin nisl nisi, lacinia nec fringilla a, tristique nec erat. Vestibulum consectetur tellus a enim mattis, eu congue ex tristique. Mauris sed enim a metus aliquet molestie. Aenean ex diam, viverra nec tortor eu, egestas vulputate arcu. In venenatis tempor libero, vel tempor felis euismod quis. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Quisque semper felis vehicula, facilisis erat vel, tempus urna.
Sed fermentum nibh a consequat elementum. Nullam ultrices nisl augue, quis commodo libero vestibulum sit amet. Morbi accumsan ut dui ac consectetur. Vivamus dignissim aliquam sem, id suscipit eros. In ut accumsan orci. Duis elit mauris, suscipit tristique ex quis, vulputate mollis massa. Vestibulum dapibus nibh sed velit tempus scelerisque. Nullam volutpat eros at risus interdum hendrerit. Curabitur quis mi venenatis, volutpat velit eu, euismod mauris.
Aenean at augue non dolor aliquet blandit. Donec tincidunt feugiat nisl et condimentum. Donec non tellus nunc. Phasellus vel bibendum turpis, ac bibendum massa. Nulla vulputate eget lectus ut porttitor. Curabitur ut nisl eu tellus ornare porta. Curabitur a efficitur purus.
Donec eget odio diam. Pellentesque et sapien pharetra, condimentum eros vel, dignissim augue. Cras et tincidunt lacus. Proin viverra et justo quis placerat. Aenean placerat ipsum vel consequat mollis. Integer rutrum libero non odio lobortis pulvinar a ac tellus. Mauris nec porttitor dui, quis ornare massa. Quisque eu elementum nulla. Proin interdum ullamcorper est. Pellentesque quis odio quis mi venenatis vestibulum. Donec iaculis semper diam, sodales eleifend sapien consequat at.
Cras at tellus augue. Donec in eros tempus, malesuada ex vel, aliquam magna. Nunc tincidunt pellentesque lorem et interdum. Morbi quis venenatis sapien, quis vestibulum tortor. Nullam dignissim lobortis elit id blandit. Vivamus ultrices enim id dolor pulvinar, ut dignissim dui commodo. Cras in tellus eu leo feugiat sollicitudin. Phasellus mattis erat eu dolor molestie, sit amet ullamcorper purus maximus. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent eget fermentum tellus, non vestibulum tellus. Aenean dapibus, massa eget scelerisque suscipit, erat urna pretium orci, ut aliquam risus lectus iaculis nibh. Vivamus cursus sapien odio, eu suscipit erat bibendum et.
Nam ut purus magna. Quisque sit amet purus neque. Nulla ut dui id est ultrices iaculis. Nulla suscipit sagittis euismod. Quisque nunc lorem, scelerisque eu tempus et, blandit eu tortor. Phasellus scelerisque congue nulla non malesuada. Fusce porttitor dolor vel nisl egestas, a bibendum augue semper. Nullam tincidunt, tortor et interdum venenatis, erat metus mollis diam, vitae finibus sem nunc et eros. Sed hendrerit neque finibus, porta dui in, cursus tellus. Nulla dignissim maximus sem vitae congue. Sed ac convallis lacus, eget placerat ipsum. Duis vulputate sapien massa, sit amet interdum metus ultricies a. Morbi id arcu blandit, porttitor metus a, dictum lorem. Praesent imperdiet erat magna, aliquam facilisis nulla congue et. Phasellus quis elit semper, congue velit vitae, laoreet sem.
Vivamus aliquet aliquam tincidunt. Duis venenatis purus felis, ut pharetra urna suscipit sit amet. Aliquam a mauris vitae odio auctor facilisis. Vivamus purus nisi, ultricies sit amet neque a, malesuada aliquet orci. Mauris gravida, magna et euismod auctor, neque est pretium nulla, in pharetra arcu nibh eu est. Nullam dictum efficitur lobortis. Duis pellentesque purus nisi, sit amet molestie dui pellentesque a. Curabitur iaculis tempus quam. Donec volutpat feugiat velit, vitae consectetur nisi auctor dapibus. Donec imperdiet enim massa, laoreet ullamcorper quam suscipit nec. Vivamus quis ex placerat, pharetra sapien sed, ullamcorper arcu. Vivamus interdum magna laoreet condimentum consectetur. Vestibulum vitae nunc leo. Etiam ornare ac ex iaculis commodo.
In auctor maximus tempus. Vestibulum aliquet, ex quis eleifend commodo, eros est mattis lectus, at ullamcorper purus elit vel risus. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Aliquam mauris mauris, molestie sit amet hendrerit ac, laoreet quis lorem. Quisque non metus in massa rutrum vulputate. Aliquam nisl dui, luctus sed erat eget, bibendum malesuada tellus. In vestibulum luctus laoreet. Quisque vel lacinia ex. Nunc consectetur sed mauris ut ultrices. Curabitur commodo sollicitudin libero in maximus.
In in sollicitudin elit. Donec in aliquet eros. Sed sit amet ipsum sed erat mattis consectetur in consectetur dui. Maecenas dignissim ac orci ut aliquam. Cras sed nibh lectus. Maecenas auctor molestie congue. Aliquam ac turpis tempus, pharetra arcu ut, varius metus. Maecenas finibus finibus condimentum. Suspendisse libero sapien, commodo a bibendum ac, ornare non ligula. Fusce at sagittis dui. Pellentesque vitae dolor vitae enim laoreet condimentum id vel massa. Etiam malesuada ullamcorper ipsum, a lobortis neque cursus congue. Nulla porta lacus a porttitor pellentesque. Nam pulvinar vehicula elit, et faucibus metus lobortis eget. Vestibulum eu metus a purus cursus venenatis. Sed maximus lacus commodo urna mattis accumsan.
Donec auctor mi sed pretium facilisis. Praesent orci arcu, sollicitudin ut maximus ac, luctus id lorem. Aliquam nunc ipsum, laoreet iaculis facilisis in, tempus non justo. Maecenas et diam libero. Cras felis enim, elementum eu gravida sed, porta eu enim. Vestibulum facilisis condimentum metus, commodo sollicitudin velit tincidunt ut. Mauris quis augue non eros ullamcorper viverra. Vivamus nisi nisi, tincidunt quis tortor eget, porta congue enim.
Nam pulvinar lorem vulputate sapien laoreet ullamcorper. Nulla vel euismod sem. Ut sed ligula erat. Sed consequat sapien leo, et accumsan mauris finibus vitae. Sed hendrerit finibus dolor, id gravida massa suscipit nec. Praesent turpis elit, blandit id tellus eu, bibendum congue turpis. Nullam eleifend ut tortor sed consequat. Ut dolor leo, pretium ut dui quis, congue aliquam diam. Nunc pharetra auctor bibendum. Phasellus arcu neque, tempor sit amet nulla sit amet, bibendum fringilla nibh. Aenean dapibus non erat eu molestie.
Pellentesque ut lorem vel dolor aliquet luctus vitae sagittis ligula. Quisque semper lectus libero, finibus dapibus augue pretium ut. Aliquam in congue libero. Maecenas volutpat, libero quis porttitor cursus, lorem nunc malesuada risus, et efficitur tortor odio eget ipsum. Suspendisse vel ligula vel eros viverra tristique et quis ex. Sed tortor quam, ultrices sed volutpat vel, porttitor ut ipsum. Suspendisse semper ligula eros, ac consectetur ante rutrum in. Nulla nisi dui, condimentum sed nibh et, vulputate ullamcorper risus. Aliquam eu est nec turpis iaculis lacinia. Sed eget sapien laoreet, consectetur felis quis, consectetur lacus.
Cras sodales ut odio fringilla imperdiet. Cras auctor urna metus, in mollis leo mattis ac. Nullam eu dui interdum, euismod neque in, dignissim augue. Donec vehicula velit a ullamcorper dictum. Mauris consectetur eros purus, quis eleifend ligula pretium non. Ut neque dolor, interdum id nunc id, consequat bibendum ex. Maecenas velit nulla, dapibus vitae odio in, finibus maximus lorem. Sed in lacus eleifend, semper sapien eget, ultrices augue.
Integer nec congue metus, vel pretium augue. Nam sed quam vestibulum, sollicitudin felis a, tempor arcu. Nulla convallis imperdiet fermentum. Nulla euismod, tellus eget cursus faucibus, urna tellus sodales elit, ut pulvinar velit magna ut metus. Nulla lobortis enim ut pretium aliquam. Sed sit amet dolor sit amet orci placerat venenatis. Maecenas rhoncus elementum diam in dapibus.
Aenean sed finibus felis. Fusce leo mauris, cursus ut tincidunt non, bibendum id leo. Vivamus eget sem ante. Sed pellentesque mi risus, sit amet lacinia eros ullamcorper at. Donec at nunc semper erat cursus egestas. Pellentesque est erat, varius a euismod vitae, ullamcorper quis odio. Vivamus pretium molestie augue et feugiat. Etiam sollicitudin auctor elit id rhoncus. Sed non dignissim nisi. Donec maximus vel sapien ac hendrerit.
Etiam feugiat odio laoreet pharetra eleifend. Nulla ipsum risus, blandit non interdum non, tempor id erat. Maecenas non dui a risus lacinia rutrum id ac sapien. Vestibulum ultrices feugiat libero, in gravida orci aliquam ultrices. Cras sit amet vulputate mauris. Phasellus viverra fermentum nibh, eget vulputate arcu commodo in. Fusce condimentum, arcu ac iaculis posuere, libero massa fringilla odio, nec blandit ipsum est non magna. Mauris leo odio, fermentum eu aliquet at, blandit eget purus. Nulla facilisi. Nullam et lorem odio.
Morbi id pellentesque ligula. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nunc bibendum nisl mauris, id malesuada libero tempor finibus. Aliquam quis porttitor nunc. Donec vitae viverra sem. In risus ligula, tincidunt dapibus egestas vel, lacinia non nunc. Nulla a elit ultrices, cursus urna non, fringilla eros. Donec luctus malesuada mi, eu tincidunt neque sodales ac. Proin interdum vestibulum lorem at vulputate. Sed justo lectus, maximus eget tempus vel, ullamcorper vel nisi. In molestie turpis eget ullamcorper rhoncus.
In gravida facilisis commodo. Nulla facilisi. Fusce auctor semper eleifend. Aenean est nulla, tincidunt vitae ligula sit amet, cursus mollis augue. Aliquam purus purus, posuere quis sodales eu, tincidunt id urna. Maecenas ultricies sit amet orci a gravida. Aliquam ut bibendum ante. Quisque vel imperdiet velit. Suspendisse potenti. Sed finibus dapibus tellus, eget mattis libero.
Proin mollis nisl ligula, in consectetur velit feugiat nec. Nullam ut elit at ipsum ultrices vestibulum vitae sit amet enim. Nunc venenatis eros ac risus gravida, a blandit urna eleifend. Donec tincidunt lectus ipsum, vitae semper magna ultrices eget. Aenean dictum viverra feugiat. Maecenas ac massa vulputate, tincidunt eros in, semper velit. Fusce vehicula vestibulum varius. Nullam nec velit eget mi lacinia accumsan. Sed consectetur ex sit amet nibh fringilla tincidunt. Proin ac laoreet nunc. Proin vitae sapien maximus, convallis purus maximus, pharetra sapien. Sed erat diam, rutrum quis porttitor ac, eleifend non turpis. Curabitur maximus vehicula orci, a maximus tortor pretium at. Etiam porttitor placerat nisi et laoreet. Aenean id orci quis dui ultricies tincidunt quis sed risus. Mauris pulvinar facilisis massa, ac pharetra libero ornare vitae.
Praesent mi lectus, accumsan vitae luctus nec, commodo sed leo. Quisque a dui ligula. Curabitur pulvinar ipsum at maximus sagittis. Nullam mauris libero, molestie ut sagittis at, fermentum a lectus. Cras sit amet tempor nunc. Quisque tempor metus est, non placerat libero sollicitudin sed. Sed ut vestibulum risus. Praesent eu suscipit sapien, vitae malesuada mi. Etiam augue ex, tristique a tempus id, consequat et elit. Pellentesque tempus sapien a eros efficitur tristique. Interdum et malesuada fames ac ante ipsum primis in faucibus. In ornare congue dapibus. Vestibulum scelerisque urna ut sollicitudin vulputate. Sed porttitor enim eu leo ornare, sit amet auctor erat accumsan. Cras nec ipsum a arcu consequat tempor. Nam at tellus ultricies, pretium ex auctor, venenatis nunc.
Nullam vehicula tellus nec porttitor porttitor. In commodo, urna et blandit ultrices, odio nisi blandit elit, eu ullamcorper odio risus vel elit. Aliquam tempor, nisi in sagittis feugiat, massa lorem congue nisi, sit amet scelerisque eros turpis sed libero. Proin sed condimentum dolor. Aliquam est justo, volutpat sit amet augue sit amet, consequat suscipit ligula. Quisque volutpat vitae tortor et convallis. Etiam dictum commodo augue.
Sed mollis nunc ornare sapien rutrum egestas. Integer et rutrum augue. Curabitur magna elit, imperdiet eget pretium vitae, rhoncus sit amet elit. Praesent ac volutpat dolor. Suspendisse ac diam eu tellus dignissim egestas vel lacinia urna. Vivamus semper facilisis tortor, non convallis risus placerat in. Proin fringilla eu magna at finibus. Etiam nec porta risus. Aenean finibus dolor leo, et aliquam eros malesuada vitae. Vestibulum faucibus interdum nisi sit amet vehicula.
Vestibulum gravida urna rutrum leo blandit feugiat. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. In mollis enim in vehicula sagittis. Donec rutrum lacus vitae ligula viverra, quis luctus nibh semper. Vestibulum et blandit lectus, quis gravida nulla. Fusce in ultricies enim. Donec mattis odio augue. Proin ac nibh odio. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed condimentum augue eu est consequat, a posuere dui cursus. Etiam consectetur laoreet quam, vitae tristique augue. Proin sit amet lacus porttitor, cursus libero ut, viverra metus. Integer efficitur iaculis diam eget imperdiet.`

func Start() {
	//go myFunc()

	// Самовызывающаяся анонимная горутина
	/*
		go func() {
			fmt.Println("Privet, ya anonymous")
		}()
	*/

	/*
		c := make(chan int, 1) // Канал с буфером размером 1
		fmt.Println(len(c), cap(c))
		c <- 1
		fmt.Println(len(c), cap(c))
		<-c
	*/

	// Буферизированные каналы - это очередь
	/* channel := make(chan string, 3)
	channel <- "a"
	channel <- "b"
	fmt.Println(<-channel)
	channel <- "c"
	fmt.Println(<-channel, <-channel)
	*/

	comparisonBufAndNoBuf()

	time.Sleep(1 * time.Second)
}

func myFunc() {
	fmt.Println("Privet")
}

// Comparison with and without a buffer channel
// Сравнение между каналом с буфером и без буфера (по времени)
func comparisonBufAndNoBuf() {
	cNoBuf := make(chan rune)
	go generate(cNoBuf)
	durationNoBuf := receiveWord(cNoBuf)
	time.Sleep(1 * time.Second)

	cBuf := make(chan rune, 2000)
	go generate(cBuf)
	durationBuf := receiveWord(cBuf)
	time.Sleep(1 * time.Second)

	fmt.Println(durationBuf.Microseconds())
	fmt.Println(durationNoBuf.Microseconds())
	fmt.Println(durationBuf.Microseconds() - durationNoBuf.Microseconds())
}

func generate(c chan rune) {
	defer close(c)
	for _, elem := range word {
		c <- elem
		//fmt.Println("Prodolzaem")
		//fmt.Println(time.Since(start))
	}
}

func receiveWord(c chan rune) time.Duration {
	word := []rune{}
	start := time.Now()
	for r := range c {
		word = append(word, r)
	}
	duration := time.Since(start)
	return duration
}
