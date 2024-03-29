package fixed

import "testing"

var bincdfTestCases_10000 = []BinCDFCase{
	{
		n:   1520,
		x:   159,
		p:   0.11873904797231849,
		cdf: 0.04620028268945807,
		s:   "+0'0bd3c81f3798ef/56",
	},
	{
		n:   1520,
		x:   1229,
		p:   0.7839039641923319,
		cdf: 0.9916878564399366,
		s:   "+0'fddf415f3ff788/56",
	},
	{
		n:   8359,
		x:   1438,
		p:   0.15563814164613257,
		cdf: 0.999978954555814,
		s:   "+0'fffe9eea8167e0/56",
	},
	{
		n:   8359,
		x:   7304,
		p:   0.8914426144413563,
		cdf: 2.0893519712083095e-07,
		s:   "+0'000003815eadbe/56",
	},
	{
		n:   6514,
		x:   755,
		p:   0.11854916867602958,
		cdf: 0.261625452704057,
		s:   "+0'42f9e2bb2a4548/56",
	},
	{
		n:   6514,
		x:   3579,
		p:   0.5207452062733783,
		cdf: 0.999998354326247,
		s:   "+0'ffffe463e292b0/56",
	},
	{
		n:   7201,
		x:   3881,
		p:   0.5154704290513524,
		cdf: 0.9999685434091704,
		s:   "+0'fffdf03ef89cd8/56",
	},
	{
		n:   7201,
		x:   6248,
		p:   0.8546136681577384,
		cdf: 0.9992980781206945,
		s:   "+0'ffd1ffb47be9c8/56",
	},
	{
		n:   2307,
		x:   448,
		p:   0.14655817727220605,
		cdf: 0.9999999997929367,
		s:   "+0'ffffffff1c54e0/56",
	},
	{
		n:   2307,
		x:   1233,
		p:   0.5271368888528801,
		cdf: 0.7658373074373637,
		s:   "+0'c40de9ed800b70/56",
	},
	{
		n:   5374,
		x:   2292,
		p:   0.4326804987523857,
		cdf: 0.1838209840343222,
		s:   "+0'2f0ee45abef6b6/56",
	},
	{
		n:   5374,
		x:   4613,
		p:   0.836922850197975,
		cdf: 0.9999933534862768,
		s:   "+0'ffff907d70e090/56",
	},
	{
		n:   1916,
		x:   912,
		p:   0.4132377210423095,
		cdf: 0.9999999871660684,
		s:   "+0'ffffffc8e0f168/56",
	},
	{
		n:   9179,
		x:   7498,
		p:   0.8169171980447082,
		cdf: 0.49904593676568776,
		s:   "+0'7fc179799c3d78/56",
	},
	{
		n:   7543,
		x:   1124,
		p:   0.14386415183145668,
		cdf: 0.9011110841688819,
		s:   "+0'e6af374c918600/56",
	},
	{
		n:   7543,
		x:   6626,
		p:   0.8584158728725929,
		cdf: 0.9999998337735162,
		s:   "+0'fffffd36100c58/56",
	},
	{
		n:   2876,
		x:   561,
		p:   0.25017555991832197,
		cdf: 1.410100517569426e-12,
		s:   "+0'00000000018ce8/56",
	},
	{
		n:   2876,
		x:   1258,
		p:   0.48695305013382106,
		cdf: 5.626247261793679e-08,
		s:   "+0'000000f1a53e2b/56",
	},
	{
		n:   2876,
		x:   2455,
		p:   0.894377619774023,
		cdf: 7.670063125541897e-12,
		s:   "+0'00000000086eee/56",
	},
	{
		n:   9689,
		x:   8475,
		p:   0.8756019198408082,
		cdf: 0.39888539306695936,
		s:   "+0'661d5a66131f40/56",
	},
	{
		n:   2290,
		x:   1286,
		p:   0.5955968987354668,
		cdf: 0.0005112133731568636,
		s:   "+0'002180bcb80f81/56",
	},
	{
		n:   2290,
		x:   1308,
		p:   0.5531668858613976,
		cdf: 0.9604933443623838,
		s:   "+0'f5e2e44e0fe610/56",
	},
	{
		n:   10000,
		x:   101,
		p:   0.01,
		cdf: 0.566227688292922,
		s:   "+0'90f44c3b4ecad0/56",
	},
	{
		n:   10000,
		x:   101,
		p:   0.02,
		cdf: 4.501474917622328e-15,
		s:   "+0'00000000000144/56",
	},
	{
		n:   10000,
		x:   201,
		p:   0.02,
		cdf: 0.547137558835646,
		s:   "+0'8c1135019cc4a0/56",
	},
	{
		n:   10000,
		x:   301,
		p:   0.02,
		cdf: 0.9999999999931125,
		s:   "+0'fffffffff86d58/56",
	},
	{
		n:   10000,
		x:   201,
		p:   0.03,
		cdf: 4.605217838591439e-10,
		s:   "+0'00000001fa595b/56",
	},
	{
		n:   10000,
		x:   301,
		p:   0.03,
		cdf: 0.5386527727429063,
		s:   "+0'89e525ead49bd8/56",
	},
	{
		n:   10000,
		x:   401,
		p:   0.03,
		cdf: 0.9999999929691475,
		s:   "+0'ffffffe1cd7ef8/56",
	},
	{
		n:   10000,
		x:   301,
		p:   0.04,
		cdf: 8.03470565061695e-08,
		s:   "+0'000001591685db/56",
	},
	{
		n:   10000,
		x:   401,
		p:   0.04,
		cdf: 0.5336000299771242,
		s:   "+0'889a02f5e57860/56",
	},
	{
		n:   10000,
		x:   501,
		p:   0.04,
		cdf: 0.9999997087529758,
		s:   "+0'fffffb1d1a82a8/56",
	},
	{
		n:   10000,
		x:   401,
		p:   0.05,
		cdf: 1.5350433390141877e-06,
		s:   "+0'000019c0f60019/56",
	},
	{
		n:   10000,
		x:   501,
		p:   0.05,
		cdf: 0.5301601713162643,
		s:   "+0'87b893b571f2e0/56",
	},
	{
		n:   10000,
		x:   601,
		p:   0.05,
		cdf: 0.9999969950013854,
		s:   "+0'ffffcd95a114f0/56",
	},
	{
		n:   10000,
		x:   501,
		p:   0.06,
		cdf: 1.0424321321290476e-05,
		s:   "+0'0000aee41e8122/56",
	},
	{
		n:   10000,
		x:   601,
		p:   0.06,
		cdf: 0.5276289290349718,
		s:   "+0'8712b082a0f310/56",
	},
	{
		n:   10000,
		x:   701,
		p:   0.06,
		cdf: 0.9999850687341421,
		s:   "+0'ffff057eb39260/56",
	},
	{
		n:   10000,
		x:   601,
		p:   0.07,
		cdf: 4.0082008251786455e-05,
		s:   "+0'0002a076ea2330/56",
	},
	{
		n:   10000,
		x:   701,
		p:   0.07,
		cdf: 0.5256687907580677,
		s:   "+0'86923ad8d571d0/56",
	},
	{
		n:   10000,
		x:   801,
		p:   0.07,
		cdf: 0.999951854692482,
		s:   "+0'fffcd8417a8fa8/56",
	},
	{
		n:   10000,
		x:   901,
		p:   0.07,
		cdf: 0.9999999999999839,
		s:   "+0'fffffffffffb78/56",
	},
	{
		n:   10000,
		x:   601,
		p:   0.08,
		cdf: 1.2510230032280923e-14,
		s:   "+0'00000000000385/56",
	},
	{
		n:   10000,
		x:   701,
		p:   0.08,
		cdf: 0.00010893479813500263,
		s:   "+0'0007239f653807/56",
	},
	{
		n:   10000,
		x:   801,
		p:   0.08,
		cdf: 0.5240951388096105,
		s:   "+0'862b19592e0b20/56",
	},
	{
		n:   10000,
		x:   901,
		p:   0.08,
		cdf: 0.9998823774745639,
		s:   "+0'fff84a9f199838/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.08,
		cdf: 0.9999999999996375,
		s:   "+0'ffffffffff99f8/56",
	},
	{
		n:   10000,
		x:   701,
		p:   0.09,
		cdf: 3.3407102663121864e-13,
		s:   "+0'00000000005e08/56",
	},
	{
		n:   10000,
		x:   801,
		p:   0.09,
		cdf: 0.00023579855723553892,
		s:   "+0'000f740b177e47/56",
	},
	{
		n:   10000,
		x:   901,
		p:   0.09,
		cdf: 0.5227973818398446,
		s:   "+0'85d60c996fc230/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.09,
		cdf: 0.99976211651867,
		s:   "+0'fff068fa3a38c8/56",
	},
	{
		n:   10000,
		x:   1101,
		p:   0.09,
		cdf: 0.9999999999957287,
		s:   "+0'fffffffffb4dc0/56",
	},
	{
		n:   10000,
		x:   801,
		p:   0.1,
		cdf: 4.383298587819392e-12,
		s:   "+0'0000000004d1c9/56",
	},
	{
		n:   10000,
		x:   901,
		p:   0.1,
		cdf: 0.0004360015331795163,
		s:   "+0'001c92e4537583/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.1,
		cdf: 0.5217047111595491,
		s:   "+0'858e70a0997198/56",
	},
	{
		n:   10000,
		x:   1101,
		p:   0.1,
		cdf: 0.999579489788011,
		s:   "+0'ffe471026451d0/56",
	},
	{
		n:   10000,
		x:   1201,
		p:   0.1,
		cdf: 0.9999999999684474,
		s:   "+0'ffffffffdd4eb8/56",
	},
	{
		n:   10000,
		x:   901,
		p:   0.11,
		cdf: 3.4778487051805204e-11,
		s:   "+0'00000000263d46/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.11,
		cdf: 0.0007195269113573255,
		s:   "+0'002f27a88d88fe/56",
	},
	{
		n:   10000,
		x:   1101,
		p:   0.11,
		cdf: 0.5207694172346744,
		s:   "+0'855124ffc7a7e0/56",
	},
	{
		n:   10000,
		x:   1201,
		p:   0.11,
		cdf: 0.9993269835703988,
		s:   "+0'ffd3e4a871f820/56",
	},
	{
		n:   10000,
		x:   1301,
		p:   0.11,
		cdf: 0.9999999998351843,
		s:   "+0'ffffffff4ac880/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.12,
		cdf: 1.9071868611833278e-10,
		s:   "+0'00000000d1b289/56",
	},
	{
		n:   10000,
		x:   1101,
		p:   0.12,
		cdf: 0.0010908260758382589,
		s:   "+0'00477d065243f3/56",
	},
	{
		n:   10000,
		x:   1201,
		p:   0.12,
		cdf: 0.5199580297702877,
		s:   "+0'851bf82d27ee50/56",
	},
	{
		n:   10000,
		x:   1301,
		p:   0.12,
		cdf: 0.9990012206480439,
		s:   "+0'ffbe8b4358f110/56",
	},
	{
		n:   10000,
		x:   1401,
		p:   0.12,
		cdf: 0.9999999993391697,
		s:   "+0'fffffffd2968d0/56",
	},
	{
		n:   10000,
		x:   1101,
		p:   0.13,
		cdf: 7.908837848106515e-10,
		s:   "+0'000000036595fe/56",
	},
	{
		n:   10000,
		x:   1201,
		p:   0.13,
		cdf: 0.0015495695970275923,
		s:   "+0'00658d76bdfae1/56",
	},
	{
		n:   10000,
		x:   1301,
		p:   0.13,
		cdf: 0.519246285452704,
		s:   "+0'84ed531696be80/56",
	},
	{
		n:   10000,
		x:   1401,
		p:   0.13,
		cdf: 0.998602426276004,
		s:   "+0'ffa4689a8fcad0/56",
	},
	{
		n:   10000,
		x:   1501,
		p:   0.13,
		cdf: 0.9999999978448786,
		s:   "+0'fffffff6be6b40/56",
	},
	{
		n:   10000,
		x:   1201,
		p:   0.14,
		cdf: 2.6408552616436843e-09,
		s:   "+0'0000000b57a6ac/56",
	},
	{
		n:   10000,
		x:   1301,
		p:   0.14,
		cdf: 0.0020917494283618026,
		s:   "+0'008915bb62e0aa/56",
	},
	{
		n:   10000,
		x:   1401,
		p:   0.14,
		cdf: 0.5186161112434509,
		s:   "+0'84c40684f82550/56",
	},
	{
		n:   10000,
		x:   1501,
		p:   0.14,
		cdf: 0.9981336557026209,
		s:   "+0'ff85aff0479928/56",
	},
	{
		n:   10000,
		x:   1601,
		p:   0.14,
		cdf: 0.9999999940376424,
		s:   "+0'ffffffe6645188/56",
	},
	{
		n:   10000,
		x:   1301,
		p:   0.15,
		cdf: 7.428782056220899e-09,
		s:   "+0'0000001fe80841/56",
	},
	{
		n:   10000,
		x:   1401,
		p:   0.15,
		cdf: 0.002710787635403502,
		s:   "+0'00b1a7783d8ebc/56",
	},
	{
		n:   10000,
		x:   1501,
		p:   0.15,
		cdf: 0.5180537337366561,
		s:   "+0'849f2b63f839a8/56",
	},
	{
		n:   10000,
		x:   1601,
		p:   0.15,
		cdf: 0.9976000043798162,
		s:   "+0'ff62b6c14cff18/56",
	},
	{
		n:   10000,
		x:   1701,
		p:   0.15,
		cdf: 0.9999999855614757,
		s:   "+0'ffffffc1fcacb8/56",
	},
	{
		n:   10000,
		x:   1401,
		p:   0.16,
		cdf: 1.8202078903030618e-08,
		s:   "+0'0000004e2d65bc/56",
	},
	{
		n:   10000,
		x:   1501,
		p:   0.16,
		cdf: 0.003398492333089429,
		s:   "+0'00deb93d6d23b0/56",
	},
	{
		n:   10000,
		x:   1601,
		p:   0.16,
		cdf: 0.5175484494719357,
		s:   "+0'847e0e2093d538/56",
	},
	{
		n:   10000,
		x:   1701,
		p:   0.16,
		cdf: 0.9970079100802202,
		s:   "+0'ff3be90fa5d9c8/56",
	},
	{
		n:   10000,
		x:   1801,
		p:   0.16,
		cdf: 0.9999999686592427,
		s:   "+0'ffffff79647910/56",
	},
	{
		n:   10000,
		x:   1501,
		p:   0.17,
		cdf: 3.983685441397754e-08,
		s:   "+0'000000ab1915ab/56",
	},
	{
		n:   10000,
		x:   1601,
		p:   0.17,
		cdf: 0.0041458119451111336,
		s:   "+0'010fb32eb838ea/56",
	},
	{
		n:   10000,
		x:   1701,
		p:   0.17,
		cdf: 0.5170917995226201,
		s:   "+0'846020cffabcb0/56",
	},
	{
		n:   10000,
		x:   1801,
		p:   0.17,
		cdf: 0.9963645861426014,
		s:   "+0'ff11bfe0600320/56",
	},
	{
		n:   10000,
		x:   1901,
		p:   0.17,
		cdf: 0.9999999378903659,
		s:   "+0'fffffef53dbc30/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.17,
		cdf: 0.999999999999998,
		s:   "+0'ffffffffffff70/56",
	},
	{
		n:   10000,
		x:   1601,
		p:   0.18,
		cdf: 7.939833075415867e-08,
		s:   "+0'0000015503634c/56",
	},
	{
		n:   10000,
		x:   1701,
		p:   0.18,
		cdf: 0.004943394332932138,
		s:   "+0'0143f864fdbda9/56",
	},
	{
		n:   10000,
		x:   1801,
		p:   0.18,
		cdf: 0.5166769999025725,
		s:   "+0'8444f1a12d4c38/56",
	},
	{
		n:   10000,
		x:   1901,
		p:   0.18,
		cdf: 0.9956775874652615,
		s:   "+0'fee4b9f38602a0/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.18,
		cdf: 0.9999998859911728,
		s:   "+0'fffffe1655f808/56",
	},
	{
		n:   10000,
		x:   2101,
		p:   0.18,
		cdf: 0.9999999999999928,
		s:   "+0'fffffffffffdf8/56",
	},
	{
		n:   10000,
		x:   1601,
		p:   0.19,
		cdf: 3.8908639605122645e-15,
		s:   "+0'00000000000118/56",
	},
	{
		n:   10000,
		x:   1701,
		p:   0.19,
		cdf: 1.4630676707434277e-07,
		s:   "+0'0000027461fdda/56",
	},
	{
		n:   10000,
		x:   1801,
		p:   0.19,
		cdf: 0.005781979959190098,
		s:   "+0'017aed86d4b283/56",
	},
	{
		n:   10000,
		x:   1901,
		p:   0.19,
		cdf: 0.5162985393757695,
		s:   "+0'842c241d976c70/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.19,
		cdf: 0.9949544937433208,
		s:   "+0'feb55673a2c3c0/56",
	},
	{
		n:   10000,
		x:   2101,
		p:   0.19,
		cdf: 0.9999998039179546,
		s:   "+0'fffffcb5d582d8/56",
	},
	{
		n:   10000,
		x:   2201,
		p:   0.19,
		cdf: 0.999999999999977,
		s:   "+0'fffffffffff988/56",
	},
	{
		n:   10000,
		x:   1701,
		p:   0.2,
		cdf: 1.3709866654875469e-14,
		s:   "+0'000000000003db/56",
	},
	{
		n:   10000,
		x:   1801,
		p:   0.2,
		cdf: 2.5225944452973685e-07,
		s:   "+0'0000043b723146/56",
	},
	{
		n:   10000,
		x:   1901,
		p:   0.2,
		cdf: 0.006652663703047633,
		s:   "+0'01b3fd2d092f28/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.2,
		cdf: 0.5159518893397638,
		s:   "+0'84156c4b0610f8/56",
	},
	{
		n:   10000,
		x:   2101,
		p:   0.2,
		cdf: 0.9942026887948009,
		s:   "+0'fe841141f80c28/56",
	},
	{
		n:   10000,
		x:   2201,
		p:   0.2,
		cdf: 0.9999996810822256,
		s:   "+0'fffffaa64232e0/56",
	},
	{
		n:   10000,
		x:   2301,
		p:   0.2,
		cdf: 0.9999999999999356,
		s:   "+0'ffffffffffede0/56",
	},
	{
		n:   10000,
		x:   1801,
		p:   0.21,
		cdf: 4.223841148386452e-14,
		s:   "+0'00000000000be3/56",
	},
	{
		n:   10000,
		x:   1901,
		p:   0.21,
		cdf: 4.1090143969979605e-07,
		s:   "+0'000006e4cee92b/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.21,
		cdf: 0.007547057564875276,
		s:   "+0'01ee9a9d6c133f/56",
	},
	{
		n:   10000,
		x:   2101,
		p:   0.21,
		cdf: 0.5156332906287335,
		s:   "+0'84008b17fab1b0/56",
	},
	{
		n:   10000,
		x:   2201,
		p:   0.21,
		cdf: 0.9934292149198009,
		s:   "+0'fe516084f8b438/56",
	},
	{
		n:   10000,
		x:   2301,
		p:   0.21,
		cdf: 0.9999995057596741,
		s:   "+0'fffff7b54103c8/56",
	},
	{
		n:   10000,
		x:   2401,
		p:   0.21,
		cdf: 0.9999999999998367,
		s:   "+0'ffffffffffd208/56",
	},
	{
		n:   10000,
		x:   1901,
		p:   0.22,
		cdf: 1.1595144000953784e-13,
		s:   "+0'000000000020a3/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.22,
		cdf: 6.372749843946544e-07,
		s:   "+0'00000ab1134164/56",
	},
	{
		n:   10000,
		x:   2101,
		p:   0.22,
		cdf: 0.008457381241947058,
		s:   "+0'022a434fd80e2f/56",
	},
	{
		n:   10000,
		x:   2201,
		p:   0.22,
		cdf: 0.5153395943624124,
		s:   "+0'83ed4bb01ed6f0/56",
	},
	{
		n:   10000,
		x:   2301,
		p:   0.22,
		cdf: 0.9926406838068317,
		s:   "+0'fe1db329a12198/56",
	},
	{
		n:   10000,
		x:   2401,
		p:   0.22,
		cdf: 0.9999992656338182,
		s:   "+0'fffff3adebd818/56",
	},
	{
		n:   10000,
		x:   2501,
		p:   0.22,
		cdf: 0.9999999999996213,
		s:   "+0'ffffffffff9568/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.23,
		cdf: 2.880321419971924e-13,
		s:   "+0'00000000005112/56",
	},
	{
		n:   10000,
		x:   2101,
		p:   0.23,
		cdf: 9.471029252047307e-07,
		s:   "+0'00000fe3c6add0/56",
	},
	{
		n:   10000,
		x:   2201,
		p:   0.23,
		cdf: 0.009376501844397105,
		s:   "+0'02667f98c5c719/56",
	},
	{
		n:   10000,
		x:   2301,
		p:   0.23,
		cdf: 0.5150681413285996,
		s:   "+0'83db817637c470/56",
	},
	{
		n:   10000,
		x:   2401,
		p:   0.23,
		cdf: 0.9918432288176502,
		s:   "+0'fde9701687e5f0/56",
	},
	{
		n:   10000,
		x:   2501,
		p:   0.23,
		cdf: 0.9999989484240986,
		s:   "+0'ffffee5b8411a8/56",
	},
	{
		n:   10000,
		x:   2601,
		p:   0.23,
		cdf: 0.9999999999991879,
		s:   "+0'ffffffffff1b68/56",
	},
	{
		n:   10000,
		x:   2101,
		p:   0.24,
		cdf: 6.557338062314973e-13,
		s:   "+0'0000000000b892/56",
	},
	{
		n:   10000,
		x:   2201,
		p:   0.24,
		cdf: 1.3559745753542883e-06,
		s:   "+0'000016bfddd004/56",
	},
	{
		n:   10000,
		x:   2301,
		p:   0.24,
		cdf: 0.010297938886576274,
		s:   "+0'02a2e2bebbead0/56",
	},
	{
		n:   10000,
		x:   2401,
		p:   0.24,
		cdf: 0.514816669214655,
		s:   "+0'83cb0675b66c80/56",
	},
	{
		n:   10000,
		x:   2501,
		p:   0.24,
		cdf: 0.9910424867284411,
		s:   "+0'fdb4f5dd71f690/56",
	},
	{
		n:   10000,
		x:   2601,
		p:   0.24,
		cdf: 0.9999985425453698,
		s:   "+0'ffffe78c47afe8/56",
	},
	{
		n:   10000,
		x:   2701,
		p:   0.24,
		cdf: 0.9999999999983752,
		s:   "+0'fffffffffe36a8/56",
	},
	{
		n:   10000,
		x:   2201,
		p:   0.25,
		cdf: 1.382709738328955e-12,
		s:   "+0'00000000018532/56",
	},
	{
		n:   10000,
		x:   2301,
		p:   0.25,
		cdf: 1.8785036603737296e-06,
		s:   "+0'00001f841c9e0d/56",
	},
	{
		n:   10000,
		x:   2401,
		p:   0.25,
		cdf: 0.011215846459344928,
		s:   "+0'02df0aadbd6668/56",
	},
	{
		n:   10000,
		x:   2501,
		p:   0.25,
		cdf: 0.514583240710393,
		s:   "+0'83bbba2debbbc0/56",
	},
	{
		n:   10000,
		x:   2601,
		p:   0.25,
		cdf: 0.9902435998121082,
		s:   "+0'fd809ac4442d50/56",
	},
	{
		n:   10000,
		x:   2701,
		p:   0.25,
		cdf: 0.9999980377495189,
		s:   "+0'ffffdf1432c788/56",
	},
	{
		n:   10000,
		x:   2801,
		p:   0.25,
		cdf: 0.9999999999969432,
		s:   "+0'fffffffffca398/56",
	},
	{
		n:   10000,
		x:   2301,
		p:   0.26,
		cdf: 2.7245967542781786e-12,
		s:   "+0'0000000002fee7/56",
	},
	{
		n:   10000,
		x:   2401,
		p:   0.26,
		cdf: 2.5275211615857834e-06,
		s:   "+0'00002a679ee817/56",
	},
	{
		n:   10000,
		x:   2501,
		p:   0.26,
		cdf: 0.012124981163175837,
		s:   "+0'031a9f6d8f7a0f/56",
	},
	{
		n:   10000,
		x:   2601,
		p:   0.26,
		cdf: 0.5143661867709277,
		s:   "+0'83ad809e596fd0/56",
	},
	{
		n:   10000,
		x:   2701,
		p:   0.26,
		cdf: 0.9894512314679317,
		s:   "+0'fd4cad082445f8/56",
	},
	{
		n:   10000,
		x:   2801,
		p:   0.26,
		cdf: 0.9999974257079443,
		s:   "+0'ffffd4cf7ff390/56",
	},
	{
		n:   10000,
		x:   2901,
		p:   0.26,
		cdf: 0.9999999999945577,
		s:   "+0'fffffffffa0420/56",
	},
	{
		n:   10000,
		x:   2401,
		p:   0.27,
		cdf: 5.054652010838062e-12,
		s:   "+0'00000000058ec2/56",
	},
	{
		n:   10000,
		x:   2501,
		p:   0.27,
		cdf: 3.3133539829669314e-06,
		s:   "+0'00003796bf3b30/56",
	},
	{
		n:   10000,
		x:   2601,
		p:   0.27,
		cdf: 0.01302066188527879,
		s:   "+0'03555274f8342c/56",
	},
	{
		n:   10000,
		x:   2701,
		p:   0.27,
		cdf: 0.5141640617363261,
		s:   "+0'83a04185efa0b0/56",
	},
	{
		n:   10000,
		x:   2801,
		p:   0.27,
		cdf: 0.9886695904103495,
		s:   "+0'fd1973455cb6f0/56",
	},
	{
		n:   10000,
		x:   2901,
		p:   0.27,
		cdf: 0.9999967005037105,
		s:   "+0'ffffc8a4c576c0/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.27,
		cdf: 0.9999999999907789,
		s:   "+0'fffffffff5dc80/56",
	},
	{
		n:   10000,
		x:   2501,
		p:   0.28,
		cdf: 8.885020514295385e-12,
		s:   "+0'0000000009c4e9/56",
	},
	{
		n:   10000,
		x:   2601,
		p:   0.28,
		cdf: 4.243226393450848e-06,
		s:   "+0'0000473084c246/56",
	},
	{
		n:   10000,
		x:   2701,
		p:   0.28,
		cdf: 0.013898725619520595,
		s:   "+0'038eddebfdee5b/56",
	},
	{
		n:   10000,
		x:   2801,
		p:   0.28,
		cdf: 0.5139756072664875,
		s:   "+0'8393e7c826bbd8/56",
	},
	{
		n:   10000,
		x:   2901,
		p:   0.28,
		cdf: 0.9879024598338368,
		s:   "+0'fce72cf49fd100/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.28,
		cdf: 0.9999958590127097,
		s:   "+0'ffffba869852e8/56",
	},
	{
		n:   10000,
		x:   3101,
		p:   0.28,
		cdf: 0.9999999999850608,
		s:   "+0'ffffffffef9300/56",
	},
	{
		n:   10000,
		x:   2601,
		p:   0.29,
		cdf: 1.4878583391474648e-11,
		s:   "+0'00000000105bf2/56",
	},
	{
		n:   10000,
		x:   2701,
		p:   0.29,
		cdf: 5.320807254439786e-06,
		s:   "+0'00005944b17206/56",
	},
	{
		n:   10000,
		x:   2801,
		p:   0.29,
		cdf: 0.014755482185191978,
		s:   "+0'03c703e96c0f69/56",
	},
	{
		n:   10000,
		x:   2901,
		p:   0.29,
		cdf: 0.5137997233026825,
		s:   "+0'838860f0476410/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.29,
		cdf: 0.9871532290136191,
		s:   "+0'fcb612f2c11950/56",
	},
	{
		n:   10000,
		x:   3101,
		p:   0.29,
		cdf: 0.999994901162947,
		s:   "+0'ffffaa74a95f40/56",
	},
	{
		n:   10000,
		x:   3201,
		p:   0.29,
		cdf: 0.9999999999767619,
		s:   "+0'ffffffffe67310/56",
	},
	{
		n:   10000,
		x:   2701,
		p:   0.3,
		cdf: 2.3846524408680217e-11,
		s:   "+0'000000001a3833/56",
	},
	{
		n:   10000,
		x:   2801,
		p:   0.3,
		cdf: 6.545913517535346e-06,
		s:   "+0'00006dd27c06e5/56",
	},
	{
		n:   10000,
		x:   2901,
		p:   0.3,
		cdf: 0.015587669711257392,
		s:   "+0'03fd8db3a17592/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.3,
		cdf: 0.5136354443606586,
		s:   "+0'837d9ccb986fd0/56",
	},
	{
		n:   10000,
		x:   3101,
		p:   0.3,
		cdf: 0.9864249255872485,
		s:   "+0'fc86580b5b4158/56",
	},
	{
		n:   10000,
		x:   3201,
		p:   0.3,
		cdf: 0.9999938300694577,
		s:   "+0'ffff987c59a050/56",
	},
	{
		n:   10000,
		x:   3301,
		p:   0.3,
		cdf: 0.9999999999651701,
		s:   "+0'ffffffffd9b440/56",
	},
	{
		n:   10000,
		x:   2801,
		p:   0.31,
		cdf: 3.6727811020129784e-11,
		s:   "+0'000000002861f5/56",
	},
	{
		n:   10000,
		x:   2901,
		p:   0.31,
		cdf: 7.914370083904349e-06,
		s:   "+0'000084c7f5ef0b/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.31,
		cdf: 0.016392412071494767,
		s:   "+0'04324b09bfe777/56",
	},
	{
		n:   10000,
		x:   3101,
		p:   0.31,
		cdf: 0.5134819200883923,
		s:   "+0'83738d15ddef28/56",
	},
	{
		n:   10000,
		x:   3201,
		p:   0.31,
		cdf: 0.9857202473334898,
		s:   "+0'fc5829814d6820/56",
	},
	{
		n:   10000,
		x:   3301,
		p:   0.31,
		cdf: 0.9999926520490765,
		s:   "+0'ffff84b8ca84f0/56",
	},
	{
		n:   10000,
		x:   3401,
		p:   0.31,
		cdf: 0.9999999999495428,
		s:   "+0'ffffffffc88590/56",
	},
	{
		n:   10000,
		x:   2901,
		p:   0.32,
		cdf: 5.454860439460474e-11,
		s:   "+0'000000003bfa11/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.32,
		cdf: 9.41801811319093e-06,
		s:   "+0'00009e02146d12/56",
	},
	{
		n:   10000,
		x:   3101,
		p:   0.32,
		cdf: 0.017167178964586284,
		s:   "+0'0465117837ac96/56",
	},
	{
		n:   10000,
		x:   3201,
		p:   0.32,
		cdf: 0.5133383992556081,
		s:   "+0'836a25349575b8/56",
	},
	{
		n:   10000,
		x:   3301,
		p:   0.32,
		cdf: 0.9850415926762542,
		s:   "+0'fc2baf91be8828/56",
	},
	{
		n:   10000,
		x:   3401,
		p:   0.32,
		cdf: 0.9999913765242567,
		s:   "+0'ffff6f52742600/56",
	},
	{
		n:   10000,
		x:   3501,
		p:   0.32,
		cdf: 0.9999999999291588,
		s:   "+0'ffffffffb21bf8/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.33,
		cdf: 7.836194186324483e-11,
		s:   "+0'000000005628ec/56",
	},
	{
		n:   10000,
		x:   3101,
		p:   0.33,
		cdf: 1.1044858170565781e-05,
		s:   "+0'0000b94d4dfc55/56",
	},
	{
		n:   10000,
		x:   3201,
		p:   0.33,
		cdf: 0.017909749001306104,
		s:   "+0'0495bbba3d7c54/56",
	},
	{
		n:   10000,
		x:   3301,
		p:   0.33,
		cdf: 0.5132042164174095,
		s:   "+0'836159fdae9cd0/56",
	},
	{
		n:   10000,
		x:   3401,
		p:   0.33,
		cdf: 0.9843910894422945,
		s:   "+0'fc010defa0e320/56",
	},
	{
		n:   10000,
		x:   3501,
		p:   0.33,
		cdf: 0.99999001582854,
		s:   "+0'ffff587e4f62d0/56",
	},
	{
		n:   10000,
		x:   3601,
		p:   0.33,
		cdf: 0.9999999999033807,
		s:   "+0'ffffffff95c418/56",
	},
	{
		n:   10000,
		x:   3101,
		p:   0.34,
		cdf: 1.0917054532252848e-10,
		s:   "+0'000000007808c6/56",
	},
	{
		n:   10000,
		x:   3201,
		p:   0.34,
		cdf: 1.2779310951898393e-05,
		s:   "+0'0000d666b8fc90/56",
	},
	{
		n:   10000,
		x:   3301,
		p:   0.34,
		cdf: 0.01861817593863444,
		s:   "+0'04c42928c48206/56",
	},
	{
		n:   10000,
		x:   3401,
		p:   0.34,
		cdf: 0.513078780854479,
		s:   "+0'835921880aa240/56",
	},
	{
		n:   10000,
		x:   3501,
		p:   0.34,
		cdf: 0.983770621603032,
		s:   "+0'fbd864368cf500/56",
	},
	{
		n:   10000,
		x:   3601,
		p:   0.34,
		cdf: 0.9999885849283047,
		s:   "+0'ffff407ca3f068/56",
	},
	{
		n:   10000,
		x:   3701,
		p:   0.34,
		cdf: 0.9999999998717217,
		s:   "+0'ffffffff72f4e0/56",
	},
	{
		n:   10000,
		x:   3201,
		p:   0.35,
		cdf: 1.4783795201259577e-10,
		s:   "+0'00000000a28caf/56",
	},
	{
		n:   10000,
		x:   3301,
		p:   0.35,
		cdf: 1.4602576383632613e-05,
		s:   "+0'0000f4fd96877f/56",
	},
	{
		n:   10000,
		x:   3401,
		p:   0.35,
		cdf: 0.019290758047134367,
		s:   "+0'04f03d36ed700b/56",
	},
	{
		n:   10000,
		x:   3501,
		p:   0.35,
		cdf: 0.5129615672070887,
		s:   "+0'835173034264e8/56",
	},
	{
		n:   10000,
		x:   3601,
		p:   0.35,
		cdf: 0.9831818538843854,
		s:   "+0'fbb1ce54743ef8/56",
	},
	{
		n:   10000,
		x:   3701,
		p:   0.35,
		cdf: 0.9999871010763451,
		s:   "+0'ffff27978b74a8/56",
	},
	{
		n:   10000,
		x:   3801,
		p:   0.35,
		cdf: 0.9999999998339113,
		s:   "+0'ffffffff496230/56",
	},
	{
		n:   10000,
		x:   3301,
		p:   0.36,
		cdf: 1.9499502889588153e-10,
		s:   "+0'00000000d66638/56",
	},
	{
		n:   10000,
		x:   3401,
		p:   0.36,
		cdf: 1.649307128006011e-05,
		s:   "+0'000114b533a671/56",
	},
	{
		n:   10000,
		x:   3501,
		p:   0.36,
		cdf: 0.01992601051321867,
		s:   "+0'0519defb7e7883/56",
	},
	{
		n:   10000,
		x:   3601,
		p:   0.36,
		cdf: 0.5128521076343482,
		s:   "+0'834a4695f96500/56",
	},
	{
		n:   10000,
		x:   3701,
		p:   0.36,
		cdf: 0.98262625422101,
		s:   "+0'fb8d64ea11f9f0/56",
	},
	{
		n:   10000,
		x:   3801,
		p:   0.36,
		cdf: 0.9999855834129009,
		s:   "+0'ffff0e213ada08/56",
	},
	{
		n:   10000,
		x:   3901,
		p:   0.36,
		cdf: 0.9999999997899536,
		s:   "+0'ffffffff190d30/56",
	},
	{
		n:   10000,
		x:   3401,
		p:   0.37,
		cdf: 2.509500860574951e-10,
		s:   "+0'0000000113ec2b/56",
	},
	{
		n:   10000,
		x:   3501,
		p:   0.37,
		cdf: 1.8426926128831942e-05,
		s:   "+0'000135270b8af6/56",
	},
	{
		n:   10000,
		x:   3601,
		p:   0.37,
		cdf: 0.020522640711382774,
		s:   "+0'0540f8c6aed5a0/56",
	},
	{
		n:   10000,
		x:   3701,
		p:   0.37,
		cdf: 0.5127499851712165,
		s:   "+0'8343954155b960/56",
	},
	{
		n:   10000,
		x:   3801,
		p:   0.37,
		cdf: 0.982105114098707,
		s:   "+0'fb6b3da249cdb8/56",
	},
	{
		n:   10000,
		x:   3901,
		p:   0.37,
		cdf: 0.9999840525292001,
		s:   "+0'fffef472226c00/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.37,
		cdf: 0.9999999997401731,
		s:   "+0'fffffffee25138/56",
	},
	{
		n:   10000,
		x:   3501,
		p:   0.38,
		cdf: 3.156111170217819e-10,
		s:   "+0'000000015b04a1/56",
	},
	{
		n:   10000,
		x:   3601,
		p:   0.38,
		cdf: 2.0378522637633343e-05,
		s:   "+0'000155e51698d3/56",
	},
	{
		n:   10000,
		x:   3701,
		p:   0.38,
		cdf: 0.021079526159227868,
		s:   "+0'056577c3781454/56",
	},
	{
		n:   10000,
		x:   3801,
		p:   0.38,
		cdf: 0.5126548280967078,
		s:   "+0'833d58c8cfd630/56",
	},
	{
		n:   10000,
		x:   3901,
		p:   0.38,
		cdf: 0.9816195668673753,
		s:   "+0'fb4b6b80cf1eb0/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.38,
		cdf: 0.9999825300076072,
		s:   "+0'fffedae6f43a18/56",
	},
	{
		n:   10000,
		x:   4101,
		p:   0.38,
		cdf: 0.999999999685241,
		s:   "+0'fffffffea5eb38/56",
	},
	{
		n:   10000,
		x:   3601,
		p:   0.39,
		cdf: 3.8842813926838486e-10,
		s:   "+0'00000001ab14cd/56",
	},
	{
		n:   10000,
		x:   3701,
		p:   0.39,
		cdf: 2.232105518858175e-05,
		s:   "+0'0001767c33b95d/56",
	},
	{
		n:   10000,
		x:   3801,
		p:   0.39,
		cdf: 0.02159569495292882,
		s:   "+0'05874ba38ea612/56",
	},
	{
		n:   10000,
		x:   3901,
		p:   0.39,
		cdf: 0.5125663051255156,
		s:   "+0'83378b9d8baf80/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.39,
		cdf: 0.9811706041335551,
		s:   "+0'fb2dff288cd8b0/56",
	},
	{
		n:   10000,
		x:   4101,
		p:   0.39,
		cdf: 0.9999810379512222,
		s:   "+0'fffec1de9ee1e0/56",
	},
	{
		n:   10000,
		x:   4201,
		p:   0.39,
		cdf: 0.9999999996261805,
		s:   "+0'fffffffe64fb28/56",
	},
	{
		n:   10000,
		x:   3701,
		p:   0.4,
		cdf: 4.683622587188514e-10,
		s:   "+0'0000000202f841/56",
	},
	{
		n:   10000,
		x:   3801,
		p:   0.4,
		cdf: 2.4227101093919894e-05,
		s:   "+0'000196769b5c2d/56",
	},
	{
		n:   10000,
		x:   3901,
		p:   0.4,
		cdf: 0.022070308482616756,
		s:   "+0'05a66655253d8b/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.4,
		cdf: 0.5124841213347535,
		s:   "+0'833228ccda20c0/56",
	},
	{
		n:   10000,
		x:   4101,
		p:   0.4,
		cdf: 0.9807590903484581,
		s:   "+0'fb13071a4d2468/56",
	},
	{
		n:   10000,
		x:   4201,
		p:   0.4,
		cdf: 0.9999795985144188,
		s:   "+0'fffea9b8496128/56",
	},
	{
		n:   10000,
		x:   4301,
		p:   0.4,
		cdf: 0.9999999995643464,
		s:   "+0'fffffffe20fe68/56",
	},
	{
		n:   10000,
		x:   3801,
		p:   0.41,
		cdf: 5.538860870144598e-10,
		s:   "+0'00000002610112/56",
	},
	{
		n:   10000,
		x:   3901,
		p:   0.41,
		cdf: 2.6069186389149978e-05,
		s:   "+0'0001b55e4d8fc0/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.41,
		cdf: 0.022502646231211034,
		s:   "+0'05c2bbc1a2f241/56",
	},
	{
		n:   10000,
		x:   4101,
		p:   0.41,
		cdf: 0.5124080146518734,
		s:   "+0'832d2bf1235988/56",
	},
	{
		n:   10000,
		x:   4201,
		p:   0.41,
		cdf: 0.9803857757158527,
		s:   "+0'fafa8fec29c610/56",
	},
	{
		n:   10000,
		x:   4301,
		p:   0.41,
		cdf: 0.9999782334443907,
		s:   "+0'fffe92d15b02d0/56",
	},
	{
		n:   10000,
		x:   4401,
		p:   0.41,
		cdf: 0.9999999995013811,
		s:   "+0'fffffffddbc340/56",
	},
	{
		n:   10000,
		x:   3901,
		p:   0.42,
		cdf: 6.430177033487483e-10,
		s:   "+0'00000002c30164/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.42,
		cdf: 2.7820335719216852e-05,
		s:   "+0'0001d2bf6e9ca6/56",
	},
	{
		n:   10000,
		x:   4101,
		p:   0.42,
		cdf: 0.022892092477980918,
		s:   "+0'05dc419487a9dc/56",
	},
	{
		n:   10000,
		x:   4201,
		p:   0.42,
		cdf: 0.5123377528898034,
		s:   "+0'832891252af740/56",
	},
	{
		n:   10000,
		x:   4301,
		p:   0.42,
		cdf: 0.9800513075361021,
		s:   "+0'fae4a47a450450/56",
	},
	{
		n:   10000,
		x:   4401,
		p:   0.42,
		cdf: 0.9999769636424128,
		s:   "+0'fffe7d8398f8a8/56",
	},
	{
		n:   10000,
		x:   4501,
		p:   0.42,
		cdf: 0.9999999994391484,
		s:   "+0'fffffffd975650/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.43,
		cdf: 7.333877041739277e-10,
		s:   "+0'00000003265e49/56",
	},
	{
		n:   10000,
		x:   4101,
		p:   0.43,
		cdf: 2.9454596612355578e-05,
		s:   "+0'0001ee2a87777c/56",
	},
	{
		n:   10000,
		x:   4201,
		p:   0.43,
		cdf: 0.02323812473614726,
		s:   "+0'05f2ef09c31a5e/56",
	},
	{
		n:   10000,
		x:   4301,
		p:   0.43,
		cdf: 0.5122731311853785,
		s:   "+0'832454f90f8b40/56",
	},
	{
		n:   10000,
		x:   4401,
		p:   0.43,
		cdf: 0.9797562401010522,
		s:   "+0'fad14e11493390/56",
	},
	{
		n:   10000,
		x:   4501,
		p:   0.43,
		cdf: 0.9999758087522134,
		s:   "+0'fffe6a2361c4a8/56",
	},
	{
		n:   10000,
		x:   4601,
		p:   0.43,
		cdf: 0.9999999993796473,
		s:   "+0'fffffffd55ea40/56",
	},
	{
		n:   10000,
		x:   4101,
		p:   0.44,
		cdf: 8.223362590923274e-10,
		s:   "+0'00000003882b14/56",
	},
	{
		n:   10000,
		x:   4201,
		p:   0.44,
		cdf: 3.094753002330417e-05,
		s:   "+0'00020736a11c8f/56",
	},
	{
		n:   10000,
		x:   4301,
		p:   0.44,
		cdf: 0.02354030377725731,
		s:   "+0'0606bcc2dc7944/56",
	},
	{
		n:   10000,
		x:   4401,
		p:   0.44,
		cdf: 0.5122139698455042,
		s:   "+0'832074690a6ff0/56",
	},
	{
		n:   10000,
		x:   4501,
		p:   0.44,
		cdf: 0.97950104324332,
		s:   "+0'fac0949320c0c0/56",
	},
	{
		n:   10000,
		x:   4601,
		p:   0.44,
		cdf: 0.9999747867816691,
		s:   "+0'fffe58fe0d4580/56",
	},
	{
		n:   10000,
		x:   4701,
		p:   0.44,
		cdf: 0.9999999993249133,
		s:   "+0'fffffffd19bc00/56",
	},
	{
		n:   10000,
		x:   4201,
		p:   0.45,
		cdf: 9.070347069226659e-10,
		s:   "+0'00000003e54b92/56",
	},
	{
		n:   10000,
		x:   4301,
		p:   0.45,
		cdf: 3.2276660470983915e-05,
		s:   "+0'00021d83337e5e/56",
	},
	{
		n:   10000,
		x:   4401,
		p:   0.45,
		cdf: 0.02379826510613721,
		s:   "+0'0617a4a25517fd/56",
	},
	{
		n:   10000,
		x:   4501,
		p:   0.45,
		cdf: 0.512160112520738,
		s:   "+0'831cecd5979db0/56",
	},
	{
		n:   10000,
		x:   4601,
		p:   0.45,
		cdf: 0.9792861096322543,
		s:   "+0'fab27e964c2fb0/56",
	},
	{
		n:   10000,
		x:   4701,
		p:   0.45,
		cdf: 0.9999739137629492,
		s:   "+0'fffe4a587709a8/56",
	},
	{
		n:   10000,
		x:   4801,
		p:   0.45,
		cdf: 0.9999999992769116,
		s:   "+0'fffffffce4f4b8/56",
	},
	{
		n:   10000,
		x:   4301,
		p:   0.46,
		cdf: 9.84624267840876e-10,
		s:   "+0'000000043a9b17/56",
	},
	{
		n:   10000,
		x:   4401,
		p:   0.46,
		cdf: 3.34218803486958e-05,
		s:   "+0'000230b9e210c6/56",
	},
	{
		n:   10000,
		x:   4501,
		p:   0.46,
		cdf: 0.024011711770412,
		s:   "+0'0625a1acc65fb3/56",
	},
	{
		n:   10000,
		x:   4601,
		p:   0.46,
		cdf: 0.5121114246650746,
		s:   "+0'8319bbfcd82ec0/56",
	},
	{
		n:   10000,
		x:   4701,
		p:   0.46,
		cdf: 0.9791117608979572,
		s:   "+0'faa711802f89e0/56",
	},
	{
		n:   10000,
		x:   4801,
		p:   0.46,
		cdf: 0.9999732034552969,
		s:   "+0'fffe3e6db783c8/56",
	},
	{
		n:   10000,
		x:   4901,
		p:   0.46,
		cdf: 0.9999999992374312,
		s:   "+0'fffffffcb98bf8/56",
	},
	{
		n:   10000,
		x:   4401,
		p:   0.47,
		cdf: 1.0523630562138254e-09,
		s:   "+0'000000048515dd/56",
	},
	{
		n:   10000,
		x:   4501,
		p:   0.47,
		cdf: 3.436580408436118e-05,
		s:   "+0'00024090013049/56",
	},
	{
		n:   10000,
		x:   4601,
		p:   0.47,
		cdf: 0.02418040840533094,
		s:   "+0'0630afef4e08ad/56",
	},
	{
		n:   10000,
		x:   4701,
		p:   0.47,
		cdf: 0.5120677922053533,
		s:   "+0'8316dff4db5ed8/56",
	},
	{
		n:   10000,
		x:   4801,
		p:   0.47,
		cdf: 0.9789782526609708,
		s:   "+0'fa9e519aac9178/56",
	},
	{
		n:   10000,
		x:   4901,
		p:   0.47,
		cdf: 0.9999726670937753,
		s:   "+0'fffe356e0fc8f8/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.47,
		cdf: 0.9999999992079847,
		s:   "+0'fffffffc992b80/56",
	},
	{
		n:   10000,
		x:   4501,
		p:   0.48,
		cdf: 1.1077718462741698e-09,
		s:   "+0'00000004c2020d/56",
	},
	{
		n:   10000,
		x:   4601,
		p:   0.48,
		cdf: 3.509406876293367e-05,
		s:   "+0'00024cc7e0abba/56",
	},
	{
		n:   10000,
		x:   4701,
		p:   0.48,
		cdf: 0.024304176431698455,
		s:   "+0'0638cc6aee2bff/56",
	},
	{
		n:   10000,
		x:   4801,
		p:   0.48,
		cdf: 0.5120291205547989,
		s:   "+0'8314572761ebb0/56",
	},
	{
		n:   10000,
		x:   4901,
		p:   0.48,
		cdf: 0.978885778518274,
		s:   "+0'fa984225416788/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.48,
		cdf: 0.9999723131865894,
		s:   "+0'fffe2f7e0ab810/56",
	},
	{
		n:   10000,
		x:   5101,
		p:   0.48,
		cdf: 0.999999999189723,
		s:   "+0'fffffffc851750/56",
	},
	{
		n:   10000,
		x:   4601,
		p:   0.49,
		cdf: 1.148769007973377e-09,
		s:   "+0'00000004ef15ba/56",
	},
	{
		n:   10000,
		x:   4701,
		p:   0.49,
		cdf: 3.559557861791754e-05,
		s:   "+0'00025531d8967b/56",
	},
	{
		n:   10000,
		x:   4801,
		p:   0.49,
		cdf: 0.024382890330102654,
		s:   "+0'063df5048cbc1f/56",
	},
	{
		n:   10000,
		x:   4901,
		p:   0.49,
		cdf: 0.5119953337563443,
		s:   "+0'8312204e302790/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.49,
		cdf: 0.9788344730359314,
		s:   "+0'fa94e561e30008/56",
	},
	{
		n:   10000,
		x:   5101,
		p:   0.49,
		cdf: 0.999972147362962,
		s:   "+0'fffe2cb5d5b668/56",
	},
	{
		n:   10000,
		x:   5201,
		p:   0.49,
		cdf: 0.9999999991833699,
		s:   "+0'fffffffc7e1b10/56",
	},
}

func Benchmark_Fixed_BinCDF_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tc := bincdfTestCases_10000[i%len(bincdfTestCases_10000)]
		bincdfResultFix = BinCDF64(tc.n, From(tc.p), tc.x)
	}

	bincdfResultFix.lo++
}

func Benchmark_Float_BinCDF_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tc := bincdfTestCases_10000[i%len(bincdfTestCases_10000)]
		bincdfResultFlt = bincdf_(tc.n, tc.p, tc.x)
	}

	bincdfResultFlt++
}

func TestFixed_BinCDF_10000(t *testing.T) {
	acc := accuracy{Epsilon: 1e-10}
	for i, tc := range bincdfTestCases_10000 {
		p := From(tc.p)
		got := BinCDF64(tc.n, p, tc.x)
		if ok := acc.update(got, tc.cdf); !ok {
			t.Errorf("%d: BinCDF(%v,%v,%v) => got %v|%v, want %v|%v", i, tc.n, tc.p, tc.x, got, got.Float(), From(tc.cdf), tc.cdf)
		}
	}
	t.Log(acc)
}
