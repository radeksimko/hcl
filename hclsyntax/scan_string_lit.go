// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclsyntax

// This file is generated from scan_string_lit.rl. DO NOT EDIT.
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
		switch cs {
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3

		}
	_ctr11:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st4
	_st4:
		if p == eof {
			goto _out4

		}
		p += 1
	st_case_4:
		if p == pe && p != eof {
			goto _out4

		}
		if p == eof {
			goto _st4

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr6

				}
			case 13:
				{
					goto _ctr7

				}
			case 36:
				{
					goto _ctr8

				}
			case 37:
				{
					goto _ctr9

				}

			}
			goto _st4

		}
	_ctr6:
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st5
	_ctr10:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st5
	_ctr12:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st5
	_st5:
		if p == eof {
			goto _out5

		}
		p += 1
	st_case_5:
		if p == pe && p != eof {
			goto _out5

		}
		if p == eof {
			goto _ctr10

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr12

				}
			case 13:
				{
					goto _ctr13

				}
			case 36:
				{
					goto _ctr14

				}
			case 37:
				{
					goto _ctr15

				}

			}
			goto _ctr11

		}
	_ctr7:
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st6
	_ctr16:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st6
	_ctr13:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st6
	_st6:
		if p == eof {
			goto _out6

		}
		p += 1
	st_case_6:
		if p == pe && p != eof {
			goto _out6

		}
		if p == eof {
			goto _ctr16

		} else {
			switch data[p] {
			case 10:
				{
					goto _st5

				}
			case 13:
				{
					goto _ctr13

				}
			case 36:
				{
					goto _ctr14

				}
			case 37:
				{
					goto _ctr15

				}

			}
			goto _ctr11

		}
	_ctr8:
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st7
	_ctr18:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st7
	_ctr14:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st7
	_st7:
		if p == eof {
			goto _out7

		}
		p += 1
	st_case_7:
		if p == pe && p != eof {
			goto _out7

		}
		if p == eof {
			goto _ctr18

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr12

				}
			case 13:
				{
					goto _ctr13

				}
			case 36:
				{
					goto _st8

				}
			case 37:
				{
					goto _ctr15

				}

			}
			goto _ctr11

		}
	_ctr20:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st8
	_st8:
		if p == eof {
			goto _out8

		}
		p += 1
	st_case_8:
		if p == pe && p != eof {
			goto _out8

		}
		if p == eof {
			goto _ctr20

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr12

				}
			case 13:
				{
					goto _ctr13

				}
			case 36:
				{
					goto _ctr14

				}
			case 37:
				{
					goto _ctr15

				}
			case 123:
				{
					goto _st5

				}

			}
			goto _ctr11

		}
	_ctr9:
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st9
	_ctr21:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st9
	_ctr15:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st9
	_st9:
		if p == eof {
			goto _out9

		}
		p += 1
	st_case_9:
		if p == pe && p != eof {
			goto _out9

		}
		if p == eof {
			goto _ctr21

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr12

				}
			case 13:
				{
					goto _ctr13

				}
			case 36:
				{
					goto _ctr14

				}
			case 37:
				{
					goto _st8

				}

			}
			goto _ctr11

		}
	_ctr29:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st10
	_st10:
		if p == eof {
			goto _out10

		}
		p += 1
	st_case_10:
		if p == pe && p != eof {
			goto _out10

		}
		if p == eof {
			goto _st10

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr23

				}
			case 13:
				{
					goto _ctr24

				}
			case 36:
				{
					goto _ctr25

				}
			case 37:
				{
					goto _ctr26

				}
			case 92:
				{
					goto _ctr27

				}

			}
			goto _st10

		}
	_ctr23:
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st11
	_ctr28:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st11
	_ctr30:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st11
	_st11:
		if p == eof {
			goto _out11

		}
		p += 1
	st_case_11:
		if p == pe && p != eof {
			goto _out11

		}
		if p == eof {
			goto _ctr28

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			goto _ctr29

		}
	_ctr24:
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st12
	_ctr35:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st12
	_ctr31:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st12
	_st12:
		if p == eof {
			goto _out12

		}
		p += 1
	st_case_12:
		if p == pe && p != eof {
			goto _out12

		}
		if p == eof {
			goto _ctr35

		} else {
			switch data[p] {
			case 10:
				{
					goto _st11

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			goto _ctr29

		}
	_ctr25:
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st13
	_ctr36:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st13
	_ctr32:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st13
	_st13:
		if p == eof {
			goto _out13

		}
		p += 1
	st_case_13:
		if p == pe && p != eof {
			goto _out13

		}
		if p == eof {
			goto _ctr36

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _st14

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			goto _ctr29

		}
	_ctr38:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st14
	_st14:
		if p == eof {
			goto _out14

		}
		p += 1
	st_case_14:
		if p == pe && p != eof {
			goto _out14

		}
		if p == eof {
			goto _ctr38

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}
			case 123:
				{
					goto _st11

				}

			}
			goto _ctr29

		}
	_ctr26:
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st15
	_ctr39:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st15
	_ctr33:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st15
	_st15:
		if p == eof {
			goto _out15

		}
		p += 1
	st_case_15:
		if p == pe && p != eof {
			goto _out15

		}
		if p == eof {
			goto _ctr39

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _st14

				}
			case 92:
				{
					goto _ctr34

				}

			}
			goto _ctr29

		}
	_ctr27:
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st16
	_ctr40:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st16
	_ctr34:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		{
			if te < p {
				ret = append(ret, data[te:p])
			}
			ts = p
		}
		goto _st16
	_st16:
		if p == eof {
			goto _out16

		}
		p += 1
	st_case_16:
		if p == pe && p != eof {
			goto _out16

		}
		if p == eof {
			goto _ctr40

		} else {
			switch data[p] {
			case 85:
				{
					goto _st17

				}
			case 117:
				{
					goto _st21

				}

			}
			if (data[p]) < 224 {
				if (data[p]) > 191 {
					{
						goto _st1

					}

				} else if (data[p]) >= 128 {
					goto _ctr29

				}

			} else if (data[p]) > 239 {
				if (data[p]) > 247 {
					{
						goto _ctr29

					}

				} else {
					goto _st3

				}

			} else {
				goto _st2

			}
			goto _st11

		}
	_ctr43:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st17
	_st17:
		if p == eof {
			goto _out17

		}
		p += 1
	st_case_17:
		if p == pe && p != eof {
			goto _out17

		}
		if p == eof {
			goto _ctr43

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			if (data[p]) < 65 {
				if 48 <= (data[p]) && (data[p]) <= 57 {
					goto _st18

				}

			} else if (data[p]) > 70 {
				if 97 <= (data[p]) && (data[p]) <= 102 {
					goto _st18

				}

			} else {
				goto _st18

			}
			goto _ctr29

		}
	_ctr45:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st18
	_st18:
		if p == eof {
			goto _out18

		}
		p += 1
	st_case_18:
		if p == pe && p != eof {
			goto _out18

		}
		if p == eof {
			goto _ctr45

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			if (data[p]) < 65 {
				if 48 <= (data[p]) && (data[p]) <= 57 {
					goto _st19

				}

			} else if (data[p]) > 70 {
				if 97 <= (data[p]) && (data[p]) <= 102 {
					goto _st19

				}

			} else {
				goto _st19

			}
			goto _ctr29

		}
	_ctr47:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st19
	_st19:
		if p == eof {
			goto _out19

		}
		p += 1
	st_case_19:
		if p == pe && p != eof {
			goto _out19

		}
		if p == eof {
			goto _ctr47

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			if (data[p]) < 65 {
				if 48 <= (data[p]) && (data[p]) <= 57 {
					goto _st20

				}

			} else if (data[p]) > 70 {
				if 97 <= (data[p]) && (data[p]) <= 102 {
					goto _st20

				}

			} else {
				goto _st20

			}
			goto _ctr29

		}
	_ctr49:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st20
	_st20:
		if p == eof {
			goto _out20

		}
		p += 1
	st_case_20:
		if p == pe && p != eof {
			goto _out20

		}
		if p == eof {
			goto _ctr49

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			if (data[p]) < 65 {
				if 48 <= (data[p]) && (data[p]) <= 57 {
					goto _st21

				}

			} else if (data[p]) > 70 {
				if 97 <= (data[p]) && (data[p]) <= 102 {
					goto _st21

				}

			} else {
				goto _st21

			}
			goto _ctr29

		}
	_ctr50:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st21
	_st21:
		if p == eof {
			goto _out21

		}
		p += 1
	st_case_21:
		if p == pe && p != eof {
			goto _out21

		}
		if p == eof {
			goto _ctr50

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			if (data[p]) < 65 {
				if 48 <= (data[p]) && (data[p]) <= 57 {
					goto _st22

				}

			} else if (data[p]) > 70 {
				if 97 <= (data[p]) && (data[p]) <= 102 {
					goto _st22

				}

			} else {
				goto _st22

			}
			goto _ctr29

		}
	_ctr52:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st22
	_st22:
		if p == eof {
			goto _out22

		}
		p += 1
	st_case_22:
		if p == pe && p != eof {
			goto _out22

		}
		if p == eof {
			goto _ctr52

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			if (data[p]) < 65 {
				if 48 <= (data[p]) && (data[p]) <= 57 {
					goto _st23

				}

			} else if (data[p]) > 70 {
				if 97 <= (data[p]) && (data[p]) <= 102 {
					goto _st23

				}

			} else {
				goto _st23

			}
			goto _ctr29

		}
	_ctr54:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st23
	_st23:
		if p == eof {
			goto _out23

		}
		p += 1
	st_case_23:
		if p == pe && p != eof {
			goto _out23

		}
		if p == eof {
			goto _ctr54

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			if (data[p]) < 65 {
				if 48 <= (data[p]) && (data[p]) <= 57 {
					goto _st24

				}

			} else if (data[p]) > 70 {
				if 97 <= (data[p]) && (data[p]) <= 102 {
					goto _st24

				}

			} else {
				goto _st24

			}
			goto _ctr29

		}
	_ctr56:
		{
			te = p
			ret = append(ret, data[ts:te])
		}
		goto _st24
	_st24:
		if p == eof {
			goto _out24

		}
		p += 1
	st_case_24:
		if p == pe && p != eof {
			goto _out24

		}
		if p == eof {
			goto _ctr56

		} else {
			switch data[p] {
			case 10:
				{
					goto _ctr30

				}
			case 13:
				{
					goto _ctr31

				}
			case 36:
				{
					goto _ctr32

				}
			case 37:
				{
					goto _ctr33

				}
			case 92:
				{
					goto _ctr34

				}

			}
			if (data[p]) < 65 {
				if 48 <= (data[p]) && (data[p]) <= 57 {
					goto _st11

				}

			} else if (data[p]) > 70 {
				if 97 <= (data[p]) && (data[p]) <= 102 {
					goto _st11

				}

			} else {
				goto _st11

			}
			goto _ctr29

		}
	_st1:
		if p == eof {
			goto _out1

		}
		p += 1
	st_case_1:
		if p == pe && p != eof {
			goto _out1

		}
		if p == eof {
			goto _st1

		} else {
			if 128 <= (data[p]) && (data[p]) <= 191 {
				goto _st11

			}
			goto _st0

		}
	_st0:
		if p == eof {
			goto _out0

		}
	st_case_0:
		goto _out0
	_st2:
		if p == eof {
			goto _out2

		}
		p += 1
	st_case_2:
		if p == pe && p != eof {
			goto _out2

		}
		if p == eof {
			goto _st2

		} else {
			if 128 <= (data[p]) && (data[p]) <= 191 {
				goto _st1

			}
			goto _st0

		}
	_st3:
		if p == eof {
			goto _out3

		}
		p += 1
	st_case_3:
		if p == pe && p != eof {
			goto _out3

		}
		if p == eof {
			goto _st3

		} else {
			if 128 <= (data[p]) && (data[p]) <= 191 {
				goto _st2

			}
			goto _st0

		}
	_out4:
		cs = 4
		goto _out
	_out5:
		cs = 5
		goto _out
	_out6:
		cs = 6
		goto _out
	_out7:
		cs = 7
		goto _out
	_out8:
		cs = 8
		goto _out
	_out9:
		cs = 9
		goto _out
	_out10:
		cs = 10
		goto _out
	_out11:
		cs = 11
		goto _out
	_out12:
		cs = 12
		goto _out
	_out13:
		cs = 13
		goto _out
	_out14:
		cs = 14
		goto _out
	_out15:
		cs = 15
		goto _out
	_out16:
		cs = 16
		goto _out
	_out17:
		cs = 17
		goto _out
	_out18:
		cs = 18
		goto _out
	_out19:
		cs = 19
		goto _out
	_out20:
		cs = 20
		goto _out
	_out21:
		cs = 21
		goto _out
	_out22:
		cs = 22
		goto _out
	_out23:
		cs = 23
		goto _out
	_out24:
		cs = 24
		goto _out
	_out1:
		cs = 1
		goto _out
	_out0:
		cs = 0
		goto _out
	_out2:
		cs = 2
		goto _out
	_out3:
		cs = 3
		goto _out
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
