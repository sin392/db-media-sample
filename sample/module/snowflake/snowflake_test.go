package snowflake_test

import (
	"context"
	"testing"

	"github.com/sin392/db-media-sample/sample/module/snowflake"
	"github.com/stretchr/testify/require"
)

func TestSnowflakeGenerator(t *testing.T) {
	t.Parallel()

	g, err := snowflake.NewSnowflakeIDGenerator(1)
	require.NoError(t, err)

	type args struct {
		name string
	}

	tests := []args{
		{
			name: "snowflakeIDの生成と取得に成功",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := g.Generate()
			require.NotEmpty(t, got)

			ctx := snowflake.SetSnowflakeID(context.Background(), got)
			require.Equal(t, got, snowflake.GetSnowflakeID(ctx))
		})
	}
}
