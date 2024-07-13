package output

import (
	"fmt"
	"testing"
)

func TestIndexToStr4c(t *testing.T) {
	const fName = "TestIndexToStr4c"

	// size over
	const expectNotOver uint = 2047
	const unexpectedSize uint = expectNotOver + 1
	_, overSizeErr := indexToStr4c(unexpectedSize)
	if overSizeErr == nil {
		t.Errorf("%s: unexpected index can over %d\n", fName, expectNotOver)
	}

	// expect result
	for i := range expectNotOver + 1 {
		input := uint(i)
		var expect string
		if input < 10 {
			expect = fmt.Sprintf("000%d", input)
		} else if input < 100 && input > 9 {
			expect = fmt.Sprintf("00%d", input)
		} else if input < 1000 && input > 99 {
			expect = fmt.Sprintf("0%d", input)
		} else if input < 2048 && input > 999 {
			expect = fmt.Sprintf("%d", input)
		} else {
			panic(fmt.Sprintf("%s: %s", fName, "expect result code else if error\n"))
		}

		result, err := indexToStr4c(input)
		if err != nil {
			t.Errorf("%s: %s", fName, err.Error())
		}
		if result != expect {
			t.Errorf("%s: input %d  expect %s, result %s\n", fName, input, expect, result)
		}
	}
}

func TestIndexListToString(t *testing.T) {
	const fName = "TestIndexListToString"
	type inputExpect struct {
		input  []uint
		expect string
	}

	// expect case
	inputExpects := []inputExpect{
		// real case get from seedsigner
		{
			input:  []uint{1924, 222, 235, 1743, 631, 1124, 378, 1770, 641, 1980, 1290, 1210},
			expect: "192402220235174306311124037817700641198012901210",
		},
		{
			input:  []uint{803, 154, 200, 626, 25, 1559, 70, 893, 1730, 788, 275, 2004},
			expect: "080301540200062600251559007008931730078802752004",
		},
		{
			input:  []uint{86, 750, 1025, 217, 1488, 23, 1715, 363, 517, 209, 1721, 1425},
			expect: "008607501025021714880023171503630517020917211425",
		},
		{
			input:  []uint{115, 1325, 1154, 127, 1190, 771, 415, 742, 1289, 1906, 2008, 870, 266, 1343, 1420, 2016, 1792, 614, 896, 1929, 300, 1524, 801, 643},
			expect: "011513251154012711900771041507421289190620080870026613431420201617920614089619290300152408010643",
		},
		{
			input:  []uint{114, 1655, 964, 1888, 73, 1119, 1572, 1887, 156, 610, 256, 1932, 1225, 1443, 573, 36, 1101, 1405, 1106, 1329, 2018, 1754, 1197, 1576},
			expect: "011416550964188800731119157218870156061002561932122514430573003611011405110613292018175411971576",
		},
		{
			input:  []uint{1662, 675, 203, 188, 1036, 1417, 658, 594, 1507, 1712, 1908, 1456, 1408, 1865, 1401, 744, 1273, 727, 1437, 994, 798, 1836, 1350, 1710},
			expect: "166206750203018810361417065805941507171219081456140818651401074412730727143709940798183613501710",
		},
		// unreal case just test function
		{
			input:  []uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expect: "000000000000000000000000000000000000000000000000",
		},
		{
			input:  []uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expect: "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
	}

	for _, inputExpect := range inputExpects {
		result, err := indexListToString(inputExpect.input)
		if err != nil {
			t.Errorf("%s: %s", fName, err.Error())
		}
		if inputExpect.expect != result {
			t.Errorf("%s: expect %s | result %s\n", fName, inputExpect.expect, result)
		}
	}

	// unexpected case
	inputUnexpected := []inputExpect{}
	//TODO:
	t.Errorf("TODO\n")

}

func TestPrintSeedSignerQTCode(t *testing.T) {
	// WARNING: use Manual Testing
	fmt.Println("output:TestPrintSeedSignerQRCode: pls manual test")
}
