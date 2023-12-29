package leetcode

import (
	"fmt"
	"testing"
)

// 判断成环？
// 如何判断成环？
// 找到入度为0的点，从这个点遍历其他所有点,从入度为0的点开始，遍历所有的点
// 如果 不能遍历到所有的点，则成环 ， 如过遍历过程中一个点出现两次，则成环

// 对了是对了
func canFinish(numCourses int, prerequisites [][]int) bool {
	in := make([]int, numCourses)
	h := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		in[prerequisites[i][1]]++
		h[prerequisites[i][0]] = append(h[prerequisites[i][0]], prerequisites[i][1])

	}
	topSort := func() bool {
		qu := make([]int, numCourses+2000)
		head, tail := 0, -1
		for i := 0; i < numCourses; i++ {
			if in[i] == 0 {
				tail++
				qu[tail] = i
			}
		}
		for head <= tail {
			t := qu[head]
			head++
			for _, v := range h[t] {
				in[v]--
				if in[v] == 0 {
					tail++
					qu[tail] = v
				}
			}
		}
		fmt.Println(tail)
		return tail == numCourses-1
	}
	return topSort()
}

func canFinish0(numCourses int, prerequisites [][]int) bool {
	degree := make([]int, 2001)
	h := make([][]int, 2001)
	for i := 0; i < 2001; i++ {
		h[i] = make([]int, 0)
	}
	// 记录一哈总共有   额，5们课程对了的
	for i := 0; i < len(prerequisites); i++ {
		degree[prerequisites[i][1]]++
		h[prerequisites[i][0]] = append(h[prerequisites[i][0]], prerequisites[i][1])
	}
	zero_degrees := make([]int, 0)
	for i := 0; i < len(degree); i++ {
		if degree[i] == 0 {
			// 这个课程是入度为0
			zero_degrees = append(zero_degrees, i)
		}
	}
	if len(zero_degrees) == 0 {
		return false
	}
	res := true
	// 遍历这些点，看是否能遍历到所有的点
	total := make(map[int]int)
	st := make([]bool, numCourses)
	var dfs func(x int)
	dfs = func(x int) {
		//	 当前遍历到哪个点了
		st[x] = true
		total[x]++
		println("x1")
		println(len(h[x]))
		fmt.Printf("%#v \n", h[x])
		for i := 0; i < len(h[x]); i++ {
			println("x2")
			if !st[h[x][i]] {
				dfs(h[x][i])
			} else {
				// 成环了
				res = false
				break
			}
		}
		st[x] = false
	}
	for _, v := range zero_degrees {
		if res {
			fmt.Println(v)
			dfs(v)
			break
		}

	}
	if len(total) != numCourses {
		res = false
	}
	return res
}

func TestCanFinish(t *testing.T) {
	//println(canFinish(700, a))
	//println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
	println(canFinish(6, b))

}

var b = [][]int{
	{5, 3},
	{3, 1},
	{1, 4},
	{4, 2},
	{4, 3},
}

var a = [][]int{
	{78, 355}, {498, 138}, {462, 346}, {57, 190}, {411, 532}, {693, 323}, {96, 76}, {299, 41}, {618, 264}, {472, 467}, {644, 127}, {175, 3}, {296, 673}, {551, 554}, {206, 128}, {317, 226}, {142, 150}, {637, 319}, {436, 433}, {421, 518}, {189, 471}, {292, 179}, {22, 206}, {202, 166}, {396, 583}, {102, 601}, {525, 445}, {33, 368}, {612, 210}, {14, 157}, {303, 37}, {24, 283}, {582, 391}, {477, 522}, {495, 552}, {381, 226}, {310, 111}, {109, 400}, {456, 589}, {62, 366}, {565, 415}, {493, 129}, {124, 250}, {259, 461}, {282, 226}, {107, 36}, {405, 178}, {198, 485}, {575, 39}, {285, 393}, {353, 674}, {444, 590}, {67, 0}, {573, 445}, {642, 339}, {548, 643}, {652, 194}, {455, 648}, {327, 253}, {540, 569}, {372, 233}, {325, 84}, {34, 648}, {310, 681}, {571, 534}, {301, 577}, {147, 41}, {410, 594}, {407, 158}, {256, 417}, {125, 566}, {30, 530}, {122, 188}, {90, 21}, {610, 289}, {105, 616}, {179, 570}, {119, 548}, {652, 589}, {304, 603}, {352, 156}, {207, 508}, {202, 15}, {116, 326}, {511, 0}, {385, 510}, {635, 623}, {358, 20}, {470, 577}, {279, 142}, {400, 241}, {195, 561}, {277, 611}, {399, 456}, {623, 431}, {406, 576}, {184, 322}, {180, 145}, {146, 101}, {475, 398}, {283, 498}, {417, 71}, {446, 688}, {683, 563}, {677, 484}, {92, 425}, {407, 585}, {516, 197}, {559, 57}, {372, 342}, {244, 618}, {150, 451}, {126, 224}, {626, 240}, {224, 211}, {2, 132}, {482, 174}, {630, 397}, {111, 457}, {338, 629}, {572, 87}, {399, 513}, {266, 51}, {239, 538}, {437, 682}, {555, 667}, {360, 229}, {413, 235}, {302, 250}, {56, 7}, {100, 515}, {264, 123}, {545, 386}, {268, 367}, {638, 199}, {432, 69}, {71, 467}, {424, 584}, {491, 134}, {69, 543}, {223, 355}, {596, 315}, {217, 619}, {550, 371}, {482, 120}, {639, 304}, {575, 89}, {116, 252}, {622, 591}, {214, 381}, {467, 73}, {253, 227}, {86, 3}, {384, 543}, {73, 610}, {169, 87}, {11, 181}, {327, 570}, {424, 124}, {139, 32}, {636, 324}, {141, 641}, {563, 12}, {654, 384}, {307, 569}, {536, 27}, {245, 176}, {423, 575}, {640, 682}, {544, 207}, {513, 653}, {241, 180}, {561, 673}, {280, 504}, {207, 144}, {613, 574}, {158, 458}, {484, 685}, {604, 416}, {533, 254}, {73, 392}, {501, 653}, {517, 624}, {299, 311}, {66, 116}, {654, 103}, {584, 362}, {110, 483}, {587, 297}, {447, 665}, {472, 484}, {561, 699}, {374, 174}, {14, 105}, {585, 250}, {78, 186}, {454, 423}, {615, 109}, {17, 635}, {643, 124}, {347, 370}, {601, 329}, {125, 117}, {36, 130}, {53, 281}, {84, 279}, {537, 470}, {391, 158}, {5, 247}, {421, 637}, {494, 529}, {192, 370}, {489, 85}, {508, 407}, {605, 185}, {325, 102}, {542, 353}, {215, 509}, {87, 55}, {40, 502}, {569, 167}, {268, 432}, {299, 113}, {330, 490}, {553, 449}, {679, 549}, {456, 420}, {82, 276}, {673, 230}, {220, 68}, {85, 679}, {432, 449}, {628, 395}, {692, 659}, {652, 510}, {5, 264}, {431, 353}, {31, 661}, {289, 205}, {575, 81}, {278, 579}, {336, 383}, {244, 223}, {139, 560}, {687, 471}, {516, 619}, {20, 647}, {292, 299}, {687, 590}, {549, 653}, {185, 579}, {92, 334}, {475, 178}, {429, 290}, {248, 226}, {154, 636}, {269, 455}, {40, 179}, {556, 21}, {193, 691}, {433, 292}, {119, 200}, {173, 287}, {515, 436}, {17, 65}, {72, 159}, {347, 305}, {264, 318}, {231, 663}, {66, 517}, {46, 633}, {512, 74}, {602, 647}, {611, 277}, {656, 187}, {390, 595}, {502, 85}, {461, 224}, {469, 36}, {598, 267}, {25, 452}, {228, 388}, {588, 487}, {663, 283}, {176, 359}, {437, 475}, {416, 253}, {349, 277}, {685, 421}, {195, 562}, {698, 457}, {167, 590}, {608, 199}, {460, 362}, {304, 163}, {245, 537}, {425, 280}, {86, 551}, {31, 76}, {17, 680}, {430, 100}, {560, 379}, {666, 163}, {316, 266}, {76, 459}, {534, 174}, {595, 10}, {302, 592}, {527, 80}, {389, 600}, {196, 61}, {247, 47}, {532, 661}, {45, 127}, {526, 228}, {672, 500}, {367, 185}, {204, 88}, {676, 243}, {642, 123}, {226, 652}, {272, 569}, {351, 489}, {411, 67}, {618, 388}, {465, 695}, {235, 632}, {387, 626}, {560, 443}, {444, 611}, {681, 434}, {109, 207}, {243, 481}, {135, 611}, {450, 2}, {391, 535}, {537, 211}, {513, 282}, {357, 405}, {9, 242}, {456, 360}, {636, 614}, {404, 572}, {418, 293}, {277, 319}, {375, 294}, {422, 612}, {485, 585}, {181, 100}, {198, 433}, {664, 37}, {673, 221}, {525, 524}, {124, 70}, {468, 122}, {277, 461}, {502, 234}, {101, 266}, {226, 616}, {27, 382}, {478, 675}, {610, 617}, {143, 254}, {552, 472}, {49, 622}, {70, 693}, {188, 398}, {569, 505}, {82, 426}, {670, 106}, {34, 186}, {665, 277}, {659, 178}, {477, 325}, {660, 139}, {540, 680}, {460, 552}, {154, 98}, {94, 471}, {238, 245}, {36, 50}, {290, 662}, {485, 91}, {600, 536}, {242, 267}, {65, 79}, {87, 232}, {642, 107}, {81, 627}, {395, 404}, {247, 394}, {60, 458}, {697, 376}, {303, 153}, {438, 318}, {386, 221}, {115, 302}, {653, 292}, {192, 29}, {488, 320}, {497, 588}, {272, 361}, {121, 317}, {577, 136}, {168, 261}, {699, 444}, {198, 390}, {631, 138}, {504, 384}, {455, 27}, {410, 298}, {656, 67}, {177, 651}, {299, 629}, {490, 420}, {156, 111}, {504, 472}, {362, 679}, {174, 578}, {174, 596}, {429, 654}, {638, 399}, {450, 62}, {317, 186}, {165, 228}, {502, 251}, {247, 681}, {138, 193}, {82, 207}, {413, 142}, {358, 626}, {671, 385}, {590, 626}, {422, 92}, {630, 94}, {418, 114}, {317, 301}, {521, 216}, {93, 284}, {77, 226}, {173, 492}, {674, 153}, {453, 146}, {426, 513}, {173, 203}, {154, 630}, {599, 583}, {23, 698}, {465, 534}, {622, 574}, {461, 3}, {377, 244}, {345, 522}, {431, 469}, {499, 48}, {90, 643}, {141, 461}, {548, 349}, {166, 215}, {91, 191}, {127, 477}, {649, 432}, {576, 238}, {623, 47}, {390, 178}, {571, 315}, {394, 184}, {133, 523}, {548, 475}, {402, 150}, {265, 546}, {53, 595}, {418, 7}, {328, 371}, {581, 261}, {10, 6}, {255, 420}, {143, 31}, {644, 50}, {64, 38}, {372, 304}, {282, 292}, {12, 110}, {152, 311}, {402, 385}, {551, 152}, {682, 206}, {617, 239}, {176, 304}, {190, 96}, {301, 132}, {638, 156}, {220, 242}, {128, 227}, {446, 306}, {678, 188}, {329, 513}, {343, 268}, {384, 212}, {218, 391}, {210, 231}, {157, 588}, {297, 217}, {101, 489}, {475, 581}, {685, 598}, {309, 82}, {383, 547}, {65, 259}, {247, 137}, {109, 315}, {203, 551}, {30, 475}, {629, 42}, {78, 17}, {128, 175}, {149, 664}, {104, 533}, {154, 623}, {542, 283}, {580, 229}, {436, 222}, {112, 386}, {84, 273}, {324, 466}, {636, 473}, {106, 282}, {246, 53}, {671, 299}, {144, 427}, {176, 242}, {213, 386}, {180, 673}, {569, 459}, {549, 571}, {528, 642}, {696, 437}, {110, 65}, {626, 399}, {14, 1}, {640, 404}, {555, 306}, {200, 340}, {23, 176}, {601, 445}, {214, 374}, {229, 683}, {299, 156}, {38, 73}, {411, 696}, {584, 89}, {145, 441}, {170, 178}, {97, 306}, {242, 129}, {347, 537}, {593, 188}, {91, 309}, {673, 136}, {101, 18}, {206, 323}, {265, 180}, {599, 333}, {92, 121}, {315, 392}, {513, 659}, {195, 435}, {21, 263}, {623, 294}, {99, 192}, {24, 61}, {169, 325}, {161, 162}, {251, 614}, {90, 122}, {512, 638}, {553, 644}, {233, 409}, {113, 624}, {293, 566}, {57, 222}, {687, 3}, {242, 219}, {194, 101}, {289, 265}, {312, 391}, {265, 358}, {111, 339}, {20, 622}, {100, 31}, {3, 521}, {114, 272}, {631, 0}, {620, 329}, {536, 135}, {282, 85}, {688, 515}, {6, 411}, {565, 122}, {226, 241}, {169, 395}, {398, 197}, {437, 25}, {532, 155}, {144, 504}, {202, 350}, {503, 543}, {388, 521}, {473, 407}, {575, 196}, {486, 251}, {571, 322}, {260, 160}, {486, 175}, {347, 629}, {676, 396}, {269, 201}, {12, 475}, {7, 266}, {218, 521}, {125, 451}, {212, 580}, {136, 521}, {357, 220}, {323, 40}, {19, 8}, {424, 152}, {31, 500}, {357, 570}, {401, 241}, {235, 211}, {658, 168}, {92, 122}, {599, 507}, {232, 665}, {237, 417}, {624, 448}, {620, 612}, {547, 361}, {648, 61}, {70, 353}, {381, 3}, {600, 206}, {147, 457}, {432, 96}, {348, 463}, {423, 339}, {134, 131}, {36, 696}, {136, 253}, {137, 375}, {358, 344}, {264, 642}, {454, 278}, {308, 187}, {674, 155}, {51, 476}, {571, 418}, {22, 17}, {665, 87}, {445, 557}, {278, 531}, {188, 690}, {80, 514}, {14, 230}, {134, 278}, {367, 584}, {599, 264}, {579, 638}, {86, 7}, {563, 2}, {150, 73}, {685, 237}, {158, 553}, {171, 593}, {64, 145}, {417, 525}, {620, 593}, {666, 137}, {514, 620}, {559, 357}, {609, 267}, {282, 415}, {417, 499}, {107, 112}, {191, 140}, {26, 458}, {467, 382}, {663, 19}, {281, 146}, {363, 515}, {179, 492}, {20, 164}, {91, 318}, {604, 335}, {50, 286}, {228, 374}, {317, 693}, {430, 280}, {422, 655}, {448, 255}, {81, 596}, {430, 326}, {482, 8}, {373, 151}, {470, 437}, {287, 124}, {410, 543}, {186, 209}, {588, 167}, {34, 363}, {50, 262}, {515, 548}, {194, 65}, {590, 617}, {282, 149}, {283, 27}, {372, 314}, {425, 685}, {563, 605}, {435, 438}, {33, 262}, {159, 657}, {246, 212}, {133, 177}, {151, 449}, {294, 45}, {62, 319}, {145, 267}, {78, 359}, {281, 70}, {289, 255}, {605, 513}, {56, 656}, {422, 519}, {649, 261}, {605, 403}, {619, 102}, {597, 362}, {23, 418}, {666, 633}, {29, 474}, {20, 261}, {680, 531}, {645, 316}, {472, 378}, {643, 410}, {534, 387}, {580, 627}, {545, 456}, {456, 694}, {462, 48}, {693, 47}, {231, 160}, {608, 674}, {547, 554}, {287, 316}, {477, 121}, {170, 104}, {534, 346}, {484, 367}, {512, 61}, {117, 103}, {22, 159}, {305, 544}, {307, 685}, {96, 196}, {669, 230}, {200, 55}, {620, 199}, {323, 267}, {248, 470}, {521, 421}, {651, 263}, {17, 512}, {569, 421}, {547, 174}, {675, 418}, {286, 531}, {191, 663}, {623, 175}, {2, 549}, {350, 588}, {324, 669}, {30, 434}, {227, 100}, {662, 140}, {418, 49}, {21, 24}, {24, 55}, {375, 608}, {550, 211}, {304, 316}, {211, 438}, {583, 256}, {507, 120}, {434, 669}, {62, 425}, {191, 0}, {574, 401}, {287, 204}, {409, 197}, {339, 9}, {451, 195}, {185, 561}, {364, 528}, {383, 149}, {479, 284}, {411, 505}, {102, 82}, {318, 637}, {602, 639}, {190, 309}, {538, 0}, {182, 62}, {663, 630}, {383, 528}, {551, 647}, {309, 415}, {567, 437}, {523, 74}, {240, 649}, {570, 623}, {458, 528}, {227, 182}, {208, 5}, {596, 359}, {344, 366}, {414, 13}, {517, 664}, {688, 258}, {279, 177}, {442, 568}, {659, 284}, {560, 424}, {487, 389}, {309, 68}, {420, 285}, {292, 92}, {243, 424}, {655, 678}, {699, 80}, {690, 549}, {262, 361}, {629, 315}, {434, 135}, {323, 289}, {696, 587}, {431, 695}, {201, 123}, {445, 91}, {526, 71}, {506, 588}, {324, 677}, {404, 47}, {61, 338}, {42, 606}, {91, 489}, {591, 347}, {479, 364}, {478, 194}, {649, 361}, {251, 627}, {312, 142}, {53, 293}, {98, 378}, {361, 289}, {553, 4}, {470, 475}, {336, 400}, {364, 8}, {403, 194}, {337, 119}, {622, 26}, {393, 309}, {541, 221}, {617, 511}, {289, 482}, {359, 581}, {311, 155}, {93, 277}, {601, 34}, {280, 564}, {57, 499}, {424, 59}, {313, 350}, {78, 194}, {453, 496}, {199, 667}, {429, 431}, {310, 474}, {332, 287}, {613, 3}, {515, 637}, {113, 106}, {617, 404}, {625, 114}, {190, 376}, {232, 306}, {568, 10}, {594, 639}, {611, 399}, {566, 414}, {254, 76}, {136, 310}, {250, 644}, {517, 330}, {530, 288}, {595, 38}, {321, 583}, {600, 383}, {244, 570}, {39, 47}, {329, 31}, {585, 158}, {547, 527}, {580, 529}, {647, 243}, {340, 56}, {150, 96}, {49, 605}, {258, 164}, {295, 155}, {473, 310}, {302, 47}, {191, 302}, {533, 310}, {462, 78}, {123, 161}, {380, 358}, {564, 218}, {601, 13}, {158, 292}, {310, 0}, {49, 240}, {584, 404}, {262, 184}, {7, 64}, {207, 68}, {430, 526}, {140, 338}, {252, 593}, {394, 274}, {467, 393}, {697, 676}, {42, 283}, {284, 418}, {460, 245}, {409, 43}, {12, 576}, {53, 314}, {492, 497}, {426, 618}, {488, 264}, {174, 649}, {496, 215}, {103, 391}, {579, 594}, {339, 303}, {537, 51}, {137, 527}, {200, 105}, {467, 9}, {79, 85}, {299, 557}, {611, 244}, {206, 97}, {527, 573}, {252, 586}, {410, 6}, {53, 465}, {610, 321}, {405, 214}, {96, 25}, {408, 495}, {554, 18}, {68, 24}, {238, 72}, {22, 538}, {477, 568}, {441, 310}, {494, 645}, {393, 254}, {4, 632}, {68, 446}, {278, 317}, {170, 490}, {689, 159}, {258, 221}, {340, 66}, {55, 356}, {330, 443}, {437, 192}, {674, 201}, {405, 344}, {100, 618}, {87, 195}, {129, 126}, {12, 432}, {273, 465}, {416, 647}, {434, 509}, {415, 538}, {291, 680}, {222, 623}, {142, 421}, {453, 232}, {250, 186}, {365, 30}, {451, 637}, {684, 437}, {86, 507}, {231, 5}, {668, 235}, {349, 541}, {444, 90}, {468, 617}, {74, 624}, {377, 571}, {671, 290}, {606, 588}, {232, 430}, {353, 49}, {44, 658}, {224, 660}, {49, 86}, {27, 236}, {636, 444}, {49, 559}, {121, 640}, {315, 214}, {146, 522}, {385, 52}, {162, 552}, {309, 37}, {231, 559}, {281, 691}, {266, 173}, {532, 512}, {292, 347}, {476, 524}, {265, 454}, {646, 625}, {629, 618}, {695, 545}, {363, 291}, {354, 277}, {154, 384}, {545, 254}, {184, 557}, {301, 629}, {275, 10}, {301, 486}, {141, 139}, {65, 299}, {285, 503}, {629, 366}, {34, 360}, {194, 529}, {245, 82}, {17, 687}, {283, 350}, {341, 260}, {398, 650}, {85, 434}, {551, 170}, {322, 514}, {264, 458}, {697, 507}, {97, 269}, {646, 141}, {395, 578}, {581, 338}, {569, 388}, {340, 537}, {241, 414}, {383, 457}, {489, 160}, {346, 273}, {183, 308}, {5, 14}, {479, 563}, {605, 342}, {370, 178}, {615, 648}, {548, 377}, {630, 405}, {697, 389}, {639, 470}, {212, 390}, {34, 355}, {635, 334}, {3, 175}, {465, 524}, {132, 435}, {466, 457}, {433, 469}, {178, 686}, {38, 699}, {367, 539}, {210, 130}, {670, 534}, {374, 182}, {35, 193}, {249, 145}, {578, 334}, {121, 343}, {287, 14}, {219, 446}, {319, 236}, {98, 632}, {646, 599}, {564, 509}, {402, 120}, {229, 649}, {416, 167}, {223, 568}, {29, 609}, {684, 138}, {540, 536}, {178, 324}, {184, 244}, {695, 243}, {143, 295}, {517, 401}, {58, 114}, {521, 413}, {458, 135}, {372, 72}, {253, 554}, {517, 111}, {663, 388}, {92, 265}, {49, 617}, {45, 275}, {352, 158}, {176, 544}, {81, 591}, {248, 429}, {14, 56}, {52, 61}, {348, 4}, {238, 565}, {612, 395}, {225, 4}, {639, 215}, {556, 342}, {525, 606}, {215, 309}, {239, 617}, {53, 206}, {409, 342}, {615, 4}, {111, 517}, {653, 431}, {550, 153}, {247, 495}, {51, 548}, {634, 529}, {179, 348}, {204, 694}, {449, 5}, {214, 371}, {588, 56}, {258, 94}, {545, 214}, {683, 108}, {82, 330}, {292, 209}, {2, 17}, {211, 580}, {428, 330}, {254, 88}, {559, 86}, {413, 43}, {518, 685}, {546, 381}, {223, 244}, {165, 159}, {441, 539}, {195, 526}, {70, 417}, {156, 308}, {492, 20}, {237, 133}, {292, 450}, {212, 545}, {493, 114}, {14, 334}, {490, 12}, {524, 426}, {634, 51}, {447, 406}, {50, 159}, {7, 641}, {430, 97}, {281, 371}, {307, 538}, {175, 181}, {540, 430}, {62, 235}, {553, 645}, {464, 626}, {5, 567}, {399, 451}, {605, 313}, {185, 519}, {605, 308}, {583, 475}, {130, 490}, {96, 650}, {577, 300}, {151, 191}, {684, 387}, {280, 413}, {558, 523}, {434, 589}, {201, 12}, {469, 591}, {648, 329}, {350, 90}, {384, 613}, {361, 238}, {334, 447}, {503, 510}, {188, 46}, {265, 607}, {357, 36}, {647, 416}, {403, 181}, {395, 631}, {135, 104}, {382, 303}, {573, 60}, {133, 82}, {118, 437}, {100, 502}, {699, 535}, {276, 348}, {183, 587}, {538, 454}, {438, 612}, {56, 48}, {300, 52}, {172, 681}, {629, 80}, {262, 197}, {291, 387}, {432, 461}, {639, 349}, {426, 310}, {418, 513}, {586, 96}, {533, 86}, {293, 620}, {273, 430}, {577, 546}, {279, 252}, {144, 303}, {577, 9}, {203, 391}, {429, 374}, {281, 73}, {641, 305}, {238, 531}, {380, 383}, {12, 434}, {61, 370}, {597, 50}, {154, 225}, {398, 331}, {205, 192}, {258, 110}, {472, 161}, {369, 491}, {119, 237}, {211, 305}, {246, 546}, {430, 539}, {440, 631}, {258, 394}, {308, 42}, {414, 175}, {698, 438}, {683, 94}, {441, 420}, {181, 58}, {518, 562}, {501, 400}, {461, 596}, {368, 665}, {188, 452}, {699, 283}, {132, 528}, {349, 238}, {314, 92}, {257, 236}}
