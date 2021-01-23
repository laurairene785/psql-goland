package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConnection(t *testing.T) {
	c := getConnection()
	require.NotNil(t, c)
	r, err := c.Query("SELECT * FROM employee")
	require.NoError(t, err)
	cols, err := r.Columns()
	require.NoError(t, err)
	require.Greater(t, 0, len(cols))
	err = c.Close()
	require.NoError(t, err)
}
