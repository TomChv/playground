--
-- Extensions
--
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

--
-- Delivery
--

CREATE TYPE DeliveryManStatus as ENUM('Available', 'Working');

DROP TABLE IF EXISTS deliveryman;
CREATE TABLE deliveryman
(
    id     uuid           DEFAULT uuid_generate_v4() PRIMARY KEY,
    name   VARCHAR,
    status DeliveryManStatus DEFAULT 'Available'
);

--
-- Packets
--
CREATE TYPE PacketStatus as ENUM('Send', 'Traveling', 'Received');

DROP TABLE IF EXISTS packet;
CREATE TABLE packet
(
    id             uuid         DEFAULT uuid_generate_v4() PRIMARY KEY,
    owner          VARCHAR,
    receiver       VARCHAR,
    destination    VARCHAR,
    status         PacketStatus DEFAULT 'Send',
    deliveryman_id uuid,
    CONSTRAINT fk_deliveryman FOREIGN KEY (deliveryman_id) REFERENCES deliveryman (id)
);