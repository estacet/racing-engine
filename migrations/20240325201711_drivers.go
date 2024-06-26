package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.Exec(`
CREATE TABLE "public"."drivers" (
  id uuid PRIMARY KEY,
  name varchar(255) NOT NULL ,
  phone_number varchar(13) NOT NULL,
  age int NULL,
  weight int NULL,
  category driver_category NOT NULL
);
`)
		return err
	}, nil)
}
