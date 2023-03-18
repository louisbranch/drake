// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package translations

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en":    &dictionary{index: enIndex, data: enData},
		"pt_BR": &dictionary{index: pt_BRIndex, data: pt_BRData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%d days ago":      8,
	"%d hours ago":     7,
	"%d mins ago":      6,
	"%d years":         41,
	"%d years or more": 42,
	"100%%%%, a technological civilization will always eventually develop if life is present": 39,
	"100%%%%, if conditions are favorable, life is inevitable":                                33,
	"2 or 3":                  28,
	"4 or 5":                  29,
	"A few dozens":            14,
	"A few hundred":           19,
	"A few hundred billion":   22,
	"A few hundred million":   21,
	"A few hundred thousand":  20,
	"A few hundred trillion":  23,
	"A few hundreds":          15,
	"A few millions":          17,
	"A few thousands":         16,
	"A very small percentage": 35,
	"A very small percentage, life is very rare": 32,
	"About":        50,
	"About 10%%%%": 38,
	"About how many stars are there in our galaxy?":                                   18,
	"Approximately what percentage of the stars in our galaxy have planets in orbit?": 24,
	"Are we alone in the Universe?":                                                   66,
	"Astronomy Education":                                                             57,
	"Back":                                                                            79,
	"Civilizations":                                                                   64,
	"Create Session":                                                                  72,
	"Created":                                                                         77,
	"Do you agree with your most recent estimation?":                                  43,
	"Dozens":              1,
	"Drake Equation":      56,
	"Drake Equation - %s": 84,
	"Estimate the number of detectable alien civilizations in the Milky Way using the Drake Equation.": 67,
	"Final Estimations": 63,
	"Home":              60,
	"How many technological advanced civilizations exist in the Milky Way?": 12,
	"Hundreds": 2,
	"If you would like to contribute to the project, for example, adding more translations, get in touch:": 54,
	"Initial Predictions":   62,
	"Internal Server Error": 58,
	"Join":                  82,
	"Join Session":          73,
	"Latest Sessions":       75,
	"Learning Goals:":       68,
	"Less than 1%%%%":       37,
	"Less than one min ago": 5,
	"Millions":              4,
	"More than 5":           30,
	"Name":                  76,
	"Next":                  85,
	"No available sessions": 78,
	"No, I am more confident with my initial prediction":                                   46,
	"No, any guess is good as mine":                                                        47,
	"On average, how many habitable (Earth-like) planets are there per planetary system?":  26,
	"On what fraction of habitable planets will life develop?":                             31,
	"On what fraction of planets on which life develops will life evolve to intelligence?": 34,
	"Only us":                   0,
	"Only us, we are all alone": 13,
	"Page Not Found":            59,
	"Participants":              65,
	"References":                51,
	"Results for Session %s":    61,
	"See Results":               83,
	"Select the choice that best agrees with what you think.": 86,
	"Session %s":  80,
	"Sessions":    74,
	"Share Link:": 81,
	"Source Code": 55,
	"The average distance to another civilization would be %d light-years.":                                                                            49,
	"The questions for the Drake Equation on the survey are from LoPresto and Hubble-Zdanowski (2012). The code is open-source using the MIT License.": 53,
	"There would be no other civilization in the galaxy.":                                                                                              48,
	"This is a difference of %d orders of magnitude from your initial prediction of %d.":                                                               11,
	"This project was created as part of the course, Principles and Practices in Science Education, at the University of Toronto with the intention of being a free resource for educators to introduce the Drake Equation to a wider audience.": 52,
	"Thousands": 3,
	"Very few (not even one in every solar system)": 27,
	"Very few stars have planets in orbit":          25,
	"What fraction of planets on which life evolves to intelligence will the intelligence develop a technological civilization capable of radio communication?": 36,
	"What is the average lifetime of a technological civilization capable of communication?":                                                                    40,
	"Yes, I am more confident using the Drake Equation calculation":                                                                                             44,
	"Yes, the Drake Equation includes factors I didn't think during my initial prediction":                                                                      45,
	"You have estimated that there are %d civilizations in the Milky Way.":                                                                                      9,
	"You have predicted that there are %d civilizations in the Milky Way.":                                                                                      10,
	"to compare your initial guess with the final value":                                                                                                        71,
	"to think about the size and composition of the galaxy and how it affects the possibility of intelligent life":                                              69,
	"to understand and estimate the terms of the Drake Equation":                                                                                                70,
}

var enIndex = []uint32{ // 88 elements
	// Entry 0 - 1F
	0x00000000, 0x00000008, 0x0000000f, 0x00000018,
	0x00000022, 0x0000002b, 0x00000041, 0x00000050,
	0x00000060, 0x0000006f, 0x00000108, 0x000001a1,
	0x00000258, 0x0000029e, 0x000002b8, 0x000002c5,
	0x000002d4, 0x000002e4, 0x000002f3, 0x00000321,
	0x0000032f, 0x00000346, 0x0000035c, 0x00000372,
	0x00000389, 0x000003d9, 0x000003fe, 0x00000452,
	0x00000480, 0x00000487, 0x0000048e, 0x0000049a,
	// Entry 20 - 3F
	0x000004d3, 0x000004fe, 0x00000535, 0x0000058a,
	0x000005a2, 0x0000063c, 0x0000064a, 0x00000655,
	0x000006ab, 0x00000702, 0x0000070e, 0x00000722,
	0x00000751, 0x0000078f, 0x000007e4, 0x00000817,
	0x00000835, 0x00000869, 0x00000951, 0x00000957,
	0x00000962, 0x00000a4d, 0x00000ade, 0x00000b43,
	0x00000b4f, 0x00000b5e, 0x00000b72, 0x00000b88,
	0x00000b97, 0x00000b9c, 0x00000bb6, 0x00000bca,
	// Entry 40 - 5F
	0x00000bdc, 0x00000bea, 0x00000bf7, 0x00000c15,
	0x00000c76, 0x00000c86, 0x00000cf3, 0x00000d2e,
	0x00000d61, 0x00000d70, 0x00000d7d, 0x00000d86,
	0x00000d96, 0x00000d9b, 0x00000da3, 0x00000db9,
	0x00000dbe, 0x00000dcc, 0x00000dd8, 0x00000ddd,
	0x00000de9, 0x00000e00, 0x00000e05, 0x00000e3d,
} // Size: 376 bytes

const enData string = "" + // Size: 3645 bytes
	"\x02Only us\x02Dozens\x02Hundreds\x02Thousands\x02Millions\x02Less than " +
	"one min ago\x02%[1]d mins ago\x02%[1]d hours ago\x02%[1]d days ago\x14" +
	"\x01\x81\x01\x00=\x01G\x02You have estimated that there is only 1 civili" +
	"zation in the Milky Way.\x00H\x02You have estimated that there are %[1]d" +
	" civilizations in the Milky Way.\x14\x01\x81\x01\x00=\x01G\x02You have p" +
	"redicted that there is only 1 civilization in the Milky Way.\x00H\x02You" +
	" have predicted that there are %[1]d civilizations in the Milky Way.\x14" +
	"\x01\x81\x01\x00=\x01T\x02This is a difference of 1 order of magnitude f" +
	"rom your initial prediction of %[2]d.\x00Y\x02This is a difference of %[" +
	"1]d orders of magnitude from your initial prediction of %[2]d.\x02How ma" +
	"ny technological advanced civilizations exist in the Milky Way?\x02Only " +
	"us, we are all alone\x02A few dozens\x02A few hundreds\x02A few thousand" +
	"s\x02A few millions\x02About how many stars are there in our galaxy?\x02" +
	"A few hundred\x02A few hundred thousand\x02A few hundred million\x02A fe" +
	"w hundred billion\x02A few hundred trillion\x02Approximately what percen" +
	"tage of the stars in our galaxy have planets in orbit?\x02Very few stars" +
	" have planets in orbit\x02On average, how many habitable (Earth-like) pl" +
	"anets are there per planetary system?\x02Very few (not even one in every" +
	" solar system)\x022 or 3\x024 or 5\x02More than 5\x02On what fraction of" +
	" habitable planets will life develop?\x02A very small percentage, life i" +
	"s very rare\x02100%%, if conditions are favorable, life is inevitable" +
	"\x02On what fraction of planets on which life develops will life evolve " +
	"to intelligence?\x02A very small percentage\x02What fraction of planets " +
	"on which life evolves to intelligence will the intelligence develop a te" +
	"chnological civilization capable of radio communication?\x02Less than 1%" +
	"%\x02About 10%%\x02100%%, a technological civilization will always event" +
	"ually develop if life is present\x02What is the average lifetime of a te" +
	"chnological civilization capable of communication?\x02%[1]d years\x02%[1" +
	"]d years or more\x02Do you agree with your most recent estimation?\x02Ye" +
	"s, I am more confident using the Drake Equation calculation\x02Yes, the " +
	"Drake Equation includes factors I didn't think during my initial predict" +
	"ion\x02No, I am more confident with my initial prediction\x02No, any gue" +
	"ss is good as mine\x02There would be no other civilization in the galaxy" +
	".\x14\x01\x81\x01\x00=\x00N\x02The average distance to another civilizat" +
	"ion would be less than 1 light-year.=\x01D\x02The average distance to an" +
	"other civilization would be 1 light-year.\x00I\x02The average distance t" +
	"o another civilization would be %[1]d light-years.\x02About\x02Reference" +
	"s\x02This project was created as part of the course, Principles and Prac" +
	"tices in Science Education, at the University of Toronto with the intent" +
	"ion of being a free resource for educators to introduce the Drake Equati" +
	"on to a wider audience.\x02The questions for the Drake Equation on the s" +
	"urvey are from LoPresto and Hubble-Zdanowski (2012). The code is open-so" +
	"urce using the MIT License.\x02If you would like to contribute to the pr" +
	"oject, for example, adding more translations, get in touch:\x02Source Co" +
	"de\x02Drake Equation\x02Astronomy Education\x02Internal Server Error\x02" +
	"Page Not Found\x02Home\x02Results for Session %[1]s\x02Initial Predictio" +
	"ns\x02Final Estimations\x02Civilizations\x02Participants\x02Are we alone" +
	" in the Universe?\x02Estimate the number of detectable alien civilizatio" +
	"ns in the Milky Way using the Drake Equation.\x02Learning Goals:\x02to t" +
	"hink about the size and composition of the galaxy and how it affects the" +
	" possibility of intelligent life\x02to understand and estimate the terms" +
	" of the Drake Equation\x02to compare your initial guess with the final v" +
	"alue\x02Create Session\x02Join Session\x02Sessions\x02Latest Sessions" +
	"\x02Name\x02Created\x02No available sessions\x02Back\x02Session %[1]s" +
	"\x02Share Link:\x02Join\x02See Results\x02Drake Equation - %[1]s\x02Next" +
	"\x02Select the choice that best agrees with what you think."

var pt_BRIndex = []uint32{ // 88 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000c, 0x00000014, 0x0000001d,
	0x00000026, 0x0000002f, 0x00000046, 0x00000058,
	0x0000006b, 0x0000007d, 0x000000ff, 0x0000018f,
	0x00000250, 0x0000029b, 0x000002c7, 0x000002d7,
	0x000002e8, 0x000002f8, 0x00000308, 0x00000334,
	0x00000345, 0x00000362, 0x0000037f, 0x0000039c,
	0x000003ba, 0x0000041c, 0x0000044b, 0x000004ac,
	0x000004d9, 0x000004e0, 0x000004e7, 0x000004f1,
	// Entry 20 - 3F
	0x00000539, 0x00000567, 0x000005a9, 0x000005f5,
	0x0000060d, 0x00000698, 0x000006a5, 0x000006b6,
	0x0000070d, 0x0000076a, 0x00000775, 0x00000788,
	0x000007b8, 0x000007fd, 0x0000085a, 0x00000894,
	0x000008c9, 0x000008f7, 0x000009d2, 0x000009d8,
	0x000009e5, 0x00000ad8, 0x00000b64, 0x00000bcb,
	0x00000bd9, 0x00000be9, 0x00000c02, 0x00000c1b,
	0x00000c33, 0x00000c3b, 0x00000c59, 0x00000c6d,
	// Entry 40 - 5F
	0x00000c80, 0x00000c8f, 0x00000c9d, 0x00000cbb,
	0x00000d24, 0x00000d3e, 0x00000da0, 0x00000dd3,
	0x00000e05, 0x00000e13, 0x00000e24, 0x00000e2d,
	0x00000e3f, 0x00000e44, 0x00000e4b, 0x00000e67,
	0x00000e6e, 0x00000e7c, 0x00000e93, 0x00000e99,
	0x00000ea8, 0x00000ec0, 0x00000ec9, 0x00000f08,
} // Size: 376 bytes

const pt_BRData string = "" + // Size: 3848 bytes
	"\x02Apenas nós\x02Dezenas\x02Centenas\x02Milhares\x02Milhões\x02Menos de" +
	" um min atrás\x02%[1]d mins atrás\x02%[1]d horas atrás\x02%[1]d dias atr" +
	"ás\x14\x01\x81\x01\x00=\x01=\x02Você estimou que há apenas 1 civilizaçã" +
	"o na Via Láctea.\x00;\x02Você estimou que há %[1]d civilizações na Via L" +
	"áctea.\x14\x01\x81\x01\x00=\x01D\x02Sua predição é de que há apenas 1 c" +
	"ivilização na Via Láctea.\x00B\x02Sua predição é de que há %[1]d civiliz" +
	"ações na Via Láctea.\x14\x01\x81\x01\x00=\x01Y\x02Há uma differença de 1" +
	" ordem de magnitude a partir da sua predição inicial de %[2]d.\x00^\x02H" +
	"á uma differença de %[1]d ordens de magnitude a partir da sua predição " +
	"inicial de %[2]d.\x02Quantas civilizações tecnologicamente avançadas exi" +
	"stem na Via Láctea?\x02Apenas nós, estamos completamente sozinhos\x02Alg" +
	"umas dezenas\x02Algumas centenas\x02Alguns milhares\x02Alguns milhões" +
	"\x02Quantas estrelas existem em nossa galáxia?\x02Algumas centenas\x02Al" +
	"gumas centenas de milhares\x02Algumas centenas de milhões\x02Algumas cen" +
	"tenas de bilhões\x02Algumas centenas de trilhões\x02Aproximadamente qual" +
	" é a porcentagem de estrelas em nossa galáxia que têm planetas em órbita" +
	"?\x02Muito poucas estrelas têm planetas em órbita\x02Na média, quantos p" +
	"lanetas habitáveis (parecidos com a Terra) existem por sistema planetári" +
	"o?\x02Muito poucos (menos de um por sistema solar)\x022 ou 3\x024 ou 5" +
	"\x02Mais de 5\x02Qual a fração de planetas habitáveis que a vida irá se " +
	"desenvolver?\x02Uma pequena porcentagem, a vida é muito rara\x02100%%, s" +
	"e as condições forem favoráveis, a vida é inevitável\x02Qual a fração de" +
	" planetas com vida na qual vida inteligente irá evoluir?\x02Uma pequena " +
	"porcentagem\x02Qual a fração de planetas com vida inteligente que irá de" +
	"senvolver uma civilização tecnologicamente capaz de se comunicar por rád" +
	"io?\x02Menos de 1%%\x02Em torno de 10%%\x02100%%, uma civilização tecnol" +
	"ogicamente capaz sempre se desenvolverá se houver vida\x02Qual é o tempo" +
	" de vida médio de uma civilização tecnologicamente capaz de comunicação?" +
	"\x02%[1]d anos\x02%[1]d anos ou mais\x02Você concoda com essa estimativa" +
	" mais recente?\x02Sim, eu estou mais confiante usando o cálculo da Equaç" +
	"ão de Drake\x02Sim, a Equação de Drake inclui fatores que eu não pensei" +
	" durante minha predição inicial\x02Não, eu estou mais confiante no minha" +
	" predição inicial\x02Não, qualquer predição é tão boa quanto a minha\x02" +
	"Não haveria outra civilização na galáxia.\x14\x01\x81\x01\x00=\x00G\x02A" +
	" distância média para outra civilização seria menos de 1 ano-luz.=\x01A" +
	"\x02A distância média para outra civilização seria de 1 ano-luz.\x00F" +
	"\x02A distância média para outra civilização seria de %[1]d anos-luz." +
	"\x02Sobre\x02Referências\x02Esse projeto foi criado como parte do curso " +
	"Princípios e Práticas em Ciência da Educação, na Unversidade de Toronto," +
	" com a intenção de ser um recurso gratuito a educadores para introduzir " +
	"a Equação de Drake para uma ampla audiência.\x02As questões para a Equaç" +
	"ão de Drake na pesquisa são de LoPresto and Hubble-Zdanowski (2012). O " +
	"código é aberto usando a Licença MIT.\x02Se você quiser contribuir com o" +
	" projeto, por exemplo, adicionando mais traduções, entre em contato:\x02" +
	"Código Fonte\x02Equação Drake\x02Educação em Astronomia\x02Erro Interno " +
	"do Servidor\x02Página Não Encontrada\x02Início\x02Resultados para Sessão" +
	" %[1]s\x02Previsões Iniciais\x02Estimativas Finais\x02Civilizações\x02Pa" +
	"rticipantes\x02Estamos sozinhos no Universo?\x02Estime o número de civil" +
	"izações alienígenas detectáveis na Via Láctea usando a Equação de Drake." +
	"\x02Objetivos de Aprendizado:\x02pensar sobre o tamanho e composição da " +
	"galaxia e como afeta a possibilidade de vida inteligente\x02entender e e" +
	"stimar os termos da Equação de Drake\x02comparar sua predição inicial co" +
	"m o valor final\x02Criar Sessão\x02Entra na Sessão\x02Sessões\x02Últimas" +
	" Sessões\x02Nome\x02Criado\x02Nenhuma sessão disponível\x02Voltar\x02Ses" +
	"são %[1]s\x02Compartilhe Endereço:\x02Entre\x02Ver Resultados\x02Equação" +
	" Drake - %[1]s\x02Próximo\x02Selecione a opção que melhor concorde com o" +
	" que você pensa."

	// Total table size 8245 bytes (8KiB); checksum: 889DC49E
