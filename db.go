package main

import (
	"context"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
)

func addFigureToDb(figure FigureData) {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:@localhost:5432/figures")
	if err != nil {
		log.Fatal(err)
	}

    _, err = conn.Exec(context.Background(), `
        INSERT INTO figure(scale, brand, origin_url)
        VALUES ($1, $2, $3)
        `, getScale(figure.TableData[4]), figure.Brand, figure.URL)

    if err != nil {
        log.Fatal(err)
    }

    conn.Close(context.Background())
}

func getScale(size string) string {
    if (strings.Contains(size, "スケール")) {
        return strings.TrimSpace(size[0:3])
    } else {
        return ""
    }
}

/*

CREATE TABLE figure
(
  id SERIAL PRIMARY KEY,
  scale VARCHAR(4),
  brand VARCHAR(255),
  origin_url VARCHAR(255) NOT NULL
);
 */
