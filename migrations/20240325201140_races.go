package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.Exec(`
CREATE TYPE driver_category AS ENUM ('junior', 'amateur', 'pro');

CREATE TABLE "public"."races" (
  id uuid PRIMARY KEY,
  title varchar(255) NOT NULL,
  description varchar(2000) NOT NULL ,
  participants_count int NOT NULL ,
  registration_at timestamp without time zone NOT NULL,
  start_at timestamp without time zone NOT NULL
);
`)
		return err
	}, nil)
}
