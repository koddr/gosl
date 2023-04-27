package gosl

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkMarshal_StructField_4(b *testing.B) {
	type user struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"-"`
	}
	
	u := &user{}
	
	for i := 0; i < b.N; i++ {
		Marshal[user](u)
	}
}

func BenchmarkMarshal_StructField_16(b *testing.B) {
	type user struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"-"`
		Attr1    string `json:"attr_1"`
		Attr2    string `json:"attr_2"`
		Attr3    string `json:"attr_3"`
		Attr4    string `json:"attr_4"`
		Attr5    string `json:"attr_5"`
		Attr6    string `json:"attr_6"`
		Attr7    string `json:"attr_7"`
		Attr8    string `json:"attr_8"`
		Rel1     string `json:"rel_1"`
		Rel2     string `json:"rel_2"`
		Rel3     string `json:"rel_3"`
		Rel4     string `json:"rel_4"`
	}
	
	u := &user{}
	
	for i := 0; i < b.N; i++ {
		Marshal[user](u)
	}
}

func BenchmarkUnmarshal_StructField_4(b *testing.B) {
	type user struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"-"`
	}
	
	u := &user{}
	d := []byte(`{"id":1,"name":"Viktor","email":"my@mail.com"}`)
	
	for i := 0; i < b.N; i++ {
		Unmarshal[user](d, u)
	}
}

func BenchmarkUnmarshal_StructField_16(b *testing.B) {
	type user struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Password string `json:"-"`
		Email    string `json:"email"`
		Attr1    string `json:"attr_1"`
		Attr2    string `json:"attr_2"`
		Attr3    string `json:"attr_3"`
		Attr4    string `json:"attr_4"`
		Attr5    string `json:"attr_5"`
		Attr6    string `json:"attr_6"`
		Attr7    string `json:"attr_7"`
		Attr8    string `json:"attr_8"`
		Rel1     string `json:"rel_1"`
		Rel2     string `json:"rel_2"`
		Rel3     string `json:"rel_3"`
		Rel4     string `json:"rel_4"`
	}
	
	u := &user{}
	d := []byte(`{"id":1,"name":"Viktor","email":"my@mail.com","attr_1":"one","attr_2":"two","attr_3":"three","attr_4":"four","attr_5":"five","attr_6":"six","attr_7":"seven","attr_8":"eight","rel_1":"one","rel_2":"two","rel_3":"three","rel_4":"four",}`)
	
	for i := 0; i < b.N; i++ {
		Unmarshal[user](d, u)
	}
}

func TestMarshal(t *testing.T) {
	type user struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Password string `json:"-"`
	}
	
	u := &user{ID: 1, Name: "Viktor"}
	
	json, err := Marshal(u)
	require.NoError(t, err)
	assert.EqualValues(t, []byte(`{"id":1,"name":"Viktor"}`), json)
	
	g := GenericUtility[user, any]{} // tests for method
	
	json, err = g.Marshal(u)
	require.NoError(t, err)
	assert.EqualValues(t, []byte(`{"id":1,"name":"Viktor"}`), json)
}

func TestUnmarshal(t *testing.T) {
	type user struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Password string `json:"-"`
	}
	
	u := &user{ID: 1, Name: "Viktor"}
	data := []byte(`{"id":1,"name":"Viktor"}`)
	
	_, err := Unmarshal(nil, u)
	require.Error(t, err)
	
	json, err := Unmarshal(data, u)
	require.NoError(t, err)
	assert.EqualValues(t, u, json)
	
	g := GenericUtility[user, any]{} // tests for method
	
	_, err = g.Unmarshal(nil, u)
	require.Error(t, err)
	
	json, err = g.Unmarshal(data, u)
	require.NoError(t, err)
	assert.EqualValues(t, u, json)
}
