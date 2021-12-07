--
-- Retrieve delivery & packets as JSON
--
CREATE OR REPLACE VIEW delivery_packets AS
SELECT d.id,
       d.name,
       d.status,
       (
           SELECT ARRAY_TO_JSON(ARRAY_AGG(ROW_TO_JSON(packetList.*)))
           FROM (
                    SELECT p.id, p.owner, p.receiver, p.destination, p.status
                    FROM packet p
                    WHERE p.deliveryman_id = d.id
                ) packetList
       ) AS packets
FROM deliveryman d;

--
-- Retrieve packets & delivery as JSON
--
CREATE OR REPLACE VIEW packets_delivery AS
SELECT p.id,
       p.receiver,
       p.owner,
       p.destination,
       p.status,
       (
           SELECT ROW_TO_JSON(deliveryman.*)
           FROM (
                    SELECT d.id, d.name, d.status
                    FROM deliveryman d
                    WHERE d.id = p.deliveryman_id
                ) deliveryman
       ) AS deliveryman
FROM packet p;