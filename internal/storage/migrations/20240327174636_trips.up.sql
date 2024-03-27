CREATE TABLE trips (
	id SERIAL PRIMARY KEY,
	vendor_id INT NOT NULL,
	pickup TIMESTAMP NOT NULL,
	dropoff TIMESTAMP NOT NULL,
	passengers INT NOT NULL,
	duration INT NOT NULL
)
