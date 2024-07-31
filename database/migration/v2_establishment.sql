CREATE TABLE IF NOT EXISTS establishment (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  document VARCHAR(14) NOT NULL,
  phone VARCHAR(15),
  qtde_motoclycles INTEGER NOT NULL DEFAULT 0,
  qtde_cars INTEGER NOT NULL DEFAULT 0,
  address_id INTEGER NOT NULL,
  FOREIGN KEY(address_id) REFERENCES addresses(id)
)
