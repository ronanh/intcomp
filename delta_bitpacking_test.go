package intcomp_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/ronanh/intcomp"
)

func TestCompressDeltaBinPackInt32(t *testing.T) {
	var testInput = []int32{0, 202, 320, math.MaxInt32, math.MinInt32, 529, 638, 838, 949, 1151, 1257, 1454, 1561, 1759, 1994, 2105, 2307, 2413, 2610, 2716, 2913, 3146, 3253, 3360, 3558, 3665, 3863, 4061, 4168, 4366, 4472, 4578, 4775, 4972, 5079, 5277, 5386,
		5586, 5694, 5801, 6000, 6106, 6304, 6501, 6627, 6844, 6950, 7059, 7256, 7456, 7563, 7672, 7870, 8070, 8309, 8418, 8618, 8857, 8961, 9156, 9269, 9473, 9579, 9776, 9883, 10081, 10187, 10384, 10491, 10689, 10826, 11054, 11180, 11287, 11504,
		11610, 11808, 11915, 12022, 12219, 12417, 12615, 12721, 12918, 13024, 13221, 13328, 13526, 13632, 13735, 13842, 14039, 14233, 14431, 14537, 14734, 14841, 15039, 15145, 15252, 15360, 15557, 15755, 15954, 16093, 16323, 16430, 16628, 16734,
		16840, 17037, 17234, 17467, 17700, 17806, 18003, 18236, 18342, 18539, 18645, 18842, 18948, 19061, 19258, 19462, 19568, 19765, 19871, 20068, 20175, 20373, 20489, 20696, 20802, 20909, 21106, 21304, 21411, 21609, 21715, 21912, 22018, 22215,
		22322, 22428, 22626, 22823, 22929, 23126, 23230, 23336, 23442, 23637, 23834, 24031, 24137, 24243, 24440, 24637, 24744, 24850, 25048, 25245, 25352, 25459, 25657, 25764, 25962, 26160, 26266, 26372, 26569, 26676, 26873, 27071, 27187, 27394,
		27500, 27606, 27712, 27819, 28016, 28213, 28410, 28608, 28726, 28832, 28939, 29046, 29243, 29452, 29650, 29756, 29954, 30062, 30168, 30365, 30562, 30761, 30867, 30973, 31079, 31276, 31383, 31580, 31777, 31900, 32098, 32204, 32310, 32524,
		32630, 32827, 33024, 33221, 33327, 33524, 33631, 33737, 33935, 34132, 34238, 34435, 34542, 34740, 34846, 34949, 35146, 35340, 35447, 35645, 35751, 35948, 36054, 36251, 36358, 36465, 36571, 36678, 36876, 37074, 37271, 37469, 37576, 37774,
		37998, 38104, 38210, 38407, 38604, 38725, 38937, 39044, 39148, 39346, 39452, 39647, 39753, 39950, 40147, 40251, 40357, 40463, 40658, 40765, 40883, 41080, 41277, 41383, 41581, 41790, 41896, 42093, 42290, 42396, 42502, 42608, 42805, 42911,
		43108, 43305, 43502, 43608, 43714, 43911, 44108, 44214, 44411, 44517, 44714, 44820, 44926, 45032, 45141, 45338, 45444, 45641, 45747, 45944, 46051, 46248, 46354, 46554, 46751, 46949, 47146, 47252, 47449, 47562, 47766, 47875, 48075, 48314,
		48423, 48623, 48862, 48968, 49165, 49271, 49468, 49572, 49767, 49880, 49986, 50190, 50387, 50493, 50600, 50797, 50912, 51110, 51316, 51422, 51531, 51731, 51928, 52034, 52231, 52337, 52534, 52640, 52746, 52943, 53140, 53246, 53443, 53676,
		53782, 53979, 54093, 54298, 54404, 54601, 54834, 54959, 55072, 55288, 55492, 55598, 55795, 55908, 56112, 56225, 56343, 56547, 56756, 56865, 57065, 57178, 57382, 57491, 57691, 57798, 57996, 58110, 58315, 58422, 58620, 58855, 58962, 59160,
		59266, 59463, 59696, 59802, 59999, 60232, 60339, 60537, 60644, 60842, 60959, 61167, 61280, 61394, 61505, 61707, 61911, 62116, 62238, 62352, 62565, 62770, 62883, 62989, 63193, 63390, 63501, 63703, 63812, 64012, 64121, 64321, 64560, 64669,
		64869, 65108, 65219, 65421, 65527, 65724, 65835, 66037, 66145, 66344, 66581, 66687, 66794, 66905, 67018, 67124, 67326, 67530, 67727, 67925, 68122, 68236, 68441, 68547, 68744, 68850, 69047, 69160, 69364, 69490, 69707, 69980, 70091, 70202,
		70313, 70515, 70717, 70919, 71025, 71222, 71455, 71568, 71772, 72019, 72128, 72328, 72435, 72633, 72857, 72968, 73170, 73277, 73475, 73710, 73836, 73942, 74159, 74356, 74463, 74570, 74768, 74966, 75072, 75269, 75382, 75586, 75692, 75889,
		75996, 76111, 76309, 76515, 76622, 76820, 76927, 77125, 77236, 77438, 77559, 77666, 77878, 77990, 78188, 78391, 78497}

	resinput, input2 := intcomp.CompressDeltaBinPackInt32(testInput, make([]uint32, 0, 100))
	resinput2, resoutput2 := intcomp.UncompressDeltaBinPackInt32(input2, nil)

	if len(resoutput2) != len(testInput)-len(resinput) {
		t.Fatalf("expected same len")
	}
	if len(resinput2) != 0 {
		t.Fatalf("resinput2 should be empty")
	}
	for i, v := range testInput[:len(testInput)-len(resinput)] {
		if resoutput2[i] != v {
			t.Fatalf("resoutput2 != input")
		}
	}
}

func TestCompressDeltaBinPackInt64_trailing(t *testing.T) {
	var testInput = []int64{0, 202, 320, 529, 638, 838, 949, 1151, 1257, 1454, 1561, 1759, 1994, 2105, 2307, 2413, 2610, 2716, 2913, 3146, 3253, 3360, 3558, 3665, 3863, 4061, 4168, 4366, 4472, 4578, 4775, 4972, 5079, 5277, 5386,
		5586, 5694, 5801, 6000, 6106, 6304, 6501, 6627, 6844, 6950, 7059, 7256, 7456, 7563, 7672, 7870, 8070, 8309, 8418, 8618, 8857, 8961, 9156, 9269, 9473, 9579, 9776, 9883, 10081, 10187, 10384, 10491, 10689, 10826, 11054, 11180, 11287, 11504,
		11610, 11808, 11915, 12022, 12219, 12417, 12615, 12721, 12918, 13024, 13221, 13328, 13526, 13632, 13735, 13842, 14039, 14233, 14431, 14537, 14734, 14841, 15039, 15145, 15252, 15360, 15557, 15755, 15954, 16093, 16323, 16430, 16628, 16734,
		16840, 17037, 17234, 17467, 17700, 17806, 18003, 18236, 18342, 18539, 18645, 18842, 18948, 19061, 19258, 19462, 19568, 19765, 19871, 20068, 20175, 20373, 20489, 20696, 20802, 20909, 21106, 21304, 21411, 21609, 21715, 21912, 22018, 22215,
		22322, 22428, 22626, 22823, 22929, 23126, 23230, 23336, 23442, 23637, 23834, 24031, 24137, 24243, 24440, 24637, 24744, 24850, 25048, 25245, 25352, 25459, 25657, 25764, 25962, 26160, 26266, 26372, 26569, 26676, 26873, 27071, 27187, 27394,
		27500, 27606, 27712, 27819, 28016, 28213, 28410, 28608, 28726, 28832, 28939, 29046, 29243, 29452, 29650, 29756, 29954, 30062, 30168, 30365, 30562, 30761, 30867, 30973, 31079, 31276, 31383, 31580, 31777, 31900, 32098, 32204, 32310, 32524,
		32630, 32827, 33024, 33221, 33327, 33524, 33631, 33737, 33935, 34132, 34238, 34435, 34542, 34740, 34846, 34949, 35146, 35340, 35447, 35645, 35751, 35948, 36054, 36251, 36358, 36465, 36571, 36678, 36876, 37074, 37271, 37469, 37576, 37774,
		37998, 38104, 38210, 38407, 38604, 38725, 38937, 39044, 39148, 39346, 39452, 39647, 39753, 39950, 40147, 40251, 40357, 40463, 40658, 40765, 40883, 41080, 41277, 41383, 41581, 41790, 41896, 42093, 42290, 42396, 42502, 42608, 42805, 42911,
		43108, 43305, 43502, 43608, 43714, 43911, 44108, 44214, 44411, 44517, 44714, 44820, 44926, 45032, 45141, 45338, 45444, 45641, 45747, 45944, 46051, 46248, 46354, 46554, 46751, 46949, 47146, 47252, 47449, 47562, 47766, 47875, 48075, 48314,
		48423, 48623, 48862, 48968, 49165, 49271, 49468, 49572, 49767, 49880, 49986, 50190, 50387, 50493, 50600, 50797, 50912, 51110, 51316, 51422, 51531, 51731, 51928, 52034, 52231, 52337, 52534, 52640, 52746, 52943, 53140, 53246, 53443, 53676,
		53782, 53979, 54093, 54298, 54404, 54601, 54834, 54959, 55072, 55288, 55492, 55598, 55795, 55908, 56112, 56225, 56343, 56547, 56756, 56865, 57065, 57178, 57382, 57491, 57691, 57798, 57996, 58110, 58315, 58422, 58620, 58855, 58962, 59160,
		59266, 59463, 59696, 59802, 59999, 60232, 60339, 60537, 60644, 60842, 60959, 61167, 61280, 61394, 61505, 61707, 61911, 62116, 62238, 62352, 62565, 62770, 62883, 62989, 63193, 63390, 63501, 63703, 63812, 64012, 64121, 64321, 64560, 64669,
		64869, 65108, 65219, 65421, 65527, 65724, 65835, 66037, 66145, 66344, 66581, 66687, 66794, 66905, 67018, 67124, 67326, 67530, 67727, 67925, 68122, 68236, 68441, 68547, 68744, 68850, 69047, 69160, 69364, 69490, 69707, 69980, 70091, 70202,
		70313, 70515, 70717, 70919, 71025, 71222, 71455, 71568, 71772, 72019, 72128, 72328, 72435, 72633, 72857, 72968, 73170, 73277, 73475, 73710, 73836, 73942, 74159, 74356, 74463, 74570, 74768, 74966, 75072, 75269, 75382, 75586, 75692, 75889,
		75996, 76111, 76309, 76515, 76622, 76820, 76927, 77125, 77236, 77438, 77559, 77666, 77878, 77990, 78188, 78391, 78497}

	for i := range testInput {
		testInput[i] *= 8
	}
	resinput, input2 := intcomp.CompressDeltaBinPackInt64(testInput, make([]uint64, 0, 100))
	resinput2, resoutput2 := intcomp.UncompressDeltaBinPackInt64(input2, nil)

	if len(resoutput2) != len(testInput)-len(resinput) {
		t.Fatalf("expected same len")
	}
	if len(resinput2) != 0 {
		t.Fatalf("resinput2 should be empty")
	}
	for i, v := range testInput[:len(testInput)-len(resinput)] {
		if resoutput2[i] != v {
			t.Fatalf("resoutput2 != input")
		}
	}
}

func TestCompressDeltaBinPackUint32(t *testing.T) {
	var testInput = []uint32{0, 202, math.MaxUint32, 0, math.MaxUint32, 320, 529, 638, 838, 949, 1151, 1257, 1454, 1561, 1759, 1994, 2105, 2307, 2413, 2610, 2716, 2913, 3146, 3253, 3360, 3558, 3665, 3863, 4061, 4168, 4366, 4472, 4578, 4775, 4972, 5079, 5277, 5386,
		5586, 5694, 5801, 6000, 6106, 6304, 6501, 6627, 6844, 6950, 7059, 7256, 7456, 7563, 7672, 7870, 8070, 8309, 8418, 8618, 8857, 8961, 9156, 9269, 9473, 9579, 9776, 9883, 10081, 10187, 10384, 10491, 10689, 10826, 11054, 11180, 11287, 11504,
		11610, 11808, 11915, 12022, 12219, 12417, 12615, 12721, 12918, 13024, 13221, 13328, 13526, 13632, 13735, 13842, 14039, 14233, 14431, 14537, 14734, 14841, 15039, 15145, 15252, 15360, 15557, 15755, 15954, 16093, 16323, 16430, 16628, 16734,
		16840, 17037, 17234, 17467, 17700, 17806, 18003, 18236, 18342, 18539, 18645, 18842, 18948, 19061, 19258, 19462, 19568, 19765, 19871, 20068, 20175, 20373, 20489, 20696, 20802, 20909, 21106, 21304, 21411, 21609, 21715, 21912, 22018, 22215,
		22322, 22428, 22626, 22823, 22929, 23126, 23230, 23336, 23442, 23637, 23834, 24031, 24137, 24243, 24440, 24637, 24744, 24850, 25048, 25245, 25352, 25459, 25657, 25764, 25962, 26160, 26266, 26372, 26569, 26676, 26873, 27071, 27187, 27394,
		27500, 27606, 27712, 27819, 28016, 28213, 28410, 28608, 28726, 28832, 28939, 29046, 29243, 29452, 29650, 29756, 29954, 30062, 30168, 30365, 30562, 30761, 30867, 30973, 31079, 31276, 31383, 31580, 31777, 31900, 32098, 32204, 32310, 32524,
		32630, 32827, 33024, 33221, 33327, 33524, 33631, 33737, 33935, 34132, 34238, 34435, 34542, 34740, 34846, 34949, 35146, 35340, 35447, 35645, 35751, 35948, 36054, 36251, 36358, 36465, 36571, 36678, 36876, 37074, 37271, 37469, 37576, 37774,
		37998, 38104, 38210, 38407, 38604, 38725, 38937, 39044, 39148, 39346, 39452, 39647, 39753, 39950, 40147, 40251, 40357, 40463, 40658, 40765, 40883, 41080, 41277, 41383, 41581, 41790, 41896, 42093, 42290, 42396, 42502, 42608, 42805, 42911,
		43108, 43305, 43502, 43608, 43714, 43911, 44108, 44214, 44411, 44517, 44714, 44820, 44926, 45032, 45141, 45338, 45444, 45641, 45747, 45944, 46051, 46248, 46354, 46554, 46751, 46949, 47146, 47252, 47449, 47562, 47766, 47875, 48075, 48314,
		48423, 48623, 48862, 48968, 49165, 49271, 49468, 49572, 49767, 49880, 49986, 50190, 50387, 50493, 50600, 50797, 50912, 51110, 51316, 51422, 51531, 51731, 51928, 52034, 52231, 52337, 52534, 52640, 52746, 52943, 53140, 53246, 53443, 53676,
		53782, 53979, 54093, 54298, 54404, 54601, 54834, 54959, 55072, 55288, 55492, 55598, 55795, 55908, 56112, 56225, 56343, 56547, 56756, 56865, 57065, 57178, 57382, 57491, 57691, 57798, 57996, 58110, 58315, 58422, 58620, 58855, 58962, 59160,
		59266, 59463, 59696, 59802, 59999, 60232, 60339, 60537, 60644, 60842, 60959, 61167, 61280, 61394, 61505, 61707, 61911, 62116, 62238, 62352, 62565, 62770, 62883, 62989, 63193, 63390, 63501, 63703, 63812, 64012, 64121, 64321, 64560, 64669,
		64869, 65108, 65219, 65421, 65527, 65724, 65835, 66037, 66145, 66344, 66581, 66687, 66794, 66905, 67018, 67124, 67326, 67530, 67727, 67925, 68122, 68236, 68441, 68547, 68744, 68850, 69047, 69160, 69364, 69490, 69707, 69980, 70091, 70202,
		70313, 70515, 70717, 70919, 71025, 71222, 71455, 71568, 71772, 72019, 72128, 72328, 72435, 72633, 72857, 72968, 73170, 73277, 73475, 73710, 73836, 73942, 74159, 74356, 74463, 74570, 74768, 74966, 75072, 75269, 75382, 75586, 75692, 75889,
		75996, 76111, 76309, 76515, 76622, 76820, 76927, 77125, 77236, 77438, 77559, 77666, 77878, 77990, 78188, 78391, 78497}

	resinput, input2 := intcomp.CompressDeltaBinPackUint32(testInput, make([]uint32, 0, 100))

	resinput2, resoutput2 := intcomp.UncompressDeltaBinPackUint32(input2, nil)

	//output2 = output2[:outpos2]
	if len(resoutput2) != len(testInput)-len(resinput) {
		t.Fatalf("expected same len")
	}
	if len(resinput2) != 0 {
		t.Fatalf("resinput2 should be empty")
	}
	for i, v := range testInput[:len(testInput)-len(resinput)] {
		if resoutput2[i] != v {
			t.Fatalf("resoutput2 != input")
		}
	}

}

func TestCompressDeltaBinPackInt64Constant(t *testing.T) {
	var testInput []int64
	for i := 0; i < 256; i++ {
		testInput = append(testInput, 123456789)
	}

	resinput, input2 := intcomp.CompressDeltaBinPackInt64(testInput, nil)

	if len(resinput) != 0 {
		t.Fatalf("expected all input consumed")
	}
	if len(input2) != 3 {
		t.Fatalf("input2 only headers expected")
	}
}

func TestCompressDeltaBinPackInt64OneBit(t *testing.T) {
	var testInput []int64
	for i := 0; i < 256; i++ {
		testInput = append(testInput, 123456789+int64(i))
	}

	resinput, input2 := intcomp.CompressDeltaBinPackInt64(testInput, nil)

	if len(resinput) != 0 {
		t.Fatalf("expected all input consumed")
	}
	if len(input2) != 3+4 {
		t.Fatalf("input2 only headers and 4 words expected")
	}
}

func TestCompressDeltaBinPackInt64TwoBit(t *testing.T) {
	var testInput []int64
	for i := 0; i < 256; i++ {
		testInput = append(testInput, 123456789+2*int64(i))
	}

	resinput, input2 := intcomp.CompressDeltaBinPackInt64(testInput, nil)

	if len(resinput) != 0 {
		t.Fatalf("expected all input consumed")
	}
	if len(input2) != 3+4 {
		t.Fatalf("input2 only headers and 4 words expected")
	}
}

func TestCompressDeltaBinPackInt64TwoBitNegativ(t *testing.T) {
	var testInput []int64
	for i := 0; i < 256; i++ {
		testInput = append(testInput, 123456789-int64(i))
	}

	resinput, input2 := intcomp.CompressDeltaBinPackInt64(testInput, nil)

	if len(resinput) != 0 {
		t.Fatalf("expected all input consumed")
	}
	if len(input2) != 3+4 {
		t.Fatalf("input2 only headers and 4 words expected")
	}
}
func TestCompressDeltaBinPackInt64(t *testing.T) {
	var testInput = []int64{0, 202, math.MaxInt64, math.MinInt64, 320, 529, 638, 838, 949, 1151, 1257, 1454, 1561, 1759, 1994, 2105, 2307, 2413, 2610, 2716, 2913, 3146, 3253, 3360, 3558, 3665, 3863, 4061, 4168, 4366, 4472, 4578, 4775, 4972, 5079, 5277, 5386,
		5586, 5694, 5801, 6000, 6106, 6304, 6501, 6627, 6844, 6950, 7059, 7256, 7456, 7563, 7672, 7870, 8070, 8309, 8418, 8618, 8857, 8961, 9156, 9269, 9473, 9579, 9776, 9883, 10081, 10187, 10384, 10491, 10689, 10826, 11054, 11180, 11287, 11504,
		11610, 11808, 11915, 12022, 12219, 12417, 12615, 12721, 12918, 13024, 13221, 13328, 13526, 13632, 13735, 13842, 14039, 14233, 14431, 14537, 14734, 14841, 15039, 15145, 15252, 15360, 15557, 15755, 15954, 16093, 16323, 16430, 16628, 16734,
		16840, 17037, 17234, 17467, 17700, 17806, 18003, 18236, 18342, 18539, 18645, 18842, 18948, 19061, 19258, 19462, 19568, 19765, 19871, 20068, 20175, 20373, 20489, 20696, 20802, 20909, 21106, 21304, 21411, 21609, 21715, 21912, 22018, 22215,
		22322, 22428, 22626, 22823, 22929, 23126, 23230, 23336, 23442, 23637, 23834, 24031, 24137, 24243, 24440, 24637, 24744, 24850, 25048, 25245, 25352, 25459, 25657, 25764, 25962, 26160, 26266, 26372, 26569, 26676, 26873, 27071, 27187, 27394,
		27500, 27606, 27712, 27819, 28016, 28213, 28410, 28608, 28726, 28832, 28939, 29046, 29243, 29452, 29650, 29756, 29954, 30062, 30168, 30365, 30562, 30761, 30867, 30973, 31079, 31276, 31383, 31580, 31777, 31900, 32098, 32204, 32310, 32524,
		32630, 32827, 33024, 33221, 33327, 33524, 33631, 33737, 33935, 34132, 34238, 34435, 34542, 34740, 34846, 34949, 35146, 35340, 35447, 35645, 35751, 35948, 36054, 36251, 36358, 36465, 36571, 36678, 36876, 37074, 37271, 37469, 37576, 37774,
		37998, 38104, 38210, 38407, 38604, 38725, 38937, 39044, 39148, 39346, 39452, 39647, 39753, 39950, 40147, 40251, 40357, 40463, 40658, 40765, 40883, 41080, 41277, 41383, 41581, 41790, 41896, 42093, 42290, 42396, 42502, 42608, 42805, 42911,
		43108, 43305, 43502, 43608, 43714, 43911, 44108, 44214, 44411, 44517, 44714, 44820, 44926, 45032, 45141, 45338, 45444, 45641, 45747, 45944, 46051, 46248, 46354, 46554, 46751, 46949, 47146, 47252, 47449, 47562, 47766, 47875, 48075, 48314,
		48423, 48623, 48862, 48968, 49165, 49271, 49468, 49572, 49767, 49880, 49986, 50190, 50387, 50493, 50600, 50797, 50912, 51110, 51316, 51422, 51531, 51731, 51928, 52034, 52231, 52337, 52534, 52640, 52746, 52943, 53140, 53246, 53443, 53676,
		53782, 53979, 54093, 54298, 54404, 54601, 54834, 54959, 55072, 55288, 55492, 55598, 55795, 55908, 56112, 56225, 56343, 56547, 56756, 56865, 57065, 57178, 57382, 57491, 57691, 57798, 57996, 58110, 58315, 58422, 58620, 58855, 58962, 59160,
		59266, 59463, 59696, 59802, 59999, 60232, 60339, 60537, 60644, 60842, 60959, 61167, 61280, 61394, 61505, 61707, 61911, 62116, 62238, 62352, 62565, 62770, 62883, 62989, 63193, 63390, 63501, 63703, 63812, 64012, 64121, 64321, 64560, 64669,
		64869, 65108, 65219, 65421, 65527, 65724, 65835, 66037, 66145, 66344, 66581, 66687, 66794, 66905, 67018, 67124, 67326, 67530, 67727, 67925, 68122, 68236, 68441, 68547, 68744, 68850, 69047, 69160, 69364, 69490, 69707, 69980, 70091, 70202,
		70313, 70515, 70717, 70919, 71025, 71222, 71455, 71568, 71772, 72019, 72128, 72328, 72435, 72633, 72857, 72968, 73170, 73277, 73475, 73710, 73836, 73942, 74159, 74356, 74463, 74570, 74768, 74966, 75072, 75269, 75382, 75586, 75692, 75889,
		75996, 76111, 76309, 76515, 76622, 76820, 76927, 77125, 77236, 77438, 77559, 77666, 77878, 77990, 78188, 78391, 78497}

	resinput, input2 := intcomp.CompressDeltaBinPackInt64(testInput, make([]uint64, 0, 100))
	resinput2, resoutput2 := intcomp.UncompressDeltaBinPackInt64(input2, nil)

	if len(resoutput2) != len(testInput)-len(resinput) {
		t.Fatalf("expected same len")
	}
	if len(resinput2) != 0 {
		t.Fatalf("resinput2 should be empty")
	}
	for i, v := range testInput[:len(testInput)-len(resinput)] {
		if resoutput2[i] != v {
			t.Fatalf("resoutput2(%d) != input(%d)", resoutput2[i], v)
		}
	}
}

func TestCompressDeltaBinPackFullInt32(t *testing.T) {
	rand.Seed(1) // nolint
	maxBits := 32
	for nBits := 0; nBits <= maxBits; nBits++ {
		for sign := 0; sign < 2; sign++ {
			var testInput []int32
			prev := int32(rand.Uint32())
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize32; i++ {
				var udiff uint32
				if nBits == maxBits {
					udiff = rand.Uint32()
				} else if nBits > 0 {
					udiff = rand.Uint32()
					udiff %= 1 << nBits
				}
				if sign == 1 {
					prev = int32(uint32(prev) + udiff)
				} else {
					// consider udiff as zigzag encoded diff
					prev += (-(int32(udiff) & 1)) ^ (int32(udiff) >> 1)
				}
				testInput = append(testInput, prev)
			}

			t.Run(fmt.Sprintf("nBits=%d, sign=%d", nBits, sign), func(t *testing.T) {
				resinput, input2 := intcomp.CompressDeltaBinPackInt32(testInput, nil)
				resinput2, resoutput2 := intcomp.UncompressDeltaBinPackInt32(input2, nil)

				if len(resinput) != 0 {
					t.Fatalf("expected all input consumed")
				}
				if len(resoutput2) != len(testInput) {
					t.Fatalf("expected same len")
				}
				if len(resinput2) != 0 {
					t.Fatalf("resinput2 should be empty")
				}
				for i, v := range testInput {
					if resoutput2[i] != v {
						t.Fatalf("resoutput2(%d) != input(%d)", resoutput2[i], v)
					}
				}
			})
		}
	}
}

func TestCompressDeltaBinPackFullUint32(t *testing.T) {
	rand.Seed(1) // nolint
	maxBits := 32
	for nBits := 0; nBits <= maxBits; nBits++ {
		for sign := 0; sign < 2; sign++ {
			var testInput []uint32
			prev := rand.Uint32()
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize32; i++ {
				var udiff uint32
				if nBits == maxBits {
					udiff = rand.Uint32()
				} else if nBits > 0 {
					udiff = rand.Uint32()
					udiff %= 1 << nBits
				}
				if sign == 1 {
					prev = uint32(prev) + udiff
				} else {
					// consider udiff as zigzag encoded diff
					prev += uint32((-(int32(udiff) & 1)) ^ (int32(udiff) >> 1))
				}
				testInput = append(testInput, prev)
			}

			t.Run(fmt.Sprintf("nBits=%d, sign=%d", nBits, sign), func(t *testing.T) {
				resinput, input2 := intcomp.CompressDeltaBinPackUint32(testInput, nil)
				resinput2, resoutput2 := intcomp.UncompressDeltaBinPackUint32(input2, nil)

				if len(resinput) != 0 {
					t.Fatalf("expected all input consumed")
				}
				if len(resoutput2) != len(testInput) {
					t.Fatalf("expected same len")
				}
				if len(resinput2) != 0 {
					t.Fatalf("resinput2 should be empty")
				}
				for i, v := range testInput {
					if resoutput2[i] != v {
						t.Fatalf("resoutput2(%d) != input(%d)", resoutput2[i], v)
					}
				}
			})
		}
	}
}

func TestCompressDeltaBinPackFullInt64(t *testing.T) {
	rand.Seed(1) // nolint
	maxBits := 64
	for nBits := 0; nBits <= maxBits; nBits++ {
		for ntz := 0; ntz < nBits; ntz++ {
			if nBits > maxBits {
				continue
			}
			if nBits > 0 && ntz >= nBits {
				continue
			}
			var testInput []int64
			prev := int64(rand.Uint64())
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize64; i++ {
				var udiff uint64
				if nBits-ntz == maxBits {
					udiff = rand.Uint64()
				} else if nBits-ntz > 0 {
					udiff = rand.Uint64()
					udiff %= 1 << (nBits - ntz)
					udiff <<= ntz
				}
				// udiff &= ^((1 << ntz) - 1)
				// consider udiff as zigzag encoded diff
				//prev += (-(int64(udiff) & 1)) ^ (int64(udiff) >> 1)
				prev -= (-(int64(udiff) & 1)) ^ (int64(udiff) >> 1)

				testInput = append(testInput, prev)
			}

			t.Run(fmt.Sprintf("nBits=%d, ntz=%d", nBits, ntz), func(t *testing.T) {
				resinput, input2 := intcomp.CompressDeltaBinPackInt64(testInput, nil)
				resinput2, resoutput2 := intcomp.UncompressDeltaBinPackInt64(input2, nil)

				if len(resinput) != 0 {
					t.Fatalf("expected all input consumed")
				}
				if len(resoutput2) != len(testInput) {
					t.Fatalf("expected same len")
				}
				if len(resinput2) != 0 {
					t.Fatalf("resinput2 should be empty")
				}
				for i, v := range testInput {
					if resoutput2[i] != v {
						intcomp.CompressDeltaBinPackInt64(testInput, nil)
						intcomp.UncompressDeltaBinPackInt64(input2, nil)

						t.Fatalf("resoutput2(%d) != input(%d)", resoutput2[i], v)
					}
				}
			})
		}
	}
}

func TestCompressDeltaBinPackFullUint64(t *testing.T) {
	rand.Seed(1) // nolint
	maxBits := 64
	for nBits := 0; nBits <= maxBits; nBits++ {
		for sign := 0; sign < 2; sign++ {
			var testInput []uint64
			prev := rand.Uint64()
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize64; i++ {
				var udiff uint64
				if nBits == maxBits {
					udiff = rand.Uint64()
				} else if nBits > 0 {
					udiff = rand.Uint64()
					udiff %= 1 << nBits
				}
				if sign == 1 {
					prev = uint64(prev) + udiff
				} else {
					// consider udiff as zigzag encoded diff
					prev += uint64((-(int64(udiff) & 1)) ^ (int64(udiff) >> 1))
				}
				testInput = append(testInput, prev)
			}

			t.Run(fmt.Sprintf("nBits=%d, sign=%d", nBits, sign), func(t *testing.T) {
				resinput, input2 := intcomp.CompressDeltaBinPackUint64(testInput, nil)
				resinput2, resoutput2 := intcomp.UncompressDeltaBinPackUint64(input2, nil)

				if len(resinput) != 0 {
					t.Fatalf("expected all input consumed")
				}
				if len(resoutput2) != len(testInput) {
					t.Fatalf("expected same len")
				}
				if len(resinput2) != 0 {
					t.Fatalf("resinput2 should be empty")
				}
				for i, v := range testInput {
					if resoutput2[i] != v {
						t.Fatalf("resoutput2(%d) != input(%d)", resoutput2[i], v)
					}
				}
			})
		}
	}
}

// BenchmarkCompressDeltaBinPackInt32 Benchmark compress one block: 128 x int32
func BenchmarkCompressDeltaBinPackInt32(b *testing.B) {
	rand.Seed(1) // nolint
	maxBits := 32
	for nBits := 0; nBits <= maxBits; nBits++ {
		for sign := 0; sign < 2; sign++ {
			var testInput []int32
			prev := int32(rand.Uint32())
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize32; i++ {
				var udiff uint32
				if nBits == maxBits {
					udiff = rand.Uint32()
				} else if nBits > 0 {
					udiff = rand.Uint32()
					udiff %= 1 << nBits
				}
				if sign == 1 {
					prev = int32(uint32(prev) + udiff)
				} else {
					// consider udiff as zigzag encoded diff
					prev += (-(int32(udiff) & 1)) ^ (int32(udiff) >> 1)
				}
				testInput = append(testInput, prev)
			}

			output := make([]uint32, 0, 1024)
			b.Run(fmt.Sprintf("nBits=%d, sign=%d", nBits, sign), func(b *testing.B) {
				b.SetBytes(int64(len(testInput) * 4))
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					intcomp.CompressDeltaBinPackInt32(testInput, output)
				}
			})
		}
	}
}

// BenchmarkDecompressDeltaBinPackInt32 Benchmark decompress one block: 128 x int32
func BenchmarkDecompressDeltaBinPackInt32(b *testing.B) {
	rand.Seed(1) // nolint
	maxBits := 32
	for nBits := 0; nBits <= maxBits; nBits++ {
		for sign := 0; sign < 2; sign++ {
			var testInput []int32
			prev := int32(rand.Uint32())
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize32; i++ {
				var udiff uint32
				if nBits == maxBits {
					udiff = rand.Uint32()
				} else if nBits > 0 {
					udiff = rand.Uint32()
					udiff %= 1 << nBits
				}
				if sign == 1 {
					prev = int32(uint32(prev) + udiff)
				} else {
					// consider udiff as zigzag encoded diff
					prev += (-(int32(udiff) & 1)) ^ (int32(udiff) >> 1)
				}
				testInput = append(testInput, prev)
			}

			resinput, input2 := intcomp.CompressDeltaBinPackInt32(testInput, nil)
			output := make([]int32, 0, 1024)

			b.Run(fmt.Sprintf("nBits=%d, sign=%d", nBits, sign), func(b *testing.B) {
				b.SetBytes(intcomp.BitPackingBlockSize32 * 4)
				b.ReportAllocs()
				b.ResetTimer()

				for n := 0; n < b.N; n++ {
					resinput2, resoutput2 := intcomp.UncompressDeltaBinPackInt32(input2, output)

					if len(resinput) != 0 {
						b.Fatalf("expected all input consumed")
					}
					if len(resoutput2) != len(testInput) {
						b.Fatalf("expected same len")
					}
					if len(resinput2) != 0 {
						b.Fatalf("resinput2 should be empty")
					}
					for i, v := range testInput {
						if resoutput2[i] != v {
							b.Fatalf("resoutput2(%d) != input(%d)", resoutput2[i], v)
						}
					}
				}
			})
		}
	}
}

// BenchmarkCompressDeltaBinPackUint32 Benchmark compress one block: 128 x uint32
func BenchmarkCompressDeltaBinPackUint32(b *testing.B) {
	rand.Seed(1) // nolint
	maxBits := 32
	for nBits := 0; nBits <= maxBits; nBits++ {
		for sign := 0; sign < 2; sign++ {
			var testInput []uint32
			prev := rand.Uint32()
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize32; i++ {
				var udiff uint32
				if nBits == maxBits {
					udiff = rand.Uint32()
				} else if nBits > 0 {
					udiff = rand.Uint32()
					udiff %= 1 << nBits
				}
				if sign == 1 {
					prev = uint32(prev) + udiff
				} else {
					// consider udiff as zigzag encoded diff
					prev += uint32((-(int32(udiff) & 1)) ^ (int32(udiff) >> 1))
				}
				testInput = append(testInput, prev)
			}

			output := make([]uint32, 0, 1024)
			b.Run(fmt.Sprintf("nBits=%d, sign=%d", nBits, sign), func(b *testing.B) {
				b.SetBytes(int64(len(testInput) * 4))
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					intcomp.CompressDeltaBinPackUint32(testInput, output)
				}
			})
		}
	}
}

// BenchmarkDecompressDeltaBinPackUint32 Benchmark decompress one block: 128 x uint32
func BenchmarkDecompressDeltaBinPackUint32(b *testing.B) {
	rand.Seed(1) // nolint
	maxBits := 32
	for nBits := 0; nBits <= maxBits; nBits++ {
		for sign := 0; sign < 2; sign++ {
			var testInput []uint32
			prev := rand.Uint32()
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize32; i++ {
				var udiff uint32
				if nBits == maxBits {
					udiff = rand.Uint32()
				} else if nBits > 0 {
					udiff = rand.Uint32()
					udiff %= 1 << nBits
				}
				if sign == 1 {
					prev = uint32(prev) + udiff
				} else {
					// consider udiff as zigzag encoded diff
					prev += uint32((-(int32(udiff) & 1)) ^ (int32(udiff) >> 1))
				}
				testInput = append(testInput, prev)
			}

			_, input2 := intcomp.CompressDeltaBinPackUint32(testInput, nil)

			output := make([]uint32, 0, 1024)

			b.Run(fmt.Sprintf("nBits=%d, sign=%d", nBits, sign), func(b *testing.B) {
				b.SetBytes(intcomp.BitPackingBlockSize32 * 4)
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					intcomp.UncompressDeltaBinPackUint32(input2, output)
				}
			})
		}
	}
}

func BenchmarkCompressDeltaBinPackInt64(b *testing.B) {
	rand.Seed(1) // nolint
	maxBits := 64
	for nBits := 0; nBits <= maxBits; nBits++ {
		for _, ntz := range []int{0, 1, 31, 62} {
			if ntz >= maxBits {
				continue
			}
			if nBits > 0 && ntz >= nBits {
				continue
			}
			var testInput []int64
			prev := int64(rand.Uint64())
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize64; i++ {
				var udiff uint64
				if nBits-ntz == maxBits {
					udiff = rand.Uint64()
				} else if nBits-ntz > 0 {
					udiff = rand.Uint64()
					udiff %= 1 << (nBits - ntz)
					udiff <<= ntz
				}
				// udiff as zigzag encoded diff
				prev -= (-(int64(udiff) & 1)) ^ (int64(udiff) >> 1)
				testInput = append(testInput, prev)
			}

			output := make([]uint64, 0, 1024)

			b.Run(fmt.Sprintf("nBits=%d, ntz=%d", nBits, ntz), func(b *testing.B) {
				b.SetBytes(int64(len(testInput) * 8))
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					intcomp.CompressDeltaBinPackInt64(testInput, output)
				}
			})

		}
	}
}

func BenchmarkDecompressDeltaBinPackInt64(b *testing.B) {
	rand.Seed(1) // nolint
	maxBits := 64
	for nBits := 0; nBits <= maxBits; nBits++ {
		for _, ntz := range []int{0, 1, 31, 62} {
			if ntz >= maxBits {
				continue
			}
			if nBits > 0 && ntz >= nBits {
				continue
			}
			var testInput []int64
			prev := int64(rand.Uint64())
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize64; i++ {
				var udiff uint64
				if nBits-ntz == maxBits {
					udiff = rand.Uint64()
				} else if nBits-ntz > 0 {
					udiff = rand.Uint64()
					udiff %= 1 << (nBits - ntz)
					udiff <<= ntz
				}
				// udiff as zigzag encoded diff
				prev -= (-(int64(udiff) & 1)) ^ (int64(udiff) >> 1)
				testInput = append(testInput, prev)
			}

			_, input2 := intcomp.CompressDeltaBinPackInt64(testInput, nil)
			output := make([]int64, 0, 1024)

			b.Run(fmt.Sprintf("nBits=%d, ntz=%d", nBits, ntz), func(b *testing.B) {
				b.SetBytes(intcomp.BitPackingBlockSize64 * 8)
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					intcomp.UncompressDeltaBinPackInt64(input2, output)
				}
			})

		}
	}
}

func BenchmarkCompressDeltaBinPackUint64(b *testing.B) {
	rand.Seed(1) // nolint
	maxBits := 64
	for nBits := 0; nBits <= maxBits; nBits++ {
		for sign := 0; sign < 2; sign++ {
			var testInput []uint64
			prev := rand.Uint64()
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize64; i++ {
				var udiff uint64
				if nBits == maxBits {
					udiff = rand.Uint64()
				} else if nBits > 0 {
					udiff = rand.Uint64()
					udiff %= 1 << nBits
				}
				if sign == 1 {
					prev = uint64(prev) + udiff
				} else {
					// consider udiff as zigzag encoded diff
					prev += uint64((-(int64(udiff) & 1)) ^ (int64(udiff) >> 1))
				}
				testInput = append(testInput, prev)
			}

			output := make([]uint64, 0, 1024)

			b.Run(fmt.Sprintf("nBits=%d, sign=%d", nBits, sign), func(b *testing.B) {
				b.SetBytes(int64(len(testInput) * 8))
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					intcomp.CompressDeltaBinPackUint64(testInput, output)
				}
			})
		}
	}
}

func BenchmarkDecompressDeltaBinPackUint64(b *testing.B) {
	rand.Seed(1) // nolint
	maxBits := 64
	for nBits := 0; nBits <= maxBits; nBits++ {
		for sign := 0; sign < 2; sign++ {
			var testInput []uint64
			prev := rand.Uint64()
			testInput = append(testInput, prev)
			for i := 1; i < intcomp.BitPackingBlockSize64; i++ {
				var udiff uint64
				if nBits == maxBits {
					udiff = rand.Uint64()
				} else if nBits > 0 {
					udiff = rand.Uint64()
					udiff %= 1 << nBits
				}
				if sign == 1 {
					prev = uint64(prev) + udiff
				} else {
					// consider udiff as zigzag encoded diff
					prev += uint64((-(int64(udiff) & 1)) ^ (int64(udiff) >> 1))
				}
				testInput = append(testInput, prev)
			}

			_, input2 := intcomp.CompressDeltaBinPackUint64(testInput, nil)

			output := make([]uint64, 0, 1024)

			b.Run(fmt.Sprintf("nBits=%d, sign=%d", nBits, sign), func(b *testing.B) {
				b.SetBytes(intcomp.BitPackingBlockSize64 * 8)
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					intcomp.UncompressDeltaBinPackUint64(input2, output)
				}
			})
		}
	}
}

func TestCompressDeltaNtz1(t *testing.T) {
	var testInput = []int64{16985, 17312, 17639, 17966, 18293, 18620, 18947, 19274, 19601, 19928, 20255, 20582, 20785, 21012, 21339, 21666, 21993, 22320, 22647, 22974, 23301, 23628, 23955, 24282, 24609, 24936, 25263, 25590, 25917, 26244, 26571, 26898, 27225, 27552, 27879, 28206, 28533, 28860, 29187, 29514, 29841, 30168, 30495, 30822, 31149, 31476, 31803, 32130, 32457, 32784, 33111, 33438, 33765, 34092, 34419, 34746, 35073, 35400, 35727, 36054, 36381, 36708, 37035, 37362, 16985, 17312, 17639, 17966, 18293, 18620, 18947, 19274, 19601, 19928, 20255, 20582, 20785, 21012, 21339, 21666, 21993, 22320, 22647, 22974, 23301, 23628, 23955, 24282, 24609, 24936, 25263, 25590, 25917, 26244, 26571, 26898, 27225, 27552, 27879, 28206, 28533, 28860, 29187, 29514, 29841, 30168, 30495, 30822, 31149, 31476, 31803, 32130, 32457, 32784, 33111, 33438, 33765, 34092, 34419, 34746, 35073, 35400, 35727, 36054, 36381, 36708, 37035, 37362, 16985, 17312, 17639, 17966, 18293, 18620, 18947, 19274, 19601, 19928, 20255, 20582, 20785, 21012, 21339, 21666, 21993, 22320, 22647, 22974, 23301, 23628, 23955, 24282, 24609, 24936, 25263, 25590, 25917, 26244, 26571, 26898, 27225, 27552, 27879, 28206, 28533, 28860, 29187, 29514, 29841, 30168, 30495, 30822, 31149, 31476, 31803, 32130, 32457, 32784, 33111, 33438, 33765, 34092, 34419, 34746, 35073, 35400, 35727, 36054, 36381, 36708, 37035, 37362, 16985, 17312, 17639, 17966, 18293, 18620, 18947, 19274, 19601, 19928, 20255, 20582, 20785, 21012, 21339, 21666, 21993, 22320, 22647, 22974, 23301, 23628, 23955, 24282, 24609, 24936, 25263, 25590, 25917, 26244, 26571, 26898, 27225, 27552, 27879, 28206, 28533, 28860, 29187, 29514, 29841, 30168, 30495, 30822, 31149, 31476, 31803, 32130, 32457, 32784, 33111, 33438, 33765, 34092, 34419, 34746, 35073, 35400, 35727, 36054, 36381, 36708, 37035, 37362}
	resinput, input2 := intcomp.CompressDeltaBinPack(testInput, make([]uint64, 0, 100))
	resinput2, resoutput2 := intcomp.UncompressDeltaBinPack[int64](input2, nil)

	if len(resoutput2) != len(testInput)-len(resinput) {
		t.Fatalf("expected same len")
	}
	if len(resinput2) != 0 {
		t.Fatalf("resinput2 should be empty")
	}
	for i, v := range testInput {
		if resoutput2[i] != v {
			t.Fatalf("expected %d, got %d", v, resoutput2[i])
		}
	}
}
