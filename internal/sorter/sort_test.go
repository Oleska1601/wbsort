package sorter_test

import (
	"reflect"
	"testing"

	"github.com/Oleska1601/wbsort/internal/parser"
	"github.com/Oleska1601/wbsort/internal/sorter"
)

// test flag -k
func TestSortByColumn(t *testing.T) {
	tests := []struct {
		name  string
		flags *parser.Flags
		lines []string
		want  []string
	}{
		{
			name:  "sort by first column",
			flags: &parser.Flags{FlagK: 1},
			lines: []string{"b\t2", "a\t1", "c\t3"},
			want:  []string{"a\t1", "b\t2", "c\t3"},
		},
		{
			name:  "sort by second column",
			flags: &parser.Flags{FlagK: 2},
			lines: []string{"z\t2", "y\t1", "x\t3"},
			want:  []string{"y\t1", "z\t2", "x\t3"},
		},
		{
			name:  "sort by non-existent column",
			flags: &parser.Flags{FlagK: 5},
			lines: []string{"banana", "apple", "cherry"},
			want:  []string{"banana", "apple", "cherry"}, // сохраняет порядок
		},
		{
			name:  "mixed column lengths",
			flags: &parser.Flags{FlagK: 2},
			lines: []string{"a\t3", "b", "c\t1"},
			want:  []string{"b", "c\t1", "a\t3"},
		},
		{
			name:  "empty lines with columns",
			flags: &parser.Flags{FlagK: 1},
			lines: []string{"", "b", "a"},
			want:  []string{"", "a", "b"},
		},
		{
			name:  "only tabs no text",
			flags: &parser.Flags{FlagK: 2},
			lines: []string{"\t", "\t\t", "\txyz"},
			want:  []string{"\t", "\t\t", "\txyz"},
		},
		{
			name:  "k=0 entire line",
			flags: &parser.Flags{FlagK: 0},
			lines: []string{"bbb", "aaa", "ccc"},
			want:  []string{"aaa", "bbb", "ccc"},
		},
	}

	runTests(t, tests)
}

// test flag -n
func TestSortNumeric(t *testing.T) {
	tests := []struct {
		name  string
		flags *parser.Flags
		lines []string
		want  []string
	}{
		{
			name:  "basic numeric sort",
			flags: &parser.Flags{FlagN: true},
			lines: []string{"10", "2", "1", "20"},
			want:  []string{"1", "2", "10", "20"},
		},
		{
			name:  "negative numbers",
			flags: &parser.Flags{FlagN: true},
			lines: []string{"-5", "10", "-10", "0"},
			want:  []string{"-10", "-5", "0", "10"},
		},
		{
			name:  "mixed numbers and text",
			flags: &parser.Flags{FlagN: true},
			lines: []string{"abc", "10", "2", "xyz"},
			want:  []string{"2", "10", "abc", "xyz"},
		},
		{
			name:  "floating point numbers",
			flags: &parser.Flags{FlagN: true},
			lines: []string{"1.5", "0.5", "2.5", "1.0"},
			want:  []string{"0.5", "1.0", "1.5", "2.5"},
		},
		{
			name:  "numeric with column",
			flags: &parser.Flags{FlagN: true, FlagK: 2},
			lines: []string{"a\t10", "b\t2", "c\t5"},
			want:  []string{"b\t2", "c\t5", "a\t10"},
		},
		{
			name:  "leading zeros",
			flags: &parser.Flags{FlagN: true},
			lines: []string{"001", "010", "100", "005"},
			want:  []string{"001", "005", "010", "100"},
		},
		{
			name:  "empty strings numeric",
			flags: &parser.Flags{FlagN: true},
			lines: []string{"", "5", "", "1"},
			want:  []string{"", "", "1", "5"},
		},
	}

	runTests(t, tests)
}

// test flag -r
func TestSortReverse(t *testing.T) {
	tests := []struct {
		name  string
		flags *parser.Flags
		lines []string
		want  []string
	}{
		{
			name:  "basic reverse sort",
			flags: &parser.Flags{FlagR: true},
			lines: []string{"a", "c", "b"},
			want:  []string{"c", "b", "a"},
		},
		{
			name:  "reverse numeric",
			flags: &parser.Flags{FlagR: true, FlagN: true},
			lines: []string{"1", "3", "2"},
			want:  []string{"3", "2", "1"},
		},
		{
			name:  "reverse with column",
			flags: &parser.Flags{FlagR: true, FlagK: 2},
			lines: []string{"x\t3", "y\t1", "z\t2"},
			want:  []string{"x\t3", "z\t2", "y\t1"},
		},
		{
			name:  "reverse empty lines",
			flags: &parser.Flags{FlagR: true},
			lines: []string{"", "b", "a", ""},
			want:  []string{"b", "a", "", ""},
		},
		{
			name:  "reverse months",
			flags: &parser.Flags{FlagR: true, FlagM: true},
			lines: []string{"Feb", "Jan", "Mar"},
			want:  []string{"Mar", "Feb", "Jan"},
		},
		{
			name:  "reverse with duplicates",
			flags: &parser.Flags{FlagR: true},
			lines: []string{"a", "c", "b", "c"},
			want:  []string{"c", "c", "b", "a"},
		},
		{
			name:  "reverse single element",
			flags: &parser.Flags{FlagR: true},
			lines: []string{"alone"},
			want:  []string{"alone"},
		},
	}

	runTests(t, tests)
}

// test flag -u
func TestSortUnique(t *testing.T) {
	tests := []struct {
		name  string
		flags *parser.Flags
		lines []string
		want  []string
	}{
		{
			name:  "basic unique",
			flags: &parser.Flags{FlagU: true},
			lines: []string{"b", "a", "b", "c", "a"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "unique with spaces",
			flags: &parser.Flags{FlagU: true},
			lines: []string{"a ", "a", " a", "a"},
			want:  []string{" a", "a", "a "}, // разные из-за пробелов
		},
		{
			name:  "unique numeric",
			flags: &parser.Flags{FlagU: true, FlagN: true},
			lines: []string{"2", "1", "2", "1", "3"},
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "unique with column",
			flags: &parser.Flags{FlagU: true, FlagK: 2},
			lines: []string{"x\t2", "y\t1", "z\t2", "w\t1"},
			want:  []string{"y\t1", "x\t2"},
		},
		{
			name:  "unique empty lines",
			flags: &parser.Flags{FlagU: true},
			lines: []string{"", "a", "", "a", ""},
			want:  []string{"", "a"},
		},
		{
			name:  "unique case sensitive",
			flags: &parser.Flags{FlagU: true},
			lines: []string{"A", "a", "A", "a"},
			want:  []string{"A", "a"}, // разные из-за регистра
		},
		{
			name:  "unique already sorted",
			flags: &parser.Flags{FlagU: true},
			lines: []string{"a", "b", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
	}

	runTests(t, tests)
}

// test flag -M
func TestSortMonth(t *testing.T) {
	tests := []struct {
		name  string
		flags *parser.Flags
		lines []string
		want  []string
	}{
		{
			name:  "basic month sort",
			flags: &parser.Flags{FlagM: true},
			lines: []string{"Feb", "Jan", "Mar"},
			want:  []string{"Jan", "Feb", "Mar"},
		},
		{
			name:  "month abbreviations",
			flags: &parser.Flags{FlagM: true},
			lines: []string{"Dec", "Nov", "Jan"},
			want:  []string{"Jan", "Nov", "Dec"},
		},
		{
			name:  "mixed month and non-month",
			flags: &parser.Flags{FlagM: true},
			lines: []string{"Feb", "abc", "Jan", "xyz"},
			want:  []string{"Jan", "Feb", "abc", "xyz"},
		},
		{
			name:  "month with column",
			flags: &parser.Flags{FlagM: true, FlagK: 2},
			lines: []string{"a\tFeb", "b\tJan", "c\tMar"},
			want:  []string{"b\tJan", "a\tFeb", "c\tMar"},
		},
		{
			name:  "case insensitive months",
			flags: &parser.Flags{FlagM: true},
			lines: []string{"feb", "JAN", "Mar"},
			want:  []string{"JAN", "feb", "Mar"},
		},
		{
			name:  "month in text",
			flags: &parser.Flags{FlagM: true},
			lines: []string{"Start Feb end", "Begin Jan end", "Mid Mar end"},
			want:  []string{"Begin Jan end", "Start Feb end", "Mid Mar end"},
		},
		{
			name:  "invalid months",
			flags: &parser.Flags{FlagM: true},
			lines: []string{"Jan", "Invalid", "Feb", "Unknown"},
			want:  []string{"Jan", "Feb", "Invalid", "Unknown"},
		},
	}

	runTests(t, tests)
}

// test flag -b
func TestSortIgnoreBlanks(t *testing.T) {
	tests := []struct {
		name  string
		flags *parser.Flags
		lines []string
		want  []string
	}{
		{
			name:  "trailing spaces ignored",
			flags: &parser.Flags{FlagB: true},
			lines: []string{"b  ", "a", "c\t"},
			want:  []string{"a", "b  ", "c\t"},
		},
		{
			name:  "leading spaces not ignored",
			flags: &parser.Flags{FlagB: true},
			lines: []string{" b", "a", " c"},
			want:  []string{" b", " c", "a"}, // leading spaces sorted
		},
		{
			name:  "spaces with column",
			flags: &parser.Flags{FlagB: true, FlagK: 2},
			lines: []string{"x\tb  ", "y\ta", "z\tc\t"},
			want:  []string{"y\ta", "x\tb  ", "z\tc\t"},
		},
		{
			name:  "mixed spaces and tabs",
			flags: &parser.Flags{FlagB: true},
			lines: []string{"b \t ", "a", "b   "},
			want:  []string{"a", "b \t ", "b   "},
		},
		{
			name:  "spaces with numeric",
			flags: &parser.Flags{FlagB: true, FlagN: true},
			lines: []string{"2  ", "1", "3\t"},
			want:  []string{"1", "2  ", "3\t"},
		},
		{
			name:  "spaces and empty",
			flags: &parser.Flags{FlagB: true},
			lines: []string{"", " b", "с ", "a "},
			want:  []string{"", " b", "a ", "с "},
		},
	}

	runTests(t, tests)
}

// test flag -c
func TestSortCheck(t *testing.T) {
	tests := []struct {
		name  string
		flags *parser.Flags
		lines []string
		want  []string
	}{
		{
			name:  "already sorted ascending",
			flags: &parser.Flags{FlagC: true},
			lines: []string{"apple", "banana", "cherry"},
			want:  nil,
		},
		{
			name:  "not sorted",
			flags: &parser.Flags{FlagC: true},
			lines: []string{"banana", "apple", "cherry"},
			want:  []string{"values are not sorted"},
		},
		{
			name:  "empty input",
			flags: &parser.Flags{FlagC: true},
			lines: []string{},
			want:  nil,
		},
		{
			name:  "single line",
			flags: &parser.Flags{FlagC: true},
			lines: []string{"single"},
			want:  nil,
		},
		{
			name:  "duplicate values",
			flags: &parser.Flags{FlagC: true, FlagU: true},
			lines: []string{"apple", "apple", "banana"},
			want:  nil, // дубликаты считаются отсортированными
		},
		{
			name:  "reverse sorted",
			flags: &parser.Flags{FlagC: true},
			lines: []string{"cherry", "banana", "apple"},
			want:  []string{"values are not sorted"},
		},
		{
			name:  "numeric check sorted",
			flags: &parser.Flags{FlagC: true, FlagN: true},
			lines: []string{"1", "2", "10"},
			want:  nil,
		},
		{
			name:  "numeric check not sorted",
			flags: &parser.Flags{FlagC: true, FlagN: true},
			lines: []string{"10", "2", "1"},
			want:  []string{"values are not sorted"},
		},
		{
			name:  "human numeric check sorted",
			flags: &parser.Flags{FlagC: true, FlagH: true},
			lines: []string{"500", "1K", "2K", "1M"},
			want:  nil,
		},
		{
			name:  "human numeric check not sorted",
			flags: &parser.Flags{FlagC: true, FlagH: true},
			lines: []string{"2K", "1K", "500"},
			want:  []string{"values are not sorted"},
		},
		{
			name:  "with column sorted",
			flags: &parser.Flags{FlagC: true, FlagK: 2},
			lines: []string{"a\tapple", "b\tbanana", "c\tcherry"},
			want:  nil,
		},
		{
			name:  "with column not sorted",
			flags: &parser.Flags{FlagC: true, FlagK: 2},
			lines: []string{"a\tbanana", "b\tapple", "c\tcherry"},
			want:  []string{"values are not sorted"},
		},
		{
			name:  "reverse check sorted",
			flags: &parser.Flags{FlagC: true, FlagR: true},
			lines: []string{"cherry", "banana", "apple"},
			want:  nil, // в обратном порядке это отсортировано
		},
		{
			name:  "reverse check not sorted",
			flags: &parser.Flags{FlagC: true, FlagR: true},
			lines: []string{"apple", "banana", "cherry"},
			want:  []string{"values are not sorted"},
		},
		{
			name:  "case insensitive sorted",
			flags: &parser.Flags{FlagC: true},
			lines: []string{"Apple", "Banana", "Cherry"},
			want:  nil,
		},
	}

	runTests(t, tests)
}

// test flag -h
func TestSortHuman(t *testing.T) {
	tests := []struct {
		name  string
		flags *parser.Flags
		lines []string
		want  []string
	}{
		{
			name:  "basic human sizes",
			flags: &parser.Flags{FlagH: true},
			lines: []string{"1K", "500", "2M", "1G"},
			want:  []string{"500", "1K", "2M", "1G"},
		},
		{
			name:  "decimal human sizes",
			flags: &parser.Flags{FlagH: true},
			lines: []string{"1.5K", "1K", "2K", "1.2K"},
			want:  []string{"1K", "1.2K", "1.5K", "2K"},
		},
		{
			name:  "mixed sizes and text",
			flags: &parser.Flags{FlagH: true},
			lines: []string{"2M", "abc", "1K", "1G"},
			want:  []string{"abc", "1K", "2M", "1G"},
		},
		{
			name:  "human with column",
			flags: &parser.Flags{FlagH: true, FlagK: 2},
			lines: []string{"x\t2K", "y\t1M", "z\t500"},
			want:  []string{"z\t500", "x\t2K", "y\t1M"},
		},
		{
			name:  "case insensitive suffixes",
			flags: &parser.Flags{FlagH: true},
			lines: []string{"1k", "2K", "1m", "2M"},
			want:  []string{"1k", "2K", "1m", "2M"},
		},
		{
			name:  "invalid human sizes",
			flags: &parser.Flags{FlagH: true},
			lines: []string{"1K", "2X", "3M", "invalid"},
			want:  []string{"2X", "invalid", "1K", "3M"},
		},
		{
			name:  "invalid human sizes 2",
			flags: &parser.Flags{FlagH: true},
			lines: []string{"1K", "cde", "2X", "3M", "abc"},
			want:  []string{"2X", "abc", "cde", "1K", "3M"},
		},
		{
			name:  "bytes without suffix",
			flags: &parser.Flags{FlagH: true},
			lines: []string{"1024", "1K", "512", "2K"},
			want:  []string{"512", "1024", "1K", "2K"},
		},
	}

	runTests(t, tests)
}

// test flag combinations
func TestSortFlagCombinations(t *testing.T) {
	tests := []struct {
		name  string
		flags *parser.Flags
		lines []string
		want  []string
	}{
		{
			name:  "numeric reverse",
			flags: &parser.Flags{FlagN: true, FlagR: true},
			lines: []string{"1", "3", "2"},
			want:  []string{"3", "2", "1"},
		},
		{
			name:  "unique reverse",
			flags: &parser.Flags{FlagU: true, FlagR: true},
			lines: []string{"c", "a", "b", "a"},
			want:  []string{"c", "b", "a"},
		},
		{
			name:  "column with unique",
			flags: &parser.Flags{FlagK: 2, FlagU: true},
			lines: []string{"x\t2", "y\t1", "z\t2", "w\t1"},
			want:  []string{"y\t1", "x\t2"},
		},
		{
			name:  "numeric with blanks",
			flags: &parser.Flags{FlagN: true, FlagB: true},
			lines: []string{"2  ", "1", "3\t"},
			want:  []string{"1", "2  ", "3\t"},
		},
		{
			name:  "month reverse",
			flags: &parser.Flags{FlagM: true, FlagR: true},
			lines: []string{"Jan", "Mar", "Feb"},
			want:  []string{"Mar", "Feb", "Jan"},
		},
		{
			name:  "human unique",
			flags: &parser.Flags{FlagH: true, FlagU: true},
			lines: []string{"1K", "2M", "1K", "500"},
			want:  []string{"500", "1K", "2M"},
		},
		{
			name:  "complex: column numeric reverse",
			flags: &parser.Flags{FlagK: 2, FlagN: true, FlagR: true},
			lines: []string{"a\t10", "b\t1", "c\t5"},
			want:  []string{"a\t10", "c\t5", "b\t1"},
		},
	}

	runTests(t, tests)
}

// help func
func runTests(t *testing.T, tests []struct {
	name  string
	flags *parser.Flags
	lines []string
	want  []string
}) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sorter.New(tt.flags, tt.lines)
			got := s.Sort()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
