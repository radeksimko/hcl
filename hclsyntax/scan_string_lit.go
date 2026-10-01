// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclsyntax

// This file is generated from scan_string_lit.rl. DO NOT EDIT.
var _hclstrtok_actions = []int8{0, 1, 0, 1, 1, 2, 1, 0, 0}
var _hclstrtok_key_offsets = []int16{0, 0, 2, 4, 6, 10, 14, 18, 22, 27, 31, 36, 41, 46, 51, 57, 62, 74, 85, 96, 107, 118, 129, 140, 151, 0}
var _hclstrtok_trans_keys = []byte{128, 191, 128, 191, 128, 191, 10, 13, 36, 37, 10, 13, 36, 37, 10, 13, 36, 37, 10, 13, 36, 37, 10, 13, 36, 37, 123, 10, 13, 36, 37, 10, 13, 36, 37, 92, 10, 13, 36, 37, 92, 10, 13, 36, 37, 92, 10, 13, 36, 37, 92, 10, 13, 36, 37, 92, 123, 10, 13, 36, 37, 92, 85, 117, 128, 191, 192, 223, 224, 239, 240, 247, 248, 255, 10, 13, 36, 37, 92, 48, 57, 65, 70, 97, 102, 10, 13, 36, 37, 92, 48, 57, 65, 70, 97, 102, 10, 13, 36, 37, 92, 48, 57, 65, 70, 97, 102, 10, 13, 36, 37, 92, 48, 57, 65, 70, 97, 102, 10, 13, 36, 37, 92, 48, 57, 65, 70, 97, 102, 10, 13, 36, 37, 92, 48, 57, 65, 70, 97, 102, 10, 13, 36, 37, 92, 48, 57, 65, 70, 97, 102, 10, 13, 36, 37, 92, 48, 57, 65, 70, 97, 102, 0}
var _hclstrtok_single_lengths = []int8{0, 0, 0, 0, 4, 4, 4, 4, 5, 4, 5, 5, 5, 5, 6, 5, 2, 5, 5, 5, 5, 5, 5, 5, 5, 0}
var _hclstrtok_range_lengths = []int8{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 3, 3, 3, 3, 3, 3, 3, 3, 0}
var _hclstrtok_index_offsets = []int16{0, 0, 2, 4, 6, 11, 16, 21, 26, 32, 37, 43, 49, 55, 61, 68, 74, 82, 91, 100, 109, 118, 127, 136, 145, 0}
var _hclstrtok_cond_targs = []int8{11, 0, 1, 0, 2, 0, 5, 6, 7, 9, 4, 5, 6, 7, 9, 4, 5, 6, 7, 9, 4, 5, 6, 8, 9, 4, 5, 6, 7, 9, 5, 4, 5, 6, 7, 8, 4, 11, 12, 13, 15, 16, 10, 11, 12, 13, 15, 16, 10, 11, 12, 13, 15, 16, 10, 11, 12, 14, 15, 16, 10, 11, 12, 13, 15, 16, 11, 10, 11, 12, 13, 14, 16, 10, 17, 21, 10, 1, 2, 3, 10, 11, 11, 12, 13, 15, 16, 18, 18, 18, 10, 11, 12, 13, 15, 16, 19, 19, 19, 10, 11, 12, 13, 15, 16, 20, 20, 20, 10, 11, 12, 13, 15, 16, 21, 21, 21, 10, 11, 12, 13, 15, 16, 22, 22, 22, 10, 11, 12, 13, 15, 16, 23, 23, 23, 10, 11, 12, 13, 15, 16, 24, 24, 24, 10, 11, 12, 13, 15, 16, 11, 11, 11, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 0}
var _hclstrtok_cond_actions = []int8{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 5, 5, 5, 5, 3, 0, 5, 5, 5, 3, 5, 5, 0, 5, 3, 5, 5, 5, 5, 0, 3, 5, 5, 5, 0, 3, 1, 1, 1, 1, 1, 0, 5, 5, 5, 5, 5, 3, 0, 5, 5, 5, 5, 3, 5, 5, 0, 5, 5, 3, 5, 5, 5, 5, 5, 0, 3, 5, 5, 5, 0, 5, 3, 0, 0, 3, 0, 0, 0, 3, 0, 5, 5, 5, 5, 5, 0, 0, 0, 3, 5, 5, 5, 5, 5, 0, 0, 0, 3, 5, 5, 5, 5, 5, 0, 0, 0, 3, 5, 5, 5, 5, 5, 0, 0, 0, 3, 5, 5, 5, 5, 5, 0, 0, 0, 3, 5, 5, 5, 5, 5, 0, 0, 0, 3, 5, 5, 5, 5, 5, 0, 0, 0, 3, 5, 5, 5, 5, 5, 0, 0, 0, 3, 0, 0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 0}
var _hclstrtok_eof_trans = []int16{155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 0}
var hclstrtok_start int = 4
var _ = hclstrtok_start
var hclstrtok_first_final int = 4
var _ = hclstrtok_first_final
var hclstrtok_error int = 0
var _ = hclstrtok_error
var hclstrtok_en_quoted int = 10
var _ = hclstrtok_en_quoted
var hclstrtok_en_unquoted int = 4
var _ = hclstrtok_en_unquoted

func scanStringLit(data []byte, quoted bool) [][]byte {
	var ret [][]byte

	// Ragel state
	p := 0          // "Pointer" into data
	pe := len(data) // End-of-data "pointer"
	ts := 0
	te := 0
	eof := pe

	var cs int // current state
	switch {
	case quoted:
		cs = hclstrtok_en_quoted
	default:
		cs = hclstrtok_en_unquoted
	}

	// Make Go compiler happy
	_ = ts
	_ = eof

	/*token := func () {
	ret = append(ret, data[ts:te])
	}*/

	{

	}
	{
		var _klen int
		var _trans uint = 0
		var _keys int
		var _acts int
		var _nacts uint
	_resume:
		{

		}
		if p == pe && p != eof {
			goto _out

		}
		if p == eof {
			if _hclstrtok_eof_trans[cs] > 0 {
				_trans = uint(_hclstrtok_eof_trans[cs]) - 1

			}

		} else {
			_keys = int(_hclstrtok_key_offsets[cs])

			_trans = uint(_hclstrtok_index_offsets[cs])
			_klen = int(_hclstrtok_single_lengths[cs])
			if _klen > 0 {
				var _lower int = _keys
				var _upper int = _keys + _klen - 1
				var _mid int
				for {
					if _upper < _lower {
						_keys += _klen
						_trans += uint(_klen)
						break

					}
					_mid = _lower + ((_upper - _lower) >> 1)
					if (data[p]) < _hclstrtok_trans_keys[_mid] {
						_upper = _mid - 1

					} else if (data[p]) > _hclstrtok_trans_keys[_mid] {
						_lower = _mid + 1

					} else {
						_trans += uint((_mid - _keys))
						goto _match

					}

				}

			}
			_klen = int(_hclstrtok_range_lengths[cs])
			if _klen > 0 {
				var _lower int = _keys
				var _upper int = _keys + (_klen << 1) - 2
				var _mid int
				for {
					if _upper < _lower {
						_trans += uint(_klen)
						break

					}
					_mid = _lower + (((_upper - _lower) >> 1) & ^1)
					if (data[p]) < _hclstrtok_trans_keys[_mid] {
						_upper = _mid - 2

					} else if (data[p]) > _hclstrtok_trans_keys[_mid+1] {
						_lower = _mid + 2

					} else {
						_trans += uint(((_mid - _keys) >> 1))
						break

					}

				}

			}
		_match:
			{

			}

		}
		cs = int(_hclstrtok_cond_targs[_trans])
		if _hclstrtok_cond_actions[_trans] != 0 {
			_acts = int(_hclstrtok_cond_actions[_trans])

			_nacts = uint(_hclstrtok_actions[_acts])
			_acts += 1
			for _nacts > 0 {
				switch _hclstrtok_actions[_acts] {
				case 0:
					{
						if te < p {
							ret = append(ret, data[te:p])
						}
						ts = p
					}

				case 1:
					{
						te = p
						ret = append(ret, data[ts:te])
					}

				}
				_nacts -= 1
				_acts += 1

			}

		}
		if p == eof {
			if cs >= 4 {
				goto _out

			}

		} else {
			if cs != 0 {
				p += 1
				goto _resume

			}

		}
	_out:
		{

		}

	}

	if te < p {
		// Collect any leftover literal characters at the end of the input
		ret = append(ret, data[te:p])
	}

	// If we fall out here without being in a final state then we've
	// encountered something that the scanner can't match, which should
	// be impossible (the scanner matches all bytes _somehow_) but we'll
	// tolerate it and let the caller deal with it.
	if cs < hclstrtok_first_final {
		ret = append(ret, data[p:len(data)])
	}

	return ret
}
