CREATE OR REPLACE FUNCTION update_ratings_on_complete()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.completed THEN
    -- Call a procedure to send notification or update ratings directly
    PERFORM update_ratings(NEW.winner_id, NEW.loser_id);
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER update_ratings_after_update
AFTER UPDATE ON matches
FOR EACH ROW
EXECUTE PROCEDURE update_ratings_on_complete();