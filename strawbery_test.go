package strawbery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StrawveryString(t *testing.T) {
	type args struct {
		kind string
		size string
	}
	tests := map[string]struct {
		args     args
		expected string
	}{
		"品種:あまおう、サイズ:Lのいちごから文字列表現「あまおう: L」を取得できる": {
			args:     args{kind: "あまおう", size: "L"},
			expected: "あまおう: L",
		},
		"品種:あまおう、サイズ:LLのいちごから文字列表現「あまおう: LL」を取得できる": {
			args:     args{kind: "あまおう", size: "LL"},
			expected: "あまおう: LL",
		},
		"品種:とちおとめ、サイズ:Lのいちごから文字列表現「とちおとめ: L」を取得できる": {
			args:     args{kind: "とちおとめ", size: "L"},
			expected: "とちおとめ: L",
		},
		"品種:もういっこ、サイズMのいちごから文字列表現「もういっこ: M」を取得できる": {
			args:     args{kind: "もういっこ", size: "M"},
			expected: "もういっこ: M",
		},
		"品種:もういっこ、サイズMのいちごから文字列表現「もういっこ: S」を取得できる": {
			args:     args{kind: "もういっこ", size: "S"},
			expected: "もういっこ: S",
		},
	}

	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			sut := Strawbery{
				kind: test.args.kind,
				size: test.args.size,
			}
			actual := sut.String()

			assert.Equal(t, test.expected, actual)
		})
	}

}

func Test_StrawveryWeight(t *testing.T) {
	t.Run("重さが0gの時エラー", func(t *testing.T) {
		_, err := New("あまおう", 0)
		assert.NotEmpty(t, err)
	})

	type args struct {
		kind   string
		weight uint
	}
	tests := map[string]struct {
		args     args
		expected string
	}{
		"重さが1gの時サイズはS": {
			args:     args{kind: "あまおう", weight: 1},
			expected: "S",
		},
		"重さが9gの時サイズはS": {
			args:     args{kind: "あまおう", weight: 9},
			expected: "S",
		},
		"重さが10gの時サイズはM": {
			args:     args{kind: "あまおう", weight: 10},
			expected: "M",
		},
		"重さが19gの時サイズはM": {
			args:     args{kind: "あまおう", weight: 19},
			expected: "M",
		},
		"重さが20gの時サイズはL": {
			args:     args{kind: "あまおう", weight: 20},
			expected: "L",
		},
		"重さが24gの時サイズはL": {
			args:     args{kind: "あまおう", weight: 24},
			expected: "L",
		},
		"重さが25gの時サイズはLL": {
			args:     args{kind: "あまおう", weight: 25},
			expected: "LL",
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			sut, _ := New(test.args.kind, test.args.weight)
			assert.Equal(t, test.expected, sut.Size())
		})
	}
}

//この辺でkindをちゃんとした名前にしたいなぁって気持ちが出てきた
func Test_あまおうとあまおうを比較した時にtrueを返す(t *testing.T) {
	berry1, _ := New("あまおう", 1)
	berry2, _ := New("あまおう", 2)
	actual := CompareKind(berry1, berry2)
	assert.Equal(t, true, actual)
}
