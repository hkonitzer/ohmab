package privacy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestViewer(t *testing.T) {
	uv := UserViewer{
		Role: Admin,
	}
	require.Equal(t, uv.Admin(), true, "TestViewer: Admin() should be true")
	uv.SetUserID("testadmin")
	require.Equalf(t, uv.GetUserID(), "testadmin", "TestViewer: GetUserID() should be testadmin")
	err := uv.SetRole(2)
	require.NoError(t, err, "TestViewer: SetRole(2) should not return an error")
	require.Equal(t, uv.Admin(), true, "TestViewer: Admin() should be true")
	err = uv.SetRole(4)
	require.NoError(t, err, "TestViewer: SetRole(4) should not return an error")
	require.Equal(t, uv.Owner(), true, "TestViewer: Owner() should be true")
	err = uv.SetRole(8)
	require.NoError(t, err, "TestViewer: SetRole(8) should not return an error")
	require.Equal(t, uv.Employee(), true, "TestViewer: Employee() should be true")
	err = uv.SetRole(16)
	require.NoError(t, err, "TestViewer: SetRole(16) should not return an error")
	require.Equal(t, uv.IsEmpty(), false, "TestViewer: IsEmpty() should be false")
	err = uv.SetRole(0)
	require.Error(t, err, "TestViewer: SetRole(0) should return an error")
	require.Equal(t, uv.HasScopes(), false, "TestViewer: HasScopes() should be false")

}
