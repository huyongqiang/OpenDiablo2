// Code generated by "stringer -linecomment -type ObjectAnimationMode"; DO NOT EDIT.

package d2enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AnimationModeObjectNeutral-0]
	_ = x[AnimationModeObjectOperating-1]
	_ = x[AnimationModeObjectOpened-2]
	_ = x[AnimationModeObjectSpecial1-3]
	_ = x[AnimationModeObjectSpecial2-4]
	_ = x[AnimationModeObjectSpecial3-5]
	_ = x[AnimationModeObjectSpecial4-6]
	_ = x[AnimationModeObjectSpecial5-7]
}

const _ObjectAnimationMode_name = "NUOPONS1S2S3S4S5"

var _ObjectAnimationMode_index = [...]uint8{0, 2, 4, 6, 8, 10, 12, 14, 16}

func (i ObjectAnimationMode) String() string {
	if i < 0 || i >= ObjectAnimationMode(len(_ObjectAnimationMode_index)-1) {
		return "ObjectAnimationMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ObjectAnimationMode_name[_ObjectAnimationMode_index[i]:_ObjectAnimationMode_index[i+1]]
}
